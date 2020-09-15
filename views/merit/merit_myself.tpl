<!-- 展示个人的价值列表：已通过，待提交 管理员查看：已通过……-->
<!-- 任何用户、管理员登录直接显示的页面 -->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>MeritMS</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <link rel="stylesheet" href="/static/froala/css/froala_editor.css">
  <link rel="stylesheet" href="/static/froala/css/froala_style.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/code_view.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/draggable.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/colors.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/emoticons.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image_manager.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/line_breaker.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/table.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/char_counter.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/video.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/fullscreen.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/file.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/quick_insert.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/help.css">
  <!-- <link rel="stylesheet" href="/static/froala/css/third_party/spell_checker.css"> -->
  <link rel="stylesheet" href="/static/froala/css/plugins/special_characters.css">
  <link rel="stylesheet" href="/static/froala/js/codemirror.min.css">
  <link rel="stylesheet" href="/static/froala/css/themes/red.css">

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
      <table id="table0" data-toggle="table" data-url="/merit/myself" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-toolbar="#toolbar1" data-query-params="queryParams" data-sort-name="ProjectName" data-sort-order="desc" data-page-size="15" data-page-list="[5, 25, 50, All]" data-unique-id="id" data-pagination="true" data-side-pagination="client" data-single-select="true" data-click-to-select="true">
        <thead>
          <tr>
            <!-- radiobox data-checkbox="true"-->
            <th data-width="10" data-radio="true"></th>
            <th data-formatter="index1" data-halign="center" data-align="center" data-valign="middle">序号</th>
            <th data-field="meritcate" data-sortable="true" data-halign="center" data-align="center" data-valign="middle">价值分类</th>
            <th data-field="merittitle" data-halign="center" data-align="center" data-valign="middle">价值名称</th>
            <th data-field="topictitle" data-halign="center" data-align="center" data-valign="middle">价值内容名称</th>
            <th data-field="choose" data-halign="center" data-align="center" data-valign="middle">价值选项</th>
            <th data-field="mark" data-halign="center" data-align="center" data-valign="middle">价值分值</th>
            <th data-field="Examined" data-halign="center" data-align="center" data-valign="middle">审核</th>
            <th data-field="updated" data-formatter="localDateFormatter" data-halign="center" data-align="center" data-valign="middle">日期</th>
            <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents" data-halign="center" data-align="center" data-valign="middle">详细</th>
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
      <table id="table" data-toggle="table" data-url="/merit/examined" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-toolbar="#examinedtoolbar" data-query-params="queryParams">
        <thead>
          <tr>
            <th data-formatter="index1" data-halign="center" data-align="center" data-valign="middle">序号</th>
            <th data-field="usernickname" data-sortable="true" data-halign="center" data-align="center" data-valign="middle">姓名</th>
            <th data-field="meritcate" data-sortable="true" data-halign="center" data-align="center" data-valign="middle">价值分类</th>
            <th data-field="merittitle" data-halign="center" data-align="center" data-valign="middle">价值名称</th>
            <th data-field="topictitle" data-halign="center" data-align="center" data-valign="middle">价值内容名称</th>
            <th data-field="choose" data-halign="center" data-align="center" data-valign="middle">价值选项</th>
            <th data-field="mark" data-halign="center" data-align="center" data-valign="middle">价值分值</th>
            <th data-field="Examined" data-halign="center" data-align="center" data-valign="middle">审核</th>
            <th data-field="updated" data-formatter="localDateFormatter" data-halign="center" data-align="center" data-valign="middle">日期</th>
            <th data-field="action" data-formatter="actionFormatter2" data-events="actionEvents" data-halign="center" data-align="center" data-valign="middle">详细</th>
          </tr>
        </thead>
      </table>
    </div>
  </div>
  <!-- 显示详细 -->
  <div class="form-horizontal">
    <div class="modal fade" id="detailedmodal">
      <div class="modal-dialog" style="width: 90%">
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
                <label class="col-sm-3 control-label">价值选项</label>
                <th>
                  <div class="col-sm-4">
                    <select id="choose1" class="form-control">
                      <option></option>
                    </select>
                  </div>
                </th>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">Active</label>
                <div class="checkbox">
                  <label>
                    <input type="checkbox" checked="checked" name="active1" id="active1"> Check me out
                  </label>
                </div>
              </div>
            </div>
            <label>内容:</label><!-- style="margin-top: 30px;" -->
            <div id='edit2' style="width: 100%;padding: 10px">
              <div id='edit' style="margin-top: 30px;">
                </div>
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
<script type="text/javascript" src="/static/froala/js/codemirror.min.js"></script>
<script type="text/javascript" src="/static/froala/js/xml.min.js"></script>
<script type="text/javascript" src="/static/froala/js/froala_editor.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/align.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/char_counter.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/code_beautifier.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/code_view.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/colors.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/draggable.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/emoticons.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/entities.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/file.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/font_size.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/font_family.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/fullscreen.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/image.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/image_manager.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/line_breaker.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/inline_style.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/link.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/lists.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/paragraph_format.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/paragraph_style.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/quick_insert.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/quote.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/table.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/save.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/url.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/video.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/help.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/print.min.js"></script>
<!-- <script type="text/javascript" src="/static/froala/js/third_party/spell_checker.min.js"></script> -->
<script type="text/javascript" src="/static/froala/js/plugins/special_characters.min.js"></script>
<script type="text/javascript" src="/static/froala/js/plugins/word_paste.min.js"></script>
<script src="/static/froala/js/languages/zh_cn.js"></script>
<script type="text/javascript">
function index1(value, row, index) {
  return index + 1
}

