require.baseUrl = "/"
require.config({
    baseUrl: require.baseUrl,
    paths: {
        "jquery": "jslib/jquery-1.10.2.min"
    }
});

require(['jquery'], function($) {
    $(function(){
		window.location.href = "/login.html";
    });
});
