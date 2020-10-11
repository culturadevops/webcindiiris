package controllers

import (
	commons "inkafarma/webcindi/common"
	"inkafarma/webcindi/model"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type RunbooksController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *RunbooksController) GetBasedatosBy(env uint) mvc.View {
	admin_user, _ := c.Session.Get("admin_user").(map[string]interface{})
	user_id, _ := admin_user["id"].(uint)
	muserenv := model.UserEnvModel{}
	mbasedato := model.BasedatoModel{}
	envs := muserenv.GetByUserId(user_id)
	var list []model.BasedatoModel
	var total, totalPages int
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	for _, item := range envs {
		if item.Env_id == env {
			list, total, totalPages = mbasedato.ListByEnvAndPage(env, page)
		}
	}

	return mvc.View{
		Name: "runbook/basedato.html",
		Data: iris.Map{
			"Title":    "Runbook de base de datos",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}
