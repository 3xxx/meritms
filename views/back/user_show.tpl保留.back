<!-- 测试页面 -->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>技术人员价值评测系统</title>
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
</head>


<div id="treeview" class="col-xs-3"></div>

<!-- <div class="col-lg-9">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#序号</th>
        <th>单位</th>
        <th>分院</th>
        <th>科室</th>
        <th>价值分类</th>
        <th>价值项</th>
        <th>选择项</th>
        <th>操作</th>
      </tr>
    </thead>

    <tbody>
      <tr>
        <th></th>
        <th>{{.Input.Danwei}}</th>
        <th></th>
        <th></th>
        <th></th>
        <th></th>
        <th></th> 
        <th></th>
        {{range $k0,$v0 :=$.Input.Fenyuan}}
                  <tr>
                  <th></th>
                  <th></th>
                  <th>{{.Department}}</th>
                  <th></th>
                  <th></th>
                  <th></th> 
                  <th></th>
                  <th></th>
                  </tr>
            {{range $k0,$v0 :=.Bumen}}
                  <tr>
                  <th></th>
                  <th></th>
                  <th></th>
                  <th>{{.Keshi}}</th>
                  <th></th>
                  <th></th> 
                  <th></th>
                  <th></th>
                  </tr>

                {{range $k0,$v0 :=.Kaohe}}
                    <tr>
                    <th></th>
                    <th></th>
                    <th></th>
                    <th></th>
                    <th>{{.Category}}</th>
                    <th></th> 
                    <th></th>
                    <th></th>
                    </tr>

                    {{range $k0,$v0 :=.Fenlei}}
                      <tr>
                      <th></th>
                      <th></th>
                      <th></th>
                      <th></th>
                      <th></th> 
                      <th>{{.Project}}</th>
                      <th></th>
                      <th>
                        <a href="/">显示成果</a>
                        <a href="/">修改</a>
                        <a href="/">删除</a></th>
                      </tr>

                      {{range $k0,$v0 :=.Xuanze}}
                        <tr>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th></th>
                        <th>{{.Choose}}</th>
                        <th></th>
                        </tr>
                      {{end}}
                    {{end}}
                {{end}}
            {{end}}    
        {{end}}
      </tr>
    </tbody>
  </table>
</div> -->
<!-- 注释掉上面的，改用页内嵌套框架 -->
<div class="col-lg-9">
        <iframe src="/secofficeshow" name='main' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onLoad="iFrameHeight()"></iframe>
</div> 

<button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span> User
</button>

