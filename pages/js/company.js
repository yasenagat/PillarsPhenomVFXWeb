//得到url地址中code参数
var code = "";
//url参数获取
var getUrlParam = function(name){
var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
var r = window.location.search.substr(1).match(reg); //匹配目标参数
if (r!=null) return unescape(r[2]); return null; //返回参数值
}

$(function(){
	//获得地址栏该外包公司code
	code = getUrlParam("code");
	var startProject = function(){
		//根据初始化左侧菜单项目列表
		var html = "";//拼接字符串
		for(var i = 0; i<5; i++){
			var code = i;
			html += "<li class='li2' name='"+code+"'><a href='javascript:void(0);'>项目"+i+"</li>";
		}
		$(".groupul").html(html);
	}
	startProject();
	
	var viewList = function(code){
		var length = 0;
		if(code=="all"){
			//TODO 取所有数据
			length = 10;
		}else{
			//TODO 取当前code数据
			length = 5;
		}
		//显示所有镜头
		var html = '';
		for(var i = 0; i<length; i++){
			var code = i;
			var imgsrc = "http://imgserv.jd-app.com/?size=200x150";
			var names = "素材名"+i;
			var size = "W"+"x"+"H";
			var speed = i+"fps";
			//根据镜头code 查询前几条制作需求
			var needhtml = "";
			for(var ii = 0; ii<3; ii ++){
				var need = "一二三四五六七八九十一"+ii;
				if(need.length>=15)
					needhtml += need.substring(0,15)+"...";
				else
					needhtml += need;
					needhtml +="<br>";
			}
			html += '<div class="post"><input type="hidden" class="code" value="'+code+'"><div class="screenshot"><img src="'+imgsrc+'"></div><div class="summary"><table width="100%" height="150" border="0" cellspacing="0" cellpadding="0"><tr height="37"><td>'+names+'</td></tr><tr height="37"><td>'+size+'</td></tr><tr height="37"><td>'+speed+'</td></tr><tr height="37"><td></td></tr></table></div><div class="tag"><h2>制作需求</h2><div class="stage">'+needhtml+'</div></div><div class="view"></div><span href="javascript:void(0);" class="spanfile"><input type="button" class="sample" value="小样"><input type="button" class="product" value="成品"><input type="file" class="updfile" onchange="updatafile(this)"></span></div>';
		}
		$(".cont").find(".rightdiv").html(html);
	}
	//初始化右侧需求列表
	var startTask = function(){
		viewList("all");
	}
	startTask();
	$(".leftcontent").on("click","a",function(){
		//得到当前列表code
		var code = $(this).parent().attr("name");
		if(code=="all"){
			//显示所有镜头
			viewList("all");
		}else{
			//显示该列表镜头
			viewList(code);
		}
	});
	$(".rightdiv").on("click",".view",function(){
		if($(".float").children(".rightdivabs").css("display")=="block"){
			$(".float").children(".rightdivabs").css("display","none");
			return;
		}else{
			$(".float").children(".rightdivabs").css("display","block");
		}
		//获得当前镜头code
		var code = $(this).siblings(".code").val();
		//初始化页面标签
		$(".withtit").find("li").css("background","#FFF");
		$(".withtit").find(".basic").css("background","#CCC");
		$(".metadata").children("div").hide();
		//赋值给弹出div
		$(".float").children(".sourceid").val(code);
		$(".metadata").children(".basicinfo").css("display","block");
		//根据code 取该code镜头信息
		
		//赋值给页面
		var names = "镜头名";//镜头名
		var size = "W"+"x"+"H";//尺寸
		var speed = "帧速率";//帧速率
		var start = "始码";//始码
		var end = "止码";//止码
		$(".ovhi").find(".names").html(names);
		$(".ovhi").find(".size").html(size);
		$(".ovhi").find(".speed").html(speed);
		$(".ovhi").find(".start").html(start);
		$(".ovhi").find(".end").html(end);
		basic(code);
	});
	//点击基本信息 制作需求 素材 note 按钮
	$(".withtit li").click(function(){
		var menu = $(this).attr("class");
		$(this).siblings("li").css("background","#FFF");
		$(this).css("background","#ccc");
		var thiscode = $(".float").children(".sourceid").val();
		if(menu=="basic"){
			basic(thiscode);
		}else if(menu=="need"){
			need(thiscode);
		}else if(menu=="material"){
			material(thiscode);
		}else{
			note(thiscode);
		}
		//显示点击的列表
		$(".metadata").children("div").hide();
		$(".metadata").children("."+menu+"info").show();
	});
	//基本信息
	var basic = function(thiscode){
		//根据code 取该code镜头信息
		
		//赋值给页面
		var names = "镜头名";//镜头名
		var size = "W"+"x"+"H";//尺寸
		var speed = "帧速率";//帧速率
		var start = "始码";//始码
		var end = "止码";//止码
		var bewrite = "描述";//描述
		$(".basicinfo").find(".names").html(names);
		$(".basicinfo").find(".size").html(size);
		$(".basicinfo").find(".speed").html(speed);
		$(".basicinfo").find(".start").html(start);
		$(".basicinfo").find(".end").html(end);
		$(".basicinfo").find(".bewrite").html(bewrite);
	}
	//制作需求
	var need = function(thiscode){
		var html = "";
		for(var i = 0; i<5;i++){
			var imgsrc = "http://imgserv.jd-app.com/?size=60x60";
			var content = "neirong"+i;
			html += '<div class="piece">   <input type="hidden" value="1" class="makecode" />   <div class="number">'+(i+1)+'</div>   <div class="news">    <div class="imgradius">     <img width="60" height="60" src="'+imgsrc+'" />    </div>    <div class="stage">'+content+'</div>   </div>  </div>';
		}
		$(".needinfo").find(".make").html(html);
	}
	//素材
	var material = function(thiscode){
		//TODO 根据镜头code 查询该镜头的参考素材列表
		var html = '<tr><td width="25%">素材名</td><td width="25%">素材格式</td><td colspan="2">描述</td></tr>';
		for(var i = 0; i<5; i++){
			var code = i+1;//素材code
			var names = "素材名"+i;//素材名
			var layout = "素材格式"+i;//素材格式
			var depict = "描述"+i;//描述
			html += '<tr><td>'+names+'</td><td>'+layout+'</td><td>'+depict+'</td><td width="5%" class="symbol2">+<div class="bolpoab2"><ul name="'+code+'"><li class="download">下载</li></ul></div></td></tr>';
		}
		$(".edittable").html(html);
	}
	//Note
	var note = function(thiscode){
		//TODO 根据该镜头code 获得该code的消息列表
		var html = "";
		for(i=0;i<5;i++){
			var name = "";
			var content =""+i;
			var lor = "";
			if("1" != "1")
				lor = "l";
			else
				lor = "r";
			html += "<div class='newinfo"+lor+"'><div class='username'>"+name+"</div><div class='buddy'>"+content+"</div></div>";
		}
		var input = document.getElementById("fileimg");
		input.addEventListener('change', readImg, false);
		$(".noteinfo").find(".chat").html(html);
		
	}
	//下载参考素材
	$(".edittable").on("click",".download",function(){
		var code = $(this).parent().attr("name");
		alert(code);
	});
	
	//发送消息
	$(".but").click(function(){
		//获得图片消息
		var imgsrc = $("#imgs").attr("src");
		//获得文本框消息
		var txt = $(".textarea").val();
		//获得镜头code
		var code = $(".sourceid").val();
		if(imgsrc=="#"&&txt==""){
			alert("请输入消息内容或图片");
			return;
		}
		var html = "";
		if(imgsrc=="#"&&txt!=""){
			html = "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>"+txt+"</div></div>";
		}else if(imgsrc!="#"&&txt!=""){
			var srchtml = "<img src='"+imgsrc+"'>";
			html = "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>"+srchtml+""+txt+"</div></div>";
		}else if(imgsrc!="#"&&txt==""){
			var srchtml = "<img src='"+imgsrc+"'>";
			html = "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>"+srchtml+"</div></div>";
		}
		
		//TODO 把消息对应该镜头code记录到数据库 图片src为imgsrc 文字消息为txt
		
		//添加到页面
		$(".chat").append(html);
		//初始化文本框
		$(".textarea").val("");
		$("#imgs").attr("src","#");
		//初始化下拉
		$(".chat").scrollTop(200000);
	});
	$(".rightdiv").on("click",".updfile",function(){
		$("#cameraCode").val($(this).parent().siblings(".code").val());
	});
	//上传小样窗口显示
	$(".rightdiv").on("click",".sample",function(){
		//得到当前镜头code
		var code = $(this).parent().siblings(".code").val();
		//显示弹出div
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({
			"height": height + "px",
			"width": width + "px"
		});
		$(".formdiv1").css({
			"left": (width / 2) - 200 + "px",
			"top": (height / 2) - 100 + "px"
		});
		$(".outer").show(500);
		$(".formdiv1").show(500);
		$(".formdiv1").find(".filesample").attr("disabled","disabled");
		$(".formdiv1").find(".code").val(code);
		$(".formdiv1").find(".description").val("");
		$(".formdiv1").find(".filesample").val("");
	});
	$(".description").blur(function(){
		if($(".formdiv1").find(".description").val()!="")
			$(".formdiv1").find(".filesample").removeAttr("disabled");
		else
			$(".formdiv1").find(".filesample").attr("disabled","disabled");
	});
	$(".outer").click(function(){
		$(".outer").hide();
		$(".formdiv1").hide();
	});
});
function readImg() {
	var file = this.files[0];
	//判断类型如果不是图片就返回 去掉就可以上传任意文件
	if (!/image\/\w+/.test(file.type)) {
		alert("请确保文件为图像类型");
		return false;
	}
	var reader = new FileReader();
	reader.readAsDataURL(file);
	reader.onload = function(e) {
		$("#imgs").attr("src",this.result);
	}
}
function updatafile(file){
	var code = $("#cameraCode").val();
	//TODO 当前镜头code
	alert(code);
}
function upsample(file){
	var code = $(".formdiv1").find(".code").val();
	var description = $(".formdiv1").find(".description").val();
	//TODO 当前镜头code
	alert(code);
	$(".outer").click();
}