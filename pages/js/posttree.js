// Node object
function Node(id, pid, name, url, title, target, icon, iconOpen, open) {
	this.id = id;
	this.pid = pid;
	this.name = name;
	this.url = url;
	this.title = title;
	this.target = target;
	this.icon = icon;
	this.iconOpen = iconOpen;
	this._io = open || false;
	this._is = false;
	this._ls = false;
	this._hc = false;
	this._ai = 0;
	this._p;
};

// Tree object
function dTree(objName) {
	this.config = {
		target: null,
		folderLinks: true,
		useSelection: true,
		useCookies: true,
		useLines: true,
		useIcons: false,
		useStatusText: false,
		closeSameLevel: false,
		inOrder: false
	}

	this.icon = {
		root: 'img/base.gif',
		folder: 'img/folder.gif',
		folderOpen: 'img/folderopen.gif',
		node: 'img/page.gif',
		empty: 'img/empty.gif',
		line: 'img/empty.gif',
		join: 'img/empty.gif',
		joinBottom: 'img/empty.gif',
		plus: 'img/empty.gif',
		plusBottom: 'img/empty.gif',
		minus: 'img/empty.gif',
		minusBottom: 'img/empty.gif',
		nlPlus: 'img/nolines_plus.gif',
		nlMinus: 'img/nolines_minus.gif'
	};

	this.obj = objName;
	this.aNodes = [];
	this.aIndent = [];
	this.root = new Node( - 1);
	this.selectedNode = null;
	this.selectedFound = false;
	this.completed = false;
};

// Adds a new node to the node array
dTree.prototype.add = function(id, pid, name, url, title, target, icon, iconOpen, open) {
	this.aNodes[this.aNodes.length] = new Node(id, pid, name, url, title, target, icon, iconOpen, open);
};

// Open/close all nodes
dTree.prototype.openAll = function() {
	this.oAll(true);
};
dTree.prototype.closeAll = function() {
	this.oAll(false);
};

// Outputs the tree to the page
dTree.prototype.toString = function() {
	var str = '<div class="dtree">\n';
	if (document.getElementById) {
		if (this.config.useCookies) this.selectedNode = this.getSelected();
		str += this.addNode(this.root);
	} else str += 'Browser not supported.';
	str += '</div>';
	if (!this.selectedFound) this.selectedNode = null;
	this.completed = true;
	return str;
};

// Creates the tree structure
dTree.prototype.addNode = function(pNode) {
	var str = '';
	var n = 0;
	if (this.config.inOrder) n = pNode._ai;
	for (n; n < this.aNodes.length; n++) {
		if (this.aNodes[n].pid == pNode.id) {
			var cn = this.aNodes[n];
			cn._p = pNode;
			cn._ai = n;
			this.setCS(cn);
			if (!cn.target && this.config.target) cn.target = this.config.target;
			if (cn._hc && !cn._io && this.config.useCookies) cn._io = this.isOpen(cn.id);
			if (!this.config.folderLinks && cn._hc) cn.url = null;
			if (this.config.useSelection && cn.id == this.selectedNode && !this.selectedFound) {
				cn._is = true;
				this.selectedNode = n;
				this.selectedFound = true;
			}
			str += this.node(cn, n);
			if (cn._ls) break;
		}
	}
	return str;
};

