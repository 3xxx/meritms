<!-- 项目列表进度 -->
<!DOCTYPE html>
{{template "tpl/T.header.tpl"}}

<title>项目进度-EngiCMS</title>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script> -->
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <!-- <script src="/static/js/tableExport.js"></script> -->
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/>

  <link href="/static/css/ganttstyle.css" type="text/css" rel="stylesheet">
  <script src="/static/js/jquery.fn.gantt.js"></script>
  <script src="/static/js/jquery.form.js"></script>
</head>

<style>
  /*body {
    margin: 0;
    padding: 0;
    font-family: "Lucida Grande",Helvetica,Arial,Verdana,sans-serif;
    font-size: 14px;
  }*/

  #script-warning {
    display: none;
    background: #eee;
    border-bottom: 1px solid #ddd;
    padding: 0 10px;
    line-height: 40px;
    text-align: center;
    font-weight: bold;
    font-size: 12px;
    color: red;
  }

  #loading {
    display: none;
    position: absolute;
    top: 10px;
    right: 10px;
  }

  #calendar {
    max-width: 900px;
    margin: 40px auto;
    padding: 0 10px;
  }

  /*body {这个导致窗口挪动，不好
    margin: 40px 10px;
    padding: 0;
    font-family: "Lucida Grande",Helvetica,Arial,Verdana,sans-serif;
    font-size: 14px;
  }*/

  /*#calendar {
    max-width: 900px;
    margin: 0 auto;
  }*/
      .fc-color-picker {
        list-style: none;
        margin: 0;
        padding: 0;
      }
      .fc-color-picker > li {
        float: left;
        font-size: 30px;
        margin-right: 5px;
        line-height: 30px;
      }
      .fc-color-picker > li .fa {
        -webkit-transition: -webkit-transform linear 0.3s;
        -moz-transition: -moz-transform linear 0.3s;
        -o-transition: -o-transform linear 0.3s;
        transition: transform linear 0.3s;
      }
      .fc-color-picker > li .fa:hover {
        -webkit-transform: rotate(30deg);
        -ms-transform: rotate(30deg);
        -o-transform: rotate(30deg);
        transform: rotate(30deg);
      }
      .ganttRed {
        color: #dd4b39 !important;
      }
      .ganttGreen {
        color: #00a65a !important;
      }
      .ganttBlue {
        color: #0073b7 !important;
      }
      .ganttOrange {
        color: #ff851b !important;
      }
      .text-red {
        color: #dd4b39 !important;
      }
      .text-yellow {
        color: #f39c12 !important;
      }
      /*.text-aqua {
        color: #00c0ef !important;
      }*/
      .text-blue {
        color: #0073b7 !important;
      }
      /*.text-black {
        color: #111111 !important;
      }
      .text-light-blue {
        color: #3c8dbc !important;
      }*/
      .text-green {
        color: #00a65a !important;
      }
      /*.text-gray {
        color: #d2d6de !important;
      }
      .text-navy {
        color: #001f3f !important;
      }
      .text-teal {
        color: #39cccc !important;
      }
      .text-olive {
        color: #3d9970 !important;
      }
      .text-lime {
        color: #01ff70 !important;
      }*/
      .text-orange {
        color: #ff851b !important;
      }
/*      .text-fuchsia {
        color: #f012be !important;
      }
      .text-purple {
        color: #605ca8 !important;
      }
      .text-maroon {
        color: #d81b60 !important;
      }*/
</style>

<body>
  <div class="navbar navba-default navbar-fixed-top">
      <div class="container-fill">{{template "tpl/T.navbar.tpl" .}}</div>
    </div>

