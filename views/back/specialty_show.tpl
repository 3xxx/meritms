<!-- 专业成本分布，测试用-->
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
<script type="text/javascript" src="/static/js/echarts.min.js"></script>
<style>
i#delete
{
color:#C71585;
}
</style>
</head>


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
        <label class="control-label">tips:(StartDay < DateRange <= EndDay)</label>
    </div>
<!-- </form>
-->
<br>

<div id="main" style="width: 800px;height:600px;"></div>
</div>
<script type="text/javascript">
        // 基于准备好的dom，初始化echarts实例
        // var myChart = echarts.init(document.getElementById('main'));
        // 指定图表的配置项和数据
        // var option = {
        //     title: {
        //         text: 'ECharts 入门示例'
        //     },
        //     tooltip: {},
        //     legend: {
        //         data:['销量']
        //     },
        //     xAxis: {
        //         data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
        //     },
        //     yAxis: {},
        //     series: [{
        //         name: '销量',
        //         type: 'bar',
        //         data: [5, 20, 36, 10, 10, 20]
        //     }]
        // };
        // 使用刚指定的配置项和数据显示图表。
        // myChart.setOption(option);


// $.get('data.json').done(function (data) {
//     myChart.setOption({
//         title: {
//             text: '异步数据加载示例'
//         },
//         tooltip: {},
//         legend: {
//             data:['销量']
//         },
//         xAxis: {
//             data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
//         },
//         yAxis: {},
//         series: [{
//             name: '销量',
//             type: 'bar',
//             data: [5, 20, 36, 10, 10, 20]
//         }]
//     });
// });



var myChart = echarts.init(document.getElementById('main'));
// 显示标题，图例和空的坐标轴
// myChart.setOption({
//     title: {
//         text: '成本分布'
//     },
//     tooltip: {},
//     legend: {
//         data:['工作量']
//     },
//     xAxis: {
//         data: []
//     },
//     yAxis: {},
//     series: [{
//         name: '工作量',
//         type: 'bar',
//         data: []
//     }]
// });

// 异步加载数据
// $(selector).get(url,data,success(response,status,xhr),dataType)
// $.ajax({
//   url: url,
//   data: data,
//   success: success,
//   dataType: dataType
// });
// $.get("test.cgi", { name: "John", time: "2pm" },
//   function(data){
//     alert("Data Loaded: " + data);
//   });
$.get('/achievement/echarts').done(function (data) {
    // 填入数据
    myChart.setOption({
        // xAxis: {
        //     data: {{.Select2}}
        // },
        title: {
        text: '水工专业成本组成'
    },
        tooltip: {
        trigger: 'item',
        formatter: "{a} <br/>{b}: {c} ({d}%)"
    },
    legend: {
        orient: 'vertical',
        x: 'right',
        data:{{.Select2}}
    },
        series: [{
            name:'水工专业成本组成',
            type:'pie',
            radius: ['10%', '60%'],
            // 根据名字对应到相应的系列
            data: data
        }]
    });
});




// app.title = '嵌套环形图';

// option = {
//     tooltip: {
//         trigger: 'item',
//         formatter: "{a} <br/>{b}: {c} ({d}%)"
//     },
//     legend: {
//         orient: 'vertical',
//         x: 'left',
//         data:['直达','营销广告','搜索引擎','邮件营销','联盟广告','视频广告','百度','谷歌','必应','其他']
//     },
//     series: [
//         {
//             name:'访问来源',
//             type:'pie',
//             selectedMode: 'single',
//             radius: [0, '30%'],

//             label: {
//                 normal: {
//                     position: 'inner'
//                 }
//             },
//             labelLine: {
//                 normal: {
//                     show: false
//                 }
//             },
//             data:[
//                 {value:335, name:'直达', selected:true},
//                 {value:679, name:'营销广告'},
//                 {value:1548, name:'搜索引擎'}
//             ]
//         },
//         {
//             name:'访问来源',
//             type:'pie',
//             radius: ['40%', '55%'],

//             data:[
//                 {value:335, name:'直达'},
//                 {value:310, name:'邮件营销'},
//                 {value:234, name:'联盟广告'},
//                 {value:135, name:'视频广告'},
//                 {value:1048, name:'百度'},
//                 {value:251, name:'谷歌'},
//                 {value:147, name:'必应'},
//                 {value:102, name:'其他'}
//             ]
//         }
//     ]
// };
</script>
</body>
</html>