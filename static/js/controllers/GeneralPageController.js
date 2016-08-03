/* Setup general page controller */
angular.module('GgcmsApp').service("AJAX", function($http, $timeout) {
    var cfg = {
        method: "POST",
        url: "",
        func: null,
        data: null,
        showmsg: true,
        params: null,
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
                    }
                    App.unblockUI();
                }).error(function(err) {
                    console.error(err);
                });
        }
    };
}).service("FileUpload", function($http, $timeout, Upload) {
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
}).controller('GeneralPageController', ['$rootScope', '$scope', 'settings', '$stateParams', function($rootScope, $scope, settings, $stateParams) {
    $scope.$on('$viewContentLoaded', function() {
        // initialize core components
        App.initAjax();
        // set default layout mode
        $rootScope.settings.layout.pageContentWhite = true;
        $rootScope.settings.layout.pageBodySolid = false;
        $rootScope.settings.layout.pageSidebarClosed = false;
    });

}]);