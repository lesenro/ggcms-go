package controllers

import (
	"ggcms/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

const (
	cnSystemConfig = "SystemConfig"
	cnSiteConfig   = "SiteConfig_"
	cnCategoryList = "CategoryList"
)

// oprations for CacheManage
type CacheManage struct {
	beego.Controller
	cacheTimeOut time.Duration
}

// @Title Get
// @Description get ClearAll
// @Success 200
// @Failure 403
// @router / [get]
func (c *CacheManage) CacheClearAll() {
	msg := models.Message{0, "成功", nil}
	if err := ggcacth.ClearAll(); err != nil {
		msg = models.Message{1, err.Error(), err}
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
func (c *CacheManage) ClearByKey(key string) {
	ggcacth.Delete(key)
}
func (this *CacheManage) Prepare() {
	var err error
	//cache超时时间为20分钟
	var ctimeout int64
	if ctimeout, err = beego.AppConfig.Int64("cachetimeout"); err != nil {
		this.cacheTimeOut = 20 * time.Minute
	} else {
		this.cacheTimeOut = time.Duration(ctimeout)
	}
}

func (c *CacheManage) CacheSystemConfigs() map[string]models.GgcmsSystemConfigs {
	cacheName := cnSystemConfig
	if !ggcacth.IsExist(cacheName) {
		ml := models.GetSystemConfigsBySiteId(0)
		cfgs := make(map[string]models.GgcmsSystemConfigs)
		for _, item := range ml {
			if v, ok := item.(models.GgcmsSystemConfigs); ok {
				cfgs[v.Varname] = v
			}
		}
		ggcacth.Put(cacheName, cfgs, c.cacheTimeOut)
	}
	imap := ggcacth.Get(cacheName)
	return imap.(map[string]models.GgcmsSystemConfigs)
}
func (c *CacheManage) CacheSiteConfigs(siteid int) map[string]models.GgcmsSystemConfigs {
	cacheName := cnSiteConfig + strconv.Itoa(siteid)
	if !ggcacth.IsExist(cacheName) {
		ml := models.GetSystemConfigsBySiteId(siteid)
		cfgs := make(map[string]models.GgcmsSystemConfigs)
		for _, item := range ml {
			if v, ok := item.(models.GgcmsSystemConfigs); ok {
				cfgs[v.Varname] = v
			}
		}
		ggcacth.Put(cacheName, cfgs, c.cacheTimeOut)
	}
	imap := ggcacth.Get(cacheName)
	return imap.(map[string]models.GgcmsSystemConfigs)
}
func (c *CacheManage) CacheCategoryList(siteid int) []interface{} {
	cacheName := cnCategoryList
	if !ggcacth.IsExist(cacheName) {
		ml, _, _ := GetAllGgcmsCategory("Id,Categoryname,Pid,Styledir,Ctempname,Atempname,Mid", "orderid", "asc", "ctype:1,siteid:"+strconv.Itoa(siteid), 0, 0, false)
		ggcacth.Put(cacheName, ml, c.cacheTimeOut)
	}
	imap := ggcacth.Get(cacheName)
	return imap.([]interface{})
}
