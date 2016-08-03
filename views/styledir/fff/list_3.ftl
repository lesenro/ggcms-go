<!DOCTYPE html>
<!--[if IE 8]> <html lang="en" class="ie8 no-js"> <![endif]-->
<!--[if IE 9]> <html lang="en" class="ie9 no-js"> <![endif]-->
<!--[if !IE]><!-->
<html lang="en">
<!--<![endif]-->
<!-- BEGIN HEAD -->

<head>
    [#include "/include/mate.ftl"]
    <title>开封天一干混砂浆有限公司</title>
    [#include "/include/webcss.ftl"]
</head>
<!-- END HEAD -->

<body class="">
    <div id="wrap">
        [#include "/include/webheader.ftl"]
        <section class="section row sectionList bg-light-gray">
            <div class="container">
                <div class="post-list clearfix">
                    <div class="page-title clearfix" id="j-post-head">
                        <ol class="breadcrumb">
                            <li class="home"><i class="fa fa-map-marker"></i> <a href="${path}/web/index">首页</a></li>
                            <li>${oneCategory.categoryName}</li>
                        </ol>
                        <h3 class="title pull-left" id="j-title"><span>${oneCategory.categoryName}</span></h3>
                    </div>
                [#list page.list   as item] 
                    <article class="post-item clearfix">
                        <header class="article-title">
                        <div class="entry-meta pull-right">
                                <span><i class="fa fa-calendar"></i> ${item.addTime}<i></i></span>
                         </div>
                         <h2 class="entry-title">
                            <a href="${path}/web/findByArticlesId?id=${item.id}">${item.articleTitle}</a>
                         </h2>
                        </header>
                    </article>
                [/#list]                      
                <hr/>
                    [#include "/include/webpagination.ftl"] 
                </div>
            </div>
        </section>
    </div>
    [#include "/include/webfooter.ftl"] [#include "/include/webjs.ftl"]
</body>

</html>

