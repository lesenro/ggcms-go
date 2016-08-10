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
