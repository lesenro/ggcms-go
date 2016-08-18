/***
Ggcms AngularJS App Main Script
***/

/* Ggcms App */
var GgcmsApp = angular.module("GgcmsApp", [
    "ui.router",
    "ui.bootstrap",
    "oc.lazyLoad",
    "ngSanitize"
]);
GgcmsApp.service("AJAX", function($http, $timeout) {
    var cfg = {
        method: "POST",
        url: "",
        func: null,
        data: null,
        showmsg: true,
        params: null,
        errfunc: null,
    };
    return {
        load: function(o) {
            var tmp = {};
            var c = $.extend(tmp, cfg, o);
            if (c.url == "") {
                return false;
            }
            $http({
                    method: c.method,
                    url: ggcmsCfg.cfg_prefixpath + c.url,
                    data: c.data, // pass in data as strings
                    params: c.params,
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded'
                    } // set the headers so angular passing info as form data (not request payload)
                })
                .success(function(retData, retcode, retfunc, retobj) {
                    //成功
                    if (retData.Code == 0) {
                        if (c.func) {
                            c.func(retData, retcode, retfunc, retobj);
                        }
                        $timeout(function() {
                            App.initUniform();
                        });
                    } else {
                        if (c.showmsg) {
                            var msg = {};
                            msg.title = "失败";
                            msg.icon = "error";
                            msg.msg = retData.Code + " : " + retData.Msg;
                            App.messageBox(msg);
                        } else {
                            console.log(retData.Code + " : " + retData.Msg);
                        }
                        if (c.errfunc) {
                            c.errfunc(retData, retcode, retfunc, retobj);
                        }
                    }
                    App.unblockUI();
                }).error(function(err) {
                    console.error(err);
                    if (c.errfunc) {
                        c.errfunc(err);
                    }
                });
        }
    };
});
GgcmsApp.service("FileUpload", function($http, $timeout, Upload) {
    var cfg = {
        file: null,
        func: null,
        inputid: "",
        siteid: 0,
        uptype: 0
    };
    return {
        send: function(o) {
            var tmp = {};
            var c = $.extend(tmp, cfg, o);
            if (c.file == null) {
                return false;
            }
            var func = c.func;
            c.func = null;
            Upload.upload({
                url: ggcmsCfg.cfg_prefixpath + "/api/ggcms_upload",
                data: c
            }).then(function(resp) {
                if (resp.data.Code == 0 && func) {
                    func(resp);
                } else {
                    console.log(resp);
                }

            }, function(resp) {
                console.log(resp.status);
            }, function(evt) {
                var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
                //console.log(progressPercentage);
            });
        }
    };
});
/* Configure ocLazyLoader(refer: https://github.com/ocombe/ocLazyLoad) */
GgcmsApp.config(['$ocLazyLoadProvider', function($ocLazyLoadProvider) {
    $ocLazyLoadProvider.config({
        // global configs go here
    });
}]);

/********************************************
 BEGIN: BREAKING CHANGE in AngularJS v1.3.x:
*********************************************/
/**
`$controller` will no longer look for controllers on `window`.
The old behavior of looking on `window` for controllers was originally intended
for use in examples, demos, and toy apps. We found that allowing global controller
functions encouraged poor practices, so we resolved to disable this behavior by
default.

To migrate, register your controllers with modules rather than exposing them
as globals:

Before:

```javascript
function MyController() {
  // ...
}
```

After:

```javascript
angular.module('myApp', []).controller('MyController', [function() {
  // ...
}]);

Although it's not recommended, you can re-enable the old behavior like this:

```javascript
angular.module('myModule').config(['$controllerProvider', function($controllerProvider) {
  // this option might be handy for migrating old apps, but please don't use it
  // in new ones!
  $controllerProvider.allowGlobals();
}]);
**/

//AngularJS v1.3.x workaround for old style controller declarition in HTML
GgcmsApp.config(['$controllerProvider', function($controllerProvider) {
    // this option might be handy for migrating old apps, but please don't use it
    // in new ones!
    $controllerProvider.allowGlobals();
}]);

/********************************************
 END: BREAKING CHANGE in AngularJS v1.3.x:
*********************************************/

/* Setup global settings */
GgcmsApp.factory('settings', ['$rootScope', 'AJAX', function($rootScope, AJAX) {
    // supported languages
    var settings = {
        layout: {
            pageSidebarClosed: false, // sidebar menu state
            pageContentWhite: true, // set page content layout
            pageBodySolid: false, // solid body color state
            pageAutoScrollOnLoad: 1000 // auto scroll to top on page load
        },
        assetsPath: '/static',
        globalPath: '/static',
        layoutPath: '/static',
        currentSite: 0,
        powers: { "allallow": 0 },
    };
    settings.adminAuth = function(pow) {
        if (settings.powers["allallow"] == 1) {
            return true;
        }
        if (settings.powers[pow]) {
            return true;
        } else {
            return false;
        }
    }
    $rootScope.settings = settings;
    AJAX.load({
        method: 'get',
        url: '/api/ggcms_tools/getpowers',
        showmsg: false,
        params: {},
        func: function(retData) {
            settings.powers = retData.Data;
        }
    });
    return settings;
}]);

