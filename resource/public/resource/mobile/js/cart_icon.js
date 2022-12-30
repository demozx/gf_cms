$(document).ready(function() {//panmove
    var tb = false;
    function panmove(){
        var pan_in=$(".site_nav");
        var scr_w=parseInt($(".allpage").width())-parseInt($(".site_btn").width());
        var scr_h=parseInt($(".allpage").height())-parseInt($(".site_btn").height());
        var half_w=parseInt(parseInt($(".allpage").width())/2);
        var half_h=parseInt(parseInt($(".allpage").height())/2);
        var isdrag=false;
        var tx,ty,x,y,n,m;
        function movemouse(e){
            if(isdrag){
                if((tx+e.touches[0].pageX-x)>scr_w){
                    n=scr_w;
                    tb = true;
                }else{
                    n=(tx+e.touches[0].pageX-x)>0?(tx+e.touches[0].pageX-x):0;
                    if(n<=half_w){ // 小于50%
                        pan_in.css("right","auto");
                        pan_in.css("left","100%");
                        tb = false;
                    }else{ // 大于50%
                        pan_in.css("left","auto");
                        pan_in.css("right","100%");
                        tb = true;
                    }
                }
                if((ty+e.touches[0].pageY-y)>scr_h){
                    m=scr_h;
                    pan_in.css("top","auto");
                    pan_in.css("bottom","100%");
                }else{
                    m=(ty+e.touches[0].pageY-y)>0?(ty+e.touches[0].pageY-y):0;
                    if(m<=half_h){ // 大于50% 屏幕底部
                        pan_in.css("top","0");
                        pan_in.css("bottom","auto");
                    }else{
                        pan_in.css("top","auto");
                        pan_in.css("bottom","100%");
                    }
                }
                $("#pbtn-nav").css("left",n);
                $("#pbtn-nav").css("top",m);
                e.preventDefault();
            }
            pan_in.removeClass("slideup");
            pan_in.removeClass("slidedown");
        }
        function selectmouse(e){
            isdrag=true;
            tx=parseInt($(".site_btn").css("left"));
            ty=parseInt($(".site_btn").css("top"));
            x=e.touches[0].pageX;
            y=e.touches[0].pageY;
            e.preventDefault();
        }
        document.getElementById("pbtn-nav").addEventListener('touchend',function(){isdrag=false;});
        document.getElementById("pbtn-nav").addEventListener('touchstart',selectmouse);
        document.getElementById("pbtn-nav").addEventListener('touchmove',movemouse);
    }
    //panmove();
    $("#pbtn-nav").click(function(){
        if(tb == true){
            $(".site-nav").toggleClass("slidedown");
        }else{
            $(".site-nav").toggleClass("slideup");
        }
    });
});