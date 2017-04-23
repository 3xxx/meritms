<!DOCTYPE html>
<html>
 
  <head>
    <meta charset="utf-8" />
    <title></title>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    <!-- <link href="css/bootstrap.css" rel="stylesheet" /> -->
  </head>
 
  <body>
    <div id="tree" style="width: 400px;height: 1000px;margin-left: auto;margin-right: auto;"></div>
    <div id="testDate"></div>
    <!-- <script src="js/jquery.js"></script> -->
    <!-- <script src="js/bootstrap-treeview.js"></script> -->
    <script type="text/javascript">
      //获取树形结构列表数据
      function getTree() {
        var tree = [{
          text: "Parent 1",
          nodes: [{
            text: "Child 1",
            nodes: [{
              text: "Grandchild 1"
            }, {
              text: "Grandchild 2"
            }]
          }, {
            text: "Child 2"
          }]
        }, {
          text: "Parent 2"
        }, {
          text: "Parent 3"
        }, {
          text: "Parent 4"
        }, {
          text: "Parent 5"
        }];
        return tree;
      }
       
      //初始化树形结构列表
      $('#tree').treeview({
        data: getTree()
      });
      $('#tree').on('nodeSelected', function(event, data) {
          // clickNode(event, data)
          alert("Hello World!");
        }); 
      //最后一次触发节点Id
      var lastSelectedNodeId = null;
      //最后一次触发时间
      var lastSelectTime = null;
       
      //自定义业务方法
      function customBusiness(data){
        alert("双击获得节点名字： "+data.text);
      }
 
      function clickNode(event, data) {
        if (lastSelectedNodeId && lastSelectTime) {
          var time = new Date().getTime();
          var t = time - lastSelectTime;
          if (lastSelectedNodeId == data.nodeId && t < 300) {
            customBusiness(data);
          }
        }
        lastSelectedNodeId = data.nodeId;
        lastSelectTime = new Date().getTime();
      }
       
      //自定义双击事件
      function customDblClickFun(){
        //节点选中时触发
        $('#tree').on('nodeSelected', function(event, data) {
          clickNode(event, data)
          alert("Hello World!");
        });
        //节点取消选中时触发
        $('#tree').on('nodeUnselected', function(event, data) {
          clickNode(event, data)
          alert("Hello World!");
        });
      }
      $('#tree').dblclick( function () { alert("Hello World!"); });
      $(document).ready(function() {
        //customDblClickFun();
      });
    </script>
  </body>
 
</html>