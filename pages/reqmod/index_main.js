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
    doc(function () {
		pm.load();
        $('#btn_login')[0].onclick = function(){user.login();}
        
    });
});