$(function(){
	//初始化消息列表
	var node = function(){
		var html = "";
		for(i=0; i<5; i++){
			var pic = "#";
			var name = "a";
			var notedetail = "neirong";
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
			html += "<div class='newinfo"+lor+"'><div class='username'>"+name+"</div><div class='buddy'>" + pic + notedetail + "</div></div>";
		}
		$(".rightdiv").find(".chatdiv").html(html);
	}
	node();
	$(".chatdiv").scrollTop(100000);
	$(".yorn").click(function(){
		var s = $(".rightdiv").css("display");
		if(s=="block"){
			$(".rightdiv").hide(500);
		}else{
			$(".rightdiv").show(500);
		}
	});
	$(".but").click(function(){
		var imgstr = $(".sendout img").attr("src");
		var textstr = $(".textarea").val();
		var name = "a";
		var html = "";
		if(imgstr.length==1){
			if(textstr==""){
				return;
			}
			html += "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>" + textstr + "</div></div>";
		}else{
			imgstr = "<img src='"+imgstr+"'>";
			html += "<div class='newinfor'><div class='username'>"+name+"</div><div class='buddy'>" + imgstr + textstr + "</div></div>";
		}
		$(".chatdiv").append(html);
		$(".chatdiv").scrollTop(100000);
		$(".sendout img").attr("src","#");
		$(".textarea").val("");
	});
	/*画图js*/
	var canvas = document.getElementById('canvas');
	var canvasnone = document.getElementById('canvasnone');
	var video=document.getElementById("video");
	var height = $("#video").height();
	var width = $("#video").width();
	canvas.setAttribute("height",height);
	canvas.setAttribute("width",width);
	var mleft = (1200-width)/2;
	$("#canvas").css("margin-left",mleft+"px");
	setTimeout('sx()',2000);
	/*画笔颜色*/
	$(".colordiv table tbody tr td").click(function(){
		var color = $(this).css("background-color");
		$("#nonediv #strokeStyle").html(color);
		$(".colordiv span").css("background-color",color);
		$(".colordiv table tbody tr").children().css("border","2px solid #ccc");
		$(this).css("border","2px solid #000");
	});
	/*画笔粗细*/
	$(".widthdiv table tbody tr td").click(function(){
		var width = $(this).html();
		$("#nonediv #lineWidth").html(width);
		$(".widthdiv span").html(width);
		$(".widthdiv table tbody tr").children().css("border","2px solid #ccc");
		$(this).css("border","2px solid #000");
	});
	$(".keepbut").click(function(){
		var bufferCanvas = document.getElementById("canvasnone");
		var buffer = bufferCanvas.getContext("2d");
		buffer.drawImage(video,0,0,bufferCanvas.width,bufferCanvas.height);
		
		var mycanvas = document.getElementById("canvas");
		var ctx = mycanvas.getContext("2d");
		buffer.drawImage(mycanvas,0,0,mycanvas.width,mycanvas.height);
		
		var mycanvas = document.getElementById("canvasnone");
		var image = mycanvas.toDataURL("image/png");
		$("#imgs").attr("src",image);
		
		var mycanvas = document.getElementById("canvas");
		var ctx = mycanvas.getContext("2d");
		ctx.clearRect(0,0,mycanvas.width,mycanvas.height);
	});
	video.addEventListener("timeupdate", function(){
		var display = $(".entry").val();
		if(display == '入点出点开'){
			var endwid = $(".cd").width();
			var end = (100/1175)*endwid;
			var ends = end/100*video.duration;
			if(ends>=Math.round(video.currentTime)){
				$(".dq").html(Math.round(video.currentTime));
				var time = video.currentTime;
				var alltime = video.duration;
				var dqtime = time*(100/alltime);
				$(".widdiv").css("width",dqtime+"%");
			}else{
				if($(".loop").val()=="循环播放开"){//循环播放
					var startwid = $(".rd").width();
					var start = (100/1175)*startwid;
					$(".widdiv").css("width",start+"%");
					document.getElementById('video').currentTime = (start/100*video.duration);
				}else{
					$(".widdiv").css("width",end+"%");
					video.pause();
				}
			}
		}
		else{
			$(".dq").html(Math.round(video.currentTime));
			var time = video.currentTime;
			var alltime = video.duration;
			var dqtime = time*(100/alltime);
			$(".widdiv").css("width",dqtime+"%");
		}
	});
	$(".play").click(function(){
		var display = $(".entry").val();
		if(display == '入点出点关'){
			video.play();
		}else{
			if($("#rdcd").html()=="关"){
				$("#rdcd").html("开");
				var startwid = $(".rd").width();
				var start = (100/1175)*startwid;
				$(".widdiv").css("width",start+"%");
				document.getElementById('video').currentTime = (start/100*video.duration);
			}
			video.play();
		}
		var bufferCanvas = document.getElementById("canvas");
		var ctx = bufferCanvas.getContext("2d");
		ctx.clearRect(0,0,bufferCanvas.width,bufferCanvas.height);
	});
	$(".paused").click(function(){
		var bufferCanvas = document.getElementById("canvas");
		var ctx = bufferCanvas.getContext("2d");
		ctx.clearRect(0,0,bufferCanvas.width,bufferCanvas.height);
		video.pause();
	});
	$(".cTimetop").click(function(){
		video.pause();
		document.getElementById('video').currentTime = document.getElementById('video').currentTime-0.02;
	});
	$(".cTimebottom").click(function(){
		video.pause();
		document.getElementById('video').currentTime = document.getElementById('video').currentTime+0.02;
	});
	yorn = false;
	$(".butdiv2").mousedown(function(){
		widjs();
		yorn = true;
	});
	$(".butdiv2").mousemove(function(){
		if(yorn){
			widjs();
		}
	});
	$(".butdiv2").mouseup(function(){
		yorn = false;
	});
	$(".butdiv2").mouseleave(function(){
		yorn = false;
	});
	
	rudian = false;
	$(".rdcd1").mousedown(function(){
		rudianjs();
		rudian = true;
	});
	$(".rdcd1").mousemove(function(){
		if(rudian){
			rudianjs();
		}
	});
	$(".rdcd1").mouseup(function(){
		rudian = false;
	});
	$(".rdcd1").mouseleave(function(){
		rudian = false;
	});
	
	chudian = false;
	$(".rdcd2").mousedown(function(){
		chudianjs();
		chudian = true;
	});
	$(".rdcd2").mousemove(function(){
		if(chudian){
			chudianjs();
		}
	});
	$(".rdcd2").mouseup(function(){
		chudian = false;
	});
	$(".rdcd2").mouseleave(function(){
		chudian = false;
	});
	
	//全屏
	$(".fullscr").click(function(){
		var element = video;
	  if(element.requestFullScreen) {
		element.requestFullScreen();
	  } else if(element.mozRequestFullScreen) {
		element.mozRequestFullScreen();
	  } else if(element.webkitRequestFullScreen) {
		element.webkitRequestFullScreen();
	  }
	});
	//入点出点
	$(".entry").click(function(){
		var str = $(this).val();
		if(str=="入点出点关")
		{
			$(this).val("入点出点开");
			$(".rddiv").slideDown("slow");
			$(".cddiv").slideDown("slow");
		}else{
			$(this).val("入点出点关");
			$(".rddiv").slideUp("slow");
			$(".cddiv").slideUp("slow");
		}
	});
	//循环播放
	$(".loop").click(function(){
		if(video.loop==true){
		video.loop = false;
		$(".loop").val("循环播放关");
		}else{
			$(".loop").val("循环播放开");
			video.loop = true;
			if(video.currentTime==video.duration){
				video.currentTime = 0;
				video.play();
			}
		}
	});
	//监听键盘事件
	document.onkeydown=function(event){
		if(document.activeElement.className=="textarea"){
			return;
		}
		e = event ? event :(window.event ? window.event : null);
		//播放/暂停
		if(e.keyCode==32){
			if(video.paused){
				$(".play").click();
			}else{
				$(".paused").click();
			}
		}
		//上一帧
		if(e.keyCode==37){
			$(".cTimetop").click();
		}
		//下一帧
		if(e.keyCode==39){
			$(".cTimebottom").click();
		}
		//全屏
		if(e.keyCode==70){
			$(".fullscr").click();
		}
		//截取
		if(e.keyCode==67){
			$(".keepbut").click();
		}
		//入点出点
		if(e.keyCode==73){
			$(".entry").click();
		}
		//循环播放开关
		if(e.keyCode==76){
			$(".loop").click();
		}	
	}
});
function widjs(){
	var x=event.offsetX;
	var time = (100/1175)*x;
	$(".widdiv").css("width",time+"%");
	document.getElementById('video').currentTime = (time/100*video.duration);
}
function rudianjs(){
	var x=event.offsetX;
	var time = (100/1175)*x;
	$(".rd").css("width",time+"%");
	$("#rdcd").html("关");
}
function chudianjs(){
	var x=event.offsetX;
	var time = (100/1175)*x;
	$(".cd").css("width",time+"%");
	$("#rdcd").html("关");
}
function sx(){
	var video=document.getElementById("video");
	var height = $("#video").height();
	var width = $("#video").width();
	canvas.setAttribute("height",height);
	canvas.setAttribute("width",width);
	var mleft = (1200-width)/2;
	$("#canvas").css("margin-left",mleft+"px");
	canvasnone.setAttribute("height",height);
	canvasnone.setAttribute("width",width);
	$(".yg").html(Math.round(video.duration));
	$(".dq").html(Math.round(video.currentTime));
}

