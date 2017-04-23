<!-- iframe里展示个人待处理的详细情况-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>待处理成果</title>
  <!-- <base target=_blank>
  -->
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
  <!-- <script type="text/javascript" src="/static/js/bootstrap-datetimepicker.min.js"></script>
-->
<!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-datetimepicker.min.css"/>
-->
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
<!-- <link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
-->
<!-- <link rel="stylesheet" type="text/css" href="/static/css/select2-bootstrap.css"/>
-->
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>

<!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-combined.min.css"/> -->
<link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
<script type="text/javascript" src="/static/js/select2.js"></script>
<!-- <script type="text/javascript" src="/static/js/select2.js"></script>
-->
<!-- <script type="text/javascript" src="/static/js/mindmup-editabletable.js"></script>
-->
<!-- <script src="/static/js/moment-with-locales.min.js"></script>
-->
<!-- <script src="/static/js/bootstrap-table-filter-control.js"></script>
-->
<style>
i#delete
{
color:#C71585;
}
</style>
</head>

<div class="container">
        <h1>Multiple Table</h1>
        <p></p>
        <div class="row">
            <div class="col-md-3">
                <table data-toggle="table"
                       data-url="/myself">
                    <thead>
                    <tr>
                        <th data-field="ProjectNumber"
                            data-formatter="operateFormatter"
                            data-events="operateEvents">Bookmark 1</th>
                    </tr>
                    </thead>
                </table>
            </div>
            <div class="col-md-3">
                <table data-toggle="table"
                       data-url="../json/data3.json">
                    <thead>
                    <tr>
                        <th data-field="github.name"
                            data-formatter="operateFormatter"
                            data-events="operateEvents">Bookmark 2</th>
                    </tr>
                    </thead>
                </table>
            </div>
            <div class="col-md-3">
                <table data-toggle="table"
                       data-url="../json/data3.json">
                    <thead>
                    <tr>
                        <th data-field="github.name"
                            data-formatter="operateFormatter"
                            data-events="operateEvents">Bookmark 3</th>
                    </tr>
                    </thead>
                </table>
            </div>
            <div class="col-md-3">
                <table data-toggle="table"
                       data-url="../json/data3.json">
                    <thead>
                    <tr>
                        <th data-field="github.name"
                            data-formatter="operateFormatter"
                            data-events="operateEvents">Bookmark 4</th>
                    </tr>
                    </thead>
                </table>
            </div>
        </div>
    </div>
<script>
    window.operateEvents = {
        'click .like': function (e, value, row) {
            alert('You click like action, row: ' + JSON.stringify(row));
        },
        'click .remove': function (e, value, row) {
            alert('You click remove action, row: ' + JSON.stringify(row));
        }
    };

    function operateFormatter(value, row, index) {
        return [
            '<div class="pull-left">',

            // '<a href="https://github.com/wenzhixin/' + value + '" target="_blank">' + value + '</a>',
            '<button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#modalTable">'
            + value + 
        '</button>',
            '</div>',
            '<div class="pull-right">',
            '<a class="like" href="javascript:void(0)" title="Like">',
            '<i class="glyphicon glyphicon-heart"></i>',
            '</a>  ',
            '<a class="remove" href="javascript:void(0)" title="Remove">',
            '<i class="glyphicon glyphicon-remove"></i>',
            '</a>',
            '</div>'
        ].join('');
    }
</script>

<script>
    function detailFormatter(index, row) {
        var html = [];
        $.each(row, function (key, value) {
            html.push('<p><b>' + key + ':</b> ' + value + '</p>');
        });
        return html.join('');
    }
</script>
<div class="container">
        <h1>Modal Table</h1>
        <!-- Button trigger modal -->
        <button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#modalTable">
            Launch modal table
        </button>
        <div class="modal fade" id="modalTable" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">Modal table</h4>
                    </div>
                    <div class="modal-body">
                        <table id="table"
                               data-toggle="table"
                               data-height="299"
                               data-url="../json/data1.json">
                            <thead>
                            <tr>
                                <th data-field="id">ID</th>
                                <th data-field="name">Item Name</th>
                                <th data-field="price">Item Price</th>
                            </tr>
                            </thead>
                        </table>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->
    </div>
