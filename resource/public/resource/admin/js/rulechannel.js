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
layui.use(['laydate', 'jquery', 'admin', 'form'], function() {
	var laydate = layui.laydate,
		$ = layui.jquery,
		form = layui.form,
		admin = layui.admin;

	//监听submit
    form.on('submit(add)', function(data){
        //console.log(data.elem) //被执行事件的元素DOM对象，一般为button对象
        //console.log(data.form) //被执行提交的form对象，一般在存在form标签时才会返回
        //console.log(data.field) //当前容器的全部表单字段，名值对形式：{name: value}
		if(data.field.channel_name==''){
			layer.msg('请输入分类名称',{icon:5,anim:6,time:2000});
		}else{
			$.get('add.html',data.field,function (res) {
				if(res['code']==1){
					//添加成功
					layer.msg(res['msg'],{icon:6,time:1000},function () {
						top.location.reload();
                    });
				}else{
					//添加失败
                    layer.msg(res['msg'],{icon:5,time:2000});
				}
            });
		}
        return false; //阻止表单跳转。如果需要表单跳转，去掉这段即可。
    });

	/*用户-删除*/
	window.member_del = function (obj, id) {
		layer.confirm('确认要删除吗？', function(index) {
			$.post('del.html',{'channel_ids':id},function (res) {
				if(res['code']==1){
                    layer.msg('删除成功!', {
                        icon: 6,
                        time: 1000
                    },function () {
						top.location.reload();
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
    /*删除全部*/
	window.delAll = function (argument) {
		var data = tableCheck.getData();
        if( data=='' ){
            layer.msg('请选择要删除的分组',{
                anim: 6
            });
            return false;
        }
		layer.confirm('确认要删除ID('+data+')吗？', function(index) {
			//console.log(data);
            $.post('del.html',{'channel_ids':data},function (res) {
                if(res['code']==1){
                    layer.msg('删除成功!', {
                        icon: 6,
                        time: 1000
                    },function () {
                        top.location.reload();
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
	/*排序*/
	window.sort = function(){
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
                        top.location.reload();
                    });
                }else{
                    layer.msg(res['msg'],{icon:5,time:2000});
                }
            },'json');
        }

    }
});