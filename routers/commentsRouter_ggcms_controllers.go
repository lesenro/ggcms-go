package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminMain",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminHome",
			`/home.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminError",
			`/error.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminSystemDict",
			`/systemdict.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminSystemDictAdd",
			`/systemdictadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminCategory",
			`/category.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminCategoryAdd",
			`/categoryadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminArticleList",
			`/articlelist.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminArticleAdd",
			`/articleadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminSiteSetting",
			`/sitesetting.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminSiteSettingAdd",
			`/sitesettingadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminStyles",
			`/styledesign.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminStylesAdd",
			`/styledesignadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminTemplate",
			`/styletemplate.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminStaticFile",
			`/stylestaticfile.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminUserLevel",
			`/userlevel.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminUserLevelAdd",
			`/userleveladd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminGroup",
			`/admingroup.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminGroupAdd",
			`/admingroupadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminPowerSet",
			`/adminpowerset.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminList",
			`/adminlist.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminAdd",
			`/adminadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminTopicList",
			`/topiclist.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminTopicAdd",
			`/topicadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminSystemConfigs",
			`/systemconfigs.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminModulesDesign",
			`/modulesdesign.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminModulesDesignAdd",
			`/modulesdesignadd.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminSiteConfigs",
			`/siteconfigs.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminDashboard",
			`/dashboard.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"AdminBtp",
			`/ui_bootstrap.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"Admin_Login",
			`/login.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:AdminPagesController"],
		beego.ControllerComments{
			"Tpls",
			`/tpl/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"],
		beego.ControllerComments{
			"GetOneByName",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminPowersController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsAdminPowersController"],
		beego.ControllerComments{
			"Edit",
			`/edit/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"GetPageInfo",
			`/pageinfo`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"Auditing",
			`/auditing`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsArticleController"],
		beego.ControllerComments{
			"UnAuditing",
			`/unauditing`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"],
		beego.ControllerComments{
			"SaveSort",
			`/savesort`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsCategoryController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsMemberController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsModulesController"],
		beego.ControllerComments{
			"GetModulesColsByMid",
			`/getcolumns/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsPowersController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSitesController"],
		beego.ControllerComments{
			"UpdateConfig",
			`/updateconfig/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"ImportStyle",
			`/import`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"ExportStyle",
			`/export/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"TemplateList",
			`/template/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"UploadTemplate",
			`/uploadtemplate`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"DeleteTemplate",
			`/deletetemplate`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"RenameTemplate",
			`/renametemplate`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"GetTemplate",
			`/gettemplate`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"SaveTemplate",
			`/savetemplate`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"UploadStatic",
			`/uploadstatic`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"DeleteStatic",
			`/staticdelete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"RenameStatic",
			`/renamestatic`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"GetStaticFile",
			`/getstaticfile`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"SaveStaticFile",
			`/savestaticfile`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"StaticList",
			`/staticfile`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"GetAllData",
			`/all`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsStylesController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"],
		beego.ControllerComments{
			"GetGroups",
			`/groups`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysEnumController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSysLogsController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"],
		beego.ControllerComments{
			"Update",
			`/update`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsSystemConfigsController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTasksController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"],
		beego.ControllerComments{
			"Add",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"],
		beego.ControllerComments{
			"Edit",
			`/edit`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsTopicController"],
		beego.ControllerComments{
			"Delete",
			`/delete`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:GgcmsUploadFileController"] = append(beego.GlobalControllerRouter["ggcms/controllers:GgcmsUploadFileController"],
		beego.ControllerComments{
			"UploadFile",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:MainController"] = append(beego.GlobalControllerRouter["ggcms/controllers:MainController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:ToolsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:ToolsController"],
		beego.ControllerComments{
			"CacheClearAll",
			`/clearcacth`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:ToolsController"] = append(beego.GlobalControllerRouter["ggcms/controllers:ToolsController"],
		beego.ControllerComments{
			"GetAdminPowers",
			`/getpowers`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"],
		beego.ControllerComments{
			"WebIndex",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"],
		beego.ControllerComments{
			"WebCategory",
			`/category/:id/?:page`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"],
		beego.ControllerComments{
			"WebTopic",
			`/topic/:id/?:page`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"],
		beego.ControllerComments{
			"WebArticle",
			`/article/:id/?:page`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"] = append(beego.GlobalControllerRouter["ggcms/controllers:WebPagesController"],
		beego.ControllerComments{
			"GetCode",
			`/getcode`,
			[]string{"get"},
			nil})

}
