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
        elem: '#adsList',
        cellMinWidth: 60,
        cellMinHeight: 100,
        url: location.href,
        cols: [
            [{
                type: 'checkbox'
            }, {
                field: 'id', title: 'ID', width: 60, sort: true
            }, {
                field: 'img_url',
                width: 110,
                title: '广告图片',
                style: 'height:100px;width:100px;',
                templet: '<div><img src="{{ d.img_url }}"></div>'
            }, {
                field: 'name', title: '广告名称'
            }, {
                field: 'channel_name', title: '分类'
            }, {
                field: 'link', title: '链接地址'
            }, {
                field: 'status', title: '状态', width: 80, templet: '#status'
            }, {
                field: 'start_time', title: '开始时间', width: 170
            }, {
                field: 'end_time', title: '结束时间', width: 170
            }, {
                field: 'sort',
                title: '排序',
                templet: '<div style="line-height: 100px;"><input lay-verify="required" autocomplete="off" class="layui-input sort" id="sort_{{ d.id }}" data-id="{{ d.id }}" type="number" value="{{ d.sort }}"></div>'
            }, {
                field: 'remarks', title: '备注'
            }, {
                field: 'operate', title: '操作', toolbar: '#operateTpl'
            }]
        ],

        event: true,
        page: {
            layout: ['prev', 'next', 'page', 'count', 'skip', 'limit'] //自定义分页布局
            //,curr: 5 //设定初始在第 5 页
            , groups: 5 //只显示 N 个连续页码
            , first: false //不显示首页
            , last: false //不显示尾页

        },
    });
    //console.table(res);

    /*
     *数据表格中form表单元素是动态插入,所以需要更新渲染下
     * http://www.layui.com/doc/modules/form.html#render
     * */
    $(function () {
        //form.render();
    });

    var active = {
        //表格内容重载
        reload: function (nowPage) {
            var nowPage = nowPage || 1;
            var catid = $('#catid');
            //执行重载
            table.reload('adsList', {
                page: {
                    curr: nowPage //重新从第 N 页开始
                }
                , where: {
                    channel_id: catid.val(),
                },
            });

        },
        //批量删除
        Del: function () { //获取选中数据
            var checkStatus = table.checkStatus('adsList'),
                data = checkStatus.data;
            //console.log(data);
            //layer.alert(JSON.stringify(data));
            var ids = new Array();
            if (data.length > 0) {
                for (var i = 0; i < data.length; i++) {
                    ids.push(data[i]['id']);
                }
                layer.confirm('确认要删除吗？ID:(' + ids + ')', function (index) {
                    $.post('del.html', {'adids': ids}, function (res) {
                        if (res['code'] == 1) {
                            //删除成功
                            layer.msg(res['msg'], {
                                icon: 6
                                , time: 1000
                            });
                            //重载当前页码的表格数据
                            //var nowPage = $('.layui-laypage-skip>.layui-input').val();
                            active.reload();
                        } else {
                            //删除失败
                            layer.msg(res['msg'], {
                                icon: 5
                                , time: 2000
                            });
                        }
                    });
                });
            } else {
                layer.msg("请先选择需要删除的广告！", {
                    anim: 6
                });
            }

        },
        //批量启用
        Open: function () {
            var checkStatus = table.checkStatus('adsList'),
                data = checkStatus.data;
            if (data.length > 0) {
                var ids = new Array();
                for (var i = 0; i < data.length; i++) {
                    ids.push(data[i]['id']);
                }
                //console.log(ids);
            } else {
                layer.msg("请选择要启用的广告", {
                    anim: 6
                });
                return false;
            }
            $.post('editstatus.html', {'adids': ids, 'status': 'on'}, function (res) {
                //console.log(res);
                var icon = 6;//笑脸
                var anim = '';//抖动动画初始变量
                if (res['code'] == 0) {
                    icon = 5;//哭脸
                }
                if (res['code'] == 2) {
                    anim = 6;
                }
                layer.msg(res['msg'], {
                    icon: icon
                    , anim: anim
                    , time: 1000
                });
                //重载当前页码的表格数据
                var nowPage = $('.layui-laypage-skip>.layui-input').val();
                active.reload(nowPage);
            });

        },
        //批量停用
        Close: function () {
            var checkStatus = table.checkStatus('adsList'),
                data = checkStatus.data;
            if (data.length > 0) {
                var ids = new Array();
                for (var i = 0; i < data.length; i++) {
                    ids.push(data[i]['id']);
                }
                //console.log(ids);
            } else {
                layer.msg("请选择要停用的广告", {
                    anim: 6
                });
                return false;
            }
            $.post('editstatus.html', {'adids': ids, 'status': 'off'}, function (res) {
                //console.log(res);
                var icon = 6;//笑脸
                var anim = '';//抖动动画初始变量
                if (res['code'] == 0) {
                    icon = 5;//哭脸
                }
                if (res['code'] == 2) {
                    anim = 6;
                }
                layer.msg(res['msg'], {
                    icon: icon
                    , anim: anim
                    , time: 1000
                });
                //重载当前页码的表格数据
                var nowPage = $('.layui-laypage-skip>.layui-input').val();
                active.reload(nowPage);
            });

        },
        /*排序*/
        Sort: function () {
            var arr = new Array();
            var work = false;
            $('.sort').each(function () {
                if ($(this).val() == '') {
                    layer.tips('不能为空', '#sort_' + $(this).attr('data-id'), {
                        tips: [4, '#78BA32']
                        , tipsMore: false
                    });
                    work = false;
                    return false;
                } else if ($(this).val() < 0) {
                    layer.tips('必须为大于等于0的数字', '#sort_' + $(this).attr('data-id'), {
                        tips: [4, '#78BA32']
                        , tipsMore: false
                    });
                    work = false;
                    return false;
                } else {
                    arr.push({'id': $(this).attr('data-id'), 'sort': $(this).val()});
                    work = true;
                }
            });
            if (work == true) {
                $.post('sort.html', {'data': arr}, function (res) {
                    if (res['code'] == 1) {
                        layer.msg('更新排序成功', {icon: 6, time: 1000}, function () {
                            //top.location.reload();
                        });
                    } else {
                        layer.msg(res['msg'], {icon: 5, time: 2000});
                    }
                }, 'json');
            }

        }

    };

    $('.we-search .layui-btn, .demoTable .layui-btn').on('click', function () {
        var type = $(this).data('type');
        active[type] ? active[type].call(this) : '';
    });

    //监听select
    form.on('select(adschannel)', function (data) {
        //console.log(data.elem); //得到select原始DOM对象
        //console.log(data.value); //得到被选中的值
        //console.log(data.othis); //得到美化后的DOM对象
        active.reload();
    });

    /*单个删除*/
    window.ad_del = function (obj, id) {
        layer.confirm('确认要删除吗？', function (index) {
            //发异步删除数据
            $.post('del.html', {'adids': id}, function (res) {
                if (res['code'] == 1) {
                    //删除成功
                    layer.msg(res['msg'], {
                        icon: 6,
                        time: 1000
                    });
                    //重载当前页码的表格数据
                    var nowPage = $('.layui-laypage-skip>.layui-input').val();
                    active.reload(nowPage);
                } else {
                    //删除失败
                    layer.msg(res['msg'], {
                        icon: 5,
                        time: 1000
                    });
                }
            });

        });
    }

    /*广告图片放大预览*/
    $('body').on('mouseenter', '.layui-table-cell>img', function () {
        //console.log('显示大图');
        var index = layer.open({
            type: 1,
            title: false,
            closeBtn: 0,
            area: '100',
            skin: 'layui-layer-nobg', //没有背景色
            shade: false,
            shadeClose: true,
            content: "<img style='max-width:600px;' src='" + $(this).prop('src') + "'>"
        });
    });
    $('body').on('mouseleave', '.layui-table-cell>img', function () {
        //console.log('取消显示大图');
        layer.close(layer.index);
    });

});
