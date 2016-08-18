package controllers

import "ggcms/models"

type ToolsController struct {
	BaseController
}

// @router /clearcacth [get]
func (c *ToolsController) CacheClearAll() {
	msg := models.Message{0, "成功", nil}
	if err := ggcacth.ClearAll(); err != nil {
		msg = models.Message{1, err.Error(), err}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @router /getpowers [get]
func (c *ToolsController) GetAdminPowers() {
	msg := models.Message{0, "成功", nil}
	if utype, ok := c.Ctx.Input.Session("utype").(string); ok {
		if utype == "0" {
			msg.Data = map[string]int{"allallow": 1}
		} else {
			pwlist := c.cacheman.CacheUserPowers(utype)
			msg.Data = pwlist
		}
	} else {
		msg = models.Message{1, "获取权限信息失败，或未登录", map[string]int{"allallow": 0}}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