<script>
    var $table = $('#table');

    $(function () {
        $('#modalTable').on('shown.bs.modal', function () {
            $table.bootstrapTable('resetView');
        });
    });
</script>

<div class="col-lg-12">
    <!-- <div class="form-group"> -->
    <!-- <label class="control-label" id="regis" for="LoginForm-UserName">   {{.UserNickname}}</label> -->
    <!-- 显示部门名称 -->
    <!-- </div> -->
    <h2>{{.UserNickname}}</h2>
<!-- <div> -->
<!-- <form class="form-inline" method="get" action="/secofficeshow" enctype="multipart/form-data">
-->

    <div class="form-inline">
        <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
        <input type="hidden" id="level" name="level" value="{{.Level}}"/>
        <input type="hidden" id="key" name="key" value="modify"/>
        <div class="form-group">
            <label for="taskNote">统计周期：</label>
            <input type="text" class="form-control" name="datefilter" id="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
        </div>
        <script type="text/javascript">
            $(function() {
              $('input[name="datefilter"]').daterangepicker({
                  autoUpdateInput: false,
                  locale: {
                      cancelLabel: 'Clear'
                  }
              });
              $('input[name="datefilter"]').on('apply.daterangepicker', function(ev,        picker)    {
                  $(this).val(picker.startDate.format('YYYY-MM-DD') + ' - ' + picker.      endDate.   format('YYYY-MM-DD'));
              });
              $('input[name="datefilter"]').on('cancel.daterangepicker', function(ev,        picker)    {
                  $(this).val('');
              });
            });
        </script>
    <!-- <button type="submit" class="btn btn-primary">提交</button>
-->
        <button id="button" class="btn btn-default">提交</button>
        <label class="control-label">tips:(StartDay <= DateRange < EndDay)</label>
    </div>
<!-- </form>
-->
<br>

<!-- 添加 ：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime    "2006-01-02"}}-->
<div class="form-inline">
      <input type='text' placeholder='项目编号' class="form-control" id='Pnumber' value='' size='4'/>
      <input type='text' placeholder='项目名称' class="form-control" id='Pname' value='' size='20'/>
      <!-- <input type='text' placeholder='阶段' class="form-control" id='txtStage' value='' size='5'/> -->
      <select class="form-control" id='Stage'>
        <option>阶段：</option>
        <option>规划</option>
        <option>项目建议书</option>
        <option>可行性研究</option>
        <option>初步设计</option>
        <option>招标设计</option>
        <option>施工图</option>
      </select>
      <input type='text' placeholder='成果编号' class="form-control" id='Tnumber' value='' size='10'/>
      <input type='text' placeholder='成果名称' class="form-control" id='Name' value='' size='25'/>
      <select class="form-control" id='Category'>
        <option>成果类型：</option>
      </select>
      <!-- <input type='text' placeholder='单位' class="form-control" id='txtPage' value='' size='1'/> -->
      <input type='text' placeholder='数量' class="form-control" id='Count' value='' size='2'/>
</div>
  <br/>
<div class="form-inline">     
      <input type='text' placeholder='绘制/编制' class="form-control" id="uname1" value='' list="cars1" size='7'/>
      <input type='text' placeholder='设计' class="form-control" id="uname2" value='' list="cars2" size='7'/>
      <input type='text' placeholder='校核' class="form-control" id="uname3" value='' list="cars3" size='7'/>
      <input type='text' placeholder='审查' class="form-control" id="uname4" value='' list="cars4" size='7'/>
      <input type='text' placeholder='绘制系数' class="form-control" id='Drawnratio' value='' size='4'/>
      <input type='text' placeholder='设计系数' class="form-control" id='Designdratio' value='' size='4'/>
      <input type='text' placeholder='出版日期' class='datepicker' id='Date' value='' size='7'/>
      <input type='button' class='btn btn-primary' name='update' value='添加' onclick='saveAddRow()'/>
      <div id='datalistDiv'>
        <datalist id="cars1" name="cars1">
        </datalist>
      </div>
      <div id='datalistDiv'>
        <datalist id="cars2" name="cars2">
        </datalist>
      </div>
      <div id='datalistDiv'>
        <datalist id="cars3" name="cars3">
        </datalist>
      </div>
      <div id='datalistDiv'>
        <datalist id="cars4" name="cars4">
        </datalist>
      </div>
