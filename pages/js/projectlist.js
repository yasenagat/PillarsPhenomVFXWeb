require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "jslib/jquery-1.10.2.min",
        "domready": "jslib/domReady"
    }
});

var updateproject;
var alter_core;
var deleteproject;

require(['domready', 'jquery'], function(doc, $) {
    doc(function(){
		updateproject = function (code) {
			window.location = "/addproject.html?a=" + code;
		}

		alter_core = function (alter, code, callback) {
		    window.res = {};
		    $.post(alter, {ProjectCode: code},
		        function(data) {
		            callback(data);
		        },
		        "json"
		    );
		}

		deleteproject = function (code) {
			alter_core('/project_del', code, function(data) {
		        if(data.FeedbackCode == 0) {
		            window.location = "/project_list";
		        } else {
		            alert("删除操作失败，请稍后重试！");
		        }
		    });
		}
	});
});
