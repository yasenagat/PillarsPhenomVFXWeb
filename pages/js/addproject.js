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
		var code = "";
		if (document.URL.indexOf("?") != -1) {
			var getval = document.URL.split('?')[1];
			if(getval.indexOf("=") != -1) {
				code = getval.split("=")[1];
			}
		}
		if(code != ""){
			$.post('/queryproject', {ProjectCode:code}, function(data) {
				if(data.FeedbackCode == undefined){
					$("input[name^='ProjectCode']").val(data.ProjectCode);
					$("input[name^='ProjectName']").val(data.ProjectName);
					$("input[name^='ProjectDescribe']").val(data.ProjectDescribe);
					$("input[name^='ProjectLeader']").val(data.ProjectLeader);
				}

			}, "json");

			$("#save").click(function () {
				$('#form1').attr("action", "updateproject");
			});
		}else{
			$("#save").click(function () {
				$('#form1').attr("action", "addproject");
			});
		}

		$('#form1').validate({
			//required:验证规则
			rules:{
				ProjectName:{required:true, minlength:2},
				ProjectDescribe:{required:true, minlength:2},
				ProjectLeader:{required:true}
			},
			//自定义错误提示信息
			messages:{
				ProjectName:{required:"项目名称必须填写", minlength:"长度至少为2个字符"},
				ProjectDescribe:{required:"项目描述必须填写", minlength:"长度至少为2个字符"},
				ProjectLeader:{required:"项目负责人必须填写"}
			},
			errorElement:"span", //span是一个html标记，用来放置错误提示信息
			success:function(label) {
				label.text("").addClass("success");
			}
		});
	});

});
