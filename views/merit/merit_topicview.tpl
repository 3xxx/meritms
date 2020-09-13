<!-- 用户价值内容详细查看界面 修改为模态框吧-->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>MeritMS</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <script src="/static/ueditor/ueditor.parse.min.js"></script>
</head>
<div class="col-lg-12">
  <table class="table table-striped">
    <thead>
      <tr>
        <!-- <th>#</th> -->
        <th>Title</th>
        <th>choose</th>
        <th>mark</th>
        <th>content</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>{{.Topic.Title}}</td>
        <td>{{.Topic.Choose}}</td>
        <td>{{.Topic.Mark}}</td>
        <td>{{.Topic.Content}}</td>
      </tr>
    </tbody>
  </table>
  <label>简介:</label>
  <div class="content">
    {{str2html .Topic.Content}}
    <!-- 项目简介如何截取html呢？ -->
  </div>
  <hr>
  <br />
  <br />
</div>
<script type="text/javascript">
// fireEvent("startUpload")
uParse('.content', {
  rootPath: '/static/ueditor/'
})
</script>
</body>

</html>