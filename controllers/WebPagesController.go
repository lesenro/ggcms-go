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
	siteid int
}

func init() {
	//注册模板函数
	//分类列表
	beego.AddFuncMap("CategoryList", CategoryList)
	//分类信息
	beego.AddFuncMap("CategoryInfo", CategoryInfo)
	//专题列表
	beego.AddFuncMap("TopicList", TopicList)
	//专题信息
	beego.AddFuncMap("TopicInfo", TopicInfo)
	//专题文章列表
	beego.AddFuncMap("TopicArticleList", TopicArticleList)
	//文章列表
	beego.AddFuncMap("ArticleList", ArticleList)
	//文章信息
	beego.AddFuncMap("ArticleInfo", ArticleInfo)
	//文章所属专题
	beego.AddFuncMap("ArticleTopics", ArticleTopics)
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
	//字符串拼接
	beego.AddFuncMap("StrCat", StrCat)
	//int转字符串
	beego.AddFuncMap("Itoa", strconv.Itoa)
	//获取链接url地址
	beego.AddFuncMap("GetLinkUrl", GetLinkUrl)
}

// @router / [get]
func (c *WebPagesController) WebIndex() {
	sid := c.getSiteid()
	temps := c.cacheman.CacheSiteConfigs(sid)
	cacheEnable := temps["cfg_cache_enable"].Value == "on"
	cpath := c.cachedir + "/" + temps["cfg_cache_dir"].Value
	hfile := cpath + "/index.html"
	if cacheEnable {
		isMob := c.isMobile()
		if isMob {
			hfile = cpath + "/m_index.html"
		}
		tout := temps["cfg_cache_timeout"].Value
		if finfo, err := os.Stat(hfile); err == nil {
			now := time.Now()
			var d time.Duration
			d, err = time.ParseDuration(tout + "m")
			ftime := finfo.ModTime().Add(d)
			if now.Before(ftime) || err != nil {
				//http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, hfile)
				if c.Ctx.Output.EnableGzip {
					c.Ctx.Output.Header("Content-Encoding", "gzip")
					c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
				}
				c.Ctx.ResponseWriter.Write(ReadFileBytes(hfile))
				return
			}
		}
	}
	c.Data["siteid"] = sid
	c.Data["configs"] = c.getConfigs()
	c.Data["categorys"] = c.cacheman.CacheCategoryList(sid)
	tpath, spath, tmplroot := c.getTemplate(0, 0)
	c.Data["stylepath"] = spath
	c.Data["tmplroot"] = tmplroot
	c.TplName = tpath
	if cacheEnable {
		if htmlbs, err := c.RenderBytes(); err == nil {
			if !Exist(cpath) {
				os.MkdirAll(cpath, os.ModePerm)
			}
			if c.Ctx.Output.EnableGzip {
				BytesToFile(GzipBytes(htmlbs), hfile)
			} else {
				BytesToFile(htmlbs, hfile)
			}
		}
	}
}

