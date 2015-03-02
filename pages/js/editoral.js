//url参数获取
var getUrlParam = function(name){
	var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
	var r = window.location.search.substr(1).match(reg);  //匹配目标参数
	if (r!=null) return unescape(r[2]); return null; //返回参数值
}
//列表数据
var materials_ajax = function(pc, lc, start, end){
	$.post("/editoral_library_materials",
		{
			ProjectCode: pc,
			LibraryCode: lc,
			Start: start,
			End: end
		},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				// TODO 列表数据已返回，组织页面显示
				$(".strdiv .videodiv").html("");
				for(var i=0;i<rs.length;i++){

					var html = '<span class="videostr"><input type="hidden" class="sourceid" value="'+rs[i]["MaterialCode"]+'"><input class="check" name="" type="checkbox" value=""><input class="play" type="button" value="回放"><div class="downdiv"><input class="downl" type="button" value="下载"><span class="disnone"><ul class="'+rs[i]["MaterialCode"]+'"><li>Source</li><li>DPX</li><li>JPG</li><li>Mov</li></ul></span></div><div class="files"><span class="name">'+rs[i]["MaterialName"]+'</span><span class="format">'+rs[i]["MaterialType"]+'</span><span class="long">'+rs[i]["Length"]+'</span></div></span>';
					$(".strdiv .videodiv").append(html);
				}
			}
        },
        "json"
    );
}
//Library数据
var filetype_ajax = function(pc){
	$.post("/editoral_library",
		{ProjectCode: pc},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				// TODO 列表数据已返回，组织页面显示
				for(var i=0; i<rs.length; i++){
					$("#materialGroup").html("<li><a href='javascript:void(0);'>"+rs[0]+"</a></li>");
				}
				setTimeout(function(){
					$("#materialGroup").on("click","a",function(){
						$(this).html()
					});
				},1000);
			}
        },
        "json"
    );
}
//素材组数据
var filetype_ajax = function(pc){
	$.post("/editoral_filetype",
		{ProjectCode: pc},
        function(data) {
            if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				// TODO 列表数据已返回，组织页面显示
				for(var i=0; i<rs.length; i++){
					$("#materialGroup").html("<li><a href='javascript:void(0);'>"+rs[0]+"</a></li>");
				}
				setTimeout(function(){
					$("#materialGroup").on("click","a",function(){
						$(this).html()
					});
				},1000);
			}
        },
        "json"
    );
}
//虚拟表单提交，下载使用
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
//下载验证
var downloadCheck_ajax = function(code, type, callback) {
    $.post("/editoral_download_file1", JSON.stringify({MaterialCode: code, SourceType: type}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//得到url地址中code参数
var projectCode = "";

$(function(){
	projectCode = getUrlParam("code");
	if (projectCode !== ""){
		filetype_ajax(projectCode);
	}
	// 搜索按钮，测试后台方法（下载文件）
	$(".strdiv .videodiv").on("click",".disnone ul li",function(){
		var code = $(this).parent().attr("class");
		code = "adsfas";
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
				// TODO 页面显示新添加的库，

				// 库保存成功，查询素材列表
				materials_ajax(projectCode, rs["LibraryCode"], 0, 20)
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