</div>
<br/>
    <form id="form1" class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
            <div class="form-group">
              <label>导入成果登记数据(Excel)
              <input type="file" class="form-control" name="catalog" id="catalog"></label>
              <br/>
              </div>
            <button type="submit" class="btn btn-primary" onclick="return import_xls_catalog();">提交</button>
    </form>
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

function saveAddRow(){
            var newPnumber = $("#Pnumber").val();    
            var newPname = $("#Pname").val();    
            var newStage = $("#Stage option:selected").text();
            var newTnumber = $("#Tnumber").val();
            var newName = $("#Name").val();
            var newCategory = $("#Category option:selected").text();
            
            var newCount = $("#Count").val();
            var newDrawn = $("#uname1").val();
            var newDesignd = $("#uname2").val();
            var newChecked = $("#uname3").val();
            var newExamined = $("#uname4").val();
            var newDrawnratio = $("#Drawnratio").val();
            var newDesigndratio = $("#Designdratio").val();
            var newDate = $("#Date").val();
          if(confirm("确定提交该行吗？")){    
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/addcatalog",
                    data: {Pnumber:newPnumber,Pname:newPname,Stage:newStage,Tnumber:newTnumber,Name:newName,Category:newCategory,Count:newCount,Drawn:newDrawn,Designd:newDesignd,Checked:newChecked,Examined:newExamined,Drawnratio:newDrawnratio,Designdratio:newDesigndratio,Date:newDate},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("添加“"+data+"”(status:"+status+".)");
                        $('#table').bootstrapTable('refresh', {url:'/myself'});
                        }
                    });  
            }   
         }
</script>         
<!-- </div> -->
<script type="text/javascript">
    // $(".datepicker").datepicker({
      $("#Date").datepicker({
            language: "zh-CN",
            autoclose: true,//选中之后自动隐藏日期选择框
            clearBtn: true,//清除按钮
            todayBtn: 'linked',//今日按钮
            format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
        });
  </script>

<h3>我发起，待提交</h3>
<div id="toolbar" class="btn-group">
        <!-- <select class="form-control">
        <option value="">Export Basic</option>
        <option value="all">Export All</option>
        <option value="selected">Export Selected</option>
        </select>
        -->
        <button type="button" class="btn btn-default"> <i class="glyphicon    glyphicon-plus"></i>
        </button>
        <button type="button" class="btn btn-default"> <i class="glyphicon        glyphicon-heart"></i>
        </button>
        <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
        </button>
    </div>
<table id="table"
      data-query-params="queryParams"
      data-toolbar="#toolbar"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
  ></table>

<h3>别人发起，我设计</h3>
<div id="designd" class="btn-group">
        <button type="button" class="btn btn-default"> <i class="glyphicon    glyphicon-plus"></i>
        </button>
        <button type="button" class="btn btn-default"> <i class="glyphicon        glyphicon-heart"></i>
        </button>
        <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
        </button>
</div>
<table id="table1" 
      data-query-params="queryParams"
      data-toolbar="#designd"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
       >
</table>

<h3>别人发起，我校核</h3>
<div id="checked" class="btn-group">
        <button type="button" class="btn btn-default"> <i class="glyphicon    glyphicon-plus"></i>
        </button>
        <button type="button" class="btn btn-default"> <i class="glyphicon        glyphicon-heart"></i>
        </button>
        <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
        </button>
</div>
<table id="table2" 
      data-detail-view="true"
      data-detail-formatter="detailFormatter"
      data-query-params="queryParams"
      data-toolbar="#checked"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
       >
