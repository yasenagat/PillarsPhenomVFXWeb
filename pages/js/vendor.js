//得到url地址中code参数
var code = "";
//url参数获取
var getUrlParam = function(name){
var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
var r = window.location.search.substr(1).match(reg); //匹配目标参数
if (r!=null) return unescape(r[2]); return null; //返回参数值
}

//项目列表加载
var project_list_ajax = function(sc, callback){
	$.post("/vendor_project_list", "",
        function(data) {
            callback(data);
        },
        "json"
    );
}

//项目列表点击查询镜头
var project_shots_ajax = function(vc, callback){
	$.post("/vendor_project_shots",
		JSON.stringify({VendorCode: vc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}

// 项目列表
var viewList = function(code) {
	project_shots_ajax(code, function(data){
		if (data.FeedbackCode == 0) {
			var rs = JSON.parse(data.Data);
			createPageList(rs);
		}
	});
}
//创建页面的列表(镜头+需求)
var createPageList = function (rs) {
	var html = '';
	if (rs != null && rs.length >0) {
		for(var i=0; i<rs.length; i++){
			//列表镜头的信息
			var shot = rs[i]["Shot"];
			var code = shot["ShotCode"];
			var pcode = shot["ProjectCode"];
			var imgsrc = shot["Picture"];
			var names = shot["ShotName"];
			var size = shot["Width"] + " X " + shot["Height"];
			var speed = shot["ShotFps"] + " fps";

			//列表制作需求的信息
			var demands = rs[i]["Demands"];
			var needhtml = "";
			if (demands != null && demands.length > 0) {
				for(var j=0; j<demands.length; j++){
					var need = demands[i]["DemandDetail"];
					if (need.length >= 15) {
						needhtml += need.substring(0,15)+"...";
					}else{
						needhtml += need;
					}
					needhtml += "<br>";
				}
			}

			html += '<div class="post"><input type="hidden" class="code" value="'+code+'"><input type="hidden" class="pcode" value="'+pcode+'"><div class="screenshot"><img src="'+imgsrc+'"></div><div class="summary"><table width="100%" height="150" border="0" cellspacing="0" cellpadding="0"><tr height="37"><td>'+names+'</td></tr><tr height="37"><td>'+size+'</td></tr><tr height="37"><td>'+speed+'</td></tr><tr height="37"><td></td></tr></table></div><div class="tag"><h2>制作需求</h2><div class="stage">'+needhtml+'</div></div><div class="view"></div><span href="javascript:void(0);" class="spanfile"><input type="button" class="sample" value="小样"><input type="button" class="product" value="成品"><input type="file" class="updfile" id="pro'+code+'" onchange="updatafile(this)"></span></div>';
		}
	}

	$(".cont").find(".rightdiv").html(html);
}

//查询单镜头基本信息
var shot_info_ajax = function(sc, callback){
	$.post("/post_shot_que",
		JSON.stringify({ShotCode: sc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头制作需求list
var shot_demandlist_ajax = function(sc, callback){
	$.post("/post_shot_demand_que",
		JSON.stringify({ShotCode: sc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头参考素材列表查询
var shot_materialque_ajax = function(sc, callback){
	$.post("/post_shot_material_que",
		JSON.stringify({ShotCode: sc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}

//镜头NOTE新增
var shot_noteadd_ajax = function(pc, sc, nd, p, callback){
	$.post("/post_shot_note_add",
		JSON.stringify({ProjectCode: pc, ShotCode: sc, NoteDetail: nd, Picture: p}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头NOTE列表查询
var shot_noteque_ajax = function(sc, callback){
	$.post("/post_shot_note_que",
		JSON.stringify({ShotCode: sc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
// 上传成品
function updatafile(selectFile){
	var code = $("#cameraCode").val();
	var pcode = $("#pcodes").val();

	var filename = selectFile.value;
	var type = filename.substr(filename.lastIndexOf(".")+1);//文件格式
	var edlfile = document.getElementById("pro"+code);
	var files = edlfile.files;
	var formData = new FormData();
	for(var i=0; i<files.length; i++){
		var file = files[i];
		formData.append("files", file, file.name);
	}
	formData.append("ShotCode", code);
	formData.append("ProjectCode", pcode);
	formData.append("ProductType", type);
	//formData.append("ProductDetail", detail);产品暂无描述
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "vendor_product_upload", true);
	xhr.onload = function(){
		if(xhr.status === 200){
			var rs = JSON.parse(xhr.responseText);
			if(rs.FeedbackCode == 0) {
				$(".outer").click();
			}else if (rs.FeedbackCode == 202) {
				alert("要上传的文件服务器已经存在!");
			}else {
				alert(rs.FeedbackText);
			}
		}else{
			alert("upload error");
		}
	}
	xhr.send(formData);
}
//上传小样
function upsample(selectFile) {
	var code = $(".formdiv1").find(".code").val();
	var pcode = $(".formdiv1").find(".pcode").val();
	var detail = $(".formdiv1").find(".description").val();// 描述

	var filename = selectFile.value;
	var type = filename.substr(filename.lastIndexOf(".")+1);//文件格式
	var edlfile = document.getElementById("demofile");
	var files = edlfile.files;
	var formData = new FormData();
	for(var i=0; i<files.length; i++){
		var file = files[i];
		formData.append("files", file, file.name);
	}
	formData.append("ProjectCode", pcode);
	formData.append("ShotCode", code);
	formData.append("DemoType", type);
	formData.append("DemoDetail", detail);
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "vendor_demo_upload", true);
	xhr.onload = function(){
		if(xhr.status === 200){
			var rs = JSON.parse(xhr.responseText);
			if(rs.FeedbackCode == 0) {
				$(".outer").click();
			}else if (rs.FeedbackCode == 202) {
				alert("要上传的文件服务器已经存在!");
			}else {
				alert(rs.FeedbackText);
			}
		}else{
			alert("upload error");
		}
	}
	xhr.send(formData);
}

$(function(){
	//获得地址栏该外包公司code
	code = getUrlParam("code");

	// 加载左侧项目列表
	project_list_ajax("", function(data){
		if (data.FeedbackCode == 0){
			var rs = JSON.parse(data.Data);
			var html = "";//拼接字符串
			if(rs != null && rs.length > 0){
				for(var i = 0; i<rs.length; i++){
					html += "<li class='li2' name='"+rs[i]["VendorCode"]+"'><a href='javascript:void(0);'>项目"+rs[i]["VendorName"]+"</li>";
				}
			}

			$(".groupul").html(html);
		}
	});

	//初始化右侧需求列表
	viewList("all");

	//列表点击
	$(".leftcontent").on("click","a",function(){
		//得到当前列表code
		var code = $(this).parent().attr("name");
		if(code=="all"){
			// 显示所有镜头
			viewList("all");
		}else{
			// 显示该列表镜头
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
		shot_info_ajax(thiscode, function(data){
			if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var names = rs["ShotName"];//镜头名
				var size = rs["Width"]+" X "+rs["Height"];//尺寸
				var speed = rs["ShotFps"];//帧速率
				var start = rs["StartTime"];//始码
				var end = rs["EndTime"];//止码
				var bewrite = rs["ShotDetail"];//描述
				//赋值给页面
				$(".basicinfo").find(".names").html(names);
				$(".basicinfo").find(".size").html(size);
				$(".basicinfo").find(".speed").html(speed);
				$(".basicinfo").find(".start").html(start);
				$(".basicinfo").find(".end").html(end);
				$(".basicinfo").find(".bewrite").html(bewrite);
				//赋值给页面
				$(".ovhi").find(".names").html(names);
				$(".ovhi").find(".size").html(size);
				$(".ovhi").find(".speed").html(speed);
				$(".ovhi").find(".start").html(start);
				$(".ovhi").find(".end").html(end);
			}
		});
	}
	//制作需求
	var need = function(code){
		shot_demandlist_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var html = "";//要拼接的html
				if(rs != null && rs.length > 0){
					for(var i=0; i<rs.length; i++){
						var imgsrc = rs[i]["Picture"];
						var content = rs[i]["DemandDetail"];
						html += '<div class="piece">   <input type="hidden" value="1" class="makecode"/>   <div class="number">'+(i+1)+'</div>   <div class="news">    <div class="imgradius">     <img width="60" height="60" src="'+imgsrc+'"/>    </div>    <div class="stage">'+content+'</div>   </div>  </div>';
					}
				}

				$(".needinfo").find(".make").html(html);
			}
		});
	}
	//素材
	var material = function(code){
		// 根据镜头code 查询该镜头的参考素材列表
		shot_materialque_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var html = "<tr><td width='25%'>素材名</td><td width='25%'>素材格式</td><td colspan='2'>描述</td></tr>";
				if (rs != null && rs.length > 0){
					for(var i=0; i<rs.length; i++){
						var code = rs[i]["MaterialCode"];//素材code
						var names = rs[i]["MaterialName"];//素材名
						var layout = rs[i]["MaterialType"];//素材格式
						var depict = rs[i]["MaterialDetail"];//描述
						html += '<tr><td>'+names+'</td><td>'+layout+'</td><td>'+depict+'</td><td width="5%" class="symbol2">+<div class="bolpoab2"><ul name="'+code+'"><li class="download">下载</li></ul></div></td></tr>';
					}
				}

				$(".edittable").html(html);
			}
		});
	}
	//Note
	var note = function(code){
		// 根据该镜头code 获得该code的消息列表
		shot_noteque_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var html = "";
				if (rs != null && rs.length > 0) {
					for(i=0; i<rs.length; i++){
						var pic = rs[i]["Picture"];
						var name = "a";
						if (pic == "#") {
							pic = ""
						}else{
							pic = "<img src='" + pic + "'><br/>"
						}
						var lor = "";
						if("1" != "1")
							lor = "l";
						else
							lor = "r";
						html += "<div class='newinfo"+lor+"'><div class='username'>"+name+"</div><div class='buddy'>" + pic + rs[i]["NoteDetail"] + "</div></div>";
					}
					var input = document.getElementById("fileimg");
					input.addEventListener('change', readImg, false);
				}

				$(".noteinfo").find(".chat").html(html);
			}
		});
	}
	//下载参考素材
	$(".edittable").on("click",".download",function(){
		var code = $(this).parent().attr("name");
		window.open("/post_shot_material_dow?MaterialCode=" + code);
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
		alert(code);
		// 把消息对应该镜头code记录到数据库 图片src为imgsrc 文字消息为txt
		shot_noteadd_ajax("aaa", code, txt, imgsrc, function(data){
			if (data.FeedbackCode == 0){
				var srchtml = "<img src='" + imgsrc + "'><br/>";
				var html = "";
				if (imgsrc == "#" && txt != "") {
					html = "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>"+txt+"</div></div>";
				} else if (imgsrc != "#" && txt != "") {
					var srchtml = "<img src='"+imgsrc+"'>";
					html = "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>"+srchtml+""+txt+"</div></div>";
				} else if (imgsrc != "#" && txt == "") {
					var srchtml = "<img src='" + imgsrc + "'>";
					html = "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>"+srchtml+"</div></div>";
				}
				//添加到页面
				$(".chat").append(html);
				//初始化文本框
				$(".textarea").val("");
				$("#imgs").attr("src","#");
				//初始化下拉
				$(".chat").scrollTop(200000);
			}
			alert(data.FeedbackCode+"--->"+data.FeedbackText);
		});
	});
	$(".rightdiv").on("click",".updfile",function(){
		$("#cameraCode").val($(this).parent().siblings(".code").val());
		$("#pcodes").val($(this).parent().siblings(".pcode").val());
	});
	//上传小样窗口显示
	$(".rightdiv").on("click",".sample",function(){
		//得到当前镜头code
		var code = $(this).parent().siblings(".code").val();
		//得到当前镜头pcode
		var pcode = $(this).parent().siblings(".pcode").val();
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
		$(".formdiv1").find(".pcode").val(pcode);
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
