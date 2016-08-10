package controllers

import (
	"fmt"
	"ggcms/models"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"

	"github.com/astaxie/beego"
	"github.com/disintegration/imaging"
)

// oprations for GgcmsAdmin
type GgcmsUploadFileController struct {
	BaseController
	sdir      string
	tempdir   string
	uploaddir string
}
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func (this *GgcmsUploadFileController) Prepare() {
	this.sdir = beego.AppConfig.String("StaticDir")
	if this.sdir == "" {
		this.sdir = "static"
	}
	this.uploaddir = beego.AppConfig.String("uploaddir")
	if this.uploaddir == "" {
		this.uploaddir = "upload"
	}
	this.uploaddir = this.sdir + "/" + this.uploaddir

	this.tempdir = beego.AppConfig.String("tempdir")
	if this.tempdir == "" {
		this.tempdir = "temp"
	}
	this.tempdir = this.uploaddir + "/" + this.tempdir

}
func (this *GgcmsUploadFileController) SaveFile(up models.UpInfo) string {
	if this.tempdir == "" {
		this.Prepare()
	}
	//原文件路径
	sfn := up.Realname
	if !Exist(sfn) {
		return sfn
	}

	val := this.GetSiteConfigVal("cfg_uploadmode")
	switch val {
	case "1": //七牛
		return this.saveToQiNiu(sfn, up.Name)
	}
	return this.saveFileLocal(sfn, up.Name)
}

//七牛缩略图，不做持久化
func (this *GgcmsUploadFileController) thumbQiNiu(pic string, width, height int) string {
	return pic + "?imageView2/1/w/" + strconv.Itoa(width) + "/h/" + strconv.Itoa(height)
}

//保存到七牛
func (this *GgcmsUploadFileController) saveToQiNiu(sfn, tfn string) string {

	//上传配置
	conf.ACCESS_KEY = this.GetSiteConfigVal("cfg_access_key")
	conf.SECRET_KEY = this.GetSiteConfigVal("cfg_secret_key")
	bucket := this.GetSiteConfigVal("cfg_bucket")
	link := this.GetSiteConfigVal("cfg_link_template")
	//创建一个Client
	c := kodo.New(0, nil)
	//上传文件名

	p := c.Bucket(bucket)
	for entry, _ := p.Stat(nil, tfn); entry.Hash != ""; entry, _ = p.Stat(nil, tfn) {
		ext := path.Ext(sfn)
		tfn = RandString() + ext
	}

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket + ":" + tfn,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet

	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	//err := uploader.PutFileWithoutKey(nil, &ret, token, sfn, nil)
	err := uploader.PutFile(nil, &ret, token, tfn, sfn, nil)
	//打印出错信息
	if err != nil {
		fmt.Println("io.Put failed:", err)
		//出错，保存本地
		return this.saveFileLocal(sfn, tfn)
	}
	os.Remove(sfn)
	return strings.Replace(link, "{fn}", ret.Key, -1)
}

//本地图片缩略图
func (this *GgcmsUploadFileController) thumbLocal(pic string, width, height int) string {
	img, err := imaging.Open(pic)
	if err != nil {
		return pic
	}
	thumb := imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)
	date := time.Now().Format("20060102")
	p := this.uploaddir + "/" + date
	//建文件夹
	if !Exist(p) {
		os.MkdirAll(p, os.ModePerm)
	}
	fi, _ := os.Stat(pic)
	tfn := fi.Name()
	tfn = p + "/thumb_" + tfn
	ext := path.Ext(pic)
	for Exist(tfn) {
		fn := RandString()
		tfn = p + "/" + fn + ext
	}
	imaging.Save(thumb, tfn)
	return tfn
}

//保存到本地
func (this *GgcmsUploadFileController) saveFileLocal(sfn, tfn string) string {

	date := time.Now().Format("20060102")
	p := this.uploaddir + "/" + date
	//建文件夹
	if !Exist(p) {
		os.MkdirAll(p, os.ModePerm)
	}
	tfn = p + "/" + tfn
	for Exist(tfn) {
		ext := path.Ext(tfn)
		fn := RandString()
		tfn = p + "/" + fn + ext
	}
	os.Rename(sfn, tfn)
	//CopyFile(tfn, sfn)
	return tfn
}

// @Title Post
// @Description create GgcmsAdmin
// @Param	body		body 	models.GgcmsAdmin	true		"body for GgcmsAdmin content"
// @Success 201 {int} models.GgcmsAdmin
// @Failure 403 body is empty
// @router / [post]
func (this *GgcmsUploadFileController) UploadFile() {
	msg := models.Message{1, "失败", nil}
	_, fh, _ := this.GetFile("file")
	t := this.GetString("uptype")
	//临时文件
	p := this.tempdir
	//uptype为0是临时文件，非0直接上传
	if t != "0" {
		date := time.Now().Format("20060102")
		p = this.uploaddir + "/" + date
	}
	//建文件夹
	if !Exist(p) {
		os.MkdirAll(p, os.ModePerm)
	}
	f := p + "/" + fh.Filename
	//重名处理
	for Exist(f) {
		ext := path.Ext(f)
		fn := RandString()
		f = p + "/" + fn + ext
	}
	//保存文件
	if err := this.SaveToFile("file", f); err != nil {
		msg.Data = err
		msg.Msg = err.Error()
		this.Data["json"] = msg
	} else {
		msg = models.Message{0, "成功", f}
		this.Data["json"] = msg
	}
	this.ServeJSON()
}

func (this *GgcmsUploadFileController) ImageThumb(pic string, width, height int) string {
	val := this.GetSiteConfigVal("cfg_uploadmode")
	switch val {
	case "1": //七牛
		return this.thumbQiNiu(pic, width, height)
	}
	return this.thumbLocal(pic, width, height)
}
func getUpInfo(inputid string, uplist models.UpInfos) *models.UpInfo {
	for _, up := range uplist.UpinfoList {
		if up.InputId == inputid {
			return &up
		}
	}
	return nil
}