</table>
<h3>别人发起，我审查</h3>
<div id="examined" class="btn-group">
        <button type="button" class="btn btn-default"> <i class="glyphicon    glyphicon-plus"></i>
        </button>
        <button type="button" class="btn btn-default"> <i class="glyphicon        glyphicon-heart"></i>
        </button>
        <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
        </button>
</div>
<table id="table3" 
      data-query-params="queryParams"
      data-toolbar="#examined"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
       >
</table>
</div>

<script type="text/javascript">
function actionFormatter(value, row, index) {
    return [
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>',
        '<a class="downsend" href="javascript:void(0)" title="退回">',
        '<i class="glyphicon glyphicon-step-backward"></i>',
        '</a>',
        
        '<a class="remove" href="javascript:void(0)" title="删除">',
        '<i id="delete" class="glyphicon glyphicon-remove"></i>',
        '</a>'
    ].join('');
}
// '<a class="edit ml10" href="javascript:void(0)" title="退回">','<i class="glyphicon glyphicon-edit"></i>','</a>'
window.actionEvents = {
    'click .send': function (e, value, row, index) {
        // alert('You click send icon, row: ' + JSON.stringify(row.Id));
        // alert(e);无值
        // alert(value);无值
        // alert(row);
        // alert(index);0~
        // console.log(value, row, index);
        if(confirm("确定提交该行吗？")){
          var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/sendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("提交“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },
    'click .downsend': function (e, value, row, index) {
        // alert('You click send icon, row: ' + JSON.stringify(row.Id));
        // alert(e);无值
        // alert(value);无值
        // alert(row);
        // alert(index);0~
        // console.log(value, row, index);
        if(confirm("确定退回该行吗？")){
        var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/downsendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("退回“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },

    // 'click .edit': function (e, value, row, index) {
    //     alert('You click edit icon, row: ' + JSON.stringify(row));
    //     console.log(value, row, index);
    // },
    'click .remove': function (e, value, row, index) {
        // alert('You click remove icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
        if(confirm("确定删除该行吗？")){  
        var removeline=$(this).parents("tr")
        //提交到后台进行删除数据库
         // alert("欢迎您：" + name) 
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/delete",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("删除“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    }
};

//不提供删除功能的操作
function actionFormatter1(value, row, index) {
    return [
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>',
        '<a class="downsend" href="javascript:void(0)" title="退回">',
        '<i class="glyphicon glyphicon-step-backward"></i>',
        '</a>',
    ].join('');
}
//不提供删除功能的操作
window.actionEvents1 = {
    'click .send': function (e, value, row, index) {
        if(confirm("确定提交该行吗？")){
          var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/sendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("提交“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },
    'click .downsend': function (e, value, row, index) {
        if(confirm("确定退回该行吗？")){
        var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/downsendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("退回“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
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
// $(function () {
//   var $table = $('#table');
//   $('#toolbar').find('select').change(function () {
//     $table.bootstrapTable('refreshOptions', {
//       exportDataType: $(this).val()
//     });
//   });
// });
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
//待选择的修改*******不要删除
//我发起
$(function () {
    $('#table').bootstrapTable({
        idField: 'Id',
        url: '/myself',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter ProjectNumber' 
            }
          },{
            field: 'ProjectName',
            title: '项目名称',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter ProjectName'  
            }
          },{
            field: 'DesignStage',
            title: '阶段',
            editable: {
                type: 'select',
                source: ["规划", "项目建议书", "可行性研究", "初步设计", "招标设计", "施工图"],
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter DesignStage'  
            }
          },{
            field: 'Tnumber',
            title: '成果编号',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter number'  
            }
          },{
            field: 'Name',
            title: '成果名称',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Name'  
            }
          },{
            field: 'Category',
            title: '成果类型',
            editable: {
                type: 'select',
                source: {{.Select2}},//["$1", "$2", "$3"],
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Category' 
            }
          },{
            field: 'Count',
            title: '数量',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Count'  
            }
          },{
            field: 'Drawn',
            title: '制图/编制',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},//'/regist/getuname1',
        // source: [
        //       {id: 'gb', text: 'Great Britain'},
        //       {id: 'us', text: 'United States'},
        //       {id: 'ru', text: 'Russia'}
        //    ],

        //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'

                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                  // multiple: true
                },//'/regist/getuname1',//这里用get方法，所以要换一个
                
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Drawn'  
            }
          },{
            field: 'Designd',
            title: '设计',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designd'  
            }
          },{
            field: 'Checked',
            title: '校核',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checked'  
            }
          },{
            field: 'Examined',
            title: '审查',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examined'  
            }
          },{
            field: 'Drawnratio',
            title: '制图比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Drawnratio'  
            }
          },{
            field: 'Designdratio',
            title: '设计比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designdratio'  
            }
          },{
            field: 'Datestring',
            title: '出版(日/月/年)',
            // formatter:localDateFormatter,
            editable: {
                type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                   }
                }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter',
            events:'actionEvents',
        }
        ]
    });
});
//我设计
$(function () {
    $('#table1').bootstrapTable({
        idField: 'Id',
        url: '/designd',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
          },{
            field: 'ProjectName',
            title: '项目名称',
          },{
            field: 'DesignStage',
            title: '阶段',
          },{
            field: 'Tnumber',
            title: '成果编号',
          },{
            field: 'Name',
            title: '成果名称',
          },{
            field: 'Category',
            title: '成果类型',
          },{
            field: 'Count',
            title: '数量',
          },{
            field: 'Drawn',
            title: '制图/编制',
          },{
            field: 'Designd',
            title: '设计',
          },{
            field: 'Checked',
            title: '校核',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checked'  
            }
          },{
            field: 'Examined',
            title: '审查',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examined'  
            }
          },{
            field: 'Drawnratio',
            title: '制图比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Drawnratio'  
            }
          },{
            field: 'Designdratio',
            title: '设计比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designdratio'  
            }
          },{
            field: 'Complex',
            title: '难度系数',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Complex'  
            }
          },{
            field: 'Datestring',
            title: '出版',
            // formatter:localDateFormatter,
            editable: {
                type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                   }
                }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter1',
            events:'actionEvents1',
        }
        ]
    });
});

