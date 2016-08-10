/* Setup Rounting For All Pages */
GgcmsApp.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
    // Redirect any unmatched url
    $urlRouterProvider.otherwise("/home.html");

    $stateProvider
    // Dashboard
        .state('home', {
        url: "/home.html",
        templateUrl: "home.html",
        data: {
            pageTitle: 'Admin Dashboard Template'
        },
        controller: "DashboardController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load({
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before a LINK element with this ID. Dynamic CSS files must be loaded between core and theme css files
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/morris/morris.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/morris/morris.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/morris/raphael-min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/jquery.sparkline.min.js',

                        ggcmsCfg.cfg_prefixpath + '/static/js/pages/dashboard.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/DashboardController.js',
                    ]
                });
            }]
        }
    })

    // Dashboard
    .state('dashboard', {
        url: "/dashboard.html",
        templateUrl: "dashboard.html",
        data: {
            pageTitle: 'Admin Dashboard Template'
        },
        controller: "DashboardController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load({
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before a LINK element with this ID. Dynamic CSS files must be loaded between core and theme css files
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/morris/morris.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/morris/morris.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/morris/raphael-min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/jquery.sparkline.min.js',

                        ggcmsCfg.cfg_prefixpath + '/static/js/pages/dashboard.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/DashboardController.js',
                    ]
                });
            }]
        }
    })

    // sitesettingList
    .state('sitesetting', {
        url: "/sitesetting.html?pagenum",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'sitesetting.html?pagenum=' + $routeParams.pagenum + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '站点设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // SystemConfigsAdd
    .state('sitesettingadd', {
        url: "/sitesettingadd.html?id",
        templateUrl: function($routeParams) {
            return 'sitesettingadd.html?id=' + $routeParams.id + '&rnd=' + Date.now();
        },
        data: {
            pageTitle: '站点设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // SystemConfigs
    .state('systemconfigs', {
        url: "/systemconfigs.html",
        templateUrl: function($routeParams) {
            return 'systemconfigs.html?rnd=' + Date.now();
        },
        data: {
            pageTitle: '系统设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/configs.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/lang/summernote-zh-CN.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/angular-summernote.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/cfgctrl.js',
                    ]
                }]);
            }]
        }
    })

    // siteconfigs
    .state('siteconfigs', {
        url: "/siteconfigs.html",
        templateUrl: function($routeParams) {
            return 'siteconfigs.html?rnd=' + Date.now();
        },
        data: {
            pageTitle: '网站设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/configs.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/lang/summernote-zh-CN.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/angular-summernote.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/cfgctrl.js',
                    ]
                }]);
            }]
        }
    })

    // category
    .state('category', {
        url: "/category.html",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'category.html?rnd=' + rnd;
        },
        data: {
            pageTitle: '栏目设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/jquery-nestable/jquery.nestable.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/jquery-nestable/jquery.nestable.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/angular-nestable.js',
                    ]
                }]);
            }]
        }
    })

    // categoryadd
    .state('categoryadd', {
        url: "/categoryadd.html?id",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'categoryadd.html?id=' + $routeParams.id + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '栏目设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/lang/summernote-zh-CN.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/angular-summernote.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/angular.treeview-master/css/angular.treeview.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/angular.treeview-master/angular.treeview.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // articlelist
    .state('articlelist', {
        url: "/articlelist.html?type&pagenum&qs",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            var qs = "";
            if ($routeParams.qs) {
                qs = '&qs=' + $routeParams.qs;
            }
            if (!$routeParams.pagenum) {
                $routeParams.pagenum = 1;
            }
            return 'articlelist.html?type=1&pagenum=' + $routeParams.pagenum + qs + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '已发布文章管理'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // rearticlelist
    .state('rearticlelist', {
        url: "/rearticlelist.html?type&pagenum&qs",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            var qs = "";
            if ($routeParams.qs) {
                qs = '&qs=' + $routeParams.qs;
            }
            if (!$routeParams.pagenum) {
                $routeParams.pagenum = 1;
            }
            return 'articlelist.html?type=-1&pagenum=' + $routeParams.pagenum + qs + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '待审核文章管理'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // onearticlelist
    .state('onearticlelist', {
        url: "/onearticlelist.html?type&pagenum&qs",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            var qs = "";
            if ($routeParams.qs) {
                qs = '&qs=' + $routeParams.qs;
            }
            if (!$routeParams.pagenum) {
                $routeParams.pagenum = 1;
            }
            return 'articlelist.html?type=0&pagenum=' + $routeParams.pagenum + qs + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '单页文章管理'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // articleadd
    .state('articleadd', {
        url: "/articleadd.html?type&id",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'articleadd.html?type=' + $routeParams.type + '&id=' + $routeParams.id + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '文章编辑'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load(
                    [{
                        name: 'ui.select',
                        insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                        files: [
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ui-select/select.min.css',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ui-select/select.min.js'
                        ]
                    }, {
                        name: 'GgcmsApp',
                        files: [
                            ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/lang/summernote-zh-CN.js',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/angular-summernote.min.js',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/angular.treeview-master/css/angular.treeview.css',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/angular.treeview-master/angular.treeview.min.js',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.css',
                            ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.js',
                            ggcmsCfg.cfg_prefixpath + '/static/js/configs.js',
                            ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                        ]
                    }]);
            }]
        }
    })

    // systemdict
    .state('systemdict', {
        url: "/systemdict.html?pagenum&qs",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            var qs = "";
            if ($routeParams.qs) {
                qs = '&qs=' + $routeParams.qs;
            }
            if (!$routeParams.pagenum) {
                $routeParams.pagenum = 1;
            }
            return 'systemdict.html?pagenum=' + $routeParams.pagenum + qs + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '系统字典'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // systemdict
    .state('systemdictadd', {
        url: "/systemdictadd.html?id",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'systemdictadd.html?id=' + $routeParams.id + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '系统字典'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // admingroup
    .state('admingroup', {
        url: "/admingroup.html?pagenum&qs",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            var qs = "";
            if ($routeParams.qs) {
                qs = '&qs=' + $routeParams.qs;
            }
            if (!$routeParams.pagenum) {
                $routeParams.pagenum = 1;
            }
            return 'admingroup.html?pagenum=' + $routeParams.pagenum + qs + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '管理员分组管理'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // admingroupadd
    .state('admingroupadd', {
        url: "/admingroupadd.html?id",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'admingroupadd.html?id=' + $routeParams.id + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '管理员分组设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // adminlist
    .state('adminlist', {
        url: "/adminlist.html?pagenum",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            if (!$routeParams.pagenum) {
                $routeParams.pagenum = 1;
            }
            return 'adminlist.html?pagenum=' + $routeParams.pagenum + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '管理员管理'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // adminadd
    .state('adminadd', {
        url: "/adminadd.html?id",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'adminadd.html?id=' + $routeParams.id + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '管理员设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // topiclist
    .state('topiclist', {
        url: "/topiclist.html?pagenum",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            if (!$routeParams.pagenum) {
                $routeParams.pagenum = 1;
            }
            return 'topiclist.html?pagenum=' + $routeParams.pagenum + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '专题管理'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // topicadd
    .state('topicadd', {
        url: "/topicadd.html?id",
        templateUrl: function($routeParams) {
            var rnd = Date.now();
            return 'topicadd.html?id=' + $routeParams.id + '&rnd=' + rnd;
        },
        data: {
            pageTitle: '专题设置'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/lang/summernote-zh-CN.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/bootstrap-summernote/angular-summernote.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/configs.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // modulesdesign
    .state('modulesdesign', {
        url: "/modulesdesign.html",
        templateUrl: function($routeParams) {
            return 'modulesdesign.html?rnd=' + Date.now();
        },
        data: {
            pageTitle: '模型设计'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // modulesdesignadd
    .state('modulesdesignadd', {
        url: "/modulesdesignadd.html?id",
        templateUrl: function($routeParams) {
            return 'modulesdesignadd.html?id=' + $routeParams.id + '&rnd=' + Date.now();
        },
        data: {
            pageTitle: '模型设计'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/configs.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // styledesign
    .state('styledesign', {
        url: "/styledesign.html",
        templateUrl: function($routeParams) {
            return 'styledesign.html?rnd=' + Date.now();
        },
        data: {
            pageTitle: '风格模板设计'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/datalist.js'
                    ]
                }]);
            }]
        }
    })

    // styledesignadd
    .state('styledesignadd', {
        url: "/styledesignadd.html?id",
        templateUrl: function($routeParams) {
            return 'styledesignadd.html?id=' + $routeParams.id + '&rnd=' + Date.now();
        },
        data: {
            pageTitle: '风格模板设计'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/configs.js',
                        ggcmsCfg.cfg_prefixpath + '/static/js/dataedit.js'
                    ]
                }]);
            }]
        }
    })

    // styletemplate
    .state('styletemplate', {
        url: "/styletemplate.html?id",
        templateUrl: function($routeParams) {
            return 'styletemplate.html?id=' + $routeParams.id + '&rnd=' + Date.now();
        },
        data: {
            pageTitle: '模板查看'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'codemirrorLoad',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/lib/codemirror.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/lib/ui-codemirror.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/lib/codemirror.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/theme/neat.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/addon/fold/foldgutter.css',
                    ]
                }, {
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                    ]
                }]);
            }]
        }
    })

    // stylestaticfile
    .state('stylestaticfile', {
        url: "/stylestaticfile.html?id",
        templateUrl: function($routeParams) {
            return 'stylestaticfile.html?id=' + $routeParams.id + '&rnd=' + Date.now();
        },
        data: {
            pageTitle: '静态文件管理'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'codemirrorLoad',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/lib/codemirror.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/lib/ui-codemirror.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/lib/codemirror.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/theme/neat.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/codemirror/addon/fold/foldgutter.css',
                    ]
                }, {
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/angularjs/plugins/ng-file-upload-12.0.4/ng-file-upload-all.min.js',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.css',
                        ggcmsCfg.cfg_prefixpath + '/static/plugins/img-preview/imgPreview.js',
                    ]
                }]);
            }]
        }
    })

    // AngularJS plugins
    .state('fileupload', {
        url: "/file_upload.html",
        templateUrl: "views/file_upload.html",
        data: {
            pageTitle: 'AngularJS File Upload'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'angularFileUpload',
                    files: [
                        '../assets/global/plugins/angularjs/plugins/angular-file-upload/angular-file-upload.min.js',
                    ]
                }, {
                    name: 'GgcmsApp',
                    files: [
                        'js/controllers/GeneralPageController.js'
                    ]
                }]);
            }]
        }
    })

    // UI Select
    .state('uiselect', {
        url: "/ui_select.html",
        templateUrl: "views/ui_select.html",
        data: {
            pageTitle: 'AngularJS Ui Select'
        },
        controller: "UISelectController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'ui.select',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/angularjs/plugins/ui-select/select.min.css',
                        '../assets/global/plugins/angularjs/plugins/ui-select/select.min.js'
                    ]
                }, {
                    name: 'GgcmsApp',
                    files: [
                        'js/controllers/UISelectController.js'
                    ]
                }]);
            }]
        }
    })

    // UI Bootstrap
    .state('uibootstrap', {
        url: "/ui_bootstrap.html",
        templateUrl: "ui_bootstrap.html",
        data: {
            pageTitle: 'AngularJS UI Bootstrap'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    files: [
                        ggcmsCfg.cfg_prefixpath + '/static/js/controllers/GeneralPageController.js'
                    ]
                }]);
            }]
        }
    })

    // Tree View
    .state('tree', {
        url: "/tree",
        templateUrl: "views/tree.html",
        data: {
            pageTitle: 'jQuery Tree View'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/jstree/dist/themes/default/style.min.css',

                        '../assets/global/plugins/jstree/dist/jstree.min.js',
                        '../assets/pages/scripts/ui-tree.min.js',
                        'js/controllers/GeneralPageController.js'
                    ]
                }]);
            }]
        }
    })

    // Form Tools
    .state('formtools', {
        url: "/form-tools",
        templateUrl: "views/form_tools.html",
        data: {
            pageTitle: 'Form Tools'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/bootstrap-fileinput/bootstrap-fileinput.css',
                        '../assets/global/plugins/bootstrap-switch/css/bootstrap-switch.min.css',
                        '../assets/global/plugins/bootstrap-markdown/css/bootstrap-markdown.min.css',
                        '../assets/global/plugins/typeahead/typeahead.css',

                        '../assets/global/plugins/fuelux/js/spinner.min.js',
                        '../assets/global/plugins/bootstrap-fileinput/bootstrap-fileinput.js',
                        '../assets/global/plugins/jquery-inputmask/jquery.inputmask.bundle.min.js',
                        '../assets/global/plugins/jquery.input-ip-address-control-1.0.min.js',
                        '../assets/global/plugins/bootstrap-pwstrength/pwstrength-bootstrap.min.js',
                        '../assets/global/plugins/bootstrap-switch/js/bootstrap-switch.min.js',
                        '../assets/global/plugins/bootstrap-maxlength/bootstrap-maxlength.min.js',
                        '../assets/global/plugins/bootstrap-touchspin/bootstrap.touchspin.js',
                        '../assets/global/plugins/typeahead/handlebars.min.js',
                        '../assets/global/plugins/typeahead/typeahead.bundle.min.js',
                        '../assets/pages/scripts/components-form-tools-2.min.js',

                        'js/controllers/GeneralPageController.js'
                    ]
                }]);
            }]
        }
    })

    // Date & Time Pickers
    .state('pickers', {
        url: "/pickers",
        templateUrl: "views/pickers.html",
        data: {
            pageTitle: 'Date & Time Pickers'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/clockface/css/clockface.css',
                        '../assets/global/plugins/bootstrap-datepicker/css/bootstrap-datepicker3.min.css',
                        '../assets/global/plugins/bootstrap-timepicker/css/bootstrap-timepicker.min.css',
                        '../assets/global/plugins/bootstrap-colorpicker/css/colorpicker.css',
                        '../assets/global/plugins/bootstrap-daterangepicker/daterangepicker-bs3.css',
                        '../assets/global/plugins/bootstrap-datetimepicker/css/bootstrap-datetimepicker.min.css',

                        '../assets/global/plugins/bootstrap-datepicker/js/bootstrap-datepicker.min.js',
                        '../assets/global/plugins/bootstrap-timepicker/js/bootstrap-timepicker.min.js',
                        '../assets/global/plugins/clockface/js/clockface.js',
                        '../assets/global/plugins/moment.min.js',
                        '../assets/global/plugins/bootstrap-daterangepicker/daterangepicker.js',
                        '../assets/global/plugins/bootstrap-colorpicker/js/bootstrap-colorpicker.js',
                        '../assets/global/plugins/bootstrap-datetimepicker/js/bootstrap-datetimepicker.min.js',

                        '../assets/pages/scripts/components-date-time-pickers.min.js',

                        'js/controllers/GeneralPageController.js'
                    ]
                }]);
            }]
        }
    })

    // Custom Dropdowns
    .state('dropdowns', {
        url: "/dropdowns",
        templateUrl: "views/dropdowns.html",
        data: {
            pageTitle: 'Custom Dropdowns'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/bootstrap-select/css/bootstrap-select.min.css',
                        '../assets/global/plugins/select2/css/select2.min.css',
                        '../assets/global/plugins/select2/css/select2-bootstrap.min.css',

                        '../assets/global/plugins/bootstrap-select/js/bootstrap-select.min.js',
                        '../assets/global/plugins/select2/js/select2.full.min.js',

                        '../assets/pages/scripts/components-bootstrap-select.min.js',
                        '../assets/pages/scripts/components-select2.min.js',

                        'js/controllers/GeneralPageController.js'
                    ]
                }]);
            }]
        }
    })

    // Advanced Datatables
    .state('datatablesAdvanced', {
        url: "/datatables/managed.html",
        templateUrl: "views/datatables/managed.html",
        data: {
            pageTitle: 'Advanced Datatables'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load({
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/datatables/datatables.min.css',
                        '../assets/global/plugins/datatables/plugins/bootstrap/datatables.bootstrap.css',

                        '../assets/global/plugins/datatables/datatables.all.min.js',

                        '../assets/pages/scripts/table-datatables-managed.min.js',

                        'js/controllers/GeneralPageController.js'
                    ]
                });
            }]
        }
    })

    // Ajax Datetables
    .state('datatablesAjax', {
        url: "/datatables/ajax.html",
        templateUrl: "views/datatables/ajax.html",
        data: {
            pageTitle: 'Ajax Datatables'
        },
        controller: "GeneralPageController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load({
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/datatables/datatables.min.css',
                        '../assets/global/plugins/datatables/plugins/bootstrap/datatables.bootstrap.css',
                        '../assets/global/plugins/bootstrap-datepicker/css/bootstrap-datepicker3.min.css',

                        '../assets/global/plugins/datatables/datatables.all.min.js',
                        '../assets/global/plugins/bootstrap-datepicker/js/bootstrap-datepicker.min.js',
                        '../assets/global/scripts/datatable.min.js',

                        'js/scripts/table-ajax.js',
                        'js/controllers/GeneralPageController.js'
                    ]
                });
            }]
        }
    })

    // User Profile
    .state("profile", {
        url: "/profile",
        templateUrl: "views/profile/main.html",
        data: {
            pageTitle: 'User Profile'
        },
        controller: "UserProfileController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load({
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/bootstrap-fileinput/bootstrap-fileinput.css',
                        '../assets/pages/css/profile.css',

                        '../assets/global/plugins/jquery.sparkline.min.js',
                        '../assets/global/plugins/bootstrap-fileinput/bootstrap-fileinput.js',

                        '../assets/pages/scripts/profile.min.js',

                        'js/controllers/UserProfileController.js'
                    ]
                });
            }]
        }
    })

    // User Profile Dashboard
    .state("profile.dashboard", {
        url: "/dashboard",
        templateUrl: "views/profile/dashboard.html",
        data: {
            pageTitle: 'User Profile'
        }
    })

    // User Profile Account
    .state("profile.account", {
        url: "/account",
        templateUrl: "views/profile/account.html",
        data: {
            pageTitle: 'User Account'
        }
    })

    // User Profile Help
    .state("profile.help", {
        url: "/help",
        templateUrl: "views/profile/help.html",
        data: {
            pageTitle: 'User Help'
        }
    })

    // Todo
    .state('todo', {
        url: "/todo",
        templateUrl: "views/todo.html",
        data: {
            pageTitle: 'Todo'
        },
        controller: "TodoController",
        resolve: {
            deps: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load({
                    name: 'GgcmsApp',
                    insertBefore: '#ng_load_plugins_before', // load the above css files before '#ng_load_plugins_before'
                    files: [
                        '../assets/global/plugins/bootstrap-datepicker/css/bootstrap-datepicker3.min.css',
                        '../assets/apps/css/todo-2.css',
                        '../assets/global/plugins/select2/css/select2.min.css',
                        '../assets/global/plugins/select2/css/select2-bootstrap.min.css',

                        '../assets/global/plugins/select2/js/select2.full.min.js',

                        '../assets/global/plugins/bootstrap-datepicker/js/bootstrap-datepicker.min.js',

                        '../assets/apps/scripts/todo-2.min.js',

                        'js/controllers/TodoController.js'
                    ]
                });
            }]
        }
    })

}]);