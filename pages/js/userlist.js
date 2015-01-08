
function deleteuser(email) {
	alter_core('/del', email, function(data) {
        if(data.FeedbackCode == 0) {
            window.location = "/userlist";
        } else {
            alert("删除操作失败，请稍后重试！");
        }
    });
}

function updateuser(email) {
	window.location = "/addUser.html?a=" + email;
}

function alter_core(alter, email, callback) {
    window.res = {};
    $.post(alter,
		{
            Email:email
		},
        function(data) {
            callback(data);
        },
        "json"
    );
}
