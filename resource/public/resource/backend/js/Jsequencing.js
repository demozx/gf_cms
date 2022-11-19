/*
 *图片排序及查看
 *j -- 2017-12-20
*/
function Jsequencing(options){
	var defaults={
		// 页面图片列表ID
		listid:"img_ul",
		// 列表图片前缀
		thumbherf:"",
		// 原图前缀
		bigherf:options.thumbherf,
		// 图片数据数组
		imgsrcarr:[],
		// 是否json格式数据
		jsondata:false,
		// 预览/查看图片
		viewimg:true,
		// 预览切换
		view_toggle:true,
		// 预览缩放
		view_zoom:true,
		// 预览旋转
		view_rotate:true,
		showtitle:true,
		// item
		IdClass:"__item__",
	};
	var opts = $.extend(defaults,options);
	if((typeof opts.imgsrcarr[0]=='string')&&opts.imgsrcarr[0].constructor==String){
		opts.jsondata=false;
	}else if((typeof opts.imgsrcarr[0]=='object')&&opts.imgsrcarr[0].constructor==Object){
		opts.jsondata=true;
	}else if(opts.imgsrcarr[0]!==undefined){
		alert("数据格式不正确")
	}
	if(!opts.jsondata) opts.showtitle=false;
	//-//-//-//-开始-//-//-//-//
	var box=this.box = $("#"+opts.listid);
	box.append( "<div class='itemlist' style='height:100%'></div>"+
							"<div class='ident'></div>"+
							"<div class='morexy'></div>"
	);
	var itemlist = box.find(".itemlist");
	// item内容
	this.itemhtml=function(imgsrc,imgtitle){
		// var titlebox = opts.showtitle ? '<div class="textbox">'+imgtitle+'</div>' : "";
		var titlebox = opts.showtitle ? '<div class="textbox">拖动进行排序</div>' : "";
		return '<input lay-ignore class="checkbox" name="" type="checkbox" value="">'+
					 '<div class="picbox">'+
						'<a class="viewimg" href="'+opts.bigherf+imgsrc+'" title="'+imgtitle+'">'+
							'<img src="'+opts.thumbherf+imgsrc+'" ondragstart="return false;" />'+
						'</a>'+
					 '</div>'+titlebox
	};
	// 页面dom绘制
	var html='';
	for(var j=0;j<opts.imgsrcarr.length;j++){
		var vimgsrc=opts.showtitle ? opts.imgsrcarr[j].src : opts.imgsrcarr[j];
		var vimgtitle=opts.showtitle ? opts.imgsrcarr[j].title : "";
		html+='<div class="item '+opts.IdClass+'" id="'+opts.listid+'_item'+j+'" item="'+j+'">'+this.itemhtml(vimgsrc,vimgtitle)+'</div>';
	}
	itemlist.html(html);
	// 取总宽度
	var box_w = box.width();
	// 页面初始化配置
	var item_w,item_h,col_len,few_len,imgnum,box_h;
	this.info=function(fun){
		if(item_w===undefined || item_w===null){
			// 每个item占的横向位置
			item_w = itemlist.find("."+opts.IdClass).outerWidth(true);
			// 每个item占的竖向位置
			item_h = itemlist.find("."+opts.IdClass).outerHeight(true);
			box.find(".ident").css({height:item_h+"px"});
			box.find(".morexy").css({width:item_w+"px",height:item_h+"px"});
		};
		// item数量
		imgnum = itemlist.find("."+opts.IdClass).length;
		// 共分多少列
		col_len = Math.floor(box_w/item_w);
		// 共分多少行
		few_len = Math.ceil(imgnum/col_len);
		box_h=item_h*few_len+20;
		box.height(box_h+"px");
		return true;
	}.bind(this)
	// 绘制/移动
	var draw=this.draw=function(dom,col,few,slidetime){
		// console.log(dom,col,few,slidetime);
		dom.css({
			"transition-duration": slidetime+"ms",
			"transform":"translate("+col+"px,"+few+"px)",
		});
	}
	// 计算位置
	this.computat=function(index,domid,slidetime){
		var item = $("#"+domid);
		item.attr({
			"item":index,
		});
		if(chekobj[domid]!==undefined){
			chekobj[domid]=index;
		}
		var col_aliquot=index%col_len;
		var row_aliquot=Math.floor(index/col_len);
		var index_col = item_w*(col_aliquot);
		var index_few = item_h*row_aliquot;
		this.draw(item,index_col,index_few,slidetime);
		item.attr({
			"col":index_col,
			"few":index_few
		})
	}.bind(this)
	// 调用绘制
	var redraw=this.redraw=function(strati,ilen,slidetime){
		for(var j=strati;j<strati+ilen;j++){
			this.computat(j,opts.listid+"_item"+itemidarr[j],slidetime);
			window.getImagesArr();
		}
//		for(var i=strati;i < strati+ilen;i++){
//			this.computat(itemidarr[i],opts.listid+"_item"+itemidarr[i],slidetime);
//		}
	}.bind(this)
	// 首次绘制
	if(this.info(this)){
		var chekobj={};
		var itemidarr=[];
		// 新数组
		this.imgnewarr=[];
		for(var i=0;i<imgnum;i++){ itemidarr.push(i) };
		this.redraw(0,itemidarr.length,0);
	}
	// 浏览器尺寸改变时
	$(window).resize(function(){
		// 取总宽度
		box_w = box.width();
		if(this.info(this)){
			this.redraw(0,itemidarr.length,0);
		}
	}.bind(this));
	// 批量删除
	this.datadel=function(){
		if(!($.isEmptyObject(chekobj))){
			for(indexi in chekobj){
				var itemi=parseInt(itemlist.find("#"+indexi).attr("item"));
				var arr2=[];
				for(var i=0;i<itemidarr.length;i++){
					if(i==itemi){
						itemidarr[i]=null;
					}else if(itemidarr[i]!=null){
						arr2.push(itemidarr[i]);
					};
				}
				itemlist.find("#"+indexi).remove();
				delete chekobj.indexi;
			}
			itemidarr=arr2;
		}
		this.redraw(0,itemidarr.length,200);
	}.bind(this)
	// 清空item
	this.dataempty=function(){
		itemlist.find("."+opts.IdClass).remove();
		opts.imgsrcarr.length = 0;
		itemidarr.length = 0;
		chekobj={};
		box_h="20";
		box.height(box_h+"px");
	}.bind(this)
	// 添加图片
	this.addimg=function(imgsrc,imgtitle){
		if(opts.jsondata){
			opts.imgsrcarr.push({src:imgsrc,title:imgtitle});
		}else{
			opts.imgsrcarr.push(imgsrc);
		}
		var imgnum=opts.imgsrcarr.length-1;
		itemidarr.push(imgnum);
		var titlebox = opts.showtitle ? '<div class="textbox">'+vimgtitle+'</div>' : "";
		if(imgtitle===undefined) imgtitle="";
		itemlist.append(
			'<div class="item '+opts.IdClass+'" id="'+opts.listid+'_item'+imgnum+'" item="'+imgnum+'">'+this.itemhtml(imgsrc,imgtitle)+'</div>'
		);
		this.redraw(itemidarr.length-1,1,0);
		this.info(this);
	}.bind(this)
	// 添加图片数组
	this.addimgarr=function(imgobj){
		if(Object.prototype.toString.call(imgobj)=='[object Array]'){
			if(opts.jsondata){
				if(!((typeof imgobj[0]=='object')&&imgobj[0].constructor==Object)){
					alert("数据格式不正确,请传入json格式数据!");
					return;
				}
			}else{
				if(!((typeof imgobj[0]=='string')&&imgobj[0].constructor==String)){
					alert("数据格式不正确,请传入字符串格式数据!");
					return;
				}
			}
			for(var i=0;i<imgobj.length;i++){
				if(opts.jsondata){
					if(imgobj[i].title===undefined) imgobj[i].title="";
					this.addimg(imgobj[i].src,imgobj[i].title);
				}else{
					this.addimg(imgobj[i]);
				}
			}
		}
	}.bind(this)
	// 获取最新数组
	this.getnewarr=function(){
		this.imgnewarr.length = 0;
		for(var i=0;i<itemidarr.length;i++){
			this.imgnewarr.push(opts.imgsrcarr[itemidarr[i]])
		}
		console.log(this.imgnewarr);
		return this.imgnewarr;
	}.bind(this)
	// 选中input
	itemlist.on("click","input.checkbox",function(){
			var thisid=$(this).parents("."+opts.IdClass).attr("id");
			if($(this).prop("checked")){
				var index=parseInt($(this).parents("."+opts.IdClass).attr("item"));
				chekobj[thisid]=index;
			}else{
				delete chekobj[thisid];
			}
		}
	)
	
	//-//-//-//-拖动排序-//-//-//-//
	// 鼠标初始位置
	var startX = null;
	var startY = null;
	// 选中元素item
	var moveDom = null;
	var checkItem = null;
	// 初始index/当前点击元素的item值
	var startIndex = null;
	// 是否多选框选中
	var isSelected = false;
	// 初始坐标位置
	var oldCol = null;
	var oldFew = null;
	// 标识
	var identDom = box.find(".ident");
	// 多选
	var multipleBox = box.find(".morexy");
	// 是否拖动
	var isDrag = false;
	// 标识初始位置
	var identMobieX = null;
	var identMobieY = null;
	// 初始化
	function mouseinit(){
		moveDom = null;
		isDrag = false;
		checkItem.css({"opacity":"1","z-index":""});
		multipleBox.hide();
		identDom.hide();
	}
	// 获取选中元素id
	function getchekobj(){
		var domidarr=[];
		for(var key in chekobj){
			domidarr.push("#"+key);
		}
		var chekDomObj = $(domidarr.join(','));
		if(domidarr.length == 0){
			chekDomObj = checkItem;
		}
		return chekDomObj;
	}
	// 松开手指返回原位
	$(document).on("mouseup",function(){
		if(moveDom != null){
			draw(moveDom,oldCol,oldFew,0);
			mouseinit();
			if(isSelected){
				for(arri in chekobj){
					var elem = 	$("#"+arri);
					draw(elem,elem.attr("col"),elem.attr("few"),300);
				}
			}
		}
	})
	// 拖动方法
	box.on({
		mousedown:function(e){
			e.preventDefault();
			if(e.target.localName=="input"){
				return false;
			}
			
			var isIdclass = $(e.target).hasClass(opts.IdClass);
			var pIdClass = $(e.target).parents("."+opts.IdClass)[0];
			
			if(!isIdclass && pIdClass==undefined){
				return false;
			}
			
			if(isIdclass){
				checkItem = $(e.currentTarget);
			}else{
				checkItem = $(pIdClass);
			}
			
			if(checkItem.find("input.checkbox").prop("checked")){
				isSelected = true;
				moveDom = multipleBox;
			}else{
				isSelected = false;
				moveDom = checkItem;
			}
			
			startX = e.pageX;
			startY = e.pageY;
			startIndex=parseInt(checkItem.attr("item"));
			oldCol = parseInt(checkItem.attr("col"));
			oldFew = parseInt(checkItem.attr("few"));
			
			checkItem = getchekobj();
			checkItem.css({"opacity":"0.8","z-index":"10"});
			
			draw(identDom,oldCol,oldFew,0);
		},
		mousemove:function(e){
			if(moveDom == null){
				return;
			}
			if(e.which != 1){
				return;
			}
			// 拖动距离
			var gapX=e.pageX-startX;
			var gapY=e.pageY-startY;
			// 偏移距离
			var mobiex=oldCol+gapX;
			var mobiey=oldFew+gapY;
			// 横向是否超出边界
			if(mobiex>box_w-item_w){
				mobiex=box_w-item_w;
			}else if(mobiex < 0){
				mobiex=0;
			}
			// 竖向是否超出边界
			if(mobiey>box_h-item_h){
				mobiey=box_h-item_h;
			}else if(mobiey < 0){
				mobiey=0;
			}
			if(Math.abs(gapX)>10 || Math.abs(gapY)>10){
				if(!isDrag){
					if(isSelected){
						draw(moveDom,oldCol,oldFew,0);
						for(arri in chekobj){
							$("#"+arri).css({"opacity":"0.25","z-index":"10"});
							draw($("#"+arri),oldCol,oldFew,300);
						}
						moveDom.show();
					}
				}
				
				// item移动
				draw(moveDom,mobiex,mobiey,0);
				// 标识移动
				identMobieX=Math.abs(Math.ceil((mobiex-item_w/2)/item_w));
				identMobieY=Math.abs(Math.ceil((mobiey-item_h/2)/item_h));
				identDom.show();
				draw(identDom,identMobieX*item_w,identMobieY*item_h,0);
				isDrag = true;
			}
		},
		mouseup:function(e){
			if(moveDom == null){
				return;
			}else if(isDrag){
				draw(moveDom,oldCol,oldFew,0);
				var toposion = identMobieY*col_len+identMobieX;
				
				// 多选
				if(isSelected){
					var min=itemidarr.length;
					var max=0;
					var transsh=[];
					for(indexi in chekobj){
						itemlist.find("#"+indexi).css({"opacity":"1","z-index":""});
						for(var i=0;i<itemidarr.length;i++){
							if(i==chekobj[indexi]){
								min = i < min ? i : min;
								max = i+1 > max ? i+1 : max;
								transsh.push(itemidarr[i]);
								itemidarr[i]=null;
							};
						};
					}
					for(var i=0;i<transsh.length;i++){
						itemidarr.splice(toposion+i,0,transsh[i]);
					}
					// transsh.length=0;
					for(var i=itemidarr.length-1;i>=0;i--){
						if(itemidarr[i]==null){
							itemidarr.splice(i,1);
						};
					}
					min = toposion < min ? toposion : min;
					max = toposion > max ? toposion : max;
					if(max-min != transsh.length){
						redraw(min,max-min,200);
					}else{
						for(arri in chekobj){
							var elem = 	$("#"+arri);
							draw(elem,elem.attr("col"),elem.attr("few"),300);
						}
					}
				}else{
					var difference=toposion-startIndex;
					if(difference > 1){
						// 往后
						var changesitem=itemidarr.splice(startIndex,1)[0];
						itemidarr.splice(toposion-1,0,changesitem);
						redraw(startIndex,startIndex+difference,200);
					}else if(difference < 0){
						// 往前
						var changesitem=itemidarr.splice(startIndex,1)[0];
						itemidarr.splice(toposion,0,changesitem);
						difference=Math.abs(difference)+1;
						redraw(toposion,toposion+difference,200);
					}
				}
				mouseinit();
			}
		},
		mouseleave:function(e){
			if(moveDom == null){
				return;
			}
				// mouseinit();
		},
	})
	
	//-//-//-//-预览/查看图片-//-//-//-//
	var ismove=false;
	var eimgx,eimgy;
	itemlist.on({
		click:function(e){
			e.preventDefault();
			if(!ismove){
				var item=parseInt($(this).parents("."+opts.IdClass).attr("item"));
				chakshow(item)
			};
		},
		mousedown:function(e){
			eimgx=e.pageX;
			eimgy=e.pageY;
			ismove=false;
		},
		mouseup:function(e){
			eimgx=Math.abs(e.pageX-eimgx);
			eimgy=Math.abs(e.pageY-eimgy);
			ismove=false;
			if(eimgx > 5 || eimgy > 5) ismove=true;
		},
		mouseleave:function(){
			ismove=false;
		},
	},"a");
	//查看图片dom添加
	if(opts.viewimg){
		var chak_btn="";
		if(opts.view_toggle) chak_btn+='<span class="btn chak_prev">上一个</span><span class="btn chak_next">下一个</span>';
		if(opts.view_rotate) chak_btn+='<span class="btn chak_turn_l">向左转</span><span class="btn chak_turn_r">向右转</span>';
		chak_btn+='<span class="btn chak_close">关闭</span>';
		box.append(
			'<div class="chak_box">'+
				'<div class="chak_title"></div>'+
				'<div class="chak_img">'+
					'<img draggable="false" src="" ondragstart="return false;" />'+
				'<p></p></div>'+
				'<div class="chak_btn">'+chak_btn+'</div>'+
			'</div>'
		);
		var chak_box=box.find(".chak_box");
		chak_box.on({click:function(e){
			chakhide();
		}});
		chak_box.find(".chak_img > img").on({click:function(e){
			e.stopPropagation();
		}});
		if(opts.view_toggle){//是否可切换图片
			chak_box.find(".chak_prev").on({click:function(e){
				e.stopPropagation();
				var item = parseInt(chak_box.find(".chak_img").attr("item"))-1;
				chakshow(item);
			}});
			chak_box.find(".chak_next").on({click:function(e){
				e.stopPropagation();
				var item = parseInt(chak_box.find(".chak_img").attr("item"))+1;
				chakshow(item);
			}});
		};
		//查看图片
		var isview=false;
		var pisshow=false;
		var chakshow=function(imgitem){
			isview=true;
			var chak_ptog=function(text){
				var chak_p=chak_box.find(".chak_img > p");
				if(!pisshow){
					pisshow=true;
					chak_p.html(text).show();
					setTimeout(function(){
						pisshow=false;
						chak_p.hide();
					},1000);
				}
			}
			if(imgitem<0){
				chak_ptog("已经是第一张图片");
				return;
			}else if(imgitem > itemidarr.length-1){
				chak_ptog("已经是最后一张图片");
				return;
			}else{
				chak_box.find(".chak_img > p").hide()
			}
			var viewimgdom=itemlist.find("."+opts.IdClass+"[item='"+imgitem+"'] .viewimg")
			var imgsrc = viewimgdom.attr("href");
			var imgtitle=viewimgdom.attr("title");
			chak_box.find(".chak_img").attr("item",imgitem);
			chak_box.find(".chak_img > img").attr({"src":imgsrc,"style":""});
			chak_box.find(".chak_title").html(imgtitle);
			$("body").css({"overflow":"hidden"});
			$("body",parent.document).css({"overflow":"hidden"});
			chak_box.show();
		};
		//关闭图片查看
		var chakhide=this.chakhide=function(){
			isview=false;
			chak_box.find(".chak_img").attr("item","");
			chak_box.find(".chak_img > img").attr("src","");
			$("body").css({"overflow":"auto"});
			chak_box.hide();
		};
		//预览图片缩放
		if(opts.view_zoom){
			var viewimg_zoom=function(solls,zoomval){//缩放
				var imgdom = chak_box.find(".chak_img > img");
				var img = new Image(); 
				img.src =imgdom.attr("src"); 
				var imgWidth = img.width; //图片实际宽度
		    var ckimg_w=parseInt(imgdom.width());
		    imgdom.css({"max-width":imgWidth+"px","max-height":"none"});
		    solls>0 ? ckimg_w+=zoomval : ckimg_w-=zoomval;
		    imgdom.css({"width":ckimg_w+"px"});
			};
			chak_box.find(".chak_img > img").on({//拖动
				click:function(event){
					event.stopPropagation();
				},
			  mousedown:function(e){
					viewmove=true;
					var marginx=parseInt($(this).css("margin-left"));
					var marginy=parseInt($(this).css("margin-top"));
					var _x=e.pageX-marginx; 
					var _y=e.pageY-marginy;
					$(this).off("mousemove").on({
						mousemove:function(e){
							if(viewmove){ 
								var x=e.pageX-_x;
								var y=e.pageY-_y;
								$(this).css({
									"margin-top":y+"px",
									"margin-left":x+"px"
								});
							} 
						}
					});
				},
				mouseup:function(e){
					viewmove=false;
				},
				mouseout:function(e){
					viewmove=false;
				},
			});
			//滚轮事件添加
			$(document).on('mousewheel', function(event){
				if(isview){
			    var solls=event.originalEvent.wheelDelta;
			    viewimg_zoom(solls,50);
				}
			});
			document.addEventListener("DOMMouseScroll", function (event) {//ff
				if(isview){
			    var solls=event.detail;
			    solls=0-solls;
				  viewimg_zoom(solls,50);
				}
			});
		};
		//预览图片旋转
		if(opts.view_rotate){
			//旋转角度
			function getmatrix(a,b,c,d,e,f){
				var aa=Math.round(180*Math.asin(a)/ Math.PI);  
				var bb=Math.round(180*Math.acos(b)/ Math.PI);  
				var cc=Math.round(180*Math.asin(c)/ Math.PI);  
				var dd=Math.round(180*Math.acos(d)/ Math.PI);  
				var deg=0;
				if(aa==bb||-aa==bb){  
				    deg=dd;  
				}else if(-aa+bb==180){  
				    deg=180+cc;  
				}else if(aa+bb==180){  
					deg=360-cc||360-dd;  
				}  
				return deg>=360?0:deg;
			}
			var chak_turn=function(dom,step){
				var _img=chak_box.find(".chak_img > img");
		    var deg=eval('get'+_img.css('transform'));
		    _img.css({'transform':'translate(-50%,-50%) rotate('+(deg+step)%360+'deg)'});
				dom.attr({"disabled":"true"});
				setTimeout(function(){
					dom.removeAttr("disabled");
				},60)
			}
			//向左转
			chak_box.find(".chak_turn_l").on({click:function(e){
				e.stopPropagation();
				chak_turn($(this),-90)
			}});
			//向右转
			chak_box.find(".chak_turn_r").on({click:function(e){
				e.stopPropagation();
				chak_turn($(this),90)
			}});
		}
	}
}
