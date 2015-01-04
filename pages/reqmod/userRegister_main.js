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

require(['domready', 'jquery','pm','user','bootstrap'], function (doc, $,pm,user){
    //binding event is ok！
    doc(function () {
        btn_register.onclick=function(){
			user.register();
		}
        pm.load();
    });
});