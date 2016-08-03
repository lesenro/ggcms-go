/* Setup general page controller */
angular.module('GgcmsApp').controller('DataListCtrl', ['$scope', 'AJAX', function($scope, AJAX) {
    $scope.delData = [];
    $scope.delete = function(data, $event) {
        var len = data.length;
        if (len == 0) {
            return false;
        }
        App.messageBox({
            msg: "是否要删除所选中的(" + len + ")个数据?",
            btn: "yes|no",
            func: function(e) {
                IdsDelete(data);
            }
        });
    }
    $scope.check = function(val, $event) {
        var chk = $event.currentTarget.checked;
        if (chk == true) {
            $scope.delData.push(val);
        } else {
            for (i in $scope.delData) {
                if ($scope.delData[i] == val) {
                    $scope.delData.splice(i, 1);
                }
            }
        }
    };

    $scope.checkAll = function($event) {
        var chk = $event.currentTarget.checked;
        App.updateUniform();
        $scope.delData = [];
        if (chk) {
            $(".data-list .item-check").each(function(e) {
                var val = parseInt($(this).val());
                $scope.delData.push(val);
            });
        }
    }
    var IdsDelete = function(ids) {
        AJAX.load({
            url: deleteUrl,
            data: ids,
            func: function(retData) {
                var msg = {};
                msg.msg = "成功删除 " + retData.Msg + " 条记录";
                msg.func = function() {
                    window.location.reload();
                };
                App.messageBox(msg);
            }
        });
    };
}]);