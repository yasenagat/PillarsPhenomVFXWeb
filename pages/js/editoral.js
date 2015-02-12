$(function(){
	$("#addsou").click(function(){
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({"height":height+"px","width":width+"px"});
		$(".formdiv").css({"left":(width/2)-200+"px","top":(height/2)-200+"px"});
		$(".outer").show(500);
		$(".formdiv").show(500);
	});
	$(".outer").click(function(){
		$(".outer").hide(500);
		$(".formdiv").hide(500);
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
		//验证通过
		alert(name+"="+url+"="+transcodingurl+"="+dpxurl+"="+movurl);
		$(".outer").hide(500);
		$(".formdiv").hide(500);
	});
});