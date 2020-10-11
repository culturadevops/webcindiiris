package controllers

import (
	"html"
	commons "inkafarma/webcindi/common"
	"inkafarma/webcindi/model"
	"inkafarma/webcindi/service"
	"strconv"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type BastionController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *BastionController) Get() mvc.View {

	admin_user, _ := c.Session.Get("admin_user").(map[string]interface{})
	admin_id, _ := admin_user["id"].(uint)
	data := model.CredecentialBastionModel{}
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	list, total, totalPages := data.List(admin_id, page)
	return mvc.View{
		Name: "bastion/list.html",
		Data: iris.Map{
			"Title":    "Credenciales de bastions",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}
func (c *BastionController) GetUpdatePasswordBy(id uint) mvc.View {
	var model model.CredecentialBastionModel
	item, _ := model.GetById(id)

	//if err != nil {
	//	return commons.MvcError(err.Error(), c.Ctx)
	//}
	return mvc.View{
		Name: "bastion/updatePassword.html",
		Data: iris.Map{
			"Title": "Nuevo password",
			"Id":    id,
			"User":  item.Account,
			"Env":   item.Env,
			//	"Account": adminInfo.Account,
		},
	}
}
func (c *BastionController) PostUpdatePassword() {
	id := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("id")))
	user := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("user")))
	env := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("env")))
	password := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("password")))
	Repassword := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("Repassword")))
	int_admin_id, _ := strconv.Atoi(id)
	var model model.CredecentialBastionModel

	uienv, _ := strconv.ParseUint(env, 10, 32)
	service.BastionUpdatePass(uint(uienv), user, password)
	if err := model.PasswodUpdate(uint(int_admin_id), password, Repassword); err == nil {
		c.Ctx.Redirect("/bastion")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
