require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "jslib/jquery-1.10.2.min",
        "domready": "jslib/domReady",
		"validate": "jslib/jquery.validate.min"
    },
	shim: {
        'validate': {
            deps: ['jquery']
        }
    }
});

require(['domready', 'jquery', 'validate'], function(doc, $) {
    doc(function(){
		var mail = "";
		if (document.URL.indexOf("?") != -1) {
			var getval = document.URL.split('?')[1];
			if(getval.indexOf("=") != -1) {
				mail = getval.split("=")[1];
			}
		}
		if(mail != ""){
			$.post('/user_sel', {Email:mail}, function(data) {
				if(data.FeedbackCode == undefined){
					$("input[name^='Email']").val(data.Email);
					$("input[name^='Email']").attr("readonly", true);
					$("input[name^='UserName']").val(data.DisplayName);
					$("input[name^='Phone']").val(data.Phone);
					$("#UserAuthority").val(data.UserAuthority);
					$("input[name^='FilePath']").val(data.FilePath);
				}

			}, "json");

			$("#save").click(function () {
				$('#form1').attr("action", "user_upd");
			});
		}else{
			$("#save").click(function () {
				$('#form1').attr("action", "user_add");
			});
		}

		//验证手机
		$.validator.addMethod("phone", function(value, element) {
    			var tel = /^(130|131|132|133|134|135|136|137|138|139|150|153|157|158|159|170|180|187|188|189)\d{8}$/;
    			return tel.test(value) || this.optional(element);
		}, "请输入正确的手机号码");

		$('#form1').validate({
			//required:验证规则
			rules:{
				Email:{required:true, email:true},
				UserName:{required:true, minlength:2},
				Phone:{required:true, phone:true}
			},
			//自定义错误提示信息
			messages:{
				Email:{required:"Email必须输入", email:"必须是合法的邮件地址"},
				UserName:{required:"用户名必须输入", minlength:"长度至少为2个字符"},
				Phone:{required:"电话必须输入", phone:"请输入正确的手机号码"}
			},
			errorElement:"span", //span是一个html标记，用来放置错误提示信息
			success:function(label) {
				label.text("").addClass("success");
			}
		});
	});
});
