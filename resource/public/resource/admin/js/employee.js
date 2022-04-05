layui.extend({
	admin: '{/}../../static/admin/js/admin'
});

layui.use(['table', 'jquery','form', 'admin', 'laydate'], function() {
	var table = layui.table,
		$ = layui.jquery,
		form = layui.form,
		admin = layui.admin,
    	laydate = layui.laydate;
    var nowPage = $('.layui-laypage-skip>.layui-input').val();
    table.render({
        elem: '#employeesList',
        cellMinWidth: 60,
        cellMinHeight: 100,
		url:location.href,
        limit:10
        ,cols: [
            [{
                field: 'id',title: 'ID',width:70, sort: true
            }, {
                field: 'avatar',width:100,title: '头像',style:'height:50px;width:50px;', templet:'<div><img src="{{ d.avatar }}"></div>'
            }, {
                field: 'name',title: '姓名'
            }, {
                field: 'sex',title: '性别',templet:'<div>{{ d.sex.text }}</div>'
            }, {
                field: 'work_id',title: '工号'
            }, {
                field: 'tel',title: '电话'
            }, {
                field: 'email',title: '邮箱'
            }, {
                field: 'dep_name',title: '所属部门'
            }, {
                field: 'createdate',title: '添加时间'
            }, {
                field: 'editdate',title: '修改时间'
            }, {
                field: 'operate',width:100,title: '操作',toolbar: '#operateTpl'
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
    //console.table(res);

	/*
	 *数据表格中form表单元素是动态插入,所以需要更新渲染下
	 * http://www.layui.com/doc/modules/form.html#render
	 * */
	$(function(){
		//form.render();
	});

	var active = {
        //表格内容重载
        reload: function (nowPage) {
            var nowPage = nowPage || 1;
            var depid = $('#depId');
            var name = $('#name');
            //执行重载
            table.reload('employeesList', {
                page: {
                    curr: nowPage //重新从第 N 页开始
                }
                , where: {
                    depid: depid.val(),
                    name: name.val()
                },
            });

        },
        /*导出excel*/
        exportExcel:function () {
            var depid = $('#depId').val();
            var name = $('#name').val();

            var index = layer.confirm('确定导出excel？', {
                btn: ['确定','取消']
            }, function(){
                window.open('exportExcel.html?depid='+depid+'&'+'name='+name, "_blank");
                layer.close(index);
            });

        }
	};

    $('.we-search .layui-btn, .demoTable .layui-btn').on('click', function(){
        var type = $(this).data('type');
        active[type] ? active[type].call(this) : '';
    });
/*
    //监听select
    form.on('select(depInfo)', function(data){
        //console.log(data.elem); //得到select原始DOM对象
        //console.log(data.value); //得到被选中的值
        //console.log(data.othis); //得到美化后的DOM对象
        //active.reload();
    });
*/
    /*单个删除*/
	window.ad_del = function(obj, id) {
		layer.confirm('确认要删除吗？', function(index) {
			//发异步删除数据
            $.post('del.html',{'empid':id},function (res) {
                if(res['code']==1){
                    //删除成功
                    layer.msg(res['msg'], {
                        icon: 6,
                        time: 1000
                    });
                    //重载当前页码的表格数据
                    var nowPage = $('.layui-laypage-skip>.layui-input').val();
                    active.reload(nowPage);
                    //active.reload();
                }else{
                    //删除失败
                    layer.msg(res['msg'], {
                        icon: 5,
                        time: 2000
                    });
                }
            });

		});
	}
    //检查同名员工
    function checkDuplicationOfName(){
	    $.get('duplicationOfName.html',function(res){
            //console.log(res);
            if(res['code']==1){
                layer.msg(res['msg'],{
                    time:10000
                    ,offset: 't' //具体配置参考：offset参数项
                    ,btnAlign: 'c' //按钮居中
                    ,shade: 0 //不显示遮罩
                    ,yes: function(){
                        layer.closeAll();
                    }
                });
            }
        });
    }
    checkDuplicationOfName();
    /*图片放大预览*/
    $('body').on('mouseenter','.layui-table-cell>img',function () {
        //console.log('显示大图');
        var index = layer.open({
            type: 1,
            title: false,
            closeBtn: 0,
            area: '100',
            skin: 'layui-layer-nobg', //没有背景色
            shade: false,
            shadeClose: true,
            content: "<img style='max-width:600px;' src='"+$(this).prop('src')+"'>"
        });
    });
    $('body').on('mouseleave','.layui-table-cell>img',function () {
        //console.log('取消显示大图');
        layer.close(layer.index);
    });
});
