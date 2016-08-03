<!DOCTYPE html>
<html lang="en">

<head>
    <title>京奇科技 - 专业的互联网开发团队</title>
	[#include "/include/mate.ftl"]
    [#include "/include/webcss.ftl"]
    <link href="${path}/resources/plugins/cubeportfolio/css/cubeportfolio.css" rel="stylesheet" type="text/css" />
    <link href="${path}/resources/css/view.css" rel="stylesheet" />
</head>

<body data-spy="scroll" data-target=".navbar" data-offset="75">
    <div id="container" class="container">
        <div class="page-header">
            <h1 class="text-center">PORTFOLIO</h1>
        </div>
        <div class="row">
            <div class="col-sm-3 col-md-2 sidebar">
                <ul class="nav nav-sidebar">
					[@categoryTagComment fatherId="035c51e3e5024f02b38efba0c80788a1" ]
                        [#list category_list as citem]
							[#if categoryInfo.id==citem.id]
                            <li class="active"><a href="${getCategoryUrl ("", citem.id)}">${citem.categoryName}</a></li>
							[#else]
                            <li><a href="${getCategoryUrl ("", citem.id)}">${citem.categoryName}</a></li>
							[/#if]
						[/#list]
                    [/@categoryTagComment]
                </ul>
            </div>
            <div class="col-sm-9 col-md-10 ">
                <div class="main">
                    <div class="row title">
                        <h2 class="sub-header">${categoryInfo.categoryName}</h2>
                    </div>
                    <div class="row">
                        <div class="portfolio-content portfolio-1">
                            <div id="js-grid-juicy-projects" class="cbp">
							 [@indexTagComment id=categoryInfo.id pageSize=10 imgShow=Y]
							 [#list indexActiclesList  as item] 
								 <div class="cbp-item print motion">
                                    <a href="${path+"/"+item.titleImg}" class="cbp-caption cbp-lightbox" data-title="${item.articleTitle}<br>${categoryInfo.categoryName}-截图">
                                        <div class="cbp-caption-defaultWrap">
                                            <img src="${path+"/"+item.titleImg}" alt=""> </div>
                                        <div class="cbp-caption-activeWrap">
                                            <div class="cbp-l-caption-alignCenter">
                                                <div class="cbp-l-caption-body">
                                                    <div class="cbp-l-caption-title">${item.articleTitle}</div>
                                                </div>
                                            </div>
                                        </div>
                                    </a>
                                </div>
							  [/#list]
							  [/@indexTagComment]

                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </div>
        <footer>
            <div class="text-center copyright">2015 &copy;开封京奇科技有限公司.保留所有权利</div>
        </footer>

    </div>
    <!-- Home Section -->

    <!-- End Home Section -->

    <!-- Navigation -->
    <nav class="navbar navbar-inverse navbar-fixed-top" id="navigation">
        <div class="container-fluid">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div id="logo">
                <a class="external" href="index.html"></a>
            </div>
            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse">
                <ul class="nav navbar-nav">
				  [@categoryTagComment fatherId="0" ]
					 [#list category_list as f]
					 [#if f.isNavigation!="N"]
						<li class="menu-${f_index+1}" data-cid="${f.id}">
							   <a class="colapse-menu1" href= '${path+"/web/index"+f.extattrib}' >${f.categoryName}</a>
						</li>
					[/#if]
                    [/#list]
                    [/@categoryTagComment]
                </ul>

            </div>
            <!-- /.navbar-collapse -->
        </div>
        <!-- /.container-fluid -->
    </nav>
    <!--/Navigation -->
	[#include "/include/webjs.ftl"]
    <script src="${path}/resources/plugins/cubeportfolio/js/jquery.cubeportfolio.min.js" type="text/javascript"></script>
    <script src='${path}/resources/js/web/view.js' type="text/javascript"></script>
</body>

</html>