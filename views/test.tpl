<!-- 测试时间轴 -->
<!doctype html>
<html lang="en-US">
<head>
  <meta charset="utf-8">
  <meta http-equiv="Content-Type" content="text/html">
  <title>Vertical Responsive Timeline UI - Template Monster Demo</title>
  <meta name="author" content="Jake Rocheleau">
  <link rel="shortcut icon" href="http://static.tmimgcdn.com/img/favicon.ico">
  <link rel="icon" href="http://static.tmimgcdn.com/img/favicon.ico">
  <link rel="stylesheet" type="text/css" media="all" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" media="all" href="/static/css/bootstrap-glyphicons.css">
  <link rel="stylesheet" type="text/css" media="all" href="/static/css/timeline.css">

<link href='https://fonts.googleapis.com/css?family=Open+Sans:400,700' rel='stylesheet' type='text/css'>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/css/font-awesome.min.css">

  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 	<!-- <style type="text/css">
		.timeline-container{
		  width:100%;
		  /*height:auto;*/
		}
		.timeline{
		  width:100%;
		  margin:auto;
		  position:relative;
		  height: auto;
		  padding:30px 0px 0px 0px;
		}
		.timeline > .timeline-block{
		  width:50%;
		}
		.timeline-block > .popover{
		  max-width: 70%;
		}
		.timeline > .timeline-block-odd{
		  float:left;
		}
		.timeline > .timeline-block-even{
		  float:right;
		}
		.timeline-block-odd>.timeline-content{
		  display: block;
		  position:relative;
		  float:right;
		  margin-right:50px;
		}
		.timeline-block-even>.timeline-content{
		  display: block;
		  float:left;
		  position:relative;
		  margin-left:50px;
		}
		
		.timeline-content > .arrow{
		  top:20px !important;
		  background: #f7f7f7;
		}
		.timeline-content > .popover-content{
		  max-height: 240px;
		  overflow-y: auto;
		  color: #696969;
		  font-size: 11px;
		}
		
		.timeline-content>.popover-footer{
		  padding:8px 14px;
		  margin:0px;
		  font-size: 11px;
		  border-top: 1px dotted #b9b9b9;
		  color: #4b4b4b;
		  background:#F2F2F2;
		}
		.timeline img{
		  border:1px solid #dddddd;
		  padding:2px;
		}
		.timeline-img{
		  width:50px;
		  height:50px;
		  -webkit-border-radius: 25px;
		  -moz-border-radius: 25px;
		  -o-border-radius: 25px;
		  border-radius: 25px;
		  border:3px solid #A3E3E8;
		  margin:auto;
		  background: #ffffff;
		  position:absolute;
		  z-index: 9;
		  left:50%;
		  margin-left:-25px;
		  background: url(../../static/img/eye.png);
		  background-size: cover;
		}
		
		.timeline-line{
		  width:4px;
		  height:100%;
		  background: #aef0f5;
		  position:absolute;
		  z-index:8;
		  left:50%;
		  border-left:1px solid #aef0f5;
		  border-right:1px solid #ffffff;
		  margin-left: -2px;
		  margin-top:-30px;
		}
		
		.timeline-block-odd>.popover.left>.arrow:after {
		  right: 2px;
		  border-left-color: #F7F7F7;
		}
		.timeline-block-even>.popover.right>.arrow:after {
		  left: 2px;
		  border-right-color: #F7F7F7;
		}
		
		
		/** mediaquery查询样式 **/
		@media screen and (max-width: 560px){
		
		  .timeline{
		    width:100%;
		    position:relative;
		    height: auto;
		    padding:30px 0px 0px 0px;
		  }
		  .timeline > .timeline-block{
		    width:100%;
		  }
		  .timeline > .timeline-block-odd{
		    float:right;
		  }
		  .timeline-block-odd>.timeline-content{
		    display: block;
		    position:relative;
		    float:left;
		    margin-left:75px;
		  }
		  .timeline-block-even>.timeline-content{
		    display: block;
		    position:relative;
		    float:left;
		    margin-left:75px;
		  }
		
		  .timeline-block-odd>.popover>.arrow, .timeline-block-odd>.popover>.arrow:		after {
		    position: absolute;
		    display: block;
		    width: 0;
		    height: 0;
		    border-color: transparent;
		    border-style: solid;
		  }
		  .timeline-block-odd>.popover.left>.arrow {
		    left: -21px;
		    bottom: -10px;
		    content: " ";
		    border-left-width: 0;
		    border-right-color: #999;
		    border-width: 10px;
		  }
		
		  .timeline-block-odd>.popover.left>.arrow:after {
		    left:1px;
		    right: 1px;
		    bottom: -10px;
		    content: " ";
		    border-left-width: 0;
		    border-right-color: #fff;
		  }
		  .timeline-block-odd>.popover>.arrow:after {
		    content: "";
		    border-width: 10px;
		  }
		
		  .timeline-img{
		    width:50px;
		    height:50px;
		    margin:auto;
		    background: #ffffff;
		    -webkit-border-radius: 25px;
		    -moz-border-radius: 25px;
		    -o-border-radius: 25px;
		    border-radius: 25px;
		    border:3px solid #8e8e8e;
		    position:absolute;
		    z-index: 9;
		    left:0;
		    margin-left:0px;
		  }
		
		  .timeline-line{
		    width:4px;
		    height:100%;
		    background: #d0d0d0;
		    border-left:1px solid #ececec;
		    border-right:1px solid #ececec;
		    position:absolute;
		    z-index:8;
		    left:0;
		    margin-left: 24px;
		    margin-top:-30px;
		  }
		}
		
		}
	</style> -->

