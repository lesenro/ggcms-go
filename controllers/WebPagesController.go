package controllers

import (
	"ggcms/models"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type WebPagesController struct {
	BaseController
}

func init() {
	//注册模板函数
	//分类列表
	beego.AddFuncMap("CategoryList", CategoryList)
	//分类信息
	beego.AddFuncMap("CategoryInfo", CategoryInfo)
	//文章列表
	beego.AddFuncMap("ArticleList", ArticleList)
	//文章信息
	beego.AddFuncMap("ArticleInfo", ArticleInfo)
	//文章附件
	beego.AddFuncMap("ArticleAttachs", ArticleAttachs)
	//文章分页列表
	beego.AddFuncMap("ArticlePageList", ArticlePageList)
	//文章分页信息
	beego.AddFuncMap("ArticlePageInfo", ArticlePageInfo)
	//文章总分页数
	beego.AddFuncMap("ArticlePageCount", ArticlePageCount)
	//文章模型信息
	beego.AddFuncMap("ArticleModules", ArticleModules)
	//获取系统字典
	beego.AddFuncMap("SystemDict", SystemDict)
}

// @router / [get]
func (c *WebPagesController) WebIndex() {
	c.Data["siteid"] = c.currentSite
	c.Data["configs"] = c.getConfigs()
	c.Data["categorys"] = c.cacheman.CacheCategoryList(c.currentSite)
	tpath, spath := c.getTemplate(0, 0)
	c.Data["stylepath"] = spath
	temps := c.cacheman.CacheSiteConfigs(c.currentSite)
	cacheEnable := temps["cfg_cache_enable"].Value == "on"
	if cacheEnable {
		render := true
		cpath := c.cachedir + "/" + temps["cfg_cache_dir"].Value
		hfile := cpath + "/index.html"
		tout := temps["cfg_cache_timeout"].Value
		if finfo, err := os.Stat(hfile); err == nil {
			now := time.Now()
			var d time.Duration
			d, err = time.ParseDuration(tout + "m")
			ftime := finfo.ModTime().Add(d)
			if now.Before(ftime) || err != nil {
				c.Ctx.WriteString(ReadFileString(hfile))
				render = false
			}
		}
		if render {
			c.TplName = tpath
			if html, err := c.RenderString(); err == nil {
				if !Exist(cpath) {
					os.MkdirAll(cpath, os.ModePerm)
				}
				StringToFile(html, hfile)
			}
		}
	} else {
		c.TplName = tpath
	}
}

// @router /category/:id/?:page [get]
func (c *WebPagesController) WebCategory() {
	var id, page int
	var err error
	idStr := c.Ctx.Input.Param(":id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		c.Abort("503")
	}
	pagestr := c.Ctx.Input.Param(":page")
	page, err = strconv.Atoi(pagestr)
	if err != nil {
		page = 1
	}
	cinfo := CategoryInfo(id)
	alist, count, _ := GetAllGgcmsArticle("", "id", "desc", "categoryid:"+idStr, int64(page), int64(cinfo.Pagesize), true)
	turl := "/category/" + idStr + "/[n]"
	pages := models.GgcmsPagination{PageNum: page, PageSize: int(cinfo.Pagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["articlelist"] = alist
	c.Data["siteid"] = c.currentSite
	c.Data["configs"] = c.getConfigs()
	c.Data["categorys"] = c.cacheman.CacheCategoryList(c.currentSite)
	c.Data["categoryinfo"] = cinfo
	tpath, spath := c.getTemplate(1, *cinfo)
	c.Data["stylepath"] = spath

	temps := c.cacheman.CacheSiteConfigs(c.currentSite)
	cacheEnable := temps["cfg_cache_enable"].Value == "on"
	if cacheEnable {
		render := true
		cpath := c.cachedir + "/" + temps["cfg_cache_dir"].Value
		hfile := cpath + "/c_" + idStr + "_" + pagestr + ".html"
		tout := temps["cfg_cache_timeout"].Value
		if finfo, err := os.Stat(hfile); err == nil {
			now := time.Now()
			var d time.Duration
			d, err = time.ParseDuration(tout + "m")
			ftime := finfo.ModTime().Add(d)
			if now.Before(ftime) || err != nil {
				c.Ctx.WriteString(ReadFileString(hfile))
				render = false
			}
		}
		if render {
			c.TplName = tpath
			if html, err := c.RenderString(); err == nil {
				if !Exist(cpath) {
					os.MkdirAll(cpath, os.ModePerm)
				}
				StringToFile(html, hfile)
			}
		}
	} else {
		c.TplName = tpath
	}
}

// @router /article/:id/?:page [get]
func (c *WebPagesController) WebArticle() {
	var id, page int
	var err error
	idStr := c.Ctx.Input.Param(":id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		c.Abort("503")
	}
	pagestr := c.Ctx.Input.Param(":page")
	page, err = strconv.Atoi(pagestr)
	if err != nil {
		page = 1
	}
	ainfo := ArticleInfo(id)
	turl := "/article/" + idStr + "/[n]"
	count := ArticlePageCount(id)
	pages := models.GgcmsPagination{PageNum: page, PageSize: 1, RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	if page > 1 {
		c.Data["pageinfo"] = ArticlePageInfo(id, page)
	} else {
		p := models.GgcmsArticlePages{Articleid: id, Content: ainfo.Content, Sortid: 1, Title: ainfo.Pagetitle}
		c.Data["pageinfo"] = &p
	}
	c.Data["attachs"] = ArticleAttachs(id)
	if ainfo.Mid > 0 {
		c.Data["moduleinfo"] = ArticleModules(ainfo.Mid, id)
	}
	c.Data["pages"] = pages
	c.Data["siteid"] = c.currentSite
	c.Data["configs"] = c.getConfigs()
	c.Data["categorys"] = c.cacheman.CacheCategoryList(c.currentSite)
	c.Data["articleinfo"] = ainfo
	tpath, spath := c.getTemplate(2, *ainfo)
	c.Data["stylepath"] = spath
	temps := c.cacheman.CacheSiteConfigs(c.currentSite)
	cacheEnable := temps["cfg_cache_enable"].Value == "on"
	if cacheEnable {
		render := true
		cpath := c.cachedir + "/" + temps["cfg_cache_dir"].Value
		hfile := cpath + "/a_" + idStr + "_" + pagestr + ".html"
		tout := temps["cfg_cache_timeout"].Value
		if finfo, err := os.Stat(hfile); err == nil {
			now := time.Now()
			var d time.Duration
			d, err = time.ParseDuration(tout + "m")
			ftime := finfo.ModTime().Add(d)
			if now.Before(ftime) || err != nil {
				c.Ctx.WriteString(ReadFileString(hfile))
				render = false
			}
		}
		if render {
			c.TplName = tpath
			if html, err := c.RenderString(); err == nil {
				if !Exist(cpath) {
					os.MkdirAll(cpath, os.ModePerm)
				}
				StringToFile(html, hfile)
			}
		}
	} else {
		c.TplName = tpath
	}
}

//获取配置
func (c *WebPagesController) getConfigs() map[string]string {
	temps := c.cacheman.CacheSiteConfigs(c.currentSite)
	cfgs := make(map[string]string)
	cfgs["cfg_prefixpath"] = beego.AppConfig.String("prefixpath")
	cfgs["cfg_logo"] = temps["cfg_logo"].Value
	cfgs["cfg_basehost"] = temps["cfg_basehost"].Value
	cfgs["cfg_indexurl"] = temps["cfg_indexurl"].Value
	cfgs["cfg_webname"] = temps["cfg_webname"].Value
	cfgs["cfg_indexname"] = temps["cfg_indexname"].Value
	return cfgs
}
func (c *WebPagesController) getTemplate(tpType int, info interface{}) (tempath, stylepath string) {
	//风格路径
	stdir := beego.AppConfig.String("styledir")
	if "" == stdir {
		stdir = "styledir"
	}
	//静态文件路径
	sdir := beego.AppConfig.String("StaticDir")
	if "" == sdir {
		sdir = "static"
	}
	temps := c.cacheman.CacheSiteConfigs(c.currentSite)
	styleName := temps["cfg_default_style"].Value
	templateName := temps["cfg_template_home"].Value

	switch tpType {
	case 0: //首页
		return stdir + "/" + styleName + "/" + templateName, sdir + "/" + stdir + "/" + styleName
	case 1: //栏目
		cinfo := info.(models.GgcmsCategory)
		return stdir + "/" + cinfo.Styledir + "/" + cinfo.Ctempname, sdir + "/" + stdir + "/" + cinfo.Styledir
	case 2: //文章
		ainfo := info.(models.GgcmsArticle)
		return stdir + "/" + ainfo.Styledir + "/" + ainfo.Tempname, sdir + "/" + stdir + "/" + ainfo.Styledir
	}
	return "", ""
}

//获取url地址
func (c *WebPagesController) getUrl(tplurl string) string {
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
func CategoryList(strfields, strquery string, siteid int) (ml []interface{}) {
	if strfields == "" {
		strfields = "Id,Categoryname,Pid,Styledir,Ctempname,Atempname,Mid"
	}
	ml, _, _ = GetAllGgcmsCategory(strfields, "orderid", "asc", "ctype:1,siteid:"+strconv.Itoa(siteid), 0, 0, false)
	return
}
func CategoryInfo(id int) (v *models.GgcmsCategory) {
	var err error
	v, err = GetOneGgcmsCategory(id)
	if err != nil {
		return nil
	}
	return
}
func ArticleList(strfields, strquery, strsort, strorder string, pagenum, pagesize int64, siteid int) (ml []interface{}) {
	if strfields == "" {
		strfields = "Id,Title,Categoryid,Dateandtime"
	}
	if strsort == "" {
		strsort = "id"
		strorder = "desc"
	}
	ml, _, _ = GetAllGgcmsArticle(strfields, strsort, strorder, strquery, pagenum, pagesize, true)
	return
}
func ArticleInfo(id int) (v *models.GgcmsArticle) {
	var err error
	v, err = GetOneGgcmsArticle(id)
	if err != nil {
		return nil
	}
	//文章未审核
	if v.Categoryid < 0 {
		return nil
	}
	return
}
func ArticleAttachs(id int) (ml []interface{}) {
	var strarr []string
	attaQuery, _ := getQueryList("articleid:" + strconv.Itoa(id))
	ml, _, _ = models.GetAllGgcmsArticleAttr(attaQuery, strarr, strarr, strarr, 0, 0, false)
	return
}
func ArticlePageList(id int) (ml []interface{}) {
	var strarr []string
	pageQuery, _ := getQueryList("articleid:" + strconv.Itoa(id))
	ml, _, _ = models.GetAllGgcmsArticlePages(pageQuery, strarr, []string{"Sortid"}, []string{"asc"}, 0, 0, false)
	return
}
func ArticlePageCount(id int) (pc int64) {
	var err error
	pageCountQuery, _ := getQueryList("articleid:" + strconv.Itoa(id))
	pc, err = models.GetCountGgcmsArticlePages(pageCountQuery)
	if err != nil {
		return 1
	}
	return pc + 1
}
func ArticlePageInfo(aid, sid int) (v *models.GgcmsArticlePages) {
	if sid == 1 {
		return nil
	}
	var err error
	v, err = models.GetGgcmsArticlePagesBySortId(aid, sid)
	if err != nil {
		return nil
	}
	return
}
func ArticleModules(mid, aid int) (ml []interface{}) {
	var maps *[]orm.Params
	v, err := models.GetGgcmsModulesById(mid)
	if err != nil {
		return nil
	}
	maps, err = models.GetGgcmsModulesInfo(aid, v.Table_name)
	if err != nil {
		return nil
	}
	var mc []models.GgcmsModulesColumns
	mc, err = GetModulesColumnsByMid(mid)
	if err != nil {
		return nil
	}
	mm := *maps
	if len(mm) > 0 {
		for _, val := range mc {
			for key, v1 := range mm[0] {
				if key == val.CName {
					val.Value = v1
				}
			}
			ml = append(ml, val)
		}
	}
	return ml
}
func SystemDict(groupname, pvalue string) (ml []interface{}) {
	if groupname == "" {
		return nil
	}
	strquery := "egroup:" + groupname
	if pvalue != "" {
		strquery = strquery + ",pvalue:" + pvalue
	}
	ml, _, _ = GetAllGgcmsSysEnum("", "Orderid", "asc", strquery, 0, 0, false)
	return
}
