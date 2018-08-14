<!-- 角色——查看科室成果统计，科室价值统计的权限 -->
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
    <!-- 角色表 -->
    <div id="h-role-info" class="col-sm-6 col-md-6 col-lg-6">
      <h3>角色表</h3>
      <div id="toolbar1" class="btn-group">
        <button type="button" id="editorButton" class="btn btn btn-primary btn-sm"> <i class="fa fa-save">保存修改</i></button>
      </div>
      <table id="table"
        data-url="/admin/role/get/"
        data-toggle="table"
        data-striped="true"
        data-toolbar="#toolbar1"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-side-pagination="client"
        data-pagination="true"
        data-click-to-select="true"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-search="false"
        data-select-item-name="role"
        >
        <thead>
          <tr><!-- data-checkbox -->
                <th data-field="state" data-radio="true"></th>
                <th data-formatter="index1">#</th>
                <th data-field="Rolenumber">角色编码</th>
                <th data-field="name">角色名称</th>
                <th data-align="center" data-formatter="StatusFormatter">状态</th>
                <!-- <th data-field="domain_desc">所属域</th> -->
                <!-- <th data-align="center"
                    data-field="create_user">创建人</th> -->
                <th data-align="center" data-field="Createtime" data-formatter="localDateFormatter" data-visible="false">创建时间</th>
                <!-- <th data-align="center"
                    data-field="modify_user">修改人</th> -->
                <th data-align="center" data-field="Updated" data-formatter="localDateFormatter" data-visible="false">修改时间</th>
                <!-- <th data-field="state-handle"
                    data-align="center"
                    data-formatter="RoleObj.formatter">资源操作</th> -->
          </tr>
        </thead>
      </table>
    </div>
  <!-- 编辑角色——查看组织架构成果统计，价值统计情况的权限 -->
    <h3>组织结构</h3>
      <div id="treeview" class="col-xs-3"></div>

<script type="text/javascript">
  function StatusFormatter(value, row, index) {
    // alert(row.Status);
    if (row.role == "0") {
        return '正常';
    }else{
      return '失效';
    }
  }

  function index1(value,row,index){
    return index+1
  }

  function stateFormatter(value, row, index) {
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
      var dHeight = iframe.contentWindow.document.documentElement.scrollHeight; 
      var height = Math.max(bHeight, dHeight); iframe.height = height;
      // console.log(height);//这个显示老是在变化
    }catch (ex){
    } 
  } 
  window.setInterval("reinitIframe()", 200);

    // 保存权限和修改权限
    $("#editorButton").click(function() {
      var selectRow=$('#table').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选角色！");
        return false;
      }
      var roleids="";
      for(var i=0;i<selectRow.length;i++){
        if(i==0){
          roleids=selectRow[i].Id;
        }else{
          roleids=roleids+","+selectRow[i].Id;
        }  
      }

      arr=$('#treeview').treeview('getChecked');
      var treenodeids="";
      if (arr.length!=0){
        var treeids="";
        for(var i=0;i<arr.length;i++){
          if(i==0){
            treeids=arr[i].id;
            treenodeids=arr[i].nodeId;
          }else{
            treeids=treeids+","+arr[i].id;
            treenodeids=treenodeids+","+arr[i].nodeId;
          }  
        };
      }
      $.ajax({
        type:"post",
        url:"/admin/role/roleachieve",
        data: {roleids:roleids,treeids:treeids,treenodeids:treenodeids},//这里加上项目id
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          // $('#table').bootstrapTable('refresh', {url:'/admin/role/get/'});
        }
      });  
    })

    //显示角色的查看成果、价值权限 
    // 每次点击角色表table，都检查是否具备查询权限
    $(function(){
      //点击角色表触发查询角色具备的项目目录权限
      $("#table").on("check.bs.table",function(e,row,ele){
          //刷新树   
          $.ajax({  //JQuery的Ajax  
            type: 'GET',    
            dataType : "json",//返回数据类型  
            // async:false, //同步会出现警告：Synchronous XMLHttpRequest on the main thread is deprecated because of its detrimental effects to the end user's experience 
            url: "/admin/role/getroleachieve",//请求的action路径 

            data: {roleid:row.Id},
             //同步请求将锁住浏览器，用户其它操作必须等待请求完成才可以执行  
            error: function () {//请求失败处理函数    
              alert('请求失败');    
            },  
            success:function(data){ //请求成功后处理函数。取到Json对象data
              // var findCheckableNodess = function() {
              // return $('#tree').treeview('search', [ data, { ignoreCase: false, exactMatch: true } ]);//忽略大小写——这个只支持名称
              // };
              $('#treeview').treeview('uncheckAll', { silent: true });
              if (data!=null){
                for(var i=0;i<data.length;i++){
                  // alert(data[i]);
                  var findCheckableNodess = function() {
                    return $('#treeview').treeview('findNodes', [data[i], 'id']);
                  }; 
                  var checkableNodes = findCheckableNodess();
                  // $('#tree').treeview('checkNode', [ checkableNodes, { silent: true } ]);
                  $('#treeview').treeview('toggleNodeChecked', [checkableNodes, { silent: true}]);
                  // alert(checkableNodes[0].id);
                } 
              }
            }
          });
        })
    });

    //显示组织架构tree
    $(function() {
          // alert(JSON.stringify({{.json}}));
         // $('#treeview').treeview('collapseAll', { silent: true });
        $('#treeview').treeview({
            showCheckbox: true,
                state: {
                  checked: true,
                  // disabled: true,
                  expanded: true,
                  // selected: false
                },
            data: {{.json}},//[{{.json}}]——有时候加个中括号就行了。defaultData,
            // data:alternateData,
            levels: 2,// expanded to 5 levels
            // enableLinks:false,
            hierarchicalCheck:true,//有效！！
            // propagateCheckEvent:true,
            highlightSearchResults:false,//搜索结果不高亮
            // showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        });  
    });

</script>
</body>
</html>