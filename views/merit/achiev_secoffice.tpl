<!-- iframe里展示科室总体情况-->
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="UTF-8">
      <title>科室情况汇总</title>
      <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
      <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
      <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
      <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
      <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    
      <script type="text/javascript" src="/static/js/moment.min.js"></script>
      <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
      <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/>
    
      <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
      <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
      <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
    
      <script src="/static/js/moment-with-locales.min.js"></script>
      <script type="text/javascript" src="/static/js/echarts.min.js"></script>
    
      <script src="/static/js/bootstrap-table-filter-control.js"></script>
      <script src="/static/js/bootstrap-table-export.min.js"></script>
      <script src="/static/js/tableExport.js"></script>
      <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
    </head>

    <div class="col-lg-12">
      <h2>{{.Sectitle}}</h2>
      <ul class="nav nav-tabs">
        <li class="active"><a href="#employee" data-toggle="tab">月份排名</a></li>
        <li><a href="#year" data-toggle="tab">年度排名</a></li>
        <li><a href="#proj" data-toggle="tab">年度项目</a></li>
        <li><a href="#achievement" data-toggle="tab">月度成果</a></li>
      </ul>
      <div class="tab-content">
        <div class="tab-pane fade in active" id="employee">
          <br>
          <div>
            <div class="form-inline">
              <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
              <input type="hidden" id="level" name="level" value="{{.Level}}"/>
              <div class="form-group">
                <label for="taskNote">统计周期：</label>
                <input type="text" class="form-control" id="datefilter" name="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
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
              <button type="submit" class="btn btn-primary" id="button">提交</button>
              <label class="control-label" id="regis" for="LoginForm-UserName">统计时间段：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}}</label>
            </div> 
          </div>
          <br>
