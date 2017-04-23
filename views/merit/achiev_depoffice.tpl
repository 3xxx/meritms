<!-- iframe里展示分院总体情况-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>分院情况汇总</title>
  <!-- <base target=_blank>
  -->
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/>
  <!-- <style type="text/css">
  a:active{text:expression(target="_blank");}
i#delete
{
color:#DC143C;
}
</style>
-->
<!-- <script type="text/javascript">
var allLinks=document.getElementsByTagName("a");
for(var i=0;i!=allLinks.length; i++){
allLinks[i].target="_blank";
}
</script>
-->
</head>

<div class="col-lg-12">
<h2>{{.Deptitle}}</h2>
<ul class="nav nav-tabs">
  <li class="active"><a href="#employee" data-toggle="tab">月份</a></li>
  <li><a href="#year" data-toggle="tab">年度</a></li>
  <li><a href="#proj" data-toggle="tab">项目</a></li>
</ul>
<div class="tab-content">
<div class="tab-pane fade in active" id="employee">
<br>

<div>
<form class="form-inline" method="get" action="/secofficeshow" enctype="multipart/form-data">
  <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
  <input type="hidden" id="level" name="level" value="{{.Level}}"/>
  <div class="form-group">
    <label for="taskNote">统计周期：</label>
    <input type="text" class="form-control" name="datefilter"  id="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
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
  <button type="submit" class="btn btn-primary" name="button">提交</button>
</form>
<br></div>

<div class="form-group">
<label class="control-label" id="regis" for="LoginForm-UserName">
  统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}
</label>
</div>
{{range $k,$v :=.Secoffice}}
<h2>{{.Name}}</h2>
<table class="table table-striped">

<thead>
  <tr>
    <th>#</th>
    <th>姓名</th>
    <th>编制</th>
    <th>设计</th>
    <th>校核</th>
    <th>审查</th>
    <th>汇总</th>
    <th>详细</th>
  </tr>
</thead>

<tbody>
  {{range $k1,$v1 :=.Employee}}
  <tr>
    <td>{{$k1|indexaddone}}</td>
    <td>{{.Name}}</td>
    <td>{{.Drawn}}</td>
    <td>{{.Designd}}</td>
    <td>{{.Checked}}</td>
    <td>{{.Examined}}</td>
    <td>{{.Sigma}}</td>
    <td>
      <a href="/secofficeshow?secid={{.Id}}&level=3"> <i class="glyphicon glyphicon-open"></i>
        详细
      </a>
    </td>
  </tr>
  {{end}}
</tbody>

</table>
{{end}}
</div>

    <div class="tab-pane fade" id="year">
      <p>这里将显示全年人员情况。</p>
    </div>

    <div class="tab-pane fade" id="proj">
      <p>这里将显示所有项目列表，每个项目的成本分布。</p>
    </div>

  </div>
</div>

<script type="text/javascript">
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
  });
// });
</script>
</body>
</html>
<!-- <button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span>
User
</button>

<button type="button" class="btn btn-primary btn-lg" style="text-shadow: black 5px 3px 3px;">
<span class="glyphicon glyphicon-user"></span>
User
</button>
-->