/*
 * @Author: https://github.com/WangEn
 * @Author: https://gitee.com/lovetime/
 * @Date:   2018-03-27
 * @lastModify 2018-3-28
 * +----------------------------------------------------------------------
 * | WeAdmin 表格table中多个删除等操作公用js
 * | 有改用时直接复制到对应页面也不影响使用
 * +----------------------------------------------------------------------
 */
layui.extend({
    admin: '{/}../../static/backend/js/backend',
});
layui.use(['laydate', 'jquery', 'admin'], function() {
	var laydate = layui.laydate,
		$ = layui.jquery,
		admin = layui.admin;
	//执行一个laydate实例
	laydate.render({
		elem: '#start' //指定元素
	});
	//执行一个laydate实例
	laydate.render({
		elem: '#end' //指定元素
	});
	/*用户-停用*/
	window.member_stop = function (obj, id) {
			var url = 'status.html';
			if($(obj).prop('title') == '启用') {
                layer.confirm('确认要启用吗？', function(index) {
                    //发异步把用户状态进行更改
					$.get(url,{'admid':id},function(res){
						//console.log(res);
						if(res['code']==1){
                            $(obj).attr('title', '停用')
                            $(obj).find('i').html('&#xe601;');
                            $(obj).parents("tr").find(".td-status").find('span').removeClass('layui-btn-disabled').html('已启用');
                            layer.msg(res['msg'], {
                                time: 1000
                                ,icon:6
                            });
						}else{
                            layer.msg(res['msg'], {
                                time: 2000
                                ,icon:5
                            });
                        }

					});
                });
			} else {
                layer.confirm('确认要停用吗？', function(index) {
                	$.get(url,{'admid':id},function(res){
                		if(res['code']==1){
                            $(obj).attr('title', '启用')
                            $(obj).find('i').html('&#xe62f;');
                            $(obj).parents("tr").find(".td-status").find('span').addClass('layui-btn-disabled').html('已停用');
                            layer.msg(res['msg'], {
                                time: 1000
                                ,icon:6
                            });
                		}else{
                            layer.msg(res['msg'], {
                                time: 2000
                                ,icon:5
                            });
                        }

					});

                });
			}

	}

	/*用户-删除*/
	window.member_del = function (obj, id) {
		layer.confirm('确认要删除吗？', function(index) {
			//发异步删除数据
			$.get('del.html',{'admid':id},function (res) {
                if(res['code']==1){
                    layer.msg(res['msg'],{
                        time: 1000
                        ,icon: 6
                    },function(){
                        location.reload();
                    });
                }else{
                    layer.msg(res['msg'],{
                        time: 2000
                        ,icon: 5
                    });
                }

            })

		});
	}

	window.delAll = function (argument) {
		var data = tableCheck.getData();
		if(data == ''){
            layer.msg('请选择要删除的管理员',{
                anim: 6
            });
            return false;
        }
		layer.confirm('确认要删除吗？' + data, function(index) {
			//捉到所有被选中的，发异步进行删除
            //console.log(data);
            $.post('del.html',{'admid':data},function (res) {
                if(res['code'] == 1){
                    layer.msg(res['msg'],{
                        time: 1000
                        ,icon: 6
                    },function () {
                        location.reload();
                    });
                }else{
                    layer.msg(res['msg'],{
                        time: 1000
                        ,icon: 5
                    });
                }

            });

		});
	}
});