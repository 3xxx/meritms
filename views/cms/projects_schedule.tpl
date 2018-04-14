<!DOCTYPE HTML>

<html>
<head>
  <meta http-equiv="X-UA-Compatible" content="IE=9; IE=8; IE=7; IE=EDGE"/>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
  <title>项目进度</title>
  </head>

<script type="text/javascript">
    //使用ajax加载动态列的
var columns = [];
$.ajax({
url: 'getColumns.action',
type: 'post',
data: data,
dataType: "json",
async: true,
success: function (returnValue) {
//异步获取要动态生成的列
var arr = returnValue;
$.each(arr, function (i, item) {
columns.push({ "field": item.colname, "title": item.colalias, "width": 100, "sortable": true });
});
});
$('#table_id').bootstrapTable('destroy').bootstrapTable({
data: data,
columns: columns
...
...
})
</script>  

</body>
</html>