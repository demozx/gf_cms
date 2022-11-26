layui.use(['table', 'jquery','form', 'laydate'], function() {
	var table = layui.table,
        form = layui.form,
		$ = layui.jquery;
    table.render({
        elem: '#employeesList',
        cellMinWidth: 90,
        cellMinHeight: 100,
		url:location.href,
        limit:10
        ,cols: [
            [{type:'checkbox',fixed: 'left'}
            ,{
                field: 'id',title: 'ID', sort: true
            }, {
                field: 'name',title: '员工姓名'
            }, {
                field: 'dep_name',title: '所属部门'
            }, {
                field: 'work_id',title: '员工工号'
            }, {
                field: 'tel',title: '员工电话',width:140
            }, {
                field: 'ip',title: '查询者IP',width:140
            }, {
                field: 'address',title: '查询者地址',width:200
            }, {
                field: 'time',title: '查询时间',width:170
            }, {
                field: 'from',title: '查询来源'
            }, {
                field: 'mode',title: '查询方式'
            },{
                field: 'read',title: '状态',templet: '#checkboxTpl', unresize: true
            }]
        ],

        event: true,
        page: {
            layout: ['prev','next','page','count','skip','limit'] //自定义分页布局
            //,curr: 5 //设定初始在第 5 页
            ,groups: 5 //只显示 N 个连续页码
            ,first: false //不显示首页
            ,last: false //不显示尾页
		},
    });
    //单个修改员工查询状态
    form.on('switch(read)', function(data){
        //console.log(data.value);
        sendIds(data.value);
    });

    var active = {
        getCheckData: function(){ //获取选中数据
            var checkStatus = table.checkStatus('employeesList')
                ,data = checkStatus.data;
            //layer.alert(JSON.stringify(data));
            var ids = new Array();
            if(data.length > 0) {
                for (var i = 0; i < data.length; i++) {
                    ids.push(data[i]['id']);
                }
                //动态修改switch值
                $('.switch').each(function(){
                    for(var i=0;i<ids.length;i++){
                        if($(this).val() == ids[i]){
                            $(this).prop('checked',true);
                        }
                    }
                });
                sendIds(ids,true);
            }else{
                layer.msg('请勾选要操作的行',{
                    'anim':6
                });
            }
            //console.log(ids);
        }
    };

    $('.demoTable .layui-btn').on('click', function(){
        var type = $(this).data('type');
        active[type] ? active[type].call(this) : '';
    });

    //提交ids
    function sendIds(ids,batch) {
		var batch = batch || false;
        $.post('read.html',{'ids':ids,'batch':batch},function (res) {
            //console.log(res);
            if(res['code']==1){
                layer.msg(res['msg'],{
                    icon: 6
                    ,time: 1000
                });
            }else{
                layer.msg(res['msg'],{
                    icon: 5
                    ,time: 1000
                });
            }
			
			form.render();
                //重新检测是否有新的员工查询
                $.get('monitorEmployeeLog.html',function (res) {
                    if(res>0){
                        $('.countemployeeLog').text(res);
                        $('.countemployeeLog').addClass('layui-badge');
                    }else{
                        $('.countemployeeLog').text('');
                        $('.countemployeeLog').removeClass('layui-badge');
                    }
                });
			
        });
    }

});
