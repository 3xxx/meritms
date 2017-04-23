<!-- iframe里，左侧点击侧栏中的价值，右侧显示用户这个价值下的价值内容列表，可以添加、删除和修改-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>MeritMS</title>

  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script src="/static/js/moment-with-locales.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/>

  <script type="text/javascript" src="/static/js/echarts.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
  <script type="text/javascript" src="/static/js/select2.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  

  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>

    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.min.js"></script>
</head>

<!-- <div class="form-group"> -->
<div class="col-lg-12">
  <h2>{{.UserNickname}}: {{.Meritcate.Title}}—{{.Merit.Title}}</h2>
  <ul class="nav nav-tabs">
    <li class="active">
      <a href="#sending" data-toggle="tab">提交</a>
    </li>
    <li>
      <a href="#completed" data-toggle="tab">完成</a>
    </li>
    <!-- <li>
      <a href="#examined" data-toggle="tab">审核</a>
    </li> -->
  </ul>

  <div class="tab-content">
    <div class="tab-pane fade in active" id="sending">
      <br>
      <div>
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
                    ranges : {
                    'Today': [moment(), moment()],
                    'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
                    'Last 7 Days': [moment().subtract(6, 'days'), moment()],
                    'Last 30 Days': [moment().subtract(29, 'days'), moment()],
                    'This Month': [moment().startOf('month'), moment().endOf('month')],
                    'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
                  },
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
          <!-- <button type="submit" class="btn btn-primary">提交</button>
        -->
          <button id="button1" class="btn btn-default">提交</button>
          <label class="control-label"> tips:(StartDay < DateRange <= EndDay)</label>
        </div>
      </div>
      <br>

      <h3>待提交</h3>
      <div id="sendingtoolbar" class="btn-group">
          <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
      </div>
      <table id="table"
              data-unique-id="id"
              data-search="true"
              data-show-refresh="true"
              data-show-toggle="true"
              data-show-columns="true"
              data-toolbar="#sendingtoolbar"
              data-query-params="queryParams"
              data-search="true"
              >
              
          <!-- <thead>        
          <tr>
              <th data-formatter="index1">#</th>
              <th data-field="MeritCate">价值分类</th>
              <th data-field="Merit">价值名称</th>
              <th data-field="Title">价值内容名称</th>
              <th data-field="Choose">价值选项</th>
              <th data-field="Mark">价值分值</th>
              <th data-field="Examined">审核</th>
              <th data-field="Created" data-formatter="localDateFormatter">日期</th>
              <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">操作</th>
            </tr>
          </thead> -->
      </table>

      <h3>已提交</h3>
      <div id="sendedtoolbar" class="btn-group">
          <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
      </div>
      <table id="table1"
              data-toggle="table"
              data-url="/merit/send/2"
              data-search="true"
              data-show-refresh="true"
              data-show-toggle="true"
              data-show-columns="true"
              data-toolbar="#sendedtoolbar"
              data-query-params="queryParams"
              >
          <thead>        
          <tr>
              <th data-formatter="index1">#</th>
              <th data-field="MeritCate">价值分类</th>
              <th data-field="Merit">价值名称</th>
              <th data-field="Title">价值内容名称</th>
              <th data-field="Choose">价值选项</th>
              <th data-field="Mark">价值分值</th>
              <th data-field="Examined">审核</th>
              <th data-field="Created" data-formatter="localDateFormatter">日期</th>
              <th data-field="action" data-formatter="actionFormatter1" data-events="actionEvents">详细</th>
            </tr>
          </thead>
      </table>

<script>
function index1(value,row,index){
  return index+1
}
function localDateFormatter(value) {
  return moment(value, 'YYYY-MM-DD').format('L');
}
function queryParams(params) {
  var date=$("#datefilter").val();
  var secid=$("#secid").val();
  var level=$("#level").val();
  var mid={{.Merit.Id}};
        params.datefilter=date;
        params.secid=secid;//传secid给后台，点击用户名，显示对应成果
        params.level=level;
        params.mid=mid;
        return params;
}
// $(function () {
//         $('#button1').click(function () {
//             //已完成
//             $('#table').bootstrapTable('refresh', {url:'/completed'});
//             //已提交
//             $('#table1').bootstrapTable('refresh', {url:'/running'});
//             //待提交
//             $('#table2').bootstrapTable('refresh', {url:'/myself'});
//             //待我处理的设计
//             $('#table3').bootstrapTable('refresh', {url:'/designd'});
//         });
//     }); 
</script>

        <tr>
          {{if .IsMe}}   
          <td colspan="4"><input type="button" class="btn btn-primary" value="处&nbsp;&nbsp;&nbsp;&nbsp;理" onclick="ModifyRow()"/></td> 
          {{end}}    
        </tr>
        <br/>
        <br/>
