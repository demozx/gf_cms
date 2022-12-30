$(function(){

	var swiper = new Swiper('.swiper-container', {
        pagination: '.swiper-pagination',
        paginationClickable: true,
        loop:true,
       /* autoplay:3000*/
    });

    $(".bangdan_c li:odd").css('background-color', '#e1e1e1');
    $(".case_i_g ul li:last").css('border', '0');
    $(".search_hl").click(function(event) {
    	var search =$(".search");
    	search.animate({top:0}, 250)
    });
    $(".xbtn").click(function(event) {
    	var search =$(".search")
    	search.animate({top:"-120%"}, 250)
    });

    $(".product_c li:odd,.zizhi_con:odd").css('margin-left', '10%');

})