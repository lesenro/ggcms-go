package controllers

import (
	"ggcms/models"
	"strings"

	"github.com/astaxie/beego/context"
)

func AdminAuth(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	url := ctx.Request.RequestURI

	if !ok && !strings.Contains(url, "ggcms_admin/login") {
		msg := models.Message{1, "没有登录", nil}
		ctx.Output.JSON(msg, true, false)
	}
}

func AdminLogin(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	url := ctx.Request.RequestURI

	if !ok && !strings.Contains(url, "login.html") {
		ctx.Redirect(302, "login.html")
	}
}
