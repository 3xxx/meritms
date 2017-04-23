<!-- 个人登录价值系统，展示价值侧栏——将来修改为管理员目录-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>Merit价值管理系统</title>

<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
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


<div id="treeview" class="col-xs-3"></div>
<div class="col-lg-9">
  <iframe src="/merit/myself" name='main' frameborder="0"  width="100%" scrolling="no" marginheight="0" marginwidth="0" id="iframepage" onload="this.height=100"></iframe> 
</div>

<script type="text/javascript">
 function reinitIframe(){//http://caibaojian.com/frame-adjust-content-height.html
  var iframe = document.getElementById("iframepage");
   try{
    var bHeight = iframe.contentWindow.document.body.scrollHeight;
     var dHeight = iframe.contentWindow.document.documentElement.scrollHeight; var height = Math.max(bHeight, dHeight); iframe.height = height;
      // console.log(height);//这个显示老是在变化
       }catch (ex){
        } 
        } 
        window.setInterval("reinitIframe()", 200);
</script>

<script type="text/javascript">
$(function() {
         // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
          data: [{{.json}}],//defaultData,
          // data:alternateData,
          levels: 5,// expanded to 5 levels
          enableLinks:true,
          showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        });

        $('#treeview').on('nodeSelected', function(event, data) {
            // alert("名称："+data.text);
            // alert("节点id："+data.nodeId);
            // alert("部门id："+data.Id);  
            // alert("部门级别："+data.Level);
            $("#regis").html(data.text);//显示部门名称
            $("#regis").css("color","black");
            document.getElementById("iframepage").src="/merit/myself?id="+data.Id+"&level="+data.Level;
        });
});
</script>
</body>
</html>