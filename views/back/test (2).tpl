<!-- iframe里展示个人待处理的详细情况-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>

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
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>

<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>

<link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
<script type="text/javascript" src="/static/js/select2.js"></script>

<link rel="stylesheet" type="text/css" href="/static/font-awesome/css/font-awesome.min.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/font-awesome.min.css"/>
<script src="/static/js/tableExport.js"></script>

<script src="/static/js/jquery.form.js"></script>

<!-- <script src="/static/js/admin/main.js"></script> -->
<!-- <script src="/static/js/admin/gridview.js"></script> -->
<!-- <script src="/static/js/admin/validate.js"></script> -->
<!-- <link rel="stylesheet" type="text/css" href="/static/css/admin/layout.css"/> -->

</head>
<div class="col-lg-2">
  <div id="tree"></div>
</div>

<script type="text/javascript">
    $(function () {
        // function getTree() {
          // text: "Node 1",
          // icon: "glyphicon glyphicon-stop",
          // selectedIcon: "glyphicon glyphicon-stop",
          // color: "#000000",
          // backColor: "#FFFFFF",
          // href: "#node-1",
          // selectable: true,
          // state: {
          //   checked: true,
          //   disabled: true,
          //   expanded: true,
          //   selected: true
          // },
          // tags: ['available'],
            // Some logic to retrieve, or generate tree structure
            var data = 
            [
              {
                text: "系统设置",
                icon: "fa fa-tachometer icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                selectable: true,
                // state: {
                  // checked: true,
                  // disabled: true,
                  // expanded: true,
                  // selected: true
                // },
                tags: ['available'],
                nodes: 
                [
                  { 
                    icon: "fa fa-cog",
                    text: "目录设置",
                    href: "/admin",
                    id: '00001',
                    nodeId: '00001'
                  }, 
                  { 
                    icon: "fa fa-bug",
                    text: "爬虫设置",
                    id: '00002'
                  }, 
                  { 
                    icon: "fa fa-th-list",
                    text: "项目权限",
                    id: '00003'
                  }, 
                  { 
                    icon: "fa fa-user",
                    text: "账号管理",
                    id: '00004',
                    nodes: 
                    [
                      { icon: "fa fa-users",
                        text: '用户组',
                        id: '00005'
                      },
                      { icon: "fa fa-user",
                        text: 'IP权限',
                        id: '00006'
                      }
                    ]
                  }
                ]
              }
            ]
            // return data;

          $('#tree').treeview({
            data: data, 
            levels: 5,
            enableLinks: true,
            // multiSelect: true
          });  
        // }
        var obj = {};
        obj.text = "123";
        

        $("#btn").click(function (e) {
            var arr = $('#tree').treeview('getSelected');
            for (var key in arr) {
                c.innerHTML = c.innerHTML + "," + arr[key].id;
            }
        })
    }) 

    function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
          }
</script>
<div class="col-lg-10">
<div id="toolbar1" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="auth" id="auth" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="password" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
</div>
<table id="table"
        data-toggle="table"
        data-url="/running"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client">
    <thead>        
    <tr>
        <th data-width="10" data-checkbox="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">类型名称</th>
        <th data-field="ProjectName" data-sortable="true">代码</th>
      </tr>
    </thead>
</table>
<div class="gridview2"></div>

<script type="text/javascript">
  $(function(){
    $("#toolbar1").on('auth',function(e, gridview, params){
      var menu_id = $(this).data('data-menu');
      if(menu_id == null ){
        return alertMsg('请先勾选要编辑的菜单！');
      }
      params.data.menu_id = menu_id
    }) 
  })

// $(function () {
//     var $result = $('#eventsResult');
//     var selectRow=$('#table').bootstrapTable('getSelections');

