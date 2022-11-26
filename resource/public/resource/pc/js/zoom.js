(function(window,undefined){
    var afx = {
        $afx: function(id){
            return document.getElementById(id);
        },
        getTags: function(obj,tag){
            if(obj){
                return obj.getElementsByTagName(tag);
            }else{
                return false;
            }
        },
        getIndex: function(obj,tag){
            var nowLength = afx.getTags(obj,tag).length;
            for(var i=0;i<nowLength;i++){
                obj[i].index = i;
                return obj[i];
            }
        },
        eListener: function(obj,event,fn){
            fn == "" ? fn = function(){} : fn;
            if(window.addEventListener){
                console.log(obj);
                obj.addEventListener(event,fn,false);
            }else{
                obj.attachEvent("on"+event,fn);
            }
        },
        conHeightAuto: function(){ // 计算左右的高度
            var conLeft = afx.$afx("contentLeft").clientHeight,
                conRight = afx.$afx("contentRight").clientHeight,
                con = afx.$afx("content");
            if(conLeft > conRight){
                con.style.height = conLeft + "px";
            }else{
                con.style.height = conRight + "px";
            }
        },
        getSortFun: function(order, sortBy) {
            var ordAlpah = (order == 'asc') ? '>' : '<';
            var sortFun = new Function('a', 'b', 'return a.' + sortBy + ordAlpah + 'b.' + sortBy + '?1:-1');
            return sortFun;
        },
        imgZoom: function(ele,parentEle,parentsEle,index){
            /*
             * ele:点击元素
             * parentEle:父级元素
             * parentsEle:祖父级元素
             * index:索引值
             */
            var picShow = "<div id='picShow'><div class='pic_show_box'><div id='pic_quit'></div><a href='javascript:;' title='上一张' id='lbtn'></a><a href='javascript:;' title='下一张' id='rbtn'></a><img width='593' height='442' alt='' /><p><a href='' id='picLink'></a></p></div></div>";
            var _this = ele,liIndex,liEle,picUrl,picShowBod,liLength,_val,_href,picShowText,lbtn,rbtn;
            liIndex = index;
            liEle = afx.getTags(parentsEle,parentEle)[liIndex]; // 获取当前点击的那个元素
            var picShowBox = document.createElement("div");
            picShowBox.innerHTML = picShow;
            document.body.appendChild(picShowBox); // 插入到页面中
            picShowBox = afx.$afx("picShow");
            picShowBod = afx.getTags(picShowBox,"img")[0]; // 获取放大的图片元素
            picShowText = afx.$afx("picLink");
            liLength = afx.getTags(parentsEle,parentEle).length;
            lbtn = afx.$afx("lbtn");
            rbtn = afx.$afx("rbtn");
            function setDate(liEle){
                picUrl = afx.getTags(liEle,"img")[0].getAttribute("src");
                _href = afx.getTags(liEle,"a")[0].getAttribute("href");
                _val = afx.getTags(liEle,"a")[1].innerHTML || afx.getTags(liEle,"a")[0].innerHTML;
				if(afx.getTags(liEle,"img")[0].parentNode.innerHTML==_val){
					_val = afx.getTags(liEle,"a")[0].innerHTML;
				}
				
            }
            function addData(picUrl,_val,_href){
                picShowBod.setAttribute("src",picUrl);
                picShowText.innerHTML = _val;
                picShowText.setAttribute("href",_href);
            }
            setDate(liEle);
            addData(picUrl,_val,_href);
            lbtn.onclick = function(){
                if(liIndex>0){
                    liEle = afx.getTags(parentsEle,parentEle)[liIndex-1];
                    liIndex--;
                }else{
                    liEle = afx.getTags(parentsEle,parentEle)[liLength-1];
                    liIndex = liLength-1;
                }
                setDate(liEle);
                addData(picUrl,_val,_href);
                listImgZoom("picShow","593");
            }
            rbtn.onclick = function(){
                if(liIndex<liLength-1){
                    liEle = afx.getTags(parentsEle,parentEle)[liIndex+1];
                    liIndex++;
                }else{
                    liEle = afx.getTags(parentsEle,parentEle)[0];
                    liIndex = 0;
                }
                setDate(liEle);
                addData(picUrl,_val,_href);
                listImgZoom("picShow","593");
            }
            afx.$afx("pic_quit").onclick = function(){
                picShowBox.parentNode.removeChild(picShowBox);
            }
        }
    }
    window.afx = afx;
    'afx' in window || (window.afx = afx);
})(window);

