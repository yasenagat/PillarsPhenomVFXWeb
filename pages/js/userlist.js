define(['reqmod/cookie','pm'],function (cookie,pm){
	var login = function () {
		console.log("login business start");
        UserID  = $('#userName')[0].value;
        passwd  = $('#userPassword')[0].value;
        if(UserID == "" || passwd == ""){
            pm.show('e',"用户名或者密码不能为空");
            return;
        }

        login_core(UserID,passwd,function(data){
            if(data.FeedbackCode ==0)
            {
                window.location="/userlist";
            }
            else
            {
                pm.show('e',"用户名或密码错误");
            }
        });
    }

    var login_core = function(UserID,passwd,callback){
        window.res = {};
		//Password=a&Email=a
        $.post('/login_action',
            {
                Email:UserID,
                Password:passwd},
            function(data)
            {
                callback(data);
            },
            "json"
        );
    }

	return {
        test: test,
        register: register,
        login: login,
        logout: logout,
        login_core:login_core,
        checkLogin:checkLogin
    };

});
