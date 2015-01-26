require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "jslib/jquery-1.10.2.min",
        "md5": "jslib/md5",
		"validate": "jslib/jquery.validate"
    },
	shim: {
		'jquery': {
			exports: '$'
		},
        'validate': {
            deps: ['jquery']
        }
    }
});

require(['jquery', 'validate', 'md5'], function($) {
    $(function(){
		var validator = $('#form1').validate({
			//required:验证规则
			rules:{
				UserName:{required:true, email:true},
				UserPassword:{required:true, minlength:6}
			},
			//自定义错误提示信息
			messages:{
				UserName:{required:"请输入用户名", email:"必须是合法的邮件地址"},
				UserPassword:{required:"请输入密码", minlength:"长度至少为6个字符"}
			},
			errorElement:"span", //span是一个html标记，用来放置错误提示信息
			success:function(label) {
				label.text("").addClass("success");
			}
		});

		var login_core = function(UserID, passwd, callback) {
	        $.post('/login_action', {Email:UserID, Password:passwd},
	            function(data) {
	                callback(data);
	            },
	            "json"
	        );
	    }

        $('#loginBtn')[0].onclick = function() {
			if(validator.form()) {
				// 密码使用MD5（32）加密
				passwd = faultylabs.MD5($('#userPassword')[0].value);
				login_core($('#userName')[0].value, passwd, function(data) {
					if(data.FeedbackCode == 0) {
		                window.location = "/" + data.FeedbackText;
		            } else {
		                alert("用户名或密码错误");
		            }
		        });
			}
		}
    });
});
