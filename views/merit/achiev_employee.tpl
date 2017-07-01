<!-- iframe里展示个人详细情况，自己查看显示处理按钮，别人查看不显示处理按钮-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>情况汇总</title>

  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/>

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>

  <script src="/static/js/moment-with-locales.min.js"></script>
  <script type="text/javascript" src="/static/js/echarts.min.js"></script>

  <script src="/static/js/bootstrap-table-filter-control.js"></script>
  <script src="/static/js/bootstrap-table-export.min.js"></script>
  <script src="/static/js/tableExport.js"></script>
</head>

<!-- <div class="form-group"> -->
<div class="col-lg-12">
  <h2>{{.UserNickname}}</h2>
  <ul class="nav nav-tabs">
    <li class="active">
      <a href="#employee" data-toggle="tab">目前</a>
    </li>
    <li>
      <a href="#year" data-toggle="tab">年度</a>
    </li>
    <li>
      <a href="#proj" data-toggle="tab">项目</a>
    </li>
  </ul>

  <div class="tab-content">
    <div class="tab-pane fade in active" id="employee">
      <br>
      <div>
        <div class="form-inline">
          <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
          <input type="hidden" id="level" name="level" value="{{.Level}}"/>
          <input type="hidden" id="key" name="key" value="modify"/>
          <div class="form-group">
            <label for="taskNote">统计周期：</label>
            <input type="text" class="form-control" name="datefilter" id="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
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
              $('input[name="datefilter"]').on('apply.daterangepicker', function(ev, picker) {
                  $(this).val(picker.startDate.format('YYYY-MM-DD') + ' - ' + picker.endDate.format('YYYY-MM-DD'));
              });
              $('input[name="datefilter"]').on('cancel.daterangepicker', function(ev, picker) {
                  $(this).val('');
              });
            });
          </script>
          <!-- <button type="submit" class="btn btn-primary">提交</button>
        -->
          <button id="button1" class="btn btn-default">提交</button>
          <label class="control-label"> tips:(StartDay < DateRange <= EndDay)</label>
        </div>
      </div>
      <br>
        <!-- ：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}} -->

      <h3>已完成</h3>
      <div id="completed" class="btn-group">
          <button type="button" class="btn btn-default">
              <i class="glyphicon glyphicon-plus"></i>
          </button>
          <button type="button" class="btn btn-default">
              <i class="glyphicon glyphicon-heart"></i>
          </button>
          <button type="button" class="btn btn-default">
              <i class="glyphicon glyphicon-trash"></i>
          </button>
      </div>
      <table id="table"
              data-toggle="table"
             data-url="/achievement/send/6"
             data-search="true"
             data-show-refresh="true"
             data-show-toggle="true"
             data-show-columns="true"
             data-toolbar="#completed"
             data-query-params="queryParams"
             >
          <thead>        
          <tr>
              <th data-formatter="index1">#</th>
              <th data-field="ProjectNumber" data-visible='false'>项目编号</th>
              <th data-field="ProjectName" data-sortable="true">项目名称</th>
              <th data-field="DesignStage" data-sortable="true">项目阶段</th>
              <th data-field="Tnumber">成果编号</th>
              <th data-field="Name">成果名称</th>
              <th data-field="Category" data-sortable="true">成果类型</th>
              <!-- <th data-field="Page">成果计量单位</th> -->
              <th data-field="Count">成果数量</th>
              <th data-field="Drawn">编制、绘制</th>
              <th data-field="Designd">设计</th>
              <th data-field="Checked">校核</th>
              <th data-field="Examined">审查</th>
              <th data-field="Date" data-formatter="localDateFormatter">出版</th>
              <!-- <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">操作</th> -->
            </tr>
          </thead>
      </table>

<!-- <div class="form-group">   
      <input type="button" class="btn btn-primary" name="insert" value="在线添加" onclick="insertNewRow()"/>
          <form id="form1" class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
            <div class="form-group">
              <label>选择成果登记数据文件(Excel)
              <input type="file" class="form-control" name="catalog" id="catalog"></label>
              <br/>
              </div>
            <button type="submit" class="btn btn-primary" onclick="return import_xls_catalog();">提交</button>
          </form>