<button type="button" class="btn btn-primary btn-lg" style="text-shadow: black 5px 3px 3px;">
<span class="glyphicon glyphicon-user"></span> User
</button>
<script type="text/javascript">
$(function() {
        var defaultData = [
          {
            "text": 'Parent 1',
            "selectable":false,
            // icon: "glyphicon glyphicon-stop",
            // selectedIcon: "glyphicon glyphicon-heart",
            href: '#parent1',
            tags: ['4'],
            // state: {
            // checked: true,
            // disabled: false,
            // expanded: false,
            // selected: true
            // },
            // tags: ['available'],
            nodes: [
              {
                text: 'Child 1',
                selectable:false,
                // icon: "glyphicon glyphicon-stop",
                // selectedIcon: "glyphicon glyphicon-heart",                
                // href: '#child1',
                tags: [2,3],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    href: '#grandchild1',
                    tags: ['0']
                  },
                  {
                    text: 'Grandchild 2',
                    href: '#grandchild2',
                    tags: ['0']
                  }
                ]
              },
              {
                text: 'Child 2',
                href: '#child2',
                tags: ['0']
              }
            ]
          },
          {
            text: 'Parent 2',
            href: '#parent2',
            tags: ['0'],
            nodes: [
              {
                text: 'Child 1',
                href: '#child1',
                tags: ['2'],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    href: '#grandchild1',
                    tags: ['0']
                  },
                  {
                    text: 'Grandchild 2',
                    href: '#grandchild2',
                    tags: ['0']
                  }
                ]
              },
              {
                text: 'Child 2',
                href: '#child2',
                tags: ['0']
              }
            ]
          },
          {
            text: 'Parent 3',
            href: '#parent3',
             tags: ['0']
          },
          {
            text: 'Parent 4',
            href: '#parent4',
            tags: ['0']
          },
          {
            text: 'Parent 5',
            href: '#parent5'  ,
            tags: ['0']
          }
        ];

        var alternateData = [
          {
            text: 'Parent 1',
            tags: ['2'],
             state: {
            checked: true,
            disabled: false,
            expanded: false,
            selected: true
            },           
            nodes: [
              {
                text: 'Child 1',
                tags: ['3'],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    tags: ['6']
                  },
                  {
                    text: 'Grandchild 2',
                    tags: ['3']
                  }
                ]
              },
              {
                text: 'Child 2',
                tags: ['3']
              }
            ]
          },
          {
            text: 'Parent 2',
            tags: ['7']
          },
          {
            text: 'Parent 3',
            icon: 'glyphicon glyphicon-earphone',
            href: '#demo',
            tags: ['11']
          },
          {
            text: 'Parent 4',
            icon: 'glyphicon glyphicon-cloud-download',
            href: '/demo.html',
            tags: ['19'],
            selected: true
          },
          {
            text: 'Parent 5',
            icon: 'glyphicon glyphicon-certificate',
            color: 'pink',
            backColor: 'red',
            href: 'http://www.tesco.com',
            tags: ['available','0']
          }
        ];

        var json=[
        {
  "text": "省水利设计分院",
  "Mark2": "总院",
  "Id":"01",
  // "selectable": false,
  "nodes": [
    {
      "text": "施工预算分院",
      "Mark2": "分院",
      "Id":"11",
      // "selectable": false,
             state: {
            checked: false,
            disabled: false,
            expanded: false,
            selected: false
            },      
      "nodes": [
        {
          "text": "水工室",
          "Id":"111",
          "Mark2": "科室",
          // "href": "",
          // "target":"main",
          // "selectable": false,
          "nodes": [
            {
              "text": "项目管理类",
              "Id":"1111",
              "Mark2": "用户1",
              // "selectable": false,
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "项目负责人",
                  "Id":"11111",
                  "Mark2": "用户2",
                  // "href": "/add?id=165",
                  "tags": [
                    3,
                    1
                  ],
                  "Xuanze": ",大型,中型,小型",
                  "Mark1": ",4,3,2"
                },
                {
                  "text": "课题研究",
                  // "href": "/add?id=166",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "4",
                  "Xuanze": "",
                  "Mark1": ""
                }
              ],
              "Parent2": ""
            },
            {
              "text": "贡献类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "获奖",
                  "href": "/add?id=168",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",国家级,省级,院级",
                  "Mark1": ",4,3,2"
                },
                {
                  "text": "开发",
                  // "href": "/add?id=169",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",系统级,标准",
                  "Mark1": ",5,2"
                }
              ],
              "Parent2": ""
            }
          ]
        },
        {
          "text": "施工室",
          "selectable": false,
          "nodes": [
            {
              "text": "项目管理类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "施工负责人",
                  // "href": "/add?id=172",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "4",
                  "Xuanze": "",
                  "Mark1": ""
                },
                {
                  "text": "课题a",
                  // "href": "/add?id=173",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "2",
                  "Xuanze": "",
                  "Mark1": ""
                }
              ],
              "Parent2": ""
            }
          ]
        }
      ]
    },
    {
      "text": "水工分院",
      "selectable": false,
      "nodes": [
        {
          "text": "水工室",
          "selectable": false,
          "nodes": [
            {
              "text": "项目管理类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "项目负责人",
                  // "href": "/add?id=177",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",大型,中型,小型",
                  "Mark1": ",6,4,2"
                }
              ],
              "Parent2": ""
            },
            {
              "text": "贡献类",
              "tags": [
                4,
                2
              ],
              "nodes": [
                {
                  "text": "获奖",
                  // "href": "/add?id=179",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",国家级,省级,院级",
                  "Mark1": ",4,3,2"
                },
                {
                  "text": "开发",
                  // "href": "/add?id=180",
                  "tags": [
                    3,
                    1
                  ],
                  "Mark2": "",
                  "Xuanze": ",系统级,标准",
                  "Mark1": ",5,2"
                }
              ],
              "Parent2": ""
            }
          ]
        }
      ]
    }
  ]
}
        ];
          // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
          data: json,//defaultData,
          // data:alternateData,
          enableLinks:true,
          showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        });

        $('#treeview').on('nodeSelected', function(event, data) {
          // clickNode(event, data)  
            // alert(JSON.stringify(data));
            alert("名称："+data.text);
            alert("节点id："+data.nodeId);
            alert("部门id："+data.Id);  
            alert("部门级别："+data.Mark2);
          document.getElementById("iframepage").src="/secofficeshow?secid="+data.nodeId;
          // document.getElementById("iframepage").src="http://www.baidu.com";

          //这句也有效哦 var arr = $('#treeview').treeview('getSelected');
            // alert(JSON.stringify(arr));

            
            // for (var key in arr) {
            //     alert(arr[key].id);
            // };


            // $("#btn").click(function (e) {
            // var arr = $('#tree').treeview('getSelected');
            // for (var key in arr) {
            //     c.innerHTML = c.innerHTML + "," + arr[key].id;
            // }

        // })
        });

         // $("#btn").click(function (e) {

            

        // })
          
});

// 自动适应高度 
function iFrameHeight() {   
var ifm= document.getElementById("iframepage");   
var subWeb = document.frames ? document.frames["iframepage"].document : ifm.contentDocument;   
if(ifm != null && subWeb != null) {
   ifm.height = subWeb.body.scrollHeight;
   ifm.width = subWeb.body.scrollWidth;
}   
} 

</script>

</body>
</html>
