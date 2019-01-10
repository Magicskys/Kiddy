
function request_get_general() {
    $.ajax({
        url: '/v1/setting/general',
        type: 'GET',
        async : false,
        dataType: 'json'
    }).then(function (res) {
        if(res.code==200){
            console.log(res);
            re=true;
        }
    }).fail(function () {
        re=false;
    });
    return re;
}