</div> --> 

<h3>已提交</h3>
<div id="running" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table1"
       data-toggle="table"
       data-url="/achievement/send/5"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-toolbar="#running"
       data-query-params="queryParams"
       data-sort-name="ProjectName"
       data-sort-order="desc"
       >
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber" data-visible='false'>项目编号</th>
        <th data-field="ProjectName" data-sortable="true">项目名称</th>
        <th data-field="DesignStage" data-sortable="true">项目阶段</th>
        <th data-field="Tnumber" data-sortable="true">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category" data-sortable="true">成果类型</th>
        <!-- <th data-field="Page">成果计量单位</th> -->
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>
<script>
function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
          }
function localDateFormatter(value) {
  return moment(value, 'YYYY-MM-DD').format('L');
}
function queryParams(params) {
  var date=$("#datefilter").val();
  var secid=$("#secid").val();
  var level=$("#level").val();
  // alert( "Data Loaded: " + date );
        params.datefilter=date;
        params.secid=secid;//传secid给后台，点击用户名，显示对应成果
        params.level=level;
        return params;
    }
$(function () {
        $('#button1').click(function () {
            //已完成
            $('#table').bootstrapTable('refresh', {url:'/completed'});
            //已提交
            $('#table1').bootstrapTable('refresh', {url:'/running'});
            //待提交
            $('#table2').bootstrapTable('refresh', {url:'/myself'});
            //待我处理的设计
            $('#table3').bootstrapTable('refresh', {url:'/designd'});
            //待我处理的校核
            $('#table4').bootstrapTable('refresh', {url:'/checked'});
            //待我处理的审查
            $('#table5').bootstrapTable('refresh', {url:'/examined'});

        });
    }); 
</script>

<h3>我发起，待提交</h3>
<div id="send" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table2"
        data-toggle="table"
       data-url="/achievement/send/1"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-toolbar="#send"
       data-query-params="queryParams"
       >
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <!-- <th data-field="Page">成果计量单位</th> -->
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>

<h3>待我处理设计</h3>
<div id="designd" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table3"
        data-toggle="table"
       data-url="/achievement/send/2"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-query-params="queryParams"
       data-toolbar="#designd"
       >
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <!-- <th data-field="Page">成果计量单位</th> -->
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>

<h3>待我处理校核</h3>
<div id="checked" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table4"
        data-toggle="table"
       data-url="/achievement/send/3"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-query-params="queryParams"
       data-toolbar="#checked">
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <!-- <th data-field="Page">成果计量单位</th> -->
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>

<h3>待我处理审查</h3>
<div id="examined" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table5"
        data-toggle="table"
       data-url="/achievement/send/4"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-query-params="queryParams"
       data-toolbar="#examined">
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <!-- <th data-field="Page">成果计量单位</th> -->
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>
  
        <tr>
        {{if .IsMe}}   
          <td colspan="4"><input type="button" class="btn btn-primary" value="处&nbsp;&nbsp;&nbsp;&nbsp;理" onclick="ModifyRow()"/></td> 
         {{end}}    
        </tr>
        <br/>
        <br/>
