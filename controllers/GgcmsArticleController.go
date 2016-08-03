package controllers

import (
	"encoding/json"
	"errors"
	"ggcms/models"
	"strconv"
	"strings"
	"time"
)

// oprations for GgcmsArticle
type GgcmsArticleController struct {
	BaseController
}

func (c *GgcmsArticleController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// @Title Post
// @Description create GgcmsArticle
// @Param	body		body 	models.GgcmsArticle	true		"body for GgcmsArticle content"
// @Success 201 {int} models.GgcmsArticle
// @Failure 403 body is empty
// @router /add [post]
func (c *GgcmsArticleController) Add() {
	err := func() (err error) {
		//文章信息
		var v models.GgcmsArticle
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
		if err != nil {
			return
		}
		//分页
		var pages models.ArticlePages
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &pages)
		if err != nil {
			return
		}

		//解析表单数据
		var f interface{}
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &f)
		if err != nil {
			return
		}
		postData := f.(map[string]interface{})
		//模型信息
		var modulesInfo map[string]interface{}
		if v.Mid > 0 {
			modulesCols := postData["modulesCols"].([]interface{})
			for _, val := range modulesCols {
				amap := val.(map[string]interface{})
				cname := amap["CName"].(string)
				value := amap["Value"]
				modulesInfo[cname] = value
			}
		}
		//有上传的文件
		upinfos := postData["UpinfoList"].([]interface{})
		upctrl := GgcmsUploadFileController{BaseController: c.BaseController}
		if len(upinfos) > 0 {
			for _, val := range upinfos {
				upmap := val.(map[string]interface{})
				upinfo := models.UpInfo{}
				upinfo.InputId = upmap["InputId"].(string)
				upinfo.Realname = upmap["Realname"].(string)
				upinfo.Name = upmap["Name"].(string)
				upinfo.Assignment = upmap["Assignment"].(bool)
				pic := upctrl.SaveFile(upinfo)
				//LOGO图上传
				if upinfo.InputId == "TitleImg" {
					v.TitleImg = pic
					continue
				}
				//富编辑框
				if upinfo.InputId == "Content" {
					for i, _ := range pages.Pages {
						if strings.Contains(pages.Pages[i].Content, upinfo.Realname) {
							pages.Pages[i].Content = strings.Replace(pages.Pages[i].Content, upinfo.Realname, pic, -1)
						}
					}
					continue
				}
				//模型信息
				for key, _ := range modulesInfo {
					if key == upinfo.InputId {
						if upinfo.Assignment {
							//直接上传
							modulesInfo[key] = pic
						} else {
							//富文本框
							val := modulesInfo[key].(string)
							modulesInfo[key] = strings.Replace(val, upinfo.Realname, pic, -1)
						}
					}
				}
			}
		}

		//补充数据
		v.Siteid = c.currentSite
		v.Dateandtime = time.Now()
		//附件
		attachs := postData["attachs"].([]interface{})
		var attalist = make([]models.GgcmsArticleAttr, 0)
		for _, val := range attachs {
			amap := val.(map[string]interface{})
			atta := models.GgcmsArticleAttr{}
			atta.AttrUrl = amap["AttrUrl"].(string)
			atta.Info = amap["Info"].(string)
			atta.Originalname = amap["Originalname"].(string)
			atta.Size = int(amap["Size"].(float64))
			atta.Addtime = time.Now()
			//保存附件
			atta.AttrUrl = upctrl.SaveFile(models.AttachmentToUpInfo(&atta))
			attalist = append(attalist, atta)
		}

		//校验
		msg := c.validation(v)
		if msg.Code != 0 {
			return errors.New("输入有误")
		}
		//保存数据
		_, err = models.AddGgcmsArticle(&v, &pages, &attalist)
		if err == nil {
			c.Ctx.Output.SetStatus(201)
		}
		return
	}()
	msg := models.Message{0, "成功", nil}
	if err != nil {
		msg = models.Message{1, err.Error(), err}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Update
// @Description update the GgcmsArticle
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsArticle	true		"body for GgcmsArticle content"
// @Success 200 {object} models.GgcmsArticle
// @Failure 403 :id is not int
// @router /edit [post]
func (c *GgcmsArticleController) Edit() {
	err := func() (err error) {
		var v models.GgcmsArticle
		var pages models.ArticlePages
		//文章信息
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
		if err != nil {
			return
		}
		//分页信息
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &pages)
		if err != nil {
			return
		}
		//解析表单数据
		var f interface{}
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &f)
		if err != nil {
			return
		}

		postData := f.(map[string]interface{})
		//模型信息
		modulesInfo := make(map[string]interface{})
		if v.Mid > 0 {
			modulesCols := postData["modulesCols"].([]interface{})
			for _, val := range modulesCols {
				amap := val.(map[string]interface{})
				cname := amap["CName"].(string)
				value := amap["Value"]
				modulesInfo[cname] = value
			}
		}
		//有上传的文件
		upinfos := postData["UpinfoList"].([]interface{})
		upctrl := GgcmsUploadFileController{BaseController: c.BaseController}
		if len(upinfos) > 0 {
			for _, val := range upinfos {
				upmap := val.(map[string]interface{})
				upinfo := models.UpInfo{}
				upinfo.InputId = upmap["InputId"].(string)
				upinfo.Realname = upmap["Realname"].(string)
				upinfo.Name = upmap["Name"].(string)
				upinfo.Assignment = upmap["Assignment"].(bool)
				pic := upctrl.SaveFile(upinfo)
				//LOGO图上传
				if upinfo.InputId == "TitleImg" {
					v.TitleImg = pic
					continue
				}
				//富编辑框
				if upinfo.InputId == "Content" {
					for i, _ := range pages.Pages {
						if strings.Contains(pages.Pages[i].Content, upinfo.Realname) {
							pages.Pages[i].Content = strings.Replace(pages.Pages[i].Content, upinfo.Realname, pic, -1)
						}
					}
					continue
				}
				//模型信息
				for key, _ := range modulesInfo {
					if key == upinfo.InputId {
						if upinfo.Assignment {
							//直接上传
							modulesInfo[key] = pic
						} else {
							//富文本框
							val := modulesInfo[key].(string)
							modulesInfo[key] = strings.Replace(val, upinfo.Realname, pic, -1)
						}
					}
				}
			}
		}
		//更新时间
		v.Dateandtime = time.Now()
		//附件
		attachs := postData["attachs"].([]interface{})
		var attalist = make([]models.GgcmsArticleAttr, 0)
		for _, val := range attachs {
			amap := val.(map[string]interface{})
			atta := models.GgcmsArticleAttr{}
			atta.AttrUrl = amap["AttrUrl"].(string)
			atta.Info = amap["Info"].(string)
			atta.Originalname = amap["Originalname"].(string)
			atta.Size = int(amap["Size"].(float64))
			atta.Articleid = int(amap["Articleid"].(float64))
			atta.Id = int(amap["Id"].(float64))
			atta.Addtime = time.Now()
			//保存附件
			if atta.Id == 0 {
				atta.AttrUrl = upctrl.SaveFile(models.AttachmentToUpInfo(&atta))
			}
			attalist = append(attalist, atta)
		}

		//校验
		msg := c.validation(v)
		if msg.Code != 0 {
			return errors.New("输入有误")
		}
		//解析删除页
		ids := postData["DeletePages"].([]interface{})
		//更新数据库
		err = models.UpdateGgcmsArticleById(&v, &pages, ids, &attalist, &modulesInfo)
		if err == nil {
			c.Ctx.Output.SetStatus(201)
		}
		return
	}()
	msg := models.Message{0, "成功", nil}
	if err != nil {
		msg = models.Message{1, err.Error(), err}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
func (this *GgcmsArticleController) validation(v models.GgcmsArticle) (msg models.Message) {
	msg = models.Message{0, "", nil}
	return
}

// @Title Get
// @Description get GgcmsArticle by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsArticle
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GgcmsArticleController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetOneGgcmsArticle(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get
// @Description get GgcmsArticle by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsArticle
// @Failure 403 :id is empty
// @router /pageinfo [get]
func (c *GgcmsArticleController) GetPageInfo() {
	msg := models.Message{1, "失败", nil}
	var aid, sid int
	var err error
	var v *models.GgcmsArticlePages
	aid, err = c.GetInt("aid")
	if err == nil {
		sid, err = c.GetInt("sid")
		if err == nil {
			v, err = models.GetGgcmsArticlePagesBySortId(aid, sid)
		}
	}
	if err != nil {
		msg.Msg = err.Error()
		msg.Data = err
	} else {
		msg = models.Message{0, "成功", v}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Get All
// @Description get GgcmsArticle
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GgcmsArticle
// @Failure 403
// @router / [get]
func (c *GgcmsArticleController) GetAll() {
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

	l, _, err := GetAllGgcmsArticle(strfields, strsortby, strorder, strquery, offset, limit, false)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	c.ServeJSON()
}

// @Title Delete
// @Description delete the GgcmsArticle
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *GgcmsArticleController) Delete() {
	msg := models.Message{1, "失败", nil}
	var ids []int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ids); err == nil {
		var num int64
		queryList, _ := getQueryList("")
		if num, err = models.MultDeleteGgcmsArticle(queryList, ids); err == nil {
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

func GetOneGgcmsArticle(id int) (v *models.GgcmsArticle, err error) {
	v, err = models.GetGgcmsArticleById(id)
	return
}
func GetAllGgcmsArticle(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
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
	return models.GetAllGgcmsArticle(query, fields, sortby, order, offset, pagesize, c)
}
