package controllers

import (
	"encoding/json"
	"ggcms/models"
	"strings"
)

// oprations for GgcmsAdminPowers
type GgcmsAdminPowersController struct {
	BaseController
}

func (c *GgcmsAdminPowersController) URLMapping() {
	c.Mapping("Edit", c.Edit)
}

// @Title Update
// @Description update the GgcmsAdminPowers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsAdminPowers	true		"body for GgcmsAdminPowers content"
// @Success 200 {object} models.GgcmsAdminPowers
// @Failure 403 :id is not int
// @router /edit/:id [post]
func (c *GgcmsAdminPowersController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	var v []int
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateGgcmsAdminPowersById(&v, idStr); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.cacheman.ClearByKey(cnUserPowers + "_" + idStr)
			msg = models.Message{0, "成功", v}
			c.Data["json"] = msg
		} else {
			msg.Msg = err.Error()
		}
	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

func GetAllGgcmsAdminPowers(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string
	// fields: col1,col2,entity.col3
	if v := strfields; v != "" {
		fields = strings.Split(v, ",")
	}

	// sortby: col1,col2
	if v := strsortby; v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := strorder; v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if query, err = getQueryList(strquery); err != nil {
		return nil, 0, err
	}
	offset := pagesize * (pagenum - 1)
	return models.GetAllGgcmsAdminPowers(query, fields, sortby, order, offset, pagesize, c)
}
