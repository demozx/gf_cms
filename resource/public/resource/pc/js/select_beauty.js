
(function($){
var	$View = function( str ){
		return  $( eval( '\'' + str.replace( /<%=([\w]+)\%>/g, '\' + $1 + \'' ) + '\'' ) );
	},
	createUiId = function( el, prefix ){
		var prefix = ( !prefix && typeof prefix != 'string' )? 'form' : prefix;
		if( el.attr('id') ){
			return prefix + '_ui_id_' + el.attr('id');
		}
		else if( el.attr('name') ){
			return	prefix + '_ui_name_' + el.attr('name')
					.replace('[]', '')
					.replace('[', '_')
					.replace('\']', '')
					.replace(']', '')
					.replace('\'', '')
					.replace('"', '');
		}
		else{
			return prefix + '_ui_the_' + 'n';
		}
	},
	cssSelect = function (config) {
		var optionsPosition;
		if(config){
			optionsPosition = config.position;
		}
		
		if($(this).length <= 0){
			return false;
		}
		return $(this).each(function(){
			var	$thisEl = $(this),
				$optsEl = $thisEl.children('option'),
				$uiEl,
				//thisEvs = $._data($thisEl[0], 'events'),
				//thisEvs = jQuery._data($thisEl[0], 'events'),
				thisEvs = $thisEl.on('events'),
				uiId = createUiId($thisEl, 'select'),
				optsWidth = 0,
                optsPosition,optionsPosition, optionsLimit,
	
				/* Ui */
				tpl = {
					wrapper: '<div class="cssSelect"></div>',
					select: '<div class="selectBox sNormal"></div>',
					selectLt: '<div class="selectLt"></div>',
					selectRt: '<div class="selectRt"></div>',
					options: '<div class="optionsBox"></div>',
					optionsInner: '<div class="optionsInnerBox"></div>',
					option: '<div class="optionBox oNormal"></div>',
					optionInner: '<span></span>'
				},
	
				/* Select box */
				$selectEl = $View(tpl.select),
				$sLtEl = $View(tpl.selectLt),
				$sRtEl = $View(tpl.selectRt),
	
				/* Options box */
				$optionsEl = $View(tpl.options),
				$optionsInnerEl = $View(tpl.optionsInner),
	
				init = function() {

					if(
							$thisEl.length <= 1
						&&	$thisEl.get(0).tagName === 'SELECT'
						&&	!$thisEl.attr('multiple')
						&&	$thisEl.children('optgroup').length <= 0
					)
					{						
						
						$uiEl = $View(tpl.wrapper).attr('id', uiId);
						$('#' + uiId).remove();
                        $thisEl.show();
						$thisEl.hide();

						$uiEl.append($selectEl).append($optionsEl.append($optionsInnerEl));
						$uiEl.click(function(event) {
							event.stopPropagation();
						});
	
						/* Render select*/
						$selectEl.append($sLtEl).append($sRtEl);
						if( $thisEl.attr('disabled') ){
							$selectEl.addClass('sDisabled');
							
							return;
						}

						$selectEl.click(events.selectClick);
						$selectEl.hover(events.selectHover, events.selectNormal);		
	
						/* Render options*/
						$optsEl.each(function(i){
							var	optEl = $(this),
								optionTxt = optEl.text(),
								
								/* Render option */
								$optionEl = $View(tpl.option).append($View(tpl.optionInner).text(optionTxt));
	
								/* Option event */
								$optionEl.css({ float : 'left' });
								$optionEl.hover(events.optionHover, events.optionNormal);
								$optionEl.click(events.optionClick);
	
								/* Render all options*/
								$optionsInnerEl.append($optionEl);
	
							/* Set selected */
							if( optEl.val() == $thisEl.val() ){
								$sLtEl.text(optionTxt);
	
								/* Set select option */
								$optionEl.addClass('selected');
							}
	
						});
	
						/* Options position */
						$thisEl.setOptions = function(){
                            if(config && config.position){
                                optionsPosition = config.position;
                            }

                            if(config && config.limit){
                                optionsLimit = config.limit;
                            }

                            $optionsInnerEl.children('.optionBox').css({ float : 'none' });

							var top,
								sltPosition = $selectEl.position(),
								sltedPosition = $optionsInnerEl.children('.selected').position(),
								eachOptHeight = $optionsInnerEl.outerHeight() / $optsEl.length,
								optsHeight = 'auto';

                            if(optionsLimit){
                                if(optionsLimit != 'auto' && $optsEl.length > optionsLimit){
                                    optsHeight = Math.round(eachOptHeight) * optionsLimit;
                                }
                                else{
									optsHeight = 'auto';
								}
							}
							else{
								if($optsEl.length > 10){
									optsHeight = Math.round(eachOptHeight) * 7;
								}
								else{
									optsHeight = 'auto';
								}
							}

                            sltPosition.top = sltPosition.top + $selectEl.outerHeight();

                            $optionsEl.css({height: Math.round(optsHeight), top: sltPosition.top + 1});
    
                            if(
                                optsPosition &&
                                ( 
                                    $(window).height() + $(document).scrollTop()
                                    <
                                    $optionsEl.offset().top + $optionsEl.outerHeight()
                                )
                            ){
								top = sltPosition.top - $optionsEl.outerHeight() - $selectEl.outerHeight() - 3;
							}else{
								top = sltPosition.top;
							}

                            $optionsEl.scrollTop(Math.round(sltedPosition.top - eachOptHeight));
							$optionsEl.css({ 'top' : top + 1 , 'left' : sltPosition.left, 'overflow-y': 'auto', 'overflow-x': 'hidden'});

                            if( !optsPosition ){
                                optsPosition = $optionsEl.position();
                            }
						};
						
                        /* Render $uiEl */
						$thisEl.before($uiEl);
						$thisEl.setOptions();
                        $optionsEl.hide();
					}
				},
				
				
				events = $.extend({
					selectNormal : function() {
						$(this).removeClass('sHover');
					},
					selectHover : function() {
						$(this).addClass('sHover');
					},
					selectClick : function() {
						if( $optionsEl.css('display') != 'none' ){
							handlers.selectOff();
						}
						else{
							handlers.selectOn();
							$thisEl.setOptions();
						}
					},
					optionNormal : function() {
						$(this).removeClass('oHover');
					},
					optionHover : function() {
						
						$(this).addClass('oHover');
					},
					optionClick : function() {
						
						handlers.selected(this);
					},
					documentClick : function() {
						handlers.selectOff();
					}
				}, function(){}),
				handlers = $.extend({
					selectOn : function(){
						this.selectOff();
	
						$selectEl.addClass('sPressDown');
						$optionsEl.show();

						$(document).one('click', events.documentClick);
					},
					selectOff : function(){
						$thisEl.unbind('click');
						$('.selectBox').removeClass('sPressDown');
						$('.optionsBox').hide();
					},
					selected : function(el){
						var i = $optionsInnerEl.children('.optionBox').index(el);
						$optionsInnerEl.children('.optionBox').removeClass('selected');
						$(el).addClass('selected');
						$sLtEl.text($(el).text());
						if(thisEvs&& thisEvs.change&& thisEvs.change.length > 0){
				
											
							$thisEl[0].selectedIndex = i;
							$.each( thisEvs.change , function(){
								$thisEl.one('click', this.handler);
								$thisEl.click();
							});
				
				
							$thisEl.children().eq(i).attr('selected',true);
							$thisEl.trigger('change');
							
							var iii=$('#identitySelect').children().eq(i).prop('selected',true);
							$.each(thisEvs.change , function(){
								$thisEl.one('click', this.handler);
								$thisEl.click();
							});
						}
						$thisEl[0].selectedIndex = i;
						this.selectOff();
	
					}
				}, function(){});
	
			return init();
		});
	}
	$.fn.extend({
		cssSelect: cssSelect
	});

})(jQuery);