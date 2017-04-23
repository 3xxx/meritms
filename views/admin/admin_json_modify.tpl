<!-- 用户添加价值内容界面 -->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <!-- <title>MeritMS</title> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap-select.min.js"></script>
 <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
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
      </tr>
    </thead>

    <tbody>
      <tr>
        <td></td>
        <td>{{.category.Title}}</td>
        <td>{{.category.List}}</td>
        <td>{{.category.ListMark}}</td>
      </tr>
    </tbody>
  </table>

  <form method="post" action="/modifyjsonpost" enctype="multipart/form-data">
 <!--  <form method="post" action="/topic/addtopic1" enctype="multipart/form-data"> -->
    <div class="form-group">
      <label>{{.category.Title}}-名称：</label>
      <input id="name" type="text" class="form-control"  placeholder="输入名称" name="title" value="{{.category.Title}}"></div> 
    <div class="form-group">
      <label>{{.category.Title}}-选项：</label>
      <input id="choose" type="text" class="form-control"  placeholder="输入选项，例子：,大型,中型,小型——要求为英文,号，并且是,号开始" name="list" value="{{.category.List}}"></div> 
    <div class="form-group">
      <label>{{.category.Title}}-分值：</label>
      <input id="mark1" type="text" class="form-control"  placeholder="输入分值，例子：,10,5,2——要求为英文,号，并且是,号开始" name="listmark" value="{{.category.ListMark}}"></div> 
    <div class="form-group">
      <label>{{.category.Title}}-单一分值：</label>
      <input id="mark2" type="text" class="form-control"  placeholder="输入唯一分值——用于分值分类层" name="mark" value="{{.category.Mark}}"></div>

    <div class="form-group">
    {{if gt (.list|len) 1}}
      <label>预览选项效果：</label>
        <th>
          <div class="form-group">
            <select id="list" class="selectpicker" name="choose">
            </select>
          </div>
        </th>
      <label>预览分值效果：</label>
        <th>
          <div class="form-group">
            <select id="mark3" class="selectpicker" name="mark3">
            </select>
          </div>
        </th>  
    {{end}}
    </div>
<hr>
      <input type="hidden" name="id" value="{{.category.Id}}"/>
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
       $("#list").append('<option>' + data[i].choose + '</option>');
       // alert(data[i].text)
     }
    }) 
$(document).ready(function(){
  var data1={{.mark}};
     for ( var i = 0; i<data1.length; i++) {  
       $("#mark3").append('<option>' + data1[i].mark1 + '</option>');
       // alert(data1[i])
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



