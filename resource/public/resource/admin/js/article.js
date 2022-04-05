layui.extend({
	admin: '{/}../../static/admin/js/admin'
});

layui.use(['table', 'jquery','form', 'admin', 'laydate'], function() {
	var table = layui.table,
		$ = layui.jquery,
		form = layui.form,
		admin = layui.admin,
    	laydate = layui.laydate;

    table.render({
        elem: '#articleList',
        cellMinWidth: 100,
        cellMinHeight: 100,
		url:location.href,
        cols: [
            [{
                type: 'checkbox',fixed:'left'
            }, {
                field: 'id',title: 'ID', width:70, sort: true,fixed:'left'
            }, {
                field: 'litpic',title: '缩略图',width:110,style:'height:100px;width:100px;', templet:'<div><img src="{{ d.litpic }}"></div>'
            }, {
                field: 'title',title: '标题',width:200,templet: '#usernameTpl'
            }, {
                field: 'createdate',title: '发布时间',width:164,sort: true
            }, {
                field: 'cat_name',title: '分类',sort: true
            },  {
                field: 'flag',title: '推荐',templet:'#recommendTpl',unresize: true
            }, {
                field: 'flag',title: '置顶',templet: '#topTpl',unresize: true
            }, {
                field: 'read',title: '审核',templet: '#reviewTpl',unresize: true
            }, {
                field: 'click',title: '点击',width:80,sort: true
            }, {
                field: 'sort',title: '排序',templet:'<div style="line-height: 100px;"><input lay-verify="required" autocomplete="off" class="layui-input sort" id="sort_{{ d.id }}" data-id="{{ d.id }}" type="number" value="{{ d.sort }}"></div>'
            }, {
                field: 'operate',title: '操作',toolbar: '#operateTpl',unresize: true,fixed:'right'
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
            var catid = $('#catid');
            var start = $('#start');
            var end = $('#end');
            var keyword = $('#keyword');
            //执行重载
            table.reload('articleList', {
                page: {
                    curr: nowPage //重新从第 N 页开始
                }
                , where: {
                    catid: catid.val(),
                    start: start.val(),
                    end: end.val(),
                    keyword: keyword.val(),
                },
            });

        },
        //添加文章
        add:function(){
            var cid = $('#catid').val();
            var str = ''
            //console.log(cid);
            if(cid){
                str = '?cid='+cid;
            }
            WeAdminShow('添加文章','add.html'+str,1000,500);
        },
        //批量删除
		getCheckData: function() { //获取选中数据
			var checkStatus = table.checkStatus('articleList'),
				data = checkStatus.data;
			// 判断物理删除和逻辑删除 start
			var delType = $('#delType').val();
			var delNotice = '';
			if(delType == 1){
                delNotice = '确认要删除吗？';
            }else{
                delNotice = '确认要移至回收站吗？';
            }
            // 判断物理删除和逻辑删除 end
			//console.log(data);
			//layer.alert(JSON.stringify(data));
            var ids = new Array();
			if(data.length > 0) {
                for(var i=0;i<data.length;i++){
                    ids.push(data[i]['id']);
                }
				layer.confirm(delNotice+'ID:(' + ids+')', function(index) {
				    $.post('del.html',{'aids':ids},function (res) {
                        if(res['code']==1){
                            //删除成功
                            layer.msg(res['msg'], {
                                icon: 6
                                ,time: 2000
                            });
                            //重载当前页码的表格数据
                            //var nowPage = $('.layui-laypage-skip>.layui-input').val();
                            active.reload();
                        }else{
                            //删除失败
                            layer.msg(res['msg'], {
                                icon: 5
                                ,time: 2000
                            });
                        }
                    });
				});
			} else {
				layer.msg("请先选择需要删除的文章！",{
				    anim: 6
                });
			}

		},
        //批量移动文章
        move:function(){
            var checkStatus = table.checkStatus('articleList'),
                data = checkStatus.data;
            if(data.length > 0){
                var aids = new Array();
                for(var i=0;i<data.length;i++){
                    aids.push(data[i]['id']);
                }
                //弹出页面
                WeAdminShow('移动文章','move.html?aids='+aids,400,200)
            }else{
                layer.msg("请选择要移动的文章",{
                    anim: 6
                });
                return false;
            }
        },
		//批量推荐
		Recommend: function() {
			var checkStatus = table.checkStatus('articleList'),
				data = checkStatus.data;
            if(data.length > 0){
                var ids = new Array();
                for(var i=0;i<data.length;i++){
                    ids.push(data[i]['id']);
                }
                //console.log(ids);
            }else {
				layer.msg("请选择要推荐的文章",{
                    anim: 6
                });
				return false;
			}
            $.post('flag.html',{'aids':ids,'flag':'r'},function (res) {
                //console.log(res);
                layer.msg(res['msg'],{
                    icon: 6
                    ,time: 1000
                });
                for(var i=0;i<ids.length;i++){
                    if(res['code']==1){
                        $('#flag_r_'+ids[i]).prop('checked',true);
                    }else if(res['code']==2){
                        $('#flag_r_'+ids[i]).prop('checked',false);
                    }

                }
                form.render();//重新渲染
            });

		},
		//批量置顶
		Top: function() {
            var checkStatus = table.checkStatus('articleList'),
                data = checkStatus.data;
            if(data.length > 0){
                var ids = new Array();
                for(var i=0;i<data.length;i++){
                    ids.push(data[i]['id']);
                }
                //console.log(ids);
            }else {
                layer.msg("请选择要置顶的文章",{
                    anim: 6
                });
                return false;
            }
            $.post('flag.html',{'aids':ids,'flag':'t'},function (res) {
                //console.log(res);
                layer.msg(res['msg'],{
                    icon: 6
                    ,time: 1000
                });
                for(var i=0;i<ids.length;i++){
                    if(res['code']==1){
                        $('#flag_t_'+ids[i]).prop('checked',true);
                    }else if(res['code']==2){
                        $('#flag_t_'+ids[i]).prop('checked',false);
                    }

                }
                form.render();//重新渲染
            });
		},
		//批量审核
		Review: function() {
            var checkStatus = table.checkStatus('articleList'),
                data = checkStatus.data;
            if(data.length > 0){
            	var ids = new Array();
				for(var i=0;i<data.length;i++){
					ids.push(data[i]['id']);
				}
			}else{
            	layer.msg('请选择要审核的文章',{
            	    anim: 6
                });
            	return false;
			}
			$.post('review.html',{'aids':ids},function (res) {
				//console.log(res);
				if(res['code']==1 || res['code']==2){
                    layer.msg(res['msg'],{
                        icon: 6
                    });
                }
                if(res['code']==0){
                    layer.msg(res['msg'],{
                        icon: 6
                        ,anim: 6
                    });
                }
				//返回res的code=0，审核失败或者已审核；code=1，审核成功；code=2，取消审核成功
				for(var i=0;i<ids.length;i++){
					if(res['code']==1){
                        $('#review_'+ids[i]).prop('checked',true);
					}else if(res['code']==2){
                        $('#review_'+ids[i]).prop('checked',false);
					}

				}
                form.render();//重新渲染
            });
		},

        /*排序*/
        Sort: function(){
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
                        //top.location.reload();
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

    //单个审核
    form.on('checkbox(review)', function(data){
        //console.log(data.value); //复选框value值，也可以通过data.elem.value得到
        $.post('review.html',{'aids':data.value},function (res) {
            //console.log(res);
            if(res['code']==1 || res['code']==2){
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
        });
    });
	//单个修改flag
    form.on('switch(flag)', function(data){
        //console.log(data.elem.className);
		var flag = '';
		if(data.elem.className == 'flag_r'){
			//推荐
			flag = 'r';
		}else if(data.elem.className == 'flag_t'){
			//置顶
			flag = 't'
		}
        $.post('flag.html',{'aids':data.value,'flag':flag},function (res) {
            //console.log(res);
            if(res['code']==1 || res['code']==2){
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
        });

    });
    /*单个删除*/
	window.member_del = function(obj, id) {
        // 判断物理删除和逻辑删除 start
        var delType = $('#delType').val();
        var delNotice = '';
        if(delType == 1){
            delNotice = '确认要删除吗？';
        }else{
            delNotice = '确认要移至回收站吗？';
        }
        // 判断物理删除和逻辑删除 end
		layer.confirm(delNotice, function(index) {
			//发异步删除数据
            $.post('del.html',{'aids':id},function (res) {
                if(res['code']==1){
                    //删除成功
                    layer.msg(res['msg'], {
                        icon: 6,
                        time: 2000
                    });
                    //重载当前页码的表格数据
                    var nowPage = $('.layui-laypage-skip>.layui-input').val();
                    active.reload(nowPage);
                }else{
                    //删除失败
                    layer.msg(res['msg'], {
                        icon: 5,
                        time: 1000
                    });
                }
            });

		});
	}
    //文章搜索开始时间选择器
    laydate.render({
        elem: '#start'
        ,type: 'datetime'
    });
    //文章搜索结束时间选择器
    laydate.render({
        elem: '#end'
        ,type: 'datetime'
    });
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
	
	//文章预览
    $('body').on('click','.preView',function () {
        var artid = $(this).attr('data-id');
        $.get('preView.html',{'artid':artid},function (res) {
            layer.open({
                type: 2,
                title: '文章预览',
                content: res,
                area: ['96%', '96%'],
                maxmin: true
            })
        });
    });

});
