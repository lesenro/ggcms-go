package controllers

import (
	"ggcms/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

const (
	cnSystemConfig = "SystemConfig"
	cnSiteConfig   = "SiteConfig"
	cnCategoryList = "CategoryList"
	cnSiteList     = "SiteList"
	cnUserPowers   = "UserPowers"
)

// oprations for CacheManage
type CacheManage struct {
	cacheTimeOut time.Duration
}

//初始化CacheManage管理器
func NewCacheManage() *CacheManage {
	cm := &CacheManage{}
	cm.Init()
	return cm
}
func (c *CacheManage) ClearByKey(key string) {
	ggcacth.Delete(key)
}

func (this *CacheManage) Init() {
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
nilerr:
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
	if imap == nil {
		c.ClearByKey(cacheName)
		goto nilerr
	}

	return imap.(map[string]models.GgcmsSystemConfigs)
}
func (c *CacheManage) CacheSiteConfigs(siteid int) map[string]models.GgcmsSystemConfigs {
	cacheName := cnSiteConfig + "_" + strconv.Itoa(siteid)
nilerr:
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
	if imap == nil {
		c.ClearByKey(cacheName)
		goto nilerr
	}
	return imap.(map[string]models.GgcmsSystemConfigs)
}
func (c *CacheManage) CacheCategoryList(siteid int) []interface{} {
	cacheName := cnCategoryList + "_" + strconv.Itoa(siteid)
nilerr:
	if !ggcacth.IsExist(cacheName) {
		ml, _, _ := GetAllGgcmsCategory("Id,Categoryname,Pid,Styledir,Ctempname,Atempname,Mob_list_temp,Mob_view_temp,Mid,Curl", "orderid", "asc", "ctype:1,siteid:"+strconv.Itoa(siteid), 0, 0, false)
		ggcacth.Put(cacheName, ml, c.cacheTimeOut)
	}
	imap := ggcacth.Get(cacheName)
	if imap == nil {
		c.ClearByKey(cacheName)
		goto nilerr
	}
	return imap.([]interface{})
}
func (c *CacheManage) CacheSiteList() []interface{} {
	cacheName := cnSiteList
nilerr:
	if !ggcacth.IsExist(cacheName) {
		ml, _, _ := GetAllGgcmsSites("", "", "", "", 0, 0, false)
		ggcacth.Put(cacheName, ml, c.cacheTimeOut)
	}
	imap := ggcacth.Get(cacheName)
	if imap == nil {
		c.ClearByKey(cacheName)
		goto nilerr
	}
	return imap.([]interface{})
}

func (c *CacheManage) CacheUserPowers(utype string) map[string]int {
	cacheName := cnUserPowers + "_" + utype
nilerr:
	if !ggcacth.IsExist(cacheName) {
		qs := make(map[string]string)
		qs["usertype"] = utype
		//fs := strings.Split("Id,Url,Params,Method,Datapower", ",")
		adminpwd := models.AdminPowers{}
		ml, _, _ := adminpwd.GetAll(qs, make([]string, 0), make([]string, 0), make([]string, 0), 0, 0, false)
		pwlist := make(map[string]int)
		for _, val := range ml {
			pw := val.(models.AdminPowers)
			pwlist[pw.Url] = pw.Id
		}
		ggcacth.Put(cacheName, pwlist, c.cacheTimeOut)
	}
	imap := ggcacth.Get(cacheName)
	if imap == nil {
		c.ClearByKey(cacheName)
		goto nilerr
	}
	return imap.(map[string]int)
}