// @router /category/:id/?:page [get]
func (c *WebPagesController) WebCategory() {
	sid := c.getSiteid()
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
	if sid != cinfo.Siteid {
		c.Abort("404")
	}

	temps := c.cacheman.CacheSiteConfigs(sid)
	cacheEnable := temps["cfg_cache_enable"].Value == "on"
	cpath := c.cachedir + "/" + temps["cfg_cache_dir"].Value
	hfile := cpath + "/c_" + idStr + "_" + pagestr + ".html"
	if cacheEnable {
		isMob := c.isMobile()
		if isMob {
			hfile = cpath + "/m_c_" + idStr + "_" + pagestr + ".html"
		}
		tout := temps["cfg_cache_timeout"].Value
		if finfo, err := os.Stat(hfile); err == nil {
			now := time.Now()
			var d time.Duration
			d, err = time.ParseDuration(tout + "m")
			ftime := finfo.ModTime().Add(d)
			if now.Before(ftime) || err != nil {
				if c.Ctx.Output.EnableGzip {
					c.Ctx.Output.Header("Content-Encoding", "gzip")
					c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
				}
				c.Ctx.ResponseWriter.Write(ReadFileBytes(hfile))
				return
			}
		}
	}
	alist, count, _ := GetAllGgcmsArticle("Id,Title,Categoryid,Dateandtime,TitleImg,TitleThumbnailImg,AUrl", "Dateandtime", "desc", "siteid="+strconv.Itoa(sid)+",categoryid:"+idStr, int64(page), int64(cinfo.Pagesize), true)
	turl := "/category/" + idStr + "/[n]"
	pages := models.GgcmsPagination{PageNum: page, PageSize: int(cinfo.Pagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["articlelist"] = alist
	c.Data["siteid"] = sid
	c.Data["configs"] = c.getConfigs()
	c.Data["categorys"] = c.cacheman.CacheCategoryList(sid)
	c.Data["categoryinfo"] = cinfo
	tpath, spath, tmplroot := c.getTemplate(1, *cinfo)
	c.Data["stylepath"] = spath
	c.Data["tmplroot"] = tmplroot
	c.TplName = tpath
	if cacheEnable {
		if htmlbs, err := c.RenderBytes(); err == nil {
			if !Exist(cpath) {
				os.MkdirAll(cpath, os.ModePerm)
			}
			if c.Ctx.Output.EnableGzip {
				BytesToFile(GzipBytes(htmlbs), hfile)
			} else {
				BytesToFile(htmlbs, hfile)
			}
		}
	}

}

// @router /topic/:id/?:page [get]
func (c *WebPagesController) WebTopic() {
	sid := c.getSiteid()
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
	tinfo := TopicInfo(id)
	if sid != tinfo.Siteid {
		c.Abort("404")
	}

	temps := c.cacheman.CacheSiteConfigs(sid)
	cacheEnable := temps["cfg_cache_enable"].Value == "on"
	cpath := c.cachedir + "/" + temps["cfg_cache_dir"].Value
	hfile := cpath + "/t_" + idStr + "_" + pagestr + ".html"
	if cacheEnable {
		isMob := c.isMobile()
		if isMob {
			hfile = cpath + "/m_t_" + idStr + "_" + pagestr + ".html"
		}
		tout := temps["cfg_cache_timeout"].Value
		if finfo, err := os.Stat(hfile); err == nil {
			now := time.Now()
			var d time.Duration
			d, err = time.ParseDuration(tout + "m")
			ftime := finfo.ModTime().Add(d)
			if now.Before(ftime) || err != nil {
				if c.Ctx.Output.EnableGzip {
					c.Ctx.Output.Header("Content-Encoding", "gzip")
					c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
				}
				c.Ctx.ResponseWriter.Write(ReadFileBytes(hfile))
				return
			}
		}
	}
	alist, count, _ := GetAllGgcmsArticleByTopic("Id,Title,Categoryid,Dateandtime,TitleImg,TitleThumbnailImg,AUrl", "Dateandtime", "desc", "siteid:"+strconv.Itoa(sid)+",categoryid.gt:0,tid:"+idStr, int64(page), int64(tinfo.Pagesize), true)
	turl := "/topic/" + idStr + "/[n]"
	pages := models.GgcmsPagination{PageNum: page, PageSize: int(tinfo.Pagesize), RowTotal: count, UrlTemplate: turl, NavigatePages: int(navigatepages)}
	pages.CalcPages()
	c.Data["pages"] = pages
	c.Data["articlelist"] = alist
	c.Data["siteid"] = sid
	c.Data["configs"] = c.getConfigs()
	c.Data["categorys"] = c.cacheman.CacheCategoryList(sid)
	c.Data["topicinfo"] = tinfo
	tpath, spath, tmplroot := c.getTemplate(3, *tinfo)
	c.Data["stylepath"] = spath
	c.Data["tmplroot"] = tmplroot
	c.TplName = tpath
	if cacheEnable {
		if htmlbs, err := c.RenderBytes(); err == nil {
			if !Exist(cpath) {
				os.MkdirAll(cpath, os.ModePerm)
			}
			if c.Ctx.Output.EnableGzip {
				BytesToFile(GzipBytes(htmlbs), hfile)
			} else {
				BytesToFile(htmlbs, hfile)
			}
		}
	}

}

// @router /article/:id/?:page [get]
func (c *WebPagesController) WebArticle() {
	sid := c.getSiteid()

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
	if ainfo == nil || sid != ainfo.Siteid {
		c.Abort("404")
	}
	turl := "/article/" + idStr + "/[n]"

	temps := c.cacheman.CacheSiteConfigs(sid)
	cacheEnable := temps["cfg_cache_enable"].Value == "on"
	cpath := c.cachedir + "/" + temps["cfg_cache_dir"].Value
	hfile := cpath + "/a_" + idStr + "_" + pagestr + ".html"
	if cacheEnable {
		isMob := c.isMobile()
		if isMob {
			hfile = cpath + "/m_a_" + idStr + "_" + pagestr + ".html"
		}
		tout := temps["cfg_cache_timeout"].Value
		if finfo, err := os.Stat(hfile); err == nil {
			now := time.Now()
			var d time.Duration
			d, err = time.ParseDuration(tout + "m")
			ftime := finfo.ModTime().Add(d)
			if now.Before(ftime) || err != nil {
				if c.Ctx.Output.EnableGzip {
					c.Ctx.Output.Header("Content-Encoding", "gzip")
					c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
				}
				c.Ctx.ResponseWriter.Write(ReadFileBytes(hfile))
				return
			}
		}
	}
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
	c.Data["siteid"] = sid
	c.Data["configs"] = c.getConfigs()
	c.Data["categorys"] = c.cacheman.CacheCategoryList(sid)
	c.Data["articleinfo"] = ainfo
	tpath, spath, tmplroot := c.getTemplate(2, *ainfo)
	c.Data["stylepath"] = spath
	c.Data["tmplroot"] = tmplroot
	topicdal := models.GgcmsTopic{}
	c.Data["topics"], _ = topicdal.TopicsById(id)

	c.TplName = tpath
	if cacheEnable {
		c.TplName = tpath
		if htmlbs, err := c.RenderBytes(); err == nil {
			if !Exist(cpath) {
				os.MkdirAll(cpath, os.ModePerm)
			}
			if c.Ctx.Output.EnableGzip {
				BytesToFile(GzipBytes(htmlbs), hfile)
			} else {
				BytesToFile(htmlbs, hfile)
			}
		}
	}
}

//生成验证码
// @router /getcode [get]
func (c *WebPagesController) GetCode() {
	ccc := NewCheckCode()
	dc, code := ccc.GetCode()
	c.Ctx.Input.CruSession.Set("checkcode", code)
	header := c.Ctx.ResponseWriter.Header()
	header.Add("Content-Type", "image/jpeg")
	dc.EncodePNG(c.Ctx.ResponseWriter)
}

//根据域名获取站点id
func (c *WebPagesController) getSiteid() (sid int) {
	if c.siteid > 0 {
		return c.siteid
	}
	host := c.Ctx.Request.Host
	sitelist := c.cacheman.CacheSiteList()
	for _, s := range sitelist {
		sinfo := s.(models.GgcmsSites)
		if sinfo.Site_domain == host {
			c.siteid = sinfo.Id
			break
		}
		if sinfo.Site_main == 1 {
			c.siteid = sinfo.Id
		}
	}
	return c.siteid
}

//获取配置
func (c *WebPagesController) getConfigs() map[string]string {
	temps := c.cacheman.CacheSiteConfigs(c.getSiteid())
	cfgs := make(map[string]string)
	cfgs["cfg_prefixpath"] = beego.AppConfig.String("prefixpath")
	cfgs["cfg_logo"] = temps["cfg_logo"].Value
	cfgs["cfg_basehost"] = temps["cfg_basehost"].Value
	cfgs["cfg_indexurl"] = temps["cfg_indexurl"].Value
	cfgs["cfg_webname"] = temps["cfg_webname"].Value
	cfgs["cfg_indexname"] = temps["cfg_indexname"].Value
	return cfgs
}
func (c *WebPagesController) getTemplate(tpType int, info interface{}) (tempath, stylepath, tmplroot string) {
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
	temps := c.cacheman.CacheSiteConfigs(c.getSiteid())
	styleName := temps["cfg_default_style"].Value
	templateName := temps["cfg_template_home"].Value
	mhometemplate := temps["cfg_template_m_home"].Value
	isMob := c.isMobile()
	tmplroot = stdir + "/" + styleName
	switch tpType {
	case 0: //首页
		tempath = stdir + "/" + styleName + "/" + templateName
		stylepath = sdir + "/" + stdir + "/" + styleName
		if isMob && mhometemplate != "" {
			tempath = stdir + "/" + styleName + "/" + mhometemplate
		}
	case 1: //栏目
		cinfo := info.(models.GgcmsCategory)
		tempath = stdir + "/" + cinfo.Styledir + "/" + cinfo.Ctempname
		stylepath = sdir + "/" + stdir + "/" + cinfo.Styledir
		if isMob && cinfo.Mob_list_temp != "" {
			tempath = stdir + "/" + cinfo.Styledir + "/" + cinfo.Mob_list_temp
		}

	case 2: //文章
		ainfo := info.(models.GgcmsArticle)
		tempath = stdir + "/" + ainfo.Styledir + "/" + ainfo.Tempname
		stylepath = sdir + "/" + stdir + "/" + ainfo.Styledir
		if isMob && ainfo.Mob_tempname != "" {
			tempath = stdir + "/" + ainfo.Styledir + "/" + ainfo.Mob_tempname
		}
	case 3: //专题
		tinfo := info.(models.GgcmsTopic)
		tempath = stdir + "/" + tinfo.Styledir + "/" + tinfo.Tempname
		stylepath = sdir + "/" + stdir + "/" + tinfo.Styledir
		if isMob && tinfo.Mob_tempname != "" {
			tempath = stdir + "/" + tinfo.Styledir + "/" + tinfo.Mob_tempname
		}
	}
	return
}

//是否是移动端
func (c *WebPagesController) isMobile() bool {
	//由客户端控制Cookie，是否浏览移动页面
	if c.Ctx.GetCookie("mob_enable") == "off" {
		return false
	}
	temps := c.cacheman.CacheSiteConfigs(c.getSiteid())
	mobEnable := temps["cfg_mob_enable"].Value == "on"
	//未开启移动端识别
	if !mobEnable {
		return false
	}
	mobflag := temps["cfg_mob_flag"].Value
	//移动端识别码为空
	if strings.TrimSpace(mobflag) == "" {
		return false
	}

	//ua := c.Ctx.Request.Header.Get("User-Agent")
	ua := c.Ctx.Input.UserAgent()
	mobreg, _ := regexp.Compile("(?i:" + mobflag + ")")
	return mobreg.MatchString(ua)
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
	if strfields == "" && strquery == "" {
		cm := NewCacheManage()
		return cm.CacheCategoryList(siteid)
	}
	if strfields == "" {
		strfields = "Id,Categoryname,Pid,Styledir,Ctempname,Atempname,Mid"
	}
	query := "ctype:1,siteid:" + strconv.Itoa(siteid)
	if strquery != "" {
		query = query + "," + strquery
	}
	ml, _, _ = GetAllGgcmsCategory(strfields, "orderid", "asc", query, 0, 0, false)
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
func TopicList(groupkey string, siteid int, ids ...int) (ml []interface{}) {
	querystr := make([]string, 0)
	querystr = append(querystr, "siteid:"+strconv.Itoa(siteid))
	if groupkey != "" {
		querystr = append(querystr, "groupkey:"+groupkey)
	}
	if len(ids) > 0 {
		arr_id := make([]string, 0)
		for _, id := range ids {
			arr_id = append(arr_id, strconv.Itoa(id))
		}
		querystr = append(querystr, "id.in:"+strings.Join(arr_id, "|"))
	}
	ml, _, _ = GetAllGgcmsTopic("Id,Topic,Groupkey,Logo,Extattrib,Turl", "id", "desc", strings.Join(querystr, ","), 0, 0, false)
	return ml
}
func TopicInfo(id int) (v *models.GgcmsTopic) {
	var err error
	v, err = GetOneGgcmsTopic(id)
	if err != nil {
		return nil
	}
	return
}
func ArticleList(strfields, strquery, strsort, strorder string, pagenum, pagesize int64, siteid int) (ml []interface{}) {
	if strfields == "" {
		strfields = "Id,Title,Categoryid,Dateandtime,TitleImg,TitleThumbnailImg,AUrl"
	}
	query := "siteid:" + strconv.Itoa(siteid)

	if !strings.Contains(strings.ToLower(strquery), "categoryid") {
		query = query + ",categoryid.gt:0"
	}
	if strquery != "" {
		query = query + "," + strquery
	}
	if strsort == "" && strorder == "" {
		strsort = "Dateandtime"
		strorder = "desc"
	}
	ml, _, _ = GetAllGgcmsArticle(strfields, strsort, strorder, query, pagenum, pagesize, true)
	return
}

//专题文章列表
func TopicArticleList(strfields, strquery, strsort, strorder string, pagenum, pagesize int64, siteid int) (ml []interface{}) {
	if strfields == "" {
		strfields = "Id,Title,Categoryid,Dateandtime,TitleImg,TitleThumbnailImg,AUrl"
	}

	query := "siteid:" + strconv.Itoa(siteid)

	if !strings.Contains(strings.ToLower(strquery), "categoryid") {
		query = query + ",categoryid.gt:0"
	}
	if strquery != "" {
		query = query + "," + strquery
	}
	if strsort == "" && strorder == "" {
		strsort = "Dateandtime"
		strorder = "desc"
	}
	ml, _, _ = GetAllGgcmsArticleByTopic(strfields, strsort, strorder, query, pagenum, pagesize, false)
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
func ArticleTopics(id, siteid int) (ml []interface{}) {
	topic := models.GgcmsTopic{}
	tids, _ := topic.TopicsById(id)
	return TopicList("", siteid, tids...)
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
func StrCat(vs ...string) string {
	return strings.Join(vs, "")
}

//itype:0首页，1栏目，2文章
//id:对应文章或栏目的id
//sid:网站id，非0，生成包含域名的链接
func GetLinkUrl(itype, id, sid int, scheme, jumpUrl string) string {
	if jumpUrl != "" {
		return jumpUrl
	}
	dm := ""
	if scheme == "" {
		scheme = "http://"
	}
	if sid > 0 {
		cm := NewCacheManage()
		slist := cm.CacheSiteList()
		for _, val := range slist {
			sinfo := val.(models.GgcmsSites)
			if sid == sinfo.Id {
				dm = scheme + sinfo.Site_domain
				break
			}
		}
	}
	url := ""
	switch itype {
	case 0:
		url = dm + "/"
	case 1:
		url = dm + "/category/" + strconv.Itoa(id) + "/"
	case 2:
		url = dm + "/article/" + strconv.Itoa(id) + "/"
	case 3:
		url = dm + "/topic/" + strconv.Itoa(id) + "/"
	}
	return url
}
