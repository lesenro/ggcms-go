function DataEditCtrl($scope, AJAX) {
    $scope.data = infoData.data;
    $scope.update = function(data) {
        var strtype = 'add';
        var strmsg = "添加成功";
        if (data.Id != 0) {
            strtype = "edit";
            strmsg = "修改成功";
        }
        if (infoData.ProcessData) {
            infoData.ProcessData($scope.data);
        }
        AJAX.load({
            url: infoData.postUrl + strtype,
            data: $scope.data,
            func: function(retData) {
                var msg = {};
                msg.msg = infoData.msgTitle + " “" + $scope.data[infoData.msgShowInput] + "” " + strmsg;
                msg.func = function() {
                    window.history.go(-1);
                };
                App.messageBox(msg);
            }
        });
    };
};


function DataEditUploadCtrl($scope, FileUpload, AJAX, $timeout) {
    $scope.data = infoData.data;
    //文件初始化
    var fileInputs = [];
    //更新
    $scope.update = function(data) {
        var strtype = 'add';
        var strmsg = "添加成功";
        if (data.Id != 0) {
            strtype = "edit";
            strmsg = "修改成功";
        }
        var ndata = angular.copy(data);
        ndata.UpinfoList = fileInputs;
        if (infoData.ProcessData) {
            infoData.ProcessData(ndata);
        }
        AJAX.load({
            url: infoData.postUrl + strtype,
            data: ndata,
            func: function(retData) {
                var msg = {};
                msg.msg = infoData.msgTitle + " “" + $scope.data[infoData.msgShowInput] + "” " + strmsg;
                msg.func = function() {
                    window.history.go(-1);
                };
                App.messageBox(msg);
            }
        });
    };
    //上传

    $scope.upload = function(file, inputId, assignment, func) {
        if (!file) {
            return false;
        }
        // console.log(file);

        FileUpload.send({
            file: file,
            inputid: inputId,
            siteid: $scope.settings.currentSite,
            func: function(resp) {
                upinfo = {
                    InputId: inputId,
                    Realname: resp.data.Data,
                    Name: file.name,
                    Assignment: assignment,
                };
                if (assignment) {
                    //主信息中包含
                    if ($scope.data.hasOwnProperty(inputId)) {
                        console.log(inputId, assignment, func);

                        $scope.data[inputId] = file.name;
                    } else {
                        //用于文章编辑中的模型处理
                        console.log($scope.data.modulesData[inputId].FileValue);
                        $scope.data.modulesData[inputId].FileValue = file.name;
                    }
                    var idx = -1;
                    for (var i = 0; i < fileInputs.length; i++) {
                        var finfo = fileInputs[i];
                        if (finfo.InputId == inputId) {
                            idx = i;
                            break;
                        }
                    }
                    if (idx > -1) {
                        fileInputs[idx] = upinfo;
                    } else {
                        fileInputs.push(upinfo);
                    }
                } else {
                    fileInputs.push(upinfo);
                }

                if (func) {
                    func(upinfo);
                }
            }
        });
    };
    //编辑器中文件上传
    $scope.imageUpload = function(files, inputId) {
        for (var i in files) {
            var ft = $.type(files[i]);
            if (ft != "object") {
                continue;
            }
            $scope.upload(files[i], inputId, false, function(upinfo) {
                $('#' + inputId).summernote('insertImage', "/" + upinfo.Realname, upinfo.Name);
            });
        }
    };
    $scope.fileclear = function(inputId) {
        //主信息中包含
        if ($scope.data.hasOwnProperty(inputId)) {
            $scope.data[inputId] = "";
        } else {
            //用于文章编辑中的模型处理
            $scope.data.modulesData[inputId].FileValue = "";
        }
        var idxs = [];
        //有新上传图片
        for (var i = 0; i < fileInputs.length; i++) {
            var finfo = fileInputs[i];
            if (finfo.InputId == inputId) {
                idxs.push(i);
            }
        }
        //删除
        for (var i = idxs.length; i > 0; i--) {
            fileInputs.splice(idxs[i] - 1, 1);
        }
    }
    $scope.preview = function(inputId) {
        var url = $scope.data[inputId];
        if (!$scope.data.hasOwnProperty(inputId)) {
            //用于文章编辑中的模型处理
            url = $scope.data.modulesData[inputId].FileValue;
        }

        for (var i = 0; i < fileInputs.length; i++) {
            var finfo = fileInputs[i];
            if (finfo.InputId == inputId) {
                url = fileInputs[i].Realname;
                break;
            }
        }
        preView.show(url);
    };
};