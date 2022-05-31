layui.extend({
	admin: '{/}../../static/backend/js/backend'
});

layui.use(['table', 'jquery','form', 'admin', 'laydate'], function() {
	var table = layui.table,
		$ = layui.jquery,
		form = layui.form,
		admin = layui.admin,
    	laydate = layui.laydate;

    table.render({
        elem: '#departmentList',
        cellMinWidth: 60,
        cellMinHeight: 100,
		url:location.href,
        limit:20
        ,cols: [
            [{
                field: 'id',title: 'ID',width:60, sort: true
            }, {
                field: 'dep_name',title: '部门名称'
            }, {
                field: 'employees_count',title: '员工数量'
            }, {
                field: 'sort',title: '排序',templet:'<div style="line-height: 100px;"><input lay-verify="required" autocomplete="off" class="layui-input sort" id="sort_{{ d.id }}" data-id="{{ d.id }}" type="number" value="{{ d.sort }}"></div>'
            }, {
                field: 'operate',title: '操作',toolbar: '#operateTpl'
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
            //执行重载
            table.reload('departmentList', {
                page: {
                    curr: nowPage //重新从第 N 页开始
                }
            });
            //console.log('重载至第'+nowPage+'页');
        },
        /*排序*/
        Sort : function(){
        var arr = new Array();
        var work = false;
        $('.sort').each(function(){
            if( $(this).val()=='' ){
                layer.tips('不能为空', '#sort_'+$(this).attr('data-id'), {
                    tips: [4, '#78BA32']
                    ,tipsMore: false
                });
                work = false;
                return false;
            }else if( $(this).val()<0 ){
                layer.tips('必须为大于等于0的数字', '#sort_'+$(this).attr('data-id'), {
                    tips: [4, '#78BA32']
                    ,tipsMore: false
                });
                work = false;
                return false;
            }else{
                arr.push({'id':$(this).attr('data-id'),'sort':$(this).val()});
                work = true;
            }
        });
        if(work == true){
            $.post('sort.html',{'data':arr},function (res) {
                if(res['code']==1){
                    layer.msg('更新排序成功',{icon:6,time:1000},function () {
                        active.reload();
                    });
                }else{
                    layer.msg(res['msg'],{icon:5,time:2000});
                }
            },'json');
        }
        }

	};

    $('.we-search .layui-btn, .demoTable .layui-btn').on('click', function(){
        var type = $(this).data('type');
        active[type] ? active[type].call(this) : '';
    });
/*
    //监听select
    form.on('select(adschannel)', function(data){
        //console.log(data.elem); //得到select原始DOM对象
        //console.log(data.value); //得到被选中的值
        //console.log(data.othis); //得到美化后的DOM对象
        active.reload();
    });
*/
    /*单个删除*/
	window.ad_del = function(obj, id) {
		layer.confirm('确认要删除吗？', function(index) {
			//发异步删除数据
            $.post('del.html',{'depid':id},function (res) {
                if(res['code']==1){
                    //删除成功
                    layer.msg(res['msg'], {
                        icon: 6,
                        time: 1000
                    });
                    //重载当前页码的表格数据
                    //var nowPage = $('.layui-laypage-skip>.layui-input').val();
                    //active.reload(nowPage);
                    active.reload();
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

});
