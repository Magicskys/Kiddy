function request_scan(idlist,action){
    var re;
    $.ajax({
        url: '/v1/task/'+action,
        type: 'POST',
        data: {"uid":idlist},
        async : false,
        dataType: 'json'
    }).then(function (res) {
        if(res.code==200){
            console.log(res);
            re=true;
        }
    }).fail(function () {
        console.log('失败');
        re=false;
    });
    return re;
}

function request_action_schema(schema,idlist,action) {
    var re;
    $.ajax({
        url: '/v1/action/'+schema,
        type: 'POST',
        data: {"action":action,"uid":idlist},
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

function request_monitor(monitor) {
    var re;
    $.ajax({
        url: '/v1/setting/monitor',
        type: 'POST',
        data: {"monitor":monitor},
        async : false,
        dataType: 'json'
    }).then(function (res) {
        if(res.code==200){
            console.log(res);
            re=true;
        }
        if (res.code==201){
            if (res.message==false){
                re=false
            }else{
                re=true
            }
        }
    }).fail(function () {
        re=false;
    });
    return re;
}
