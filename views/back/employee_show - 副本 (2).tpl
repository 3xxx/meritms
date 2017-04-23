<!-- iframe里展示个人详细情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>情况汇总</title>

<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />
</head>

<div class="form-group">
        <label class="control-label" id="regis" for="LoginForm-UserName">{{.UserNickname}}</label><!-- 显示部门名称 -->
    </div>
<div class="col-lg-12">

<div>
<form class="form-inline" method="get" action="/secofficeshow" enctype="multipart/form-data">
  <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
  <input type="hidden" id="level" name="level" value="{{.Level}}"/>
  <div class="form-group">
    <label for="taskNote">统计周期：</label>
    <input type="text" class="form-control" name="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
  </div>
  <script type="text/javascript">
$(function() {
  $('input[name="datefilter"]').daterangepicker({
      autoUpdateInput: false,
      locale: {
          cancelLabel: 'Clear'
      }
  });
  $('input[name="datefilter"]').on('apply.daterangepicker', function(ev, picker) {
      $(this).val(picker.startDate.format('YYYY-MM-DD') + ' - ' + picker.endDate.format('YYYY-MM-DD'));
  });
  $('input[name="datefilter"]').on('cancel.daterangepicker', function(ev, picker) {
      $(this).val('');
  });
});
</script>
  <button type="submit" class="btn btn-primary" name="button">提交</button>
</form>
<br></div>

<div class="form-group">
<label class="control-label" id="regis" for="LoginForm-UserName">
  统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}
</label>
</div>
<h3>已完成</h3>
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>项目编号</th>
        <th>项目名称</th>
        <th>项目阶段</th>
        <th>成果编号</th>
        <th>成果名称</th>
        <th>成果类型</th>
        <th>成果计量单位</th>
        <th>成果数量</th>
        <th>编制、绘制</th>
        <th>设计</th>
        <th>校核</th>
        <th>审查</th>
      </tr>
    </thead>

    <tbody>
    {{range $k,$v :=.Ratio}}
      <tr><th colspan=13>{{$v.Category}}</th></tr>
      {{range $k1,$v1 :=$.Catalogs}}
      {{if eq $v.Category $v1.Category}}
      {{if eq $v1.State "4"}}
      <tr>
        <td>{{$k1|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>
        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Drawn }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td> 
      </tr>
      {{end}}
      {{end}}
      {{end}}
    {{end}}
    </tbody>
  </table>

  <div class="form-group">   
      <input type="button" class="btn btn-primary" name="insert" value="在线添加" onclick="insertNewRow()"/>
          <form id="form1" class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
            <div class="form-group">
              <label>选择成果登记数据文件(Excel)
              <input type="file" class="form-control" name="catalog" id="catalog"></label>
              <br/>
              </div>
            <button type="submit" class="btn btn-primary" onclick="return import_xls_catalog();">提交</button>
          </form>
    </div> 

<h3>需要提交给校核</h3>
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>项目编号</th>
        <th>项目名称</th>
        <th>项目阶段</th>
        <th>成果编号</th>
        <th>成果名称</th>
        <th>成果类型</th>
        <th>成果计量单位</th>
        <th>成果数量</th>
        <th>编制、绘制</th>
        <th>设计</th>
        <th>校核</th>
        <th>审查</th>
      </tr>
    </thead>

    <tbody>
      {{range $k1,$v1 :=$.Catalogs}}
      {{if eq $v1.State "1" "2"}}
      {{if eq $.UserNickname $v1.Drawn $v1.Designd}}
      <tr>
        <td>{{$k1|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>
        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Drawn }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td> 
      </tr>
      {{end}}
      {{end}}
      {{end}}

    </tbody>
  </table>
  <h3>需要处理校核</h3>
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>项目编号</th>
        <th>项目名称</th>
        <th>项目阶段</th>
        <th>成果编号</th>
        <th>成果名称</th>
        <th>成果类型</th>
        <th>成果计量单位</th>
        <th>成果数量</th>
        <th>编制、绘制</th>
        <th>设计</th>
        <th>校核</th>
        <th>审查</th>
      </tr>
    </thead>

    <tbody>
      {{range $k1,$v1 :=$.Catalogs}}
      {{if eq $v1.State "2"}}
      {{if eq $.UserNickname $v1.Checked}}
      <tr>
        <td>{{$k1|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>
        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Drawn }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td> 
      </tr>
      {{end}}
      {{end}}
      {{end}}

    </tbody>
  </table>
  <h3>需要处理审查</h3>
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>项目编号</th>
        <th>项目名称</th>
        <th>项目阶段</th>
        <th>成果编号</th>
        <th>成果名称</th>
        <th>成果类型</th>
        <th>成果计量单位</th>
        <th>成果数量</th>
        <th>编制、绘制</th>
        <th>设计</th>
        <th>校核</th>
        <th>审查</th>
      </tr>
    </thead>

    <tbody>
      {{range $k1,$v1 :=$.Catalogs}}
      {{if eq $v1.State "3"}}
      {{if eq $.UserNickname $v1.Examined}}
      <tr>
        <td>{{$k1|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>
        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Drawn }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td> 
      </tr>
      {{end}}
      {{end}}
      {{end}}

    </tbody>
  </table>
<tr>    
       <td colspan="4"><input type="button" class="btn btn-primary" name="insert" value="处&nbsp;&nbsp;&nbsp;&nbsp;理" onclick="ModifyRow()"/></td>    
       </tr>
</div>

<script type="text/javascript">
function import_xls_catalog(){
  var form1 = window.document.getElementById("form1");//获取form1对象
  form1.submit();
  $.ajax({
                        success:function(data,status){//数据提交成功时返回数据
                        // alert("添加“"+data+"”成功！(status:"+status+".)");
                        window.location.reload();
                        }
                    });
    return true;  //这个return必须放最后，前面的值才能传到后台    
   }

function insertNewRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        window.open('/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}}+'&key=editor');
        }
 function ModifyRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        window.open('/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}}+'&key=modify');
        }       

</script>
</body>
</html>
