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

var download_core = function(code, type, callback) {
    $.post("/editoral_download_file1", JSON.stringify({MaterialCode: code, SourceType: type}),
        function(data) {
            callback(data);
        },
        "json"
    );
}

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


	// 搜索按钮，测试后台方法
	$(".butsearch").click(function(){
		var code = "sadgasgewe0";
		var type = "Source";
		var a = post("/editoral_download_file", {MaterialCode: code, SourceType: type});
		alert(a)
	});

});
