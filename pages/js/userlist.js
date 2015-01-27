require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "jslib/jquery-1.10.2.min",
        "domready": "jslib/domReady"
    }
});

var updateuser;
var alter_core;
var deleteuser;

require(['domready', 'jquery'], function(doc, $) {
    doc(function(){
		updateuser = function (email) {
			window.location = "/addUser.html?a=" + email;
		}

		alter_core = function (alter, email, callback) {
		    window.res = {};
		    $.post(alter, {Email: email},
		        function (data) {
		            callback(data);
		        },
		        "json"
		    );
		}

		deleteuser = function (email) {
			alter_core('/user_del', email, function(data) {
		        if(data.FeedbackCode == 0) {
		            window.location = "/user_list";
		        } else {
		            alert("删除操作失败，请稍后重试！");
		        }
			});
		}
	});
});