<body>
	<div class="container">
	  <header class="page-header">
	    <h1>Dark Responsive Timeline with Bootstrap</h1>
	    <h1><small>黑色Bootstrap响应式时间轴</small></h1>
	  </header>
	  
	  <ul class="timeline">
	    <!-- <li><div class="tldate">Apr 2014</div></li> -->
	    
	    <!-- <li class="timeline-inverted"> 
      	<div class="timeline-icon timeline-icon-hide-border"><i style="color:#c23b22" class="fa fa-bell-o fa-lg"></i></div>

      	<div class="timeline-panel">
      	  <div class="tl-heading">
      	    <h4>Surprising Headline Right Here</h4>
      	    <p><small class="text-muted"><i class="glyphicon glyphicon-time"></i> 3 hours ago</small></p>
      	  </div>
      	  <div class="tl-body">
      	    <p>Lorem Ipsum and such.</p>
      	  </div>
      	</div>
    	</li> -->
    
    	<!-- <li class="timeline-inverted">
      	<div class="tl-circ"></div>
      	<div class="timeline-panel">
      	  <div class="tl-heading">
          <h4>Breaking into Spring!</h4>
          <p><small class="text-muted"><i class="glyphicon glyphicon-time"></i> 4/07/2014</small></p>
      	  </div>
      	  <div class="tl-body">
      	    <p>Hope the weather gets a bit nicer...</p>
      	    <p>Y'know, with more sunlight.</p>
      	  </div>
      	</div>
    	</li> -->
  	</ul>
	</div>


<!-- <div class="timeline-container"></div> -->

<!-- <div class="timeline-container">
	<div class="row">
		<div class="timeline">
			<div class="timeline-block timeline-block-odd">
				<div class="popover timeline-content left" style="left: 0px;">
					<div class="arrow"></div>
					<h3 class="popover-title">2017/03/18</h3>
					<div class="popover-content">
						<div style="margin-bottom: 10px;">
							<img src="/static/img/1.jpg" style="width: 60px; height: 60px; margin-right: 10px;">
							<img src="/static/img/2.jpg" style="width: 60px; height: 60px; margin-right: 10px;">
							<img src="/static/img/3.jpg" style="width: 60px; height: 60px; margin-right: 10px;"></div>
						<table class="table table-bordered table-condensed">
							<tbody>
								<tr>
									<td nowrap="">眼象特征</td>
									<td>通过图像识别获得眼像特征</td>
								</tr>
								<tr>
									<td nowrap="">匹配结果</td>
									<td>知识库自动获取的饼子</td>
								</tr>
								<tr>
									<td nowrap="">结论说明</td>
									<td>
										对综合揭露进行行详细描述
									</td>
								</tr>
							</tbody>
						</table>
					</div>
					<div class="popover-footer">根据病症信息分析结果</div>
				</div>
			</div>
			<div class="timeline-img"></div>
			<div class="timeline-line"></div>
			<div class="clearfix"></div>
		</div>
	</div>
</div> -->

