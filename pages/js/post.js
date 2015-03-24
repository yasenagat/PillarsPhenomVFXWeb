/* --------------- TODO 后期添加页面数据验证 ----------------
	1. 镜头基本信息"帧速率","宽","高"必须录入数字格式
	2. 添加镜头"帧速率","宽","高"必须录入数字格式
   --------------------------- END -------------------------*/

// ----------------------- modify by chengxz -----------------------start
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
				if (shotList == null || shotList.length == 0){
					alert("EDL文件未能获取镜头数据!");
					return;
				}
				createShotPage(shotList,"");
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
//新增镜头
var shot_add_ajax = function(pc, sn, wh, ht, sf, st, sd, callback){
	$.post("/post_shot_add",
		JSON.stringify({ProjectCode: pc, ShotName: sn, Width:wh, Height:ht, ShotFps:sf, ShotType: st, ShotDetail:sd}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//删除镜头(暂定只能删除手动添加的)
var shot_del_ajax = function(sc, callback){
	$.post("/post_shot_del",
		JSON.stringify({ShotCode: sc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//修改镜头名称
var shot_updshotname_ajax = function(sc, sn, callback){
	$.post("/post_shot_updshotname",
		JSON.stringify({ShotCode: sc, ShotName: sn}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头基本信息更新保存
var shot_infoupd_ajax = function(sc, sn, wh, ht, sf, st, et, sd, callback){
	$.post("/post_shot_upd",
		JSON.stringify({ShotCode: sc, ShotName: sn, Width:wh, Height:ht, ShotFps:sf, StartTime:st, EndTime:et, ShotDetail:sd}),
        function(data) {
            callback(data);
        },
        "json"
    );
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
//镜头制作需求新增
var shot_demandadd_ajax = function(pc, sc, dd, p, callback){
	$.post("/post_shot_demand_add",
		JSON.stringify({ProjectCode: pc, ShotCode: sc, DemandDetail: dd, Picture: p}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头制作需求删除
var shot_demanddel_ajax = function(dc, callback){
	$.post("/post_shot_demand_del",
		JSON.stringify({DemandCode: dc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头制作需求更新
var shot_demandupd_ajax = function(dc, dd, p, callback){
	$.post("/post_shot_demand_upd",
		JSON.stringify({DemandCode: dc, DemandDetail: dd, Picture: p}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//创建页面镜头视图列表
var createShotPage = function(rs,flags){
	var html = "";
	if (rs != null && rs.length > 0) {
		for(i=0; i<rs.length; i++){
			var code = rs[i]["ShotCode"];//镜头id
			var names = rs[i]["ShotName"];//镜头名
			var pic = rs[i]["Picture"];//图片
			var liInfo = "";
			var shotflag = "";
			if(rs[i]["SourcePath"] == "Y") {
				liInfo += "<li>Source</li>";
			}
			if(rs[i]["DpxPath"] == "Y") {
				liInfo += "<li>DPX</li>";
			}
			if(rs[i]["JpgPath"] == "Y") {
				liInfo += "<li>JPG</li>";
			}
			if(rs[i]["MovPath"] == "Y") {
				liInfo += "<li>Mov</li>";
			}
			if(flags=="1"){
				if(rs[i]["ShotFlag"] == "1"){
					shotflag = "<div class='dels'>X</div>";
				}
			}
			html += "<span class='videoimg' id='span"+code+"'><div class='view'></div><input type='hidden' id='code' value='"+code+"'><input class='check' name='checks' type='checkbox' value='"+code+"'><div class='state'></div><input class='play' type='button' value=''><h2 class='names'>"+names+"</h2><div class='downdiv'><input class='downl' type='button' value='下载'><span class='disnone'><ul class='"+code+"'>"+liInfo+"</ul></span></div><div class='files'><img src='"+pic+"'></div>"+shotflag+"</span>";
		}
	}
	$(".videodiv").html(html);
}

//上传镜头参考素材
function uploadMaterial(selectFile) {
	var code = $(".float").find(".sourceid").val();
	var detail = $(".explain").val();// 描述
	var filename = selectFile.value;
	var type = filename.substr(filename.lastIndexOf(".")+1);//文件格式
	var edlfile = document.getElementById("materialfile");
	var files = edlfile.files;
	var formData = new FormData();
	for(var i=0; i<files.length; i++){
		var file = files[i];
		formData.append("files", file, file.name);
	}
	formData.append("ProjectCode", projectCode);
	formData.append("ShotCode", code);
	formData.append("MaterialType", type);
	formData.append("MaterialDetail", detail);
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "post_shot_material_add", true);
	xhr.onload = function(){
		if(xhr.status === 200){
			var rs = JSON.parse(xhr.responseText);
			if(rs.FeedbackCode == 0) {
				$("#materialfile").attr("disabled", "true");
				$(".outer").click();
				$("#f7").find("input").val("");
				material("");
			}else if (rs.FeedbackCode == 202) {
				alert("要上传的文件服务器已经存在!");
			}else {
				alert(rs.FeedbackText);
			}
		}else{//查询"Load EDL"的镜头
			alert("upload error");
		}
	}
	xhr.send(formData);
}
//镜头参考素材删除
var shot_materialdel_ajax = function(mc, callback){
	$.post("/post_shot_material_del",
		JSON.stringify({MaterialCode: mc}),
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
//镜头版本列表查询
var shot_demo_version_ajax = function(sc, callback){
	$.post("/post_shot_demo_version",
		JSON.stringify({ShotCode: sc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}

//左侧树形列表点击
var folders_click_ajax = function(pc, fi, callback){
	$.post("/post_shot_folder_shots",
		{ProjectCode: pc, FolderId: fi},
        function(data) {
			callback(data);
        },
        "json"
    );
}

//自定义分组点击事件
var li = function(string) {
	//点击行
	$("#treeflag").val("1");
	$(".dtree").find(".dTreeNode").css("color","none");
	$(".dtree").children("div.dTreeNode").siblings(".node").css("color","#b6b6b6");
	$(".dtree").find("."+string).parent().siblings(".node").css("color","#b6b6b6");
	// 查找该分组string的素材,返回素材列表
	folders_click_ajax(projectCode, string, function(data){
		if(data.FeedbackCode == 0){
			var rs = JSON.parse(data.Data);
			var flag = true;//是否包含素材,包含素材
			if(rs == null || rs.length == 0){//若列表length为0,flag=false
				$(".videodiv").html("");
				flag = false;
			}
			if(flag && $("#dd"+string).css("display")=="none"){//如果包含素材,且当前目录为隐藏
				$(".strdiv .videodiv").html("");
				createShotPage(rs,"");
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

//镜头外包商列表添加
var vendor_add_ajax = function(pc, vn, vd, callback){
	$.post("/post_shot_vendor_add",
		JSON.stringify({ProjectCode: pc, VendorName: vn, VendorDetail: vd}),
        function(data) {
			callback(data);
        },
        "json"
    );
}
//镜头外包商列表删除
var vendor_del_ajax = function(vc, callback){
	$.post("/post_shot_vendor_del",
		JSON.stringify({VendorCode: vc}),
        function(data) {
			callback(data);
        },
        "json"
    );
}
//镜头外包商列表添加镜头
var shot_vendor_addshots_ajax = function(pc, vc, scs, callback){
	$.post("/post_shot_vendor_addShots",
		JSON.stringify({ProjectCode: pc, VendorCode: vc, ShotCodes: scs}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头外包商列表删除镜头
var shot_vendor_delshots_ajax = function(pc, vc, scs, callback){
	$.post("/post_shot_vendor_delShots",
		JSON.stringify({ProjectCode: pc, VendorCode: vc, ShotCodes: scs}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头外包商列表查询镜头
var shot_vendor_queshots_ajax = function(pc, vc, callback){
	$.post("/post_shot_vendor_queShots",
		JSON.stringify({ProjectCode: pc, VendorCode: vc}),
        function(data) {
            callback(data);
        },
        "json"
    );
}
//镜头外包商列表指定外包商
var vendor_specify_ajax = function(vc, vu, callback){
	$.post("/post_shot_vendor_specify",
		JSON.stringify({VendorCode: vc, VendorUser: vu}),
        function(data) {
			callback(data);
        },
        "json"
    );
}
//镜头外包商列表描述修改
var vendor_detail_ajax = function(vc, vd, callback){
	$.post("/post_shot_vendor_detail",
		JSON.stringify({VendorCode: vc, VendorDetail: vd}),
        function(data) {
			callback(data);
        },
        "json"
    );
}
//镜头外包商列表设置权限
var vendor_auth_ajax = function(vc, odl, od, dm, ud, up, callback){
	$.post("/post_shot_vendor_auth",
		JSON.stringify({VendorCode: vc, OpenDetail: odl, OpenDemo: od, DownMaterial: dm, UpDemo: ud, UpProduct: up}),
        function(data) {
			callback(data);
        },
        "json"
    );
}
//镜头外包商列表加载
var vendor_list_ajax = function(pc, callback){
	$.post("/post_shot_vendor_list",
		JSON.stringify({ProjectCode: pc}),
        function(data) {
			callback(data);
        },
        "json"
    );
}
//镜头外包商列表单条信息查询
var vendor_que_ajax = function(vc, callback){
	$.post("/post_shot_vendor_que",
		JSON.stringify({VendorCode: vc}),
        function(data) {
			callback(data);
        },
        "json"
    );
}

//查询分包商列表
var user_vendor_ajax = function(callback){
	$.post("/user_vendor_list","",
        function(data) {
            callback(data);
        },
        "json"
    );
}

// ----------------------- modify by chengxz -----------------------start



// JavaScript Document
$(function(){

	user_vendor_ajax(function(data){
		if (data.FeedbackCode == 0){
			var rs = JSON.parse(data.Data);
			var html = "<option value=''>-- 请选择 --</option>";
			if(rs != null && rs.length > 0){
				for(var i = 0; i<rs.length; i++){
					html += "<option value="+rs[i]["VendorUser"]+">"+rs[i]["UserName"]+"</option>";
				}
			}

			$("#factorysel").html(html);
		}
	});
	projectCode = getUrlParam("code");
	if (projectCode !== ""){
		$(".editoralpage").attr("href","editoral.html?code=" + projectCode);
		folders_ajax(projectCode);
	}else{
		return;
	}
	vendor_list_ajax(projectCode, function(data){
		if (data.FeedbackCode == 0) {
			var rs = JSON.parse(data.Data);
			var html = "";
			if(rs != null && rs.length > 0){
				for(var i = 0; i<rs.length; i++){
					html += '<li class="li1 li'+rs[i]["VendorCode"]+'" name="'+rs[i]["VendorCode"]+'"><a href="javascript:void(0);">'+rs[i]["VendorName"]+'</a><span class="addvgr">+<div><ul><li class="addlens">添加镜头</li><li class="factory" id="'+rs[i]["VendorUser"]+'">指定外包商</li><li class="with">描述<input type="hidden" value="'+rs[i]["VendorDetail"]+'"></li><li class="power">设置权限</li><li class="delfactory">删除</li></ul></div></span></li>';
				}
			}
			$(".vendiv").find(".venul").html(html);
		}
	});

	shots_ajax(projectCode, function(data){
		if(data.FeedbackCode == 0) {
			var rs = JSON.parse(data.Data);
			if(rs == null || rs.length === 0){
				// 没有数据,"Load EDL"可以点击
				$(".updedl").css("display","block");
				return;
			}
			// 有数据,"Load EDL"禁止点击,创建页面镜头视图列表
			$(".updedl").css("display","none");
			$(".loadedl").html("Shot List");
			createShotPage(rs,"1");
		}
	});

	$(".dtree").children("div.dTreeNode").css("background","#e3e3e3");
	$(".loadedl").click(function(){
		if($(this).html().trim()=="Shot List"){
			shots_ajax(projectCode, function(data){
				if(data.FeedbackCode == 0) {
					var rs = JSON.parse(data.Data);
					createShotPage(rs,"1");
				}
			});
		}
	});
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
			//获取img的url
			$(".roughimg").children("img").attr("src",$(this).siblings(".files").children("img").attr("src"));
			// 根据当前选中镜头的id thiscode，查询该镜头id的信息:镜头名 尺寸 帧速率 始码 止码 描述
			shot_info_ajax(thiscode, function(data){
				if(data.FeedbackCode == 0) {
					var rs = JSON.parse(data.Data);
					var names = rs["ShotName"];//镜头名
					var size = rs["Width"]+" X "+rs["Height"];//尺寸
					var speed = rs["ShotFps"];//帧速率
					var start = rs["StartTime"];//始码
					var end = rs["EndTime"];//止码
					//处理长字符串
					if(names.length>12)
						names = names.substring(0,11)+"...";
					if(size.length>12)
						size = size.substring(0,11);
					if(speed.length>12)
						speed = speed.substring(0,11);
					if(start.length>12)
						start = start.substring(0,11);
					if(end.length>12)
						end = end.substring(0,11);
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
		//alert("要删除的镜头code：" + str);
		// TODO 后台只接受一个镜头的删除
		var code = str.replace(',','');
		shot_del_ajax(code, function(data){
			if (data.FeedbackCode == 0){
				//成功删除
			}else{
				//没有删除
			}
		});
	});
	//便捷的修改素材名
	$(".bgimgdiv").on("click",".names",function(){
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
	$(".bgimgdiv").on("blur",".updlab",function(){//释放焦点时 修改数据库中该code的镜头名
		//获得当前镜头的code
		var code = $(this).parent().siblings("#code").val();
		//获得修改后文本框的文字
		var names = $(this).val();
		var temp = $(this);
		// 往数据库中根据该code修改该镜头的镜头名称
		shot_updshotname_ajax(code, names, function(data){
			if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				//修改页面显示内容
				temp.parent().html(rs["ShotName"]);
			}
		});
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
		// 根据code获得该镜头信息
		shot_info_ajax(code, function(data){
			if(data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				$(".tab2").find(".names").html(rs["ShotName"]);//镜头名
				$(".tab2").find(".size").find(".sizeW").html(rs["Width"]);//宽
				$(".tab2").find(".size").find(".sizeH").html(rs["Height"]);//高
				$(".tab2").find(".speed").html(rs["ShotFps"]);//帧速率
				$(".tab2").find(".start").html(rs["StartTime"]);//始码
				$(".tab2").find(".end").html(rs["EndTime"]);//止码
				$(".tab2").find(".bewrite").html(rs["ShotDetail"]);//描述
			}
		});
	}
	var need = function(code){
		//制作需求
		$(".needinfo .make").html("");
		// 根据镜头id查询制作需求list,并遍历
		shot_demandlist_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var html = "";//要拼接的html
				if(rs != null && rs.length > 0){
					for(i = 0; i < rs.length; i++){
						var code = rs[i]["DemandCode"];
						var imgurl = rs[i]["Picture"];
						var mation = rs[i]["DemandDetail"];
						var length = (i*1)+1;
						html += '<div class="piece"><input type="hidden" class="makecode" value="'+code+'"><div class="number">'+length+'<div class="del">X</div></div><div class="news"><div class="imgradius"><img src="'+imgurl+'" width="60" height="60"></div><div class="stage">'+mation+'</div><input type="button" value="编辑" class="edit"></div>                        </div>';
					}
				}
				$(".needinfo .make").html(html);
			}
		});
	}

	// 镜头参考素材列表加载
	var material = function(code){
		$(".edittable").html("");
		code = $(".float").find(".sourceid").val();
		// 根据该镜头code查询该镜头的参考素材列表
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
						html += "<tr><td>"+names+"</td><td>"+layout+"</td><td>"+depict+"</td><td class='symbol2'  width='5%'>+<div class='bolpoab2'><ul name='"+code+"'><li class='download'>下载</li><li class='delete'>删除</li></ul></div></td></tr>";
					}
				}

				$(".edittable").html(html);
			}
		});
	}

	//编辑详细信息
	$("#submits").click(function(){
		//获得该按钮上的文字 判断该编辑还是直接保存
		var butval = $(this).val();
		if(butval == "编辑"){
			//获得当前镜头ID和字段内容 并变成可编辑的文本框
			var names = $(".basicinfo").find(".names").html().trim();//镜头名
			var sizew = $(".basicinfo").find(".size").find(".sizeW").html().trim();//Width
			var sizeh = $(".basicinfo").find(".size").find(".sizeH").html().trim();//Height
			var speed = $(".basicinfo").find(".speed").html().trim();//帧速率
			var start = $(".basicinfo").find(".start").html().trim();//始码
			var end = $(".basicinfo").find(".end").html().trim();//止码
			var bewrite = $(".basicinfo").find(".bewrite").html().trim();//描述
			$(".basicinfo").find(".names").html("<input type='text' id='namesinp' value='"+names+"'>");
			$(".basicinfo").find(".size").html("W<span class='sizeW'><input type='text' value='"+sizew+"'></span>H<span class='sizeH'><input type='text' value='"+sizeh+"'></span>");
			$(".basicinfo").find(".speed").html("<input type='text' id='speedinp' value='"+speed+"'>");
			$(".basicinfo").find(".start").html("<input type='text' id='startinp' value='"+start+"'>");
			$(".basicinfo").find(".end").html("<input type='text' id='endinp' value='"+end+"'>");
			$(".basicinfo").find(".bewrite").html("<input type='text' id='bewriteinp' value='"+bewrite+"'>");
			$(this).val("保存");
		}else{//保存
			var names = $(".basicinfo").find("#namesinp").val();//镜头名
			var sizew = parseInt($(".basicinfo").find(".sizeW").find("input").val());//尺寸width
			var sizeh = parseInt($(".basicinfo").find(".sizeH").find("input").val());//尺寸height
			var speed = parseInt($(".basicinfo").find("#speedinp").val());//帧速率
			var start = $(".basicinfo").find("#startinp").val();//始码
			var end = $(".basicinfo").find("#endinp").val();//止码
			var bewrite = $(".basicinfo").find("#bewriteinp").val();//描述
			var code = $(".sourceid").val();
			var temp = $(this);
			// 后台保存方法
			shot_infoupd_ajax(code, names, sizew, sizeh, speed, start, end, bewrite, function(data){
				if (data.FeedbackCode == 0) {
					var rs = JSON.parse(data.Data);
					//文本框改成字符串
					$(".basicinfo").find(".names").html(rs["ShotName"]);
					$(".basicinfo").find(".size").html("<span class='sizeW'>"+rs["Width"]+"</span> X <span class='sizeH'>"+rs["Height"]+"</span>");
					$(".basicinfo").find(".speed").html(rs["ShotFps"]);
					$(".basicinfo").find(".start").html(rs["StartTime"]);
					$(".basicinfo").find(".end").html(rs["EndTime"]);
					$(".basicinfo").find(".bewrite").html(rs["ShotDetail"]);
					temp.val("编辑");
				}
			});
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
			// 当前需求id makecode 当前需求字符串 makeval 更新到数据库
			var p = "";//图片不作为更新内容,保留
			var temp = $(this);
			shot_demandupd_ajax(makecode, makeval, p, function(data){
				if (data.FeedbackCode == 0) {
					alert("更新成功");
					temp.siblings(".stage").html(makeval);
					temp.val("编辑");
				}else{
					alert("保存失败,请稍后重试!");
				}
			});
		}
	});

	//详细信息页面 编辑制作需求
	$(".make").on("click",".addedit",function(){
		var makeval = $(this).siblings(".stage").children("#demand").val();//当前制作需求输入框
		if(makeval==""){
			alert("请输入需求");
			$("#demand").focus();
			return;
		}
		// 当前镜头id code src地址 srcstr 和 需求字符串 makeval 添加到数据库
		var code = $(".sourceid").val();
		var srcstr = $(this).siblings(".imgradius").children("img").attr("src");
		var temp = $(this);
		shot_demandadd_ajax(projectCode, code, makeval, srcstr, function(data){
			if (data.FeedbackCode == 0) {
				temp.siblings(".stage").html(makeval);
				temp.attr("class","edit");
				temp.val("编辑");
				temp.siblings(".imgradius").children("#file_input").remove();
			}else{
				alert("保存失败,请稍后重试!");
			}
		});
	});

	//制作需求的删除
	$(".make").on("click",".del",function(){
		var temp = $(this);
		//获得该制作需求code
		var code = temp.parent().siblings(".makecode").val();
		// 从数据库中删除该code
		shot_demanddel_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				temp.parents(".piece").remove();
			}
		});
	});

	//
	$(".addmark").click(function(){
		if($(".needinfo").find("#demand").length>0){//判断是否已经打开判断的窗口
			alert("您正在添加");
			$("#demand").focus();
			return;
		}
		var length = ($(".make").children(".piece").length*1)+1;
		var html = '<div class="piece"><input type="hidden" class="makecode" value="1"><div class="number">'+length+'<div class="del">X</div></div><div class="news"><div class="imgradius"><img src="#" width="60" height="60"><input type="file" id="file_input"/></div><div class="stage"><textarea id="demand" rows="" cols="" name=""></textarea></div><input type="button" value="保存" class="addedit"></div></div>';
		$(".needinfo .make").append(html);
		var input = document.getElementById("file_input");
		input.addEventListener('change', readFile, false);
		$("#demand").focus();
	});
	//参考素材的下载
	$(".edittable").on("click",".download",function(){
		//获取当前参考素材的code
		var code = $(this).parent().attr("name");
		// 数据库根据该code获取相应的下载地址
		window.open("/post_shot_material_dow?MaterialCode=" + code);
	});
	//参考素材的删除
	$(".edittable").on("click",".delete",function(){
		//获取当前参考素材的code
		var code = $(this).parent().attr("name");
		var temp = $(this);
		// 数据库删除该code
		shot_materialdel_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				//删除该页面对应的数据
				temp.parents("tr").remove();
			}else{
				alert("删除失败,请稍后重试!(ErrorCode:" + data.FeedbackCode + ")");
			}
		});
	});
	//note消息读取
	var node = function(code){
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
	//发送消息
	$(".but").click(function(){
		//获得图片消息
		var imgsrc = $("#imgs").attr("src");
		//获得文本框消息
		var txt = $(".textarea").val();
		//获得镜头code
		var code = $(".sourceid").val();
		var name = "a";
		if(imgsrc=="#"&&txt==""){
			alert("请输入消息内容或图片");
			return;
		}
		// 把消息对应该镜头code记录到数据库 消息为news
		shot_noteadd_ajax(projectCode, code, txt, imgsrc, function(data){
			if (data.FeedbackCode == 0){
				var srchtml = "<img src='"+imgsrc+"'><br/>";
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
				//添加到页面
				$(".chat").append(html);
				//初始化文本框
				$(".textarea").val("");
				$("#imgs").attr("src","#");
				//初始化下拉
				$(".chat").scrollTop(200000);
			}
		});
	});
	//根据镜头code 查询版本列表信息
	var edition = function(code){
		//TODO 根据code 获得版本信息
		shot_demo_version_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var html = "";
				if(rs != null && rs.length > 0){
					for(var i=0; i<rs.length; i++){
						var code = rs[i]["VersionCode"];//当前版本code
						var num = "V " + rs[i]["VersionNum"];//版本号
						var img = rs[i]["Picture"];//缩略图url
						var detail = rs[i]["DemoDetail"];//描述
						html += '<div class="tag" name="'+code+'"><div class="num">'+num+'</div><div class="frame"><div class="imgs"><img src="'+img+'" alt="缩略图"><br/>'+detail+'</div><div class="menu_tab">+<div class="bolpoabs"><ul><li class="viewhue">查看小样</li><li class="down">下载成品</li></ul></div></div></div></div>';
					}
				}
				$(".editioninfo").html(html);
			}
		});
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
		$(".formdiv3").find(":text").val("");
		$(".formdiv3").find("textarea").val("");
		$(".formdiv3").hide(500);
		$(".formdiv4").find(":text").val("");
		$(".formdiv4").find("textarea").val("");
		$(".formdiv4").hide(500);
		$(".formdiv5").find(":text").val("");
		$(".formdiv5").find("textarea").val("");
		$(".formdiv5").hide(500);
		$(".formdiv5").find(".modify").val("编辑");
		$(".formdiv6").hide(500);
		$(".formdiv7").hide(500);
		$(".formdiv1").find(".size").val("1*2");
		$(".formdiv1").find("#sizew").attr("disabled","true");
		$(".formdiv1").find("#sizeh").attr("disabled","true");
	});//外包商弹出层
	$(".venul").on("click",".addvgr",function(){
		if($(this).children("div").css("display")=="block"){
			$(".venul").find(".addvgr div").css("display","none");
			$(this).children("div").css("display","none");
		}
		else{
			$(".venul").find(".addvgr div").css("display","none");
			$(this).children("div").css("display","inline-block");
		}

	});
	//添加外包商
	$("#addven").click(function(){
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({
			"height": height + "px",
			"width": width + "px"
		});
		$(".formdiv3").css({
			"left": (width / 2) - 200 + "px",
			"top": (height / 2) - 100 + "px"
		});
		$(".outer").show(500);
		$(".formdiv3").show(500);
	});
	//点击添加列表的确定按钮
	$(".addvenbut").click(function(){
		//获得列表名称
		var names = $("#packinginp").val();
		//获得外包商描述
		var descrinp = $(".formdiv3").find("#descrinp").val();
		if(names==""){
			alert("请输入外包商");
			return;
		}
		if(descrinp==""||descrinp==null){
			alert("请输入描述");
			return;
		}
		//TODO 添加外包商 names 跟描述 descrinp 到数据库 并返回刚添加外包商之后的id
		vendor_add_ajax(projectCode, names, descrinp, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var code = rs["VendorCode"];
				// TODO 增加2隐藏字段存储1分包商2描述
				var html = '<li class="li1 li'+rs[i]["VendorCode"]+'" name="'+code+'"><a href="javascript:void(0);">'+names+'</a><span class="addvgr">+<div><ul><li class="addlens">添加镜头</li><li class="factory">指定外包商</li><li class="with">描述</li><li class="power">设置权限</li><li class="delfactory">删除</li></ul></div></span></li>';
				$(".venul").append(html);
				$(".outer").click();
			}
		});
	});
	//分包商列表添加镜头
	$(".venul").on("click",".addlens",function(){
		var code = $(this).parents(".li1").attr("name");//获得当前组的code
		//获得当前选中镜头的id
		var str = new Array();
		$('input[class="check"]:checked').each(function() {
			str.push($(this).val());
		});
		if(str.length == 0){
			alert("请选中镜头后再添加");
			return;
		}
		// 添加镜头
		shot_vendor_addshots_ajax(projectCode, code, str, function(data){
			if (data.FeedbackCode == 0) {
				alert("添加成功！");
			}
		});
	});
	//指定外包商窗口
	$(".venul").on("click",".factory",function(){
		//获得当前分组的code
		var code = $(this).parents(".li1").attr("name");
		//code赋给隐藏文本框
		$(".formdiv4").find(".code").val(code);
		//获得当前分组绑定的外包商code
		var vencode = $(this).attr("id");
		//赋值给隐藏文本框
		$(".formdiv4").find(".vencode").val(vencode);
		$("#factorysel").val(vencode);
		//显示窗口
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({
			"height": height + "px",
			"width": width + "px"
		});
		$(".formdiv4").css({
			"left": (width / 2) - 200 + "px",
			"top": (height / 2) - 100 + "px"
		});
		$(".outer").show(500);
		$(".formdiv4").show(500);
	});
	$(".addoutsidebut").click(function(){
		//获得当前列表code
		var listcode = $(".formdiv4").find(".code").val();
		//获得当前选中外包公司code
		var company = $(".formdiv4").find("#factorysel").val();
		if(company=="")
		{
			alert("请选择外包商");
			return;
		}
		//TODO 设置列表code 外包公司code
		vendor_specify_ajax(listcode, company, function(data){
			if (data.FeedbackCode == 0) {
				$(".li"+listcode).find(".factory").attr("id",company);
				$(".outer").click();
			}
		});
	});
	$(".venul").on("click",".with",function(){
		//获得列表code
		var listcode = $(this).parents(".li1").attr("name");
		//TODO 根据列表code 获得描述信息
		var description = $(this).find("input").val();//"这是描述内容";
		//赋值列表code给描述页面
		$(".formdiv5").find(".code").val(listcode);
		//赋值描述信息给描述页面
		$(".formdiv5").find(".descrtr").html(description);
		//显示窗口
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({
			"height": height + "px",
			"width": width + "px"
		});
		$(".formdiv5").css({
			"left": (width / 2) - 200 + "px",
			"top": (height / 2) - 100 + "px"
		});
		$(".outer").show(500);
		$(".formdiv5").show(500);
	});
	$(".modify").click(function(){
		var txt = $(this).val();
		if(txt=="编辑"){//编辑 把html变为编辑框
			//获取描述的文字
			var html = $(".descrtr").html().trim();
			$(".descrtr").html("<input type='text' value="+html+">");
			$(this).val("保存");
		}else{
			//获得该列表值的code
			var listcode = $(".formdiv5").find(".code").val();
			var value = $(".descrtr").find("input").val();
			// 保存列表值code和描述
			vendor_detail_ajax(listcode, value, function(data){
				if (data.FeedbackCode == 0) {
					$(".descrtr").html(value);
					$(".modify").val("编辑");
					$(".li"+listcode).find(".with").find("input").val(value);
				}
			});

		}
	});
	$(".upddescbut").click(function(){
		if($(".formdiv5").find(".descrtr input").val()!=null){
			alert("请保存后再确定");
			$(".formdiv5").find(".descrtr input").focus();
			return;
		}
		$(".outer").click();
	});
	//权限设置
	$(".venul").on("click",".power",function(){
		//取该列表id设置到页面
		var code = $(this).parents(".li1").attr("name");
		$(".formdiv6").find(".code").val(code);
		// 根据该列表id 获取该id权限标识
		vendor_que_ajax(code, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				var flog = rs["UserCode"];
				var result = flog.split(",");
				for(var i = 0; i<result.length; i++){
					if(result[i] == "0"){
						$(".formdiv6").find("input[type='checkbox']").eq(i).attr("checked",false);
					}else{
						$(".formdiv6").find("input[type='checkbox']").eq(i).attr("checked",true);
					}
				}
				//显示窗口
				var height = $(window).height();
				var width = $(window).width();
				$(".outer").css({
					"height": height + "px",
					"width": width + "px"
				});
				$(".formdiv6").css({
					"left": (width / 2) - 200 + "px",
					"top": (height / 2) - 100 + "px"
				});
				$(".outer").show(500);
				$(".formdiv6").show(500);
			}
		});
	});
	$(".yesbut").click(function(){
		//获取当前列表code
		var code = $(".formdiv6").find(".code").val();
		//获得当前选中权限的标识
		var str = new Array();

		$('input[name="droit"]').each(function() {
			if($(this).prop("checked")){
				str.push(1);
			}else{
				str.push(0);
			}
		});
		vendor_auth_ajax(code, str[0], str[1], str[2], str[3], str[4], function(data){
			if (data.FeedbackCode == 0) {
				$(".outer").click();
			}
		});

	});

	//
	$(".venul").on("click",".delfactory",function(){
		//获得当前列表code，
		var code = $(this).parents(".li1").attr("name");

		//TODO 从数据库中删除该列表code

		//从页面中删除
		$(this).parents(".li1").remove();
	});
	//
	$(".venul").on("click",".li1 a",function(){
		//获得当前code
		var code = $(this).parent().attr("name");
		// 根据该code 从后台取 并便利到页面
		shot_vendor_queshots_ajax(projectCode, code, function(data){
			if (data.FeedbackCode == 0) {
				var rs = JSON.parse(data.Data);
				createShotPage(rs,"");
			}
		});
	});

	//添加镜头
	$(".subut").click(function(){
		var names = $(".formdiv1").find(".name").val();//获取镜头名
		if(names==""){
			alert("请输入镜头名");
			$(".formdiv1").find(".name").focus();
			return;
		}
		//获取尺寸 宽 sizew 高 sizeh
		var sizes = $(".formdiv1").find(".size").val();
		var sizew = "";
		var sizeh = "";
		if(sizes=="自定义"){
			sizew = $(".formdiv1").find("#sizew").val();
			sizeh = $(".formdiv1").find("#sizeh").val();
			if(sizew==""){
				alert("请输入宽");
				$(".formdiv1").find("#sizew").focus();
				return;
			}
			if(sizeh==""){
				alert("请输入高");
				$(".formdiv1").find("#sizeh").focus();
				return;
			}
			//匹配正则表达式
			var reg = /^[1-9]\d*$/;
			if(!reg.test(sizew)){
				alert("请输入数字");
				$(".formdiv1").find("#sizew").focus();
				return;
			}
			if(!reg.test(sizeh)){
				alert("请输入数字");
				$(".formdiv1").find("#sizeh").focus();
				return;
			}
		}else{
			//获取*号的下标
			var index = sizes.indexOf("*");
			sizew = sizes.substr(0,index);
			sizeh = sizes.substr(index+1,sizes.length);
		}
		//获取帧速率
		var speed = $(".formdiv1").find(".speed").val();
		if(speed==""){
			alert("请输入帧速率");
			$(".formdiv1").find(".speed").focus();
			return;
		}
		//获取类型
		var types = $(".formdiv1").find(".types").val();
		//获取描述
		var description = $(".formdiv1").find(".description").val();
		// 添加到数据库
		shot_add_ajax(projectCode, names, parseInt(sizew), parseInt(sizeh), parseInt(speed), types, description, function(data){
			if (data.FeedbackCode == 0) {
				$(".outer").click();
				//TODO 新增镜头添加到页面显示
			}
		});
	});

	//
	$(".addmaterial").click(function(){
		//显示窗口
		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({
			"height": height + "px",
			"width": width + "px"
		});
		$(".formdiv7").css({
			"left": (width / 2) - 200 + "px",
			"top": (height / 2) - 100 + "px"
		});
		$(".outer").show(500);
		$(".formdiv7").show(500);
	});
	//参考素材描述框焦点
	$(".explain").blur(function(){
		if($(this).val()!=""){
			$("#materialfile").removeAttr('disabled');
		}else{
			$("#materialfile").attr("disabled", "true");
		}
	});

	//下载
	$(".videodiv").on("click",".disnone ul li",function(){
		var source = $(this).html().trim();
		var code = $(this).parent().attr("class");
	});
	//删除
	$(".videodiv").on("click",".dels",function(){
		//得到当前镜头code
		var code = $(this).siblings("#code").val();
		var temp = $(this);
		shot_del_ajax(code, function(data){
			if (data.FeedbackCode == 0){
				temp.parents(".videoimg").remove();
			}
		});
	});
});
function selectChange(val){
	if(val!="自定义"){
		$("#sizew").attr("disabled","true");
		$("#sizeh").attr("disabled","true");
	}else{
		$("#sizew").val("");
		$("#sizeh").val("");
		$("#sizew").removeAttr('disabled');
		$("#sizeh").removeAttr('disabled');
	}
}
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
function readFile() {
	var file = this.files[0];
	//判断类型如果不是图片就返回 去掉就可以上传任意文件
	if (!/image\/\w+/.test(file.type)) {
		alert("请确保文件为图像类型");
		return false;
	}
	var reader = new FileReader();
	reader.readAsDataURL(file);
	reader.onload = function(e) {
		$("#file_input").siblings("img").attr("src",this.result);
	}
}
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
		$(".sendout").children("#imgs").attr("src",this.result);
	}
}
