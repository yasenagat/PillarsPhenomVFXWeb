require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "jslib/jquery-1.4.3",
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

require(['jquery', 'validate'], function($) {
    $(function(){
		//验证手机
		$.validator.addMethod("phone", function(value, element) {
    			var tel = /^(130|131|132|133|134|135|136|137|138|139|150|153|157|158|159|180|187|188|189)\d{8}$/;
    			return tel.test(value) || this.optional(element);
		}, "请输入正确的手机号码");

		$('#form1').validate({
			//required:验证规则
			rules:{
				Email:{required:true, email:true},
				UserName:{required:true, minlength:2},
				Phone:{required:true, phone:true},
				UserAuthority:{required:true}
			},
			//自定义错误提示信息
			messages:{
				Email:{required:"Email必须输入", email:"必须是合法的邮件地址"},
				UserName:{required:"用户名必须输入", minlength:"长度至少为2个字符"},
				Phone:{required:"电话必须输入"},
				UserAuthority:{required:"分组必须输入"}
			},
			errorElement:"span", //span是一个html标记，用来放置错误提示信息
			success:function(label) {
				label.text("").addClass("success");
			}
		});
	});

});