// Creates the node icon, url and text
dTree.prototype.node = function(node, nodeId) {
	if (node.target == "0") { //底层节点
		var str = '<div class="dTreeNode">' + this.indent(node, nodeId) + '<span class=\'' + node.id + '\'>+<div><ul><li class="addGroup">创建镜头组</li><li class="removeGroup">删除镜头组</li><li class="addMaterial">添加镜头</li><li class="removeMaterial">移除镜头</li><li class="updGroup">修改</li></ul></div></span>';
	} else { //非底层节点
		var str = '<div class="dTreeNode">' + this.indent(node, nodeId) + '<span class=\'' + node.id + '\'>+<div><ul><li class="addGroup">创建镜头组</li><li class="removeGroup">删除镜头组</li><li class="addMaterial">添加镜头</li><li class="removeMaterial">移除镜头</li><li class="updGroup">修改</li></ul></div></span>';
	}
	if (this.config.useIcons) {
		if (!node.icon) node.icon = (this.root.id == node.pid) ? this.icon.root: ((node._hc) ? this.icon.folder: this.icon.node);
		if (!node.iconOpen) node.iconOpen = (node._hc) ? this.icon.folderOpen: this.icon.node;
		if (this.root.id == node.pid) {
			node.icon = this.icon.root;
			node.iconOpen = this.icon.root;
		}
		str += '<img id="i' + this.obj + nodeId + '" src="' + ((node._io) ? node.iconOpen: node.icon) + '" alt="" />';
	}
	if (node.url) {
		//带htef的被注释
		//str += '<a id="s' + this.obj + nodeId + '" class="' + ((this.config.useSelection) ? ((node._is ? 'nodeSel' : 'node')) : 'node') + '" href="' + node.url + '"';
		str += '<a id="s' + this.obj + nodeId + '" class="' + ((this.config.useSelection) ? ((node._is ? 'nodeSel': 'node')) : 'node') + '"';
		if (node.title) str += ' title="' + node.title + '"';
		if (node.target) str += ' target="' + node.target + '"';
		if (this.config.useStatusText) str += ' onmouseover="window.status=\'' + node.name + '\';return true;" onmouseout="window.status=\'\';return true;" ';
		if (this.config.useSelection && ((node._hc && this.config.folderLinks) || !node._hc))
		//str += ' onclick="javascript: ' + this.obj + '.s(' + nodeId + ');"';
		str += ' href="javascript:void(0);" onclick="javascript: urlstr(\'' + node.id + '\');"';
		str += '>';
	}

	//else if ((!this.config.folderLinks || !node.url) && node._hc && node.pid != this.root.id)
	if (node.target == "1")
	//str += '<a href="javascript: ' + this.obj + '.o(' + nodeId + ');" class="node">';
	str += '<a href="javascript: li(' + node.id + ');" class="node">';
	else {
		str += '<a href="javascript: urlstr(' + node.id + ')" class="node">';
	}
	str += node.name;

	//	if (node.url || ((!this.config.folderLinks || !node.url) && node._hc))
	str += '</a>';
	str += '</div>';
	//if (node._hc) {
	str += '<div id="d' + this.obj + node.id + '" class="clip" style="display:' + ((this.root.id == node.pid || node._io) ? 'block': 'none') + ';">';
	str += this.addNode(node);
	str += '</div>';
	//}
	this.aIndent.pop();
	return str;
};

// Adds the empty and line icons
dTree.prototype.indent = function(node, nodeId) {
	var str = '';
	if (this.root.id != node.pid) {
		for (var n = 0; n < this.aIndent.length; n++)
		str += '<img src="' + ((this.aIndent[n] == 1 && this.config.useLines) ? this.icon.line: this.icon.empty) + '" alt="" />';
		(node._ls) ? this.aIndent.push(0) : this.aIndent.push(1);
		if (node._hc) {
			str += '<a href="javascript: li(' + node.id + ');"><img id="j' + this.obj + nodeId + '" src="';
			if (!this.config.useLines) str += (node._io) ? this.icon.nlMinus: this.icon.nlPlus;
			else str += ((node._io) ? ((node._ls && this.config.useLines) ? this.icon.minusBottom: this.icon.minus) : ((node._ls && this.config.useLines) ? this.icon.plusBottom: this.icon.plus));
			str += '" alt="" /></a>';
		} else {str += '<img src="' + ((this.config.useLines) ? ((node._ls) ? this.icon.joinBottom: this.icon.join) : this.icon.empty) + '" alt="" />';
		}
	}

	return str;
};

// Checks if a node has any children and if it is the last sibling
dTree.prototype.setCS = function(node) {
	var lastId;
	for (var n = 0; n < this.aNodes.length; n++) {
		if (this.aNodes[n].pid == node.id) node._hc = true;
		if (this.aNodes[n].pid == node.pid) lastId = this.aNodes[n].id;
	}
	if (lastId == node.id) node._ls = true;
};

// Returns the selected node
dTree.prototype.getSelected = function() {
	var sn = this.getCookie('cs' + this.obj);
	return (sn) ? sn: null;
};

// Highlights the selected node
dTree.prototype.s = function(id) {
	if (!this.config.useSelection) return;
	var cn = this.aNodes[id];
	if (cn._hc && !this.config.folderLinks) return;
	if (this.selectedNode != id) {
		if (this.selectedNode || this.selectedNode == 0) {
			eOld = document.getElementById("s" + this.obj + this.selectedNode);
			eOld.className = "node";
		}
		eNew = document.getElementById("s" + this.obj + id);
		eNew.className = "nodeSel";
		this.selectedNode = id;
		if (this.config.useCookies) this.setCookie('cs' + this.obj, cn.id);
	}
};

