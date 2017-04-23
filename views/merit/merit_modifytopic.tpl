<!-- 用户修改价值内容界面 -->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>MeritMS</title>
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap-select.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-select.min.css"/>

    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.min.js"></script>
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

      <tr>
        <td></td>
        <td>{{.Topic.Title}}</td>
        <td>{{.Topic.Choose}}</td>
        <td>{{.Topic.Mark}}</td>
        <td>{{.Topic.Content}}</td>  
      </tr>

    </tbody>
  </table>

  <form method="post" action="/ModifyPost" enctype="multipart/form-data">
 <!--  <form method="post" action="/topic/addtopic1" enctype="multipart/form-data"> -->
    <div class="form-group">
      <label>{{.category.Title}}-名称：</label>
      <input id="name" type="text" class="form-control"  placeholder="输入名称" name="name" value="{{.Topic.Title}}"></div> 

    <div class="form-group">
    {{if gt (.list|len) 1}}
      <label>选项：</label>
        <th>
          <div class="form-group">
            <select id="cars" class="selectpicker" name="choose">
            <option>{{.Topic.Choose}}</option></select>
          </div>
        </th>
    {{end}}
    </div>
    <label>简介:</label>
    <div id="content">
    <script id="editor" type="text/plain" style="height:500px;"></script>
</div>

<hr>
      <input type="hidden" name="id" value="{{.Topic.Id}}"/>
      <button type="submit" class="btn btn-primary" onclick="return checkInput();"> 修 改 </button>
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

    var ue = UE.getEditor('editor');
ue.addListener("ready", function () {
uParse('.content', {
    rootPath: '/static/ueditor/'
});
});

$(function(){
        var content =$('#content').val();
        //判断ueditor 编辑器是否创建成功
        ue.addListener("ready", function () {
        // editor准备好之后才可以使用
        ue.setContent({{str2html .Topic.Content}});
        });
    });
// fireEvent("startUpload")

</script>
<script type="text/javascript">
// $(document).ready(function(){
//   var data={{.list}};
//      for ( var i = 0; i<data.length; i++) {  
//        $("#cars").append('<option value="' + data[i].text + '"></option>');
//        // alert(data[i].text)
//      }
//     }) 

$(document).ready(function(){
  var data={{.list}};
     for ( var i = 0; i<data.length; i++) {  
       $("#cars").append('<option>' + data[i].choose + '</option>');
       // alert(data[i].text)
     }
    }) 

window.onload=function(){
        $('.selectpicker').selectpicker({
          style: 'btn-info',
          size: 4
        });
      };
// $('.selectpicker').selectpicker('val',$('.selectpicker').attr('value'));
</script>

</body>
</html>