<!--           <div class="form-group">
          </div> -->

          <!-- <table class="table table-striped">
            <thead>
              <tr>
                <th>#</th>
                <th>姓名</th>
                <th>编制</th>
                <th>设计</th>
                <th>校核</th>
                <th>审查</th>
                <th>汇总</th>
                <th>详细</th>
              </tr>
            </thead>

            <tbody>
              {{range $k,$v :=.Employee}}
              <tr>
                <td>{{$k|indexaddone}}</td>
                <td>{{.Name}}</td>
                <td>{{.Drawn}}</td>
                <td>{{.Designd}}</td>
                <td>{{.Checked}}</td>
                <td>{{.Examined}}</td>
                <td>{{.Sigma}}</td>
                <td>
                 <a href="/secofficeshow?secid={{.Id}}&level=3"><i class="glyphicon glyphicon-open"></i>详细</a>
                </td>  
              </tr>
              {{end}}
            </tbody>
          </table>
          <table class="table table-striped">
            <thead>
              <tr>
                <th>#</th>
                <th>姓名</th>
                <th>编制</th>
                <th>设计</th>
                <th>校核</th>
                <th>审查</th>
                <th>汇总</th>
                <th>详细</th>
              </tr>
            </thead>
            <tbody>
              {{range $k,$v :=.Employeereal}}
              <tr>
                <td>{{$k|indexaddone}}</td>
                <td>{{.Name}}</td>
                <td>{{.Drawn}}</td>
                <td>{{.Designd}}</td>
                <td>{{.Checked}}</td>
                <td>{{.Examined}}</td>
                <td>{{.Sigma}}</td>
                <td>
                 <a href="/secofficeshow?secid={{.Id}}&level=3"><i class="glyphicon glyphicon-open"></i>详细</a>
                </td>  
              </tr>
              {{end}}
            </tbody>
          </table>
          </div> -->

          <h3>成果统计</h3>
            <!-- <div id="completed" class="btn-group">
              <button type="button" class="btn btn-default">
                  <i class="glyphicon glyphicon-plus"></i>
              </button>
              <button type="button" class="btn btn-default">
                  <i class="glyphicon glyphicon-heart"></i>
              </button>
              <button type="button" class="btn btn-default">
                  <i class="glyphicon glyphicon-trash"></i>
              </button>
            </div> data-toolbar="#completed"-->
            <table id="table"
                  data-toggle="table"
                  data-url="/achievement/secofficedata"
                  data-search="true"
                  data-show-refresh="true"
                  data-show-toggle="true"
                  data-show-columns="true"
                  data-show-export="true"
                  data-query-params="queryParams"
                  >
              <thead>        
              <tr>
                  <th data-formatter="index1">#</th>
                  <th data-field="Name" data-sortable="true">姓名</th>
                  <th data-field="Drawn">编制</th>
                  <th data-field="Designd">设计</th>
                  <th data-field="Checked">校核</th>
                  <th data-field="Examined">审查</th>
                  <th data-field="Sigma" data-sortable="true">全部工作量汇总</th>
                  <th data-field="MaterialSigma" data-sortable="true">实物工作量汇总</th>
                  <th data-field="action" data-formatter="actionFormatter1">详细</th>
                </tr>
              </thead>
            </table>

            <script type="text/javascript">
              function actionFormatter1(value, row, index) {
                return '<a href="/achievement/secofficeshow?secid='+row.Id+'&level=3"><i class="fa fa-list"></i>详细</a>'
              }
              function queryParams(params) {
                var date=$("#datefilter").val();
                var secid=$("#secid").val();
                var level=$("#level").val();
                // alert( "Data Loaded: " + date );
                params.datefilter=date;
                params.secid=secid;//传secid给后台，点击用户名，显示对应成果
                params.level=level;
                    return params;
              }
              $(function () {
                $('#button').click(function () {
                  $('#table').bootstrapTable('refresh', {url:'/secofficedata'});
                });
              });
            </script>
        </div>

        <div class="tab-pane fade" id="year">
          <p>这里将显示全年人员情况。</p>
        </div>
    
        <div class="tab-pane fade" id="proj">
          <p>这里将显示年度项目列表，每个项目的成果列表。</p>
          <h3>科室年度项目</h3>
            <div id="secparticipate" class="btn-group">
              <button type="button" class="btn btn-default">
                 <i class="glyphicon glyphicon-plus"></i>
              </button>
              <button type="button" class="btn btn-default">
                 <i class="glyphicon glyphicon-heart"></i>
              </button>
              <button type="button" class="btn btn-default">
                 <i class="glyphicon glyphicon-trash"></i>
              </button>
            </div>
            <table id="table10"
              data-toggle="table"
              data-url="/achievement/secparticipate"
              data-search="true"
              data-show-refresh="true"
              data-show-toggle="true"
              data-show-columns="true"
              data-toolbar="#secparticipate"
              data-query-params="queryParams"
              data-filter-control="true"
              data-filter-show-clear="true"
              data-show-export="true"
              >
            <thead>        
              <tr>
              <th data-formatter="index1">#</th>
              <th data-field="ProjectNumber" data-filter-control="select">项目编号</th>
              <th data-field="ProjectName" data-sortable="true" data-filter-control="select">项目名称</th>
              <th data-field="DesignStage" data-sortable="true" data-filter-control="select">项目阶段</th>
              <th data-field="Value" data-sortable="true">成果总分</th>
              <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">详细</th>
              </tr>
            </thead>
          </table>

      <div class="container">
        <div class="modal fade" id="modalTable" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="bottom:auto">
            <div class="modal-dialog" style="width:auto;height:800px;overflow:auto;">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">本专业项目全部成果列表</h4>
                    </div>
                    <div class="modal-body">
                      <div id="projachievement" class="btn-group">
                        <button type="button" class="btn btn-default">
                           <i class="glyphicon glyphicon-plus"></i>
                        </button>
                        <button type="button" class="btn btn-default">
                           <i class="glyphicon glyphicon-heart"></i>
                        </button>
                        <button type="button" class="btn btn-default">
                           <i class="glyphicon glyphicon-trash"></i>
                        </button>
                      </div>
                      <!-- 表格冲突，也不知道咋回事，导致筛选重复，排序会死机……<table id="table11"
                        data-search="true"
                        data-show-refresh="true"
                        data-show-toggle="true"
                        data-show-columns="true"
                        data-toolbar="#projachievement"
                        data-query-params="queryParams"
                        data-filter-control="true"
                        data-filter-show-clear="true"
                        data-show-export="true"
                        >
                        <thead>
                          <tr>
                            <th data-formatter="index1">#</th>
                            <th data-field="Tnumber" data-sortable="true">成果编号</th>
                            <th data-field="Name">成果名称</th>
                            <th data-field="Category" data-filter-control="select">成果类型</th>
                            <th data-field="Count">成果数量</th>
                            <th data-field="Drawn" data-filter-control="select">编制、绘制</th>
                            <th data-field="Designd" data-filter-control="select">设计</th>
                             <th data-field="Checked" data-filter-control="select">校核</th>
                            <th data-field="Examined" data-filter-control="select">审查</th>
                          </tr>
                        </thead>
                      </table> -->
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->

        <div class="modal fade" id="modalTable1" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="bottom:auto">
            <div class="modal-dialog" style="width:auto">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">本专业项目个人贡献率</h4>
                    </div>
                    <div class="modal-body">
                    
                    <!-- 插入每个人贡献比例饼图（全年） -->
                    <div class="col-lg-12" id="main4" style="width: 800px;height:600px;"></div>
                    
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->

        <div class="modal fade" id="modalTable2" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" style="bottom:auto">
            <div class="modal-dialog" style="width:auto">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">本专业项目成果类型比例</h4>
                    </div>
                    <div class="modal-body">
                    <!-- 插入成果类型比例饼图（全年） -->
                    <div class="col-lg-12" id="main3" style="width: 800px;height:600px;"></div>
                    
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->
      </div>
    </div>
    <div class="tab-pane fade" id="achievement">
      <p>这里将显示月度实物成果情况。</p>
      <h3>科室本月实物成果表</h3>
        <div id="sec" class="btn-group">
          <button type="button" class="btn btn-default">
             <i class="glyphicon glyphicon-plus"></i>
          </button>
          <button type="button" class="btn btn-default">
             <i class="glyphicon glyphicon-heart"></i>
          </button>
          <button type="button" class="btn btn-default">
             <i class="glyphicon glyphicon-trash"></i>
          </button>
        </div>
        <table id="table12"
          data-toggle="table"
          data-url="/achievement/secprojectachievement"
          data-search="true"
          data-show-refresh="true"
          data-show-toggle="true"
          data-show-columns="true"
          data-toolbar="#sec"
          data-query-params="queryParams"
          data-filter-control="true"
          data-filter-show-clear="true"
          data-show-export="true"
          >
        <thead>        
          <tr>
          <th data-formatter="index1">#</th>
          <!--表格冲突 <th data-field="ProjectNumber" data-filter-control="select">项目编号</th> -->
          <!-- <th data-field="ProjectName" data-sortable="true" data-filter-control="select">项目名称</th> -->
          <!-- <th data-field="DesignStage" data-sortable="true" data-filter-control="select">项目阶段</th> -->
          <th data-field="Tnumber" data-sortable="true">成果编号</th>
          <th data-field="Name">成果名称</th>
          <th data-field="Category" data-filter-control="select">成果类型</th>
          <th data-field="Count">成果数量</th>
          <th data-field="Drawn" data-filter-control="select">编制、绘制</th>
          <th data-field="Designd" data-filter-control="select">设计</th>
          <th data-field="Checked" data-filter-control="select">校核</th>
          <th data-field="Examined" data-filter-control="select">审查</th>
          </tr>
        </thead>
      </table>
    </div>
  </div>
