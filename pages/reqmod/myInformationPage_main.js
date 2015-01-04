require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "js/jquery-1.10.2.min",
        "domready" : "reqmod/domReady",
        "pm" : "reqmod/pmodal",
		"cookie":"reqmod/cookie",
		"user":"reqmod/user",
		"bootstrap": "reqmod/bootstrap.min",
    },
	shim: {
        'bootstrap': {
            deps: ['jquery']
        }
    }
});

require(['domready', 'jquery', 'user','pm'], function (doc, $, user,pm){
    //binding event is ok！
	item_template = "<div class='horizontal-List'  >"+
"<div  id='listTitleContent'><a href='/wiki/TTTitleTTT'>"+
"TTTitleTTT"+
"</a></div>"+
"<div  id='dateTimeContent'>"+
"TTTimeTTT"+
"<button onclick=\"window.location='/updateWiki.html?t=SSSegmentCodeSSS'\" class='btn btn-warning'>修改</button>"+
"<button data-title='TTTitleTTT' class='btn btn-danger delete'>删除</button>"+
"</div>"+
"</div> "
	
    doc(function () {
		pm.load();
        // 查询我的wiki
		$.post('/wiki_get_list',{},function(data){
			if(data.FeedbackCode == 1)
			{
				pm.show('e','您没有登录，或者登录超时。请您重新登录');
			}
			else
			{
				var knowledges = eval('(' + data.FeedbackText + ')');
				console.log(knowledges);
				res  = ""
				for(i=0 ; i<knowledges.length ; i++){
					res += item_template
						.replace("TTTitleTTT",knowledges[i].Title)
						.replace("TTTitleTTT",knowledges[i].Title)
						.replace("TTTitleTTT",knowledges[i].Title)
						.replace("TTTimeTTT","创建时间："+knowledges[i].InsertDatetime
						+"&nbsp;&nbsp;&nbsp;最后修改时间："
						+knowledges[i].UpdateDatetime+"&nbsp;&nbsp;&nbsp;")
						.replace("SSSegmentCodeSSS",knowledges[i].Title)
				}
				wiki_list.innerHTML=res;
			}
			btn_delete = document.getElementsByClassName('delete');
			for (i=0 ;i< btn_delete.length ; i++)
			{
				btn_delete[i].onclick=function(){
					console.log(this.attributes['data-title'].value);
					$.post('/wiki_delete',{
						Title:this.attributes['data-title'].value
					},function(data){
						if(data.FeedbackCode != 0)
						{
							pm.show('通知','删除失败');
						}else
						{
							pm.show('通知','删除成功');
							this.innerHTML='已删除';
							this.onclick=function(){console.log('已经解绑')}
						}
					},"json");
					
				}
			}
		},"json");

	});
});