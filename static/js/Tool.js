var Tools = (function () {
    Service = {
        getServiceData:function (parameter) {
            var Result;
            $.ajax({
                async:parameter.async === undefined ? true : parameter.async ,
                type: parameter.type === undefined ? "POST" : parameter.type,
                url: parameter.url,
                data:parameter.data === undefined ? {} : parameter.data,
                dataType: parameter.dataType === undefined ? "JSON" : parameter.dataType,
                beforeSend:parameter.beforeSend === undefined ? function () {
                }:parameter.beforeSend,
                success:parameter.success === undefined ? function (data) {
                    Result = data;
                }:parameter.success,
                complete:parameter.complete === undefined ? function () {
                }:parameter.complete,
                error:function (error) {
                    console.error("403.2系统异常请稍后重试");
                }
            });
            if((parameter.async !== undefined && !parameter.async))
                return Result;
        },
    }
    var instance;
    function init(){
        return Service;
    }
    return {
        getInstance:function () {
            return !instance ? init() : instance;
        }
    }
})();

var tools = Tools.getInstance();