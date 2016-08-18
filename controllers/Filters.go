package controllers

import (
	"ggcms/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func AdminAuth(ctx *context.Context) {
	url := ctx.Request.RequestURI
	if !strings.Contains(url, "ggcms_admin/login") {
		var utype string
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			msg := models.Message{1, "没有登录或登录已超时", nil}
			ctx.Output.JSON(msg, true, false)
			return
		}
		utype, ok = ctx.Input.Session("utype").(string)
		//登录信息不正确
		if !ok {
			msg := models.Message{1, "没有登录或登录已超时", nil}
			ctx.Output.JSON(msg, true, false)
			return
		}
		//为0时是管理拥有超级权限
		if utype != "0" {
			//没有权限跳转
			if !adminAuthPower(ctx, true) {
				msg := models.Message{1, "无权使用本功能", nil}
				ctx.Output.JSON(msg, true, false)
				return
			}
		}

	}
	beego.Debug(ctx.Input.Scheme(), ctx.Input.URI(), ctx.Input.URL(), ctx.Input.Domain(), ctx.Input.Method(), ctx.Input.IsAjax(), ctx.Request.Host)
}

func AdminLogin(ctx *context.Context) {
	url := ctx.Input.URL()
	if !strings.Contains(url, "login.html") {
		var ok bool
		var utype string
		_, ok = ctx.Input.Session("uid").(int)
		//未登录
		if !ok {
			ctx.Redirect(302, "login.html")
			return
		}
		utype, ok = ctx.Input.Session("utype").(string)
		//登录信息不正确
		if !ok {
			ctx.Redirect(302, "login.html")
			return
		}
		//为0时是管理拥有超级权限
		if utype != "0" {
			//没有权限跳转
			if !adminAuthPower(ctx, false) {
				ctx.Redirect(302, "error.html?errcode=501")
				return
			}
		}
	}
}