</div>

    <div class="tab-pane fade" id="year">
      <!-- 全年个人每个月总分及排名 -->
      <div class="col-lg-12" id="main1" style="width: 800px;height:600px;"></div>
      <!-- 当月成果类型组成 -->
      <div class="col-lg-12" id="main" style="width: 800px;height:600px;"></div>
      <!-- 当年成果类型组成 -->
      <div class="col-lg-12" id="main2" style="width: 800px;height:600px;"></div>
    </div>

    <div class="tab-pane fade" id="proj">
      <p>这里将显示参与的所有项目列表，每个项目的贡献程度。</p>

      <h3>我参与的项目</h3>
        <div id="participate" class="btn-group">
          <button type="button" class="btn btn-default">
             <i class="glyphicon glyphicon-plus"></i>
          </button>
          <button type="button" class="btn btn-default">
             <i class="glyphicon glyphicon-heart"></i>
          </button>
          <button type="button" class="btn btn-default">
             <i class="glyphicon glyphicon-trash"></i>
          </button>
        </div>
          <table id="table6"
          data-toggle="table"
          data-url="/achievement/participate"
          data-search="true"
          data-show-refresh="true"
          data-show-toggle="true"
          data-show-columns="true"
          data-toolbar="#participate"
          data-query-params="queryParams"
          data-filter-control="true"
          data-filter-show-clear="true"
          data-show-export="true"
          >
        <thead>        
          <tr>
          <th data-formatter="index1">#</th>
          <th data-field="ProjectNumber" data-filter-control="select">项目编号</th>
          <th data-field="ProjectName" data-sortable="true" data-filter-control="select">项目名称</th>
          <th data-field="DesignStage" data-sortable="true" data-filter-control="select">项目阶段</th>
          <th data-field="Value" data-sortable="true">成果总分</th>
          <th data-field="Myvalue" data-sortable="true">我的贡献</th>
          <th data-field="Percent" data-sortable="true">百分比</th>
          <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">详细</th>
          </tr>
        </thead>
      </table>

      <div class="container">
        <div class="modal fade" id="modalTable" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="bottom:auto">
            <div class="modal-dialog" style="width:auto;height:800px;overflow:auto;">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">本专业项目全部成果列表</h4>
                    </div>
                    <div class="modal-body">
                    <div id="projachievement" class="btn-group">
                      <button type="button" class="btn btn-default">
                         <i class="glyphicon glyphicon-plus"></i>
                      </button>
                      <button type="button" class="btn btn-default">
                         <i class="glyphicon glyphicon-heart"></i>
                      </button>
                      <button type="button" class="btn btn-default">
                         <i class="glyphicon glyphicon-trash"></i>
                      </button>
                    </div>
                        <table id="table7"
                              data-search="true"
                              data-show-refresh="true"
                              data-show-toggle="true"
                              data-show-columns="true"
                              data-toolbar="#projachievement"
                              data-query-params="queryParams"
                              data-filter-control="true"
                              data-filter-show-clear="true"
                              data-show-export="true"
                              >
                            <thead>
                            <tr>
                                <th data-formatter="index1">#</th>
                                <th data-field="Tnumber" data-sortable="true">成果编号</th>
                                <th data-field="Name">成果名称</th>
                                <th data-field="Category" data-filter-control="select">成果类型</th>
                                <th data-field="Count">成果数量</th>
                                <th data-field="Drawn" data-filter-control="select">编制、绘制</th>
                                <th data-field="Designd" data-filter-control="select">设计</th>
                                <th data-field="Checked" data-filter-control="select">校核</th>
                                <th data-field="Examined" data-filter-control="select">审查</th>
                            </tr>
                            </thead>
                        </table>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->

        <div class="modal fade" id="modalTable1" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="bottom:auto">
            <div class="modal-dialog" style="width:auto">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">本专业项目个人贡献率</h4>
                    </div>
                    <div class="modal-body">
                    
                    <!-- 插入每个人贡献比例饼图（全年） -->
                    <div class="col-lg-12" id="main4" style="width: 800px;height:600px;"></div>
                    
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->

        <div class="modal fade" id="modalTable2" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="bottom:auto">
            <div class="modal-dialog" style="width:auto">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">本专业项目成果类型比例</h4>
                    </div>
                    <div class="modal-body">
                    <!-- 插入成果类型比例饼图（全年） -->
                    <div class="col-lg-12" id="main3" style="width: 800px;height:600px;"></div>
                    
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->
      </div>

    </div>
  </div>
</div>

<script type="text/javascript">
function import_xls_catalog(){
  var form1 = window.document.getElementById("form1");//获取form1对象
  form1.submit();
  $.ajax({
          success:function(data,status){//数据提交成功时返回数据
              // alert("添加“"+data+"”成功！(status:"+status+".)");
              window.location.reload();
          }
        });
    return true;  //这个return必须放最后，前面的值才能传到后台    
   }

function insertNewRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        window.open('/achievement/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}}+'&key=editor');
        }
 function ModifyRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        window.open('/achievement/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}}+'&key=modify');
        }       

