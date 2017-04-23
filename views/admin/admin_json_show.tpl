
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <!-- <title>MeritMS</title> -->
  <!-- <base target=_blank> -->
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

    <tbody>
      {{range $k,$v :=.Input.Fenyuan}}
      <tr>
        <td>{{$k|indexaddone}}</td>
        <td>{{.Department}}</td>
        <td></td>
        <td></td> 
        <td></td>
        
        <td>
          <!-- <input type="hidden" id="{{.Id}}" value="{{.Pid}}"/> -->
          <a href="" onclick="prom('{{.Pid}}')">添加同级</a><!-- href="/addjson?pid=0" -->
          <a href="" onclick="prom('{{.Id}}')">添加下级</a>
          <a href="/modifyjson?id={{.Id}}">修改</a>
          <a href="/deletejson?id={{.Id}}">删除</a>
        </td>
                  {{range $k1,$v1 :=$.Input.Fenyuan}}
                 {{range $k2,$v2 :=.Bumen}}
                 {{if eq $v2.Pid $v.Id}}
                 <tr>
                   <td></td>
                   <td></td>
                 <td>{{.Keshi}}</td>
                 <td></td>
                 <td></td>
                 
                 <td>
                  <!-- <input type="hidden" id="{{.Id}}" value="{{$v.Id}}"/> -->
                   <a href="" onclick="prom('{{$v.Id}}')">添加同级</a>
                   <a href="" onclick="prom('{{.Id}}')">添加下级</a>
                  <a href="/modifyjson?id={{.Id}}">修改</a>
                  <a href="/deletejson?id={{.Id}}">删除</a>
                 </td>
                 </tr>
                  {{range $k3,$v3 :=$.Input.Fenyuan}}
                 {{range $k4,$v4 :=.Bumen}}
                 {{range $k5,$v5 :=.Kaohe}}
                 {{if eq $v5.Pid $v2.Id}}

                 <tr>
                   <td></td>
                   <td></td>
                   <td></td>
                  <td>{{.Category}}</td>  
                  <td></td>
                                  
                  <td>
                  <!-- <input type="hidden" id="{{.Id}}" value="{{$v2.Id}}"/> -->
                  <a href="" onclick="prom('{{$v2.Id}}')">添加同级</a>
                  <a href="" onclick="prom('{{.Id}}')">添加下级</a>
                  <a href="/modifyjson?id={{.Id}}">修改</a>
                  <a href="/deletejson?id={{.Id}}">删除</a>
                  </td>
                 </tr>
                  {{range $k6,$v6 :=$.Input.Fenyuan}}
                 {{range $k7,$v7 :=.Bumen}}
                 {{range $k8,$v8 :=.Kaohe}}
                 {{range $k9,$v9 :=.Fenlei}}
                 {{if eq $v9.Pid $v5.Id}}
                  <tr>
                   <td></td>
                   <td></td>
                   <td></td>
                  <td></td>  
                  <td>{{.Project}}</td>
                                  
                  <td>
                  <a href="/modifyjson?id={{.Id}}">修改</a>
                  <a href="/deletejson?id={{.Id}}">删除</a>
                  </td>
                 </tr>

                  {{end}} 
                 {{end}}
                 {{end}} 
                  {{end}} 
                 {{end}}
                 {{end}} 
                 {{end}}
                 {{end}} 
                  {{end}} 
                 {{end}}
                 {{end}} 
                 {{end}}
      </tr>
      {{end}}

    </tbody>
  </table>
</div>


<script type="text/javascript">
$(function() {
         // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
          data: [{{.json}}],//defaultData,
          // data:alternateData,
          levels: 3,// expanded to 5 levels
          enableLinks:true,
          showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        });
});


//弹出一个输入框，输入一段文字，可以提交 
//添加同级/下级，通过id来区分统计或下级
    function prom(id) {  
        var name = prompt("请输入名称", ""); //将输入的内容赋给变量 name ，  
        //这里需要注意的是，prompt有两个参数，前面是提示的话，后面是当对话框出来后，在对话框里的默认值  
        if (name)//如果返回的有内容  
        {  
          // var pid = $('#'+id).val();
            // alert("欢迎您：" + name) 
            $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/addjson",
                data: {pid:id,title:name},
                success:function(data,status){//数据提交成功时返回数据
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        }  
    } 
 //弹出一个输入框，输入一段文字，可以提交
 //添加下级  
    // function prom1(id) {  
        // var name = prompt("请输入名称", ""); //将输入的内容赋给变量 name ，  
        //这里需要注意的是，prompt有两个参数，前面是提示的话，后面是当对话框出来后，在对话框里的默认值  
        // if (name)//如果返回的有内容  
        // {  
          // var pid = $('#'+id).val();
            // alert("欢迎您：" + name) 
            // $.ajax({
                // type:"post",//这里是否一定要用post？？？
                // url:"/addjson",
                // data: {pid:id,title:name},
                // success:function(data,status){//数据提交成功时返回数据
                  // alert("添加“"+data+"”成功！(status:"+status+".)");
                 // }
            // });  
        // }  
  
    // }   
$(document).ready(function(){//这个没有用到吧
  // var roletitle1=document.getElementsByName("roletitle");
  // $("#uname").focus(function(){获得焦点
     $("input").blur(function(){//其失去焦点
        var pwd=document.getElementsByName("password");
        var nickname=document.getElementsByName("nickname");
        var email=document.getElementsByName("email");
        // var roletitle2=document.getElementsByName("roletitle");
        // alert(pwd[0].value.length);//什么时候是逗号，什么时候是分号？
        if (pwd[0].value.length<1)
        {
          // alert("请输入密码。")
             $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/user/UpdateUser",
                data: { userid:{{.User.Id}},username:{{.User.Username}},nickname: nickname[0].value,email: email[0].value},
                success:function(data,status){//数据提交成功时返回数据
                  // alert(data,status);
                 }
            });         
        }else{
            $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/user/UpdateUser",
                data: { userid:{{.User.Id}},username:{{.User.Username}},password: pwd[0].value,nickname: nickname[0].value,email: email[0].value},
                success:function(data,status){//数据提交成功时返回数据
                  alert('success modified~');
                  // alert(data,status);
                 }
            });
       }     
 });
});
</script>
</body>
</html>
