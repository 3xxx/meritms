<!-- 首页界面 -->
<!DOCTYPE html>
<html>
<head>
  <title>MeritMS</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <!--[if lt IE 9]>
    <script src="https://www.novamind.com/wp-content/themes/novamind/js/html5.js"></script>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->

<!-- <link rel="stylesheet" id="genericons-css" href="/static/NovaMind/genericons.css" type="text/css" media="all"> -->
<link rel="stylesheet" id="novamind-style-css" href="/static/NovaMind/style.css" type="text/css" media="all">


<link rel="stylesheet" id="Novamindtwelve-style-css" href="/static/NovaMind/style3.css" type="text/css" media="all"><!--这个是上面的样式-->
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<style type="text/css">
label {
color: #E6E6FA;
}
</style>
</head>

<body>
  <!--<div class="col-lg-12">
    <div class="form-group">
      <form method="post" action="/importjson" enctype="multipart/form-data">
        <div class="input-group">
          <label>
            选择json数据文件：
            <input type="file" name="json" id="json" />
          </label>
          <br/>
        </div>
        <button type="submit" class="btn btn-default" >导入结构数据</button>
      </form>
    </div>

    <div class="form-group">
      <button type="button" class="btn btn-default" id="import">初始化评测结构数据</button>
    </div>

    <form class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
      <div class="form-group">
        <label>选择excel</label>
        <input type="file" class="form-control" name="catalog" id="catalog"></div>
      <button type="submit" class="btn btn-default">提交</button>
    </form>

    <div class="form-group">
      {{if .IsLogin}}
      <a href="/login?exit=true">
        <button type="button" class="btn btn-primary">
          <span class="glyphicon glyphicon-user"></span>
          管理员退出
        </button>
      </a>
      {{else}}
      <a href="/login?url=/admin">
        <button type="button" class="btn btn-primary">
          <span class="glyphicon glyphicon-user"></span>
          管理员登录
        </button>
      </a>
      {{end}}
    </div>
    <div class="form-group">
      <a href="/json">
        <button type="button" class="btn btn-primary">
          <span class="glyphicon glyphicon-user"></span>
          查看价值结构
        </button>
      </a>
    </div>
    <div class="form-group">
      <a href="/getperson">
        <button type="button" class="btn btn-primary">
          <span class="glyphicon glyphicon-user"></span>
          价值排序
        </button>
      </a>
    </div>
    <div class="form-group">
      <a href="/user">
        <button type="button" class="btn btn-primary">
          <span class="glyphicon glyphicon-user"></span>
          查看个人
        </button>
      </a>
    </div>
  </div>-->


<div class="container-fluid blue-bg-rpt">
  <div class="row">
    <h2 class="how-it-works">How can Merit help me?</h2>
    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box presentation home-how-box" id="home-presentation" style="cursor: pointer;">
        <h3>成果登记</h3>
        <p>
          管理人员轻松进行设计人员的成果登记和分析。
        </p>
      </div>
    </div>
    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box visual-pl home-how-box" id="visualpl" style="cursor: pointer;">
        <h3>价值管理</h3>
        <p> 发现技术范们的价值，展示价值<br/>
        努力积累、提升价值<br/>
        为价值自豪！
        </p>
      </div>
    </div>
    
    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box studying home-how-box" id="home-studying" style="cursor: pointer;">
        <h3>奖金分配</h3>
        <p>根据项目产值、成果排名和个人价值进行分配，即双轨制奖金分配方案。</p>
      </div>
    </div>
    <div class="col-xs-6 col-sm-3 col-md-3 col-lg-3">
      <div class="hlp-box pitching home-how-box" id="home-pitching" style="cursor: pointer;">
        <h3>优雅地管理</h3>
        <p>
          工程师利用HydroCMS对个人项目和设计成果进行网络化管理。
        </p>
      </div>
    </div>
    
  </div>
