<!-- 管理员查看用户--> 
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

</head>
<body>


<input type="button" id="btn_addtr" value="增行">
<div class="col-lg-12">
<table class="table table-striped">
  <thead>
      <tr>
      <!-- <th><span style="cursor: pointer">Id</span></th> -->
      <th><span style="cursor: pointer">Username</span></th>
      <th><span style="cursor: pointer">Password</span></th>
      <th><span style="cursor: pointer">Nickname</span></th>
      <th><span style="cursor: pointer">Email</span></th>
      <!-- <th><span style="cursor: pointer">Remark</span></th> -->
      <!-- <th><span style="cursor: pointer">Status</span></th> -->
      <th><span style="cursor: pointer">Lastlogintime</span></th>
      <th><span style="cursor: pointer">Createtime</span></th>
      <!-- <th><span style="cursor: pointer">RoleId</span></th> -->
      <th><span style="cursor: pointer">RoleTitle</span></th>
      <th><span style="cursor: pointer">RoleName</span></th>
      <th><span style="cursor: pointer">Remark</span></th>      
      <th>操作</th>
    </tr>
  </thead>

    <tr id="row{{.User.Id}}">
      <td>{{.User.Username}}</td>
      <td><input type="password" id="input" name="password"  size='18'/></td>
      <td><input type="text" id="input" name="nickname" value="{{.User.Nickname}}" size='6'/></td>
      <td><input type="text" id="input" name="email" value="{{.User.Email}}" size='20'/></td>
      <td>{{dateformat .User.Lastlogintime "2006-01-02 T 15:04:05"}}</td>
      <td>{{dateformat .User.Createtime "2006-01-02 T 15:04:05"}}</td>
      <td><input type="text" id="input" name="roletitle2" value="{{.User.Role}}" size='6'/></td>
      <td>{{.User.Nickname}}</td>
      <td>{{.User.Role}}</td>
      <td><input type="button" id="btn_deltr" onclick="deltr()" value="删行"></td>
    </tr>
      <input type="hidden" id="input" name="roletitle1" value="{{.User.Role}}"/>

 </table>

</div>

</body>
</html>
