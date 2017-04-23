<!-- iframe里展示总院总体情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>总院情况汇总</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

<script type="text/javascript" src="/static/js/moment.min.js"></script>
<script type="text/javascript" src="/static/js/daterangepicker.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />

</head>

<div class="container-fluid blue-bg-rpt">
  <div class="row">
<img style="-webkit-user-select: none; cursor: zoom-in;" src="/static/成果登记系统.png" width="100%" ><hr/>
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
