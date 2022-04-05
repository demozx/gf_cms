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
    admin: '{/}../../static/admin/js/admin',
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

			if($(obj).prop('title') == '启用') {
                layer.confirm('确认要启用吗？', function(index) {
                	$.get('status.html',{'roleid':id},function (res) {
						if(res['code']==1){
							//启用成功
                            $(obj).attr('title', '停用')
                            $(obj).find('i').html('&#xe601;');

                            $(obj).parents("tr").find(".td-status").find('span').removeClass('layui-btn-disabled').html('已启用');
                            layer.msg('已启用', {
                                icon: 6,
                                time: 1000
                            });
						}else{
                            layer.msg(res['msg'], {
                                icon: 5,
                                time: 2000
                            });
						}
                    });

                });
			} else {
                layer.confirm('确认要停用吗？', function(index) {
                    $.get('status.html',{'roleid':id},function (res) {
						if(res['code']==1){
							//停用成功
                            $(obj).attr('title', '启用')
                            $(obj).find('i').html('&#xe62f;');

                            $(obj).parents("tr").find(".td-status").find('span').addClass('layui-btn-disabled').html('已停用');
                            layer.msg('已停用!', {
                                icon: 6,
                                time: 1000
                            });
						}else{
                            layer.msg(res['msg'], {
                                icon: 5,
                                time: 2000
                            });
						}
					});

                });
			}

	}

	/*角色-删除*/
	window.member_del = function (obj, id) {
		layer.confirm('确认要删除吗？', function(index) {
			$.post('del.html',{'roleids':id},function (res) {
				if(res['code']==1){
					//删除成功
                    layer.msg(res['msg'], {
                        icon: 6,
                        time: 1000
                    });
                    $(obj).parents("tr").remove();
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

	window.delAll = function (argument) {
		var data = tableCheck.getData();
		if( data=='' ){
            layer.msg('请选择要删除的角色',{
                anim: 6
            });
            return false;
        }
		layer.confirm('确认要删除ID(' + data + ')的角色吗？', function(index) {
			$.post('del.html',{'roleids':data},function (res) {
				if(res['code']==1){
                    layer.msg(res['msg'], {
                        icon: 6
                    });
                    $(".layui-form-checked").not('.header').parents('tr').remove();
				}else{
                    layer.msg(res['msg'], {
                        icon: 5
						,time:4000
                    });

				}
            });

		});
	}
});