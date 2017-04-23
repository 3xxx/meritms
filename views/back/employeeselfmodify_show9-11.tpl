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

<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/select2-bootstrap.css"/>

<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> 
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
<script type="text/javascript" src="/static/js/select2.js"></script>  
<!-- <script type="text/javascript" src="/static/js/mindmup-editabletable.js"></script> -->

<script src="/static/js/moment-with-locales.min.js"></script>

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
    <input type="text" class="form-control" name="datefilter" id="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
  </div>
  <input type='text' placeholder='计量单位' id='txtPage' value='55' size='1'/>
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
  <!-- <button type="submit" class="btn btn-primary">提交</button> -->
  <button id="button" class="btn btn-default">Refresh from url</button>
</form>
<br></div>

<div class="form-group">
<label class="control-label" id="regis" for="LoginForm-UserName">
  统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}
</label>
</div>
<h3>需要提交给校核</h3>
<table id="table" data-query-params="queryParams"></table>

<!-- <div id="toolbar" class="btn-group">
<select class="form-control">
    <option value="">Export Basic</option>
    <option value="all">Export All</option>
    <option value="selected">Export Selected</option>
  </select>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div> -->

<!-- <table id="table" 
      data-toggle="table"
      data-toolbar="#toolbar"
       data-url="/addinline"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
      data-query-params="queryParams"
       data-striped="true"
       data-clickToSelect="true"
       data-show-export="true"

       data-icon-size="sm"
       >
    <thead>        
    <tr> -->
        <!-- <th data-field="state" data-checkbox="true" data-formatter="stateFormatter"></th> -->
       <!--  <th data-field="state" data-checkbox="true"></th>
        <th data-field="Id">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <th data-field="Page">成果计量单位</th>
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Drawnratio">绘制系数</th>
        <th data-field="Data" data-formatter="localDateFormatter">出版</th>
        <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">操作</th>
      </tr>
    </thead>
</table>  -->

<script>
function localDateFormatter(value) {
                return moment(value, 'YYYY-MM-DD').format('L');
            }
 // moment(yourDateVar).format('yyyy-MM-dd')
    // moment().format();
</script>

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
     
</div>

<script type="text/javascript">

  function actionFormatter(value, row, index) {
    return [
        '<a class="like" href="javascript:void(0)" title="Like">',
        '<i class="glyphicon glyphicon-heart"></i>',
        '</a>',
        '<a class="edit ml10" href="javascript:void(0)" title="Edit">',
        '<i class="glyphicon glyphicon-edit"></i>',
        '</a>',
        '<a class="remove ml10" href="javascript:void(0)" title="Remove">',
        '<i class="glyphicon glyphicon-remove"></i>',
        '</a>'
    ].join('');
}