function actionFormatter(value, row, index) {
    return [
        '<a class="like" href="javascript:void(0)" title="成果列表">',
        '<i class="glyphicon glyphicon-list-alt"></i>',
        '</a>&nbsp;',
        '<a class="edit ml10" href="javascript:void(0)" title="贡献组成">',
        '<i class="glyphicon glyphicon-user"></i>',
        '</a>&nbsp;',
        '<a class="remove ml10" href="javascript:void(0)" title="成果类型组成">',
        '<i class="glyphicon glyphicon-adjust"></i>',//fa-pie-chart
        '</a>'
    ].join('');
}

window.actionEvents = {
    'click .like': function (e, value, row, index) {
     //提交到后台
     //*******错误！！！！直接用表格获取即可。不用ajax
       $.ajax({
       // type:"post",
       url:"/achievement/projectachievement",
       data: {CatalogId:row.Id},
            success:function(data,status){
             var data1 =data ;
             $('#table7').bootstrapTable({data: data1});
             $('#table7').bootstrapTable('refresh', {url:'/achievement/projectachievement?CatalogId='+row.Id});
            }
       });

        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
        // alert('You click like icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
    },
    'click .edit': function (e, value, row, index) {
      //展示这个项目（阶段、专业）一年来的每个人贡献比例 
      var myChart4 = echarts.init(document.getElementById('main4'));  
      $.get('/achievement/projectuserparticipate?CatalogId='+row.Id).done(function (data) {
            // 填入数据
            myChart4.setOption({
                title: {
                text: '一年来该项目贡献组成'
            },
                tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b}: {c} ({d}%)"
            },
            legend: {
                orient: 'vertical',
                x : 'left',
                // data:{{.UserName}}
            },
            toolbox: {
                show : true,
                feature : {
                    mark : {show: true},
                    dataView : {show: true, readOnly: false},
                    magicType : {
                        show: true, 
                        type: ['pie', 'funnel'],
                        option: {
                            funnel: {
                                x: '25%',
                                width: '50%',
                                funnelAlign: 'left',
                                max: 1548
                            }
                        }
                    },
                    restore : {show: true},
                    saveAsImage : {show: true}
                }
            },
            calculable : true,
            series: [{
                name:'贡献组成',
                type:'pie',
                radius: ['10%', '60%'],
                // 根据名字对应到相应的系列
                data: data
            }]
        });
      });

      $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
        // alert('You click edit icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
    },
    'click .remove': function (e, value, row, index) {
      //展示这个项目（阶段、专业）一年来的成果类型
      var myChart3 = echarts.init(document.getElementById('main3'));
      $.get('/achievement/echarts3?CatalogId='+row.Id).done(function (data) {
            // 填入数据
        myChart3.setOption({
                title: {
                text: '一年来该项目成果类型组成',
                x:'center'
            },
                tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b}: {c} ({d}%)"
            },
            legend: {
                orient: 'vertical',
                x : 'left',
                data:{{.Select2}}
            },
            toolbox: {
                show : true,
                feature : {
                    mark : {show: true},
                    dataView : {show: true, readOnly: false},
                    magicType : {
                        show: true, 
                        type: ['pie', 'funnel'],
                        option: {
                            funnel: {
                                x: '25%',
                                width: '50%',
                                funnelAlign: 'left',
                                max: 1548
                            }
                        }
                    },
                    restore : {show: true},
                    saveAsImage : {show: true}
                }
            },
            calculable : true,
            series: [{
                name:'成果类型组成',
                type:'pie',
                radius: ['10%', '60%'],
                // 根据名字对应到相应的系列
                data: data
            }]
        });
      }); 

        $('#modalTable2').modal({
        show:true,
        backdrop:'static'
        });
        // alert('You click remove icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
    }
};

//模态框垂直水平居中——这个绝对居中，没有用。
// function centerModals(){
//   $('#modalTable').each(function(i){
//     var $clone=$(this).clone().css('diplay','block').appendTo('body');
//     var top=Math.round(($clone.height('content').height())/2);
//     top=top>0?top:0;
//     $clone.remove();
//     $(this).find('.modal-content').css("margin-top",top);

