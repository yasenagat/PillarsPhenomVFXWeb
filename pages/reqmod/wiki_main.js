require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "/js/jquery-1.10.2.min",
        "domready" : "/reqmod/domReady",
        "pm" : "/reqmod/pmodal",
		"cookie":"/reqmod/cookie",
		"user":"/reqmod/user",
		"bootstrap": "/reqmod/bootstrap.min",
    },
	shim: {
        'bootstrap': {
            deps: ['jquery']
        }
    }
});

require(['domready', 'jquery', 'user','pm'], function (doc, $, user,pm){
	//follow_category,follow_wiki
	pm.load();
	$('#follow_category')[0].onclick=function(){
		console.log('category binding ok');
		console.log($(this).attr('source_code'))
		$.post('/subscription',{
			SorceCode:$(this).attr('source_code'),
			subType:"3"
		},function(data){
			if(data.FeedbackCode == 0){
				pm.show('','订阅成功');
			}
			else
			{
				pm.show('','订阅失败');
			}
		},"json"
		);
	}
	$('#follow_wiki')[0].onclick=function(){
		console.log('wiki binding ok');
		console.log($(this).attr('source_code'));
		$.post('/subscription',{
			SorceCode:$(this).attr('source_code'),
			subType:"2"
		},function(data){
			if(data.FeedbackCode == 0){
				pm.show('','订阅成功');
			}
			else
			{
				pm.show('','订阅失败');
			}
		},"json"
		);
	}
});