</div>

    <div class="tab-pane fade" id="completed">
      <h3>已完成</h3>
      <div id="completedtoolbar" class="btn-group">
          <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
      </div>
      <table id="table2"
              data-toggle="table"
              data-url="/merit/send/3"
              data-search="true"
              data-show-refresh="true"
              data-show-toggle="true"
              data-show-columns="true"
              data-toolbar="#completedtoolbar"
              data-query-params="queryParams"
              >
          <thead>        
          <tr>
              <th data-formatter="index1">#</th>
              <th data-field="MeritCate">价值分类</th>
              <th data-field="Merit">价值名称</th>
              <th data-field="Title">价值内容名称</th>
              <th data-field="Choose">价值选项</th>
              <th data-field="Mark">价值分值</th>
              <th data-field="Examined">审核</th>
              <th data-field="Created" data-formatter="localDateFormatter">日期</th>
              <th data-field="action" data-formatter="actionFormatter1" data-events="actionEvents">详细</th>
            </tr>
          </thead>
      </table>
    </div>

  </div>
</div>

<script type="text/javascript">
    uParse('.content',{
      rootPath : '/static/ueditor/'
    })

    var ue = UE.getEditor('container', {
      autoHeightEnabled: true,
      autoFloatEnabled: true
    });

    var ue1 = UE.getEditor('container1', {
      autoHeightEnabled: true,
      autoFloatEnabled: true
    });

  $(document).ready(function() {
    $("#addButton").click(function() {
        $('#addmodal').modal({
        show:true,
        backdrop:'static'
        });
    })
  })

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

//详细，提交，删除
function actionFormatter(value, row, index) {
    return [
        '<a class="like" href="javascript:void(0)" title="详细">',
        '<i class="glyphicon glyphicon-list-alt"></i></a>&nbsp;',
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>&nbsp;',
        '<a class="remove" href="javascript:void(0)" title="删除">',
        '<i id="delete" class="glyphicon glyphicon-remove"></i>',
        '</a>'
    ].join('');
}