// Toggle Open or close
dTree.prototype.o = function(id) {
	var cn = this.aNodes[id];
	this.nodeStatus(!cn._io, id, cn._ls);
	cn._io = !cn._io;
	if (this.config.closeSameLevel) this.closeLevel(cn);
	if (this.config.useCookies) this.updateCookie();
};

// Open or close all nodes
dTree.prototype.oAll = function(status) {
	for (var n = 0; n < this.aNodes.length; n++) {
		if (this.aNodes[n]._hc && this.aNodes[n].pid != this.root.id) {
			this.nodeStatus(status, n, this.aNodes[n]._ls)
			this.aNodes[n]._io = status;
		}
	}
	if (this.config.useCookies) this.updateCookie();
};

// Opens the tree to a specific node
dTree.prototype.openTo = function(nId, bSelect, bFirst) {
	if (!bFirst) {
		for (var n = 0; n < this.aNodes.length; n++) {
			if (this.aNodes[n].id == nId) {
				nId = n;
				break;
			}
		}
	}
	var cn = this.aNodes[nId];
	if (cn.pid == this.root.id || !cn._p) return;
	cn._io = true;
	cn._is = bSelect;
	if (this.completed && cn._hc) this.nodeStatus(true, cn._ai, cn._ls);
	if (this.completed && bSelect) this.s(cn._ai);
	else if (bSelect) this._sn = cn._ai;
	this.openTo(cn._p._ai, false, true);
};

// Closes all nodes on the same level as certain node
dTree.prototype.closeLevel = function(node) {
	for (var n = 0; n < this.aNodes.length; n++) {
		if (this.aNodes[n].pid == node.pid && this.aNodes[n].id != node.id && this.aNodes[n]._hc) {
			this.nodeStatus(false, n, this.aNodes[n]._ls);
			this.aNodes[n]._io = false;
			this.closeAllChildren(this.aNodes[n]);
		}
	}
}

// Closes all children of a node
dTree.prototype.closeAllChildren = function(node) {
	for (var n = 0; n < this.aNodes.length; n++) {
		if (this.aNodes[n].pid == node.id && this.aNodes[n]._hc) {
			if (this.aNodes[n]._io) this.nodeStatus(false, n, this.aNodes[n]._ls);
			this.aNodes[n]._io = false;
			this.closeAllChildren(this.aNodes[n]);
		}
	}
}

// Change the status of a node(open or closed)
dTree.prototype.nodeStatus = function(status, id, bottom) {
	eDiv = document.getElementById('d' + this.obj + id);
	eJoin = document.getElementById('j' + this.obj + id);
	if (this.config.useIcons) {
		eIcon = document.getElementById('i' + this.obj + id);
		eIcon.src = (status) ? this.aNodes[id].iconOpen: this.aNodes[id].icon;
	}
	//eJoin.src = (this.config.useLines)?
//((status)?((bottom)?this.icon.minusBottom:this.icon.minus):((bottom)?this.icon.plusBottom:this.icon.plus)):
	//((status)?this.icon.nlMinus:this.icon.nlPlus);
	eDiv.style.display = (status) ? 'block': 'none';
};

// [Cookie] Clears a cookie
dTree.prototype.clearCookie = function() {
	var now = new Date();
	var yesterday = new Date(now.getTime() - 1000 * 60 * 60 * 24);
	this.setCookie('co' + this.obj, 'cookieValue', yesterday);
	this.setCookie('cs' + this.obj, 'cookieValue', yesterday);
};

// [Cookie] Sets value in a cookie
dTree.prototype.setCookie = function(cookieName, cookieValue, expires, path, domain, secure) {
	document.cookie =
	escape(cookieName) + '=' + escape(cookieValue)
	+ (expires ? '; expires=' + expires.toGMTString() : '')
	+ (path ? '; path=' + path: '')
	+ (domain ? '; domain=' + domain: '')
	+ (secure ? '; secure': '');
};

// [Cookie] Gets a value from a cookie
dTree.prototype.getCookie = function(cookieName) {
	var cookieValue = '';
	var posName = document.cookie.indexOf(escape(cookieName) + '=');
	if (posName != -1) {
		var posValue = posName + (escape(cookieName) + '=').length;
		var endPos = document.cookie.indexOf(';', posValue);
		if (endPos != -1) cookieValue = unescape(document.cookie.substring(posValue, endPos));
		else cookieValue = unescape(document.cookie.substring(posValue));
	}
	return (cookieValue);
};

