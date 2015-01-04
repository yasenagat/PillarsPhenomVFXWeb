define(['jquery','reqmod/cookie','bootstrap'],function ($,cookie){
    var modal_load = function () {
        $.get('/modal.html', function(result){
            $('body').append(result);
        });
    }

    var modal_show = function(title,contain){
        if(title=='e'){modal_title.innerHTML="<a style='color: red'>错误</a>";}
        if(title=='w'){modal_title.innerHTML="<a style='color: orange'>警告</a>";}
        if(title=='i'){modal_title.innerHTML="提示";}
        modal_contain.innerHTML=contain;
        $('#myModal').modal('show');
		//btn_modal.click();

    }

    return {
        load :modal_load,
        show :modal_show
    };
});