<script type="text/javascript">
	$(function(){
  	var _timeline_date_ = $("<li><div class='tldate'>Apr 2017<div><li>");
  	$(".timeline").append(_timeline_date_);
  	var loadData=function(){
    	$.getJSON("/test", function (data) {
      	$.each(data, function (i, tl) {
					if(i%2==1){
						var _timeline_invert_ = $("<li></li>");
					}else{
						var _timeline_invert_ = $("<li></li>").addClass("timeline-inverted");
					}

        	$(".timeline").append(_timeline_invert_);

        	var _timeline_icon_ = $("<div></div>").addClass("timeline-icon timeline-icon-hide-border");
        	var _timeline_fa_ = $("<i style='color:#c23b22'></i>").addClass("fa fa-github fa-lg");
        	_timeline_icon_.append(_timeline_fa_);
        	_timeline_invert_.append(_timeline_icon_);
        	/**
        	 * 设置显示内容
        	 */

					var _timeline_panel_ = $("<div></div>").addClass("timeline-panel");
					var _timeline_head_ = $("<div></div>").addClass("tl-heading");
					var _timeline_body_ = $("<div></div>").addClass("tl-body");

					var _timeline_body_p1_ = $("<p></p>").text(tl["feature"]);
					_timeline_body_.append(_timeline_body_p1_);
					var _timeline_body_p2_ = $("<p></p>");
					// var _timeline_body_p2_img_ = $("<img>");
					_timeline_body_p2_.append($("<img/>").attr("src",tl["images"][0]));
					// <p><img src="http://lorempixel.com/600/300/nightlife/" alt="lorem pixel"></p>
					_timeline_body_.append(_timeline_body_p1_);
					_timeline_body_.append(_timeline_body_p2_);

        	var _timeline_head_h4_ = $("<h4></h4>").text(tl["diagDoc"]);
					var _timeline_head_p_ = $("<p></p>");
					var _timeline_head_p_small_ = $("<small class='text-muted'></small>");
					var _timeline_head_p_small_i_ = $("<i class='glyphicon glyphicon-time'></i>");
					_timeline_head_p_small_.append(_timeline_head_p_small_i_);
					_timeline_head_p_small_.append(tl["diagTime"]);

					_timeline_head_p_.append(_timeline_head_p_small_);

					_timeline_head_.append(_timeline_head_h4_);
					_timeline_head_.append(_timeline_head_p_);

					_timeline_panel_.append(_timeline_head_);
					_timeline_panel_.append(_timeline_body_);

					_timeline_invert_.append(_timeline_panel_);

        	/**
        	 * 主页展示内容布局
        	 */
          // _popover_content_.append(_img_container_).append(_text_container_);

        	// $(_timeline_).append(_time_block_)
        	//   .append($("<div></div>").addClass("timeline-img"))
        	//   .append($("<div></div>").addClass("timeline-line"))
        	//   .append($("<div></div>").addClass("clearfix"));
        	// if ($(_timeline_).prev().find(".timeline-block").hasClass("timeline-block-odd")) {
         //  	$(_timeline_).find(".timeline-block").addClass("timeline-block-even");
         //  	$(_timeline_).find(".timeline-block-even>.timeline-content").addClass("right").css({"left": "150px"});
        	// } else {
        	//   $(_timeline_).find(".timeline-block").addClass("timeline-block-odd");
        	//   $(_timeline_).find(".timeline-block-odd>.timeline-content").addClass("left").css({"left": "-150px"});
        	// }

        	// $(_timeline_).find(".timeline-block>.timeline-content").animate({
        	//   left: "0px"
        	// }, 1000);

      	});

      	if($(window).height()>=document.documentElement.scrollHeight){
      	  //没有出现滚动条,继续加载下一页
      	  loadData();
      	}
    	});
  	}

  var tcScroll=function(){
    $(window).on('scroll', function () {
      var scrollTop = $(this).scrollTop();
      var scrollHeight = $(document).height();
      var windowHeight = $(this).height();
      if (scrollTop + windowHeight == scrollHeight) {
        //此处是滚动条到底部时候触发的事件，在这里写要加载的数据，或者是拉动滚动条的操作
        loadData();
      }
    })
  }
  loadData();
  tcScroll();

});

	</script>

	</body>
</html>

