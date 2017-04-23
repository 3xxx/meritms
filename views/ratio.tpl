<!--编辑成果类型折标系数表——作废-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>成果类型折标系数表</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />
</head>

<div class="col-lg-12">
<div class="form-group">
        <label class="control-label" id="regis" for="LoginForm-UserName"></label>
</div>
  <table class="table table-striped" id="orderTable" name="orderTable">
    <thead>
      <tr>
        <th>#</th>
        <th>成果类型</th>
        <th>折算A2图纸系数</th>
        <th>是否实物工作量</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
      {{range $k,$v :=.Ratio}}
      <tr id="row{{.Id}}">
        <td>{{$k|indexaddone}}</td>
        <td>{{.Category}}</td>
        <td>{{.Rationum}}</td>
        <td>{{.Ismaterial}}</td>
        <td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow("row{{.Id}}")'/> 
        <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow("row{{.Id}}")' /></td> 
      </tr>
      {{end}}
    </tbody>
  </table>
  <!-- <input type="hidden" id="CategoryId" name="CategoryId" value="{{.CategoryId}}"/> -->
     <tr>    
       <td colspan="4"><input type="button" class="btn btn-default" name="insert" value="增加行" onclick="insertNewRow()"/></td>    
       </tr>
</div>



<script type="text/javascript">
//*********这个是编辑表格
var flag = 0;  //标志位，标志第几行  
        /*    
         *添加一个新行    
         */    
        function insertNewRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        // window.open('/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}});
        // window.location.href="http://www.jb51.net";    
            //获得表格有多少行    
            var rowLength = $("#orderTable tr").length;  
            //这里的rowId 就是row加上标志位组合的，为了方便起见所以分开好一点。  
            var rowId = "row" + flag;  
            //每次都往低flag+1的下标出添加tr，因为append是往标签内追加，所以用after
            //"<td>￥<input type='text' id='txtDrawn"+flag+"' value='' size='10'/></td>"  
            var insertStr = "<tr id="+rowId+">" 
                         + "<td><input type='text' placeholder='序号' id='txtIndex"+flag+"' value='' size='10'/></td>" 
                         + "<td><input type='text' placeholder='成果类型' id='txtCategory"+flag+"' value='' size='10'/></td>"   
                         + "<td><input type='text' placeholder='折算系数' id='txtRatio"+flag+"' value='' size='10'/></td>"
                         + "<td><select id='txtIsMaterial"+flag+"'><option value='volvo'>是否实物工作量</option><option value='true'>true</option><option value='false'>false</option></select></td>"
                         + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='确定' onclick='saveAddRow(\""+rowId+"\",\""+flag+"\")' /></td>"                   
                         + "</tr>";  
            $("#orderTable tr:eq("+(rowLength-1)+")").after(insertStr);  //这里之所以减2 ，是因为减去底部的一行和顶部一行，剩下的为开始插入的索引。  
            flag++;  
        }    
  
        /*    
         *删除选中的行    
         */    
         function deleteSelectedRow(rowId){    
            //根据rowId查询出该行所在的行索引    
            if(confirm("确定删除该行吗？")){    
                // $("#"+rowId).remove();    //这里需要注意删除一行之后 我的标志位没有-1，因为如果减一，那么我再增加一行的话，可能会导致我的tr的id重复，不好维护。
                // 提交到后台进行删除数据库
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/deleteratio",
                    data: {RatioId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("删除“"+data+"”成功！(status:"+status+".)");
                        $("#"+rowId).remove();
                        }
                    });  
            }       
         }    
          
         /*    
          *修改选中的行    
          */    
         function updateSelectedRow(rowId){
            var oldIndex = $("#"+rowId+" td:eq(0)").html();
            var oldCategory = $("#"+rowId+" td:eq(1)").html();  
            // var oldUnit = $("#"+rowId+" td:eq(2)").html();  
            var oldRatio = $("#"+rowId+" td:eq(2)").html();
            var oldMaterial = $("#"+rowId+" td:eq(3)").html();
            // if(oldPrice != ""){//去掉第一个人民币符号  
            //     oldPrice = oldPrice.substring(1);  
            // }  
            var uploadStr = "<td><input type='text' id='txtIndex"+flag+"' value='"+oldIndex+"' size='10'/></td>"
                        + "<td><input type='text' id='txtCategory"+flag+"' value='"+oldCategory+"' size='10'/></td>"  
                        // + "<td><input type='text' id='txtUnit"+flag+"' value='"+oldUnit+"' size='10'/></td>"  
                        + "<td><input type='text' id='txtRatio"+flag+"' value='"+oldRatio+"' size='10'/></td>"
                        + "<td><select id='txtIsMaterial"+flag+"'><option value='volvo'>是否实物工作量</option><option value='true'>true</option><option value='false'>false</option></select></td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='确定' onclick='saveUpdateRow(\""+rowId+"\",\""+flag+"\")' /></td>";  
            $("#"+rowId).html(uploadStr);  
         }    
  
         /*    
          *保存添加    
          */    
          function saveAddRow(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).val();
            var newCategory = $("#txtCategory"+flag).val();    
            // var newUnit = $("#txtUnit"+flag).val();    
            var newRatio = $("#txtRatio"+flag).val();
            var newMaterial =$("#txtIsMaterial"+flag+" option:selected").text();
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newCategory + "</td>"  
                        // + "<td>" + newUnit + "</td>"  
                        + "<td>" + newRatio + "</td>"
                        + "<td>" + newMaterial + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newCategory)//如果返回的有内容  
                {  
                 // var pid = $('#CategoryId').val();
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/addratio",
                    data: {Category:newCategory,Ratio:newRatio,IsMaterial:newMaterial},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("添加“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }

          /*    
          *保存修改    
          */
        function saveUpdateRow(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).val();
            var newCategory = $("#txtCategory"+flag).val();    
            // var newUnit = $("#txtUnit"+flag).val();    
            var newRatio = $("#txtRatio"+flag).val();
            var newMaterial =$("#txtIsMaterial"+flag+" option:selected").text();
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newCategory + "</td>"  
                        // + "<td>" + newUnit + "</td>"  
                        + "<td>" + newRatio + "</td>"
                        + "<td>" + newMaterial + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow(\""+rowId+"\")'/></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newCategory)//如果返回的有内容  
                {  
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/modifyratio",
                    data: {Category:newCategory,Ratio:newRatio,IsMaterial:newMaterial,RatioId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("修改“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }

</script>

<script type="text/javascript">
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
