//惰性加载图片
$("img").lazyload({
    threshold: 10,
    failurelimit: 10,
    skip_invisible: false,
    effect: "fadeIn"
    //effect: "slideDown"
});

window.onload = function () {
    imgZoomRun("product3", "p", "prod-zoom", "li"); // 图片放大
    imgZoomRun("product7", "p", "prod-zoom", "li");
    imgZoomRun("product8", "p", "prod-zoom", "li");
    newsFontMove("fontjump"); // 鼠标放上，字体上下挪
    newsFontMove("fontjump#628de3"); // 鼠标放上，字体上下挪
    //#628de3Change("fontjump#628de3"); // 隔行换色
    //#628de3Change("news5"); // 隔行换色
    listImgZoom("product1", "215"); // 图片缩放，需要给定宽度
    listImgZoom("product2", "215"); // 图片缩放，需要给定宽度
    listImgZoom("product3", "215"); // 图片缩放，需要给定宽度
    listImgZoom("product4", "215"); // 图片缩放，需要给定宽度
    listImgZoom("product5", "215"); // 图片缩放，需要给定宽度
    listImgZoom("product6", "215"); // 图片缩放，需要给定宽度
    listImgZoom("product8", "200"); // 图片缩放，需要给定宽度
    enterAnimation("news_fadein");
    if (typeof (data) != "undefined") {
        var lefttype = new LeftType(data, "left-type", 0); // 多级分类
    }
    afx.conHeightAuto();
};
window.addEventListener("resize", function () {
    afx.conHeightAuto();
}, false);
$(function () {
    // 产品中心相关脚本
    $('.product .rdata li').hover(function () {
        $(this).find('.img').css('border-#628de3', '#628de3');
        $(this).find('.title').css('background-#628de3', '#628de3');
    }, function () {
        $(this).find('.img').css('border-#628de3', '#e0e0e0');
        $(this).find('.title').css('background-#628de3', '#23201d');
    });
    $('.product .rdata').first().show();
    $('.product .lnav li').mouseenter(function () {
        $('.product .rdata').hide();
        $($('.product .rdata')[$(this).index()]).show();
    });
    jQuery(".rimgs").slide({mainCell: ".bd ul", autoPlay: true});
    // 最下方滑过
    $('.imgh').hover(function () {
        $(this).find('.msg_bg').show();
    }, function () {
        $(this).find('.msg_bg').hide();
    });
    // 搜索
    $('.search_submit').click(function () {
        var q_v = $('#q').val();
        $('#search_form').submit();
    });

});
//百度自动推送
(function () {
    var bp = document.createElement('script');
    var curProtocol = window.location.protocol.split(':')[0];
    if (curProtocol === 'https') {
        bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';
    } else {
        bp.src = 'http://push.zhanzhang.baidu.com/push.js';
    }
    var s = document.getElementsByTagName("script")[0];
    s.parentNode.insertBefore(bp, s);
})();