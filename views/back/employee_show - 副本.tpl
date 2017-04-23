<!-- iframe里展示个人详细情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>情况汇总</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />
<!-- <style type="text/css">
a:active{text:expression(target="_blank");}
i#delete
{
color:#DC143C;
}
</style>
<script type="text/javascript">
  var allLinks=document.getElementsByTagName("a");
for(var i=0;i!=allLinks.length; i++){
allLinks[i].target="_blank";
}
</script> -->
</head>


<!-- <div id="treeview" class="col-xs-3"></div> -->

<div class="col-lg-12">
<div>
<form class="form-inline" method="get" action="/secofficeshow" enctype="multipart/form-data">
  <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
  <input type="hidden" id="level" name="level" value="{{.Level}}"/>
  <div class="form-group">
    <label for="taskNote">统计周期：</label>
    <input type="text" class="form-control" name="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
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
  <button type="submit" class="btn btn-primary" name="button">提交</button>
</form>
<br></div>

<div class="form-group">
<label class="control-label" id="regis" for="LoginForm-UserName">
  统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}
</label>
</div>

  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>项目编号</th>
        <th>项目名称</th>
        <th>项目阶段</th>
        <th>成果编号</th>
        <th>成果名称</th>
        <th>成果类型</th>
        <th>成果计量单位</th>
        <th>成果数量</th>
        <th>编制、绘制</th>
        <th>设计</th>
        <th>校核</th>
        <th>审查</th>
      </tr>
    </thead>

    <tbody>
      <tr><th colspan=13>图纸</th></tr>
      {{range $k,$v :=.Catalogtuzhi}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>

        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Drawn }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
         <td>{{.Examined}}</td> 
      </tr>
      {{end}}

<tr><th colspan=13>报告</th></tr>
    {{range $k,$v :=.Catalogbaogao}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>

        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count}}</td>
        <td>{{.Drawn}}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>  
      </tr>
      {{end}}

<tr><th colspan=13>计算书</th></tr>
      {{range $k,$v :=.Catalogjisuanshu}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>

        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count}}</td>
        <td>{{.Drawn}}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>  
      </tr>
      {{end}}
 
 <tr><th colspan=13>修改单</th></tr>
      {{range $k,$v :=.Catalogxiugaidan}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>

        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count}}</td>
        <td>{{.Drawn}}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>  
      </tr>
      {{end}}

<tr><th colspan=13>大纲</th></tr>
      {{range $k,$v :=.Catalogdagang}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>

        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count}}</td>
        <td>{{.Drawn}}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>  
      </tr>
      {{end}}

<tr><th colspan=13>标书</th></tr>
      {{range $k,$v :=.Catalogbiaoshu}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>

        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Drawn }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>  
      </tr>
      {{end}}
    </tbody>
  </table>
  <tr>    
       <td colspan="4"><input type="button" class="btn btn-default" name="insert" value="在线编辑" onclick="insertNewRow()"/></td>    
       </tr>
</div>

<script type="text/javascript">
function insertNewRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        window.open('/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}}+'&key=editor');
        }
// $(function() {
         // $('#treeview').treeview('collapseAll', { silent: true });
          // $('#treeview').treeview({
          // data: [{{.json}}],//defaultData,
          // data:alternateData,
          // levels: 5,// expanded to 5 levels
          // enableLinks:true,
          // showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
//         });
// });


  $(document).ready(function() {
  $("table").tablesorter({sortList: [[6,1]]});
  // $("#ajax-append").click(function() {
  //    $.get("assets/ajax-content.html", function(html) {
  //     // append the "ajax'd" data to the table body
  //     $("table tbody").append(html);
  //     // let the plugin know that we made a update
  //     $("table").trigger("update");
  //     // set sorting column and direction, this will sort on the first and third column
  //     var sorting = [[2,1],[0,0]];
  //     // sort on the first column
  //     $("table").trigger("sorton",[sorting]);
  //   });
  //   return false;
  // });
});
</script>
</body>
</html>
<!-- <button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span> User
</button>

<button type="button" class="btn btn-primary btn-lg" style="text-shadow: black 5px 3px 3px;">
<span class="glyphicon glyphicon-user"></span> User
</button> -->