window.actionEvents = {
    'click .like': function (e, value, row, index) {
        alert('You click like icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    },
    'click .edit': function (e, value, row, index) {
        alert('You click edit icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    },
    'click .remove': function (e, value, row, index) {
        alert('You click remove icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    }
};
//这个是指定哪几个不能选的
function stateFormatter(value, row, index) {
    if (index === 2) {
        return {
            disabled: true
        };
    }
    if (index === 0) {
        return {
            disabled: true,
            checked: true
        }
    }
    return value;
}
//这个是导出的
$(function () {
  var $table = $('#table');
  $('#toolbar').find('select').change(function () {
    $table.bootstrapTable('refreshOptions', {
      exportDataType: $(this).val()
    });
  });
});
//这个是编辑表-2方法
// $(function () {
//     $('#table').bootstrapTable({
//         idField: 'ProjectNumber',
//         // pagination: true,
//         // search: true,
//         url: '/addinline',
//         columns: [{
//             field: 'Id',
//             title: '编号'
//         },
//         {
//             field: 'ProjectNumber',
//             title: '项目编号'
//         }, {
//             field: 'ProjectName',
//             title: '项目名称'
//         }],
//         onPostBody: function () {
//             $('#table').editableTableWidget({editor: $('<textarea>')});
//         }
//     });
// });
//在线编辑
// $(function () {
//   $('#table').bootstrapTable({
//     idField: 'ProjectNumber',
//     url: '/addinline',
//     columns: [{
//       field: 'Id',
//             title: '编号'
//         },
//         {
//       field: 'ProjectNumber',
//       title: 'ProjectNumber',
//       editable: {
//         type: 'text'
//       }
//     }, {
//       field: 'ProjectName',
//       title: 'ProjectName',
//       editable: {
//         type: 'address',
//         // var value={{.Ratio}}
//         display: function(value) {
//           if(!value) {
//             $(this).empty();
//             return; 
//           }
//           var html = '<b>' + $('<div>').text(value.Category).html() + '</b>, ' + $('<div>').text(value.Category).html() + ' st., bld. ' + $('<div>').text(value.Category).html();
//           $(this).html(html); 
//         }
//       }
//     }, {
//       field: 'description',
//       title: 'Description'
//     }]
//   });
// });
//待选择的修改
$(function () {
    $('#table').bootstrapTable({
        idField: 'Id',
        url: '/addinline',
        striped: "true",
        columns: [
        {
            field: 'Number',
            title: 'Number',
            formatter:function(value,row,index){
            return index+1
    }
          },
          {
            field: 'Id',
            title: '编号',
          },
        {
            field: 'Name',
            title: '成果名称',
            editable: {
                type: 'select2',
                inputclass: 'input-large',
                select2: {
                    tags: {{.Select2}},//这个select必须是在secoffice控制器中发过来['bootstrap-table', 'multiple-select', 'bootstrap-show-password', 'blog', 'scutech-redmine']
                    tokenSeparators: [',', ' ']
                }
            }
        },  
        {
            field: 'ProjectName',
            title: 'ProjectName',
            editable: {
                type: 'text'
            }
        }]
    });


});
function queryParams(params) {
  // var newPage = $("#txtPage").val();
  var date=$("#datefilter").val();
  params.datefilter=date;//"2016-09-10 - 2016-09-15";
        // params.your_param1 = 1; // add param1
        // params.your_param2 = 2; // add param2
        // console.log(JSON.stringify(params));
        // {"limit":10,"offset":0,"order":"asc","your_param1":1,"your_param2":2}
        return params;
    }

    var $table = $('#table'),
        $button = $('#button');

    $(function () {
        $button.click(function () {
            $table.bootstrapTable('refreshOptions', {
                showColumns: true,
                search: true,
                showRefresh: true,
                url: '/addinline2'
            });
        });
    });    
// $(function () {
 // $('#button').click(function () {
      // var newPage = $("#txtPage").val();
            // var date=$("#datefilter").val();
            // params.datefilter=date;
            // alert( "Date Loaded: " + newPage);
            // $table.bootstrapTable('refresh', {url:'/addinline2'});
            // return params;
    // }); 
// });
// function queryParams() {
//         var params = {};
//         $('#toolbar').find('input[name]').each(function () {
//             params[$(this).attr('name')] = $(this).val();
//         });
//         return params;
//     }

// function queryParams(params) {
//             return {
//                 pageSize: params.pageSize,
//                 pageIndex: params.pageNumber,
//                 UserName: $("#txtName").val(),
//                 Birthday: $("#txtBirthday").val(),
//                 Gender: $("#Gender").val(),
//                 Address: $("#txtAddress").val(),
//                 name: params.sortName,
//                 order: params.sortOrder
//             };
//         }        
// 使用jQuery.post()方法传修改的数据到后台，这实际上是小菜一碟。

// $('#editable td').on('change', function(evt, newValue) {
//     $.post( "script.php", { value: newValue })
//     .done(function( data ) {
//         alert( "Data Loaded: " + data );
//     });
// });


</script>

</body>
</html>
