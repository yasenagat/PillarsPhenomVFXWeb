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
		"editor":"reqmod/editor"
    },
	shim: {
        'bootstrap': {
            deps: ['jquery']
        },
		'editor':{
			deps:['jquery']
		}
    }
});

require(['domready', 'jquery','pm','user','bootstrap','editor'], function (doc, $,pm,user){
    //binding event is ok！
    doc(function () {
        pm.load();
		console.log('dom is ready');
		$('#editor').wysiwyg();
    });
});