
// ---------------------- start ----------------------
//得到url地址中code参数
var projectCode = "";
//url参数获取
var getUrlParam = function(name){
	var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
	var r = window.location.search.substr(1).match(reg);  //匹配目标参数
	if (r!=null) return unescape(r[2]); return null; //返回参数值
}
//上传edl文件
function uploadEdl(selectFile) {
	var filename = selectFile.value;
	var mime = filename.toLowerCase().substr(filename.lastIndexOf("."));
    if (mime != ".edl") {
        alert("请选择edl格式的文件上传");
        selectFile.outerHTML = selectFile.outerHTML;
        return false;
    }
	var form = document.getElementById("edl-form");
	var edlfile = document.getElementById("edlfile");
	var files = edlfile.files;
	var formData = new FormData();
	for(var i=0; i<files.length; i++){
		var file = files[i];
		formData.append("files", file, file.name);
	}
	formData.append("ProjectCode", projectCode);
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "post_upload_edl", true);
	xhr.onload = function(){
		if(xhr.status === 200){
			var rs = JSON.parse(xhr.responseText);
			if(rs.FeedbackCode == 0) {
				var shotList = JSON.parse(rs.Data);
				alert(shotList.length+"");
				// TODO 创建页面内容
				for (var i=0; i<shotList.length; i++){

				}
			}else{
				alert(rs.FeedbackText);
			}
		}else{
			alert("upload error");
		}
	}
	xhr.send(formData);
}
//自定义镜头组列表加载
var folders_ajax = function(pc){
	$.post("/post_shot_folder",
		JSON.stringify({ProjectCode: pc}),
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
//查询"Load EDL"的镜头
var shots_ajax = function(pc, callback){
	$.post("/post_shot_list",
		JSON.stringify({ProjectCode: pc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}

// ---------------------- End -----------------------

// JavaScript Document
$(function(){
	// ---------------------- start ----------------------
	projectCode = getUrlParam("code");
	if (projectCode !== ""){
		$(".editoralpage").attr("href","editoral.html?code=" + projectCode);
		folders_ajax(projectCode);
	}else{
		return;
	}

	shots_ajax(projectCode, function(data){
		if(data.FeedbackCode == 0) {
			var rs = JSON.parse(data.Data);
			if(rs == null || rs.length === 0){
				alert("NO Data!");
				// TODO 没有数据,"Load EDL"可以点击
				return;
			}
			alert(rs.length+"");
			//rs[0]["ShotCode"]
			//TODO 有数据,"Load EDL"禁止点击,创建页面镜头视图列表
		}
	});

	// ---------------------- End ----------------------


	$(".dtree").children("div.dTreeNode").css("background","#e3e3e3");

	//初始化右侧窗口 隐藏
	$(".metadata").children(".basicinfo").css("display","block");
	$(".videodiv").on("click",".view",function(){
		var rightabs = $(".rightdivabs").css("display");
		if(rightabs=="block"){
			$(".rightdivabs").css("display","none");
		}else{
			$(".rightdivabs").css("display","block");
			//获得当前镜头的id
			var thiscode = $(this).siblings("#code").attr("value");
			$(".float").find(".sourceid").val(thiscode);
			//TODO 根据当前选中镜头的id thiscode，查询该镜头id的信息:镜头名 尺寸 帧速率 始码 止码 描述
			var names = "123456789123456789";//镜头名
			var size = "";//尺寸
			var speed = "";//帧速率
			var start = "";//始码
			var end = "";//止码
			//处理长字符串
			if(names.length>10)
				names = names.substring(0,9)+"...";
			if(size.length>10)
				size = size.substring(0,9);
			if(speed.length>10)
				speed = speed.substring(0,9);
			if(start.length>10)
				start = start.substring(0,9);
			if(end.length>10)
				end.substring(0,9);
			//初始化右侧窗口 隐藏
			$(".metadata").children("div").css("display","none");
			$(".metadata").find(".basicinfo").css("display","block");
			$(".withtit").children("li").css("background","#FFF");
			$(".withtit").find(".basic").css("background","#ccc");
			$(".tab1").find(".names").html(names);//镜头名
			$(".tab1").find(".size").html(size);//尺寸
			$(".tab1").find(".speed").html(speed);//帧速率
			$(".tab1").find(".start").html(start);//始码
			$(".tab1").find(".end").html(end);//止码
			basic(thiscode);
		}
	});
	$(".dels").click(function(){
		//获得当前选中复选框的id
		var str = "";
		$('input[class="check"]:checked').each(function() {
			str += $(this).val() + ",";
		});
		if (str == "") {
			alert("请选中镜头复选框再删除镜头");
			return;
		}
		//添加镜头id到该镜头组id，镜头组id是str，镜头id是thiscode
		alert("要删除的镜头code：" + str);
	});
	//便捷的修改素材名
	$(".videoimg").on("click",".names",function(){
		//判断该h2标签是否已经包含input标签 若有 说明正在编辑
		if($(this).children("input").length>0){
			return;
		}
		//获得当前镜头的镜头名
		var names = $(this).html().trim();

		$(this).html("<input type='text' class='updlab'>");
		$(this).children("input").focus();//获得焦点
		$(this).children("input").val(names);
	})
	$(".videoimg").on("blur",".updlab",function(){//释放焦点时 修改数据库中该code的镜头名
		//获得当前镜头的code
		var code = $(this).parent().siblings("#code").val();
		//获得修改后文本框的文字
		var names = $(this).val();
		//TODO 往数据库中根据该code修改该镜头的镜头名称

		//修改页面显示内容
		$(this).parent().html(names);
	});
	$(".withtit li").click(function(){
		//获得点击元素class
		var names = $(this).attr("class");
		//初始化点击按钮
		$(".withtit").children("li").css("background","none");
		$(this).css("background","#ccc");
		//初始化右侧窗口 隐藏
		$(".metadata").children("div").css("display","none");
		$("."+names+"info").show();
		//获取当前选中镜头的id
		var code = $(".float").find(".sourceid").val();
		if(names == "basic"){
			basic(code);
		}else if(names == "need"){
			need(code);
		}else if(names == "material"){
			material(code);
		}else if(names == "note"){
			node(code);
		}else if(names == "edition"){
			edition(code);
		}
	});
	var basic = function(code){//详细信息
		$(".tab2").find(".names").html("a");//镜头名
		$(".tab2").find(".size").html("a");//尺寸
		$(".tab2").find(".speed").html("a");//帧速率
		$(".tab2").find(".start").html("a");//始码
		$(".tab2").find(".end").html("a");//止码
		$(".tab2").find(".bewrite").html("a");//描述
	}
	var need = function(code){//制作需求
		//TODO 根据镜头id查询制作需求list,并遍历
		var html = "";//要拼接的html

		for(i=0;i<2;i++){
			var code = (i*1)+1;
			var imgurl = "";
			var mation = "ab";
			var length = (i*1)+1;
			html += '<div class="piece"><input type="hidden" class="makecode" value="'+code+'"><div class="number">'+length+'<div class="del">X</div></div><div class="news"><div class="imgradius"><img src="'+imgurl+'" width="60" height="60"></div><div class="stage">'+mation+'</div><input type="button" value="编辑" class="edit"></div>                        </div>';
		}
		$(".needinfo .make").html(html);

	}
	var material = function(code){
		//TODO 根据该镜头code查询该镜头的参考素材列表
		var html = "<tr><td width='25%'>素材名</td><td width='25%'>素材格式</td><td colspan='2'>描述</td></tr>";
		for(i=0;i<2;i++){
			var code = i+1;//素材code
			var names = "素材名"+i;//素材名
			var layout = "素材格式"+i;//素材格式
			var depict = "描述"+i;//描述
			html += "<tr><td>"+names+"</td><td>"+layout+"</td><td>"+depict+"</td><td class='symbol2'  width='5%'>+<div class='bolpoab2'><ul name='"+code+"'><li class='download'>下载</li><li class='delete'>删除</li></ul></div></td></tr>";
		}
		$(".edittable").html(html);
	}
	//编辑详细信息
	$("#submits").click(function(){
		//获得该按钮上的文字 判断该编辑还是直接保存
		var butval = $(this).val();
		if(butval == "编辑"){
			//获得当前镜头ID和字段内容 并变成可编辑的文本框
			var names = $(".basicinfo").find(".names").html().trim();//镜头名
			var size = $(".basicinfo").find(".size").html().trim();//尺寸
			var speed = $(".basicinfo").find(".speed").html().trim();//帧速率
			var start = $(".basicinfo").find(".start").html().trim();//始码
			var end = $(".basicinfo").find(".end").html().trim();//止码
			var bewrite = $(".basicinfo").find(".bewrite").html().trim();//描述
			$(".basicinfo").find(".names").html("<input type='text' id='namesinp' value='"+names+"'>");
			$(".basicinfo").find(".size").html("<input type='text' id='sizeinp' value='"+size+"'>");
			$(".basicinfo").find(".speed").html("<input type='text' id='speedinp' value='"+speed+"'>");
			$(".basicinfo").find(".start").html("<input type='text' id='startinp' value='"+start+"'>");
			$(".basicinfo").find(".end").html("<input type='text' id='endinp' value='"+end+"'>");
			$(".basicinfo").find(".bewrite").html("<input type='text' id='bewriteinp' value='"+bewrite+"'>");
			$(this).val("保存");
		}else{//保存
			var names = $(".basicinfo").find("#namesinp").val();//镜头名
			var size = $(".basicinfo").find("#sizeinp").val();//尺寸
			var speed = $(".basicinfo").find("#speedinp").val();//帧速率
			var start = $(".basicinfo").find("#startinp").val();//始码
			var end = $(".basicinfo").find("#endinp").val();//止码
			var bewrite = $(".basicinfo").find("#bewriteinp").val();//描述
			//后台保存方法

			//文本框改成字符串
			$(".basicinfo").find(".names").html(names);
			$(".basicinfo").find(".size").html(size);
			$(".basicinfo").find(".speed").html(speed);
			$(".basicinfo").find(".start").html(start);
			$(".basicinfo").find(".end").html(end);
			$(".basicinfo").find(".bewrite").html(bewrite);
			$(this).val("编辑");
		}
	});
	//详细信息页面 编辑制作需求
	$(".make").on("click",".edit",function(){
		var value = $(this).val();
		if(value=="编辑"){
			if($(".needinfo").find("#demand").length>0){//判断是否已经打开判断的窗口
				alert("您正在编辑");
				$("#demand").focus();
				return;
			}
			var makeval = $(this).siblings(".stage").html().trim();
			$(this).siblings(".stage").html("<textarea name='' cols='' rows='' id='demand'>"+makeval+"</textarea>");
			$(this).val("保存");
		}else{
			var makeval = $(this).siblings(".stage").children("#demand").val();//当前制作需求输入框
			var makecode = $(this).parent().siblings(".makecode").val();//当前制作需求id
			//TODO 当前需求id makecode 当前需求字符串 makeval 更新到数据库
			$(this).siblings(".stage").html(makeval);
			$(this).val("编辑");
		}
	});
	//制作需求的删除
	$(".make").on("click",".del",function(){
		//获得该制作需求code
		var code =$(this).parent().siblings(".makecode").val();
		//TODO 从数据库中删除该code

		$(this).parents(".piece").remove();
	});
	$(".addmark").click(function(){
		if($(".needinfo").find("#demand").length>0){//判断是否已经打开判断的窗口
			alert("您正在添加");
			$("#demand").focus();
			return;
		}
		var length = ($(".make").children(".piece").length*1)+1;
		var html = '<div class="piece"><input type="hidden" class="makecode" value="1"><div class="number">'+length+'<div class="del">X</div></div><div class="news"><div class="imgradius"><img src="#" width="60" height="60"></div><div class="stage"><textarea id="demand" rows="" cols="" name=""></textarea></div><input type="button" value="保存" class="edit"></div>                        </div>';
		$(".needinfo .make").append(html);
		$("#demand").focus();
	});
	//参考素材的下载
	$(".edittable").on("click",".download",function(){
		//获取当前参考素材的code
		var code = $(this).parent().attr("name");
		//TODO 数据库根据该code获取相应的下载地址

	});
	//参考素材的删除
	$(".edittable").on("click",".delete",function(){
		//获取当前参考素材的code
		var code = $(this).parent().attr("name");
		//TODO 数据库删除该code

		//删除该页面对应的数据
		$(this).parents("tr").remove();
	});
	//note消息读取
	var node = function(code){
		//TODO 根据该镜头code 获得该code的消息列表
		var html = "";
		for(i=0;i<10;i++){
			var content =""+i;
			html += "<div class='newinfo'>"+content+"</div>";
		}
		$(".noteinfo").find(".chat").html(html);
	}
	//发送消息
	$(".but").click(function(){
		//获得文本框消息
		var txt = $(".textarea").val();
		//获得镜头code
		var code = $(".sourceid").val();
		//TODO 把消息对应该镜头code记录到数据库

		var html = '<div class="newinfo">'+txt+'</div>';
		//添加到页面
		$(".chat").append(html);
		//初始化文本框
		$(".textarea").val("");
		//初始化下拉
		$(".chat").scrollTop(200000);
	});
	//根据镜头code 查询版本列表信息
	var edition = function(code){
		//TODO 根据code 获得版本信息
		var html = "";
		for(i=0;i<3;i++){
			var code = i;//当前版本code
			var num = "版本号";//版本号
			var img = "缩略图";//缩略图url
			html += '<div class="tag" name="'+code+'"><div class="num">'+num+'</div><div class="frame"><div class="imgs"><img src="'+img+'" alt="缩略图"></div><div class="menu_tab">+<div class="bolpoabs"><ul><li class="viewhue">查看小样</li><li class="down">下载成品</li></ul></div></div></div></div>';
		}
		$(".editioninfo").html(html);
	}
	//查看版本小样
	$(".editioninfo").on("click",".viewhue",function(){
		//获得该版本的code
		var code = $(this).parents(".tag").attr("name");
		alert(code);
	});
	//下载版本成品
	$(".editioninfo").on("click",".down",function(){
		//获得该版本的code
		var code = $(this).parents(".tag").attr("name");
		alert(code);
	});
	$(".updfile").click(function(){
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({"height":height+"px","width":width+"px"});
		$(".formdiv1").css({"left":(width/2)-200+"px","top":(height/2)-200+"px"});
		$(".outer").show(500);
		$(".formdiv1").show(500);
	});
	$(".outer").click(function(){
		$(".outer").hide(500);
		$(".formdiv1").find(":text").val("");
		$(".formdiv1").hide(500);
		$(".formdiv2").find(":text").val("");
		$(".formdiv2").hide(500);
	});
});
function upload(fnUpload){
	var filename = fnUpload.value;
	var mime = filename.toLowerCase().substr(filename.lastIndexOf("."));
	if(mime!=".edl")
	{
		alert("请选择edl格式的文件上传");
		fnUpload.outerHTML=fnUpload.outerHTML;
		return false;
	}else{
		$("#f1").submit();
	}
}
