//得到url地址中code参数
var projectCode = "";
//url参数获取
var getUrlParam = function(name){
	var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
	var r = window.location.search.substr(1).match(reg);  //匹配目标参数
	if (r!=null) return unescape(r[2]); return null; //返回参数值
}

// JavaScript Document
$(function(){
	projectCode = getUrlParam("code");
	$("#updedl").click();
	//初始化右侧窗口 隐藏
	$(".metadata").children(".basicinfo").css("display","block");
	$(".videodiv").on("click",".videoimg",function(){
		var rightabs = $(".rightdivabs").css("display");
		if(rightabs=="block"){
			$(".rightdivabs").css("display","none");
		}else{
			$(".rightdivabs").css("display","block");
		}
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
	});

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
	$(".del").click(function(){
		$(this).parents(".piece").remove();
	});
	$(".addmark").click(function(){
		var html = '<div class="piece">                    	<div class="number">2                        	<div class="del">X</div>                        </div>                        <div class="news">                        	<div class="imgradius">                            	<img src="#" width="60" height="60">                            </div>                            <div class="stage">                            fas                            </div>                        </div>                    </div>';
		$(".needinfo").append(html);
	});
	//发送消息
	$(".but").click(function(){
		//获得文本框内容
		var txt = $(".textarea").val();
		var html = '<div class="newinfo">'+txt+'</div>';
		//添加到页面
		$(".chat").append(html);
		//初始化文本框
		$(".textarea").val("");
		//初始化下拉
		$(".chat").scrollTop(100000);
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
	});
});
function upload(fnUpload) {
    var filename = fnUpload.value;
    var mime = filename.toLowerCase().substr(filename.lastIndexOf("."));
    if (mime != ".edl") {
        alert("请选择edl格式的文件上传");
        fnUpload.outerHTML = fnUpload.outerHTML;
        return false;
    } else {
        $("#f1").submit();
    }
}
function uploadAndSubmit() {
    var form = document.forms["f1"];
    if (form["file"].files.length > 0) {
        // 寻找表单域中的 <input type="file" ... /> 标签
        var file = form["file"].files[0];
        // try sending
        var reader = new FileReader();

        reader.onloadstart = function() {
            // 这个事件在读取开始时触发
            console.log("onloadstart");
        }
        reader.onprogress = function(p) {
            // 这个事件在读取进行中定时触发
            console.log("onprogress");
        }

        reader.onload = function() {
            // 这个事件在读取成功结束后触发
            //alert("读取成功结束");
        }

        reader.onloadend = function() {
            // 这个事件在读取结束后，无论成功或者失败都会触发
            if (reader.error) {
                console.log(reader.error);
            } else {
                // 构造 XMLHttpRequest 对象，发送文件 Binary 数据
                var xhr = new XMLHttpRequest();
                xhr.open("POST", "/post_upload_edl");
                //xhr.overrideMimeType("application/octet-stream");
                //xhr.sendAsBinary(reader.result);
				/*
                xhr.onreadystatechange = function() {
                    if (xhr.readyState == 4) {
                        if (xhr.status == 200) {
                            console.log("upload complete");
                            console.log("response: " + xhr.responseText);
                        }
                    }
                }
				*/
				xhr.onload = function(){};

				var formData = new FormData();
				formData.append('name', 'file');
				formData.append('file', reader.result);
				xhr.send(formData);
            }
        }

        reader.readAsBinaryString(file);
    } else {
        alert("Please choose a file.");
    }
}
