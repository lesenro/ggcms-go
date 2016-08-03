var preView = {
    template: '<div class="imgPreview" id="{id}" style="display:none;"><div class="close">&times;</div><div class="img"><img src="{url}"/></div></div>',
    show: function name(url) {
        var self = this;
        if (url.indexOf("://") == -1) {
            url = "/" + url;
        }
        var rimg = new Image();
        rimg.src = url;

        if (rimg.complete) {
            self.view(rimg.width, rimg.height, url);
            return;
        }
        rimg.onload = function() {
            self.view(rimg.width, rimg.height, url);
            return;
        };
    },
    view: function(iw, ih, url) {
        var per = 0.8;
        var id = "ipv_" + Date.now();
        var html = this.template.replace("{id}", id).replace("{url}", url);
        $("body").append(html);
        $("#" + id).fadeIn();
        var close = "#" + id + " .close";
        var img = "#" + id + " .img img";
        var bw = Math.floor($(window).width() * per);
        var bh = Math.floor($(window).height() * per);
        var nw = iw;
        var nh = ih;
        //图宽超出可视区
        if (iw > bw) {
            nw = bw;
            nh = Math.floor(ih * nw / iw);
        }
        //图高超出可视区
        if (nh > bh) {
            nh = bh;
            nw = Math.floor(iw * nh / ih);
        }
        if (nw > 0 && nh > 0) {
            $(img).width(nw);
            $(img).height(nh);
        }
        $(close + "," + img).click(function() {
            $("#" + id).fadeOut(function() {
                $("#" + id).remove();
            });
        });
    }
}