function imgZoomRun(box,tag1,name,tag2){
    var pro = afx.$afx(box);
    var pEle = afx.getTags(pro,tag1);
    var pArr = []; // 存放类名是prod-zoom的p标签
    for(var i=0; i<pEle.length; i++){
        if(pEle[i].className == name){ // 筛选
            pArr.push(pEle[i]);
        }
    }
    for(var j=0;j<pArr.length;j++){
        pArr[j].index = j;
        pArr[j].onclick = function(){
            afx.imgZoom(this,tag2,pro,this.index);
            listImgZoom("picShow","593");
        }
    }
}

function setInto(a){
    var returnHtml = "";
    for(var b=0; b< a.innerHTML.length; b++){
        returnHtml += "<font>" + a.innerHTML[b] + "</font>";
    }
    a.innerHTML = returnHtml;

}
function nextPrev(_this,numb){
    if(_this.previousSibling){
        _this.previousSibling.style.top = numb + "px";
    }
    if(_this.nextSibling){
        _this.nextSibling.style.top = numb + "px";
    }
}
function newsFontMove(newsId){
    var news5 = afx.$afx(newsId);
    var news5li = afx.getTags(news5,"li");
    for(var i=0; i<news5li.length; i++){
        var newsLiA = afx.getTags(news5li[i],"a")[0];
        setInto(newsLiA);
        var newsFont = afx.getTags(newsLiA,"font");
        for(var c=0; c<newsFont.length; c++){
            newsFont[c].index = c;
            newsFont[c].style.position = "relative";
            newsFont[c].onmouseover = function(){
                var _this = this;
                _this.style.top = "-8px";
                nextPrev(_this,"-5");
            }
            newsFont[c].onmouseout = function(){
                var _this = this;
                _this.style.top = "6px";
                nextPrev(_this,"4");
                setTimeout(function(){
                    _this.style.top = "0";
                    nextPrev(_this,"0");
                },100);
            }
        }
    }
}
function colorChange(newsId){
    var news4 = afx.$afx(newsId);
    var news4Li = afx.getTags(news4,"li");
    for(var n=0; n<news4Li.length; n++){
        if(n%2 == 0){
            news4Li[n].className += "news4-bg";
        }
    }
}

