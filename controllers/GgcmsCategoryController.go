package controllers

import (
	"encoding/json"
	"ggcms/models"
	"strconv"
	"strings"
	"time"
)

// oprations for GgcmsCategory
type GgcmsCategoryController struct {
	BaseController
}

func (c *GgcmsCategoryController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description create GgcmsCategory
// @Param	body		body 	models.GgcmsCategory	true		"body for GgcmsCategory content"
// @Success 201 {int} models.GgcmsCategory
// @Failure 403 body is empty
// @router /add [post]
func (c *GgcmsCategoryController) Add() {
	var v models.GgcmsCategory
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			v.Lastupdate = time.Now()
			v.Siteid = c.currentSite
			//有上传的文件
			var upinfos models.UpInfos
			if err := json.Unmarshal(c.Ctx.Input.RequestBody, &upinfos); err == nil {
				upctrl := GgcmsUploadFileController{BaseController: c.BaseController}
				for _, up := range upinfos.UpinfoList {
					//LOGO图上传
					if up.InputId == "Logo" {
						v.Logo = upctrl.SaveFile(up)
						continue
					}
					//富编辑框
					if up.InputId == "Content" {
						pic := upctrl.SaveFile(up)
						v.Content = strings.Replace(v.Content, up.Realname, pic, -1)
						continue
					}
				}
			}
			if _, err := models.AddGgcmsCategory(&v); err == nil {
				sid := c.currentSite
				c.cacheman.ClearByKey(cnCategoryList + "_" + strconv.Itoa(sid))

				c.Ctx.Output.SetStatus(201)
				msg = models.Message{0, "成功", v}
				c.Data["json"] = msg
			} else {
				msg.Msg = err.Error()
			}
		}

	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsCategory
// @Param	body		body 	models.GgcmsCategory	true		"body for GgcmsCategory content"
// @Success 201 {int} models.GgcmsCategory
// @Failure 403 body is empty
// @router /savesort [post]
func (c *GgcmsCategoryController) SaveSort() {
	var v []models.GgcmsCategory
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err = models.UpdateSortData(v); err == nil {
			sid := c.currentSite
			c.cacheman.ClearByKey(cnCategoryList + "_" + strconv.Itoa(sid))

			msg = models.Message{0, "成功", nil}
		} else {
			msg.Msg = err.Error()
		}
	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Update
// @Description update the GgcmsCategory
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsCategory	true		"body for GgcmsCategory content"
// @Success 200 {object} models.GgcmsCategory
// @Failure 403 :id is not int
// @router /edit [post]
func (c *GgcmsCategoryController) Edit() {
	var v models.GgcmsCategory
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		msg = c.validation(v)
		if msg.Code == 0 {
			v.Lastupdate = time.Now()
			//有上传的文件
			var upinfos models.UpInfos
			if err := json.Unmarshal(c.Ctx.Input.RequestBody, &upinfos); err == nil {
				upctrl := GgcmsUploadFileController{BaseController: c.BaseController}
				for _, up := range upinfos.UpinfoList {
					//LOGO图上传
					if up.InputId == "Logo" {
						v.Logo = upctrl.SaveFile(up)
						continue
					}
					//富编辑框
					if up.InputId == "Content" {
						pic := upctrl.SaveFile(up)
						v.Content = strings.Replace(v.Content, up.Realname, pic, -1)
						continue
					}
				}
			}
			if err := models.UpdateGgcmsCategoryById(&v); err == nil {
				sid := c.currentSite
				c.cacheman.ClearByKey(cnCategoryList + "_" + strconv.Itoa(sid))

				c.Ctx.Output.SetStatus(201)
				msg = models.Message{0, "成功", v}
				c.Data["json"] = msg
			} else {
				msg.Msg = err.Error()
			}
		}
	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
func (this *GgcmsCategoryController) validation(v models.GgcmsCategory) (msg models.Message) {
	msg = models.Message{0, "", nil}
	return
}

// @Title Get
// @Description get GgcmsCategory by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsCategory
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GgcmsCategoryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetOneGgcmsCategory(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get GgcmsCategory
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GgcmsCategory
// @Failure 403
// @router / [get]
func (c *GgcmsCategoryController) GetAll() {
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

	l, _, err := GetAllGgcmsCategory(strfields, strsortby, strorder, strquery, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	c.ServeJSON()
}

// @Title Delete
// @Description delete the GgcmsCategory
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *GgcmsCategoryController) Delete() {
	msg := models.Message{1, "失败", nil}
	var ids []int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ids); err == nil {
		var num int64
		queryList, _ := getQueryList("")
		if num, err = models.MultDeleteGgcmsCategory(queryList, ids); err == nil {
			sid := c.currentSite
			c.cacheman.ClearByKey(cnCategoryList + "_" + strconv.Itoa(sid))
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

func GetOneGgcmsCategory(id int) (v *models.GgcmsCategory, err error) {
	v, err = models.GetGgcmsCategoryById(id)
	return
}
func GetAllGgcmsCategory(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
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
	return models.GetAllGgcmsCategory(query, fields, sortby, order, offset, pagesize, c)
}
