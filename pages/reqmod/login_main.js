require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "js/jquery-1.10.2.min",
        "user" : "reqmod/user",
        "domready" : "reqmod/domReady",
        "pm" : "reqmod/pmodal"
    }
});

require(['domready', 'jquery', 'user','pm'], function (doc, $, user,pm){
    //binding event is ok！
    doc(function () {
        $('#login_button')[0].onclick = function(){user.login();}
        pm.load();
    });
});