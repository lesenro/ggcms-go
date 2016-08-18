/* GgcmsApp App */
var GgcmsApp = angular.module("GgcmsApp", [
    "ui.bootstrap",
]);
/* Setup Layout Part - Header */
GgcmsApp.controller('LoginController', ['$scope', '$http', function($scope, $http) {
    $scope.data = {
        username: "",
        password: "",
        checkcode: "",
        remember: false
    };
    $scope.ChangCode = function() {
        $scope.chkcodeurl = ggcmsCfg.cfg_prefixpath + "/getcode?rnd=" + App.getUniqueID();
    };
    $scope.login = function() {
        var data = angular.copy($scope.data);
        data.password = $.md5($.md5(data.password) + data.checkcode.toLowerCase());
        $http({
            method: "post",
            url: ggcmsCfg.cfg_prefixpath + "/api/ggcms_admin/login",
            data: data, // pass in data as strings
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            } // set the headers so angular passing info as form data (not request payload)
        }).success(function(retData) {
            //成功
            if (retData.Code == 0) {
                location.href = "./"
            } else {
                loginError(retData);
            }

        }).error(function() {
                loginError({Code:"",Msg:"登录失败"});
        });
    };
    var loginError = function(retData) {
        var msg = {};
        msg.title = "失败";
        msg.icon = "error";
        msg.msg = retData.Code + " : " + retData.Msg;
        App.messageBox(msg);
        $scope.ChangCode();
    }
    $scope.ChangCode();
}]);