//我校核
$(function () {
    $('#table2').bootstrapTable({
        idField: 'Id',
        url: '/checked',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
          },{
            field: 'ProjectName',
            title: '项目名称',
          },{
            field: 'DesignStage',
            title: '阶段',
          },{
            field: 'Tnumber',
            title: '成果编号',
          },{
            field: 'Name',
            title: '成果名称',
          },{
            field: 'Category',
            title: '成果类型',
          },{
            field: 'Count',
            title: '数量',
          },{
            field: 'Designd',
            title: '设计',
          },{
            field: 'Checked',
            title: '校核',
          },{
            field: 'Examined',
            title: '审查',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examined'  
            }
          },{
            field: 'Designdratio',
            title: '设计比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designdratio'  
            }
          },{
            field: 'Checkedratio',
            title: '校核比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checkedratio'  
            }
          },{
            field: 'Complex',
            title: '难度系数',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Complex'  
            }
          },{
            field: 'Datestring',
            title: '出版',
            // formatter:localDateFormatter,
            editable: {
                type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                   }
                }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter1',
            events:'actionEvents1',
        }
        ]
    });
});

//我审查
$(function () {
    $('#table3').bootstrapTable({
        idField: 'Id',
        url: '/examined',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
          },{
            field: 'ProjectName',
            title: '项目名称',
          },{
            field: 'DesignStage',
            title: '阶段',
          },{
            field: 'Tnumber',
            title: '成果编号',
          },{
            field: 'Name',
            title: '成果名称',
          },{
            field: 'Category',
            title: '成果类型',
          },{
            field: 'Count',
            title: '数量',
          },{
            field: 'Checked',
            title: '校核',
          },{
            field: 'Examined',
            title: '审查',
          },{
            field: 'Checkedratio',
            title: '校核比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checkedratio'  
            }
          },{
            field: 'Examinedratio',
            title: '审查比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examinedratio'  
            }
          },{
            field: 'Complex',
            title: '难度系数',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Complex'  
            }
          },{
            field: 'Datestring',
            title: '出版',
            // formatter:localDateFormatter,
            editable: {
            type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                }
            }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter1',
            events:'actionEvents1',
        }
        ]
    });
});

