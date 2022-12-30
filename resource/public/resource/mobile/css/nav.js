(function (doc, win) {
    var docEl = doc.documentElement,
        resizeEvt = "orientationchange" in window ? "orientationchange" : "resize",
        recalc = function () {
            var clientWidth = docEl.clientWidth;
            //var foot = document.getElementById("foot");
            if (!clientWidth) return;
            if (clientWidth<640){
                docEl.style.fontSize = 120 * (clientWidth / 640) + "px";
                console.log(120 * (clientWidth / 640) + "px");
            }else{
                docEl.style.fontSize = "120px";
            }
        };

    if (!doc.addEventListener) return;
    win.addEventListener(resizeEvt, recalc, false);
    doc.addEventListener('DOMContentLoaded', recalc, false);
})(document, window);

$(document).ready(function() {

    $(".common_news").click(function(event) {
        //图片加载完获取高度
        var t_img1; // 定时器
        var isLoad = false; // 控制变量
         
        // 判断图片加载状况，加载完成后回调
        isImgLoad(function(){
            
            var ccliH = $(".common_news li a img").outerHeight();
            var spanH = $(".common_news li a span").outerHeight()*1.3;

            var maxHH =0;
            $(".common_news li a img").each(function(index, el) {
                if(maxHH < $(this).height()){
                    maxHH = $(this).height()
                    return maxHH;
                }
            });

            var myHeight = Math.floor(maxHH+spanH);
            $(".common_news li").height(myHeight);
           
        });
         
        // 判断图片加载的函数
        function isImgLoad(callback){

            // 注意我的图片类名都是cover，因为我只需要处理cover。其它图片可以不管。
            // 查找所有封面图，迭代处理
            $(".common_news li img").each(function(){
                // 找到为0就将isLoad设为false，并退出each
                if(this.height === 0){
                    isLoad = false;
                    return false;
                }
            });
            // 为true，没有发现为0的。加载完毕
            if(isLoad){

                clearTimeout(t_img1); // 清除定时器
                // 回调函数
                callback();
            // 为false，因为找到了没有加载完成的图，将调用定时器递归
            }else{
                console.log(2)
                isLoad = true;
                t_img1 = setTimeout(function(){
                    isImgLoad(callback); // 递归扫描
                },100); // 我这里设置的是500毫秒就扫描一次，可以自己调整
            }
        }

    //图片加载完获取高度 end
    });




//图片加载完获取高度

    var t_img1; // 定时器
    var isLoad = false; // 控制变量
     
    // 判断图片加载状况，加载完成后回调
    isImgLoad(function(){
        
        var ccliH = $(".common_news li a img").outerHeight();
        var spanH = $(".common_news li a span").outerHeight()*1.3;

        var maxHH =0;
        $(".common_news li a img").each(function(index, el) {
            if(maxHH < $(this).height()){
                maxHH = $(this).height()
                return maxHH;
            }
        });

        var myHeight = Math.floor(maxHH+spanH);
        $(".common_news li").height(myHeight);
       
    });
     
    // 判断图片加载的函数
    function isImgLoad(callback){

        // 注意我的图片类名都是cover，因为我只需要处理cover。其它图片可以不管。
        // 查找所有封面图，迭代处理
        $(".common_news li img").each(function(){
            // 找到为0就将isLoad设为false，并退出each
            if(this.height === 0){
                isLoad = false;
                return false;
            }
        });
        // 为true，没有发现为0的。加载完毕
        if(isLoad){

            clearTimeout(t_img1); // 清除定时器
            // 回调函数
            callback();
        // 为false，因为找到了没有加载完成的图，将调用定时器递归
        }else{
            isLoad = true;
            t_img1 = setTimeout(function(){
                isImgLoad(callback); // 递归扫描
            },100); // 我这里设置的是500毫秒就扫描一次，可以自己调整
        }
    }
    


    
    var navbtn = $('.nav-btn');
    var box = $('.allpage'),
        blackFixed = $(".black-fixed");
    function navShow() {
        if(box.hasClass('clicked')){
            blackFixed.removeClass('black-clicked');
            box.removeClass('clicked');
            $(".head,.footer,.type").removeClass('clicked');
            $(".nav").removeClass('fixed');
        }else{
            box.addClass('clicked');
            $(".head,.footer,.type").addClass('clicked');
            $(".nav").addClass('fixed');
            blackFixed.addClass('black-clicked');
        }
    };
    navbtn.click(navShow);
    blackFixed.click(navShow);

    $('.top-search').click(function(){     // 搜索
        var search = $(".search"),
            _this = $(this);
        if(search.css("display") == "none"){
            search.show();
            _this.html("&#xe609;");
            $(".search-input").focus();
        }else{
            search.hide();
            _this.html("&#xe60f;");
            $(".search-input").blur();
        }
    });

    $('.class-btn').click(function(){     // 分类
        $(".type").toggle();
    });
    $('.common-search-btn').click(function(){     // 分类
        $(".common-search").toggle();
    });


    // 新增
    $(".cart-select").click(function(){
        $(this).toggleClass("on");
    });
    $(".cart_order_type span").on("click",function(){
        if($(this).hasClass("type_company")){
            $(".cart_type_company").show();
            $(this).parents(".cart_order_type").find("span").removeClass("on");
            $(this).addClass("on");
        }else{
            $(".cart_type_company").hide();
            $(this).parents(".cart_order_type").find("span").removeClass("on");
            $(this).addClass("on");
        }
    });


    // 星星
    var starPF = $(".discuss_list_top_right font"),
        starArr = {},
        starColor = $(".discuss_list_top_right");
    for(var i=0;i<starPF.length;i++){
        starArr[i] = starPF.eq(i).text();
        starColor.eq(i).find("div").find("p").css("width",starArr[i]/5*100 + 2 + "%");
    }
    $(".sub_name").click(function(){
        $(".top_order").toggle();
    });

    $(".vip_backup").click(function(){
        $('body,html').animate({scrollTop:0},500)
    });

    $(".pro-img-big").click(function(){
        $(".product li").css("width","100%");
        $(this).addClass("show");
        $(".pro-img-small").removeClass("show");
    });
    $(".pro-img-small").click(function(){
        $(".product li").css("width","50%");
        $(this).addClass("show");
        $(".pro-img-big").removeClass("show");
    });
});

