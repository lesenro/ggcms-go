package controllers

import (
	"errors"
	"ggcms/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/session"
)

type BaseController struct {
	beego.Controller
	cacheman    *CacheManage
	currentSite int
	cachedir    string
}

var adminpagesize int64
var navigatepages int64
var globalSessions *session.Manager
var ggcacth cache.Cache

func (this *BaseController) Prepare() {
	if sid := this.Ctx.GetCookie("currentSite"); sid == "" {
		ml, _, _ := GetAllGgcmsSites("", "", "", "site_main:1", 1, 1, false)
		if len(ml) > 0 {
			sinfo := ml[0].(models.GgcmsSites)
			this.currentSite = sinfo.Id
			this.Ctx.SetCookie("currentSite", strconv.Itoa(sinfo.Id), "/")
		}
	} else {
		this.currentSite, _ = strconv.Atoi(sid)
	}
	this.cacheman = NewCacheManage()
	this.cachedir = beego.AppConfig.String("cachedir")
	if "" == this.cachedir {
		this.cachedir = "cachedir"
	}
}
func (this *BaseController) GetSiteConfigVal(key string) string {
	cfgs := this.cacheman.CacheSiteConfigs(this.currentSite)
	return cfgs[key].Value
}
func init() {
	//初始化session
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"beegosessionID", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
	//管理页面每页记录数
	adminpagesize, _ = beego.AppConfig.Int64("adminpagesize")
	navigatepages, _ = beego.AppConfig.Int64("navigatepages")

	var err error

	//清理过期cache
	cacheinterval := 60

	if cacheinterval, _ = beego.AppConfig.Int("cacheinterval"); err != nil {
		cacheinterval = 60
	}
	ggcacth, _ = cache.NewCache("memory", `{"interval":`+strconv.Itoa(cacheinterval)+`"}`)
}
func getQueryList(strquery string) (query map[string]string, err error) {
	query = make(map[string]string)
	if v := strquery; v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				return nil, errors.New("Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	return query, nil
}