//   });
// }
// $('#modalTable').on('show.bs.modal',centerModals);
// $(window).on('resize',centerModals);

//当月成果类型组成
        var myChart = echarts.init(document.getElementById('main'));
        $.get('/achievement/echarts').done(function (data) {
            // 填入数据
          myChart.setOption({
                title: {
                text: '累计一个月来成果类型组成',
                x:'center'
            },
                tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b}: {c} ({d}%)"
            },
            legend: {
                orient: 'vertical',
                x:'left',
                data:{{.Select2}}
            },
            toolbox: {
                show : true,
                feature : {
                    mark : {show: true},
                    dataView : {show: true, readOnly: false},
                    magicType : {
                        show: true, 
                        type: ['pie', 'funnel'],
                        option: {
                            funnel: {
                                x: '25%',
                                width: '50%',
                                funnelAlign: 'left',
                                max: 1548
                            }
                        }
                    },
                    restore : {show: true},
                    saveAsImage : {show: true}
                }
            },
            calculable : true,
            series: [{
                name:'成果类型',
                type:'pie',
                radius: ['10%', '60%'],
                // 根据名字对应到相应的系列
                data: data
            }]
          });
        });

//当年成果类型组成
        var myChart2 = echarts.init(document.getElementById('main2'));
        $.get('/achievement/echarts2').done(function (data) {
            // 填入数据
            myChart2.setOption({
                title: {
                text: '累计一年来成果类型组成',
                x:'center'
            },
                tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b}: {c} ({d}%)"
            },
            legend: {
                orient: 'vertical',
                x : 'left',
                data:{{.Select2}}
            },
            toolbox: {
                show : true,
                feature : {
                    mark : {show: true},
                    dataView : {show: true, readOnly: false},
                    magicType : {
                        show: true, 
                        type: ['pie', 'funnel'],
                        option: {
                            funnel: {
                                x: '25%',
                                width: '50%',
                                funnelAlign: 'left',
                                max: 1548
                            }
                        }
                    },
                    restore : {show: true},
                    saveAsImage : {show: true}
                }
            },
            calculable : true,
                series: [{
                    name:'成果类型',
                    type:'pie',
                    radius: ['10%', '60%'],
                    // 根据名字对应到相应的系列
                    data: data
                }]
            });
        });

      var myChart1 = echarts.init(document.getElementById('main1'));
        // option = {
     // $.get('/achievement/echarts').done(function (data) {     
          myChart1.setOption({
            title: {
                text: '个人每月工作量和排名',
                x:'left'
            },
          tooltip : {
              trigger: 'axis'
          },
          toolbox: {
              show : true,
              feature : {
                  mark : {show: true},
                  dataView : {show: true, readOnly: false},
                  magicType: {show: true, type: ['line', 'bar']},
                  restore : {show: true},
                  saveAsImage : {show: true}
              }
          },
          calculable : true,
          legend: {
              data:['全部工作量','实物工作量','排名']
          },
          xAxis : [
              {
                  type : 'category',
                  name : '月份',
                  data : {{.Value3}}//['1月','2月','3月','4月','5月','6月','7月','8月','9月','10月','11月','12月']
              }
          ],
          yAxis : [
              {
                  type : 'value',
                  name : '工作量',
                  axisLabel : {
                      formatter: '{value}'
                  }
              },
              {
                  type : 'value',
                  name : '排名',
                  axisLabel : {
                      formatter: '{value}'
                  }
              }
          ],
          series : [
              {
                  name:'全部工作量',
                  type:'bar',
                  data:{{.Value1}}
              },
              {
                  name:'实物工作量',
                  type:'bar',
                  data:{{.realValue1}}
              },
              {
                  name:'全部工作量排名',
                  type:'line',
                  yAxisIndex: 1,
                  data:{{.Value2}}
              },
              {
                  name:'实物工作量排名',
                  type:'line',
                  yAxisIndex: 1,
                  data:{{.realValue2}}
              }
          ]
        });
     // }); 
   
</script>
</body>
</html>