// banner

(function(window){if(window.addEventListener){window.addEventListener('DOMContentLoaded',function(){slider('.big-pic-in','.pic-list','a','.slide-dot','span','slide-dot-cur',300,8000);slider('.slide','.slide-con','.slide-item','.tab-nav','a','tab-nav-cur',300);function slider(slide,slideCon,slideItem,nav,navItem,navCur,delay,autoTime){var slides=document.querySelectorAll(slide),navs=document.querySelectorAll(nav),length=slides.length;for(var i=0;i<length;i++){new Slider(slides[i],slideCon,slideItem,navs[i],navItem,navCur,delay,autoTime);}}
    function Slider(slide,slideCon,slideItem,nav,navItem,navCur,delay,autoTime){var slide=slide,slideCon=slide.querySelector(slideCon),slideItem=slide.querySelectorAll(slideItem),nav=nav,navItem=nav.querySelectorAll(navItem),navCur=navCur,delay=delay,autoTime=autoTime||false,length=slideItem.length,temp,cur=0,x,y,startX,startY,dx,dy,dir={left:false},isTouch=false,isBegin=false,isMove=false,isEnd=true,autoT;var width=slide.clientWidth;window.addEventListener('resize',function(){width=slide.clientWidth;},false);function autoPlay(){if(autoTime<4*delay||isBegin||isMove||!isEnd){return;}
        autoT=setTimeout(function(){isEnd=false;temp=cur;cur=cur===length-1?0:cur+1;Slider.addClass(slideCon,'transition');Slider.removeClass(navItem[temp],navCur);Slider.addClass(navItem[cur],navCur);Slider.transform(slideCon,(-cur*100/length)+'%');setTimeout(function(){Slider.removeClass(slideCon,'transition');isEnd=true;autoPlay();},delay);},autoTime);}
        function clearPlay(){try{clearTimeout(autoT);}catch(e){return;}}
        var aaa=0;this.start=function(e){clearPlay();if(!isEnd){return;}
            isBegin=true;x=y=0;if(e.targetTouches){isTouch=true;startX=e.targetTouches[0].clientX;startY=e.targetTouches[0].clientY;}else{startX=e.clientX;startY=e.clientY;}};this.move=function(e){clearPlay();if(!isBegin){return;}
            isMove=true;var tempX,tempY;if(isTouch){tempX=e.targetTouches[0].clientX;tempY=e.targetTouches[0].clientY;}else{tempX=e.clientX;tempY=e.clientY;}
            dx=tempX-startX;dy=tempY-startY;startX=tempX;startY=tempY;x+=dx;y+=dy;if(dir.horizontal===undefined){if(Math.abs(dx)>Math.abs(dy)){dir.horizontal=true;}else{dir.horizontal=false;}}
            if(dir.horizontal){e.preventDefault();if((cur===0&&x>0)||(cur===length-1&&x<0)){Slider.transform(slideCon,((x/6/width-cur)*100/length)+'%');}
            else{Slider.transform(slideCon,((x/width-cur)*100/length)+'%');}}else{isBegin=false;isMove=false;isEnd=true;delete dir.horizontal;return;}};this.end=function(e){if(!isBegin||!isMove){isBegin=false;isMove=false;isEnd=true;delete dir.horizontal;autoPlay();return;}
            isEnd=false;isBegin=false;isMove=false;Slider.addClass(slideCon,'transition');temp=cur;if(x>0){cur=cur===0?0:cur-1;}else{cur=cur===length-1?length-1:cur+1;}
            Slider.removeClass(navItem[temp],navCur);Slider.addClass(navItem[cur],navCur);Slider.transform(slideCon,(-cur*100/length)+'%');setTimeout(function(){x=0;Slider.removeClass(slideCon,'transition');isEnd=true;autoPlay();},delay);delete dir.horizontal;};slideCon.addEventListener('touchstart',this.start,false);slideCon.addEventListener('touchmove',this.move,false);slideCon.addEventListener('touchend',this.end,false);slideCon.addEventListener('touchcancel',this.end,false);slideCon.addEventListener('mousedown',this.start,false);slideCon.addEventListener('mousemove',this.move,false);slideCon.addEventListener('mouseup',this.end,false);slideCon.addEventListener('mouseout',this.end,false);autoPlay();}
    Slider.transform=function(element,x,y,z){x=x||0;y=y||0;z=z||0;var style=element.style;if(typeof x==='string'&&(x.indexOf('%')||y.indexOf('%'||z.indexOf('%')))){style.WebkitTransform='translate3d('+x+', '+y+', '+z+')';style.MozTransform='translate('+x+', '+y+')';style.OTransform='translate('+x+', '+y+')';style.transform='translate('+x+', '+y+')';}else{style.WebkitTransform='translate3d('+x+'%, '+y+'%, '+z+'%)';style.MozTransform='translate('+x+'%, '+y+'%)';style.OTransform='translate('+x+'%, '+y+'%)';style.transform='translate('+x+'%, '+y+'%)';}};Slider.addClass=function(element,className){var temp=element.className;element.className=temp+' '+className;}
    Slider.removeClass=function(element,className){var temp=element.className;temp=temp.split(' ');for(var i=0,length=temp.length;i<length;i++){if(temp[i]===className){temp.splice(i,1);break;}}
        element.className=temp.join(' ');}},false);}})(window);