/**
 * Theme functions file
 *
 * Contains handlers for navigation, accessibility, header sizing
 * footer widgets and Featured Content slider
 *
 */
(function($) {
    var body = $('body'), _window = $(window);
// start of script

    $('#main').on("click", function() {
        $('#primary-navigation').removeClass('toggled-on');
    });
    $('.container').on("click", function() {
        $('#primary-navigation').removeClass('toggled-on');
    });

    $("#primary-navigation .menu-toggle").on("click", function() {
        if ($('#primary-navigation').hasClass('toggled-on')) {
            $('#primary-navigation').removeClass('toggled-on');
        } else {
            $('#primary-navigation').addClass('toggled-on');
        }

    });

//load content when scroll down
    if ($('.container-prices-page').length > 0) {
        $(window).scroll(function() {
            if ($(window).scrollTop() + $(window).height() > $('.container-prices-page').offset().top) {
                if (!$('.container-prices-page').hasClass('loaded')) {
                    $.ajax({
                        type: 'GET',
                        data: {action: 'nm_get_page', page_id : 415},
                        url: ajaxurl,
                        beforeSend: function() {
                            $('.container-prices-page').html('<div style="text-align:center; padding: 10px 0;"><b>Loading...</b></div>');
                        },
                        success: function(data) {
                            setTimeout(function() {
                                $('.container-prices-page').html();

                                $('.container-prices-page').fadeIn('slow').html(data);
                                $('.container-prices-page').addClass('loaded');

 				$(".monthly-cost").text(monthly_recurring.toFixed(2));
    				$(".annual-cost").text(annual_recurring.toFixed(2));
    				$(".perpetual-cost").text(perpetual_signup.toFixed(2));

                                $('.monthly_subs_block').css({'cursor': 'pointer'});
                                $('.monthly_subs_block').on("click", function() {
                                    location.href = '/subscription/monthly/';
                                });

                                $('.annual_subs_block').css({'cursor': 'pointer'});
                                $('.annual_subs_block').on("click", function() {
                                    location.href = '/subscription/annual/';
                                });

                                $('.perpetual_subs_block').css({'cursor': 'pointer'});
                                $('.perpetual_subs_block').on("click", function() {
                                    location.href = '/subscription/perpetual/';
                                });

                                $('#lite_subs_block').css({'cursor': 'pointer'});

                                $('#lite_subs_block').on("click", function() {
                                    location.href = '/subscription/lite/';
                                });

                                $('#monthly_block_addcart').css({'cursor': 'pointer'});
                                $('#monthly_block_addcart').on("click", function() {
                                    location.href = '/subscription/monthly';
                                });

                                $('#annual_block_addcart').css({'cursor': 'pointer'});
                                $('#annual_block_addcart').on("click", function() {
                                    location.href = '/subscription/annual';
                                });

                                $('#perpetual_block_addcart').css({'cursor': 'pointer'});
                                $('#perpetual_block_addcart').on("click", function() {
                                    location.href = '/subscription/perpetual';
                                });
                            },100);
                        }
                    });
                }
            }
        });
    }
//end Script

    // Enable menu toggle for small screens.
    /*	( function() {
     var nav = $( '#primary-navigation' ), button, menu;
     if ( ! nav ) {
     return;
     }
     
     button = nav.find( '.menu-toggle' );
     if ( ! button ) {
     return;
     }
     
     // Hide button if menu is missing or empty.
     menu = nav.find( '.nav-menu' );
     if ( ! menu || ! menu.children().length ) {
     button.hide();
     return;
     }
     
     $( '.menu-toggle' ).on( 'click.novamind', function() {
     nav.toggleClass( 'toggled-on' );
     } );
     } )();
     */
    /*
     * Makes "skip to content" link work correctly in IE9 and Chrome for better
     * accessibility.
     *
     * @link http://www.nczonline.net/blog/2013/01/15/fixing-skip-to-content-links/
     */
    _window.on('hashchange.novamind', function() {
        var element = document.getElementById(location.hash.substring(1));

        if (element) {
            if (!/^(?:a|select|input|button|textarea)$/i.test(element.tagName)) {
                element.tabIndex = -1;
            }

            element.focus();

            // Repositions the window on jump-to-anchor to account for header height.
            window.scrollBy(0, -80);
        }
    });

    $(function() {
        // Search toggle.
        $('.search-toggle').on('click.novamind', function(event) {
            var that = $(this),
                    wrapper = $('.search-box-wrapper');

            that.toggleClass('active');
            wrapper.toggleClass('hide');

            if (that.is('.active') || $('.search-toggle .screen-reader-text')[0] === event.target) {
                wrapper.find('.search-field').focus();
            }
        });
	
	var width = $('.field-search').width();
	//$('.field-search').hide();
 	$('.field-search').css('width', '30px');
	$('.toggle-search').on('click', function (event){
		event.preventDefault();
		 if ($('.field-search').hasClass("search") ){
			$('.field-search').animate({width: '30px'}, 400);
			$('.field-search').removeClass("search");
		}else {
		    	$('.field-search').animate({ width: width }, 400);
			$('.field-search').addClass("search");
			$('.field-search').focus();
		}		
	});
	$('.field-search').on('blur', function (){
		if($(this).val()=='' && $(this).hasClass("search")) $('.toggle-search').click();
	});
        /*
         * Fixed header for large screen.
         * If the header becomes more than 48px tall, unfix the header.
         *
         * The callback on the scroll event is only added if there is a header
         * image and we are not on mobile.
         */
        if (_window.width() > 781) {
            var mastheadHeight = $('#masthead').height(),
                    toolbarOffset, mastheadOffset;

            if (mastheadHeight > 48) {
                body.removeClass('masthead-fixed');
            }

            if (body.is('.header-image')) {
                toolbarOffset = body.is('.admin-bar') ? $('#wpadminbar').height() : 0;
                mastheadOffset = $('#masthead').offset().top - toolbarOffset;

                _window.on('scroll.novamind', function() {
                    if ((window.scrollY > mastheadOffset) && (mastheadHeight < 49)) {
                        body.addClass('masthead-fixed');
                    } else {
                        body.removeClass('masthead-fixed');
                    }
                });
            }
        }

        // Focus styles for menus.
        $('.primary-navigation, .secondary-navigation').find('a').on('focus.novamind blur.novamind', function() {
            $(this).parents().toggleClass('focus');
        });
    });

    _window.load(function() {
        // Arrange footer widgets vertically.
        if ($.isFunction($.fn.masonry)) {
            $('#footer-sidebar').masonry({
                itemSelector: '.widget',
                columnWidth: function(containerWidth) {
                    return containerWidth / 4;
                },
                gutterWidth: 0,
                isResizable: true,
                isRTL: $('body').is('.rtl')
            });
        }

        // Initialize Featured Content slider.
        if (body.is('.slider')) {
            $('.featured-content').featuredslider({
                selector: '.featured-content-inner > article',
                controlsContainer: '.featured-content'
            });
        }
    });
})(jQuery);

//Init tool tip
$('a[data-toggle="tooltip"]').tooltip();

$('body').on('keyup', "#query", function(event) {
    if (event.keyCode == 13) {
        $(".btnsupportsearch").click();
    }
});
$('body').on('keypress', '.clearable', function() {
    if ($(this).val() != '' && $('#cleare-search').length == 0)
        $(this).after('<span id="cleare-search" class="glyphicon glyphicon-remove-circle" style="position:absolute; right:16px; padding: 10px;"></span>');
})
$('body').on('click', '#cleare-search', function() {
    $('.clearable').val('');
    $(this).remove()
})
