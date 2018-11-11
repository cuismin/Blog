var NavTitleTemplate = `<span class="layui-breadcrumb" style="visibility: visible;">   
                            <span onclick="Blog.getArticle().article()">网站首页</span>  
                            <span lay-separator="">/</span>  
                            <span onclick="Blog.getArticle().article()">博客文章</span>  
                            <span lay-separator="">/</span>     
                            <cite><#Title></cite>
                         </span>`

//文章类型
var ArticletypeTemplate =  `<div class="layui-col-xs6" uuid="<#uuid>">
                                <div class="grid-demo">
                                    <p><i class="layui-icon"></i><#Type></p>
                                </div>
                            </div>`;
//文章列表
var ArticleTemplate = ` <div class="layui-collapse layui-panel layui-article" ArticleOneOfTheLists>
                            <div class="layui-colla-item">
                                <div class="layui-colla-content layui-show layui-article">
                                    <fieldset class="layui-elem-field layui-field-title">
                                        <legend class="center-to-head">
                                            <span class="layui-badge layui-bg-green"><#ArticleType></span>
                                            <a rhref="/article/info/?" uuid="<#uuid>" title="<#title>"><#title></a>
                                        </legend>
                                        <div class="layui-field-box has-pic">
                                            <div class="layui-row layui-col-space10">
                                                <div class="layui-col-lg10 layui-col-mdContent layui-col-sm10 layui-col-xs12">
                                                   &emsp;&emsp;<#content>
                                                    <a class="loading" rhref="/article/info/?"  uuid="<#uuid>" title="<#title>">
                                                        更多
                                                        <i class="fa fa-angle-double-right">
                                                        </i>
                                                    </a>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="operation">
                                            <div class="tags">
                                                <span class="layui-badge-rim" uuid="<#userId>" type="User"><i class="fa fa-tag"></i>一花一世界</span>
                                                <span class="layui-badge-rim" uuid="<#ArticleTypeId>" type="ArticleType"><i class="fa fa-tag"></i><#ArticleType>博客</span>
                                                <span class="layui-badge-rim layui-hide"><i class="fa fa-tag"></i>开源</span>
                                            </div>
                                            <div class="info">
                                                <span class="datetime"><i class="fa fa-bullhorn"></i><#CreateDate></span>
                                            </div>
                                        </div>
                                    </fieldset>
                                </div>
                            </div>
                        </div>`;


var NoDataTemplate  = ` <div class="layui-flow-more" NoDataTemplate><a">
                            <cite>没有更多了</cite></a>
                        </div>`;


var NoteArticleTemplate =`<div class="layui-colla-item">
                                <h2 class="layui-colla-title" uuid="<#uuid>"><#Title></h2>
                                <div class="layui-colla-content" id="<#uuid>">
                                    <textarea  style="display:none;"><#content></textarea>
                                </div>
                          </div>`