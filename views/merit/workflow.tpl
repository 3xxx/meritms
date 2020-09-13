<!-- workflow试验 -->
<!DOCTYPE html>
<html>
{{template "tpl/T.header.tpl"}}
<meta charset="UTF-8">
<title>待处理成果</title>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
<script type="text/javascript" src="/static/js/daterangepicker.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />
<script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script>
<script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.js"></script>
<link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker3.css" />
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" />
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/select2.css" />
<script type="text/javascript" src="/static/js/select2.js"></script>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
<script src="/static/js/tableExport.js"></script>
<script src="/static/js/jquery.form.js"></script>
<style>
  i#delete
        {
          color:#C71585;
        }
    </style>
</head>
<div class="container-fill">{{template "tpl/T.navbar.tpl" .}}</div>

<body>
  <div class="col-lg-12">
    <h3>我发起，待提交</h3>
    <div id="toolbar" class="btn-group">
      <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
      </button>
      <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
      </button>
      <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
      </button>
    </div>
    <table id="table" data-query-params="queryParams" data-toolbar="#toolbar" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-clickToSelect="true" data-show-export="true" data-filter-control="true" data-page-size="15" data-page-list="[10,15, 50, 100, All]" data-pagination="true" data-side-pagination="client">
    </table>
    <h3>别人发起，我设计</h3>
    <table id="table1" data-query-params="queryParams" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-striped="true" data-clickToSelect="true" data-show-export="true" data-filter-control="true" data-page-size="15" data-page-list="[10,15, 50, 100, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client">
    </table>
  </div>
  <script type="text/javascript">
  function index1(value, row, index) {
    return index + 1
  }

  function actionFormatter(value, row, index) {
    return [
      '<a class="send" href="javascript:void(0)" title="提交">',
      '<i class="fa fa-step-forward"></i>',
      '</a>&nbsp;&nbsp;',
      '<a class="downsend" href="javascript:void(0)" title="退回">',
      '<i class="fa fa-step-backward"></i>',
      '</a>&nbsp;&nbsp;',
      '<a class="remove" href="javascript:void(0)" title="删除">',
      '<i id="delete" class="fa fa-remove"></i>',
      '</a>'
    ].join('');
  }

  //别人发起，我设计，不提供删除功能的操作
  function actionFormatter1(value, row, index) {
    return [
      '<a class="send" href="javascript:void(0)" title="提交">',
      '<i class="fa fa-step-forward"></i>',
      '</a>&nbsp;&nbsp;',
      '<a class="downsend" href="javascript:void(0)" title="退回">',
      '<i class="fa fa-step-backward"></i>',
      '</a>',
    ].join('');
  }

  //别人发起，我设计，不提供删除功能的操作
  window.actionEvents1 = {
    'click .send': function(e, value, row, index) {
      var selectRow3 = $('#table1').bootstrapTable('getSelections');
      if (selectRow3.length == 0) {
        var mycars = new Array()
        mycars[0] = row;
        var selectRow3 = mycars
      }
      if (confirm("确定提交吗？")) {
        var ids = $.map($('#table1').bootstrapTable('getSelections'), function(row) {
          return row.id;
        });
        if (ids.length == 0) {
          ids = $.map(mycars, function(row) {
            return row.id;
          });
        }
        // var removeline=$(this).parents("tr")
        //提交到后台进行修改数据库状态修改
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: "/achievement/sendcatalog",
          data: JSON.stringify(selectRow3), //JSON.stringify(row),
          success: function(data, status) { //数据提交成功时返回数据
            $('#table1').bootstrapTable('remove', {
              field: 'id',
              values: ids
            });
            // removeline.remove();
            alert("提交“" + data + "”成功！(status:" + status + ".)");
            // $('#table1').bootstrapTable('refresh', {url:'/admin/merit/meritlist/1'});
          }
        });
      }
    },
    'click .downsend': function(e, value, row, index) {
      var selectRow3 = $('#table1').bootstrapTable('getSelections');
      if (selectRow3.length == 0) {
        var mycars = new Array()
        mycars[0] = row;
        var selectRow3 = mycars
      }
      if (confirm("确定退回吗？")) {
        var ids = $.map($('#table1').bootstrapTable('getSelections'), function(row) {
          return row.id;
        });
        if (ids.length == 0) {
          ids = $.map(mycars, function(row) {
            return row.id;
          });
        }
        $.ajax({
          type: "post", //这里是否一定要用post？？？
          url: "/achievement/downsendcatalog",
          data: JSON.stringify(selectRow3), //JSON.stringify(row),
          success: function(data, status) { //数据提交成功时返回数据
            $('#table1').bootstrapTable('remove', {
              field: 'id',
              values: ids
            });
            alert("退回“" + data + "”成功！(status:" + status + ".)");
          }
        });
      }
    }
  };

  //待选择的修改*******不要删除
  //我发起
  $(function(value, sourceData) {
    $('#table').bootstrapTable({
      idField: 'ID',
      uniqueId: 'ID',
      url: '/v1/flow/workflowdata',
      // striped: "true",
      columns: [{
          checkbox: true,
          width: 10
        },
        {
          title: '#',
          align: "center",
          valign: "middle",
          formatter: function(value, row, index) {
            return index + 1
          }
        },
        {
          field: 'Name',
          title: '名称',
          align: "center",
          valign: "middle",
          sortable: true,
        }, {
          field: 'action',
          title: '操作',
          align: "center",
          valign: "middle",
          formatter: 'actionFormatter',
          events: 'actionEvents',
        }
      ],
    });
  });
  //我设计
  $(function() {
    $('#table1').bootstrapTable({
      idField: 'ID',
      uniqueId: 'ID',
      url: '/v1/flow/flowgetdoctypebyname',
      // striped: "true",
      columns: [{
        checkbox: true,
        width: 10
      }, {
        // field: 'Number',
        title: '#',
        align: "center",
        valign: "middle",
        formatter: function(value, row, index) {
          return index + 1
        }
      }, {
        field: 'Name',
        title: '项目编号',
        align: "center",
        valign: "middle",
        sortable: true,
      }, {
        field: 'action',
        title: '操作',
        align: "center",
        valign: "middle",
        formatter: 'actionFormatter1',
        events: 'actionEvents1',
      }]
    });
  });

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('L');
  }

  function nameFormatter(value) {
    return '<a href="https://github.com/wenzhixin/' + value + '">' + value + '</a>';
  }
  //这个是显示时间选择
  function datepicker(value) {
    $(".datepicker").datepicker({
      language: "zh-CN",
      autoclose: true, //选中之后自动隐藏日期选择框
      clearBtn: true, //清除按钮
      todayBtn: 'linked', //今日按钮
      format: "yyyy-mm-dd" //日期格式，详见 http:// atepicker.readthedocs.org/en/release/options.html#format
    });
  }

  function queryParams(params) {
    var date = $("#datefilter").val();
    params.datefilter = date; //"2016-09-10 - 2016-09-15";
    return params;
  }

  $(function() {
    $('#button').click(function() {
      $('#table').bootstrapTable('refresh', { url: '/achievement/send/1' });
      $('#table1').bootstrapTable('refresh', { url: '/achievement/send/2' });
      $('#table2').bootstrapTable('refresh', { url: '/achievement/send/3' });
      $('#table3').bootstrapTable('refresh', { url: '/achievement/send/4' });
    });
  });
  </script>
</body>

</html>