//详细
function actionFormatter1(value, row, index) {
    return '<a class="like" href="javascript:void(0)" title="详细"><i class="glyphicon glyphicon-list-alt"></i></a>&nbsp;'
}
//审核：详细，提交，回退
function actionFormatter2(value, row, index) {
    return [
        '<a class="like" href="javascript:void(0)" title="详细">',
        '<i class="glyphicon glyphicon-list-alt"></i></a>&nbsp;',
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>&nbsp;',
        '<a class="downsend" href="javascript:void(0)" title="退回">',
        '<i class="glyphicon glyphicon-step-backward"></i>',
        '</a>&nbsp;',
    ].join('');
}
// '<a class="edit ml10" href="javascript:void(0)" title="退回">','<i class="glyphicon glyphicon-edit"></i>','</a>'
window.actionEvents = {
  //详细
    'click .like': function (e, value, row, index) {
        $("#title1").val(row.Title);
        // alert(row.Choose);
        $("#choose1 option:selected").text(row.Choose);
        ue1.setContent(row.Content);
        $("input#mid").remove();
        var th1="<input id='mid' type='hidden' value='" +row.Id+"'/>"
        $(".modal-body").append(th1);

        $('#detailmodal').modal({
          show:true,
          backdrop:'static'
        });
    },
    //提交
    'click .send': function (e, value, row, index) {
        if(confirm("确定提交该行吗？")){
          var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/merit/sendmerit",
            data: {mid:row.Id,state:row.State},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                $('#table1').bootstrapTable('refresh', {url:'/merit/send/2'});
                alert("提交“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },
    //回退
    // 'click .downsend': function (e, value, row, index) {
    //     if(confirm("确定退回该行吗？")){
    //     var removeline=$(this).parents("tr")
    //       //提交到后台进行修改数据库状态修改
    //         $.ajax({
    //         type:"post",//这里是否一定要用post？？？
    //         url:"/merit/downsendmerit",
    //         data: {mid:row.Id,state:row.State},
    //             success:function(data,status){//数据提交成功时返回数据
    //             removeline.remove();
    //             alert("退回“"+data+"”成功！(status:"+status+".)");
    //             }
    //         });  
    //     }
    // },
    //删除
    'click .remove': function (e, value, row, index) {
        // alert('You click remove icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
        if(confirm("确定删除该行吗？")){  
        var removeline=$(this).parents("tr")
        //提交到后台进行删除数据库
         // alert("欢迎您：" + name) 
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/merit/delete",
            data: {mid:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("删除“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    }
};

</script>

  <!-- 添加价值内容 -->
  <div class="form-horizontal">
    <div class="modal fade" id="addmodal">
      <div class="modal-dialog" style="width: 80%">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加价值</h3>
            <h4 class="modal-title">价值类型：{{.Meritcate.Title}}——价值名称：{{.Merit.Title}}</h4>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="title"></div>
              </div>
              <div class="form-group must">
                {{if gt (.list|len) 1}}
                <label class="col-sm-3 control-label">价值选项</label>
                  <th>
                    <div class="col-sm-4">
                      <select id="choose" class="form-control">
                      <option>价值选项：</option></select>
                    </div>
                  </th>
                {{end}}
              </div>
            </div>
            <label>内容:</label>
              <div>
                <script id="container" type="text/plain" style="height:200px;width: 100%"></script>
              </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="save()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- 显示详细 -->
  <div class="form-horizontal">
    <div class="modal fade" id="detailmodal">
      <div class="modal-dialog" style="width: 80%">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">价值详细</h3>
            <h4 class="modal-title"></h4>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="text" readonly="true" class="form-control" id="title1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">价值选项</label>
                  <th>
                    <div class="col-sm-4">
                      <select id="choose1" class="form-control">
                      <option></option></select>
                    </div>
                  </th>
              </div>
            </div>
            <label>内容:</label>
              <div id="content1">
                <script id="container1" type="text/plain" style="height:200px;width: 100%"></script><!-- width:1024px; -->
              </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="updatemerit()">修改</button>
          </div>
        </div>
      </div>
    </div>
  </div>

<script type="text/javascript">

      // var data={{.list}};
      // for ( var i = 0; i<data.length; i++) {  
      //   $("#list").append('<option>' + data[i].choose + '</option>');
      // }
    $(document).ready(function(){
      $.each({{.list}},function(i,d){
        $("#choose").append('<option value="' + i + '">'+d.choose+'</option>');
      });
    });

    ue1.ready(function () {
        ue1.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
            var pid = $('#pid').val();
            // var html = ue.getContent();
            ue1.execCommand('serverparam', {
              "pid":pid 
            });
        });
    });
  //添加价值
  function save(){
    // var radio =$("input[type='radio']:checked").val();
    // var meritcateid = {{.Meritcate.Title}};//$('#mcid').val();
    var meritid = {{.Merit.Id}};//$('#mid').val();
    var title = $('#title').val();
    // var content = $('#subtext1').val();
    var choose = $('#choose option:selected').text();
    // var userid = $('#prodlabel2').val();
    var html = ue.getContent();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (title&&choose){  
      $.ajax({
        type:"post",
        url:"/merit/addmerit",
        data: {mid:meritid,title:title,choose:choose,content:html},//父级id
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          $('#addmodal').modal('hide');
          $('#table').bootstrapTable('refresh', {url:'/merit/send/1'});
        }, 
      });
    }else{
      alert("请填写标题和选择！");
      return;
    }
  }

  //修改价值
  function updatemerit(){
    // var radio =$("input[type='radio']:checked").val();
    // var meritcateid = {{.Meritcate.Title}};//$('#mcid').val();
    var meritid = $('#mid').val();
    // var title = $('#title1').val();
    // var content = $('#subtext1').val();
    // var choose = $('#choose1 option:selected').text();
    // var userid = $('#prodlabel2').val();
    var html = ue1.getContent();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (title&&choose){  
      $.ajax({
        type:"post",
        url:"/merit/updatemerit",
        data: {pk:meritid,name:"Content",value:html},//父级id
        success:function(data,status){
          alert("修改“"+data+"”成功！(status:"+status+".)");
          $('#detailmodal').modal('hide');
          $('#table').bootstrapTable('refresh', {url:'/merit/send/1'});
        }, 
      });
    }else{
      alert("请填写标题和选择！");
      return;
    }
  }

  $(function () {
    $('#table').bootstrapTable({
        idField: 'Id',
        url: '/merit/send/1',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
              return index+1
            }
          },{
            field: 'MeritCate',
            title: '价值分类',
            // sortable:'true',
            // editable: {
            //     type: 'text',
            //     pk: 1,
            //     url: '/achievement/modifycatalog',
            //     title: 'Enter ProjectNumber' 
            // }
          },{
            field: 'Merit',
            title: '价值名称',
            // editable: {
            //     type: 'text',
            //     pk: 1,
            //     url: '/achievement/modifycatalog',
            //     title: 'Enter ProjectName'  
            // }
          },{
            field: 'Title',
            title: '价值内容名称',
            sortable:'true',
            editable: {
                type: 'text',
                pk: 1,
                url: '/merit/updatemerit',
                title: 'Enter Ttitle'  
            }
          },{
            field: 'Choose',
            title: '价值选项',
            sortable:'true',
            editable: {
                type: 'select',
                source: {{.select2}},//["$1", "$2", "$3"],
                pk: 1,
                url: '/merit/updatemerit',
                title: 'Enter Choose' 
            }
          },{
            field: 'Mark',
            title: '价值分值',
            // editable: {
            //     type: 'text',
            //     pk: 1,
            //     url: '/achievement/modifycatalog',
            //     title: 'Enter ProjectName'  
            // }
          },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter',
            events:'actionEvents',
          }
        ]
    });
  });

</script>
</body>
</html>
