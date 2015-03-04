//得到url地址中code参数
var projectCode = "";
//url参数获取
var getUrlParam = function(name){
	var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
	var r = window.location.search.substr(1).match(reg);  //匹配目标参数
	if (r!=null) return unescape(r[2]); return null; //返回参数值
}
//Library数据
var librarys_ajax = function(pc){
	$.post("/editoral_library",
		{ProjectCode: pc},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				// 组织页面显示
				for(var i=0; i<rs.length; i++){
					$(".library").append("<li class='li1'><a href='javascript:void(0);' class='"+rs[0]["LibraryCode"]+"'>"+rs[0]["LibraryName"]+"</a></li>");
				}
				setTimeout(function(){
					library_click();
				},1000);
			}
        },
        "json"
    );
}
var library_click = function(){
	$(".library").on("click","a",function(){
		var code = $(this).attr("class");
		materials_ajax("1", projectCode, code, 0, 10, function(data){
			if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				$(".strdiv .videodiv").html("");
				fileList_create(rs);
				$("#pageflag").val("Library");
				$("#pagecode").val(code);
			}
		});
	});
}
//素材组数据
var filetypes_ajax = function(pc){
	$.post("/editoral_filetype",
		{ProjectCode: pc},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				// 组织页面显示
				for(var i=0; i<rs.length; i++){
					$("#materialGroup").html("<li><a href='javascript:void(0);'>"+rs[0]+"</a></li>");
				}
				setTimeout(function(){
					$("#materialGroup").on("click","a",function(){
						var code = $(this).html();
						materials_ajax("2", projectCode, code, 0, 10, function(data){
							if(data.FeedbackCode == 0) {
								var rs = JSON.parse(data.Data);
								$(".strdiv .videodiv").html("");
								fileList_create(rs);
								$("#pageflag").val("Group");
								$("#pagecode").val(code);
							}
						});
					});
				},1000);
			}
        },
        "json"
    );
}
//左侧列表点击
var materials_ajax = function(flag, pc, lc, start, end, callback){
	$.post("/editoral_library_materials",
		{
			Flag: flag,
			ProjectCode: pc,
			LibraryCode: lc,
			Start: start,
			End: end
		},
        function(data) {
            callback(data);
        },
        "json"
    );
}
//搜索
var find_materials_ajax = function(pc, args, callback){
	$.post("/editoral_find_materials",
		{ProjectCode: pc, Args: args},
        function(data) {
			callback(data);
        },
        "json"
    );
}
//素材信息查询
var materialInfo_ajax = function(mc, callback){
	$.post("/editoral_material", {MaterialCode: mc},
        function(data) {
			callback(data);
        },
        "json"
    );
}
//列表数据创建
var fileList_create = function(rs){
	for(var i=0;i<rs.length;i++){
		var liInfo = "";
		if(rs[i]["DpxPath"] == "Y") {
			liInfo += "<li>DPX</li>";
		}
		if(rs[i]["JpgPath"] == "Y") {
			liInfo += "<li>JPG</li>";
		}
		if(rs[i]["MovPath"] == "Y") {
			liInfo += "<li>Mov</li>";
		}
		var html = '<span class="videostr">';
		html += '<input type="hidden" class="sourceid" value="'+rs[i]["MaterialCode"]+'">';
		html += '<input class="check" name="" type="checkbox" value="">';
		html += '<input class="play" type="button" value="回放">';
		html += '<div class="downdiv">';
		html += '<input class="downl" type="button" value="下载">';
		html += '<span class="disnone">';
		html += '<ul class="'+rs[i]["MaterialCode"]+'">';
		html += '<li>Source</li>'+liInfo+'</ul></span>';
		html += '</div><div class="files">';
		html += '<span class="name">'+rs[i]["MaterialName"]+'</span>';
		html += '<span class="format">'+rs[i]["MaterialType"]+'</span>';
		html += '<span class="long">'+rs[i]["Length"]+'</span></div></span>';
		$(".strdiv .videodiv").append(html);
	}
	$(".files").on("click",function(){
		var rightdivabs = $(".rightdivabs").css("display");//当前是否悬浮
		var sourceid = $(this).siblings(".sourceid").val();//选中的id
		var absid = $(".float .sourceid").val();//右边悬浮的id
		if(rightdivabs == "block" && sourceid == absid) {
			$(".rightdivabs").css("display","none");
		} else {
			$(".rightdivabs").css("display","block");
			$(".float .sourceid").val(sourceid);
			//根据sourceid从后台查询数据
			materialInfo_ajax(sourceid, function(data){
				if(data.FeedbackCode == 0) {
					var rs = JSON.parse(data.Data);
					$(".float .roughimg img").attr("src", rs["Picture"]);//缩略图
					$(".float .names").html(rs["MaterialName"]);//素材名
					$(".float .size").html(rs["Size"]);//尺寸
					$(".float .long").html(rs["Length"]);//长度
					$(".float .speed").html(rs["VideoAudioFramerate"]);//帧速率
					var metaData = JSON.parse(rs["MetaData"]);
					var metasHtml = "";
					for (var s in metaData) {
						// TODO div加样式带滚动条,小区域的
						metasHtml += '<div><span>' + s + '</span>';
						metasHtml += " : ";
						metasHtml += '<span>' + metaData[s] + '</span></div>';
					}
					$(".float .metadata").html(metasHtml);//基本信息
				}
			});
		}
	});
}
//Library添加
var addLibrary_ajax = function(ln, lp, dp, jp, mp, pc, callback){
	$.post("/editoral_library_add",
		{
			LibraryName: ln,
			LibraryPath: lp,
			DpxPath: dp,
			JpgPath: jp,
			MovPath: mp,
			ProjectCode: pc
		},
        function(data) {
            callback(data);
        },
        "json"
    );
}
//下载文件
var downloadCheck_ajax = function(code, type) {
    $.post("/editoral_download_file_check", JSON.stringify({MaterialCode: code, SourceType: type}),
        function(data) {
            if(data.FeedbackCode == 0) {
				post("/editoral_download_file", {MaterialCode: code, SourceType: type});
			}else{
				alert("下载的文件不存在，请确认库的文件有效性！");
			}
        },
        "json"
    );
}
//下载文件（此方式不可用）
var downloadFile = function(code, type) {
    $.post("/editoral_download_file", {MaterialCode: code, SourceType: type});
}
//虚拟表单提交，下载文件使用
var post = function(URL, PARAMS) {
    var temp = document.createElement("form");
    temp.action = URL;
    temp.method = "post";
    temp.style.display = "none";
    for (var x in PARAMS) {
        var opt = document.createElement("textarea");
        opt.name = x;
        opt.value = PARAMS[x];
        temp.appendChild(opt);
    }
    document.body.appendChild(temp);
    temp.submit();
    return temp;
}


