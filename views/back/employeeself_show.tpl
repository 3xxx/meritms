<!-- iframe里展示个人可添加的详细情况，只显示状态为1的，这个作废了-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>添加、编辑和提交</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />
 <script type="text/javascript" src="/static/js/bootstrap-select.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-select.min.css"/>  

<script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script>
<script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.js"></script>
<link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker3.css"/> 

<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script> 

<!-- <script src="/static/js/moment-with-locales.min.js"></script> -->
</head>


<div class="col-lg-12">
<div class="form-group" id="div2">
        <label class="control-label" id="regis" for="LoginForm-UserName">{{.UserNickname}}</label><!-- 显示部门名称 -->
</div>

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
  <!-- <button type="submit" class="btn btn-primary">提交</button> -->
  <button id="button1" class="btn btn-default">提交</button>
  </div>
<!-- </form> -->
<br></div>

<div class="form-group">
<label class="control-label" id="regis" for="LoginForm-UserName">
  统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}
</label>
</div>

<div id="toolbar" class="btn-group">
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
<table data-toggle="table"
       data-url="/myfornextlevel"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-query-params="queryParams"
       data-toolbar="#toolbar">
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
        <th data-field="Drawnratio">绘制系数</th>
        <th data-field="Data" data-formatter="localDateFormatter">出版</th>
        <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">操作</th>
      </tr>
    </thead>
</table>

<script>
function index1(value,row,index){
  alert( "Data Loaded: " + index );
            return index+1
          }
function localDateFormatter(value) {
  return moment(value, 'YYYY-MM-DD').format('L');
}
function queryParams(params) {
  var date=$("#datefilter").val();
  // alert( "Data Loaded: " + date );
        params.datefilter=date;
        return params;
    }
$(function () {
        $('#button1').click(function () {
            $('#table').bootstrapTable('refresh', {url:'/myfornextlevel'});
        });
    }); 
</script>
<script>
    // var $table = $('#table');
    // $(function () {
        // $table.bootstrapTable({
          // url: '/secofficeshow1'
            // data: [{
            //     "Id": 1,
            //     "Name": "Item 1",
            //     "Section": "$10"
            // },
            // {
            //     "Id": 2,
            //     "Name": "Item 2",
            //     "Section": "$20"
            // }
            // ]
        // });
        // $('#button').click(function () {
            // $table.bootstrapTable('refresh', {url: '/secofficeshow1'});
        // });
    // });


function actionFormatter(value, row, index) {
    return [
        '<a class="like" href="javascript:void(0)" title="Like">',
        '<i class="glyphicon glyphicon-heart"></i>',
        '</a>',
        '<a class="edit ml10" href="javascript:void(0)" title="Edit">',
        '<i class="glyphicon glyphicon-edit"></i>',
        '</a>',
        '<a class="remove ml10" href="javascript:void(0)" title="Remove">',
        '<i class="glyphicon glyphicon-remove"></i>',
        '</a>'
    ].join('');
}

window.actionEvents = {
    'click .like': function (e, value, row, index) {
        alert('You click like icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    },
    'click .edit': function (e, value, row, index) {
        alert('You click edit icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    },
    'click .remove': function (e, value, row, index) {
        alert('You click remove icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    }
};

</script>
  </div>

</body>
</html>
