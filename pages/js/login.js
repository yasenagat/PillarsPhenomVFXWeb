require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery" : "jslib/jquery-1.10.2.min",
		"bootstrap": "jslib/bootstrap.min",
        "domready" : "jslib/domReady",
        "pm" : "reqmod/pmodal",
        "user" : "js/user"
    },
	shim: {
        'bootstrap': {
            deps: ['jquery']
        }
    }
});

require(['domready', 'jquery', 'user','pm'], function (doc, $, user, pm){
    //binding event is ok！
    doc(function () {
        $('#loginPage_loginButton')[0].onclick = function(){user.login();}
        pm.load();
    });
});
