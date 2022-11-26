/*
 *  date - 数据，name - .className 或者 #id，width - 宽度，height - 高度
 *  slide： banner效果，需要传递相应的.classname或者#id，不然默认插入网页底部；
 *  couplet： 对联广告，最好不指定name值，默认插入body；
 *  floatP： 浮动广告，乱飘的那个，最好不指定name值，默认插入body，开完CPU飞起来，慎开；
 *  mount： 泰山压顶，不指定name值，默认插到网页顶部；
 *  bottomCenter： 网页底部居中广告；
 *  bottomRight ： 网页底部居右广告；
 *  randomPic： 随机单张，需给定name值，不然默认插到网页底部；
 *  picList： 图片罗列，所有传过来的数据全部罗列出来；
 */
function bindClose(closeBtn,closeBox){
    $(closeBtn).click(function(){
        $(closeBox).remove();
    });
}

function errorsAlert(width,height){
    if(width == "" || height==""){
        alert("广告未填写高度或宽度，请重新填写！");
    }
}
function noneLink(link){
    if(link == ""){
        links = "javascript:;";
        target = "_self";
    }else{
        links = link;
        target = "_blank";
    }
}
function Atm(){
    var links="",
        target="";
}
Atm.prototype = {
    slide:function(data,name,width,height){ // 幻灯片
        errorsAlert(width,height);
        name == "" ? name = "body" : name;
        var kLi = "";
        !function(a){a.fn.slide=function(b){return a.fn.slide.defaults={type:"slide",effect:"fade",autoPlay:!1,delayTime:500,interTime:2500,triggerTime:150,defaultIndex:0,titCell:".hd li",mainCell:".bd",targetCell:null,trigger:"mouseover",scroll:1,vis:1,titOnClassName:"on",autoPage:!1,prevCell:".atm_prev",nextCell:".atm_next",pageStateCell:".pageState",opp:!1,pnLoop:!0,easing:"swing",startFun:null,endFun:null,switchLoad:null,playStateCell:".playState",mouseOverStop:!0,defaultPlay:!0,returnDefault:!1},this.each(function(){var c=a.extend({},a.fn.slide.defaults,b),d=a(this),e=c.effect,f=a(c.prevCell,d),g=a(c.nextCell,d),h=a(c.pageStateCell,d),i=a(c.playStateCell,d),j=a(c.titCell,d),k=j.size(),l=a(c.mainCell,d),m=l.children().size(),n=c.switchLoad,o=a(c.targetCell,d),p=parseInt(c.defaultIndex),q=parseInt(c.delayTime),r=parseInt(c.interTime);parseInt(c.triggerTime);var Q,t=parseInt(c.scroll),u=parseInt(c.vis),v="false"==c.autoPlay||0==c.autoPlay?!1:!0,w="false"==c.opp||0==c.opp?!1:!0,x="false"==c.autoPage||0==c.autoPage?!1:!0,y="false"==c.pnLoop||0==c.pnLoop?!1:!0,z="false"==c.mouseOverStop||0==c.mouseOverStop?!1:!0,A="false"==c.defaultPlay||0==c.defaultPlay?!1:!0,B="false"==c.returnDefault||0==c.returnDefault?!1:!0,C=0,D=0,E=0,F=0,G=c.easing,H=null,I=null,J=null,K=c.titOnClassName,L=j.index(d.find("."+K)),M=p=-1==L?p:L,N=p,O=p,P=m>=u?0!=m%t?m%t:t:0,R="leftMarquee"==e||"topMarquee"==e?!0:!1,S=function(){a.isFunction(c.startFun)&&c.startFun(p,k,d,a(c.titCell,d),l,o,f,g)},T=function(){a.isFunction(c.endFun)&&c.endFun(p,k,d,a(c.titCell,d),l,o,f,g)},U=function(){j.removeClass(K),A&&j.eq(N).addClass(K)};if("menu"==c.type)return A&&j.removeClass(K).eq(p).addClass(K),j.hover(function(){Q=a(this).find(c.targetCell);var b=j.index(a(this));I=setTimeout(function(){switch(p=b,j.removeClass(K).eq(p).addClass(K),S(),e){case"fade":Q.stop(!0,!0).animate({opacity:"show"},q,G,T);break;case"slideDown":Q.stop(!0,!0).animate({height:"show"},q,G,T)}},c.triggerTime)},function(){switch(clearTimeout(I),e){case"fade":Q.animate({opacity:"hide"},q,G);break;case"slideDown":Q.animate({height:"hide"},q,G)}}),B&&d.hover(function(){clearTimeout(J)},function(){J=setTimeout(U,q)}),void 0;if(0==k&&(k=m),R&&(k=2),x){if(m>=u)if("leftLoop"==e||"topLoop"==e)k=0!=m%t?(0^m/t)+1:m/t;else{var V=m-u;k=1+parseInt(0!=V%t?V/t+1:V/t),0>=k&&(k=1)}else k=1;j.html("");var W="";if(1==c.autoPage||"true"==c.autoPage)for(var X=0;k>X;X++)W+="<li>"+(X+1)+"</li>";else for(var X=0;k>X;X++)W+=c.autoPage.replace("$",X+1);j.html(W);var j=j.children()}if(m>=u){l.children().each(function(){a(this).width()>E&&(E=a(this).width(),D=a(this).outerWidth(!0)),a(this).height()>F&&(F=a(this).height(),C=a(this).outerHeight(!0))});var Y=l.children(),Z=function(){for(var a=0;u>a;a++)Y.eq(a).clone().addClass("clone").appendTo(l);for(var a=0;P>a;a++)Y.eq(m-a-1).clone().addClass("clone").prependTo(l)};switch(e){case"fold":l.css({position:"relative",width:D,height:C}).children().css({position:"absolute",width:E,left:0,top:0,display:"none"});break;case"top":l.wrap('<div class="tempWrap" style="overflow:hidden; position:relative; height:'+u*C+'px"></div>').css({top:-(p*t)*C,position:"relative",padding:"0",margin:"0"}).children().css({height:F});break;case"left":l.wrap('<div class="tempWrap" style="overflow:hidden; position:relative; width:'+u*D+'px"></div>').css({width:m*D,left:-(p*t)*D,position:"relative",overflow:"hidden",padding:"0",margin:"0"}).children().css({"float":"left",width:E});break;case"leftLoop":case"leftMarquee":Z(),l.wrap('<div class="tempWrap" style="overflow:hidden; position:relative; width:'+u*D+'px"></div>').css({width:(m+u+P)*D,position:"relative",overflow:"hidden",padding:"0",margin:"0",left:-(P+p*t)*D}).children().css({"float":"left",width:E});break;case"topLoop":case"topMarquee":Z(),l.wrap('<div class="tempWrap" style="overflow:hidden; position:relative; height:'+u*C+'px"></div>').css({height:(m+u+P)*C,position:"relative",padding:"0",margin:"0",top:-(P+p*t)*C}).children().css({height:F})}}var $=function(a){var b=a*t;return a==k?b=m:-1==a&&0!=m%t&&(b=-m%t),b},_=function(b){var c=function(c){for(var d=c;u+c>d;d++)b.eq(d).find("img["+n+"]").each(function(){var b=a(this);if(b.attr("src",b.attr(n)).removeAttr(n),l.find(".clone")[0])for(var c=l.children(),d=0;d<c.size();d++)c.eq(d).find("img["+n+"]").each(function(){a(this).attr(n)==b.attr("src")&&a(this).attr("src",a(this).attr(n)).removeAttr(n)})})};switch(e){case"fade":case"fold":case"top":case"left":case"slideDown":c(p*t);break;case"leftLoop":case"topLoop":c(P+$(O));break;case"leftMarquee":case"topMarquee":var d="leftMarquee"==e?l.css("left").replace("px",""):l.css("top").replace("px",""),f="leftMarquee"==e?D:C,g=P;if(0!=d%f){var h=Math.abs(0^d/f);g=1==p?P+h:P+h-1}c(g)}},ab=function(a){if(!A||M!=p||a||R){if(R?p>=1?p=1:0>=p&&(p=0):(O=p,p>=k?p=0:0>p&&(p=k-1)),S(),null!=n&&_(l.children()),o[0]&&(Q=o.eq(p),null!=n&&_(o),"slideDown"==e?(o.not(Q).stop(!0,!0).slideUp(q),Q.slideDown(q,G,function(){l[0]||T()})):(o.not(Q).stop(!0,!0).hide(),Q.animate({opacity:"show"},q,function(){l[0]||T()}))),m>=u)switch(e){case"fade":l.children().stop(!0,!0).eq(p).animate({opacity:"show"},q,G,function(){T()}).siblings().hide();break;case"fold":l.children().stop(!0,!0).eq(p).animate({opacity:"show"},q,G,function(){T()}).siblings().animate({opacity:"hide"},q,G);break;case"top":l.stop(!0,!1).animate({top:-p*t*C},q,G,function(){T()});break;case"left":l.stop(!0,!1).animate({left:-p*t*D},q,G,function(){T()});break;case"leftLoop":var b=O;l.stop(!0,!0).animate({left:-($(O)+P)*D},q,G,function(){-1>=b?l.css("left",-(P+(k-1)*t)*D):b>=k&&l.css("left",-P*D),T()});break;case"topLoop":var b=O;l.stop(!0,!0).animate({top:-($(O)+P)*C},q,G,function(){-1>=b?l.css("top",-(P+(k-1)*t)*C):b>=k&&l.css("top",-P*C),T()});break;case"leftMarquee":var c=l.css("left").replace("px","");0==p?l.animate({left:++c},0,function(){l.css("left").replace("px","")>=0&&l.css("left",-m*D)}):l.animate({left:--c},0,function(){l.css("left").replace("px","")<=-(m+P)*D&&l.css("left",-P*D)});break;case"topMarquee":var d=l.css("top").replace("px","");0==p?l.animate({top:++d},0,function(){l.css("top").replace("px","")>=0&&l.css("top",-m*C)}):l.animate({top:--d},0,function(){l.css("top").replace("px","")<=-(m+P)*C&&l.css("top",-P*C)})}j.removeClass(K).eq(p).addClass(K),M=p,y||(g.removeClass("nextStop"),f.removeClass("prevStop"),0==p&&f.addClass("prevStop"),p==k-1&&g.addClass("nextStop")),h.html("<span>"+(p+1)+"</span>/"+k)}};A&&ab(!0),B&&d.hover(function(){clearTimeout(J)},function(){J=setTimeout(function(){p=N,A?ab():"slideDown"==e?Q.slideUp(q,U):Q.animate({opacity:"hide"},q,U),M=p},300)});var bb=function(a){H=setInterval(function(){w?p--:p++,ab()},a?a:r)},cb=function(a){H=setInterval(ab,a?a:r)},db=function(){z||(clearInterval(H),bb())},eb=function(){(y||p!=k-1)&&(p++,ab(),R||db())},fb=function(){(y||0!=p)&&(p--,ab(),R||db())},gb=function(){clearInterval(H),R?cb():bb(),i.removeClass("pauseState")},hb=function(){clearInterval(H),i.addClass("pauseState")};if(v?R?(w?p--:p++,cb(),z&&l.hover(hb,gb)):(bb(),z&&d.hover(hb,gb)):(R&&(w?p--:p++),i.addClass("pauseState")),i.click(function(){i.hasClass("pauseState")?gb():hb()}),"mouseover"==c.trigger?j.hover(function(){var a=j.index(this);I=setTimeout(function(){p=a,ab(),db()},c.triggerTime)},function(){clearTimeout(I)}):j.click(function(){p=j.index(this),ab(),db()}),R){if(g.mousedown(eb),f.mousedown(fb),y){var ib,jb=function(){ib=setTimeout(function(){clearInterval(H),cb(0^r/10)},150)},kb=function(){clearTimeout(ib),clearInterval(H),cb()};g.mousedown(jb),g.mouseup(kb),f.mousedown(jb),f.mouseup(kb)}"mouseover"==c.trigger&&(g.hover(eb,function(){}),f.hover(fb,function(){}))}else g.click(eb),f.click(fb)})}}(jQuery),jQuery.easing.jswing=jQuery.easing.swing,jQuery.extend(jQuery.easing,{def:"easeOutQuad",swing:function(a,b,c,d,e){return jQuery.easing[jQuery.easing.def](a,b,c,d,e)},easeInQuad:function(a,b,c,d,e){return d*(b/=e)*b+c},easeOutQuad:function(a,b,c,d,e){return-d*(b/=e)*(b-2)+c},easeInOutQuad:function(a,b,c,d,e){return(b/=e/2)<1?d/2*b*b+c:-d/2*(--b*(b-2)-1)+c},easeInCubic:function(a,b,c,d,e){return d*(b/=e)*b*b+c},easeOutCubic:function(a,b,c,d,e){return d*((b=b/e-1)*b*b+1)+c},easeInOutCubic:function(a,b,c,d,e){return(b/=e/2)<1?d/2*b*b*b+c:d/2*((b-=2)*b*b+2)+c},easeInQuart:function(a,b,c,d,e){return d*(b/=e)*b*b*b+c},easeOutQuart:function(a,b,c,d,e){return-d*((b=b/e-1)*b*b*b-1)+c},easeInOutQuart:function(a,b,c,d,e){return(b/=e/2)<1?d/2*b*b*b*b+c:-d/2*((b-=2)*b*b*b-2)+c},easeInQuint:function(a,b,c,d,e){return d*(b/=e)*b*b*b*b+c},easeOutQuint:function(a,b,c,d,e){return d*((b=b/e-1)*b*b*b*b+1)+c},easeInOutQuint:function(a,b,c,d,e){return(b/=e/2)<1?d/2*b*b*b*b*b+c:d/2*((b-=2)*b*b*b*b+2)+c},easeInSine:function(a,b,c,d,e){return-d*Math.cos(b/e*(Math.PI/2))+d+c},easeOutSine:function(a,b,c,d,e){return d*Math.sin(b/e*(Math.PI/2))+c},easeInOutSine:function(a,b,c,d,e){return-d/2*(Math.cos(Math.PI*b/e)-1)+c},easeInExpo:function(a,b,c,d,e){return 0==b?c:d*Math.pow(2,10*(b/e-1))+c},easeOutExpo:function(a,b,c,d,e){return b==e?c+d:d*(-Math.pow(2,-10*b/e)+1)+c},easeInOutExpo:function(a,b,c,d,e){return 0==b?c:b==e?c+d:(b/=e/2)<1?d/2*Math.pow(2,10*(b-1))+c:d/2*(-Math.pow(2,-10*--b)+2)+c},easeInCirc:function(a,b,c,d,e){return-d*(Math.sqrt(1-(b/=e)*b)-1)+c},easeOutCirc:function(a,b,c,d,e){return d*Math.sqrt(1-(b=b/e-1)*b)+c},easeInOutCirc:function(a,b,c,d,e){return(b/=e/2)<1?-d/2*(Math.sqrt(1-b*b)-1)+c:d/2*(Math.sqrt(1-(b-=2)*b)+1)+c},easeInElastic:function(a,b,c,d,e){var f=1.70158,g=0,h=d;if(0==b)return c;if(1==(b/=e))return c+d;if(g||(g=.3*e),h<Math.abs(d)){h=d;var f=g/4}else var f=g/(2*Math.PI)*Math.asin(d/h);return-(h*Math.pow(2,10*(b-=1))*Math.sin((b*e-f)*2*Math.PI/g))+c},easeOutElastic:function(a,b,c,d,e){var f=1.70158,g=0,h=d;if(0==b)return c;if(1==(b/=e))return c+d;if(g||(g=.3*e),h<Math.abs(d)){h=d;var f=g/4}else var f=g/(2*Math.PI)*Math.asin(d/h);return h*Math.pow(2,-10*b)*Math.sin((b*e-f)*2*Math.PI/g)+d+c},easeInOutElastic:function(a,b,c,d,e){var f=1.70158,g=0,h=d;if(0==b)return c;if(2==(b/=e/2))return c+d;if(g||(g=e*.3*1.5),h<Math.abs(d)){h=d;var f=g/4}else var f=g/(2*Math.PI)*Math.asin(d/h);return 1>b?-.5*h*Math.pow(2,10*(b-=1))*Math.sin((b*e-f)*2*Math.PI/g)+c:.5*h*Math.pow(2,-10*(b-=1))*Math.sin((b*e-f)*2*Math.PI/g)+d+c},easeInBack:function(a,b,c,d,e,f){return void 0==f&&(f=1.70158),d*(b/=e)*b*((f+1)*b-f)+c},easeOutBack:function(a,b,c,d,e,f){return void 0==f&&(f=1.70158),d*((b=b/e-1)*b*((f+1)*b+f)+1)+c},easeInOutBack:function(a,b,c,d,e,f){return void 0==f&&(f=1.70158),(b/=e/2)<1?d/2*b*b*(((f*=1.525)+1)*b-f)+c:d/2*((b-=2)*b*(((f*=1.525)+1)*b+f)+2)+c},easeInBounce:function(a,b,c,d,e){return d-jQuery.easing.easeOutBounce(a,e-b,0,d,e)+c},easeOutBounce:function(a,b,c,d,e){return(b/=e)<1/2.75?d*7.5625*b*b+c:2/2.75>b?d*(7.5625*(b-=1.5/2.75)*b+.75)+c:2.5/2.75>b?d*(7.5625*(b-=2.25/2.75)*b+.9375)+c:d*(7.5625*(b-=2.625/2.75)*b+.984375)+c},easeInOutBounce:function(a,b,c,d,e){return e/2>b?.5*jQuery.easing.easeInBounce(a,2*b,0,d,e)+c:.5*jQuery.easing.easeOutBounce(a,2*b-e,0,d,e)+.5*d+c}});
        for(var i=0; i<data.length; i++){
            noneLink(data[i].href);
            kLi += "<li><a href='"+ links +"' title='"+ data[i].title +"' target='"+ target +"'><img src='"+ data[i].url +"' alt='"+ data[i].title +"' width='"+ width +"' height='"+ height +"' /></a></li>";
        }
        var atmSlide = "<div class='atm_banner'><div class='atm_banner_box'><ul class='atm_banner_pic'>"+ kLi +"</ul><div class='atm_btns'><div><a class='atm_prev' href='javascript:void(0)'></a><a class='atm_next' href='javascript:void(0)'></a></div></div></div></div>";
        $(name).append(atmSlide);
        var atm_banner = $(".atm_banner");
        atm_banner.css("height",height);
        $(".atm_banner_box").css({"width":width,"margin-left":-width/2+"px"});
        atm_banner.slide({mainCell:".atm_banner_pic" , effect:"fold", autoPlay:true, delayTime:700 , autoPage:true });
    },
    couplet:function(data,name,width,height){ // 对联广告
        errorsAlert(width,height);
        name == "" ? name = "body" : name;
        noneLink(data[0].href);
        var atmCouplet = "<span class='atm_couplet_left_close'>关闭</span><span class='atm_couplet_right_close'>关闭</span><a href='"+ links +"' title='"+ data[0].title +"' class='atm_couplet_left' target='"+ target +"'><img src='"+ data[0].url +"' width='"+ width +"' height='"+ height +"' alt='"+ data[0].title +"'/></a><a href='"+ links +"' title='"+ data[1].title +"' class='atm_couplet_right' target='"+ target +"'><img src='"+ data[1].url +"' width='"+ width +"' height='"+ height +"' alt='"+ data[1].title +"'/></a>";
        $(name).append(atmCouplet);
        bindClose(".atm_couplet_left_close,.atm_couplet_right_close",".atm_couplet_left_close,.atm_couplet_right_close,.atm_couplet_left,.atm_couplet_right");
    },
    floatP:function(data,name,width,height){ // 浮动广告
        errorsAlert(width,height);
        noneLink(data[0].href);
        name == "" ? name = "body" : name;
        (function ( $, window, document, undefined ) {
            var pluginName = 'floatPic';
            var defaults = {
                step: 1,
                delay: 50,
                isLinkClosed: false,
                onClose: function(elem){}
            };
            var ads = {
                linkUrl: '#',
                'z-index': '100',
                'closed-icon': '',
                imgHeight: '',
                imgWidth: '',
                title: data[0].title,
                img: '#',
                linkWindow: target,
                headFilter: 0.2
            };

            function Plugin(element, options) {
                this.element = element;
                this.options = $.extend(
                    {},
                    defaults,
                    options,
                    {
                        width: $(window).width(),
                        height: $(window).height(),
                        xPos: this.getRandomNum(0, $(window).width() - $(element).innerWidth()),
                        yPos: this.getRandomNum(0, 300),
                        yOn: this.getRandomNum(0, 1),
                        xOn: this.getRandomNum(0, 1),
                        yPath: this.getRandomNum(0, 1),
                        xPath: this.getRandomNum(0, 1),
                        hOffset: $(element).innerHeight(),
                        wOffset: $(element).innerWidth(),
                        fn: function(){},
                        interval: 0
                    }
                );
                this._defaults = defaults;
                this._name = pluginName;

                this.init();
            }

            Plugin.prototype = {
                init: function () {
                    var elem = $(this.element);
                    var defaults = this.options;
                    var p = this;
                    var xFlag = 0;
                    var yFlag = 0;

                    elem.css({"left": defaults.xPos + p.scrollX(), "top": defaults.yPos + p.scrollY()});
                    defaults.fn = function(){
                        defaults.width = $(window).width();
                        defaults.height = $(window).height();

                        if(xFlag == p.scrollX() && yFlag == p.scrollY()){
                            elem.css({"left": defaults.xPos + p.scrollX(), "top": defaults.yPos + p.scrollY()});
                            if (defaults.yOn)
                                defaults.yPos = defaults.yPos + defaults.step;
                            else
                                defaults.yPos = defaults.yPos - defaults.step;

                            if (defaults.yPos <= 0) {
                                defaults.yOn = 1;
                                defaults.yPos = 0;
                            }
                            if (defaults.yPos >= (defaults.height - defaults.hOffset)) {
                                defaults.yOn = 0;
                                defaults.yPos = (defaults.height - defaults.hOffset);
                            }

                            if (defaults.xOn)
                                defaults.xPos = defaults.xPos + defaults.step;
                            else
                                defaults.xPos = defaults.xPos - defaults.step;

                            if (defaults.xPos <= 0) {
                                defaults.xOn = 1;
                                defaults.xPos = 0;
                            }
                            if (defaults.xPos >= (defaults.width - defaults.wOffset)) {
                                defaults.xOn = 0;
                                defaults.xPos = (defaults.width - defaults.wOffset);
                            }
                        }
                        yFlag = $(window).scrollTop();
                        xFlag = $(window).scrollLeft();
                    };
                    this.run(elem, defaults);
                },
                run: function(elem, defaults){
                    this.start(elem, defaults);
                    this.adEvent(elem,defaults);
                },
                start: function(elem, defaults){
                    elem.find('div.floatPicClose').hide();
                    defaults.interval = window.setInterval(defaults.fn,  defaults.delay);
                    window.setTimeout(function(){elem.show();}, defaults.delay);
                },
                getRandomNum: function (Min, Max){
                    var Range = Max - Min;
                    var Rand = Math.random();
                    return(Min + Math.round(Rand * Range));
                },
                getPath: function(on){
                    return on ? 0 : 1;
                },
                clear: function(elem, defaults){
                    elem.find('div.floatPicClose').show();
                    window.clearInterval(defaults.interval);
                },
                close: function(elem, defaults, isClose){
                    //elem.unbind('hover');
                    elem.unbind("mouseenter mouseleave");
                    elem.hide();
                    if(isClose)
                        defaults.onClose.call(elem);
                },
                adEvent: function(elem, defaults){
                    var obj = {
                        elem: this,
                        fn_close: function() {
                            this.elem.close(elem, defaults, true);
                        },
                        fn_clear: function() {
                            if(this.elem.options.isLinkClosed)
                                this.elem.close(elem, defaults, false);
                        }
                    };

                    elem.find('div.floatPicClose').bind('click', jQuery.proxy(obj, "fn_close"));

                    elem.find('a').bind('click', jQuery.proxy(obj, "fn_clear"));

                    var stop = {
                        elem: this,
                        over: function(){
                            this.elem.clear(elem, defaults);
                        },
                        out: function(){
                            this.elem.start(elem, defaults);
                        }
                    };
                    elem.bind("mouseenter", jQuery.proxy(stop, "over"));
                    elem.bind("mouseleave", jQuery.proxy(stop, "out"));
                },
                scrollX: function(){
                    var de = document.documentElement;
                    return self.pageXOffset || (de && de.scrollLeft) || document.body.scrollLeft;
                },
                scrollY: function(){
                    var de = document.documentElement;
                    return self.pageYOffset || (de && de.scrollTop) || document.body.scrollTop;
                }
            };
            $.fn.floatPic = function(options) {
                return this.children("div").each(function (i, elem) {
                    if (!$.data(this, 'plugin_' + pluginName)) {
                        $.data(this, 'plugin_' + pluginName, new Plugin(this, options));
                    }
                });
            };
            $.floatPic = function(options){

                if(options){
                    if(options.ad){
                        var adDiv = $('#' + pluginName);

                        if(adDiv.length <= 0)
                            adDiv = $('<div>', {
                                'id': pluginName,
                                'class': pluginName
                            }).appendTo('body');

                        for(var i in options.ad){

                            var ad = options.ad[i];
                            ad = $.extend({}, ads, ad);
                            var div = $('<div>', {
                                'class': 'floatImg'
                            });

                            div.css("z-index", ad['z-index']);
                            var closeDiv = $('<div>', {
                                'class': 'floatPicClose'
                            });
                            closeDiv.appendTo(div);
                            var content = $('<div>');

                            $('<a>', {
                                href: ad.linkUrl,
                                target: ad.linkWindow,
                                title: ad.title
                            }).append(
                                $('<img>', {
                                    'src': ad.img,
                                    'style': (ad.imgHeight ? 'height:' + ad.imgHeight + 'px;' : '') +
                                        (ad.imgWidth ? 'width:' + ad.imgWidth + 'px;' : '')
                                })
                            ).appendTo(content);

                            content.appendTo(div);

                            div.appendTo(adDiv);
                        }
                        delete options.ad;
                        $('#' + pluginName).floatPic(options);
                    }
                }
                else
                    $.error('浮动模块出错!');
            };
        })(jQuery, window, document);
        $.floatPic({
            delay: 10,
            //isLinkClosed: true,
            ad:	[{
                'img': data[0].url,
                'imgHeight': width,
                'imgWidth': height,
                'linkUrl': links,
                'z-index': 100
            }]
        });
        $(".floatPicClose").text("X");
    },
    mount:function(data,name,width,height){ // 泰山压顶
        errorsAlert(width,height);
        name == "" ? name = "body" : name;
        noneLink(data[0].href);
        var atmMount = "<div class='mount'><span class='mount_close'>X</span><a href='"+ links +"' target='"+ target +"' title='"+ data[0].title+"'><img src='"+ data[0].url +"' alt='"+ data[0].title+"' width='"+ width +"' height='"+ height +"'/></a></div>";
        $(name).prepend(atmMount);
        $(".mount").css("height",height);
        $(".mount a").css("margin-left",-width/2);
        $(".mount").slideDown(1200,function(){
            setTimeout(function(){
                $(".mount").slideUp(1000);
            },5000);
        });
        bindClose(".mount_close",".mount");
    },
    bottomCenter:function(data,name,width,height){ // 底部浮动居中广告
        errorsAlert(width,height);
        name == "" ? name = "body" : name;
        noneLink(data[0].href);
        var atmBottomCenter = "<div class='atm_bottom_pic_bg'><div class='atm_bottom_pic'><span class='atm_bottom_pic_close'>X</span><a href='"+ links +"' target='"+ target +"' title='"+ data[0].title+"'><img src='"+ data[0].url +"' alt='"+ data[0].title+"' width='"+ width +"' height='"+ height +"'/></a></div></div>";
        $(name).prepend(atmBottomCenter);
        $(".atm_bottom_pic").css("height",height);
        $(".atm_bottom_pic a").css("margin-left",-width/2);
        bindClose(".atm_bottom_pic_close",".atm_bottom_pic_bg");
    },
    bottomRight:function(data,name,width,height){  // 底部浮动居右广告
        errorsAlert(width,height);
        name == "" ? name = "body" : name;
        noneLink(data[0].href);
        var atmBottomRight = "<div class='atm_bottom_right'><div class='atm_bottom_right_pic'><span class='atm_bottom_pic_close'>X</span><a href='"+ links +"' target='"+ target +"' title='"+ data[0].title+"'><img src='"+ data[0].url +"' alt='"+ data[0].title+"' width='"+ width +"' height='"+ height +"'/></a></div></div>";
        $(name).prepend(atmBottomRight);
        bindClose(".atm_bottom_pic_close",".atm_bottom_right");
    },
    randomPic:function(data,name,width,height){  // 随机单张广告
        errorsAlert(width,height);
        name == "" ? name = "body" : name;
        var atmRandom,
            i = parseInt(data.length*Math.random());
            noneLink(data[i].href);
            atmRandom = "<a href='"+ links +"' target='"+ target +"' title='"+ data[i].title +"' class='random_pic'><img src='"+ data[i].url +"' alt='"+ data[i].title +"' width='"+ width +"' height='"+ height +"'/></a>";
        $(name).append(atmRandom);
    },
    picList:function(data,name,width,height){  // 图片罗列广告
        errorsAlert(width,height);
        name == "" ? name = "body" : name;
        var i=0,
            picList="";
        for(i;i<data.length;i++){
            noneLink(data[i].href);
            picList += "<a href='"+ links +"' target='"+ target +"' title='"+ data[i].title +"' class='random_pic'><img src='"+ data[i].url +"' alt='"+ data[i].title +"' width='"+ width +"' height='"+ height +"'/></a>";
        }
        $(name).append(picList);
    }
}