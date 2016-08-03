ggcmsCfg.defaultOptions = {
    type: "text",
    required: false,
    min: 0,
    max: 0,
    minLength: 0,
    maxLength: 0,
    pattern: "",
    message: "",
    requiredMessage: "",
    minMessage: "",
    maxMessage: "",
    minLengthMessage: "",
    maxLengthMessage: "",
    patternMessage: "",
    helpMessage: "",
    preview: false,
    datasource: "",
    egroup: "",
    multiple: false,
    onColor: "info",
    offColor: "default",
    onText: "ON",
    offText: "OFF",
    minDate: "",
    minDateMessage: "",
    maxDate: "",
    maxDateMessage: "",
    extension: "",
    extensionMessage: "",
    fileSize: 0,
    fileSizenMessage: "",
}
ggcmsCfg.datePicker = {
    options: {
        formatYear: 'yyyy',
        startingDay: 1,
        formatMonth: "MMMM",
        formatMonthTitle: "yyyy",
        formatDayTitle: "MM yyyy"
    },
    format: "yyyy-MM-dd",
    currentText: "今天",
    clearText: "清除",
    closeText: "关闭",
    open: function($event, opts) {
        $event.preventDefault();
        $event.stopPropagation();
        opts.opened = true;
    },

}

ggcmsCfg.keyValues = [
    //日期控件类型
    {
        "key": "n",
        "val": "当前日期",
        "type": "dtype"
    }, {
        "key": "v",
        "val": "变量日期",
        "type": "dtype"
    }, {
        "key": "c",
        "val": "自定义日期",
        "type": "dtype"
    },
    //日期加减单位 
    {
        "key": "m",
        "val": "月",
        "type": "kind"
    }, {
        "key": "y",
        "val": "年",
        "type": "kind"
    }, {
        "key": "d",
        "val": "日",
        "type": "kind"
    },
    //switch 颜色
    {
        "key": "primary",
        "val": "深蓝色",
        "type": "color"
    }, {
        "key": "info",
        "val": "浅蓝色",
        "type": "color"
    }, {
        "key": "success",
        "val": "青绿色",
        "type": "color"
    }, {
        "key": "warning",
        "val": "暗黄色",
        "type": "color"
    }, {
        "key": "danger",
        "val": "红色",
        "type": "color"
    }, {
        "key": "default",
        "val": "灰色",
        "type": "color"
    },
    //字段类型
    {
        "key": "varchar",
        "val": "字符串",
        "type": "columntype"
    }, {
        "key": "int",
        "val": "整数",
        "type": "columntype"
    }, {
        "key": "text",
        "val": "文本",
        "type": "columntype"
    }, {
        "key": "datetime",
        "val": "日期时间",
        "type": "columntype"
    }, {
        "key": "decimal",
        "val": "数字",
        "type": "columntype"
    },
    //输入框类型 
    {
        "key": "text",
        "val": "普通文字",
        "type": "inputtype"
    }, {
        "key": "url",
        "val": "url地址",
        "type": "inputtype"
    }, {
        "key": "number",
        "val": "数字",
        "type": "inputtype"
    }, {
        "key": "email",
        "val": "电子邮箱",
        "type": "inputtype"
    }, {
        "key": "password",
        "val": "加密输入",
        "type": "inputtype"
    }, {
        "key": "file",
        "val": "文件上传",
        "type": "inputtype"
    }, {
        "key": "textarea",
        "val": "多行文本",
        "type": "inputtype"
    }, {
        "key": "richeditor",
        "val": "富文本输入(全)",
        "type": "inputtype"
    }, {
        "key": "simpleditor",
        "val": "富文本输入(简)",
        "type": "inputtype"
    }, {
        "key": "select",
        "val": "列表选择",
        "type": "inputtype"
    }, {
        "key": "checkbox",
        "val": "复选框",
        "type": "inputtype"
    }, {
        "key": "radio",
        "val": "单选框",
        "type": "inputtype"
    }, {
        "key": "switch",
        "val": "开关",
        "type": "inputtype"
    }, {
        "key": "date",
        "val": "日期",
        "type": "inputtype"
    }, {
        "key": "time",
        "val": "时间",
        "type": "inputtype"
    }, {
        "key": "datetime",
        "val": "日期时间",
        "type": "inputtype"
    },
];
ggcmsCfg.getDateExp = function(str, val) {
    var d = new Date();
    //以当前值为标准
    if (/v$/.test(str) && val) {
        d = new Date(val);
        str = str.replace(/v$/, "");
    }
    if (/c$/.test(str)) {
        str = str.replace(/c$/, "");
        var dd = str.split(/[,;:|]/);
        if (dd.length == 2) {
            d = new Date(dd[0]);
            str = dd[1];
        }
    }
    //或取增减天数
    var v = parseInt(str);
    //月
    if (/m$/.test(str)) {
        return d.setMonth(d.getMonth() + v);
    }
    //年
    if (/y$/.test(str)) {
        return d.setFullYear(d.getFullYear() + v);
    }
    //天
    if (/d$/.test(str)) {
        return d.setDate(d.getDate() + v);
    }
    return Date.parse(d);
}
ggcmsCfg.timePicker = {
    hourStep: 1,
    minuteStep: 1,
}