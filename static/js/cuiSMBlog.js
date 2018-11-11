(function ($) {
    $.fn.cuiSMBlog = function (options) {
        var defaults = {
            website:"http://www.cuism.com"
        };
        defaults = $.extend(defaults, options);//两个对象合并
        var _loadData = function () {
            cuismPageServlce.getServiceData({});
        };
        var cuismPageServlce = {
            render:function(){

            },
            //加载动画
            LoadAnimation:function(url) {
                $("section").removeClass("animated fadeInLeft");
                $("#section_iframe").attr("src",url);
                setTimeout(function () {
                    $("section").addClass("animated fadeInLeft");
                },10)
            },
            //更新title
            UpdateTitle:function (title) {
                localStorage.setItem("title",$("blockquote").html());
                $("blockquote").empty().html(title);
            },
            UpdateTitleNumber:function (num) {
                $("blockquote").find("span[class='sum-font']").empty().html(num);
            },
            addEventListener:function (element,event,callback) {
               $(element).on(event,callback);
            },
            getServiceData:function (parameter) {
                var Result;
                $.ajax({
                    async:parameter.async === undefined ? true : parameter.async ,
                    type: parameter.type === undefined ? "POST" : parameter.type,
                    url: parameter.url,
                    data:parameter.data === undefined ? {} : parameter.data,
                    dataType: parameter.dataType === undefined ? "JSON" : parameter.dataType,
                    beforeSend:parameter.beforeSend === undefined ? function () {
                        //console.info("加载......")
                    }:parameter.beforeSend,
                    success:parameter.success === undefined ? function (data) {
                        Result = data;
                    }:parameter.success,
                    complete:parameter.complete === undefined ? function () {
                        //console.info("加载......")
                    }:parameter.complete,
                    error:function (error) {
                        //console.error("403.2系统异常请稍后重试");
                    }
                });
                if((parameter.async !== undefined && !parameter.async))
                    return Result;
            },
            getArticle:function () {
                return{
                    article:function () {
                        $("blockquote").html(localStorage.getItem("title"));
                        cuismPageServlce.LoadAnimation("page/home*aside.html")
                    }
                };
            },
            getArticleType:function () {
                return{

                };
            }
        }
        return (function () {
            var instance;
            function init() {
                var section_iframe = document.getElementById("section_iframe");
                section_iframe.height = document.documentElement.clientHeight-122+"px";
                layui.use(['layer','util','element'], function(){
                    var util = layui.util ,element = layui.element , layer = layui.layer;
                    $(window).resize(function () {
                        window.location.reload(true);
                    });
                    //固定块
                    util.fixbar({
                        bar1: true
                        ,bar2: true
                        ,css: {right: 20, bottom: 30}
                        ,bgcolor: '#393D49'
                        ,click: function(type){
                            if(type === 'bar1'){
                                layer.msg('icon是可以随便换的')
                            } else if(type === 'bar2') {
                                layer.msg('两个bar都可以设定是否开启')
                            }
                        }
                    });
                });
                $(window).resize(function () {
                    window.location.reload(true);
                });
                return cuismPageServlce;
            }
            return {
                getInstance : function () {
                    return !instance ? init() : instance;
                }
            };
        })();
    }
})(jQuery);
