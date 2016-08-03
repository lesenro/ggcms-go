package controllers

import (
	"github.com/astaxie/beego/context"
)

func AdminAuth(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok {
		m1 := make(map[string]string)
		m1["type"]="success";
		ctx.Output.JSON(m1,true,false)
	}
}

func AdminLogin(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "login" {
		ctx.Redirect(302,"login");
	}
}