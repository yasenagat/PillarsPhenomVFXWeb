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

var flag = "add";
// TODO 缩略图后期实现，前端解析为：base64 image
var pic = "data:image/png;base64,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";

require(['domready', 'jquery', 'validate'], function(doc, $) {
    doc(function() {
		// 搜索(条件仅限:"项目名","负责人","项目类型")
		$(".inpsearch").click(function() {
			var str = $(".inptext").val();
			if (str == "") {
				alert("请输入搜索条件!");
				return;
			}
			$("#searchflag").val("Y");
			projectfind_ajax(str, function(data) {
				if(data.FeedbackCode == 0) {
					var rs = JSON.parse(data.Data);
					if (rs == null || rs.length == 0){
						alert("没有搜索到符合条件的项目信息!");
						return;
					}
					$(".pro").html("");
					createProject(rs);
				}
			});
		});
		// 添加项目
		$(".inpbutton").click(function() {
			flag = "add";
			var height = $(window).height();
			var width = $(window).width();
			$(".outer").css({"height":height+"px","width":width+"px"});
			$(".formdiv").css({"left":(width/2)-200+"px","top":(height/2)-200+"px"});
			againjs();
			$(".outer").show(500);
			$(".formdiv").show(500);
		});
		$(".outer").click(function(){
			$(".outer").hide(500);
			$(".formdiv").hide(500);
		});
		$(".again").click(function(){
			againjs();
		});

		//url跳转到素材页面
		$(".pro").on("click","table",function(){
			var code = $(this).find(".ProjectCode_").val();
			window.location.href="editoral.html?code="+code;
		});

		// 操作
		$(".pro").on("click",".cz",function(){
			flag = "upd";
			var thistab = $(this).parent().children();
			var xmm = thistab.find(".ProjectName_").html().trim();
			var fzr = thistab.find(".ProjectLeader_").html().trim();
			var kssj = thistab.find(".StartDatetime_").html().trim();
			var jssj = thistab.find(".EndDatetime_").html().trim();
			var bz = thistab.find(".ProjectDetail_").html().trim();
			var xmlx = thistab.find(".ProjectType_").html().trim();
			var xmdm = thistab.find(".ProjectCode_").val().trim();
			var height = $(window).height();
			var width = $(window).width();
			$(".ProjectName").val(xmm);
			$(".ProjectLeader").val(fzr);
			$(".StartDatetime").val(kssj);
			$(".EndDatetime").val(jssj);
			$(".ProjectDetail").val(bz);
			$(".ProjectType").val(xmlx);
			$(".ProjectCode").val(xmdm);
			$(".outer").css({"height":height+"px","width":width+"px"});
			$(".formdiv").css({"left":(width/2)-200+"px","top":(height/2)-200+"px"});
			$(".outer").show(500);
			$(".formdiv").show(500);
		});

		// 删除
		$(".pro").on("click",".del",function(){
			var thistab = $(this).parent().children();
			var xmdm = thistab.find(".ProjectCode_").val().trim();
			if(confirm("关闭项目后不可恢复，确定要删除吗？")) {
				projectdel_ajax(xmdm, function(data) {
					if(data.FeedbackCode == 0) {
						$(".tab"+xmdm).remove();
		            } else {
						// TODO 删除失败
		            }
				});
			}
		})

		// 表单验证
		var validator = $('#projectForm').validate({
			//required:验证规则
			rules:{
				ProjectName:{required:true, minlength:2},
				ProjectLeader:{required:true, minlength:2},
				StartDatetime:{required:true},
				EndDatetime:{required:true},
				ProjectDetail:{required:true},
				ProjectType:{required:true}
			},
			//自定义错误提示信息
			messages:{
				ProjectName:{required:"项目名称必须填写", minlength:"长度至少为2个字符"},
				ProjectLeader:{required:"项目负责人必须填写", minlength:"长度至少为2个字符"},
				StartDatetime:{required:"开始时间必须填写"},
				EndDatetime:{required:"结束时间必须填写"},
				ProjectDetail:{required:"备注必须填写"},
				ProjectType:{required:"项目类型必须填写"}
			},
			errorElement:"span", //span是一个html标记，用来放置错误提示信息
			success:function(label) {
				label.text("").addClass("success");
			}
		});

		// 提交
		$(".submit").click(function() {
			if(validator.form()) {
				var pc = $(".ProjectCode").val();
				var pn = $(".ProjectName").val();
				var pi = pic;
				var pl = $(".ProjectLeader").val();
				var pt = $(".ProjectType").val();
				var sd = $(".StartDatetime").val();
				var ed = $(".EndDatetime").val();
				var pd = $(".ProjectDetail").val();
				var action = "/project_upd";
				if(flag == "add"){
					action = "/project_add";
				}

				projectsave_ajax(action, pc, pn, pi, pl, pt, sd, ed, pd, function(data) {
					if(data.FeedbackCode == 0) {
						if(flag == "add"){
							var rs = JSON.parse(data.Data);
							pc = rs["ProjectCode"];
							var addhtml = "<div class='protab'><table width='100%' border='0' cellspacing='0' cellpadding='0' class='tab" + pc + "'><tr><td rowspan='3' width='150' height='200'><img width='120' height='160' src='" + pic + "'></td><td width='15%' align='right'>项目名：<input type='hidden' class='ProjectCode_' value='" + pc + "'></td><td class='ProjectName_'>" + pn + "</td><td width='15%' align='right'>负责人：</td><td class='ProjectLeader_'>" + pl + "</td></tr><tr><td align='right'>开始时间：</td><td class='StartDatetime_'>" + sd + "</td><td align='right'>结束时间：</td><td class='EndDatetime_'>" + ed + "</td></tr><tr><td align='right'>项目类型：</td><td class='ProjectType_'>" + pt + "</td><td align='right'>备注：</td><td class='ProjectDetail_'>" + pd + "</td></tr></table><div class='cz'>操作</div><div class='del'>删除</div></div>";
							$(".pro").prepend(addhtml);
							againjs();
							$(".outer").hide(500);
							$(".formdiv").hide(500);
						}else{
							$(".tab"+pc+" .ProjectName_").html(pn);
							$(".tab"+pc+" .ProjectLeader").html(pl);
							$(".tab"+pc+" .ProjectType").html(pt);
							$(".tab"+pc+" .StartDatetime").html(sd);
							$(".tab"+pc+" .EndDatetime").html(ed);
							$(".tab"+pc+" .ProjectDetail").html(pd);
							againjs();
							$(".outer").hide(500);
							$(".formdiv").hide(500);
						}
		            } else {
						// TODO 保存失败显示位置
		                alert("保存失败，请稍后重试！");
		            }
	        	});
			}
		});

		$(window).scroll(function() {
			if($("#searchflag").val() == "Y") {
				return;
			}
			var scrollTop = $(this).scrollTop();
			var scrollHeight = $(document).height();
			var windowHeight = $(this).height();
			if(scrollTop + windowHeight == scrollHeight) {
				var start = $(".pro").children(".protab").length;
				var end = start + 6;
				projectload_ajax(start, end, function(data) {
					if(data.FeedbackCode == 0) {
						var rs = JSON.parse(data.Data);
						createProject(rs);
		            }
	        	});
			}
		});

		// 创建项目信息
		var createProject = function(rs) {
			if (rs == null || rs.length == 0){
				return;
			}
			var addhtml;
			for (var i=0; i<rs.length; i++) {
				addhtml = "";
				addhtml += "<div class='protab'>";
				addhtml += "<table width='100%' border='0' cellspacing='0' cellpadding='0' class='tab" + rs[i]["ProjectCode"] + "'>";
				addhtml += "<tr><td rowspan='3' width='150' height='200'><img width='120' height='160' src='" + rs[i]["Picture"] + "'></td>";
				addhtml += "<td width='15%' align='right'>项目名：<input type='hidden' class='ProjectCode_' value='" + rs[i]["ProjectCode"] + "'></td>";
				addhtml += "<td class='ProjectName_'>" + rs[i]["ProjectName"] + "</td>";
				addhtml += "<td width='15%' align='right'>负责人：</td>";
				addhtml += "<td class='ProjectLeader_'>" + rs[i]["ProjectLeader"] + "</td></tr>";
				addhtml += "<tr><td align='right'>开始时间：</td><td class='StartDatetime_'>" + rs[i]["StartDatetime"] + "</td>";
				addhtml += "<td align='right'>结束时间：</td>";
				addhtml += "<td class='EndDatetime_'>" + rs[i]["EndDatetime"] + "</td>";
				addhtml += "</tr><tr><td align='right'>项目类型：</td>";
				addhtml += "<td class='ProjectType_'>" + rs[i]["ProjectType"] + "</td>";
				addhtml += "<td align='right'>备注：</td>";
				addhtml += "<td class='ProjectDetail_'>" + rs[i]["ProjectDetail"] + "</td></tr>";
				addhtml += "</table><div class='cz'>操作</div><div class='del'>删除</div></div>";
				$(".pro").append(addhtml);
			}
		}
		// 重置表单内容
		var againjs = function() {
			$(".ProjectCode").val("");
			$(".ProjectName").val("");
			$(".ProjectLeader").val("");
			$(".ProjectType").val("");
			$(".StartDatetime").val("");
			$(".EndDatetime").val("");
			$(".ProjectDetail").val("");
		}
		// 项目删除
		var projectdel_ajax = function(xmdm, callback) {
	        $.post("/project_del", {ProjectCode: xmdm},
	            function(data) {
	                callback(data);
	            },
	            "json"
	        );
	    }
		// 项目保存,新增和修改
		var projectsave_ajax = function(action, pc, pn, pi, pl, pt, sd, ed, pd, callback) {
	        $.post(action,
				{
					ProjectCode: pc,
					ProjectName: pn,
					Picture: pi,
					ProjectLeader: pl,
					ProjectType: pt,
					StartDatetime: sd,
					EndDatetime: ed,
					ProjectDetail: pd
				},
	            function(data) {
	                callback(data);
	            },
	            "json"
	        );
	    }
		// 瀑布流加载项目信息
		var projectload_ajax = function(start, end, callback) {
	        $.post("/project_load", {Start: start, End: end},
	            function(data) {
	                callback(data);
	            },
	            "json"
	        );
	    }
		// 搜索项目
		var projectfind_ajax = function(args, callback) {
	        $.post("/project_find", {Args: args},
	            function(data) {
	                callback(data);
	            },
	            "json"
	        );
	    }
	});
});
