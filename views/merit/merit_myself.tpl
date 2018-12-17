<!-- 展示个人的价值列表：已通过，待提交 管理员查看：已通过……-->
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>MeritMS</title>

<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>

  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>

    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.min.js"></script>
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

<div class="col-lg-12">
  <h2>{{.UserNickname}}</h2>
  <ul class="nav nav-tabs">
    <li class="active"><a href="#completed" data-toggle="tab">已通过</a></li>
    <li><a href="#sending" data-toggle="tab">提交</a></li>
    <li>
      <a href="#examined" data-toggle="tab">审核</a>
    </li>
  </ul>

  <div class="tab-content">
    <div class="tab-pane fade in active" id="completed">
      <!-- <div class="tab-content"> -->
      <div id="toolbar1" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
      </div>

      <table id="table0"
        data-toggle="table"
        data-url="/merit/myself"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="15"
        data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        >
      <thead>        
      <tr>
        <!-- radiobox data-checkbox="true"-->
        <th data-width="10" data-radio="true"></th>
        <th data-field="MeritCate" data-sortable="true">价值分类</th>
        <th data-field="Merit">价值名称</th>
        <th data-field="Title">价值内容名称</th>
        <th data-field="Choose">价值选项</th>
        <th data-field="Mark">价值分值</th>
        <th data-field="Examined">审核</th>
        <th data-field="Created" data-formatter="localDateFormatter">日期</th>
        <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">详细</th>
      </tr>
      </thead>
    </table>
    
    </div>
    <div class="tab-pane fade" id="sending">
      <p>这里将显示待提交的价值记录。</p>
    </div>

    <div class="tab-pane fade" id="examined">
      <div id="examinedtoolbar" class="btn-group">
          <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
      </div>
      <table id="table"
              data-toggle="table"
              data-url="/merit/examined"
              data-search="true"
              data-show-refresh="true"
              data-show-toggle="true"
              data-show-columns="true"
              data-toolbar="#examinedtoolbar"
              data-query-params="queryParams"
              >
          <thead>        
          <tr>
              <th data-formatter="index1">#</th>
              <th data-field="UserNickName" data-sortable="true">姓名</th>
              <th data-field="MeritCate" data-sortable="true">价值分类</th>
              <th data-field="Merit">价值名称</th>
              <th data-field="Title">价值内容名称</th>
              <th data-field="Choose">价值选项</th>
              <th data-field="Mark">价值分值</th>
              <th data-field="Examined">审核</th>
              <th data-field="Created" data-formatter="localDateFormatter">日期</th>
              <th data-field="action" data-formatter="actionFormatter2" data-events="actionEvents">详细</th>
            </tr>
          </thead>
      </table>
    </div>

  </div>

  <!-- 显示详细 -->
  <div class="form-horizontal">
    <div class="modal fade" id="detailedmodal">
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
                  <input type="text" class="form-control" id="title1"></div>
              </div>
              <div class="form-group must">
                <!-- <label class="col-sm-3 control-label">价值选项</label> -->
                  <!-- <div class="col-sm-4">
                    <select class="form-control" id='choose'>
                      <option>价值选项：</option>
                    </select>
                  </div> -->
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
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</div>

<script type="text/javascript">
  function index1(value,row,index){
      return index+1
  }

  function localDateFormatter(value) {
  return moment(value, 'YYYY-MM-DD').format('L');
  }
  function actionFormatter(value, row, index) {
    return  '<a class="like" href="javascript:void(0)" title="成果列表"><i class="glyphicon glyphicon-list-alt"></i></a>&nbsp;'
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

    uParse('.content',{
      rootPath : '/static/ueditor/'
    })

    var ue1 = UE.getEditor('container1', {
      autoHeightEnabled: true,
      autoFloatEnabled: true
    });

  window.actionEvents = {
    'click .like': function (e, value, row, index) {
        $("#title1").val(row.Title);
        $("#choose1 option:selected").text(row.Choose);
        ue1.setContent(row.Content);
        $('#detailedmodal').modal({
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
                alert("提交“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },
    //回退
    'click .downsend': function (e, value, row, index) {
        if(confirm("确定退回该行吗？")){
        var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/merit/downsendmerit",
            data: {mid:row.Id,state:row.State},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("退回“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },
  }
</script>
</body>
</html>