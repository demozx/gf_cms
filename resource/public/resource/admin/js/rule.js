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
    //admin: '{/}../../static/admin/js/admin',
    admin: '/static/admin/js/admin',
});
layui.use(['laydate', 'jquery', 'form', 'admin'], function() {
	var laydate = layui.laydate,
		form = layui.form,
		$ = layui.jquery,
        admin = layui.admin;
	//当前模块地址
	var _controller = $('#_controller_').val();
	//console.log(_controller);
	//执行一个laydate实例
	laydate.render({
		elem: '#start' //指定元素
	});
	//执行一个laydate实例
	laydate.render({
		elem: '#end' //指定元素
	});
    //监听控制器select
    form.on('select(controller)', function(data){
        $('#action').find('option').not('option:first').remove();
        var v = data.value;
        if(v=='0'){
            $('#action').prop('disabled',true);
            form.render();
        	return;
		}
        var index = layer.load();
		$.get(_controller+'getactions.html',{'controller':data.value},function(res){
			//console.log(res);
			if(res['code']==0){
                layer.msg(res['msg'],{
                    icon: 5
                });
			}else{
                $('#action').prop('disabled',false);
                var html = '';
                for(var i=0;i<res.length;i++){
                    html += "<option value='"+res[i]+"'>"+res[i]+"</option>";
                }
                $('#action').find('option').eq(0).after(html);
            }
            form.render();
            layer.close(index);
        },'json');
    });
    //监听菜单显示
    form.on('switch(show_channel)',function (data) {
        //console.log($(this).attr('data-id'));
        $.post(_controller+'showchannel.html',{'ruleid':$(this).attr('data-id')},function (res) {
            if(res['code']==1){
                layer.msg(res['msg'],{
                    icon:6,
                    time:1000
                },function () {
                    top.location.reload();//刷新最顶端对象（用于多开窗口）
                });
            }else{
                layer.msg(res['msg'],{
                    icon:5,
                    time:2000
                });
            }
        });
    });
    /*启用-停用*/
    window.member_stop = function (obj, id) {
        var url = _controller+'status.html';
        if($(obj).prop('title') == '启用') {
            layer.confirm('确认要启用吗？', function(index) {
                //发异步把用户状态进行更改
                $.get(url,{'ruleid':id},function(res){
                    //console.log(res);
                    if(res['code']==1){
                        $(obj).attr('title', '停用')
                        $(obj).find('i').html('&#xe601;');
                        $(obj).parents("tr").find(".td-status").find('span').removeClass('layui-btn-disabled').html('已启用');
                    }
                    layer.msg(res['msg'], {
                        time: 1000
                        ,icon:6
                    });
                });
            });
        } else {
            layer.confirm('确认要停用吗？', function(index) {
                $.get(url,{'ruleid':id},function(res){
                    if(res['code']==1){
                        $(obj).attr('title', '启用')
                        $(obj).find('i').html('&#xe62f;');
                        $(obj).parents("tr").find(".td-status").find('span').addClass('layui-btn-disabled').html('已停用');
                    }
                    layer.msg(res['msg'], {
                        time: 1000
                        ,icon:6
                    });
                });

            });
        }

    }
    //监听提交
    form.on('submit(add)', function(data){
    	//console.log(data.field);
    	if(data.field['channel_id']==0){
			layer.tips('请选择规则分类',$('#channel_id').next(),{
                tips: [3, '#78BA32']
                ,time: 1000
            });
			return false;
		}
        if(data.field['controller']==0){
            layer.tips('请选择控制器名',$('#controller').next(),{
                tips: [3, '#78BA32']
                ,time: 1000
            });
            return false;
        }
        if(data.field['action']==0){
            layer.tips('请选择方法名',$('#action').next(),{
                tips: [3, '#78BA32']
                ,time: 1000
            });
            return false;
        }
        if(data.field['title']==''){
            layer.tips('请填写权限名',$('#rule_name'),{
                tips: [3, '#78BA32']
                ,time: 1000
            });
            return false;
        }
    	$.post(_controller+'add.html',data.field,function (res) {
			if(res['code']==1){
				layer.msg('添加成功,请手动刷新页面',{icon:6,time:2000},function () {
                    //top.location.reload();//刷新最顶端对象（用于多开窗口）
                    //由于添加权限的时候自动刷新会影响操作效率，所以改成提示手动刷新
                });
			}else{
                layer.msg(res['msg'],{icon:5,time:3000});
			}
        });
        return false; //阻止表单跳转。如果需要表单跳转，去掉这段即可。
    });
    //监听分类筛选
    form.on('select(channel)', function(data){
        //console.log(data.elem[data.elem.selectedIndex].getAttribute('_href')); //得到被选中的_href
        var href = data.elem[data.elem.selectedIndex].getAttribute('_href');
        location.href = href;
    });
    //监听方法名，判断控制器和方法名组合是否存在
    form.on('select(action)', function(data){
        //console.log('控制器名：'+$('#controller').val());
        //console.log('方法名：'+data.value);
        var controller = $('#controller').val();
        var action = data.value;
        //控制器名或方法名为0时
        if(controller==0 || action==0){
            return false;
        }
        var index = layer.load();
        $.post(_controller+'findrule.html',{'controller':controller,'action':action},function (res) {
            //console.log(res);
            layer.close(index);
            if(res['code']==1){
                //规则存在
                layer.msg('规则：'+res['msg']['rule']+'<br>规则名称：'+res['msg']['name']+'<br>所属分类：'+res['msg']['channel']+'<br>已经存在',{
                    time: 8000,
                    icon: 5
                });
            }else{
                //规则不存在
                layer.msg('规则：Admin/'+controller+'/'+action+'<br>检测通过',{
                    icon: 6
                });
            }
        });
    });
    /*删除*/
	window.member_del = function (obj, id) {
		layer.confirm('确认要删除ID:'+id+'吗？', function(index) {
			$.post(_controller+'del.html',{'ruleids':id},function (res) {
				//console.log(res);
				if(res['code']==1){
                    layer.msg('删除成功', {
                        icon: 6,
						time:1000
                    },function () {
                        top.location.reload();//刷新最顶端对象（用于多开窗口）
                    });
				}else{
                    layer.msg(res['msg'], {
                        icon: 5,
                        time:2000
                    });
				}
            });
		});
	}
    /*批量删除*/
	window.delAll = function (argument) {
		var data = tableCheck.getData();
        if( data=='' ){
            layer.msg('请选择要删除的权限',{
                anim: 6
            });
            return false;
        }
		layer.confirm('确认要删除ID('+data+')吗？', function(index) {
            $.post(_controller+'del.html',{'ruleids':data},function (res) {
                //console.log(res);
                if(res['code']==1){
                    layer.msg('删除成功', {
                        icon: 6,
                        time:1000
                    },function () {
                        top.location.reload();//刷新最顶端对象（用于多开窗口）
                    });
                }else{
                    layer.msg(res['msg'], {
                        icon: 5,
                        time:2000
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
        //console.log(arr);
        if(work == true){
            $.post(_controller+'sort.html',{'data':arr},function (res) {
                if(res['code']==1){
                    layer.msg('更新排序成功',{icon:6,time:1000},function () {
                        top.location.reload();//刷新最顶端对象（用于多开窗口）
                    });
                }else{
                    layer.msg(res['msg'],{icon:5,time:2000});
                }
            },'json');
        }

    }
});