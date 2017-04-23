<!-- 用户添加价值内容界面 -->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>MeritMS</title>
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
 <script type="text/javascript" src="/static/js/bootstrap-select.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-select.min.css"/>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
</head>

<div class="col-lg-12">
<table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>Title</th>
        <th>choose</th>
        <th>mark</th>
        <th>content</th>
      </tr>
    </thead>

    <tbody>
      {{range $k,$v :=.topics}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td><a href="/view?id={{.Id}}">{{.Title}}</a></td>
        <td>{{.Choose}}</td>
        <td>{{.Mark}}</td>
        <td>{{.Content}}</td>  
      </tr>
      {{end}}
    </tbody>
  </table>

  <form method="post" action="/AddMeritTopic" enctype="multipart/form-data">
 <!--  <form method="post" action="/topic/addtopic1" enctype="multipart/form-data"> -->
    <div class="form-group">
      <label>{{.category.Title}}-名称：</label>
      <input id="name" class="form-control"  placeholder="输入名称" name="name"></div> 

    <div class="form-group">
    {{if gt (.list|len) 1}}
      <label>选项：</label>
        <th>
          <div class="form-group">
            <select id="list" class="selectpicker" name="choose">
            <option></option></select>
          </div>
        </th>
    {{end}}
    </div>

    <label>简介:</label>
<div>
    <!-- <h1>项目简介:</h1> -->
    <script id="editor" type="text/plain" style="height:500px;"></script><!-- width:1024px; -->
</div>
<hr>
      <input type="hidden" name="id" value="{{.category.Id}}"/>
      <button type="submit" class="btn btn-primary" onclick="return checkInput();"> 添  加 </button>
  <!--必须加return才能不跳转-->
</form>
<br />
<br />
</div>

<script type="text/javascript">
// document.getElementById()返回对拥有指定 id 的第一个对象的引用。
// document.getElementsByName()返回带有指定名称的对象集合。
// document.getElementsByTagName()返回带有指定标签名的对象集合。
   function checkInput(){
    //是下面这段代码出了问题，等下修改
      var name=document.getElementById("name");
      if (name.value.length==0){
        alert("请输入名称");
        return false;
      }
    return true;  //这个return必须放最后，前面的值才能传到后台
   }
    //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    var ue = UE.getEditor('editor');
    /* 1.传入函数,命令里执行该函数得到参数表,添加到已有参数表里 */
    ue.ready(function () {
    ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
      var name = $('#name').val();
      // var number = $('#number').val();
        ue.execCommand('serverparam', {
        // "number":number,
        'name': name,
        });
      });
    });
// fireEvent("startUpload")
</script>
<script type="text/javascript">
$(document).ready(function(){
  var data={{.list}};
     for ( var i = 0; i<data.length; i++) {  
       $("#list").append('<option>' + data[i].choose + '</option>');
       // alert(data[i].text)
     }
    }) 
</script>

</body>
</html>