// [Cookie] Returns ids of open nodes as a string
dTree.prototype.updateCookie = function() {
	var str = '';
	for (var n = 0; n < this.aNodes.length; n++) {
		if (this.aNodes[n]._io && this.aNodes[n].pid != this.root.id) {
			if (str) str += '.';
			str += this.aNodes[n].id;
		}
	}
	this.setCookie('co' + this.obj, str);
};

// [Cookie] Checks if a node id is in a cookie
dTree.prototype.isOpen = function(id) {
	var aOpen = this.getCookie('co' + this.obj).split('.');
	for (var n = 0; n < aOpen.length; n++)
	if (aOpen[n] == id) return true;
	return false;
};

// If Push and pop is not implemented by the browser
if (!Array.prototype.push) {
	Array.prototype.push = function array_push() {
		for (var i = 0; i < arguments.length; i++)
		this[this.length] = arguments[i];
		return this.length;
	}
};

if (!Array.prototype.pop) {
	Array.prototype.pop = function array_pop() {
		lastElement = this[this.length - 1];
		this.length = Math.max(this.length - 1, 0);
		return lastElement;
	}
};

// ----------------------- modify by chengxz -----------------------start
//新增镜头组
var folder_add_ajax = function(pc, fn, fc, fd, callback){
	$.post("/post_shot_folder_add",
		{ProjectCode: pc, FolderName: fn, FatherCode: fc, FolderDetail: fd},
        function(data) {
			callback(data);
        },
        "json"
    );
}

//删除镜头组
var folder_del_ajax = function(pc, fc, callback){
	$.post("/post_shot_folder_del",
		{ProjectCode: pc, FolderCode: fc},
        function(data) {
			callback(data);
        },
        "json"
    );
}

//修改镜头组
var folder_upd_ajax = function(fc, fn, fd, callback){
	$.post("/post_shot_folder_upd",
		{FolderCode: fc, FolderName: fn, FolderDetail: fd},
        function(data) {
			callback(data);
        },
        "json"
    );
}

//查询镜头组
var folder_que_ajax = function(fc, callback){
	$.post("/post_shot_folder_que", {FolderCode: fc},
        function(data) {
			callback(data);
        },
        "json"
    );
}

//镜头组添加镜头
var folder_addfiles_ajax = function(pc, fc, mc, callback){
	$.post("/post_shot_folder_addfiles",
		JSON.stringify({ProjectCode: pc, FolderCode: fc, ShotCodes: mc}),
        function(data) {
			callback(data);
        },
        "json"
    );
}

//镜头组删除镜头
var folder_delfiles_ajax = function(pc, fc, mc, callback){
	$.post("/post_shot_folder_delfiles",
		JSON.stringify({ProjectCode: pc, FolderCode: fc, ShotCodes: mc}),
        function(data) {
			callback(data);
        },
        "json"
    );
}

//镜头组查询是否添加镜头
var folder_countfiles_ajax = function(pc, fc, callback){
	$.post("/post_shot_folder_countfiles",
		JSON.stringify({ProjectCode: pc, FolderCode: fc,}),
        function(data) {
			callback(data);
        },
        "json"
    );
}
// ----------------------- modify by chengxz -----------------------end


