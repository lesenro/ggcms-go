package controllers

import (
	"encoding/json"
	"strings"

	"ggcms/models"
	"strconv"
)

// oprations for GgcmsSystemConfigs
type GgcmsSystemConfigsController struct {
	BaseController
}

func (c *GgcmsSystemConfigsController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description create GgcmsSystemConfigs
// @Param	body		body 	models.GgcmsSystemConfigs	true		"body for GgcmsSystemConfigs content"
// @Success 201 {int} models.GgcmsSystemConfigs
// @Failure 403 body is empty
// @router /add [post]
func (c *GgcmsSystemConfigsController) Add() {
	var v models.GgcmsSystemConfigs
	msg := models.Message{1, "失败", nil}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("")
			exist := models.ExistGgcmsSystemConfigs(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if _, err := models.AddGgcmsSystemConfigs(&v); err == nil {
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
// @Description update the GgcmsSystemConfigs
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsSystemConfigs	true		"body for GgcmsSystemConfigs content"
// @Success 200 {object} models.GgcmsSystemConfigs
// @Failure 403 :id is not int
// @router /edit [post]
func (c *GgcmsSystemConfigsController) Edit() {
	var v models.GgcmsSystemConfigs
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("")
			exist := models.ExistGgcmsSystemConfigs(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if err := models.UpdateGgcmsSystemConfigsById(&v); err == nil {
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
// @Description update the GgcmsSystemConfigs
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsSystemConfigs	true		"body for GgcmsSystemConfigs content"
// @Success 200 {object} models.GgcmsSystemConfigs
// @Failure 403 :id is not int
// @router /update [post]
func (c *GgcmsSystemConfigsController) Update() {
	var v map[string]models.GgcmsSystemConfigs
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validationUpdate(v)
		if msg.Code == 0 {
			//有上传的文件
			if _, ok := v["UpinfoList"]; ok {
				byts := []byte(v["UpinfoList"].Value)
				var upinfos []models.UpInfo
				if err := json.Unmarshal(byts, &upinfos); err == nil {
					upctrl := GgcmsUploadFileController{BaseController: c.BaseController}
					for _, up := range upinfos {
						val := v[up.InputId]
						if up.Assignment {
							val.Value = upctrl.SaveFile(up)
						} else {
							pic := upctrl.SaveFile(up)
							val.Value = strings.Replace(val.Value, up.Realname, pic, -1)
						}
						v[up.InputId] = val
					}
				}
				delete(v, "UpinfoList")
			}
			//获取当前操作配置id
			sid := -1
			for _, v := range v {
				sid = v.Siteid
				break
			}
			models.MultUpdateGgcmsSystemConfig(v)
			if sid > 0 {
				c.cacheman.ClearByKey(cnSiteConfig + "_" + strconv.Itoa(sid))
			} else {
				c.cacheman.ClearByKey(cnSystemConfig)
			}
			//清理缓存
		}
	} else {
		msg.Msg = err.Error()
		msg.Data = err
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
func (this *GgcmsSystemConfigsController) validationUpdate(v map[string]models.GgcmsSystemConfigs) (msg models.Message) {
	msg = models.Message{0, "", nil}
	//cfgs := this.cacheman.CacheSystemConfigs(this.siteid)
	// for _, val := range cfgs {
	// 	var jdata map[string]interface{}
	// 	if val.Others != "" {
	// 		reg := regexp.MustCompile(`\s+`)
	// 		jsonstr := reg.ReplaceAllString(val.Others, "")
	// 		if err := json.Unmarshal([]byte(jsonstr), &jdata); err == nil {
	// 			if _, ok := jdata["type"]; ok {
	// 				regstr := ``
	// 				switch jdata["type"] {
	// 				case "url":
	// 					regstr = `^(\w+:\/\/)?\w+(\.\w+)+.*$`
	// 				case "number":
	// 					regstr = `^\d+$`
	// 				case "email":
	// 					regstr = `^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	// 				}
	// 				reg = regexp.MustCompile(regstr)
	// 				if !reg.MatchString(val.Value) {
	// 					return models.Message{1, "输入错误", val}
	// 				}
	// 			}
	// 		} else {
	// 			return models.Message{2, err.Error(), err}
	// 		}
	// 	}
	// }
	return
}
func (this *GgcmsSystemConfigsController) validation(v models.GgcmsSystemConfigs) (msg models.Message) {
	msg = models.Message{0, "", nil}

	return
}

// @Title Get
// @Description get GgcmsSystemConfigs by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsSystemConfigs
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GgcmsSystemConfigsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetOneGgcmsSystemConfigs(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get GgcmsSystemConfigs
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GgcmsSystemConfigs
// @Failure 403
// @router / [get]
func (c *GgcmsSystemConfigsController) GetAll() {
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

	l, _, err := GetAllGgcmsSystemConfigs(strfields, strsortby, strorder, strquery, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	c.ServeJSON()
}

// @Title Delete
// @Description delete the GgcmsSystemConfigs
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *GgcmsSystemConfigsController) Delete() {
	msg := models.Message{1, "失败", nil}
	var ids []int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ids); err == nil {
		var num int64
		queryList, _ := getQueryList("")
		if num, err = models.MultDeleteGgcmsSystemConfigs(queryList, ids); err == nil {
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

func GetOneGgcmsSystemConfigs(id int) (v *models.GgcmsSystemConfigs, err error) {
	v, err = models.GetGgcmsSystemConfigsById(id)
	return
}
func GetAllGgcmsSystemConfigs(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
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
	return models.GetAllGgcmsSystemConfigs(query, fields, sortby, order, offset, pagesize, c)
}
