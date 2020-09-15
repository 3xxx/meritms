<!-- iframe里，左侧点击侧栏中的价值，右侧显示用户这个价值下的价值内容列表，可以添加、删除和修改-->
<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>MeritMS</title>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script src="/static/js/moment-with-locales.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />
  <script type="text/javascript" src="/static/js/echarts.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css" />
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> -->
  <link rel="stylesheet" type="text/css" href="/static/css/select2.css" />
  <script type="text/javascript" src="/static/js/select2.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
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
  <style type="text/css">
  div#addmodal {
    /*.modal .fade .in*/
    z-index: 3;
  }
  div#detailmodal{
    z-index: 3;
  }
  </style>
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
          <input type="hidden" id="secid" name="secid" value="{{.Secid}}" />
          <input type="hidden" id="level" name="level" value="{{.Level}}" />
          <input type="hidden" id="key" name="key" value="modify" />
          <div class="form-group">
            <label for="taskNote">统计周期：</label>
            <input type="text" class="form-control" name="datefilter" id="datefilter" value="" placeholder="选择时间段(默认最近一个月)" />
          </div>
          <button id="button1" class="btn btn-default">提交</button>
          <label class="control-label"> tips:(StartDay < DateRange <=EndDay)</label> </div> </div> <br>
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
              <table id="table" data-unique-id="id" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-toolbar="#sendingtoolbar" data-query-params="queryParams" data-search="true">
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
              <table id="table1" data-toggle="table" data-url="/merit/send/2" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-toolbar="#sendedtoolbar" data-query-params="queryParams">
                <thead>
                  <tr>
                    <th data-formatter="index1" data-halign="center" data-align="center" data-valign="middle">序号</th>
                    <th data-field="meritcate" data-halign="center" data-align="center" data-valign="middle" data-halign="center" data-align="center" data-valign="middle">价值分类</th>
                    <th data-field="merittitle" data-halign="center" data-align="center" data-valign="middle">价值名称</th>
                    <th data-field="topictitle" data-halign="center" data-align="center" data-valign="middle">价值内容名称</th>
                    <th data-field="choose" data-halign="center" data-align="center" data-valign="middle">价值选项</th>
                    <th data-field="mark" data-halign="center" data-align="center" data-valign="middle">价值分值</th>
                    <th data-field="Examined" data-halign="center" data-align="center" data-valign="middle">审核</th>
                    <th data-field="updated" data-formatter="localDateFormatter" data-halign="center" data-align="center" data-valign="middle">日期</th>
                    <th data-field="action" data-formatter="actionFormatter1" data-events="actionEvents" data-halign="center" data-align="center" data-valign="middle">详细</th>
                  </tr>
                </thead>
              </table>
              <tr>
                {{if .IsMe}}
                <td colspan="4"><input type="button" class="btn btn-primary" value="处&nbsp;&nbsp;&nbsp;&nbsp;理" onclick="ModifyRow()" /></td>
                {{end}}
              </tr>
              <br />
              <br />
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
          <table id="table2" data-toggle="table" data-url="/merit/send/3" data-search="true" data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-toolbar="#completedtoolbar" data-query-params="queryParams">
            <thead>
              <tr>
                <th data-formatter="index1" data-halign="center" data-align="center" data-valign="middle">序号</th>
                <th data-field="meritcate" data-halign="center" data-align="center" data-valign="middle">价值分类</th>
                <th data-field="merittitle" data-halign="center" data-align="center" data-valign="middle">价值名称</th>
                <th data-field="topictitle" data-halign="center" data-align="center" data-valign="middle">价值内容名称</th>
                <th data-field="choose" data-halign="center" data-align="center" data-valign="middle">价值选项</th>
                <th data-field="mark" data-halign="center" data-align="center" data-valign="middle">价值分值</th>
                <th data-field="Examined" data-halign="center" data-align="center" data-valign="middle">审核</th>
                <th data-field="updated" data-formatter="localDateFormatter" data-halign="center" data-align="center" data-valign="middle">日期</th>
                <th data-field="action" data-formatter="actionFormatter1" data-events="actionEvents" data-halign="center" data-align="center" data-valign="middle">详细</th>
              </tr>
            </thead>
          </table>
        </div>
      </div>
    </div>
    <script type="text/javascript">
    $(function() {
      $('input[name="datefilter"]').daterangepicker({
        ranges: {
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

    function index1(value, row, index) {
      return index + 1
    }

    function localDateFormatter(value) {
      return moment(value, 'YYYY-MM-DD').format('L');
    }

    function queryParams(params) {
      var date = $("#datefilter").val();
      var secid = $("#secid").val();
      var level = $("#level").val();
      var meritid = {{.Merit.Id }};
      params.datefilter = date;
      params.secid = secid; //传secid给后台，点击用户名，显示对应成果
      params.level = level;
      params.meritid = meritid;
      return params;
    }

    $(document).ready(function() {
      $("#addButton").click(function() {
        $('#addmodal').modal({
          show: true,
          backdrop: 'static'
        });
      })
    })

    function import_xls_catalog() {
      var form1 = window.document.getElementById("form1"); //获取form1对象
      form1.submit();
      $.ajax({
        success: function(data, status) { //数据提交成功时返回数据
          window.location.reload();
        }
      });
      return true; //这个return必须放最后，前面的值才能传到后台    
    }

    function insertNewRow() {
      // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
      window.open('/secofficeshow?secid=' + {{.Secid }} + '&level=' + {{.Level }} + '&key=editor');
    }

    function ModifyRow() {
      // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
      window.open('/secofficeshow?secid=' + {{.Secid }} + '&level=' + {{.Level }} + '&key=modify');
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
    // 添加价值
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
        $('#edit2').froalaEditor('html.set', row.content);

        $("input#meritid").remove();
        $("input#meritstate").remove();
        var th1 = "<input id='meritid' type='hidden' value='" + row.id + "'/>"
        $(".modal-body").append(th1);
        var th2 = "<input id='meritstate' type='hidden' value='" + row.state + "'/>"
        $(".modal-body").append(th1);
        $(".modal-body").append(th2);

        $('#detailmodal').modal({
          show: true,
          backdrop: 'static'
        });
      },
      //提交
      'click .send': function(e, value, row, index) {
        if (confirm("确定提交该行吗？")) {
          var removeline = $(this).parents("tr")
          alert(row.id)
          //提交到后台进行修改数据库状态修改
          $.ajax({
            type: "post", //这里是否一定要用post？？？
            url: "/merit/sendmerit",
            data: { meritid: row.id, state: row.state },
            success: function(data, status) { //数据提交成功时返回数据
              removeline.remove();
              $('#table1').bootstrapTable('refresh', { url: '/merit/send/2' });
              alert("提交“" + data + "”成功！(status:" + status + ".)");
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
      'click .remove': function(e, value, row, index) {
        // alert('You click remove icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
        if (confirm("确定删除该行吗？")) {
          var removeline = $(this).parents("tr")
          //提交到后台进行删除数据库
          // alert("欢迎您：" + name) 
          $.ajax({
            type: "post", //这里是否一定要用post？？？
            url: "/merit/delete",
            data: { meritid: row.Id },
            success: function(data, status) { //数据提交成功时返回数据
              removeline.remove();
              alert("删除“" + data + "”成功！(status:" + status + ".)");
            }
          });
        }
      }
    };
    </script>
    <!-- 添加价值内容 -->
    <div class="form-horizontal">
      <div class="modal fade" id="addmodal">
        <div class="modal-dialog" style="width: 90%">
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
                    <input type="text" class="form-control" id="title">
                  </div>
                </div>
                <div class="form-group must">
                  {{if gt (.list|len) 1}}
                  <label class="col-sm-3 control-label">价值选项</label>
                  <th>
                    <div class="col-sm-4">
                      <select id="choose" class="form-control">
                        <option>价值选项：</option>
                      </select>
                    </div>
                  </th>
                  {{end}}
                </div>
                <div class="form-group must">
                  <label class="col-sm-3 control-label">Active</label>
                  <div class="checkbox">
                    <label>
                      <input type="checkbox" checked="checked" name="active"> Check me out
                    </label>
                  </div>
                </div>
              </div>
              <label>内容:</label>
              <!-- <div>
                <script id="container" type="text/plain" style="height:200px;width: 100%"></script>
              </div> -->
              <div id="editor" style="width: 100%;padding: 10px">
                <div id='edit' style="margin-top: 30px;">
                </div>
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
                    <input type="text" readonly="true" class="form-control" id="title1"></div>
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
              <label>内容:</label>
              <!-- <div id="content1">
                <script id="container1" type="text/plain" style="height:200px;width: 100%"></script>
              </div> -->
              <div id='editor2' style="width: 100%;padding: 10px">
                <div id='edit2' style="margin-top: 30px;">
                </div>
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
    <script>
    // $(function(){
    //   $('#edit').froalaEditor()
    // });
    // 显示详细-兼修改
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
      $('#edit2').froalaEditor({
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

    // var data={{.list}};
    // for ( var i = 0; i<data.length; i++) {  
    //   $("#list").append('<option>' + data[i].choose + '</option>');
    // }
    $(document).ready(function() {
      $.each({{.list }}, function(i, d) {
        $("#choose").append('<option value="' + d.id + '">' + d.title + '</option>');
      });
    });
    $(document).ready(function() {
      $.each({{.list }}, function(i, d) {
        $("#choose1").append('<option value="' + d.id + '">' + d.title + '</option>');
      });
    });

    //添加价值
    function save() {
      var meritid = {{.Merit.Id }};
      var title = $('#title').val();
      // alert(title);
      // var choose = $('#choose option:selected').text();
      var choose = $('#choose option:selected').val();
      var active = document.getElementsByName('active')[0].checked
      // alert(active);
      // return;
      // var html = ue.getContent();
      var html = $('div#edit').froalaEditor('html.get');
      if (title && choose) {
        $.ajax({
          type: "post",
          url: "/merit/addmerit",
          data: { meritid: choose, title: title, content: html,active:active },
          success: function(data, status) {
            alert("添加“" + data + "”成功！(status:" + status + ".)");
            $('#addmodal').modal('hide');
            $('#table').bootstrapTable('refresh', { url: '/merit/send/1' });
          },
        });
      }else if(title){
        $.ajax({
          type: "post",
          url: "/merit/addmerit",
          data: { meritid: meritid, title: title, content: html,active:active },
          success: function(data, status) {
            alert("添加“" + data + "”成功！(status:" + status + ".)");
            $('#addmodal').modal('hide');
            $('#table').bootstrapTable('refresh', { url: '/merit/send/1' });
          },
        });
      } else {
        alert("请填写标题和选择！");
        return;
      }
    };

    //修改价值
    function updatemerit() {
      var title = $('#title1').val();
      // alert(title)
      var choose = $('#choose1 option:selected').val();
      var active = document.getElementsByName('active1')[0].checked
      // alert(choose)
      var meritid = $('#meritid').val();
      var state = $('#meritstate').val();
      // alert(state)
      var html = $('div#edit2').froalaEditor('html.get');
      if (title && choose && state == 1) {
        $.ajax({
          type: "post",
          url: "/merit/updatemerit",
          data: { pk: meritid, name: "Content", value: html,active:active,choose:choose }, //父级id
          success: function(data, status) {
            alert("修改“" + data + "”成功！(status:" + status + ".)");
            $('#detailmodal').modal('hide');
            $('#table').bootstrapTable('refresh', { url: '/merit/send/1' });
          },
        });
      } else if (state != 1) {
        alert("提交后的价值不允许再修改！");
        return;
      } else {
        alert("请填写标题和选择！");
        return;
      }
    };

    $(function() {
      $('#table').bootstrapTable({
        idField: 'id',
        url: '/merit/send/1',
        //编辑单元格事件
        // onEditableSave: function (field, row, oldValue, $el) {
        //   $.ajax({
        //     type : "POST",
        //     url : '/merit/updatemerit',
        //     data : row,
        //     dataType : 'json',
        //     cache : false,
        //     success : function(data) {
        //       if ("success" == data) {
        //         alert("编辑成功");
        //         $('#dataGrid').bootstrapTable('refresh');
        //         } else {
        //           alert(data);
        //         }
        //       },
        //     error: function () {
        //       alert('编辑失败');
        //     },
        //     complete: function () {
                
        //     }
        //   });
        // },
        columns: [{
          title: '序号',
          formatter: function(value, row, index) {
            return index + 1
          },
          halign: "center",
          align: "center",
          valign: "middle"
        }, {
          field: 'meritcate',
          title: '价值分类',
          sortable: 'true',
          halign: "center",
          align: "center",
          valign: "middle",
          // editable: {
          //   type: 'text',
          //   pk: 1,
          //   url: '/merit/updatemerit',
          //   title: 'Enter Ttitle'
          // }
        }, {
          field: 'merittitle',
          title: '价值名称',
          sortable: 'true',
          halign: "center",
          align: "center",
          valign: "middle",
          // editable: {
          //   type: 'text',
          //   pk: 1,
          //   url: '/merit/updatemerit',
          //   title: 'Enter Ttitle'
          // }
        }, {
          field: 'topictitle',
          title: '价值内容名称',
          sortable: 'true',
          halign: "center",
          align: "center",
          valign: "middle",
          editable: {
            type: 'text',
            pk: 1,
            url: '/merit/updatemerit',
            title: 'Enter Ttitle'
          }
        }, {
          field: 'choose',
          title: '价值选项',
          sortable: 'true',
          halign: "center",
          align: "center",
          valign: "middle",
          // editable: {
          //   type: 'select',
          //   source: {{.select2 }},
          //   pk: 1,
          //   url: '/merit/updatemerit',
          //   title: 'Enter Choose',
          //   params: function(params) {
          //     //originally params contain pk, name and value
          //     //取出选择的值
          //     //与list对比，得到listid
          //     //要么使用select2方式，type:'select2'
          //     var choose = $('#choose option:selected').val();
          //     params.a = 1;
          //     return params;
          //   }
          // }
        }, {
          field: 'mark',
          title: '价值分值',
          halign: "center",
          align: "center",
          valign: "middle"
        }, {
          field: 'active',
          title: 'Active',
          halign: "center",
          align: "center",
          valign: "middle",
          editable: {
            type: 'select',
            source: [true,false],
            pk: 1,
            url: '/merit/updatemerit',
            title: 'Enter Active',
          }
        },{
          field: 'action',
          title: '操作',
          formatter: 'actionFormatter',
          events: 'actionEvents',
          halign: "center",
          align: "center",
          valign: "middle"
        }]
      });
    });
    </script>
    </body>

</html>