// var date={{.Starttime}};
// function list(value, row, index) {
             // return '<i class="glyphicon ' + icon + '"></i> ' + value;
            // return "<select data-index='row'><option>成果类型：</option></select>";
        // }
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
               autoclose: true,//选中之后自动隐藏日期选择框
               clearBtn: true,//清除按钮
               todayBtn: 'linked',//今日按钮
               format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
            });
}

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

    // var $table = $('#table'),
    // $button = $('#button');
    $(function () {
        $('#button').click(function () {
            $('#table').bootstrapTable('refresh', {url:'/myself'});
            $('#table1').bootstrapTable('refresh', {url:'/designd'});
            $('#table2').bootstrapTable('refresh', {url:'/checked'});
            $('#table3').bootstrapTable('refresh', {url:'/examined'});
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

// <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account" list="cars"></div>
//         <div id='datalistDiv'>
//           <datalist id="cars" name="cars">//           </datalist>
//         </div>
$(document).ready(function(){
        $("#sel_Province").change(function(){
            $.ajax({
                url: '<%=basePath%>areaAjax/getCity.do',
                data: "procode="+$("#sel_Province").val(),
                type: 'get',
                dataType:'json',
                error: function(data)
                {
                    alert("加载json 文件出错！");
                },
                success: function(data)
                {
                    for (var one in data)
                    {
                        var name = data[one].name;
                        var code = data[one].code;
                        $("#sel_City").append("<option value="+code+">"+name+"</option>");
                    }
                },
            });
       });
});

$(document).ready(function(){
//   $(array).each(function(index){
//     alert(this);
// });
 
// $.each(array,function(index){
//     alert(this);
// });
$.each({{.Select2}},function(i,d){
  // alert(this);
  // alert(i);
  // alert(d);
   $("#Category").append('<option value="' + i + '">'+d+'</option>');
   });
});

$('#uname1').attr("autocomplete","off"); 
$(document).ready(function(){
  $("#uname1").keyup(function(event){
    // alert(event.keyCode);
    var uname1=document.getElementById("uname1");
  // if (uname.value.length==0)
   if (event.keyCode != 38 && event.keyCode != 40 && uname1.value.length==2){
    $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars1").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
                // $("#uname1").keydown(function(){
                //   $("option").remove();
                // }); 
    }
 });
});  
$('#uname2').attr("autocomplete","off"); 
$(document).ready(function(){
  $("#uname2").keyup(function(event){
    var uname2=document.getElementById("uname2");
    // alert(event.keyCode);
   if (event.keyCode != 38 && event.keyCode != 40 && uname2.value.length==2){
    $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars2").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
                // $("#uname2").keydown(function(){
                //   $("option").remove();
                // }); 
    }
 });
}); 
$('#uname3').attr("autocomplete","off"); 
$(document).ready(function(){
  $("#uname3").keyup(function(event){
    var uname3=document.getElementById("uname3");
    // alert(event.keyCode);
   if (event.keyCode != 38 && event.keyCode != 40 && uname3.value.length==2){
    $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars3").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
                // $("#uname3").keydown(function(){
                //   $("option").remove();
                // }); 
    }
 });
}); 
$('#uname4').attr("autocomplete","off"); 
$(document).ready(function(){
  $("#uname4").keyup(function(event){
    var uname4=document.getElementById("uname4");
    // alert(event.keyCode);
   if (event.keyCode != 38 && event.keyCode != 40 && uname4.value.length==2){
    $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars4").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
    //             $("#uname4").keydown(function(){
    //               $("option").remove();
    //             }); 
    }
 });
}); 
</script>

</body>
</html>