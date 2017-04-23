<!-- 用户登录后自己的资料列表-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>MeritMS</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <!-- <script src="/static/js/jquery.form.js"></script> -->
  <link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
  <script type="text/javascript" src="/static/js/select2.js"></script>
</head>
<body>
  <div class="navbar navba-default navbar-fixed-top">
    <div class="container-fill">{{template "navbar" .}}</div>
  </div>

  <div class="col-lg-12">
  <h3>用户表-{{.User}}</h3>
    <div id="toolbar1" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="importButton" id="importButton" class="btn btn-default"> <i class="fa fa-plus">导入</i>
        </button>
        <!-- <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button> -->
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
    </div>

    <table id="table0"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-striped="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="Username"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        data-show-export="true"
        >
    </table>

<script type="text/javascript">
  $(function () {
    $('#table0').bootstrapTable({
        idField: 'Id',
        url: '/usermyself',
        // striped: "true",
        columns: [
          {
            radio: 'true',
            width: '10'
          },
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
            }
          },{
            field: 'Username',
            title: '用户名',
            sortable:'true',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter ProjectNumber' 
            }
          },{
            field: 'Nickname',
            title: '昵称',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter ProjectName'  
            }
          },{
            field: 'Password',
            title: '密码',
            editable: {
              type: 'text',
                // type: 'select',
                // source: ["规划", "项目建议书", "可行性研究", "初步设计", "招标设计", "施工图"],
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Password'  
            }
          },{
            field: 'Email',
            title: '邮箱',
            // sortable:'true',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Email'  
            }
          },{
            field: 'Department',
            title: '部门',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Department'  
            }
          },{
            field: 'Secoffice',
            title: '科室',
            sortable:'true',
            editable: {
                type: 'text',
                // source: {{.Select2}},//["$1", "$2", "$3"],
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Category' 
            }
          },{
            field: 'Ip',
            title: 'IP',
            editable: {
                type: 'text',
                pk: 1,
                url: '/admin/user/updateuser',
                title: 'Enter Count'  
            }
          },{
            field: 'Status',
            title: '状态',
            editable: {
              type: 'select2',
            //   // source:{{.Userselect}},//'/regist/getuname1',
              source: [
                {id: '1', text: '显示',value:1},
                {id: '2', text: '隐藏',value:2},
                {id: '3', text: '禁止',value:3}
              ],
            //   //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'
            //   select2: {
            //     allowClear: true,
            //     width: '150px',
            //     placeholder: '请选择状态',
            //     // multiple: true
            //   },//'/regist/getuname1',//这里用get方法，所以要换一个
            //   pk: 1,
            //   url: '/admin/user/updateuser',
              title: 'Enter Status'  
            }
          },{
            field: 'Lastlogintime',
            title: '最后登录',
            formatter:localDateFormatter,
          },{
            field: 'Createtime',
            title: '建立',
            formatter:localDateFormatter,
          },{
            field: 'Role',
            title: '权限',
            // editable: {
            //   type: 'select2', 
            //   // source:{{.Userselect}},//'/regist/getuname1',
            //   source: [
            //     {id: '1', text: '1级',value:1},
            //     {id: '2', text: '2级',value:2},
            //     {id: '3', text: '3级',value:3}
            //   ],
            //   //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'
            //   select2: {
            //     allowClear: true,
            //     width: '150px',
            //     placeholder: '请选择权限',
            //     // multiple: true
            //   },//'/regist/getuname1',//这里用get方法，所以要换一个
            //   pk: 1,
            //   url: '/admin/user/updateuser',
            //   title: 'Enter Status'  
            // }
          }
        ]
    });
  });

  function index1(value,row,index){
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  // 改变点击行颜色
  $(function(){
     // $("#table").bootstrapTable('destroy').bootstrapTable({
     //     columns:columns,
     //     data:json
     // });
     $("#table0").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
     });
  });

</script>

</div>

</body>
</html>