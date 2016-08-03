package controllers

import (
	"encoding/json"
	"errors"
	"ggcms/models"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for GgcmsStyles
type GgcmsStylesController struct {
	beego.Controller
	styledir string
	viewdir  string
	tempdir  string
}

func (c *GgcmsStylesController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}
func (this *GgcmsStylesController) Prepare() {
	//风格路径
	stdir := beego.AppConfig.String("styledir")
	if "" == stdir {
		stdir = "styledir"
	}
	//模板路径
	vdir := beego.AppConfig.String("ViewsPath")
	if "" == vdir {
		vdir = "views"
	}
	//静态文件路径
	sdir := beego.AppConfig.String("StaticDir")
	if "" == sdir {
		sdir = "static"
	}
	//临时文件夹
	tmp := beego.AppConfig.String("tempdir")
	if tmp == "" {
		tmp = "temp"
	}
	this.tempdir = tmp
	if !Exist(tmp) {
		os.MkdirAll(tmp, os.ModePerm)
	}
	//创建风格路径
	this.viewdir = vdir + "/" + stdir
	if !Exist(this.viewdir) {
		os.MkdirAll(this.viewdir, os.ModePerm)
	}
	this.styledir = sdir + "/" + stdir

	if !Exist(this.styledir) {
		os.MkdirAll(this.styledir, os.ModePerm)
	}
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /add [post]
func (c *GgcmsStylesController) Add() {
	var v models.GgcmsStyles
	msg := models.Message{1, "失败", nil}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("")
			exist := models.ExistGgcmsStyles(query)
			if exist {
				msg.Code = 101
				msg.Msg = "已存在，不能重复添加。"
			} else {
				if _, err := models.AddGgcmsStyles(&v); err == nil {
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

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /import [post]
func (c *GgcmsStylesController) ImportStyle() {
	msg := models.Message{1, "失败", nil}
	var v map[string]string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		fn := v["fn"]
		if fn == "" || !Exist(fn) {
			msg.Msg = "文件上传失败"
		} else {
			tmp := c.tempdir
			dir := RandString()
			to := tmp + "/" + dir
			for Exist(to) {
				dir = RandString()
				to = tmp + "/" + dir
			}
			//创建临时目录
			os.MkdirAll(to, os.ModePerm)
			//解压
			if err := DeCompress(fn, to); err != nil {
				msg.Msg = err.Error()
				msg.Data = err
			} else {
				//读配置
				cfg := to + "/config.json"
				v := models.GgcmsStyles{Style_name: "新风格", Style_folder: dir, Style_descrip: "风格配置文件不存在,"}
				if Exist(cfg) {
					if err := json.Unmarshal(ReadFileBytes(cfg), &v); err != nil {
						v = models.GgcmsStyles{Style_name: "新风格", Style_folder: dir, Style_descrip: "风格配置文件格式不正确,"}
					}
					//删除风格配置文件
					os.Remove(cfg)
				}

				vpath := c.viewdir + "/" + v.Style_folder
				spath := c.styledir + "/" + v.Style_folder
				//路径是否存在
				for Exist(vpath) || Exist(spath) {
					v.Style_folder = RandString()
					vpath = c.viewdir + "/" + v.Style_folder
					spath = c.styledir + "/" + v.Style_folder
				}
				if Exist(to + "/template") {
					//移动模板文件
					err = os.Rename(to+"/template", vpath)
				} else {
					os.MkdirAll(vpath, os.ModePerm)
					v.Style_descrip = v.Style_descrip + "模板文件不存在"
				}
				//移动风格文件
				if err == nil {
					err = os.Rename(to, spath)
				} else {
					//出错后删除已移动的模板文件
					os.RemoveAll(vpath)
				}
				if err == nil {
					msg.Code = 0
					msg.Data = "成功"
					models.AddGgcmsStyles(&v)
				} else {
					msg.Data = err
				}
			}

		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /export/:id [get]
func (c *GgcmsStylesController) ExportStyle() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if v, err := GetOneGgcmsStyles(id); err == nil {
		//临时文件夹
		tmp := c.tempdir
		//压缩文件
		zf := RandString()
		zipfile := tmp + "/" + zf + ".zip"
		for Exist(zipfile) {
			zf = RandString()
			zipfile = tmp + "/" + zf + ".zip"
		}
		//风格路径
		spath := c.styledir + "/" + v.Style_folder
		//风格配置文件
		cfg := spath + "/config.json"
		ToJsonFile(&v, cfg)
		//模板路径
		vpath := c.viewdir + "/" + v.Style_folder
		//在风格路径中创建模板路径，方便压缩
		svpath := spath + "/template"
		//删除风格路径中的模板文件
		os.RemoveAll(svpath)
		os.MkdirAll(svpath, os.ModePerm)
		//复制模板到风格路径
		list, _ := DirList(vpath)
		for _, fn := range list {
			name := path.Base(fn)
			tofile := svpath + "/" + name
			CopyFile(tofile, fn)
		}
		//添加风格路径到压缩文件列表
		var files []*os.File
		list, _ = DirList(spath)
		for _, fn := range list {
			f, _ := os.Open(fn)
			defer f.Close()
			files = append(files, f)
		}

		//压缩
		err = Compress(files, zipfile)
		//CCC(spath, zipfile)
		//删除风格路径中的模板文件
		os.RemoveAll(svpath)
		os.RemoveAll(cfg)

		if err == nil {
			c.Ctx.Output.Download(zipfile, v.Style_folder+".zip")
		}

	}
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /template/:id [get]
func (c *GgcmsStylesController) TemplateList() {

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	msg := models.Message{1, "失败", nil}
	v, err := GetOneGgcmsStyles(id)
	if err != nil {
		msg.Data = err
	} else {
		vpath := c.viewdir + "/" + v.Style_folder
		flist, err := FileAttrList(vpath)
		if err != nil {
			msg.Data = err
		} else {
			msg = models.Message{0, "成功", nil}
			msg.Data = flist
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /uploadtemplate [post]
func (c *GgcmsStylesController) UploadTemplate() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		fn := v["fn"]
		name := v["name"]
		if fn == "" || name == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.viewdir + "/" + info.Style_folder
			tfn := vpath + "/" + name
			ext := path.Ext(tfn)
			for Exist(tfn) {
				tfn = vpath + "/" + RandString() + ext
			}
			err = os.Rename(fn, tfn)
			if err != nil {
				msg.Data = err
			} else {
				if flist, err := FileAttrList(vpath); err != nil {
					msg.Data = err
				} else {
					msg = models.Message{0, "成功", flist}
				}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /deletetemplate [post]
func (c *GgcmsStylesController) DeleteTemplate() {
	var v map[string]interface{}
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"].(string))
		fs := v["fs"].([]interface{})
		if len(fs) == 0 {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.viewdir + "/" + info.Style_folder
			for _, f := range fs {
				tfn := vpath + "/" + f.(string)
				os.Remove(tfn)
			}
			if flist, err := FileAttrList(vpath); err != nil {
				msg.Data = err
			} else {
				msg = models.Message{0, "成功", flist}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /renametemplate [post]
func (c *GgcmsStylesController) RenameTemplate() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		name := v["name"]
		sname := v["sname"]
		if name == "" || sname == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.viewdir + "/" + info.Style_folder
			tfn := vpath + "/" + name
			sfn := vpath + "/" + sname
			if Exist(tfn) {
				msg.Msg = "文件已存在，修改失败"
			} else {
				os.Rename(sfn, tfn)
				if flist, err := FileAttrList(vpath); err != nil {
					msg.Data = err
				} else {
					msg = models.Message{0, "成功", flist}
				}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /gettemplate [post]
func (c *GgcmsStylesController) GetTemplate() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		name := v["name"]
		if name == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.viewdir + "/" + info.Style_folder
			sfn := vpath + "/" + name
			if !Exist(sfn) {
				msg.Msg = "文件不存在"
			} else {
				msg = models.Message{0, "成功", ReadFileString(sfn)}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /savetemplate [post]
func (c *GgcmsStylesController) SaveTemplate() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		name := v["name"]
		tcode := v["tcode"]
		if name == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.viewdir + "/" + info.Style_folder
			sfn := vpath + "/" + name
			if !Exist(sfn) {
				msg.Msg = "文件不存在"
			} else {
				err = StringToFile(tcode, sfn)
				if err != nil {
					msg.Data = err
				} else {
					msg = models.Message{0, "成功", nil}
				}

			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /uploadstatic [post]
func (c *GgcmsStylesController) UploadStatic() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		fn := v["fn"]
		name := v["name"]
		p := v["path"]
		if p != "" {
			p = "/" + p
		}

		if fn == "" || name == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.styledir + "/" + info.Style_folder + p
			tfn := vpath + "/" + name
			ext := path.Ext(tfn)
			for Exist(tfn) {
				tfn = vpath + "/" + RandString() + ext
			}
			err = os.Rename(fn, tfn)
			if err != nil {
				msg.Data = err
			} else {
				if flist, err := FileAttrList(vpath); err != nil {
					msg.Data = err
				} else {
					msg = models.Message{0, "成功", flist}
				}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @router /staticdelete [post]
func (c *GgcmsStylesController) DeleteStatic() {
	var v map[string]interface{}
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"].(string))
		fs := v["fs"].([]interface{})
		p := v["path"].(string)
		if p != "" {
			p = "/" + p
		}
		if len(fs) == 0 {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.styledir + "/" + info.Style_folder + p
			for _, f := range fs {
				tfn := vpath + "/" + f.(string)
				finfo, _ := os.Stat(tfn)
				if finfo.IsDir() {
					//目录不允许删除
					//os.RemoveAll(tfn)
				} else {
					os.Remove(tfn)
				}

			}
			if flist, err := FileAttrList(vpath); err != nil {
				msg.Data = err
			} else {
				msg = models.Message{0, "成功", flist}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /renamestatic [post]
func (c *GgcmsStylesController) RenameStatic() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		name := v["name"]
		sname := v["sname"]
		p := v["path"]
		if p != "" {
			p = "/" + p
		}
		if name == "" || sname == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.styledir + "/" + info.Style_folder + p
			tfn := vpath + "/" + name
			sfn := vpath + "/" + sname
			if Exist(tfn) {
				msg.Msg = "文件已存在，修改失败"
			} else {
				os.Rename(sfn, tfn)
				if flist, err := FileAttrList(vpath); err != nil {
					msg.Data = err
				} else {
					msg = models.Message{0, "成功", flist}
				}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /getstaticfile [post]
func (c *GgcmsStylesController) GetStaticFile() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		name := v["name"]
		p := v["path"]
		if p != "" {
			p = "/" + p
		}
		if name == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.styledir + "/" + info.Style_folder + p
			sfn := vpath + "/" + name
			if !Exist(sfn) {
				msg.Msg = "文件不存在"
			} else {
				msg = models.Message{0, "成功", ReadFileString(sfn)}
			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /savestaticfile [post]
func (c *GgcmsStylesController) SaveStaticFile() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		name := v["name"]
		tcode := v["tcode"]
		p := v["path"]
		if p != "" {
			p = "/" + p
		}
		if name == "" {
			msg.Msg = "参数错误"
		} else if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.styledir + "/" + info.Style_folder + p
			sfn := vpath + "/" + name
			if !Exist(sfn) {
				msg.Msg = "文件不存在"
			} else {
				err = StringToFile(tcode, sfn)
				if err != nil {
					msg.Data = err
				} else {
					msg = models.Message{0, "成功", nil}
				}

			}
		} else {
			msg.Data = err
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /staticfile [post]
func (c *GgcmsStylesController) StaticList() {
	var v map[string]string
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		msg.Data = err
	} else {
		id, _ := strconv.Atoi(v["id"])
		p := v["path"]
		if info, err := GetOneGgcmsStyles(id); err == nil {
			vpath := c.styledir + "/" + info.Style_folder
			sfn := vpath + "/" + p
			if !Exist(sfn) {
				msg.Msg = "访问路径不存在"
			} else {
				flist, err := FileAttrList(sfn)
				if err != nil {
					msg.Data = err
				} else {
					msg = models.Message{0, "成功", flist}
				}
			}
		}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}

// @Title Post
// @Description create GgcmsStyles
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 201 {int} models.GgcmsStyles
// @Failure 403 body is empty
// @router /all [get]
func (c *GgcmsStylesController) GetAllData() {
	msg := models.Message{1, "失败", nil}
	data, err := styleAndTemplate()
	if err == nil {
		msg = models.Message{0, "成功", data}
	} else {
		msg.Data = err
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
func styleAndTemplate() (data map[string]interface{}, err error) {
	l, _, err := GetAllGgcmsStyles("", "", "", "", 0, 0, false)
	if err != nil {
		return nil, err
	}
	//创建风格路径
	stdir := beego.AppConfig.String("styledir")
	if "" == stdir {
		stdir = "styledir"
	}
	vdir := beego.AppConfig.String("ViewsPath")
	if "" == vdir {
		vdir = "views"
	}
	vdir = vdir + "/" + stdir

	data = make(map[string]interface{})
	data["style"] = l
	for _, v := range l {
		info := v.(models.GgcmsStyles)
		vpath := vdir + "/" + info.Style_folder
		flist, _ := FileAttrList(vpath)
		data["t_"+info.Style_folder] = flist
	}
	return data, nil
}

// @Title Update
// @Description update the GgcmsStyles
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GgcmsStyles	true		"body for GgcmsStyles content"
// @Success 200 {object} models.GgcmsStyles
// @Failure 403 :id is not int
// @router /edit [post]
func (c *GgcmsStylesController) Edit() {
	var v models.GgcmsStyles
	msg := models.Message{1, "失败", nil}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		msg = c.validation(v)
		if msg.Code == 0 {
			query, _ := getQueryList("id.ne:" + strconv.Itoa(v.Id) + ",style_folder:" + v.Style_folder)
			exist := models.ExistGgcmsStyles(query)
			if exist {
				msg.Code = 101
				msg.Msg = "目录已存在，不能重复。"
			} else {

				if old, err := GetOneGgcmsStyles(v.Id); err == nil {
					//目录修改
					if old.Style_folder != v.Style_folder {
						vpath := c.viewdir + "/" + v.Style_folder
						spath := c.styledir + "/" + v.Style_folder
						if Exist(vpath) || Exist(spath) {
							err = errors.New("目录已存在，不能重复")
							msg.Code = 101
						} else {
							//移动目录
							ovpath := c.viewdir + "/" + old.Style_folder
							ospath := c.styledir + "/" + old.Style_folder
							err = os.Rename(ovpath, vpath)
							if err == nil {
								err = os.Rename(ospath, spath)
							}
						}
						//修改风格
						if err == nil {
							if err := models.UpdateGgcmsStylesById(&v); err == nil {
								c.Ctx.Output.SetStatus(201)
								msg = models.Message{0, "成功", v}
								c.Data["json"] = msg
							} else {
								msg.Msg = err.Error()
							}
						} else {
							msg.Msg = err.Error()
						}
					}
				}
			}
		}
	} else {
		msg.Msg = err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
func (this *GgcmsStylesController) validation(v models.GgcmsStyles) (msg models.Message) {
	msg = models.Message{0, "", nil}
	return
}

// @Title Get
// @Description get GgcmsStyles by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GgcmsStyles
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GgcmsStylesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := GetOneGgcmsStyles(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get GgcmsStyles
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GgcmsStyles
// @Failure 403
// @router / [get]
func (c *GgcmsStylesController) GetAll() {
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

	l, _, err := GetAllGgcmsStyles(strfields, strsortby, strorder, strquery, offset, limit, false)
	if err != nil {
		c.Data["json"] = models.Message{1, "失败", err}
	} else {
		c.Data["json"] = models.Message{0, "成功", l}
	}

	c.ServeJSON()
}

// @Title Delete
// @Description delete the GgcmsStyles
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete [post]
func (c *GgcmsStylesController) Delete() {
	msg := models.Message{1, "失败", nil}
	var ids []int
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ids); err == nil {
		var num int64
		queryList, _ := getQueryList("")
		for _, id := range ids {
			if v, err := GetOneGgcmsStyles(id); err == nil {
				vpath := c.viewdir + "/" + v.Style_folder
				if Exist(vpath) {
					os.RemoveAll(vpath)
				}
				spath := c.styledir + "/" + v.Style_folder
				if Exist(spath) {
					os.RemoveAll(spath)
				}
			}
		}
		if num, err = models.MultDeleteGgcmsStyles(queryList, ids); err == nil {
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

func GetOneGgcmsStyles(id int) (v *models.GgcmsStyles, err error) {
	v, err = models.GetGgcmsStylesById(id)
	return
}
func GetAllGgcmsStyles(strfields, strsortby, strorder, strquery string, pagenum int64, pagesize int64, c bool) (ml []interface{}, count int64, err error) {
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
	return models.GetAllGgcmsStyles(query, fields, sortby, order, offset, pagesize, c)
}
