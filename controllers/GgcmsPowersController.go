package controllers

import (
	"encoding/json"
	"ggcms/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for GgcmsPowers
type GgcmsPowersController struct {
	beego.Controller
}

func (c *GgcmsPowersController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description create GgcmsPowers
// @Param	body		body 	models.GgcmsPowers	true		"body for GgcmsPowers content"
// @Success 201 {int} models.GgcmsPowers
// @Failure 403 body is empty
// @router /add [post]
func (c *GgcmsPowersController) Add() {
	var v models.GgcmsPowers
	msg := models.Message{1, "失败", nil}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("")
			exist := models.ExistGgcmsPowers(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if _, err := models.AddGgcmsPowers(&v); err == nil {
					c.Ctx.Output.SetStatus(201)
					msg = models.Message{0, "成功", v}
					c.Data["json"] = msg
				} else {
					msg.Msg = err.Error()
				}
			}
		}
	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Update
// @Description update the GgcmsPowers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsPowers	true		"body for GgcmsPowers content"
// @Success 200 {object} models.GgcmsPowers
// @Failure 403 :id is not int
// @router /edit [post]
func (c *GgcmsPowersController) Edit() {
	var v models.GgcmsPowers
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("")
			exist := models.ExistGgcmsPowers(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if err := models.UpdateGgcmsPowersById(&v); err == nil {
					c.Ctx.Output.SetStatus(201)
					msg = models.Message{0, "成功", v}
					c.Data["json"] = msg
				} else {
					msg.Msg = err.Error()
				}
			}
		}
	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
func (this *GgcmsPowersController) validation(v models.GgcmsPowers) (msg models.Message) {
	msg = models.Message{0, "", nil}
	return
}

// @Title Get
// @Description get GgcmsPowers by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsPowers
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GgcmsPowersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetOneGgcmsPowers(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get GgcmsPowers
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GgcmsPowers
// @Failure 403
// @router / [get]
func (c *GgcmsPowersController) GetAll() {
	var limit int64 = adminpagesize
	var offset int64 = 0
	// fields: col1,col2,entity.col3
	strfields := c.GetString("fields")
	// sortby: col1,col2
	strsortby := c.GetString("sortby")
	// order: desc,asc
	strorder := c.GetString("order")
	// query: k:v,k:v
	strquery := c.GetString("query")
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	l, _, err := GetAllGgcmsPowers(strfields, strsortby, strorder, strquery, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	c.ServeJSON()
}

// @Title Delete
// @Description delete the GgcmsPowers
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *GgcmsPowersController) Delete() {
	msg := models.Message{1, "失败", nil}
	var ids []int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ids); err == nil {
		var num int64
		queryList, _ := getQueryList("")
		if num, err = models.MultDeleteGgcmsPowers(queryList, ids); err == nil {
			msg = models.Message{0, strconv.Itoa(int(num)), ids}
		} else {
			msg.Msg = err.Error()
		}
	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

func GetOneGgcmsPowers(id int) (v *models.GgcmsPowers, err error) {
	v, err = models.GetGgcmsPowersById(id)
	return
}
func GetAllGgcmsPowers(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
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
	return models.GetAllGgcmsPowers(query, fields, sortby, order, offset, pagesize, c)
}
