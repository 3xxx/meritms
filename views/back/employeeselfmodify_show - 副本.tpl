<!-- iframe里展示个人待处理的详细情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>待处理成果</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />

  <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script>
<script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.js"></script>
<link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker3.css"/> 
<!-- <style type="text/css">
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
</script> -->
</head>


<!-- <div id="treeview" class="col-xs-3"></div> -->

<div class="col-lg-12">
<div class="form-group">
        <label class="control-label" id="regis" for="LoginForm-UserName">{{.UserNickname}}</label><!-- 显示部门名称 -->
</div>
<div>
<form class="form-inline" method="get" action="/secofficeshow" enctype="multipart/form-data">
  <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
  <input type="hidden" id="level" name="level" value="{{.Level}}"/>
  <input type="hidden" id="key" name="key" value="modify"/>
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
  <button type="submit" class="btn btn-primary">提交</button>
</form>
<br></div>

<div class="form-group">
<label class="control-label" id="regis" for="LoginForm-UserName">
  统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}
</label>
</div>
<h3>需要提交给校核</h3>
  <table class="table table-striped" id="orderTable" name="orderTable">
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
        <th>绘制系数</th>
        <th>出版</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
    {{range $k1,$v1 :=$.Catalogs}}
      {{if eq $v1.State "1"}}
      {{if eq $.UserNickname $v1.Drawn $v1.Designd}}
      <tr id="row{{.Id}}">
        <td>{{$k1|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>
        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count}}</td>
        <td>{{.Drawn}}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>
        <td>{{.Drawnratio}}</td>
        <td>{{dateformat .Data "2006-01-02"}}</td>
        <td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow("row{{.Id}}")'/> 
        <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow("row{{.Id}}")' />
        <input type='button' class='btn btn-default' name='update' value='提交' onclick='sendSelectedRow("row{{.Id}}")' /></td> 
      </tr>
      {{end}}
      {{end}}
      {{end}}
    </tbody>
  </table>

  <h3>需要处理校核</h3>
  <table class="table table-striped" id="orderTable" name="orderTable">
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
        <th>设计</th>
        <th>校核</th>
        <th>设计系数</th>
        <th>校核系数</th>
        <th>出版</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
    {{range $k1,$v1 :=$.Catalogs}}
      {{if eq $v1.State "2"}}
      {{if eq $.UserNickname $v1.Checked}}
      <tr id="row{{.Id}}">
        <td>{{$k1|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>
        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Designdratio}}</td>
        <td>{{.Checkedratio}}</td>
        <td>{{dateformat .Data "2006-01-02"}}</td>
        <td><input type='button' class='btn btn-default' name='delete' value='退回' onclick='downsendSelectedRow("row{{.Id}}")'/> 
        <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow2("row{{.Id}}")' />
        <input type='button' class='btn btn-default' name='update' value='提交' onclick='sendSelectedRow("row{{.Id}}")' /></td> 
      </tr>
      {{end}}
      {{end}}
      {{end}}
    </tbody>
  </table>
  <h3>需要处理审查</h3>
  <table class="table table-striped" id="orderTable" name="orderTable">
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
        <th>难度系数</th>
        <th>校核</th>
        <th>审查</th>
        <th>校核系数</th>
        <th>审查系数</th>
        <th>出版</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
    {{range $k1,$v1 :=$.Catalogs}}
      {{if eq $v1.State "3"}}
      {{if eq $.UserNickname $v1.Examined}}
      <tr id="row{{.Id}}">
        <td>{{$k1|indexaddone}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>
        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count}}</td>
        <td>{{.Complex}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>
        <td>{{.Checkedratio}}</td>
        <td>{{.Examinedratio}}</td>
        <td>{{dateformat .Data "2006-01-02"}}</td>
        <td><input type='button' class='btn btn-default' name='delete' value='退回' onclick='downsendSelectedRow("row{{.Id}}")'/> 
        <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow3("row{{.Id}}")' />
        <input type='button' class='btn btn-default' name='update' value='提交' onclick='sendSelectedRow("row{{.Id}}")' /></td> 
      </tr>
      {{end}}
      {{end}}
      {{end}}
    </tbody>
  </table>
  <!-- <input type="hidden" id="CategoryId" name="CategoryId" value="{{.CategoryId}}"/> -->
     
</div>



<script type="text/javascript">
//*********这个是编辑表格
var flag = 0;  //标志位，标志第几行  
         /*    
         *删除选中的行    
         */    
         function deleteSelectedRow(rowId){    
            //根据rowId查询出该行所在的行索引    
            if(confirm("确定删除该行吗？")){    
                $("#"+rowId).remove();    //这里需要注意删除一行之后 我的标志位没有-1，因为如果减一，那么我再增加一行的话，可能会导致我的tr的id重复，不好维护。
                // 提交到后台进行删除数据库
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/delete",
                    data: {CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("删除“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
            }       
         }    
          /*    
         *退回选中的行    
         */
         function downsendSelectedRow(rowId){
          if(confirm("确定退回该行吗？")){    
                $("#"+rowId).remove();    //这里需要注意删除一行之后 我的标志位没有-1，因为如果减一，那么我再增加一行的话，可能会导致我的tr的id重复，不好维护。
                // 提交到后台进行修改数据库状态为降一个
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/downsendcatalog",
                    data: {CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("退回“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
            }   
         }
         /*    
          *用户直接传递行   
          */
        function sendSelectedRow(rowId){
          if(confirm("确定提交该行吗？")){    
                $("#"+rowId).remove();    //这里需要注意删除一行之后 我的标志位没有-1，因为如果减一，那么我再增加一行的话，可能会导致我的tr的id重复，不好维护。
                // 提交到后台进行修改数据库状态为降一个
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/sendcatalog",
                    data: {CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("提交“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
            }   
         }
         /*    
          *修改-提交给校核-选中的行    
          */  
         function updateSelectedRow(rowId){
            var oldIndex = $("#"+rowId+" td:eq(0)").html();
            var oldPnumber = $("#"+rowId+" td:eq(1)").html();  
            var oldPname = $("#"+rowId+" td:eq(2)").html();  
            var oldStage = $("#"+rowId+" td:eq(3)").html();
            var oldTnumber = $("#"+rowId+" td:eq(4)").html();
            var oldName = $("#"+rowId+" td:eq(5)").html();
            var oldCategory = $("#"+rowId+" td:eq(6)").html();
            var oldPage = $("#"+rowId+" td:eq(7)").html();
            var oldCount = $("#"+rowId+" td:eq(8)").html();
            var oldDrawn = $("#"+rowId+" td:eq(9)").html();
            var oldDesignd = $("#"+rowId+" td:eq(10)").html();
            var oldChecked = $("#"+rowId+" td:eq(11)").html();
            var oldExamined = $("#"+rowId+" td:eq(12)").html();
            var oldDrawnratio = $("#"+rowId+" td:eq(13)").html();
            var oldData = $("#"+rowId+" td:eq(14)").html();
            // if(oldPrice != ""){//去掉第一个人民币符号  
            //     oldPrice = oldPrice.substring(1);  
            // }  
            var uploadStr = "<td><input type='text' id='txtIndex"+flag+"' value='"+oldIndex+"' size='1'/></td>"
                        + "<td><input type='text' id='txtPnumber"+flag+"' value='"+oldPnumber+"' size='3'/></td>"  
                        + "<td><input type='text' id='txtPname"+flag+"' value='"+oldPname+"' size='14'/></td>"  
                        + "<td><input type='text' id='txtStage"+flag+"' value='"+oldStage+"' size='3'/></td>"
                        + "<td><input type='text' id='txtTnumber"+flag+"' value='"+oldTnumber+"' size='12'/></td>"
                        + "<td><input type='text' id='txtName"+flag+"' value='"+oldName+"' size='20'/></td>"
                        + "<td><input type='text' id='txtCategory"+flag+"' value='"+oldCategory+"' size='4'/></td>"
                        + "<td><input type='text' id='txtPage"+flag+"' value='"+oldPage+"' size='1'/></td>"
                        + "<td><input type='text' id='txtCount"+flag+"' value='"+oldCount+"' size='1'/></td>"
                        + "<td><input type='text' id='txtDrawn"+flag+"' value='"+oldDrawn+"' size='2'/></td>"
                        + "<td><input type='text' id='txtDesignd"+flag+"' value='"+oldDesignd+"' size='2'/></td>"
                        + "<td><input type='text' id='txtChecked"+flag+"' value='"+oldChecked+"' size='2'/></td>"
                        + "<td><input type='text' id='txtExamined"+flag+"' value='"+oldExamined+"' size='2'/></td>"
                        + "<td><input type='text' id='txtDrawnratio"+flag+"' value='"+oldDrawnratio+"' size='2'/></td>"
                        + "<td><input type='text' id='txtData"+flag+"' class='datepicker' value='"+oldData+"'/></td>"
                        + "<td><input type='button' class='btn btn-default' name='update' value='保存' onclick='saveUpdateRow(\""+rowId+"\",\""+flag+"\")'/></td>";  
            $("#"+rowId).html(uploadStr); 
            $(".datepicker").datepicker({
               language: "zh-CN",
               autoclose: true,//选中之后自动隐藏日期选择框
               clearBtn: true,//清除按钮
               todayBtn: 'linked',//今日按钮
               format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
            }); 
         }    
  
         /*    
          *保存提交给校核修改行    
          */
        function saveUpdateRow(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).val();
            var newPnumber = $("#txtPnumber"+flag).val();    
            var newPname = $("#txtPname"+flag).val();    
            var newStage = $("#txtStage"+flag).val();
            var newTnumber = $("#txtTnumber"+flag).val();
            var newName = $("#txtName"+flag).val();
            var newCategory = $("#txtCategory"+flag).val();
            var newPage = $("#txtPage"+flag).val();
            var newCount = $("#txtCount"+flag).val();
            var newDrawn = $("#txtDrawn"+flag).val();
            var newDesignd = $("#txtDesignd"+flag).val();
            var newChecked = $("#txtChecked"+flag).val();
            var newExamined = $("#txtExamined"+flag).val();
            var newDrawnratio = $("#txtDrawnratio"+flag).val();
            var newData = $("#txtData"+flag).val();
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newPnumber + "</td>"  
                        + "<td>" + newPname + "</td>"  
                        + "<td>" + newStage + "</td>"
                        + "<td>" + newTnumber + "</td>"
                        + "<td>" + newName + "</td>"
                        + "<td>" + newCategory + "</td>"
                        + "<td>" + newPage + "</td>"
                        + "<td>" + newCount + "</td>"
                        + "<td>" + newDrawn + "</td>"
                        + "<td>" + newDesignd + "</td>"
                        + "<td>" + newChecked + "</td>"
                        + "<td>" + newExamined + "</td>"
                        + "<td>" + newDrawnratio + "</td>"
                        + "<td>" + newData + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow(\""+rowId+"\")' /><input type='button' class='btn btn-default' name='update' value='提交' onclick='sendSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newName)//如果返回的有内容  
                {  
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/modifycatalog",
                    data: {Pnumber:newPnumber,Pname:newPname,Stage:newStage,Tnumber:newTnumber,Name:newName,Category:newCategory,Page:newPage,Count:newCount,Drawn:newDrawn,Designd:newDesignd,Checked:newChecked,Examined:newExamined,Drawnratio:newDrawnratio,Data:newData,CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("修改“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }
          /*    
          *修改-提交给审查-选中的行    
          */ 
         function updateSelectedRow2(rowId){
            var oldIndex = $("#"+rowId+" td:eq(0)").html();
            var oldPnumber = $("#"+rowId+" td:eq(1)").html();  
            var oldPname = $("#"+rowId+" td:eq(2)").html();  
            var oldStage = $("#"+rowId+" td:eq(3)").html();
            var oldTnumber = $("#"+rowId+" td:eq(4)").html();
            var oldName = $("#"+rowId+" td:eq(5)").html();
            var oldCategory = $("#"+rowId+" td:eq(6)").html();
            var oldPage = $("#"+rowId+" td:eq(7)").html();
            var oldCount = $("#"+rowId+" td:eq(8)").html();
            var oldDesignd = $("#"+rowId+" td:eq(9)").html();
            var oldChecked = $("#"+rowId+" td:eq(10)").html();
            var oldDesigndratio = $("#"+rowId+" td:eq(11)").html();
            var oldCheckedratio = $("#"+rowId+" td:eq(12)").html();
            var oldData = $("#"+rowId+" td:eq(13)").html();
            // if(oldPrice != ""){//去掉第一个人民币符号  
            //     oldPrice = oldPrice.substring(1);  
            // }  
            var uploadStr = "<td id='txtIndex"+flag+"'>"+oldIndex+"</td>"
                        + "<td id='txtPnumber"+flag+"'>"+oldPnumber+"</td>"  
                        + "<td id='txtPname"+flag+"'>"+oldPname+"</td>"  
                        + "<td id='txtStage"+flag+"'>"+oldStage+"</td>"
                        + "<td id='txtTnumber"+flag+"'>"+oldTnumber+"</td>"
                        + "<td id='txtName"+flag+"'>"+oldName+"</td>"
                        + "<td id='txtCategory"+flag+"'>"+oldCategory+"</td>"
                        + "<td id='txtPage"+flag+"'>"+oldPage+"</td>"
                        + "<td id='txtCount"+flag+"'>"+oldCount+"</td>"

                        + "<td id='txtDesignd"+flag+"'>"+oldDesignd+"</td>"
                        + "<td id='txtChecked"+flag+"'>"+oldChecked+"</td>"
                        + "<td><input type='text' id='txtDesigndratio"+flag+"' value='"+oldDesigndratio+"' size='1'/></td>"
                        + "<td><input type='text' id='txtCheckedratio"+flag+"' value='"+oldCheckedratio+"' size='1'/></td>"
                        + "<td id='txtData"+flag+"'>"+oldData+"</td>"
                        + "<td><input type='button' class='btn btn-default' name='update' value='保存' onclick='saveUpdateRow2(\""+rowId+"\",\""+flag+"\")'/></td>";  
            $("#"+rowId).html(uploadStr); 
            
         }    
  
         /*    
          *保存提交给审查的修改行    
          */
        function saveUpdateRow2(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).text();
            var newPnumber = $("#txtPnumber"+flag).text();    
            var newPname = $("#txtPname"+flag).text();    
            var newStage = $("#txtStage"+flag).text();
            var newTnumber = $("#txtTnumber"+flag).text();
            var newName = $("#txtName"+flag).text();
            var newCategory = $("#txtCategory"+flag).text();
            var newPage = $("#txtPage"+flag).text();
            var newCount = $("#txtCount"+flag).text();

            var newDesignd = $("#txtDesignd"+flag).text();
            var newChecked = $("#txtChecked"+flag).text();
            var newDesigndratio = $("#txtDesigndratio"+flag).val();
            var newCheckedratio = $("#txtCheckedratio"+flag).val();
            var newData = $("#txtData"+flag).text();
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newPnumber + "</td>"  
                        + "<td>" + newPname + "</td>"  
                        + "<td>" + newStage + "</td>"
                        + "<td>" + newTnumber + "</td>"
                        + "<td>" + newName + "</td>"
                        + "<td>" + newCategory + "</td>"
                        + "<td>" + newPage + "</td>"
                        + "<td>" + newCount + "</td>"

                        + "<td>" + newDesignd + "</td>"
                        + "<td>" + newChecked + "</td>"
                        + "<td>" + newDesigndratio + "</td>"
                        + "<td>" + newCheckedratio + "</td>"
                        + "<td>" + newData + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='退回' onclick='downsendSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow2(\""+rowId+"\")' /><input type='button' class='btn btn-default' name='update' value='提交' onclick='sendSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newName)//如果返回的有内容  
                {  
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/modifycatalog",
                    data: {Checked:newChecked,Designdratio:newDesigndratio,Checkedratio:newCheckedratio,CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("修改“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }
          /*    
          *修改-审查人员提交给统计-选中的行    
          */ 
         function updateSelectedRow3(rowId){
            var oldIndex = $("#"+rowId+" td:eq(0)").html();
            var oldPnumber = $("#"+rowId+" td:eq(1)").html();  
            var oldPname = $("#"+rowId+" td:eq(2)").html();  
            var oldStage = $("#"+rowId+" td:eq(3)").html();
            var oldTnumber = $("#"+rowId+" td:eq(4)").html();
            var oldName = $("#"+rowId+" td:eq(5)").html();
            var oldCategory = $("#"+rowId+" td:eq(6)").html();
            var oldPage = $("#"+rowId+" td:eq(7)").html();
            var oldCount = $("#"+rowId+" td:eq(8)").html();
            var oldComplex = $("#"+rowId+" td:eq(9)").html();
            var oldChecked = $("#"+rowId+" td:eq(10)").html();
            var oldExamined = $("#"+rowId+" td:eq(11)").html();  
            var oldCheckedratio = $("#"+rowId+" td:eq(12)").html();
            var oldExaminedratio = $("#"+rowId+" td:eq(13)").html();
            var oldData = $("#"+rowId+" td:eq(14)").html();
            // if(oldPrice != ""){//去掉第一个人民币符号  
            //     oldPrice = oldPrice.substring(1);  
            // }  
            var uploadStr = "<td id='txtIndex"+flag+"'>"+oldIndex+"</td>"
                        + "<td id='txtPnumber"+flag+"'>"+oldPnumber+"</td>"  
                        + "<td id='txtPname"+flag+"'>"+oldPname+"</td>"  
                        + "<td id='txtStage"+flag+"'>"+oldStage+"</td>"
                        + "<td id='txtTnumber"+flag+"'>"+oldTnumber+"</td>"
                        + "<td id='txtName"+flag+"'>"+oldName+"</td>"
                        + "<td id='txtCategory"+flag+"'>"+oldCategory+"</td>"
                        + "<td id='txtPage"+flag+"'>"+oldPage+"</td>"
                        + "<td id='txtCount"+flag+"'>"+oldCount+"</td>"
                        + "<td><input type='text' id='txtComplex"+flag+"' value='"+oldComplex+"' size='1'/></td>"
                        + "<td id='txtChecked"+flag+"'>"+oldChecked+"</td>"
                        + "<td id='txtExamined"+flag+"'>"+oldExamined+"</td>"  
                        + "<td><input type='text' id='txtCheckedratio"+flag+"' value='"+oldCheckedratio+"' size='1'/></td>"
                        + "<td><input type='text' id='txtExaminedratio"+flag+"' value='"+oldExaminedratio+"' size='1'/></td>"
                        + "<td><input type='text' id='txtData"+flag+"' class='datepicker' value='"+oldData+"'/></td>"
                        + "<td><input type='button' class='btn btn-default' name='update' value='保存' onclick='saveUpdateRow3(\""+rowId+"\",\""+flag+"\")'/></td>";  
            $("#"+rowId).html(uploadStr);
            $(".datepicker").datepicker({
            language: "zh-CN",
            autoclose: true,//选中之后自动隐藏日期选择框
            clearBtn: true,//清除按钮
            todayBtn: 'linked',//今日按钮
            format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
            });  
         }    
  
         /*    
          *审查人员保存修改行    
          */
        function saveUpdateRow3(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).text();
            var newPnumber = $("#txtPnumber"+flag).text();    
            var newPname = $("#txtPname"+flag).text();    
            var newStage = $("#txtStage"+flag).text();
            var newTnumber = $("#txtTnumber"+flag).text();
            var newName = $("#txtName"+flag).text();
            var newCategory = $("#txtCategory"+flag).text();
            var newPage = $("#txtPage"+flag).text();
            var newCount = $("#txtCount"+flag).text();
            var newComplex = $("#txtComplex"+flag).val();
            var newChecked = $("#txtChecked"+flag).text();
            var newExamined = $("#txtExamined"+flag).text();    
            var newCheckedratio = $("#txtCheckedratio"+flag).val();
            var newExaminedratio = $("#txtExaminedratio"+flag).val();
            var newData = $("#txtData"+flag).val();
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>"+newPnumber+"</td>"  
                        + "<td>"+newPname+"</td>"  
                        + "<td>"+newStage+"</td>"
                        + "<td>"+newTnumber+"</td>"
                        + "<td>"+newName+"</td>"
                        + "<td>"+newCategory+"</td>"
                        + "<td>"+newPage+"</td>"
                        + "<td>"+newCount+"</td>"
                        + "<td>" + newComplex + "</td>"
                        + "<td>" + newChecked + "</td>"
                        + "<td>" + newExamined + "</td>"
                        + "<td>" + newCheckedratio + "</td>"
                        + "<td>" + newExaminedratio + "</td>"
                        + "<td>" + newData + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='退回' onclick='downsendSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow3(\""+rowId+"\")' /><input type='button' class='btn btn-default' name='update' value='提交' onclick='sendSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newName)//如果返回的有内容  
                {  
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/modifycatalog",
                    data: {Complex:newComplex,Examined:newExamined,Checkedratio:newCheckedratio,Examinedratio:newExaminedratio,Data:newData,CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("修改“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }
</script>




<script type="text/javascript">
// $(function() {
         // $('#treeview').treeview('collapseAll', { silent: true });
          // $('#treeview').treeview({
          // data: [{{.json}}],//defaultData,
          // data:alternateData,
          // levels: 5,// expanded to 5 levels
          // enableLinks:true,
          // showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
//         });
// });
  $(document).ready(function() {
  $("table").tablesorter({sortList: [[13,0]]});
  // $("#ajax-append").click(function() {
  //    $.get("assets/ajax-content.html", function(html) {
  //     // append the "ajax'd" data to the table body
  //     $("table tbody").append(html);
  //     // let the plugin know that we made a update
  //     $("table").trigger("update");
  //     // set sorting column and direction, this will sort on the first and third column
  //     var sorting = [[2,1],[0,0]];
  //     // sort on the first column
  //     $("table").trigger("sorton",[sorting]);
  //   });
  //   return false;
  // });
});
</script>
</body>
</html>
