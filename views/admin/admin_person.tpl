<!-- 展示所有用户的价值排名-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <!-- <title>MeritMS</title> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
<!-- jquery一定要放前面 -->
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<style type="text/css">
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
</script>
</head>

<!-- <div id="treeview" class="col-xs-3"></div> -->

<div class="col-lg-10">
  <table class="table table-striped">
    <thead>
      <tr>
        <th><span style="cursor: pointer">#</span></th>
        <th><span style="cursor: pointer">Name</span></th>
        <th><span style="cursor: pointer">Numbers</span></th>
        <th><span style="cursor: pointer">Marks</span></th>
        <th><span style="cursor: pointer">Secoffic</span></th>
        <th><span style="cursor: pointer">Department</span></th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      {{range $k,$v :=.person}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.Name}}</td>
        <td>{{.Numbers}}</td>
        <td>{{.Marks}}</td>
        <td>{{.Keshi}}</td>
        <td>{{.Department}}</td>
        <td>
         <a href="/user?uid={{.Id}}"><i class="glyphicon glyphicon-open"></i>详细</a>
        </td>  
      </tr>
      {{end}}
    </tbody>

  </table>
</div>

<!-- <script type="text/javascript">
$(function() {
          $('#treeview').treeview({
          data: [{{.json}}],
          levels: 5,
          enableLinks:true,
          showTags:true,
        });
});
</script> -->
<script type="text/javascript">
  $(document).ready(function() {
  $("table").tablesorter( {sortList: [[5,1],[4,1], [3,1]]} );//[列索引,排序方向] 0 asc 1 desc 
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
