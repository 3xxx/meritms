<!-- 这个是显示科室树状，来自achievement -->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <!-- <title>MeritMS</title> -->
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
</head>

<body>
  <!-- <div class="col-xs-12"> -->
  <!-- 编辑科室价值 -->
    <h3>组织结构</h3>
      <div id="treeview" class="col-xs-3"></div>

    <div id="toolbar2" class="btn-group">
        <button type="button" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">保存修改</i>
        </button>
    </div>
<!--   -->
  <div class="col-xs-9" id="details" style="display:none">
  <h3 id="rowtitle"></h3>
    <table id="table"
          data-toggle="table"
          data-search="true"
          data-show-toggle="true"
          data-show-columns="true"
          data-toolbar="#toolbar2"
          data-query-params="queryParams"
          data-sort-name="ProjectName"
          data-sort-order="desc"
          data-page-size="5"
          data-page-list="[5, 25, 50, All]"
          data-unique-id="id"
          data-pagination="true"
          data-side-pagination="client"
          data-click-to-select="true"
          >
      <thead>        
        <tr>
          <!-- radiobox data-checkbox="true"-->
          <th data-width="10" data-checkbox="true" data-formatter="stateFormatter"></th>
          <th data-formatter="index1">#</th>
          <th data-field="text">价值分类名称</th>
          <!-- <th data-field="Mark">价值分值</th>
          <th data-field="List">价值选项</th>
          <th data-field="ListMark">选项分值</th> -->
          <!-- <th data-field="Iprole" data-title-tooltip="1-管理员;2-下载任意附件;3-下载pdf;4-查看成果">权限等级</th> -->
        </tr>
      </thead>
    </table>
  </div>

<script type="text/javascript">
  function index1(value,row,index){
      return index+1
  }

      function stateFormatter(value, row, index) {
        // if (index === 2) {
        //     return {
        //         disabled: true
        //     };
        // }
        if (row.Level === "1") {
            return {
                // disabled: true,
                checked: true
            }
        }
        return value;
      }

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
          // alert(JSON.stringify({{.json}}));
         // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
            showCheckbox: true,
                state: {
                  checked: true,
                  disabled: true,
                  expanded: true,
                  selected: true
                },
            data: {{.json}},//[{{.json}}]——有时候加个中括号就行了。defaultData,
            // data:alternateData,
            levels: 2,// expanded to 5 levels
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
          $("#rowtitle").html(data.text+"-价值列表");
          $("#details").show();
          dataid=data.id;//全局变量
          // alert(dataid);
         //得到选择的节点
        var arr = new Array();
        arr=$('#treeview').treeview('getChecked');
        // arr=$('#tree').treeview('getSelected',0);
        // alert(arr[0].nodeId);//节点顺序号0.0.0.1这样的
        // alert(arr[0].id);
         $('#table').bootstrapTable('refresh', {url:'/admin/merit/secoffice/'+dataid});
          // document.getElementById("iframepage").src="/achievement/secofficeshow?secid="+data.Id+"&level="+data.Level;
        });   
    });

    // 保存修改
    $("#editorButton").click(function() {
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      var selectRow=$('#table').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
      }
      // if(confirm("确定删除吗？一旦删除将无法恢复！")){
        var title=$.map(selectRow,function(row){
          return row.Title;
        })
        var ids="";
        for(var i=0;i<selectRow.length;i++){
          if(i==0){
            ids=selectRow[i].Id;
          }else{
            ids=ids+","+selectRow[i].Id;
          }  
        }
        $.ajax({
          type:"post",
          url:"/admin/merit/secoffice/addsecofficemerit",
          data: {sid:dataid,ids:ids},
          success:function(data,status){
            alert("添加“"+data+"”成功！(status:"+status+".)");
            $('#table').bootstrapTable('refresh', {url:'/admin/merit/secoffice/'+dataid});
          }
        });
      // }  
    })

</script>
</body>
</html>