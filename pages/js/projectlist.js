
function deleteproject(code) {
	alter_core('/delproject', code, function(data) {
        if(data.FeedbackCode == 0) {
            window.location = "/projectlist";
        } else {
            alert("删除操作失败，请稍后重试！");
        }
    });
}

function updateproject(code) {
	window.location = "/addproject.html?a=" + code;
}

function alter_core(alter, code, callback) {
    window.res = {};
    $.post(alter,
		{
            ProjectCode:code
		},
        function(data) {
            callback(data);
        },
        "json"
    );
}
