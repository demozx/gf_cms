layui.extend({
	admin: '{/}../../static/admin/js/admin'
});

layui.use(['table', 'jquery','form', 'admin'], function() {
	var table = layui.table,
		$ = layui.jquery,
		form = layui.form,
		admin = layui.admin;

    table.render({
        elem: '#smsLogsList',
        cellMinWidth: 60,
        cellMinHeight: 100,
		url:location.href,
        cols: [
            [{
                field: 'id',title: 'ID',width:60, sort: true
            }, {
                field: 'dep_name',title: '部门名称'
            }, {
                field: 'emp_name',title: '员工姓名'
            }, {
                field: 'work_id',title: '员工工号'
            }, {
                field: 'sms_content',title: '短信内容'
            }, {
                field: 'send_to',title: '接收手机号'
            }, {
                field: 'RequestId',title: '请求id'
            }, {
                field: 'Code',title: '返回状态码'
            }, {
                field: 'Message',title: '返回信息'
            }, {
                field: 'smsStatus',title: '发送情况'
            }, {
                field: 'send_time',title: '发送时间',width:170
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

});
