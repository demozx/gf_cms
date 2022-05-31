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
layui.use(['laydate', 'jquery', 'admin', 'form'], function() {
	var laydate = layui.laydate,
		$ = layui.jquery,
		form = layui.form,
		admin = layui.admin;

    /*修改启用状态*/
    form.on('switch(status)', function(data){
        //console.log(data.elem); //得到checkbox原始DOM对象
        //console.log(data.elem.checked); //开关是否开启，true或者false
        //console.log(data.value); //开关value值，也可以通过data.elem.value得到
        //console.log(data.othis); //得到美化后的DOM对象
        $.post('read.html',{'id':data.value},function (res) {
            if(res['code']==1){
                //已读成功
                layer.msg(res['msg'],{icon:6,time:1000},function () {
                    //self.location.reload();
                });
            }else if(res['code']==2){
                //未读成功
                layer.msg(res['msg'],{icon:6,time:1000},function () {
                    //self.location.reload();
                });
            }else{
                //操作失败
                layer.msg(res['msg'],{icon:5,time:2000});
            }
        });
    });
    /*删除全部*/
	window.delAll = function () {
		var data = tableCheck.getData();
        if( data=='' ){
            layer.msg('请选择要删除的分组',{
                anim: 6
            });
            return false;
        }
		layer.confirm('确认要删除ID('+data+')吗？', function(index) {
            $.post('del.html',{'ids':data},function (res) {
                if(res['code']==1){
                    layer.msg('删除成功!', {
                        icon: 6,
                        time: 1000
                    },function () {
                        self.location.reload();
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
                        self.location.reload();
                    });
                }else{
                    layer.msg(res['msg'],{icon:5,time:2000});
                }
            },'json');
        }

    }

});