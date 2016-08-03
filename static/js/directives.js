/***
GLobal Directives
***/

// Route State Load Spinner(used on page or content load)
GgcmsApp.directive('ngSpinnerBar', ['$rootScope',
    function($rootScope) {
        return {
            link: function(scope, element, attrs) {
                // by defult hide the spinner bar
                element.addClass('hide'); // hide spinner bar by default

                // display the spinner bar whenever the route changes(the content part started loading)
                $rootScope.$on('$stateChangeStart', function() {
                    element.removeClass('hide'); // show spinner bar
                });

                // hide the spinner bar on rounte change success(after the content loaded)
                $rootScope.$on('$stateChangeSuccess', function() {
                    element.addClass('hide'); // hide spinner bar
                    $('body').removeClass('page-on-load'); // remove page loading indicator
                    Layout.setSidebarMenuActiveLink('match'); // activate selected link in the sidebar menu

                    // auto scorll to page top
                    setTimeout(function() {
                        App.scrollTop(); // scroll to the top on content load
                    }, $rootScope.settings.layout.pageAutoScrollOnLoad);
                });

                // handle errors
                $rootScope.$on('$stateNotFound', function() {
                    element.addClass('hide'); // hide spinner bar
                });

                // handle errors
                $rootScope.$on('$stateChangeError', function() {
                    element.addClass('hide'); // hide spinner bar
                });
            }
        };
    }
])

// Handle global LINK click
GgcmsApp.directive('a', function() {
    return {
        restrict: 'E',
        link: function(scope, elem, attrs) {
            if (attrs.ngClick || attrs.href === '' || attrs.href === '#') {
                elem.on('click', function(e) {
                    e.preventDefault(); // prevent link click for above criteria
                });
            }
        }
    };
});

// Handle Dropdown Hover Plugin Integration
GgcmsApp.directive('dropdownMenuHover', function() {
    return {
        link: function(scope, elem) {
            elem.dropdownHover();
        }
    };
});

GgcmsApp.directive('valiData', function($compile, $timeout) {
    return function(scope, element, attrs) {
        var forid = attrs["forid"] || attrs["name"];
        //取数据是设置还是模型
        var inputdata = scope.datamap || scope.data.modulesData;
        var opt = inputdata[forid].opts;
        var val = inputdata[forid].Value;
        if (opt.required) {
            attrs.$set('required', true);
        }
        if (opt.min > 0) {
            attrs.$set('min', opt.min);
        }
        if (opt.max > 0) {
            attrs.$set('max', opt.max);
        }
        if (opt.minLength > 0) {
            attrs.$set('ng-minlength', opt.minLength);
        }
        if (opt.maxLength > 0) {
            attrs.$set('ng-maxlength', opt.maxLength);
        }
        if (opt.pattern != "") {
            attrs.$set('ng-pattern', opt.pattern);
        }
        if (opt.minDate) {
            attrs.$set('min-date', ggcmsCfg.getDateExp(opt.minDate, val));
        }
        if (opt.maxDate) {
            attrs.$set('max-date', ggcmsCfg.getDateExp(opt.maxDate, val));
        }
        if (opt.type == "file") {
            attrs.$set('ngf-select', "upload($file,item.Varname,true)");
            attrs.$set('ng-model', forid);
            if (opt.fileSize > 0) {
                attrs.$set('ngf-max-size', opt.fileSize + "kb");
            }
            if (opt.extension != "") {
                attrs.$set('ngf-pattern', opt.extension);
            }
        }
        element.removeAttr("vali-data");
        $compile(element)(scope);
    }
});