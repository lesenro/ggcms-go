package controllers

import (
	"encoding/json"
	"ggcms/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type AdminPagesController struct {
	BaseController
}

// @router / [get]
func (c *AdminPagesController) AdminMain() {
	//temps := c.cacheman.CacheSystemConfigs()
	cfgs := make(map[string]string)
	cfgs["cfg_prefixpath"] = beego.AppConfig.String("prefixpath")
	c.Data["configs"] = cfgs
	c.TplName = "admin/index.html"
}

// @router /home.html [get]
func (c *AdminPagesController) AdminHome() {
	c.TplName = "admin/home.html"
}

//系统字典
// @router /systemdict.html [get]
func (c *AdminPagesController) AdminSystemDict() {
	var tplName = "systemdict.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}

	qsstr := c.GetString("qs")
	var qs models.Query
	if err := json.Unmarshal([]byte(qsstr), &qs); err != nil {
		qs.QueryString = "issystem:0"
	} else {
		qs.QueryString = "issystem:0," + qs.QueryString
	}
	beego.Alert(qs.QueryString)
	turl := c.getUrl(tplName)
	var count int64
	count = 0
	list, count, _ := GetAllGgcmsSysEnum("", "id", "desc", qs.QueryString, int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//系统字典
// @router /systemdictadd.html [get]
func (c *AdminPagesController) AdminSystemDictAdd() {
	var tplName = "systemdictadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	groups, _ := models.GetGgcmsSysEnumGroups()
	c.Data["groups"] = groups
	if info, err := GetOneGgcmsSysEnum(id); err != nil {
		c.Data["info"] = models.GgcmsSysEnum{}
	} else {
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

//栏目设置
// @router /category.html [get]
func (c *AdminPagesController) AdminCategory() {
	var tplName = "category.html"
	if c.currentSite > 0 {
		c.Data["list"] = c.cacheman.CacheCategoryList(c.currentSite)
	} else {
		c.Data["list"] = nil
	}

	c.TplName = "admin/" + tplName
}

//栏目设置
// @router /categoryadd.html [get]
func (c *AdminPagesController) AdminCategoryAdd() {
	var tplName = "categoryadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	c.Data["list"] = c.cacheman.CacheCategoryList(c.currentSite)
	c.Data["style"], _ = styleAndTemplate()
	c.Data["modules"], _, _ = GetAllGgcmsModules("", "", "", "", 0, 0, false)
	if info, err := GetOneGgcmsCategory(id); err != nil {
		temps := c.cacheman.CacheSiteConfigs(c.currentSite)
		imgw, _ := strconv.Atoi(temps["cfg_ddimg_width"].Value)
		imgh, _ := strconv.Atoi(temps["cfg_ddimg_height"].Value)
		tmplview := temps["cfg_template_view"].Value
		tmpllist := temps["cfg_template_list"].Value
		mtmplview := temps["cfg_template_m_view"].Value
		mtmpllist := temps["cfg_template_m_list"].Value
		tmplstyle := temps["cfg_default_style"].Value
		c.Data["info"] = models.GgcmsCategory{
			Pagesize:      int(adminpagesize),
			NavPages:      int(navigatepages),
			Imgwidth:      imgw,
			Imgheight:     imgh,
			Styledir:      tmplstyle,
			Ctempname:     tmpllist,
			Atempname:     tmplview,
			Mob_list_temp: mtmpllist,
			Mob_view_temp: mtmplview,
			Ctype:         1,
		}
	} else {
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

//文章管理
// @router /articlelist.html [get]
func (c *AdminPagesController) AdminArticleList() {
	var tplName = "articlelist.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	atype := 1
	if v, err := c.GetInt("type"); err == nil {
		atype = v
	}
	csite := "siteid:" + strconv.Itoa(c.currentSite)
	typeof := "categoryid.gt:0"
	if atype == 0 {
		typeof = "categoryid:0"
	} else if atype < 0 {
		typeof = "categoryid.lt:0"
	}
	qsstr := c.GetString("qs")
	var qs models.Query
	if err := json.Unmarshal([]byte(qsstr), &qs); err != nil {
		qs.QueryString = csite + "," + typeof
	} else {
		qs.QueryString = csite + "," + typeof + "," + qs.QueryString
	}
	turl := c.getUrl(tplName)
	var count int64
	count = 0
	list, count, _ := GetAllGgcmsArticle("Id,Title,Categoryid,Dateandtime,Siteid,AUrl", "Dateandtime", "desc", qs.QueryString, int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["list"] = list
	c.Data["atype"] = atype
	c.Data["clist"] = c.cacheman.CacheCategoryList(c.currentSite)
	c.TplName = "admin/" + tplName
}

//文章添加
// @router /articleadd.html [get]
func (c *AdminPagesController) AdminArticleAdd() {
	var tplName = "articleadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	c.Data["alltopic"], _, _ = GetAllGgcmsTopic("Id,Topic", "id", "desc", "siteid:"+strconv.Itoa(c.currentSite), 0, 0, false)
	c.Data["list"] = c.cacheman.CacheCategoryList(c.currentSite)
	c.Data["style"], _ = styleAndTemplate()
	var strarr []string
	attaQuery, _ := getQueryList("articleid:" + strconv.Itoa(id))
	c.Data["attachs"], _, _ = models.GetAllGgcmsArticleAttr(attaQuery, strarr, strarr, strarr, 0, 0, false)
	var topicids []int
	if info, err := GetOneGgcmsArticle(id); err != nil {
		c.Data["info"] = models.GgcmsArticle{}
		c.Data["pagecount"] = 1
		c.Data["minfo"] = nil
	} else {
		c.Data["info"] = info
		pageCountQuery, _ := getQueryList("articleid:" + strconv.Itoa(info.Id))
		pc, _ := models.GetCountGgcmsArticlePages(pageCountQuery)
		c.Data["pagecount"] = pc + 1
		c.Data["minfo"] = GetModulesInfo(info.Mid, info.Id)
		topicdal := models.GgcmsTopic{}
		topicids, _ = topicdal.TopicsById(id)
	}
	c.Data["topic"] = topicids
	c.TplName = "admin/" + tplName
}

//站点设置
// @router /sitesetting.html [get]
func (c *AdminPagesController) AdminSiteSetting() {
	var tplName = "sitesetting.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	turl := c.getUrl(tplName)
	var count int64
	count = 0

	list, count, _ := GetAllGgcmsSites("", "", "", "", int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//站点设置
// @router /sitesettingadd.html [get]
func (c *AdminPagesController) AdminSiteSettingAdd() {
	var tplName = "sitesettingadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	if info, err := GetOneGgcmsSites(id); err != nil {
		c.Data["info"] = models.GgcmsSites{}
	} else {
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

//风格设置
// @router /styledesign.html [get]
func (c *AdminPagesController) AdminStyles() {
	var tplName = "styledesign.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	turl := c.getUrl(tplName)
	var count int64
	count = 0

	list, count, _ := GetAllGgcmsStyles("", "", "", "", int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//风格设置
// @router /styledesignadd.html [get]
func (c *AdminPagesController) AdminStylesAdd() {
	var tplName = "styledesignadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	if info, err := GetOneGgcmsStyles(id); err != nil {
		c.Data["info"] = models.GgcmsStyles{}
	} else {
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

//模板设置
// @router /styletemplate.html [get]
func (c *AdminPagesController) AdminTemplate() {
	var tplName = "styletemplate.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	if info, err := GetOneGgcmsStyles(id); err != nil {
		c.Data["info"] = models.GgcmsStyles{}
	} else {
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

//静态文件设置
// @router /stylestaticfile.html [get]
func (c *AdminPagesController) AdminStaticFile() {
	var tplName = "stylestaticfile.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	if info, err := GetOneGgcmsStyles(id); err != nil {
		c.Data["info"] = models.GgcmsStyles{}
	} else {
		sdir := beego.AppConfig.String("StaticDir")
		if "" == sdir {
			sdir = "static"
		}
		stdir := beego.AppConfig.String("styledir")
		if "" == stdir {
			stdir = "styledir"
		}
		c.Data["styledir"] = sdir + "/" + stdir
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

// @router /userlevel.html [get]
func (c *AdminPagesController) AdminUserLevel() {
	var tplName = "userlevel.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	turl := c.getUrl(tplName)
	strquery := "egroup:userlevel"
	var count int64
	count = 0
	list, count, _ := GetAllGgcmsSysEnum("", "", "", strquery, int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//站点设置
// @router /userleveladd.html [get]
func (c *AdminPagesController) AdminUserLevelAdd() {
	var tplName = "userleveladd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	if info, err := GetOneGgcmsSysEnum(id); err != nil {
		c.Data["info"] = models.GgcmsSysEnum{Egroup: "userlevel"}
	} else {
		c.Data["info"] = info
	}
	//	strquery := "egroup:userlevel"
	c.TplName = "admin/" + tplName
}

//管理员分组
// @router /admingroup.html [get]
func (c *AdminPagesController) AdminGroup() {
	var tplName = "admingroup.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	turl := c.getUrl(tplName)
	strquery := "egroup:admingroup,issystem:1"
	var count int64
	count = 0
	list, count, _ := GetAllGgcmsSysEnum("", "", "", strquery, int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//管理员分组设置
// @router /admingroupadd.html [get]
func (c *AdminPagesController) AdminGroupAdd() {
	var tplName = "admingroupadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	if info, err := GetOneGgcmsSysEnum(id); err != nil {
		c.Data["info"] = models.GgcmsSysEnum{Egroup: "admingroup", Issystem: 1}
	} else {
		c.Data["info"] = info
	}
	//	strquery := "egroup:admingroup"
	c.TplName = "admin/" + tplName
}

//管理员管理
// @router /adminlist.html [get]
func (c *AdminPagesController) AdminList() {
	var tplName = "adminlist.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	turl := c.getUrl(tplName)
	var count int64
	count = 0
	list, count, _ := GetAllGgcmsAdmin("", "", "", "", int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	admingroup, _, _ := GetAllGgcmsSysEnum("Id,Ename,Evalue", "", "", "egroup:admingroup", 0, 0, false)
	uid, ok := c.Ctx.Input.Session("uid").(int)
	if !ok {
		uid = 0
	}
	c.Data["userid"] = uid
	c.Data["pages"] = pages
	c.Data["admingroup"] = admingroup
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//管理员设置
// @router /adminadd.html [get]
func (c *AdminPagesController) AdminAdd() {
	var tplName = "adminadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	admingroup, _, _ := GetAllGgcmsSysEnum("Id,Ename,Evalue", "Orderid", "asc", "egroup:admingroup", 0, 0, false)
	c.Data["admingroup"] = admingroup

	if info, err := GetOneGgcmsAdmin(id); err != nil {
		c.Data["info"] = models.GgcmsAdmin{}
	} else {
		info.Pwd = ""
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

//专题管理
// @router /topiclist.html [get]
func (c *AdminPagesController) AdminTopicList() {
	var tplName = "topiclist.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	turl := c.getUrl(tplName)
	var count int64
	count = 0
	list, count, _ := GetAllGgcmsTopic("Id,Topic,Groupkey", "id", "desc", "siteid:"+strconv.Itoa(c.currentSite), int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	topic := models.GgcmsTopic{}
	groups, _ := topic.GetGroups()
	c.Data["pages"] = pages
	c.Data["groups"] = groups
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//专题设置
// @router /topicadd.html [get]
func (c *AdminPagesController) AdminTopicAdd() {
	var tplName = "topicadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	topic := models.GgcmsTopic{}
	groups, _ := topic.GetGroups()
	c.Data["groups"] = groups
	c.Data["style"], _ = styleAndTemplate()

	if info, err := GetOneGgcmsTopic(id); err != nil {
		topic := models.GgcmsTopic{
			Pagesize: int(adminpagesize),
		}
		c.Data["info"] = topic
	} else {
		c.Data["info"] = info
	}
	c.TplName = "admin/" + tplName
}

//系统设置
// @router /systemconfigs.html [get]
func (c *AdminPagesController) AdminSystemConfigs() {
	var tplName = "systemconfigs.html"
	c.Data["configs"] = c.cacheman.CacheSystemConfigs()
	c.TplName = "admin/" + tplName
}

//模块设计
// @router /modulesdesign.html [get]
func (c *AdminPagesController) AdminModulesDesign() {
	var tplName = "modulesdesign.html"
	pagenum := 1
	if v, err := c.GetInt("pagenum"); err == nil {
		pagenum = v
	}
	turl := c.getUrl(tplName)
	var count int64
	count = 0

	list, count, _ := GetAllGgcmsModules("", "", "", "", int64(pagenum), adminpagesize, true)
	pages := models.GgcmsPagination{PageNum: pagenum, PageSize: int(adminpagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["list"] = list
	c.TplName = "admin/" + tplName
}

//模块编辑
// @router /modulesdesignadd.html [get]
func (c *AdminPagesController) AdminModulesDesignAdd() {
	var tplName = "modulesdesignadd.html"
	id := 0
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}
	var info models.GgcmsModulesPostData
	if mInfo, err := GetOneGgcmsModules(id); err != nil {
		info.GgcmsModules = models.GgcmsModules{}
	} else {
		info.GgcmsModules = *mInfo
	}
	info.Columns, _ = GetModulesColumnsByMid(id)
	c.Data["info"] = info
	c.TplName = "admin/" + tplName
}

//网站设置
// @router /siteconfigs.html [get]
func (c *AdminPagesController) AdminSiteConfigs() {
	var tplName = "siteconfigs.html"
	c.Data["configs"] = c.cacheman.CacheSiteConfigs(c.currentSite)
	c.TplName = "admin/" + tplName
}

// @router /dashboard.html [get]
func (c *AdminPagesController) AdminDashboard() {
	c.TplName = "admin/dashboard.html"
}

// @router /ui_bootstrap.html [get]
func (c *AdminPagesController) AdminBtp() {
	c.TplName = "admin/ui_bootstrap.html"
}

// @router /login.html [get]
func (c *AdminPagesController) Admin_Login() {
	cfgs := make(map[string]string)
	cfgs["cfg_prefixpath"] = beego.AppConfig.String("prefixpath")
	c.Data["configs"] = cfgs

	c.TplName = "admin/login.html"
}

//头尾模板
// @router /tpl/:key [get]
func (c *AdminPagesController) Tpls() {
	key := c.Ctx.Input.Param(":key")
	c.TplName = "admin/tpl/" + key
}
func (c *AdminPagesController) getUrl(tplurl string) string {
	turl := c.Ctx.Request.URL.String()
	urlArr := strings.Split(turl, "?")
	if len(urlArr) == 2 {
		turl = urlArr[1]
	}
	reg := regexp.MustCompile(`pagenum=[\d|undefined]+`)
	if reg.MatchString(turl) {
		turl = reg.ReplaceAllString(turl, "pagenum=[n]")
	} else {
		if turl == "" {
			turl = "pagenum=[n]"
		} else {
			turl = turl + "&pagenum=[n]"
		}
	}
	turl = "#/" + tplurl + "?" + turl
	return turl
}
