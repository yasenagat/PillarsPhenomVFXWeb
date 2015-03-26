//得到url地址中code参数
var projectCode = "";
//url参数获取
var getUrlParam = function(name){
	var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
	var r = window.location.search.substr(1).match(reg);  //匹配目标参数
	if (r!=null) return unescape(r[2]); return null; //返回参数值
}
//左侧Library数据
var librarys_ajax = function(pc){
	$.post("/editoral_library",
		{ProjectCode: pc},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				if (rs == null) {
					return;
				}
				// 组织页面显示
				for(var i=0; i<rs.length; i++){
					$(".library").append("<li class='li1'><a href='javascript:void(0);' class='"+rs[0]["LibraryCode"]+"'>"+rs[0]["LibraryName"]+"</a></li>");
				}

				library_click();
			} else if (data.FeedbackCode == 404){
				window.location.href = "404.html";
			}
        },
        "json"
    );
}
var library_click = function(){
	$(".library").on("click","a",function(){
		$("#treeflag").val("0");
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
//左侧素材组数据
var filetypes_ajax = function(pc){
	$.post("/editoral_filetype",
		{ProjectCode: pc},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				if (rs == null) {
					return;
				}
				// 组织页面显示
				for(var i=0; i<rs.length; i++){
					$("#materialGroup").html("<li><a href='javascript:void(0);'>"+rs[0]+"</a></li>");
				}

				$("#materialGroup").on("click","a",function(){
					$("#treeflag").val("0");
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
			}
        },
        "json"
    );
}
//左侧列表点击,flag为判断点击的哪一个列表
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
	if (rs == null || rs.length == 0) {
		return;
	}

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
		var html = '<span class="videostr" id="span'+rs[i]["MaterialCode"]+'">';
		html += '<input type="hidden" class="sourceid" value="'+rs[i]["MaterialCode"]+'">';
		html += '<input class="check" name="" type="checkbox" value="'+rs[i]["MaterialCode"]+'">';
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
var post_download = function(URL, PARAMS) {
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
//左侧树形列表加载
var folders_ajax = function(pc){
	$.post("/editoral_folder",
		{ProjectCode: pc},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				d = new dTree('d');
				//遍历查询,格式为d.add(id,父id,'素材名称','','','1');
				for(var i=0; i<rs.length; i++){
					d.add(rs[i]["FolderCode"],rs[i]["FatherCode"],rs[i]["FolderName"],'','','1');
				}
				$(".tree").html(d + '');
				$(".dtree").children("div.dTreeNode").addClass("grouptit");
			}
        },
        "json"
    );
}
//左侧树形列表点击
var folders_click_ajax = function(pc, fi, callback){
	$.post("/editoral_folder_materials",
		{ProjectCode: pc, FolderId: fi},
        function(data) {
			callback(data);
        },
        "json"
    );
}

var dos = function(str) {
	if ($("#dd"+str).css("display")=="none") {
		$("#dd"+str).css("display","block");
	} else {
		$("#dd"+str).css("display","none");
	}
}
// 自定义分组点击事件
var li = function(string) {
	//点击行
	$("#treeflag").val("1");
	$(".dtree").find(".dTreeNode").css("background","none");
	$(".dtree").children("div.dTreeNode").css("background","#E3E3E3");
	$(".dtree").find("."+string).parent().css("background","#979797");
	// 查找该分组string的素材,返回素材列表
	folders_click_ajax(projectCode, string, function(data){
		if(data.FeedbackCode == 0){
			var rs = JSON.parse(data.Data);
			var flag = true;//是否包含素材,包含素材
			if(rs == null || rs.length == 0){//若列表length为0,flag=false
				$(".strdiv .videodiv").html("");
				flag = false;
			}
			if(flag && $("#dd"+string).css("display")=="none"){//如果包含素材,且当前目录为隐藏
				$(".strdiv .videodiv").html("");
				fileList_create(rs);
			}else if( flag && $("#dd"+string).css("display")=="block"){
				$("#dd"+string).css("display","none");
			}else if(!flag && $("#dd"+string).css("display")=="none"){
				$("#dd"+string).css("display","block");
			}else if(!flag && $("#dd"+string).css("display")=="block") {
				$("#dd"+string).css("display","none");
			}
		}
	});
}

$(function(){
	projectCode = getUrlParam("code");
	if (projectCode !== ""){
		librarys_ajax(projectCode);
		filetypes_ajax(projectCode);
		folders_ajax(projectCode);
		$(".postpage").attr("href","post.html?code="+projectCode)
	}

	// All列表点击
	$(".li1").on("click", "a", function(){
		$("#treeflag").val("0");
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
		post_download("/editoral_download_file", {MaterialCode: code, SourceType: type});
	});
	// 批量下载文件
	$("#butdownloads").click(function(){
		//获得当前选中复选框的id
		var strs = new Array();
		$('input[class="check"]:checked').each(function(){
			strs.push($(this).val());
		});
		if(strs.length == 0){
			alert("请先勾选要下载的素材!");
			return;
		}
		var type = "Source";

		for(var i=0; i<strs.length; i++){
			window.open("/editoral_download_file?MaterialCode="+strs[i]+"&SourceType="+type);
		}

		/*
		function sub(i){
			if(i == strs.length) {
				return;
			}
			var temp = document.createElement("form");
		    temp.action = "/editoral_download_file";
		    temp.method = "get";
		    temp.style.display = "none";
		    var opt = document.createElement("textarea");
		    opt.name = "MaterialCode";
		    opt.value = strs[i];
		    temp.appendChild(opt);
			var opt2 = document.createElement("textarea");
			opt2.name = "SourceType";
		    opt2.value = type;
		    temp.appendChild(opt2);
		    document.body.appendChild(temp);
			i += 1;
			alert("第" + i + "个文件可以下载了!");
		    temp.submit();
			setTimeout(sub(i),10);
		}
		sub(0);
		*/
		//downloadCheck_ajax("/editoral_download_file", {MaterialCode: strs[i], SourceType: type});
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
			var rs = [];
			if(pageflag == "Library") {
				materials_ajax("1", projectCode, pagecode, start, end, function(data){
					if(data.FeedbackCode == 0) {
						rs = JSON.parse(data.Data);
						fileList_create(rs);
					}
				});
			}else if(pageflag == "Group") {
				materials_ajax("2", projectCode, pagecode, start, end, function(data){
					if(data.FeedbackCode == 0) {
						rs = JSON.parse(data.Data);
						fileList_create(rs);
					}
				});
			}
		}
	});

	$(".strdiv .videodiv").on("click",".files",function(){
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
						metasHtml += '<div><span>' + s + '</span>';
						metasHtml += " : ";
						metasHtml += '<span>' + metaData[s] + '</span></div>';
					}
					$(".float .metadata").html(metasHtml);//基本信息
				}
			});
		}
	});

	$("#addsou").click(function(){
		$(".url").val("/home/pillars/Videos");// TODO delete 设置默认路径
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({"height":height+"px","width":width+"px"});
		$(".formdiv").css({"left":(width/2)-200+"px","top":(height/2)-200+"px"});
		$(".outer").show(500);
		$(".formdiv").show(500);
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
				filetypes_ajax(projectCode);
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
