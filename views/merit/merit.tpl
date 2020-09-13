<!-- 这个是显示左侧栏，右边显示secoffice和employee或employeework -->
<!DOCTYPE html>
<html>
<!-- <head> -->
{{template "tpl/T.header.tpl"}}
<meta charset="UTF-8">
<title>MeritMS</title>
<!-- <base target=_blank>
  -->
<!-- <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script> -->
<!-- <script type="text/javascript" src="/static/js/bootstrap.min.js"></script> -->
<script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css" />
<!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/> -->
<style type="text/css">
a:active {
  text: expression(target="_blank");
}

i#delete {
  color: #DC143C;
}
</style>
<script type="text/javascript">
var allLinks = document.getElementsByTagName("a");
for (var i = 0; i != allLinks.length; i++) {
  allLinks[i].target = "_blank";
}
</script>
</head>
<!-- <div class="navbar navba-default navbar-fixed-top"> -->
<div class="container-fill">{{template "tpl/T.navbar.tpl" .}}</div>
<!-- </div> -->

<body>
  <!-- 侧栏 -->
  <div id="treeview" class="col-xs-3 col-sm-3 col-md-3 col-lg-3"></div>
  <!-- <a href="/category/view?id={{.Id}}" target='main'> -->
  <div class="col-xs-9 col-sm-9 col-md-9 col-lg-9">
    <!-- <div class="form-group"> -->
    <!-- <label class="control-label" id="regis" for="LoginForm-UserName"></label> 显示部门名称  -->
    <!-- </div> -->
    <!-- <iframe src="/secofficeshow" name='main' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onLoad="iFrameHeight()"></iframe> -->
    <!-- <iframe src="/secofficeshow" name='main' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onload="changeFrameHeight()"></iframe> -->
    <iframe src="/merit/secofficeshow" name='main' frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" id="iframepage" onload="this.height=100"></iframe>
  </div>
  <script type="text/javascript">
  function reinitIframe() { //http://caibaojian.com/frame-adjust-content-height.html
    var iframe = document.getElementById("iframepage");
    try {
      var bHeight = iframe.contentWindow.document.body.scrollHeight;
      var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
      var height = Math.max(bHeight, dHeight, 800);
      iframe.height = height;
      // console.log(height);//这个显示老是在变化
    } catch (ex) {}
  }
  window.setInterval("reinitIframe()", 200);
  </script>
  <!--   <div class="col-lg-9">
    <table class="table table-striped">
      <thead>
        <tr>
          <th>#序号</th>
          <th>部门</th>
          <th>科室</th>
          <th>价值分类</th>
          <th>价值</th>

          <th>操作</th>
        </tr>
      </thead>

      <tbody></tbody>
    </table>
  </div> -->
  <script type="text/javascript">
  $(function() {
    // alert(JSON.stringify({{.json}}));
    // $('#treeview').treeview('collapseAll', { silent: true });
    $('#treeview').treeview({
      data: {{.json }}, //[{{.json}}]——有时候加个中括号就行了。defaultData,
      // data:alternateData,
      levels: 2, // expanded to 5 levels
      enableLinks: true,
      showTags: true,
      // collapseIcon:"glyphicon glyphicon-chevron-up",
      // expandIcon:"glyphicon glyphicon-chevron-down",
    });
    $('#treeview').on('nodeSelected', function(event, data) {
      // alert("名称："+data.text);
      // alert("节点id："+data.nodeId);
      // alert("部门id："+data.Id);  
      // alert("部门级别："+data.Level);
      $("#regis").html(data.text); //显示部门名称
      $("#regis").css("color", "black");
      document.getElementById("iframepage").src = "/merit/secofficeshow?secid=" + data.Id + "&level=" + data.Level;
    });
  });
  // document.getElementById()返回对拥有指定 id 的第一个对象的引用。
  // document.getElementsByName()返回带有指定名称的对象集合。
  // document.getElementsByTagName()返回带有指定标签名的对象集合。
  // 自动适应高度 

  // function iFrameHeight() {   
  // var ifm= document.getElementById("iframepage");   
  // var subWeb = document.frames ? document.frames["iframepage"].document : ifm.contentDocument;   
  // if(ifm != null && subWeb != null) {
  //    ifm.height = subWeb.body.scrollHeight;
  //    ifm.width = subWeb.body.scrollWidth;
  // }   
  // }   

  // function changeFrameHeight(){
  //     var ifm= document.getElementById("iframepage"); 
  //     ifm.height=document.documentElement.clientHeight;
  // }
  // window.onresize=function(){  
  //      changeFrameHeight();  
  // } 

  // document.getElementById("changeUrl").onclick = function(){
  //     document.getElementById("iframepage").src="http://www.baidu.com";
  // }
  //弹出一个输入框，输入一段文字，可以提交 
  //添加同级/下级，通过id来区分统计或下级
  // function prom(id) {  
  //     var name = prompt("请输入名称", ""); //将输入的内容赋给变量 name ，  
  //     //这里需要注意的是，prompt有两个参数，前面是提示的话，后面是当对话框出来后，在对话框里的默认值  
  //     if (name)//如果返回的有内容  
  //     {  
  //       // var pid = $('#'+id).val();
  //         // alert("欢迎您：" + name) 
  //         $.ajax({
  //             type:"post",//这里是否一定要用post？？？
  //             url:"/addjson",
  //             data: {pid:id,title:name},
  //             success:function(data,status){//数据提交成功时返回数据
  //               alert("添加“"+data+"”成功！(status:"+status+".)");
  //              }
  //         });  
  //     }  
  // }
  </script>
</body>

</html>