/* Setup App Main Controller */
GgcmsApp.controller('AppController', ['$scope', '$rootScope', '$http', function($scope, $rootScope, $http) {
    $scope.$on('$viewContentLoaded', function() {
        App.initComponents(); // init core components
        //Layout.init(); //  Init entire layout(header, footer, sidebar, etc) on page load if the partials included in server side instead of loading with ng-include directive 
    });

}]);

/***
Layout Partials.
By default the partials are loaded through AngularJS ng-include directive. In case they loaded in server side(e.g: PHP include function) then below partial 
initialization can be disabled and Layout.init() should be called on page load complete as explained above.
***/

/* Setup Layout Part - Header */
GgcmsApp.controller('HeaderController', ['$scope', 'AJAX', function($scope, AJAX) {
    $scope.$on('$includeContentLoaded', function() {
        Layout.initHeader(); // init header
    });
    $scope.clearCache = function() {
        App.messageBox({
            msg: "是否要清理缓存?",
            icon: "question",
            btn: "yes|no",
            func: function(e) {
                AJAX.load({
                    method: 'get',
                    url: '/api/ggcms_tools/clearcacth',
                    params: {},
                    func: function(retData) {
                        //成功
                        var msg = {};
                        if (retData.Code == 0) {
                            msg.msg = "缓存清理成功!!!";
                        } else {
                            msg.title = "失败";
                            msg.icon = "error";
                        }
                        App.messageBox(msg);
                    }
                });
            }
        });
    };
    $scope.siteSelect = function(site, reload) {
        $scope.settings.currentSite = site.Id;
        $scope.cursite = site.Site_name;
        Cookies.set("currentSite", site.Id, {
            expires: 7,
            path: '/'
        });
        if (reload) {
            window.location.reload();
        }
    };
    AJAX.load({
        method: 'get',
        url: '/api/ggcms_sites',
        showmsg: false,
        params: {
            limit: 0,
            offset: 0
        },
        func: function(retData) {
            //成功
            $scope.sitelist = retData.Data;
            $scope.sitetotal = retData.Data.length;
            var cid = Cookies.get("currentSite");
            for (i in $scope.sitelist) {
                if ($scope.sitelist[i].Id == cid) {
                    $scope.siteSelect($scope.sitelist[i]);
                    break;
                }
            }
            if ($scope.settings.currentSite == 0) {
                for (i in $scope.sitelist) {
                    if ($scope.sitelist[i].Site_main == 1) {
                        $scope.siteSelect($scope.sitelist[i]);
                        break;
                    }
                }
            }
        }
    });
    // $http({
    //         method: 'get',
    //         url: '/api/ggcms_sites',
    //         params: {
    //             limit: 0,
    //             offset: 0
    //         }, // pass in data as strings
    //         headers: {
    //             'Content-Type': 'application/x-www-form-urlencoded'
    //         } // set the headers so angular passing info as form data (not request payload)
    //     })
    //     .success(function(retData) {
    //         //成功
    //         $scope.sitelist = retData;
    //         $scope.sitetotal = retData.length;
    //         var cid = Cookies.get("currentSite");
    //         for (i in $scope.sitelist) {
    //             if ($scope.sitelist[i].Id == cid) {
    //                 $scope.siteSelect($scope.sitelist[i]);
    //                 break;
    //             }
    //         }
    //         if ($scope.settings.currentSite == 0) {
    //             for (i in $scope.sitelist) {
    //                 if ($scope.sitelist[i].Site_main == 1) {
    //                     $scope.siteSelect($scope.sitelist[i]);
    //                     break;
    //                 }
    //             }
    //         }
    //     });

}]);

/* Setup Layout Part - Sidebar */
GgcmsApp.controller('SidebarController', ['$scope', function($scope) {
    $scope.$on('$includeContentLoaded', function() {
        Layout.initSidebar(); // init sidebar
        console.log($scope.settings);
    });
}]);

/* Setup Layout Part - Sidebar */
GgcmsApp.controller('PageHeadController', ['$scope', function($scope) {
    $scope.$on('$includeContentLoaded', function() {
        Demo.init(); // init theme panel
    });
}]);

/* Setup Layout Part - Quick Sidebar */
GgcmsApp.controller('QuickSidebarController', ['$scope', function($scope) {
    $scope.$on('$includeContentLoaded', function() {
        setTimeout(function() {
            QuickSidebar.init(); // init quick sidebar        
        }, 2000)
    });
}]);

/* Setup Layout Part - Theme Panel */
GgcmsApp.controller('ThemePanelController', ['$scope', function($scope) {
    $scope.$on('$includeContentLoaded', function() {
        Demo.init(); // init theme panel
    });
}]);

/* Setup Layout Part - Footer */
GgcmsApp.controller('FooterController', ['$scope', function($scope) {
    $scope.$on('$includeContentLoaded', function() {
        Layout.initFooter(); // init footer
    });
}]);


/* Init global settings and run the app */
GgcmsApp.run(["$rootScope", "settings", "$state", function($rootScope, settings, $state) {
    $rootScope.$state = $state; // state to be accessed from view
    $rootScope.$settings = settings; // state to be accessed from view
}]);