$(function(){
	$(".tree").on("blur",".texinp",function(){
		var inputstr = $(this).val();//获得文本框内容
		var inputid = $(this).attr("id");//获得文本框id
		var pid = $(this).attr("name");//获得父id
		if(inputstr!="") {
			var html = "<a href='javascript:li("+inputid+");' class='node'>"+inputstr+"</a>";
			$(this).parent().append(html);
			$(this).siblings("span").css("display","");
			$(this).remove();
			//传输节点数据:节点字符串，节点id，该节点的父节点id
			number(inputstr,inputid,pid);
		}else{
			$(this).parent().remove();
			$("#dd"+inputid).remove();
		}
	});
	//添加镜头组
	$(".tree").on("click",".addGroup",function(){
		var nodesize = "";
		$(".tree").find(".dTreeNode").each(function() {
            if (nodesize = "") {
				nodesize = $(this).children("span").attr("class");
			} else {
				if(nodesize*1<$(this).children("span").attr("class")*1){
					nodesize = $(this).children("span").attr("class");
				}
			}
        });
		nodesize++;
		var thisclass = $(this).parents("span").attr("class");
		//同等级的img 的数量
		imglength = $(this).parents("span").siblings("img").length +
					$(this).parents("span").siblings("a").children("img").length;
		//$(".texinp").focus();

		var height = $(window).height();
		var width = $(window).width();
		$(".outer").css({"height":height+"px","width":width+"px"});
		$(".formdiv2").css({"left":(width/2)-200+"px","top":(height/2)-100+"px"});
		$(".formdiv2").find(".id").val(nodesize);
		$(".formdiv2").find(".pid").val(thisclass);
		$(".formdiv2").find(".imgsize").val(imglength);
		$(".outer").show(500);
		$(".formdiv2").show(500);
	});
	//重置添加镜头组表单input
	$(".formdiv2").find(".again").click(function(){
		$(".formdiv2").find(":text").val("");
	})
	//点击添加镜头组弹窗的确定按钮之后
	$(".formdiv2").find(".submit").click(function(){
		var thisclass = $(".formdiv2").find(".pid").val();//父ID
		var names = $(".formdiv2").find(".names").val();//镜头组的名称
		var depict = $(".formdiv2").find(".depict").val();//镜头的描述
		//根据是否有父id判断是添加还是修改,若父ID不等于空则为添加，若父ID等于空则为修改
		if(thisclass != ""){//添加
			var imglength = $(".formdiv2").find(".imgsize").val();//层级关系，界面需要
			// 往后台添加数据,thisclass是父id，names是镜头组的名称，depict是镜头组的描述，添加完成后
			folder_add_ajax(projectCode, names, thisclass, depict, function(data){
				if(data.FeedbackCode == 0){
					var rs = JSON.parse(data.Data);
					var nodesize = rs["FolderCode"];//后台传过来的添加完之后的id
					//rs["FolderName"]
					var imgstr = "";
					for (i=0; i<imglength*1+1; i++) {
						imgstr += "<img src='img/empty.gif'>";
					}
					$("#dd"+thisclass).css("display","block");
					$("#dd"+thisclass).append("<div class='dTreeNode'>"+imgstr+"<span class=\'"+nodesize+"\' style='display:;'>+<div><ul><li class='addGroup'>创建镜头组</li><li class='removeGroup'>删除镜头组</li><li class='addMaterial'>添加镜头</li><li class='removeMaterial'>移除镜头</li><li class='updGroup'>修改</li></ul></div></span><a href='javascript:li("+nodesize+");' class='node'>"+rs["FolderName"]+"</a></div><div id=\'dd"+nodesize+"\' class='clip' style='display:block;'></div>");
				}
			});
		}else{//修改
			// 往后台传数据，thisclass是当前镜头组要修改的id，names是当前镜头组名称，depice是镜头的描述
			var fid = $(".formdiv2").find(".id").val();
			folder_upd_ajax(fid, names, depict, function(data){
				if(data.FeedbackCode == 0){
					var rs = JSON.parse(data.Data);
					$(".tree").find("."+rs["FolderCode"]).siblings(".node").html(rs["FolderName"]);
				}
			});
		}
		$(".outer").hide(500);
		$(".formdiv2").find(":text").val("");
		$(".formdiv2").hide(500);
	});

	//点击右侧加号
	$(".tree").on("click","span",function(){
		//判断当前菜单是否显示
		if($(this).children("div").css("display")=="block"){
			$(this).children("div").css("display","none");
			return;
		}
		//初始化所有span
		$(".dTreeNode span div").css("display","none");
		var code = $(this).attr("class");//得到当前镜头组的id
		//得到当前镜头组的id，根据该id，判断该id是否有镜头
		var thisDom = $(this);
		// 根据该id code，求该镜头组id是否有镜头 有为true
		//求该镜头组id是否有下级目录
		folder_countfiles_ajax(projectCode, code, function(data){
			if(data.FeedbackCode == 0){
				var rs = JSON.parse(data.Data);
				//rs["IsHaveShot"],rs["IsHaveLeaf"],rs["FatherCode"]
				// 根据当前组id code 查询该id 的父id是否是-1,赋值给thispid
				if (rs["FatherCode"] == "-1") {
					thisDom.find(".removeGroup").css("display","none");
					thisDom.find(".updGroup").css("display","none");
				}else{
					thisDom.find(".removeGroup").css("display","block");
					thisDom.find(".updGroup").css("display","block");
				}
				//rs["IsHaveShot"]//该分组包含镜头,true是有镜头
				//rs["IsHaveLeaf"]//该分组是否有子目录,true是有
				if (rs["IsHaveShot"]) {
				//若该分组包含镜头 不能添加镜头组 可以添加镜头
					thisDom.find(".addGroup").css("display","none");
					thisDom.find(".addMaterial").css("display","block");
					thisDom.find(".removeMaterial").css("display","block");
				} else if (!rs["IsHaveShot"] && !rs["IsHaveLeaf"]){
				//若无镜头 没子目录 添加镜头 添加组
					thisDom.find(".addGroup").css("display","block");
					thisDom.find(".addMaterial").css("display","block");
					thisDom.find(".removeMaterial").css("display","none");
				} else if (!rs["IsHaveShot"] && rs["IsHaveLeaf"]){
				//若该分组不包含镜头 有子目录 能添加镜头组 不可以添加镜头
					thisDom.find(".addGroup").css("display","block");
					thisDom.find(".addMaterial").css("display","none");
					thisDom.find(".removeMaterial").css("display","none");
				}
				//判断该镜头组下拉菜单是否显示
				if(thisDom.children("div").css("display")=='none'){
					thisDom.children("div").css("display","inline-block");
				}else{
					thisDom.children("div").css("display","none");
				}
				if($("#treeflag").val()=="0"){
					thisDom.find(".removeMaterial").css("display","none");
				}
			}
		});
	});

	//添加镜头
	$(".tree").on("click",".addMaterial",function(){
		var thiscode = $(this).parents("span").attr("class");//当前镜头组id
		//获得当前选中复选框的id
		var strs = new Array();
		$('input[class="check"]:checked').each(function(){
			strs.push($(this).val());
		});
		if(strs.length == 0){
			alert("请选中右侧镜头复选框再添加镜头");
			return;
		}
		// 添加镜头id到该镜头组id,镜头组id是thiscode,选中的镜头id是strs
		folder_addfiles_ajax(projectCode, thiscode, strs, function(data){
			if(data.FeedbackCode == 0){
				alert("镜头添加成功!");
			}
		});
	});

	//移除镜头
	$(".tree").on("click",".removeMaterial",function(){
		var thiscode = $(this).parents("span").attr("class");//当前镜头组id
		//获得当前选中复选框的id
		var strs = new Array();
		$('input[class="check"]:checked').each(function(){
    		strs.push($(this).val());
		});
		if(strs.length==0){
			alert("请选中镜头再删除");
			return;
		}

		// 从镜头组id中删除选中镜头镜头的id:strs
		folder_delfiles_ajax(projectCode, thiscode, strs, function(data){
			if(data.FeedbackCode == 0){
				$('input[class="check"]:checked').each(function(){
					$("#span"+$(this).val()).remove();
				});
			}
		});
	});

	//删除镜头组
	$(".tree").on("click",".removeGroup",function(){
		var thiscode = $(this).parents("span").attr("class");//当前镜头组id
		var del = $(this).parents("span").parent();
		if(window.confirm('你确定要删除该分组吗？')){
			// 删除该镜头组id thiscode
			folder_del_ajax(projectCode, thiscode, function(data){
				if(data.FeedbackCode == 0){
					del.next("div").remove();
					del.remove();
				}
			});
		}
	});

	//修改信息
	$(".tree").on("click",".updGroup",function(){
		var thiscode = $(this).parents("span").attr("class");//当前镜头组id
		// 根据该镜头组id，查询数据库该镜头组的名称和描述信息，显示到界面文本框
		folder_que_ajax(thiscode, function(data){
			if(data.FeedbackCode == 0){
				var rs = JSON.parse(data.Data);
				var height = $(window).height();
				var width = $(window).width();
				$(".outer").css({"height":height+"px","width":width+"px"});
				$(".formdiv2").css({"left":(width/2)-200+"px","top":(height/2)-100+"px"});
				$(".formdiv2").find(".id").val(thiscode);
				$(".formdiv2").find(".pid").val("");
				$(".formdiv2").find(".names").val(rs["FolderName"]);
				$(".formdiv2").find(".depict").val(rs["FolderDetail"]);
				$(".outer").show(500);
				$(".formdiv2").show(500);
			}
		});
	});
});