function localDateFormatter(value) {
  return moment(value, 'YYYY-MM-DD').format('L');
}

function actionFormatter(value, row, index) {
  return '<a class="like" href="javascript:void(0)" title="成果列表"><i class="glyphicon glyphicon-list-alt"></i></a>&nbsp;'
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

window.actionEvents = {
  //详细
  'click .like': function(e, value, row, index) {
    $("#title1").val(row.topictitle);
    // alert(row.Choose);
    $("#choose1 option:selected").text(row.choose);
    if (row.active==true){
      document.getElementById("active1").checked=true;
    }else{
      document.getElementById("active1").checked=false;
    }
    // ue1.setContent(row.Content);
    $('#edit').froalaEditor('html.set', row.content);
    $("input#meritid").remove();
    var th1 = "<input id='meritid' type='hidden' value='" + row.Id + "'/>"
    $(".modal-body").append(th1);
    $('#detailedmodal').modal({
      show: true,
      backdrop: 'static'
    });
  },
  // 'click .like': function(e, value, row, index) {
  //   $("#title1").val(row.topictitle);
  //   $("#choose1 option:selected").text(row.choose);
  //   $('#detailedmodal').modal({
  //     show: true,
  //     backdrop: 'static'
  //   });
  //提交
  'click .send': function(e, value, row, index) {
    if (confirm("确定提交该行吗？")) {
      var removeline = $(this).parents("tr")
      //提交到后台进行修改数据库状态修改
      $.ajax({
        type: "post", //这里是否一定要用post？？？
        url: "/merit/sendmerit",
        data: { meritid: row.Id, state: row.State },
        success: function(data, status) { //数据提交成功时返回数据
          removeline.remove();
          alert("提交“" + data + "”成功！(status:" + status + ".)");
        }
      });
    }
  },
  //回退
  'click .downsend': function(e, value, row, index) {
    if (confirm("确定退回该行吗？")) {
      var removeline = $(this).parents("tr")
      //提交到后台进行修改数据库状态修改
      $.ajax({
        type: "post", //这里是否一定要用post？？？
        url: "/merit/downsendmerit",
        data: { meritid: row.Id, state: row.State },
        success: function(data, status) { //数据提交成功时返回数据
          removeline.remove();
          alert("退回“" + data + "”成功！(status:" + status + ".)");
        }
      });
    }
  },
}

$(function() {
  //超大屏幕'fullscreen',
  var toolbarButtons = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', '|', 'color', 'emoticons', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', '-', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting', 'selectAll', 'html', 'help'];
  //大屏幕
  var toolbarButtonsMD = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', '|', 'fontFamily', 'fontSize', 'color', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', '-', 'quote', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', '|', 'specialCharacters', 'insertHR', 'undo', 'redo', 'clearFormatting', '|', 'html', 'help'];
  //小屏幕'fullscreen',
  var toolbarButtonsSM = ['bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'insertLink', 'insertImage', 'insertTable', 'undo', 'redo'];
  //手机
  var toolbarButtonsXS = ['bold', 'italic', 'fontFamily', 'fontSize', 'undo', 'redo'];
  // var pid = $('#pid').val();
  //编辑器初始化并赋值 
  $('#edit').froalaEditor({
    placeholderText: '请输入内容',
    charCounterCount: true, //默认
    // charCounterMax         : -1,//默认
    saveInterval: 0, //不自动保存，默认10000
    // theme                    : "red",
    height: "300px",
    toolbarBottom: false, //默认
    toolbarButtonsMD: toolbarButtons, //toolbarButtonsMD,
    toolbarButtonsSM: toolbarButtonsMD, //toolbarButtonsSM,
    toolbarButtonsXS: toolbarButtonsXS,
    toolbarInline: false, //true选中设置样式,默认false
    imageUploadMethod: 'POST',
    heightMin: 450,
    charCounterMax: 3000,
    // imageUploadURL: "uploadImgEditor",
    imageParams: { postId: "123" },
    params: {
      acl: '01',
      AWSAccessKeyId: '02',
      policy: '03',
      signature: '04',
    },
    autosave: true,
    autosaveInterval: 2500,
    saveURL: 'hander/FroalaHandler.ashx',
    saveParams: { postId: '1' },
    spellcheck: false,
    imageUploadURL: '/v1/wx/uploadmeritimg', //上传到本地服务器
    imageUploadParams: { meritid: '{{.Secid}}' },
    imageDeleteURL: 'lib/delete_image.php', //删除图片
    imagesLoadURL: 'lib/load_images.php', //管理图片
    videoUploadURL: '/v1/wx/uploadmeritimg',
    videoUploadParams: { meritid: '{{.Secid}}' },
    fileUploadURL: '/v1/wx/uploadmeritimg',
    fileUploadParams: { meritid: '{{.Secid}}' },
    enter: $.FroalaEditor.ENTER_BR,
    language: 'zh_cn',
    // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
  });
})
</script>
</body>

</html>