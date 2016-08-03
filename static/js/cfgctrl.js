function ConfigsCtrl($scope, FileUpload, AJAX, $timeout, $ocLazyLoad) {
    $scope.datamap = configs;
    $scope.data = [];
    $scope.dict = {};
    $scope.style = {};
    //文件初始化
    var fileInputs = [];

    //初始化基础数据
    for (i in $scope.datamap) {
        var option = $.trim($scope.datamap[i].Others) == "" ? "{}" : $scope.datamap[i].Others;
        option = eval("(" + option + ")");
        var tmp = {};
        $scope.datamap[i].opts = $.extend(tmp, ggcmsCfg.defaultOptions, option);
        //数字转换
        if (tmp.type == "number") {
            if ($scope.datamap[i].Value.indexOf(".") == -1) {
                $scope.datamap[i].Value = parseInt($scope.datamap[i].Value) || "0";
            } else {
                $scope.datamap[i].Value = parseFloat($scope.datamap[i].Value) || "0";
            }
        }
        //多选转为数组
        if ($scope.datamap[i].opts.type == "select" && $scope.datamap[i].opts.multiple) {
            $scope.datamap[i].Value = $scope.datamap[i].Value.split(",");
        }
        //日期转换
        if (tmp.type == "date" || tmp.type == "datetime") {
            $scope.datamap[i].Value = Date.parse($scope.datamap[i].Value);
        }
        //获取字典数据
        if ((tmp.type == "select" || tmp.type == "checkbox" || tmp.type == "radio") && tmp.datasource == "sysenum" && "" != tmp.egroup) {
            AJAX.load({
                url: "/api/ggcms_sys_enum",
                method: 'GET',
                showmsg: false,
                params: {
                    query: "egroup:" + tmp.egroup
                },
                func: function(retData, a, b, c) {
                    var reg = /egroup:(\w+)[^;|,]*/gi;
                    var egroup = reg.exec(c.params.query)[1];
                    $scope.dict[egroup] = retData.Data;
                }
            });
        }
        //获取风格数据
        if ((tmp.type == "select" || tmp.type == "radio") && tmp.datasource == "style") {
            var defdir = $scope.datamap[i].Value;
            AJAX.load({
                url: "/api/ggcms_styles/all",
                method: 'GET',
                showmsg: false,
                func: function(retData) {
                    $scope.style = retData.Data;
                    $scope.styleChange(defdir);
                }
            });
        }
        //文件
        if (tmp.type == "file") {
            $scope.datamap[i].FileValue = $scope.datamap[i].Value;
        }
        //添加为数组，方便ng-repeat加载
        $scope.data.push($scope.datamap[i]);
    }
    //保存修改
    $scope.update = function(data) {
        //复制数据
        var ndata = angular.copy($scope.datamap);
        //转换数字和多选内容
        for (i in ndata) {
            if (ndata[i].opts.type == "number") {
                ndata[i].Value = ndata[i].Value.toString();
            }
            if (ndata[i].opts.type == "select" && ndata[i].opts.multiple) {
                ndata[i].Value = ndata[i].Value.join(",");
            }
            if (ndata[i].opts.type == "switch") {
                ndata[i].Value = $("#" + ndata[i].Varname).bootstrapSwitch("state") ? "on" : ""
            }
            //日期转换
            if (ndata[i].opts.type == "date" || ndata[i].opts.type == "datetime") {
                ndata[i].Value = ndata[i].Value.toString();
            }
            //文件
            if (tmp.type == "file") {
                ndata[i].Value = ndata[i].FileValue;
            }
            //清理数据，减少提交数据大小
            delete ndata[i].Others;
            delete ndata[i].opts;
        }
        //有上传的文件
        if (fileInputs.length > 0) {
            ndata.UpinfoList = {
                Value: angular.toJson(fileInputs)
            };
        }
        //提交表单
        AJAX.load({
            url: "/api/system_configs/update",
            data: ndata,
            func: function(retData) {
                App.messageBox({
                    msg: "恭喜，您的修改已保存成功",
                    func: function() {
                        window.location.reload();
                    }
                });
            }
        });
    };
    //初始化checkbox判断是否选中
    $scope.cbxselect = function(k, v) {
        return ("," + k + ",").indexOf("," + v + ",") > -1;
    };
    //checkbox修改
    $scope.cbxchange = function(k, v, $event) {
        //更新样式
        App.updateUniform();
        var sel = [];
        $("input[name='" + k + "']:checked").each(function(e) {
            sel.push($(this).val());
        });
        $scope.datamap[k].Value = sel.join(",");
    };
    //日期
    $scope.datePicker = ggcmsCfg.datePicker;
    $scope.timePicker = ggcmsCfg.timePicker;
    //时间
    //初始化富文本编辑框
    $scope.richeditor = {
        lang: "zh-CN",
    };
    $scope.simpleditor = {
        lang: "zh-CN",
        toolbar: [
            ['color', ['color']],
            ['font', ['bold', 'underline', 'clear']],
            ['para', ['ul', 'paragraph']],
            ['table', ['table']],
            ['insert', ['link']]
        ],
    };
    //页面加载完成后，处理内容
    $timeout(function() {
        //更新表单控件样式
        App.initAjax();
    });
    //风格修改
    $scope.styleChange = function(dir) {
        $scope.style.template = $scope.style["t_" + dir];
    };
    //模板过滤
    $scope.templateFilter = function(name, val) {
        var reg = eval("/" + val + "_/i");
        return reg.test(name);
    };
    //文件上传相关
    // upload on file select or drop
    $scope.upload = function(file, inputId, assignment, func) {
        if (!file) {
            return false;
        }
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
                fileInputs.push(upinfo);
                if (assignment) {
                    $scope.datamap[inputId].FileValue = file.name;
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
    //清除文件
    $scope.fileclear = function(inputId) {
        $scope.datamap[inputId].FileValue = "";
        var idxs = [];
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
    };
    $scope.preview = function(inputId) {
        var url = $scope.datamap[inputId].FileValue;
        //有新上传图片
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