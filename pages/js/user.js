define(['reqmod/cookie','pm'],function (cookie,pm){
    var test = function (){
        console.log('test successful');
    };

    var checkLogin =function(){
        $.post(window.location.protocol + '//' + window.location.host + "/ci.php/app",{
            action:'checkLogin',
            handler: cookie.getCookie('handler')
        },function(data){
            if(data!=1){
                pm.show('e','请您先登录系统');
            }
        },"json");
    }
    var register = function () {
        console.log('register business start');
        var Phone,passwd,Email,NickName;
        
        if($('#tb_pwd_conform')[0].value!=$('#tb_pwd')[0].value)
        {
            pm.show('e',"两次密码不一致");
            return;
        }
        else{
            passwd = $('#tb_pwd_conform')[0].value
        }
        var Email = $('#tb_email')[0].value;
		Phone=$('#tb_phone')[0].value;
		NickName = $('#tb_NickName')[0].value;
		//check value exist
        if(NickName == "" || passwd == "" || Email == "")
        {
            pm.show('e',"请填入必要信息");return;
        }
        if(passwd.length < 6)
        {
            pm.show('e',"密码有点短");return;
        }
        var reg = /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;
        if(!reg.test(Email))
        {
            pm.show('e',"email 错误");return;
        }
		//Password=a&DisplayName=a&Picture=a&Email=a&Phone=a
        $.post('/register_action',
            {
				Picture:"none",
				Phone:Phone,
				Password:passwd,
				Email:Email,
				DisplayName:NickName
			},
            function(data)
            {
               	if(data.FeedbackCode == 0)
				{
					window.location="/";
					return;
				}
				else
				{
					pm.show('e','注册失败');
					return;
				}
            },
            "json"
        );

    }

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
                window.location="/cameraLens.html";
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

    var logout =function(){
        cookie.deleteCookie('handler');
        location.href=require.baseUrl+'login.html';
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
