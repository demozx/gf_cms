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
                field: 'sort',title: '排序',templet:'<div style="line-height: 100px;"><input lay-verify="required" readonly autocomplete="off" class="layui-input sort" id="sort_{{ d.id }}" data-id="{{ d.id }}" type="number" value="{{ d.sort }}"></div>'
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
        //物理删除
		getCheckData: function() { //获取选中数据
			var checkStatus = table.checkStatus('articleList'),
				data = checkStatus.data;
			//console.log(data);
			//layer.alert(JSON.stringify(data));
            var ids = new Array();
			if(data.length > 0) {
                for(var i=0;i<data.length;i++){
                    ids.push(data[i]['id']);
                }
				layer.confirm('确认要永久删除吗？ID:(' + ids+')', function(index) {
				    $.post('del.html',{'aids':ids},function (res) {
                        if(res['code']==1){
                            //删除成功
                            layer.msg(res['msg'], {
                                icon: 6
                                ,time: 1000
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
				layer.msg("请先选择需要永久删除的文章！",{
				    anim: 6
                });
			}

		},
		//还原文章
        Restore: function() {
			var checkStatus = table.checkStatus('articleList'),
				data = checkStatus.data;
            if(data.length > 0){
                var ids = new Array();
                for(var i=0;i<data.length;i++){
                    ids.push(data[i]['id']);
                }
                layer.confirm('确认要还原文章吗？ID:(' + ids+')', function(index) {
                    $.post('restore.html',{'aids':ids},function (res) {
                        if(res['code']==1){
                            //还原成功
                            layer.msg(res['msg'], {
                                icon: 6
                                ,time: 2000
                            });
                            //重载当前页码的表格数据
                            //var nowPage = $('.layui-laypage-skip>.layui-input').val();
                            active.reload();
                        }else{
                            //还原失败
                            layer.msg(res['msg'], {
                                icon: 5
                                ,time: 2000
                            });
                        }
                    });
                });
                //console.log(ids);
            }else {
				layer.msg("请选择要还原的文章！",{
                    anim: 6
                });
				return false;
			}
		}

	};

    $('.we-search .layui-btn, .demoTable .layui-btn').on('click', function(){
        var type = $(this).data('type');
        active[type] ? active[type].call(this) : '';
    });
    /*单个删除*/
	window.member_del = function(obj, id) {
		layer.confirm('确认要永久删除吗？', function(index) {
			//发异步删除数据
            $.post('del.html',{'aids':id},function (res) {
                if(res['code']==1){
                    //删除成功
                    layer.msg(res['msg'], {
                        icon: 6,
                        time: 1000
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
    /*广告图片放大预览*/
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
