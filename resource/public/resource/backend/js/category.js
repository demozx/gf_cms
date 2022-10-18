
/*分类-停用*/
function member_stop(obj, id) {
    var confirmTip;
    var $ = layui.jquery;
    if($(obj).attr('title') == '启用') {
        confirmTip = '确认要停用吗？';
    } else {
        confirmTip = '确认要启用吗？';
    }
    layer.confirm(confirmTip, function(index) {
        if($(obj).attr('title') == '启用') {
            //发异步把用户状态进行更改
			$.post('editstate.html',{'id':id,'state':'off'},function(res){
				if(res['code']==1){
                    $(obj).attr('title', '停用')
                    $(obj).find('i').html('&#xe62f;');
                    $(obj).parents("tr").find(".td-status").find('span').addClass('layui-btn-disabled').html('已停用');
                    layer.msg('已停用!', {
                        icon: 6,
                        time: 1000
                    });
				}else{
					//操作失败
                    layer.msg(res['msg'], {
                        icon: 5,
                        time: 1000
                    });
				}
			});

        } else {
            //发异步把用户状态进行更改
		$.post('editstate.html',{'id':id,'state':'on'},function(res){
			if(res['code']==1){
                $(obj).attr('title', '启用')
                $(obj).find('i').html('&#xe601;');

                $(obj).parents("tr").find(".td-status").find('span').removeClass('layui-btn-disabled').html('已启用');
                layer.msg('已启用!', {
                    icon: 6,
                    time: 1000
                });
			}else{
                //操作失败
                layer.msg(res['msg'], {
                    icon: 5,
                    time: 1000
                });
			}
       	 });
		}
    });
}
//自定义的render渲染输出多列表格
var layout = [
	{
		name: '栏目名称',
		treeNodes: true,
		headerClass: 'value_col',
		colClass: 'value_col',
		style: 'width: 60%'
	},
	{
		name: '状态',
		headerClass: 'td-status',
		colClass: 'td-status',
		style: 'width: 10%',
		render: function(row) {
			//启用禁用状态值
			//console.log('----'+row.state);
			if(row.state==1){
                return '<span class="layui-btn layui-btn-normal layui-btn-xs">已启用</span>';
			}
			return '<span class="layui-btn layui-btn-normal layui-btn-xs layui-btn-disabled">已停用</span>';
		}
	},
	{
		name: '操作',
		headerClass: 'td-manage',
		colClass: 'td-manage',
		style: 'width: 20%',
		render: function(row) {
            if(row.state==1){
                return '<a onclick="member_stop(this,' + row.id + ')" href="javascript:;" title="启用"><i class="layui-icon">&#xe601;</i></a>' +
                    '<a title="添加子类" onclick="WeAdminEdit(\'添加\',\'../category/add.html?cid=' + row.id + '\')" href="javascript:;"><i class="layui-icon">&#xe654;</i></a>' +
                    '<a title="编辑" onclick="WeAdminEdit(\'编辑\',\'../category/edit.html?cid='+ row.id +'\')" href="javascript:;"><i class="layui-icon">&#xe642;</i></a>' +
                    '<a title="删除" onclick="del(' + row.id + ')" href="javascript:;">\<i class="layui-icon">&#xe640;</i></a>';
            }
			return '<a onclick="member_stop(this,' + row.id + ')" href="javascript:;" title="停用"><i class="layui-icon">&#xe62f;</i></a>' +
				'<a title="添加子类" onclick="WeAdminEdit(\'添加\',\'../category/add.html?cid=' + row.id + '\')" href="javascript:;"><i class="layui-icon">&#xe654;</i></a>' +
				'<a title="编辑" onclick="WeAdminEdit(\'编辑\',\'../category/edit.html?cid='+ row.id +'\')" href="javascript:;"><i class="layui-icon">&#xe642;</i></a>' +
				'<a title="删除" onclick="del(' + row.id + ')" href="javascript:;">\<i class="layui-icon">&#xe640;</i></a>';
			//return '<a class="layui-btn layui-btn-danger layui-btn-mini" onclick="del(' + row.id + ')"><i class="layui-icon">&#xe640;</i> 删除</a>'; //列渲染
		}
	},
];
//加载扩展模块 treeGird
//		layui.config({
//			  base: './static/js/'
//			  ,version: '101100'
//			}).use('backend');
layui.extend({
	admin: '{/}../../static/backend/js/backend',
	treeGird: '{/}../../static/lib/layui/lay/treeGird' // {/}的意思即代表采用自有路径，即不跟随 base 路径
});
layui.use(['treeGird', 'jquery', 'admin', 'layer'], function() {
	var layer = layui.layer,
		$ = layui.jquery,
		admin = layui.admin,
		treeGird = layui.treeGird;

	$.get('#',function (res) {
		//console.log(res);
        var tree1 = layui.treeGird({
            elem: '#demo', //传入元素选择器
            spreadable: true, //设置是否全展开，默认不展开
            nodes:res,
			/*
			 nodes: [{
			 "id": "1",
			 "name": "父节点1",
			 "children": [{
			 "id": "11",
			 "name": "子节点11"
			 },
			 {
			 "id": "12",
			 "name": "子节点12"
			 }
			 ]
			 },
			 {
			 "id": "2",
			 "name": "父节点2",
			 "children": [{
			 "id": "21",
			 "name": "子节点21",
			 "children": [{
			 "id": "211",
			 "name": "子节点211"
			 }]
			 }]
			 }
			 ],
			 */
            layout: layout
        });

        $('#collapse').on('click', function() {
            layui.collapse(tree1);
        });

        $('#expand').on('click', function() {
            layui.expand(tree1);
        });
    },'json');

    //点击栏目名称
    $('body').on('click','cite',function () {
        //获取栏目id
        var cateId = $(this).parents('tr').attr('id');
        var title = $(this).text()
        //console.log(cateId);
        //self.location.href = '../article/index.html?catid='+cateId;
        //WeAdminEdit(title,'../article/index.html?catid='+cateId);
        var index = layer.open({
            type: 2,
            content: '../article/index.html?catid='+cateId,
            area: ['90%', '90%'],
            maxmin: true
        });
        layer.full(index);

        var url = $(this).children('a').attr('_href');
        var title = $(this).find('cite').html();
        var index = $('.left-nav #nav li').index($(this));

        for(var i = 0; i < $('.x-iframe').length; i++) {
            if($('.x-iframe').eq(i).attr('tab-id') == index + 1) {
                tab.tabChange(index + 1);
                event.stopPropagation();
                return;
            }
        };

        tab.tabAdd(title, url, index + 1);
        tab.tabChange(index + 1);


    });

});
//删除栏目
function del(Id) {
	var $ = layui.$;
    layer.confirm('确定要删除栏目？', {
            btn: ['确定','取消'] //按钮
        },function () {
		delpost();
    });
       function delpost() {
           $.post('del.html',{'id':Id},function (res) {
               console.log(res);
               if(res['code']==1){
                   layer.msg(res['msg'], {
                       icon: 6,
                       time: 1000
                   });
                   //删除元素节点
                   $('#cate_id_'+Id).parents('tr').remove();
                   //刷新栏目数量
                   reloadCatcount();
               }else{
                   layer.msg(res['msg'], {
                       icon: 5,
                       time: 2000
                   });
               }
           });
       }
}
//刷新栏目数量方法
function reloadCatcount() {
    layui.use(['jquery'], function(){
        var $ = layui.$;
        //获取栏目总数量
        $.get('countcates.html?'+Math.random(),function (res) {
            //console.log(res);
            $('#count_cate').text(res);
        });
    });
}
