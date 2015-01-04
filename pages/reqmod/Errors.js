define(['reqmod/cookie','pm'],function (cookie,pm){
    var noLogin = function()
    {
        pm.login();
    }
    return {
        noLogin :noLogin
    };
});