$(function(){
	projectCode = getUrlParam("code");
	if (projectCode !== ""){
		librarys_ajax(projectCode);
		filetypes_ajax(projectCode);
	}
	// All列表点击
	$(".li1").on("click", "a", function(){
		var code = $(this).html();
		materials_ajax("2", projectCode, code, 0, 10, function(data){
			if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				$(".strdiv .videodiv").html("");
				fileList_create(rs);
				$("#pageflag").val("Group");
				$("#pagecode").val(code);
			}
		});
	});
	// 搜索
	$(".butsearch").click(function(){
		find_materials_ajax(projectCode, $(".inpsearch").val(), function(data){
			if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				$(".strdiv .videodiv").html("");
				fileList_create(rs);
				$("#pageflag").val("");
				$("#pagecode").val("");
			}
		});

	});
	// 下载文件
	$(".strdiv .videodiv").on("click",".disnone ul li",function(){
		var code = $(this).parent().attr("class");
		var type = $(this).html().trim();
		post("/editoral_download_file", {MaterialCode: code, SourceType: type});
	});

	$("#addsou").click(function(){
		$(".url").val("/home/pillars/Videos");//设置默认路径
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({"height":height+"px","width":width+"px"});
		$(".formdiv").css({"left":(width/2)-200+"px","top":(height/2)-200+"px"});
		$(".outer").show(500);
		$(".formdiv").show(500);
	});
	//瀑布流加载
	$(window).scroll(function() {
		var scrollTop = $(this).scrollTop();
		var scrollHeight = $(document).height();
		var windowHeight = $(this).height();
		if(scrollTop + windowHeight == scrollHeight) {
			var pageflag = $("#pageflag").val();
			var pagecode = $("#pagecode").val();
			if(pageflag == "" || pagecode == "") {
				return;
			}

			var start = $(".strdiv .videodiv").children(".videostr").length;
			var end = start + 10;
			if(pageflag == "Library") {
				materials_ajax("1", projectCode, pagecode, start, end, function(data){
					if(data.FeedbackCode == 0) {
						var rs = JSON.parse(data.Data);
						fileList_create(rs);
					}
				});
			}else if(pageflag == "Group") {
				materials_ajax("2", projectCode, pagecode, start, end, function(data){
					if(data.FeedbackCode == 0) {
						var rs = JSON.parse(data.Data);
						fileList_create(rs);
					}
				});
			}
		}
	});

	$(".outer").click(function(){
		$(".outer").hide(500);
		$(".formdiv").find(":text").val("");
		$(".formdiv").hide(500);
		$(".formdiv2").find(":text").val("");
		$(".formdiv2").hide(500);
	});
	$(".submit").click(function(){
		var name = $(".names").val();
		var url = $(".url").val();
		var transcodingurl = $(".transcodingurl").val();
		var dpxurl = $(".dpxurl").val();
		var movurl = $(".movurl").val();
		if(name==""||url==""){
			return false;
		}
		//验证通过，ajax方式提交数据
		addLibrary_ajax(name, url, dpxurl, transcodingurl, movurl, projectCode, function(data) {
			if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				// 页面显示新添加的库
				$(".library").append("<li class='li1'><a href='javascript:void(0);' class='"+rs["LibraryCode"]+"'>"+rs["LibraryName"]+"</a></li>");
				library_click();
				// 库保存成功，查询素材列表
				materials_ajax("1", projectCode, rs["LibraryCode"], 0, 10, function(data){
					if(data.FeedbackCode == 0) {
						var mrs = JSON.parse(data.Data);
						$(".strdiv .videodiv").html("");
						fileList_create(mrs);
						$("#pageflag").val("Library");
						$("#pagecode").val(rs["LibraryCode"]);
					}
				});
			} else {
				// TODO 保存失败显示位置
				alert(data.FeedbackText);
			}
		});

		$(".outer").hide(500);
		$(".formdiv").hide(500);
	});
	$(".again").click(function(){
		$("#f1").find(":text").val("");
	});
	$(".files").click(function(){
		var rightdivabs = $(".rightdivabs").css("display");//当前是否悬浮
		var sourceid = $(this).siblings(".sourceid").val();//选中的id
		var absid = $(".float .sourceid").val();//右边悬浮的id

		if(rightdivabs=="block"&&sourceid==absid){
			$(".rightdivabs").css("display","none");
		}else{
			$(".rightdivabs").css("display","block");
			$(".float .sourceid").val(sourceid);
			//根据sourceid从后台查询数据
			$(".float .roughimg img").attr("src","");//缩略图
			$(".float .names").html("name"+sourceid);//素材名
			$(".float .size").html("size"+sourceid);//尺寸
			$(".float .long").html("long"+sourceid);//长度
			$(".float .speed").html("speed"+sourceid);//帧速率
			$(".float .metadata").html("metadata"+sourceid);//基本信息
		}
	});
	$(".addli").mouseenter(function(){
		$(this).parent().css("background","#ccc");
	});
	$(".addli").mouseleave(function(){
		$(this).parent().css("background","none");
	});
	$("#addmaterial").click(function(){
		$(".distreenode").append("<div class='TreeNode'><input type='text' class='nameinp'><span class='delbtn'>x</span></div>");
	});
	$(".distreenode").on("click",".delbtn",function(){
		$(this).parent().remove();
	});

	$(".distreenode").on("blur",".nameinp",function(){
		var inputstr = $(this).val();//获得文本框内容
		if(inputstr!=""){
			var nodelength = $(".distreenode").find(".TreeNode").length;
			var lihtml = "<a href='javascript:dos("+nodelength+")' target='"+nodelength+"'>"+inputstr+"</a><span class='addbtn'>+</span>";//拼接当前元素的内容
			$(this).parent().html(lihtml);
			var txt="<div class='clip' id='dd"+nodelength+"'>";
			$(".distreenode").append(txt);
		}
		//ltree("1",inputstr);//参数1目录层级，参数2文字内容
	});
	$(".distreenode").on("click",".addbtn",function(){
		var length = $(this).siblings("a").attr("target");
		var html = "<input type='text' class='nameinp2'><span class='delbtn2'>x</span>";
		$("#dd"+length).append(html);
		$("#dd"+length).css("display","block");
	});

	$(".distreenode").on("blur",".nameinp2",function(){
		var inputstr = $(this).val();//获得文本框内容
		var clipid = $(this).parent().attr("id");
		var cliplength = clipid.substring(2,clipid.length);
		var nodelength = $(".distreenode").find(".clip").length+1;
		if(inputstr!=""){
			$(this).siblings("span").remove();
			$(this).remove();/*
			var html = "<a href='javascript:dos("+nodelength+")' target='"+nodelength+"'>"+inputstr+"</a><span class='addbtn'>+</span>";//拼接当前元素的内容
			$(this).parent().html(lihtml);
			var txt="<div class='clip' id='dd"+nodelength+"'>";
			$(".distreenode").append(txt);*/
			var html = "<div class='treenode1'><a href='javascript:dos("+nodelength+")' target='"+nodelength+"'>"+inputstr+"</a><span class='addbtn'>+</span></div>";
			$("#dd"+cliplength).append(html);
			var txt="<div class='clip' id='dd"+nodelength+"'></div>";
			$("#dd"+cliplength).append(txt);
		}
		//ltree("1",inputstr);//参数1目录层级，参数2文字内容
	});
});
function dos(str)
{
	if($("#dd"+str).css("display")=="none")
	{
		$("#dd"+str).css("display","block");
	}else{
		$("#dd"+str).css("display","none");
	}
}