<div class="col-lg-12">
  <h3>项目进度</h3>
  <div id="toolbar1" class="btn-group">
        <button type="button" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" id="importButton" class="btn btn-default"> <i class="fa fa-plus-square">导入</i>
        </button>
        <!-- <button type="button" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button> -->
  </div>
  <div class="gantt"></div>

  <script type="text/javascript">
        $(function() {
            "use strict";
            $(".gantt").gantt({
              source:"/projectgant/getprojgants",
                // source: [{
                //     name: "珠三角",
                //     desc: "可研",
                //     id:2,
                //     values: [{
                //         from: "/Date(1492790400000)/",
                //         to: "/Date(1501257600000)/",
                //         desc: '<b>Task #</b>3<br><b>Data</b>: [2011-02-01 15:30:00 - 2011-02-01 16:00:00]',
                //         label: "label是啥",
                //         customClass: "ganttRed",
                //         dataObj: ['ha','ha2']
                //     }]
                //   },
                months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"], //月份显示的语言
                dow: ["日", "一", "二", "三", "四", "五", "六"], //星期显示的语言
                navigate: "scroll",
                scale: "days",
                maxScale: "months",
                minScale: "hours",
                itemsPerPage: 20,
                useCookie: true,
                scrollToToday:true,
                markNow:true,
                onItemClick: function(data) {
                    // alert("Item clicked - show some details"+data);
                    $('#modalTable1').modal({
                    show:true,
                    backdrop:'static'
                    });
                },
                onAddClick: function(dt, rowId) {
                    // alert("Empty space clicked - add an item!"+dt+","+rowId);
                    if ({{.role}}!=1){
                      alert("权限不够！");
                      return;
                    }
                    $("label#info").remove();
                    $("#saveproj").removeClass("disabled")
                    $('#modalTable').modal({
                    show:true,
                    backdrop:'static'
                    });
                },
                onRender: function() {
                    if (window.console && typeof console.log === "function") {
                        console.log("chart rendered");
                    }
                }
            });

            // $(".gantt").popover({
            //     selector: ".bar",
            //     title: "I'm a popover",
            //     content: "And I'm the content of said popover.",
            //     trigger: "hover",
            //     placement: "auto right"
            // });

            // prettyPrint();
        });

  // RGB 转16进制
  var rgbToHex = function(rgb) {
    // rgb(x, y, z)
    var color = rgb.toString().match(/\d+/g); // 把 x,y,z 推送到 color 数组里
    var hex = "#";
    for (var i = 0; i < 3; i++) {
      // 'Number.toString(16)' 是JS默认能实现转换成16进制数的方法.
      // 'color[i]' 是数组，要转换成字符串.
      // 如果结果是一位数，就在前面补零。例如： A变成0A
      hex += ("0" + Number(color[i]).toString(16)).slice(-2);
    }
    return hex;
  }
  // 16进制 转 RGB
  
  // 能处理 #axbycz 或 #abc 形式
  var hexToRgb = function(hex) {
    var color = [], rgb = [];
    hex = hex.replace(/#/,"");
    if (hex.length == 3) { // 处理 "#abc" 成 "#aabbcc"
      var tmp = [];
      for (var i = 0; i < 3; i++) {
        tmp.push(hex.charAt(i) + hex.charAt(i));
      }
      hex = tmp.join("");
    }
    for (var i = 0; i < 3; i++) {
      color[i] = "0x" + hex.substr(i+2, 2);
      rgb.push(parseInt(Number(color[i])));
    }
    return "rgb(" + rgb.join(",") + ")";
  }

  var currColor = "#3c8dbc"; //Red by default
  var customclass = "ganttRed";
  $(function () {
    /* ADDING EVENTS */
    //Color chooser button
    // var colorChooser = $("#color-chooser-btn");
    $("#color-chooser > li > a").click(function (e) {
      e.preventDefault();
      //Save color
      currColor = $(this).css("color");
      customclass = $(this).attr("class");
      alert(customclass);
      //Add color effect to button
      $('#saveproj').css({"background-color": currColor, "border-color": currColor});
    });
    $("#color-chooser1 > li > a").click(function (e) {
      e.preventDefault();
      //Save color
      currColor = $(this).css("color");
      customclass = $(this).attr("class");
      //Add color effect to button
      $('#add-new-event1').css({"background-color": currColor, "border-color": currColor});
    });
  });

  $(document).ready(function() {
    $("#addButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("labe1#info").remove();
      $("#saveproj").removeClass("disabled")
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#importButton").click(function() {
        $('#importgants').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#editorButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选类别！");
        return;
      }
      if (selectRow.length>1){
        alert("请不要勾选一个以上！");
        return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#projcode1").val(selectRow[0].Code);
      $("#projname1").val(selectRow[0].Title);
      $("#projlabel1").val(selectRow[0].Label);
      $("#projprincipal1").val(selectRow[0].Principal);

        $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
    })

    // 删除项目
    $("#deleteButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选项目！");
        return false;
      }
      if(confirm("确定删除项目吗？第一次提示！")){
      }else{
        return false;
      }
      if(confirm("确定删除项目吗？第二次提示！")){
      }else{
        return false;
      }
      if(confirm("确定删除项目吗？一旦删除将无法恢复！")){
        var title=$.map(selectRow,function(row){
          return row.Title;
        })
        var ids="";
        for(var i=0;i<selectRow.length;i++){
          if(i==0){
            ids=selectRow[i].Id;
          }else{
            ids=ids+","+selectRow[i].Id;
          }  
        }
        $.ajax({
          type:"post",
          url:"/project/deleteproject",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#table0').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      }  
    })
  })