//     $('#table').on('all.bs.table', function (e, name, args) {
//         console.log('Event:', name, ', data:', args);
//     })
//     .on('click-row.bs.table', function (e, row, $element) {
//       alert("选择！"+row.Id);
//       if (selectRow.length<1){
//         selectRow=$('#table').bootstrapTable('getSelections');
//         alert("请选择"+selectRow.length);
//         // return;
//         }
//         $result.text('Event: click-row.bs.table');
//     })
//     .on('dbl-click-row.bs.table', function (e, row, $element) {
//         $result.text('Event: dbl-click-row.bs.table');
//     })
//     .on('sort.bs.table', function (e, name, order) {
//         $result.text('Event: sort.bs.table');
//     })
//     .on('check.bs.table', function (e, row) {
//         $result.text('Event: check.bs.table');
//     })
//     .on('uncheck.bs.table', function (e, row) {
//         $result.text('Event: uncheck.bs.table');
//     })
//     .on('check-all.bs.table', function (e) {
//         $result.text('Event: check-all.bs.table');
//     })
//     .on('uncheck-all.bs.table', function (e) {
//         $result.text('Event: uncheck-all.bs.table');
//     })
//     .on('load-success.bs.table', function (e, data) {
//         $result.text('Event: load-success.bs.table');
//     })
//     .on('load-error.bs.table', function (e, status) {
//         $result.text('Event: load-error.bs.table');
//     })
//     .on('column-switch.bs.table', function (e, field, checked) {
//         $result.text('Event: column-switch.bs.table');
//     })
//     .on('page-change.bs.table', function (e, number, size) {
//         $result.text('Event: page-change.bs.table');
//     })
//     .on('search.bs.table', function (e, text) {
//         $result.text('Event: search.bs.table');
//     });
// });


// $(document).ready(function() {
  // $("#addButton").click(function() {
    
  // var selectRow=$('#table').bootstrapTable('getSelections');  
// if (selectRow.length<1){
  // selectRow=$('#table').bootstrapTable('getSelections');
  // alert("请选择"+selectRow.length);
  // return;
// }
        // $('#modalTable').modal({
        // show:true,
        // backdrop:'static'
        // });
    // })
  // })

  $(document).ready(function() {
    $("#auth").click(function() {
      var selectRow=$('#table').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选分类名称！");
        return;
      }
      // alert('请先勾选要编辑的菜单！');
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })


  $("#addButton").click(function() {
    // alert('请先勾选要编辑的菜单！');
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })
  })
</script>

<!-- 添加菜单类别 -->
<div class="container">
  <form data-method="post" data-action="/{$Request.module}/{$Request.controller}/{$Request.action}" data-submit="ajax" data-validate="true" class="form-horizontal">

    <div class="modal fade" id="modalTable" tabindex="-1" role="dialog" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加账号</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">姓名</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" name="username" maxlength="8"  placeholder="至多8个字符" required></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">手机号码</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control mobile" name="mobile" maxlength="11" placeholder="至多11个字符" required></div>
              </div>
              <!--  <div class="form-group">
              <label class="col-sm-3 control-label">数字</label>
              <div class="col-sm-7">
                <input type="number" class="form-control digits" name="num" maxlength="20" placeholder="至多20个字符" required></div>
            </div>
            -->
            <div class="form-group must">
              <label class="col-sm-3 control-label">登录密码</label>
              <div class="col-sm-7">
                <input type="password" class="form-control" name="password" id="password" maxlength="32" placeholder="至多32个字符" required></div>
            </div>

            <div class="form-group must">
              <label class="col-sm-3 control-label">确认密码</label>
              <div class="col-sm-7">
                <input type="password" class="form-control equalto" name="password2" maxlength="32" placeholder="至多32个字符" required data-rule-equalto="#password" data-msg-equalto="密码不一致"></div>
            </div>
            <div class="form-group must">
              <label class="col-sm-3 control-label">状态</label>
              <div class="col-sm-7">
                <select name="status" class="form-control" required>
                  <option value="1" >显示</option>
                  <option value="2" >隐藏</option>
                  <option value="0" >禁用</option>
                </select>
              </div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="submit" class="btn btn-primary">保存</button>
        </div>
      </div>
    </div>
  </div>
