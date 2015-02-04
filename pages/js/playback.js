$(function(){
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
		if(imgstr.length==1){
			if(textstr==""){
				return;
			}
			var html = "<div class='news'>"+textstr+"</div>";
		}else{
			var html = "<div class='news'><img src='"+imgstr+"'>"+textstr+"</div>";
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
		/*keepraw();
		var mycanvas = document.getElementById("canvas");
		var ctx = mycanvas.getContext("2d");
		var bufferCanvas = document.getElementById("canvasnone");
		buffer.drawImage(bufferCanvas,10,0,ctx.width,ctx.height);

		var image = mycanvas.toDataURL("image/png");
		var imghtml = "<img src="+image+">";
		$(".imgs").append(imghtml);
		var height = $("#video").height();
		var width = $("#video").width();
		ctx.clearRect(0,0,width,height);
		*/
	});
	video.addEventListener("timeupdate", function(){
		$(".dq").html(Math.round(video.currentTime));
		var time = video.currentTime;
		var alltime = video.duration;
		console.log(time);
		var dqtime = time*(100/alltime);
		$(".widdiv").css("width",dqtime+"%");
//		$(".widdiv").animate({width:dqtime+'px'},50);
	});
	$(".play").click(function(){
		video.play();
	});
	$(".paused").click(function(){
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
});
function widjs(){
	var x=event.offsetX;
	var time = (100/1175)*x;
	$(".widdiv").css("width",time+"%");
	document.getElementById('video').currentTime = (time/100*video.duration);
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
		console.log(point);
	}
}
});


