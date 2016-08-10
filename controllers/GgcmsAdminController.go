package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"ggcms/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for GgcmsAdmin
type GgcmsAdminController struct {
	beego.Controller
}

func (c *GgcmsAdminController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description create GgcmsAdmin
// @Param	body		body 	models.GgcmsAdmin	true		"body for GgcmsAdmin content"
// @Success 201 {int} models.GgcmsAdmin
// @Failure 403 body is empty
// @router /add [post]
func (c *GgcmsAdminController) Add() {
	var v models.GgcmsAdmin
	msg := models.Message{1, "失败", nil}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("userid:" + v.Userid)
			exist := models.ExistGgcmsAdmin(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if _, err := models.AddGgcmsAdmin(&v); err == nil {
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
// @Description update the GgcmsAdmin
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsAdmin	true		"body for GgcmsAdmin content"
// @Success 200 {object} models.GgcmsAdmin
// @Failure 403 :id is not int
// @router /edit [post]
func (c *GgcmsAdminController) Edit() {
	var v models.GgcmsAdmin
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("userid:" + v.Userid + ",id.ne:" + strconv.Itoa(v.Id))
			exist := models.ExistGgcmsAdmin(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if strings.TrimSpace(v.Pwd) != "" {
					m := md5.New()
					m.Write([]byte(v.Pwd))
					v.Pwd = hex.EncodeToString(m.Sum(nil))
				} else {
					v.Pwd = ""
				}
				if err := models.UpdateGgcmsAdminById(&v); err == nil {
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
func (this *GgcmsAdminController) validation(v models.GgcmsAdmin) (msg models.Message) {
	msg = models.Message{0, "", nil}
	return
}

// @Title Get
// @Description get GgcmsAdmin by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsAdmin
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GgcmsAdminController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetOneGgcmsAdmin(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @router /login [post]
func (c *GgcmsAdminController) GetOneByName() {
	msg := models.Message{1, "失败", nil}
	//解析表单数据
	var f interface{}
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &f)
	if err != nil {
		msg = models.Message{1, err.Error(), err}
	} else {
		postData := f.(map[string]interface{})
		name := postData["username"].(string)
		pass := postData["password"].(string)
		code := strings.ToLower(postData["checkcode"].(string))
		sesscode := c.GetSession("checkcode")
		c.DelSession("checkcode")
		if sesscode == nil || strings.ToLower(sesscode.(string)) != code {
			msg = models.Message{1, "验证码错误，请检查:", nil}
		} else {
			var v *models.GgcmsAdmin
			v, err = models.GetGgcmsAdminByName(name)
			if err != nil {
				msg = models.Message{1, "用户名或密码错误，请检查", err}
			} else {
				m := md5.New()
				m.Write([]byte(v.Pwd + code))
				md5pass := hex.EncodeToString(m.Sum(nil))
				if pass != md5pass {
					msg = models.Message{1, "用户名或密码错误，请检查", nil}
				} else {
					c.SetSession("uid", v.Id)
					msg = models.Message{0, "成功", nil}
				}
			}
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Get All
// @Description get GgcmsAdmin
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GgcmsAdmin
// @Failure 403
// @router / [get]
func (c *GgcmsAdminController) GetAll() {
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

	l, _, err := GetAllGgcmsAdmin(strfields, strsortby, strorder, strquery, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	c.ServeJSON()
}

// @Title Delete
// @Description delete the GgcmsAdmin
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *GgcmsAdminController) Delete() {
	msg := models.Message{1, "失败", nil}
	var ids []int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ids); err == nil {
		var num int64
		queryList, _ := getQueryList("")
		if num, err = models.MultDeleteGgcmsAdmin(queryList, ids); err == nil {
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

func GetOneGgcmsAdmin(id int) (v *models.GgcmsAdmin, err error) {
	v, err = models.GetGgcmsAdminById(id)
	return
}
func GetAllGgcmsAdmin(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
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
	return models.GetAllGgcmsAdmin(query, fields, sortby, order, offset, pagesize, c)
}
