<!-- 测试用 -->
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
<!-- <div id="toolbar" class="btn-group">
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
       data-url="/secofficeshow"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-toolbar="#toolbar">
    <thead>
    <tr>
        <th data-field="name">Name</th>
        <th data-field="stargazers_count">Stars</th>
        <th data-field="forks_count">Forks</th>
        <th data-field="description">Description</th>
    </tr>
    </thead>
</table> 
<table data-toggle="table" data-url="{{.json}}">
    <thead>
        <tr>
            <th data-field="Id">Item ID</th>
            <th data-field="ProjectNumber">Item Name</th>
        </tr>
    </thead>
</table> -->
<!-- <div class="container">
    <h1>Refresh from url after use data option.</h1>
    <p><button id="button" class="btn btn-default">Refresh from url</button></p>
    <table id="table">
        <thead>
        <tr>
            <th data-field="Id">ID</th>
            <th data-field="ProjectNumber">Item Name</th>
            <th data-field="Tnumber">Item Price</th>
        </tr>
        </thead>
    </table>
</div> -->
<table data-toggle="table">
    <thead>
        <tr>
            <th data-field="id">Item ID</th>
            <th data-field="name">Item Name</th>
            <th data-field="price">Item Price</th>
        </tr>
    </thead>
</table>

<script> 
$(document).ready(function () { 
// $("#table").bootstrapTable("load", {{.json}});         
          //调用函数，初始化表格  
          // initTable();  
  $('#table').bootstrapTable({
          url: '/test'
          //当点击查询按钮的时候执行  
          $("#search").bind("click", initTable);  
      });

 }); 
    // var $table = $('#table');  
    // $(function () {  
    //     $table.bootstrapTable({ 
    //     method: "get",  //使用get请求到服务器获取数据  
    //         url: "/test", //获取数据的Servlet地址 
            // data: [
            //         {
            //           "Id": 34,
            //           "ProjectNumber": "SL000",
            //           "ProjectName": "广东",
            //           "DesignStage": "项目建议书",
            //           "Section": "水工",
            //           "Tnumber": "DZ122D.5-1-1",
            //           "Name": "工程系统示意图"
            //         },
            //         {
            //           "Id": 35,
            //           "ProjectNumber": "SL3999",
            //           "ProjectName": "惠来县中东部供水工程",
            //           "DesignStage": "初步设计",
            //           "Section": "",
            //           "Tnumber": "DZ122D.1-1-1",
            //           "Name": "邦山取水泵站枢纽布置图"
            //         }
            //  ]        
    //     });  
    //     $('#button').click(function () {  
    //         $table.bootstrapTable('refresh', {url: ''});  
    //     });  
    // });  
</script> 

</body>
</html>