//填充select选项
$(document).ready(function(){
  //   $(array).each(function(index){
  //     alert(this);
  // });
   
  // $.each(array,function(index){
  //     alert(this);
  // });
  if ({{.Select2}}){
    $.each({{.Select2}},function(i,d){
    // alert(this);
    // alert(i);
    // alert(d);
      // $("#admincategory").append('<option value="' + i + '">'+d+'</option>');
    });
  }
});

//根据选择，刷新表格
// function refreshtable(){
  // var newcategory = $("#admincategory option:selected").text();
  // alert("你选的是"+newcategory);
  //根据名字，查到id，或者另外写个categoryname的方法
  // $('#table1').bootstrapTable('refresh', {url:'/admin/categorytitle?title='+newcategory});
  // $("#details").show();
// }

  //保存
  function save(){
      // var radio =$("input[type='radio']:checked").val();
      var code = $('#code').val();
      var name = $('#title').val();
      var designstage = $("#designstage option:selected").text();
      var section = $("#section option:selected").text();//专业
      var label = $('#label').val();
      var desc = $('#desc').val();
      // var customclass = $('#customclass').val();
      var dataobj = $('#dataobj').val();
      var datefilter = $('#datefilter').val();
      // var end = $('#end').val();
      // var selectRow3=$('#table1').bootstrapTable('getSelections');
      // if (selectRow3.length<1){
      //   alert("请先勾选目录！");
      //   return;
      // }
      
      // var ids="";
      // for(var i=0;i<selectRow3.length;i++){
      //   if(i==0){
      //     ids=selectRow3[i].Id;
      //   }else{
      //     ids=ids+","+selectRow3[i].Id;
      //   }  
      // }

      // $('#myModal').on('hide.bs.modal', function () {  
      if (name&&code)
        {
          var lab="<label id='info'>添加项目进度中，请耐心等待几秒/分钟……</label>";
          $(".modal-footer").prepend(lab);//这里是否要换名字$("p").remove();
          $("#saveproj").addClass("disabled");
          $.ajax({
            type:"post",
            url:"/projectgant/addprojgant",
            data: {code:code,title:name,designstage:designstage,section:section,label:label,desc:desc,customclass:customclass,dataobj:dataobj,datefilter:datefilter},//
            success:function(data,status){
              alert("添加“"+data+"”成功！(status:"+status+".)");
              //按确定后再刷新
              $('#modalTable').modal('hide');
            }
          });  
        }else{
          alert("请填写编号和名称！");
          return;
        } 
        // $(function(){$('#myModal').modal('hide')}); 
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
  }

    //导入用户数据表
    function importgants(){
        var file=$("#gantsexcel").val();
        if(file!=""){  
            var form = $("form[id=form1]");
            var options  = {    
                url:'/projectgant/importprojgant',    
                type:'post', 
                success:function(data)    
                {    
                  alert("导入数据："+data+"！")
                }    
            };
           form.ajaxSubmit(options);
           return false;
        }else{
            alert("请选择文件！");
            return false; 
        }
    }
  //修改项目
  function update(){
      // var radio =$("input[type='radio']:checked").val();
      var pid = $('#cid').val();
      var projcode = $('#projcode1').val();
      var projname = $('#projname1').val();
      var projlabel = $('#projlabel1').val();
      var projprincipal = $('#projprincipal1').val();

      if (projname&&projcode)
        {  
          $.ajax({
            type:"post",
            url:"/project/updateproject",
            data: {code:projcode,name:projname,label:projlabel,principal:projprincipal,pid:pid},//
            success:function(data,status){
              alert("修改“"+data+"”成功！(status:"+status+".)");
              //按确定后再刷新
              $('#modalTable1').modal('hide');
              // $('#table0').bootstrapTable('refresh', {url:'/project/getprojects'});
            }
          });  
        }else{
          alert("请填写编号和名称！");
          return;
        } 
  }
  //模态框可拖曳—要引入ui-jquery.js
  // $("#modalTable").draggable({
  //   handle:".modal-header",
  //   cusor:"move",
  //   refreshPositions:false
  // });
  // 来自群，保留，批量
  // var rows= $('#account-table').bootstrapTable('getSelections');
  //       if(rows.length==0) {
  //           layer.alert('请您选择要删除的子账号！', {
  //               title:'提示信息',
  //               closeBtn: 0,
  //               icon: 0,
  //               skin: 'layui-layer-lan',
  //               shift:0 //动画类型
  //           });
  //           return false;
  //       }
  //           var ids="";
  //           for(var i=0;i<rows.length;i++){
  //               if(i==0){
  //                   ids=rows[i].frontUserId;
  //               }else{
  //                   ids=ids+","+rows[i].frontUserId;
  //               }
  //           }