<br/>
<br/>
<br/>

  <div class="col-md-12">
    <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
      <div class=" broad-feature-12">
        <div class="broad-feature feature1-box">
          <div class="col-480-12">
            <h3 class="col-480-12">初始化部门、价值结构</h3>
          <div class="form-group">
            <form class="form-inline" method="post" action="/importjson" enctype="multipart/form-data">
              <div class="form-group">
                <label>
                  选择部门、价值结构数据(json格式)：
                  <input type="file" class="form-control" name="json" id="json" />  
                </label>
                <br/> 
              </div>
              <button type="submit" class="btn btn-default" >提交</button>
            </form>
          </div>
          <div class="form-group">
            <form method="post" action="/user/importexcel" enctype="multipart/form-data">
              <div class="form-inline" class="form-group">
                <label>选择用户数据文件(Excel)：
                  <input type="file" class="form-control" name="excel" id="excel" /> </label>
                <br/>          
              </div>
              <button type="submit" class="btn btn-default">提交</button>
            </form>
          </div>
          <!-- 保留<div class="form-group">
            <button type="button" class="btn btn-primary" id="import">初始化评测结构数据</button>
          </div> 保留-->
        </div>
      </div>
    </div>
  </div>
    <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
      <div class=" broad-feature-12">
        <div class="broad-feature feature2-box">
          <div class="col-480-12">
            <h3 class="col-480-12">导入成果数据</h3>
         <form id="form1" class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
            <div class="form-group">
              <label>选择成果登记数据文件(Excel)
              <input type="file" class="form-control" name="catalog" id="catalog"></label>
              <br/>
              </div>
            <button type="submit" class="btn btn-default" onclick="return import_xls_catalog();">提交</button>
          </form>
          <br/>
          <div class="form-group">
            <a href="/achievement/ratio">
              <button type="button" class="btn btn-primary">
                <span class="glyphicon glyphicon-user"></span>
                编辑成果类型、折标系数表
              </button>
            </a>
          </div>
        </div>
      </div>
    </div>
    </div>

    <div class="col-480-12 col-xs-4 col-sm-4 col-md-4">
      <div class="col-480-12 broad-feature-12 ">
        <div class=" broad-feature feature3-box">
          <div class="col-480-12">
            <h3 class="col-480-12">价值管理</h3>
            <!-- <p class="col-480-12 broad-box-p">利用网络化、自动化管理，体高品质.</p> -->
            <div class="form-group">
            {{if .IsLogin}}
            <a href="/login?exit=true">
              <button type="button" class="btn btn-primary">
                <span class="glyphicon glyphicon-user"></span>
                管理员退出
              </button>
            </a>
            {{else}}
            <a href="/login?url=/admin">
              <button type="button" class="btn btn-primary">
                <span class="glyphicon glyphicon-user"></span>
                管理员登录
              </button>
            </a>
            {{end}}
          </div>
          <div class="form-group">
            <a href="/json">
              <button type="button" class="btn btn-primary">
                <span class="glyphicon glyphicon-user"></span>
                编辑部门、价值结构
              </button>
            </a>
          </div>
          <div class="form-group">
            <a href="/getperson">
              <button type="button" class="btn btn-primary">
                <span class="glyphicon glyphicon-user"></span>
                查看所有人价值排序
              </button>
            </a>
          </div>
          <div class="form-group">
            <a href="/user">
              <button type="button" class="btn btn-primary">
                <span class="glyphicon glyphicon-user"></span>
                查看自己价值
              </button>
            </a>
          </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>


<script>
// $('#getstarted').click(function () {
        
//      location.href='/download/?id='+$('#download_app').val();
//      });
$('#home-presentation').css({'cursor': 'pointer'});
    
$('#home-presentation').click(function () {
    
location.href='/achievement';
  
});
$('#visualpl').css({'cursor': 'pointer'});
    
$('#visualpl').click(function () {
    
location.href='/merit';
  
});

$('#home-studying').css({'cursor': 'pointer'});
    
$('#home-studying').click(function () {
    
location.href='/dollars';
  
});

$('#home-pitching').css({'cursor': 'pointer'});
    
$('#home-pitching').click(function () {
    
location.href='http://192.168.9.13';
  
});

$(document).ready(function(){
$("#import").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
            $.ajax({
                type:"POST",
                url:"/importjson",
                success:function(data){//数据提交成功时返回数据
                    alert("导入成功！")
                }
            });
            return true;//这里true和false结果都一样。不刷新页面的意思？
 });
});

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
  </script>
</body>
</html>
<!-- <button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span>
User
</button>
-->