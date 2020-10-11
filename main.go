package main

import (
	commons "inkafarma/webcindi/common"
	"inkafarma/webcindi/libs"
	"inkafarma/webcindi/model"
	"inkafarma/webcindi/route"
	"log"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	config "github.com/spf13/viper"
)

func init() {
	config.AddConfigPath("./configs")
	config.SetConfigName("mysql")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbConfig := libs.DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
	}
	libs.DB = dbConfig.InitDB()
	if config.GetBool("default.sql_log") {
		libs.DB.LogMode(true)
	}
}

func main() {
	app := iris.New()
	config.SetConfigName("app")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer el archivo de configuración, %s", err)
	}
	tmpl := iris.HTML("./views", ".html").Layout(config.GetString("site.DefaultLayout"))
	if config.GetBool("site.APPDebug") == true {
		app.Logger().SetLevel("debug") //设置debug
		tmpl.Reload(true)
	}

	tmpl.AddFunc("TimeToDate", libs.TimeToDate)
	tmpl.AddFunc("strToHtml", libs.StrToHtml)

	app.RegisterView(tmpl)
	app.Favicon("./favicon.ico")
	app.Use(iris.Gzip)

	// (opcional) agrega dos controladores integrados
	// Puede recuperarse de cualquier pánico relacionado con http
	// Grabe la solicitud en la terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	app.StaticWeb("/public", "./public")   //设置静态文件目录
	app.StaticWeb("/uploads", "./uploads") //设置静态文件目录

	//设置公共页面输出
	app.Use(func(ctx iris.Context) {
		if auth := commons.SessManager.Start(ctx).Get("admin_user"); auth != nil {
			admin_user, _ := auth.(map[string]interface{})
			var admin_model model.Admin
			admin_id, _ := admin_user["id"].(uint)
			adminInfo, _ := admin_model.AdminInfo(admin_id)
			if adminInfo.Headico == "" {
				adminInfo.Headico = "/public/adminlit/dist/img/user2-160x160.jpg"
			}
			ctx.ViewData("adminInfo", adminInfo)
			var modeluserenv model.UserEnvModel
			envs := modeluserenv.GetByUserId(admin_id)

			ctx.ViewData("ci01", "no")
			ctx.ViewData("ci02", "no")
			ctx.ViewData("qa01", "no")
			ctx.ViewData("qa02", "no")
			ctx.ViewData("uat", "no")
			ctx.ViewData("uat02", "no")
			ctx.ViewData("prd", "no")
			for _, envs := range envs {
				if envs.Env_id == 1 {
					ctx.ViewData("ci01", "si")
				}
				if envs.Env_id == 2 {
					ctx.ViewData("ci02", "si")
				}
				if envs.Env_id == 3 {
					ctx.ViewData("qa01", "si")
				}
				if envs.Env_id == 4 {
					ctx.ViewData("qa02", "si")
				}
				if envs.Env_id == 5 {
					ctx.ViewData("uat", "si")
				}
				if envs.Env_id == 6 {
					ctx.ViewData("uat02", "si")
				}
				if envs.Env_id == 7 {
					ctx.ViewData("prd", "si")
				}
			}

		}
		ctx.ViewData("Title", config.GetString("site.DefaultTitle"))
		ctx.ViewData("Version", config.GetString("site.version"))
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.Next()
	})

	//设置错误模版
	app.OnAnyErrorCode(func(ctx iris.Context) {
		_, err := ctx.HTML("<center>Lo siento! Error de página actual, código de error::" + strconv.Itoa(ctx.GetStatusCode()) + "</center>")
		if err != nil {
			log.Fatalf("内部错误,错误代码 %s", err)
		}
	})

	route.Routes(app)

	//应用配置文件
	app.Configure(iris.WithConfiguration(iris.YAML("./configs/iris.yml")))

	//Run
	www := app.Party("www.")
	{
		currentRoutes := app.GetRoutes()
		for _, r := range currentRoutes {
			www.Handle(r.Method, r.Tmpl().Src, r.Handlers...)
		}
	}
	err := app.Run(iris.Addr(config.GetString("server.domain") + ":" + config.GetString("server.port")))
	if err != nil {
		log.Fatalf("El servicio no pudo iniciarse, código de error %s", err)
	}
}
