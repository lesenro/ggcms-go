// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ggcms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		//beego.NSBefore(controllers.AdminAuth),
		beego.NSNamespace("/ggcms_admin",
			beego.NSInclude(
				&controllers.GgcmsAdminController{},
			),
		),

		beego.NSNamespace("/ggcms_admin_powers",
			beego.NSInclude(
				&controllers.GgcmsAdminPowersController{},
			),
		),

		beego.NSNamespace("/ggcms_article",
			beego.NSInclude(
				&controllers.GgcmsArticleController{},
			),
		),

		beego.NSNamespace("/ggcms_category",
			beego.NSInclude(
				&controllers.GgcmsCategoryController{},
			),
		),

		beego.NSNamespace("/ggcms_member",
			beego.NSInclude(
				&controllers.GgcmsMemberController{},
			),
		),

		beego.NSNamespace("/ggcms_powers",
			beego.NSInclude(
				&controllers.GgcmsPowersController{},
			),
		),
		beego.NSNamespace("/ggcms_sites",
			beego.NSInclude(
				&controllers.GgcmsSitesController{},
			),
		),
		beego.NSNamespace("/ggcms_sys_enum",
			beego.NSInclude(
				&controllers.GgcmsSysEnumController{},
			),
		),

		beego.NSNamespace("/ggcms_sys_logs",
			beego.NSInclude(
				&controllers.GgcmsSysLogsController{},
			),
		),

		beego.NSNamespace("/ggcms_tasks",
			beego.NSInclude(
				&controllers.GgcmsTasksController{},
			),
		),

		beego.NSNamespace("/ggcms_topic",
			beego.NSInclude(
				&controllers.GgcmsTopicController{},
			),
		),
		beego.NSNamespace("/system_configs",
			beego.NSInclude(
				&controllers.GgcmsSystemConfigsController{},
			),
		),
		beego.NSNamespace("/ggcms_cache",
			beego.NSInclude(
				&controllers.CacheManage{},
			),
		),
		beego.NSNamespace("/ggcms_modules",
			beego.NSInclude(
				&controllers.GgcmsModulesController{},
			),
		),
		beego.NSNamespace("/ggcms_upload",
			beego.NSInclude(
				&controllers.GgcmsUploadFileController{},
			),
		),
		beego.NSNamespace("/ggcms_styles",
			beego.NSInclude(
				&controllers.GgcmsStylesController{},
			),
		),
	)
	nsAdmin := beego.NewNamespace("/"+beego.AppConfig.String("adminpath"),
		//beego.NSBefore(controllers.AdminPagesController),
		beego.NSInclude(
			&controllers.AdminPagesController{},
		),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(nsAdmin)
	beego.Include(&controllers.WebPagesController{})

}
