<!-- 展示科室人员价值排名，显示个人详细价值列表==myself.tpl-->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>MeritMS</title>
      <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
      <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
      <script src="/static/js/bootstrap-treeview.js"></script>
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
      <style type="text/css">
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
      </script>
  </head>

  <div class="col-lg-12">
    <h2>{{.Sectitle}}</h2>
    <ul class="nav nav-tabs">
      <li class="active"><a href="#employee" data-toggle="tab">已通过</a></li>
      <li><a href="#year" data-toggle="tab">待提交</a></li>
      <li><a href="#proj" data-toggle="tab">分布</a></li>
    </ul>

    <div class="tab-content">
      <div class="tab-pane fade in active" id="employee">
      <br>
        <div class="form-inline">
            <table id="table"
                  data-toggle="table"
                  data-url="/merit/secofficedata"
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
                  <th data-field="Numbers">项数</th>
                  <th data-field="Marks" data-sortable="true">汇总</th>
                  <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">详细</th>
                </tr>
              </thead>
            </table>
      </div>
    </div>

    <div class="tab-pane fade" id="year">
      <p>这里将显示待提交的价值记录。</p>
    </div>

    <div class="tab-pane fade" id="proj">
      <p>这里将显示价值分布。</p>
    </div>
  </div>

      <div class="modal fade" id="modalTable" style="bottom:auto">
        <div class="modal-dialog" style="width:auto;height:800px;overflow:auto;">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span></button>
              <h4 class="modal-title">全部价值列表</h4>
            </div>
            <div class="modal-body">
              <table id="table7"
                data-toggle="table"
                data-search="true"
                data-show-refresh="true"
                data-show-toggle="true"
                data-show-columns="true"
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
                    <th data-width="10" data-radio="true"></th>
                    <th data-formatter="index1">#</th>
                    <th data-field="MeritCate">价值分类</th>
                    <th data-field="Merit">价值名称</th>
                    <th data-field="Title">价值内容名称</th>
                    <th data-field="Choose">价值选项</th>
                    <th data-field="Mark">价值分值</th>
                  </tr>
                </thead>
              </table>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
            </div>
          </div><!-- /.modal-content -->
        </div><!-- /.modal-dialog -->
      </div>


</div>

<script type="text/javascript">
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
  }
  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('L');
  }
  function queryParams(params) {
  // var date=$("#datefilter").val();
  // var secid=$("#secid").val();
  // var level=$("#level").val();
  // var mid={{.Merit.Id}};
        // params.datefilter=date;
        params.secid={{.Secid}};//传 secid=科室id 给后台
        // params.level=level;
        // params.mid=mid;
        return params;
  }
  function actionFormatter(value, row, index) {
    return  '<a class="like" href="javascript:void(0)" title="详细"><i class="glyphicon glyphicon-list-alt"></i></a>&nbsp;'
  }

    window.actionEvents = {
      'click .like': function (e, value, row, index) {
        $('#table7').bootstrapTable('refresh', {url:'/merit/myself?userid='+row.UserId});
        $('#modalTable').modal({
          show:true,
          backdrop:'static'
        });
        // alert('You click like icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
      }
    }
</script>
</body>
</html>