$(function(){
var canvas = document.createElement('canvas');
var point = {};
point.notFirst = false;
var clickX = [];
var clickY = [];
var clickDrag = [];
var paint;

var context = canvas.getContext("2d");
$("#canvas").mousedown(function(e){//鼠标点击画布
	//获得当前在画布的坐标
	var x = $(".all").css("margin-left");//.indexOf("px");
	var xpxvar = x.indexOf("px");
	x = x.substr(0,xpxvar);
	var sx = (x*1)+(this.offsetLeft*1);
	var mouseX = e.pageX - sx;//当前坐标＝鼠标距离窗口的坐标-画布距离窗口的坐标
	var mouseY = e.pageY - 60;
	paint = true;
	addClick(mouseX,mouseY);
	redraw();
});
$("#canvas").mousemove(function(e){//鼠标点下去不放拖动
	if(paint){
		var x = $(".all").css("margin-left");//.indexOf("px");
		var xpxvar = x.indexOf("px");
		x = x.substr(0,xpxvar);
		var sx = (x*1)+(this.offsetLeft*1);
		addClick(e.pageX - sx, e.pageY - 60, true);
		redraw();
	}
});
$('#canvas').mouseup(function(e){
  paint = false;
});

$('#canvas').mouseleave(function(e){
  paint = false;
});



function addClick(x,y,dragging){
	var sx = clickX.push(x);
	var sy = clickY.push(y);
	clickDrag.push(dragging);
}
function redraw(){
	var canvas = document.getElementById('canvas');
	var context = canvas.getContext("2d");
	context.lineJoin = "round";
	context.strokeStyle = $("#strokeStyle").html();
	context.lineWidth = $("#lineWidth").html();
	while(clickX.length > 0){
		context.beginPath();
		point.tx = point.x;
		point.ty = point.y;
		point.x = clickX.pop();
		point.y = clickY.pop();
		point.drag = clickDrag.pop();
		if(point.drag && point.notFirst){
			context.moveTo(point.tx, point.ty);
		}else{
			point.notFirst = true;
			context.moveTo(point.x-1, point.y);
		}
		context.lineTo(point.x, point.y);
		context.stroke();
		context.closePath();
	}
}
});