<!-- $(function(){
  	var _timeline_row_ = $("<div></div>").addClass("row");
  	$(".timeline-container").append(_timeline_row_);
  	var loadData=function(){
    	$.getJSON("/test", function (data) {
      	$.each(data, function (i, tl) {
        	var _timeline_ = $("<div></div>").addClass("timeline");
        	_timeline_row_.append(_timeline_);

        	var _time_block_ = $("<div></div>").addClass("timeline-block");
        	var _time_content_ = $("<div></div>").addClass("popover timeline-content");
        	_time_block_.append(_time_content_);
        	/**
        	 * 设置显示内容
        	 */

        	var _popover_title_ = $("<h3></h3>").addClass("popover-title").text(tl["diagTime"]);
        	var _popover_footer_ = $("<div></div>").addClass("popover-footer").text(tl["result"]);
        	var _popover_content_ = $("<div></div>").addClass("popover-content");
        	_time_content_.append($("<div></div>").addClass("arrow"))
          .append(_popover_title_)
          .append(_popover_content_)
          .append(_popover_footer_);
        	/**
        	 * 主页展示内容布局
        	 */

        if (tl["images"].length > 1 && tl["images"] != "" && tl["images"] != null && tl["images"] != "undefined") {
          var _img_container_ = $("<div></div>").css("margin-bottom", "10px");
          var _table_container_ = $("<table></table>").addClass("table table-bordered table-condensed");
          for (var m = 0; m < tl["images"].length; m++) {
            _img_container_.append($("<img/>").attr("src", tl["images"][m]).css({"width":"60px","height":"60px","margin-right":"10px"}));
          }
          _popover_content_.append(_img_container_);
          _table_container_.append($("<tr></tr>")
              .append($("<td nowrap></td>").text("眼象特征"))
              .append($("<td></td>").text(tl["feature"]))
          );

          _table_container_.append($("<tr></tr>")
              .append($("<td nowrap></td>").text("匹配结果"))
              .append($("<td></td>").text(tl["matchList"]))
          );

          _table_container_.append($("<tr></tr>")
              .append($("<td nowrap></td>").text("结论说明"))
              .append($("<td></td>").text(tl["desc"]))
          );

          _popover_content_.append(_img_container_).append(_table_container_);

        } else if (tl["images"].length == 1 && tl["images"] != "" && tl["images"] != null && tl["images"] != "undefined") {
          var _img_container_ = $("<div></div>")
            .addClass("pull-left")
            .append($("<img/>").attr("src",tl["images"][0]).css({"margin": "5px 10px","width": "100px", "height": "100px"}));
          var _text_container_ = $("<p></p>").css({"margin-left": "10px", "margin-top": "5px", "font-size": "12px"})
            .append("眼象特征: " + tl["feature"]).append("<br/>")
            .append("匹配结果: " + tl["matchList"]).append("<br/>")
            .append("结论说明: " + tl["desc"]).append("<br/>");
          _popover_content_.append(_img_container_).append(_text_container_);
        } else if (tl["images"].length < 1 || tl["images"] == "" || tl["images"] == null || tl["images"] == "undefined") {
          var _text_container_ = $("<p></p>").css({"margin-left": "10px", "margin-top": "5px", "font-size": "12px"})
            .append("眼象特征: " + tl["feature"]).append("<br/>")
            .append("匹配结果: " + tl["matchList"]).append("<br/>")
            .append("结论说明: " + tl["desc"]).append("<br/>");
          _popover_content_.append(_img_container_).append(_text_container_);
        }

        $(_timeline_).append(_time_block_)
          .append($("<div></div>").addClass("timeline-img"))
          .append($("<div></div>").addClass("timeline-line"))
          .append($("<div></div>").addClass("clearfix"));
        if ($(_timeline_).prev().find(".timeline-block").hasClass("timeline-block-odd")) {
          $(_timeline_).find(".timeline-block").addClass("timeline-block-even");
          $(_timeline_).find(".timeline-block-even>.timeline-content").addClass("right").css({"left": "150px"});
        } else {
          $(_timeline_).find(".timeline-block").addClass("timeline-block-odd");
          $(_timeline_).find(".timeline-block-odd>.timeline-content").addClass("left").css({"left": "-150px"});
        }
        $(_timeline_).find(".timeline-block>.timeline-content").animate({
          left: "0px"
        }, 1000);
      });
      if($(window).height()>=document.documentElement.scrollHeight){
        //没有出现滚动条,继续加载下一页
        loadData();
      }
    });
  } -->