</div>

  <script type="text/javascript">
    function index1(value,row,index){
      return index+1
    }
    function actionFormatter(value, row, index) {
        return [
            '<a class="like" href="javascript:void(0)" title="成果列表">',
            '<i class="glyphicon glyphicon-list-alt"></i>',
            '</a>&nbsp;',
            '<a class="edit ml10" href="javascript:void(0)" title="贡献组成">',
            '<i class="glyphicon glyphicon-user"></i>',
            '</a>&nbsp;',
            '<a class="remove ml10" href="javascript:void(0)" title="成果类型组成">',
            '<i class="glyphicon glyphicon-adjust"></i>',
            '</a>'
        ].join('');
    }

    window.actionEvents = {
      'click .like': function (e, value, row, index) {
       //提交到后台
         $.ajax({
         // type:"post",
         url:"/achievement/projectachievement",
         data: {CatalogId:row.Id},
              success:function(data,status){
               var data1 =data ;
               $('#table11').bootstrapTable({data: data1});
               $('#table11').bootstrapTable('refresh', {url:'/achievement/projectachievement?CatalogId='+row.Id});
              }
         });
          $('#modalTable').modal({
          show:true,
          backdrop:'static'
          });
      },
      'click .edit': function (e, value, row, index) {
        //展示这个项目（阶段、专业）一年来的每个人贡献比例 
        var myChart4 = echarts.init(document.getElementById('main4'));  
        $.get('/achievement/projectuserparticipate?CatalogId='+row.Id).done(function (data) {
            // 填入数据
            myChart4.setOption({
                title: {
                text: '一年来该项目贡献组成'
            },
                tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b}: {c} ({d}%)"
            },
            legend: {
                orient: 'vertical',
                x : 'left',
                // data:{{.UserName}}
            },
            toolbox: {
                show : true,
                feature : {
                    mark : {show: true},
                    dataView : {show: true, readOnly: false},
                    magicType : {
                        show: true, 
                        type: ['pie', 'funnel'],
                        option: {
                            funnel: {
                                x: '25%',
                                width: '50%',
                                funnelAlign: 'left',
                                max: 1548
                            }
                        }
                    },
                    restore : {show: true},
                    saveAsImage : {show: true}
                }
            },
            calculable : true,
            series: [{
                name:'贡献组成',
                type:'pie',
                radius: ['10%', '60%'],
                // 根据名字对应到相应的系列
                data: data
            }]
        });
      });

      $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
        // alert('You click edit icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
    },
    'click .remove': function (e, value, row, index) {
      //展示这个项目（阶段、专业）一年来的成果类型
      var myChart3 = echarts.init(document.getElementById('main3'));
      $.get('/achievement/echarts3?CatalogId='+row.Id).done(function (data) {
            // 填入数据
        myChart3.setOption({
                title: {
                text: '一年来该项目成果类型组成',
                x:'center'
            },
                tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b}: {c} ({d}%)"
            },
            legend: {
                orient: 'vertical',
                x : 'left',
                data:{{.Select2}}
            },
            toolbox: {
                show : true,
                feature : {
                    mark : {show: true},
                    dataView : {show: true, readOnly: false},
                    magicType : {
                        show: true, 
                        type: ['pie', 'funnel'],
                        option: {
                            funnel: {
                                x: '25%',
                                width: '50%',
                                funnelAlign: 'left',
                                max: 1548
                            }
                        }
                    },
                    restore : {show: true},
                    saveAsImage : {show: true}
                }
            },
            calculable : true,
            series: [{
                name:'成果类型组成',
                type:'pie',
                radius: ['10%', '60%'],
                // 根据名字对应到相应的系列
                data: data
            }]
        });
      }); 

        $('#modalTable2').modal({
        show:true,
        backdrop:'static'
        });
    }
};
</script>
</body>
</html>