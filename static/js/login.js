/* GgcmsApp App */
var GgcmsApp = angular.module("GgcmsApp", [
    "ui.bootstrap",
]);
/* Setup Layout Part - Header */
GgcmsApp.controller('LoginController', ['$scope', 'AJAX', function($scope, AJAX) {
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
        AJAX.load({
            url: "/api/ggcms_admin/login",
            data: data,
            func: function(retData) {
                location.href = "./"
            },
            errfunc: function() {
                $scope.ChangCode();
            }
        });
    };
    $scope.ChangCode();
}]);