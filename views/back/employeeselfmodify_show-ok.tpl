<!-- iframe里展示个人待处理的详细情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>待处理成果</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />

  <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script>
<script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.js"></script>
<link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker3.css"/> 

<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/select2-bootstrap.css"/>

<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> 
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
<script type="text/javascript" src="/static/js/select2.js"></script>  
<!-- <script type="text/javascript" src="/static/js/mindmup-editabletable.js"></script> -->

<script src="/static/js/moment-with-locales.min.js"></script>

</head>


<div class="col-lg-12">
<div class="form-group">
        <label class="control-label" id="regis" for="LoginForm-UserName">{{.UserNickname}}</label><!-- 显示部门名称 -->
</div>
<div>

  <div class="form-group">
    <label for="taskNote">统计周期：</label>
    <input type="text" class="form-control" name="datefilter" id="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
  </div>

  <button id="button" class="btn btn-default">Refresh from url</button>

<br></div>

<div class="form-group">
<label class="control-label" id="regis" for="LoginForm-UserName">
  统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}
</label>
</div>
<h3>需要提交给校核</h3>

<div id="toolbar" class="btn-group">
<select class="form-control">
    <option value="">Export Basic</option>
    <option value="all">Export All</option>
    <option value="selected">Export Selected</option>
  </select>
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
      data-toolbar="#toolbar"
       data-url="/addinline"
       
      data-query-params="queryParams"
       data-striped="true"
       
       >
    <thead>        
    <tr> 

     <th data-field="state" data-checkbox="true"></th>
        <th data-field="Id">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <th data-field="Page">成果计量单位</th>
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
     
</div>

<script type="text/javascript">
function queryParams(params) {
  // var date=$("#datefilter").val();
  params.datefilter="2016-09-10 - 2016-09-15";//"2016-09-10 - 2016-09-15";
        return params;
    }
    var $table = $('#table'),
        $button = $('#button');

    $(function () {
        $button.click(function () {
            $table.bootstrapTable('refresh');
        });
    });   

</script>

</body>
</html>
