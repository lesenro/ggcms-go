<!DOCTYPE html>
<!--[if IE 8]> <html lang="en" class="ie8 no-js"> <![endif]-->
<!--[if IE 9]> <html lang="en" class="ie9 no-js"> <![endif]-->
<!--[if !IE]><!-->
<html lang="en">
<!--<![endif]-->
<!-- BEGIN HEAD -->

<head>
    [#include "/include/mate.ftl"]
    <title>${webActicles.articleTitle}-开封天一干混砂浆有限公司</title>
<link rel="stylesheet" href="${path}/resources/plugins/bootstrap/css/bootstrap.min.css">
<link rel="stylesheet" href="${path}/resources/css/components-md.css">    
    [#include "/include/webcss.ftl"]
</head>
<!-- END HEAD -->

<body class="">
    <div id="wrap">
        [#include "/include/webheader.ftl"]
        <section class="section sectionView bg-light-gray">
            <div class="container">
                <div class="row">
                <div class="col-sm-3">
                        <ul class="list-group side-menu">
                            <a data-id="095cf059bc9d4bfe96dba24b9699d85f" href="${path}/web/findByArticlesCategoryId?id=095cf059bc9d4bfe96dba24b9699d85f&pageNum=1&type=one" class="list-group-item"> 关于我们  </a>
                            <a data-id="283aad0799e6404aa11e21698620f988" href="${path}/web/findByArticlesCategoryId?id=283aad0799e6404aa11e21698620f988&pageNum=1&type=one" class="list-group-item"> 联系我们  </a>
                            <a data-id="865ef377fd3f42bfb9321a6f4fa8df6d" href="${path}/web/findByArticlesCategoryId?id=865ef377fd3f42bfb9321a6f4fa8df6d&pageNum=1&type=one" class="list-group-item"> 组织架构  </a>
                        </ul>
                    </div>
                    <div class="col-sm-9">
                        <div class="main">
                            <div class="page-title clearfix" id="j-post-head">
                                <ol class="breadcrumb">
                                    <li class="home"><i class="fa fa-map-marker"></i> <a href="${path}/web/index">首页</a></li>
                                    <li>${oneCategory.categoryName}</li>
                                </ol>
                                <h3 class="title pull-left" id="j-title"><span>${oneCategory.categoryName}</span></h3>
                            </div>
                            <div class="entry">
                                <div class="entry-content">
                                    ${webActicles.articleContent}
                                </div>
                            </div>
                        </div>
                        </div>
                     </div>
                </div>
        </section>
        </div>
       [#include "/include/webfooter.ftl"] [#include "/include/webjs.ftl"]
    <script>
    $(".side-menu a").each(function(){
        var cid=$(this).attr("data-id");
        if(currentpath.indexOf(cid)>0){
                $(this).addClass("active");
                return false;
        }
     });
   </script>
</body>

</html>