// 多级分类
function LeftType(arr,ul,id){
    this.parentDom = afx.$afx(ul); // 传入的ul#id
    this.allDom = ""; // 最后输入到页面的dom元素
    arr.sort(afx.getSortFun('desc', 't_order'));
    var _0 = new Array(),
        _1 = new Array(),
        _2 = new Array(),
        _3 = new Array(),
        _4 = new Array();
    for(var i=0; i<arr.length; i++){
        switch(arr[i].t_depth){
            case "0":
                _0.push(arr[i]);
                break;
            case "1":
                _1.push(arr[i]);
                break;
            case "2":
                _2.push(arr[i]);
                break;
            case "3":
                _3.push(arr[i]);
                break;
            case "4":
                _4.push(arr[i]);
                break;
        }
    }
    this.allArr = [_0,_1,_2,_3,_4];

    this.domPrint(_0); // 先把一级分类输出到页面
    this.parentDom.innerHTML = this.allDom;
    for(var z=0; z<this.allArr.length-1; z++){
        this.allDomArr(this.allArr[z+1],this.allArr[z]);
    }
    if(id !=0){
        this.pageDisplay(id,arr);
        var nowDOM = afx.$afx(id);
        afx.getTags(nowDOM,"a")[0].className = "on";
        for(var b=0; b<nowDOM.childNodes.length; b++){
            nowDOM.childNodes[b].style.display = "block"
        }
    }
}
LeftType.prototype = {
    allDomArr: function(subArr,parentArr){ // 根据层级（t_depth）绑定dom元素
        for(var x=0; x<subArr.length; x++){
            for(var y=0; y<parentArr.length; y++){
                if(subArr[x].t_pid == parentArr[y].t_id){
                    var subDom = "";
                    subDom += "<li id='"+ subArr[x].t_id +"'><a href='"+ subArr[x].t_url +"' title='"+ subArr[x].t_name +"'>" + subArr[x].t_name + "</a></li>";
                    var subUL = document.createElement("ul");
                    var domId = parentArr[y].t_id.toString();
                    subUL.innerHTML = subDom;
                    afx.$afx(domId).appendChild(subUL);
                }
            }
        }
    },
    domPrint: function(domArr){ // 根据传递的数组生成dom元素
        var domTags = "";
        for(var j=0; j<domArr.length; j++){
            domTags += "<li id='"+domArr[j].t_id+"'><a href='"+ domArr[j].t_url +"' title='"+ domArr[j].t_name +"'>"+domArr[j].t_name + "</a></li>";
        }
        this.allDom += domTags;
    },
    pageDisplay: function(id,arr){ // 根据当前id显示dom元素
        var tempId,
            parentDOM,
            thisId = id; // 当前元素的id
        function idShow(thisId){
            for(var j=0; j<arr.length; j++){
                if(thisId == arr[j].t_id){ // 找到当前元素
                    tempId = arr[j].t_pid; // 找到父元素
                    parentDOM = afx.$afx(thisId).parentNode;
                    var parentDOMs = parentDOM.parentNode; // 获取同级的ul显示
                    parentDOM.style.display = "block";
                    for(var b=0; b<parentDOMs.childNodes.length; b++){
                        if(parentDOMs.childNodes[b].nodeType == 1){
                            parentDOMs.childNodes[b].style.display = "block";
                        }
                    }
                    parentDOM.className += "show";
                    tempId != 1 ? idShow(tempId) : "";   // 到第一层返回
                }
            }
        }
        idShow(thisId);
    }
}
// 入场动画
function enterAnimation(className){
    if(document.querySelector){
        var boxDOM = document.querySelector("ul."+className);
        var showLi = afx.getTags(boxDOM,"li");
        for(var i=0; i<showLi.length; i++){
            showLi[i].className += "show-animation";
        }
    }
}
// 图片缩放
function listImgZoom(dom,w){
    var listBox = afx.$afx(dom);
    var listImgDOM = afx.getTags(listBox,"img"),
        imgWidth = w,
        imgHeight = Math.round(imgWidth * 0.75);
    for(var i=0; i<listImgDOM.length; i++){
        var img = new Image();
        img.src = listImgDOM[i].src;
        if(img.height/img.width > 0.75){ // 高了,窄了
            listImgDOM[i].height=imgHeight;
            listImgDOM[i].width=imgHeight/img.height * img.width;
            listImgDOM[i].style.padding = 0;
            listImgDOM[i].style.paddingLeft = Math.round((imgWidth-img.width/img.height * imgHeight)/2) + "px";
            listImgDOM[i].style.paddingRight = Math.round((imgWidth-img.width/img.height * imgHeight)/2) + "px";
        }else if(img.height/img.width < 0.75){ // 宽了,矮了
            listImgDOM[i].width = imgWidth;
            listImgDOM[i].height = imgWidth/img.width * img.height;
            listImgDOM[i].style.padding = 0;
            listImgDOM[i].style.paddingTop = Math.round((imgHeight-listImgDOM[i].height)/2) + "px";
            listImgDOM[i].style.paddingBottom = Math.round((imgHeight-listImgDOM[i].height)/2) + "px";
        }else{
            listImgDOM[i].height=imgHeight;
            listImgDOM[i].width=imgWidth;
        }
    }
}