</form>
</div>

<!-- onClickRow  click-row.bs.table  row, $element 当用户点击某一行的时候触发，参数包括：
row：点击行的数据，
$element：tr 元素，
field：点击列的 field 名称 -->
<script type="text/javascript">
  // var toolbarUrl = '/myself'; //$menuTable, $btn_gridview;
  // $(function() {
  //   $menuTable = $('#table')
  //       .on('clickRow',function(e, row, $element) {
  //             if ($btn_gridview == null) {
  //               var $gridview2 = $('.gridview2');
  //               $.ajax({
  //                 url: toolbarUrl, //+ row.id,
  //                 type: 'get',
  //                       dataType: 'html',
  //                 success:function(html){
  //                   $html = $(html);
  //                   $html.appendTo($gridview2);
  //                   win.init($gridview2)

  //                   $btn_gridview = $('#btn_gridview').gridView();

  //                   $menuTable.gridView('resetView');
  //                   $('#current_menu').val(row.id);
  //                   $btn_gridview.data('data-menu',row.id);
  //                 },
  //               })
  //               return;
  //             } else if ($btn_gridview.data('data-menu') == row.id) {
  //               return;
  //             }

  //             $btn_gridview.data('data-menu', row.id);
  //             console.log($btn_gridview.data('bootstrap.table'))
  //             $btn_gridview.data('bootstrap.table').options.url = toolbarUrl //+ row.id;
  //             $btn_gridview.bootstrapTable('refresh');
  //           })
  //     });
$('#table').bootstrapTable({
    onClickRow: function (row, $element) {
        // alert( "选择了行Id为: " + row.Id );
        $('#btn_gridview').bootstrapTable('refresh', {url:'/myself'});
    }
});

 /*初始化table数据*/
        $(function(){
            // $("#table").bootstrapTable('destroy').bootstrapTable({
            //     columns:columns,
            //     data:json
            // });
            $("#table").on("click-row.bs.table",function(e,row,ele){
                $(".info").removeClass("info");
                $(ele).addClass("info");
            });
            $("#get").click(function(){
                alert("商品名称：" + getContent().TuanGouName);
            })
        });
 // $('#editable td').on('change', function(evt, newValue) {
//     $.post( "script.php", { value: newValue })
//     .done(function( data ) {
//         alert( "Data Loaded: " + data );
//     });
// }); 

function format_status(status,row,index) {
  if(status == 1){
    return '显示'
  }else if(status == 2){
    return  '隐藏'
  }else if(status == 0){
    return  '禁止'
  }
}

</script>

<toolbar id="btn_toolbar1" class="toolbar" data-module="/admin/menu">
<div class="btn-group">
        <button type="button" data-name="addButton" class="btn btn-default" data-event-type="view" data-event-value="" data-target="modal"><i class="fa fa-plus" aria-hidden="true"> </i>添加</button>
        <button type="button" data-name="editButton" class="btn btn-default" data-event-type="view" data-event-value="" data-target="modal"><i class="fa fa-edit" aria-hidden="true"> </i>编辑</button>
        <button type="button" data-name="deleteButton" class="btn btn-default" data-event-type="default" data-event-value="" data-target="default"><i class="fa fa-trash" aria-hidden="true"> </i>删除</button>
    </div>
</toolbar>

<table id="btn_gridview"
        data-toggle="table"
       data-url="/"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-toolbar="#btn_toolbar1"
       data-query-params="queryParams"
       data-sort-name="ProjectName"
       data-sort-order="desc"

 data-page-size="5"
       data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
         data-pagination="true"
          data-side-pagination="client">
    <thead>        
    <tr>
        <th data-width="10" data-checkbox="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">名称</th>
        <th data-field="ProjectName" data-sortable="true">代码</th>
        <th data-field="DesignStage" data-sortable="true">级别</th>
      </tr>
    </thead>
</table>


<br/>
<br/>
</div>



</body>
</html>