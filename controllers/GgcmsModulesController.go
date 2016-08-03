package controllers

import (
	"encoding/json"
	"ggcms/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// oprations for GgcmsModules
type GgcmsModulesController struct {
	beego.Controller
}

func (c *GgcmsModulesController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description create GgcmsModules
// @Param	body		body 	models.GgcmsModules	true		"body for GgcmsModules content"
// @Success 201 {int} models.GgcmsModules
// @Failure 403 body is empty
// @router /add [post]
func (c *GgcmsModulesController) Add() {
	var v models.GgcmsModulesPostData
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if obj, err := models.RunSql("SELECT IFNULL(MAX(ggcms_modules.id),0)+1  as maxid FROM ggcms_modules", make([]interface{}, 0), "maxid"); err == nil {
			v.Table_name = "model_tab_" + obj[0]["maxid"].(string)
			v.View_name = "model_view_" + obj[0]["maxid"].(string)
			msg = c.validation(v.GgcmsModules)
			if msg.Code == 0 {
				query, _ := getQueryList("table_name:" + v.Table_name)
				exist := models.ExistGgcmsModules(query)
				if exist {
					msg.Code = 101
					msg.Msg = "已存在，不能重复添加。"
				} else {
					if _, err := models.AddGgcmsModules(&v); err == nil {
						c.Ctx.Output.SetStatus(201)
						msg = models.Message{0, "成功", v.Id}
						c.Data["json"] = msg
					} else {
						msg.Msg = err.Error()
					}
				}
			}
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
// @Description update the GgcmsModules
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsModules	true		"body for GgcmsModules content"
// @Success 200 {object} models.GgcmsModules
// @Failure 403 :id is not int
// @router /edit [post]
func (c *GgcmsModulesController) Edit() {
	var v models.GgcmsModulesPostData
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v.GgcmsModules)
		if msg.Code == 0 {
			query, _ := getQueryList("id.ne:" + strconv.Itoa(v.Id) + ",table_name:" + v.Table_name)
			exist := models.ExistGgcmsModules(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if err := models.UpdateGgcmsModulesById(&v); err == nil {
					c.Ctx.Output.SetStatus(201)
					msg = models.Message{0, "成功", v.Id}
					c.Data["json"] = msg
				} else {
					msg.Code = 1
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
func (this *GgcmsModulesController) validation(v models.GgcmsModules) (msg models.Message) {
	msg = models.Message{0, "", nil}
	if len(v.Table_name) < 4 && len(v.Table_name) > 50 {
		msg = models.Message{201, "", nil}
	}
	return
}

// @Title Get
// @Description get GgcmsModules by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsModules
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GgcmsModulesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetOneGgcmsModules(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get GgcmsModules
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GgcmsModules
// @Failure 403
// @router / [get]
func (c *GgcmsModulesController) GetAll() {
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

	l, _, err := GetAllGgcmsModules(strfields, strsortby, strorder, strquery, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	c.ServeJSON()
}

// @Title Delete
// @Description delete the GgcmsModules
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *GgcmsModulesController) Delete() {
	msg := models.Message{1, "失败", nil}
	var ids []int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ids); err == nil {
		var num int64
		queryList, _ := getQueryList("")
		if num, err = models.MultDeleteGgcmsModules(queryList, ids); err == nil {
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

func GetOneGgcmsModules(id int) (v *models.GgcmsModules, err error) {
	v, err = models.GetGgcmsModulesById(id)
	return
}

// @router /getcolumns/:id [get]
func (c *GgcmsModulesController) GetModulesColsByMid() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetModulesColumnsByMid(id)
	if err != nil {
		c.Data["json"] = models.Message{1, err.Error(), err}
	} else {
		c.Data["json"] = models.Message{0, "成功", v}
	}
	c.ServeJSON()
}
func GetModulesColumnsByMid(id int) (ml []models.GgcmsModulesColumns, err error) {
	var strs []string
	strquery := "tid:" + strconv.Itoa(id)
	// query: k:v,k:v
	if query, err := getQueryList(strquery); err != nil {
		return nil, err
	} else {
		ml, _, err = models.GetAllGgcmsModulesColumns(query, strs, strs, strs, 0, 0, false)
	}
	return
}
func GetModulesInfo(mid, aid int) (maps *[]orm.Params) {
	v, err := models.GetGgcmsModulesById(mid)
	if err != nil {
		return nil
	}
	maps, err = models.GetGgcmsModulesInfo(aid, v.Table_name)
	if err == nil {
		return maps
	}
	return nil
}
func GetAllGgcmsModules(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
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
	return models.GetAllGgcmsModules(query, fields, sortby, order, offset, pagesize, c)
}