</script>
<!-- 添加项目进度 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加项目进度</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">项目编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="code"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">项目名称</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="title" placeholder="用简称，不超过3个字"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">阶段</label>
                <div class="col-sm-7">
                  <select id="designstage" class="form-control">
                    <!-- <option>选择类别：</option> -->
                    <option value="0">选择阶段：</option>
                    <option value="1">规划</option>
                    <option value="2">项建</option>
                    <option value="3">可研</option>
                    <option value="4">初设</option>
                    <option value="5">招标</option>
                    <option value="6">施工图</option>
                    <option value="7">竣工图</option>
                  </select>
                </div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">专业</label>
                <div class="col-sm-7">
                  <select id="section" class="form-control">
                    <!-- <option>选择类别：</option> -->
                    <option value="0">选择专业：</option>
                    <option value="1">水工</option>
                    <option value="2">施工</option>
                    <option value="3">预算</option>
                  </select>
                </div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="label" placeholder="进度条上显示的文字"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">描述</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="desc" placeholder="鼠标移到进度条上显示文字"></div>
              </div>

              <div class="form-group must">
                <label class="col-sm-3 control-label">进度条颜色</label>
                <div class="col-sm-7">
                  <div class="btn-group" style="width: 100%; margin-bottom: 10px;">
                    <ul class="fc-color-picker" id="color-chooser">
                      <!-- <li><a class="text-aqua" href="#"><i class="fa fa-square"></i></a></li> -->
                      <li><a class="ganttBlue" href="#"><i class="fa fa-square"></i></a></li>
                      <!-- <li><a class="text-light-blue" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-teal" href="#"><i class="fa fa-square"></i></a></li> -->
                      <!-- <li><a class="text-yellow" href="#"><i class="fa fa-square"></i></a></li> -->
                      <li><a class="ganttOrange" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="ganttGreen" href="#"><i class="fa fa-square"></i></a></li>
                      <!-- <li><a class="text-lime" href="#"><i class="fa fa-square"></i></a></li> -->
                      <li><a class="ganttRed" href="#"><i class="fa fa-square"></i></a></li>
                      <!-- <li><a class="text-purple" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-fuchsia" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-muted" href="#"><i class="fa fa-square"></i></a></li>
                      <li><a class="text-navy" href="#"><i class="fa fa-square"></i></a></li> -->
                    </ul>
                  </div>
                </div>
              </div>

              <div class="form-group must">
                <label class="col-sm-3 control-label">dataobj</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="dataobj" placeholder="备选说明，用英文逗号,隔开"></div>
              </div>
              <div class="form-group">
                <label class="col-sm-3 control-label" for="taskNote">周期：</label>
                <div class="col-sm-7">
                <input type="text" class="form-control" name="datefilter" id="datefilter" value="" placeholder="选择时间段(默认最近一个月)"></div>
              </div>
              <script type="text/javascript">
                $(function() {
          $('input[name="datefilter"]').daterangepicker({
            ranges : {
              'Today': [moment(), moment()],
              'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
              'Last 7 Days': [moment().subtract(6, 'days'), moment()],
              'Last 30 Days': [moment().subtract(29, 'days'), moment()],
              'This Month': [moment().startOf('month'), moment().endOf('month')],
              'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
            },
            autoUpdateInput: false,
            locale: {
              cancelLabel: 'Clear'
            }
          });
          $('input[name="datefilter"]').on('apply.daterangepicker', function(ev,        picker){
            $(this).val(picker.startDate.format('YYYY-MM-DD') + ' - ' + picker.      endDate.   format('YYYY-MM-DD'));
          });
          $('input[name="datefilter"]').on('cancel.daterangepicker', function(ev,        picker)    {
            $(this).val('');
          });
                });
              </script>
              <!-- <div class="form-group">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="number" class="form-control digits" name="label" maxlength="20" placeholder="至多20个字符" required></div>
              </div> -->
              <!-- <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                  <div class="col-sm-7">
                    <input type="password" class="form-control" name="password" id="password" maxlength="32" placeholder="至多32个字符" required></div>
              </div> -->

            <!-- <div class="form-group must">
              <label class="col-sm-3 control-label">确认密码</label>
              <div class="col-sm-7">
                <input type="password" class="form-control equalto" name="password2" maxlength="32" placeholder="至多32个字符" required data-rule-equalto="#password" data-msg-equalto="密码不一致"></div>
            </div> -->
            </div>

            <!-- <div id="details" style="display:none">
              <h3>工程目录分级</h3>
              <table id="table1"
                    data-page-size="5"
                    data-page-list="[5, 25, 50, All]"
                    data-unique-id="id"
                    data-sort-name="Grade"
                    data-pagination="true"
                    data-side-pagination="client"
                    data-click-to-select="true">
                  <thead>        
                  <tr>
                    <th data-width="10" data-checkbox="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="Title">名称</th>
                    <th data-field="Code">代码</th>
                    <th data-field="Grade" data-sortable="true">级别</th>
                  </tr>
                </thead>
              </table>
            </div> -->
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" id="saveproj" onclick="save()">保存</button><!--  style="display:none" -->
          </div>
        </div>
      </div>
    </div>
  </div>

<!-- 导入用户数据 -->
<div class="container form-horizontal">
    <div class="modal fade" id="importgants">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加用户</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content"> 
              <div class="form-group">
                <form method="post" id="form1" action="/projectgant/importprojgant" enctype="multipart/form-data">
                  <div class="form-inline" class="form-group">
                    <label>选择项目进度数据文件(Excel)：
                      <input type="file" class="form-control" name="gantsexcel" id="gantsexcel"/> </label>
                    <br/>          
                  </div>
                  <!-- <button type="submit" class="btn btn-default">提交</button> -->
                </form>
              </div>
            </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="submit" class="btn btn-primary" onclick="return importgants();">导入</button>
          <!-- <button type="submit" class="btn btn-primary" onclick="return import_xls_catalog();">提交</button> -->
        </div>
      </div>
    </div>
  </div>
</div>

  <!-- 编辑项目 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑项目</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="projcode1" name="projcode"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">名称</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projname1" name="projname"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projlabel1" name="projlabel"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="projprincipal1" name="projprincipal"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" id="updateproj" onclick="update()">更新</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</div>
</body>
</html>