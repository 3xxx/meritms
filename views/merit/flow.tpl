<!-- 这个是显示左侧栏，右边index_user显示用户的cms情况 -->
<!DOCTYPE html>
<head>
  <title>珠三角水资源配置工程设代</title>
  <!-- <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script> -->
  <script src="/static/vue.js/vue.js"></script>
  <script src="/static/vue.js/vue-router.js"></script>
  <!-- 引入样式 -->
  <link rel="stylesheet" href="https://unpkg.com/element-ui@2.4.11/lib/theme-chalk/index.css">
  <!-- <link rel="stylesheet" href="/static/vue.js/index.css"> -->
  <!-- <script src="https://unpkg.com/element-ui@2.4.11/lib/index.js"></script> -->
  <script src="/static/vue.js/index.js"></script>
  <!-- 引入组件库 -->
  <!-- <script src="https://unpkg.com/element-ui/lib/index.js"></script> -->
  <!-- <script src="https://unpkg.com/axios/dist/axios.min.js"></script> -->
  <script src="/static/vue.js/axios.js"></script>

  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <!-- <script src="https://unpkg.com/vue/dist/vue.js"></script> -->

<style lang="scss">
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
    line-height:36px;
    text-indent : 25px;
  }

  .grid-content .cusheader {
    min-height: 60px;
    line-height:60px;
    text-indent : 5px;
  }

  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }
 
  
</style>
</head>

<body>
<div id="app">
  <el-container>

<el-menu default-active="activeIndex" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" :collapse="isCollapse" router>
  <el-menu-item index="index">
    <i class="el-icon-menu"></i>
    <span slot="title">首页</span>
  </el-menu-item>
  <el-submenu index="1">
    <template slot="title">
      <i class="el-icon-location"></i>
      <span slot="title" @click="jump('project')">项目</span>
    </template>
    <!-- <el-menu-item-group> -->
      <el-menu-item index="bar" @click="jump('design')">设计</el-menu-item>
      <el-menu-item index="1-2">建设</el-menu-item>
      <el-menu-item index="1-3">监理</el-menu-item>
      <!-- <el-menu-item index="1-4">施工</el-menu-item> -->
    <!-- </el-menu-item-group> -->
  </el-submenu>
  <el-submenu index="2">
    <template slot="title">
      <i class="el-icon-printer"></i>
      <span slot="title">标段</span>
    </template>
    <!-- <el-menu-item-group title="A标"> -->
      <el-submenu index="2-1">
        <span slot="title">A标</span>
        <el-menu-item index="2-1-1">A1标</el-menu-item>
        <el-menu-item index="2-1-2">A2标</el-menu-item>
        <el-menu-item index="2-1-3">A3标</el-menu-item>
        <el-menu-item index="2-1-4">A4标</el-menu-item>
        <el-menu-item index="2-1-5">A5标</el-menu-item>
        <el-menu-item index="2-1-6">A6标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-2">
        <span slot="title">B标</span>
        <el-menu-item index="2-2-1">B1标</el-menu-item>
        <el-menu-item index="2-2-2">B2标</el-menu-item>
        <el-menu-item index="2-2-3">B3标</el-menu-item>
        <el-menu-item index="2-2-4">B4标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-3">
        <span slot="title">C标</span>
        <el-menu-item index="2-3-1">C1标</el-menu-item>
        <el-menu-item index="2-3-2">C2标</el-menu-item>
        <el-menu-item index="2-3-3">C3标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-4">
        <span slot="title">D标</span>
        <el-menu-item index="2-4-1">D1标</el-menu-item>
        <el-menu-item index="2-4-2">D2标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group>     -->
  </el-submenu>

  <el-menu-item index="3" @click="jump('onlyoffice')">
    <i class="el-icon-upload"></i>
    <span slot="title">onlyoffice</span>
  </el-menu-item>
  <el-menu-item index="4" disabled>
    <i class="el-icon-document"></i>
    <span slot="title">视频</span>
  </el-menu-item>
  <el-menu-item index="5">
    <i class="el-icon-setting"></i>
    <span slot="title">日记</span>
  </el-menu-item>
</el-menu>
    <!-- </el-aside> v-model语法糖，相当于v-bind:value="isCollapse" v-on:input="isCollapse = $event.target.value-即label的true或false"-->
<el-container>
  <el-header class="grid-content bg-purple-light">
    <el-row>
      <el-col :span="12" style="text-align: left;" class="cusheader">
        <!-- <el-switch v-model="isCollapse" active-color="#13ce66" inactive-color="#ff4949">
        </el-switch> -->
        <i v-model="isCollapse" @click="changeCollapse(isCollapse)" :class="{'el-icon-d-arrow-right':isCollapse,'el-icon-d-arrow-left':!isCollapse}"></i>
      </el-col>
      <el-col :span="12" style="text-align: right;" class="cusheader">
        <el-dropdown>
          <i class="el-icon-setting" style="margin-right: 5px"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item @click.native="dialogFormVisible = true">登陆</el-dropdown-item>
            <el-dropdown-item>大事记</el-dropdown-item>
            <el-dropdown-item>查阅</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <span>设代处</span>
      </el-col>
    </el-row>
  </el-header>

  <el-main>

  <el-row>
    <el-col :span="24" style="text-align: left;">
      <div class="grid-content bg-purple-light">最新消息</div>
    </el-col>
  </el-row>

  <el-row>
    <el-col :span="24" style="text-align: left;">
      <div class="grid-content bg-purple-light">人员简介</div>
    </el-col>
  </el-row>

  </el-main>


  </el-container>
</el-container>
  <el-footer style="text-align: center;">Copyright © 2016-2019 EngineerCMS</el-footer>

<!-- <el-button round @click="dialogFormVisible = true">登录</el-button> -->
<el-dialog title="系统登录" :visible.sync="dialogFormVisible" center>
  <!-- 插入测试 -->
  <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
    <el-form-item label="账号" prop="num">
      <el-input v-model.number="ruleForm2.num" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="pass">
      <el-input type="password" v-model="ruleForm2.pass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
    <el-form-item label="记住密码" prop="delivery">
      <el-switch v-model="ruleForm2.delivery"></el-switch>
    </el-form-item> 
    <span><a>忘记密码？</a></span>
  </el-form>
   <!-- 插入测试 -->
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false; resetForm('ruleForm2')">取 消</el-button>
    <el-button type="primary" @click="submitForm('ruleForm2')">登 录</el-button>
  </div>
</el-dialog>

  <transition>
    <router-view>afadf</router-view>
  </transition>

</div>


<script type="text/javascript">

  const Foo = { template: '<div>foo</div>' };
  const Bar = resolve => require(['../views/login.tpl'],resolve);
  const Main = { template: '<div></div>' };
  const routes = [
    { path: '/foo', component: Foo },
    { path: '/bar', component: Bar },
    { path: '/', component: Main }
  ];
  const router = new VueRouter({
    mode: 'history',
    routes // (缩写) 相当于 routes: routes
  });
  var app = new Vue({
      router,
      delimiters: ['[[', ']]'],
      el:'#app',
      data:{
        users:{
            name:'',
            age:''
        },
        message: 'Hello Vue!',
        info: null,

        posts:[],
      },
      data() {
        var checkNum = (rule, value, callback) => {
          if (!value) {
            return callback(new Error('账号不能为空'));
          }
          setTimeout(() => {
            if (!Number.isInteger(value)) {
              callback(new Error('请输入数字值'));
            } else {
              var myreg=/^[1][3,4,5,7,8][0-9]{9}$/;
              if (!myreg.test(value) ) {
                callback(new Error('请输入正确的手机号码'));
              } else {
                callback();
              }
  
            }
          }, 1000);
        };
        var validatePass = (rule, value, callback) => {
          if (value === '') {
            callback(new Error('请输入密码'));
          } else {
  
            callback();
          }
        };

        return {
          loginPower:false,
          checked: true,
          /*插入form方法*/
          /*设定规则指向*/
          ruleForm2: {
            pass: '',
            num: '',
             delivery: false,
          },
          rules2: {
            pass: [
              { validator: validatePass, trigger: 'blur' }
            ],
  
            num: [
              { validator: checkNum, trigger: 'blur' }
            ]
          },
          /*插入form方法*/
          dialogTableVisible: false,
          dialogFormVisible: false,
          form: {
            name: '',
            type: [],
            resource: '',
            desc: ''
          },
          formLabelWidth: '120px',

          isCollapse: true,
          bannerHeight:200,

          clientHeight:'',
          // calleft:0
          posts:[],
          numbers:0,

          dialogFormVisible: false,
          form: {
            name: '',
            region: '',
            date1: '',
            date2: '',
            delivery: false,
            type: [],
            resource: '',
            desc: ''
          },
          formLabelWidth: '120px'
        };
      },
      //注意这里我是将它的改为这样的，效果是一样的，但使用这个可以在页面任何部位设置跳转时头部导航条部分样式会跟这变化即为选中，而不是不变的例如：你设置跳转到产品页导航条的样式也会跟着选中产品项
      computed:{
        activeIndex(){
          return this.$route.path.replace('/','')
        }
      },

      mounted:function () {
        this.sendGetByObj();
        this.setSize();
        const that = this;
        window.addEventListener('resize', function() {
          that.screenWidth = $(window).width();
          that.setSize();
        }, false);

        // 获取浏览器可视区域高度
        this.clientHeight =   `${document.documentElement.clientHeight}`          //document.body.clientWidth;
        console.log(this.clientHeight);
        window.onresize = function temp() {
          this.clientHeight = `${document.documentElement.clientHeight}`;
        };
      },
      watch: {
        initData () {
          let H = document.querySelector('.boxShadow')
          H.style.height = ''
          setTimeout(() => {
            console.log(H.offsetHeight)
            if (H.offsetHeight < window.innerHeight) {
              document.body.style.height = window.innerHeight + 'px'
              H.style.height = window.innerHeight - 30 + 'px'
            } else {
              document.body.style.height = H.offsetHeight + 'px'
              H.style.height = ''
            }
          }, 300)
        },

      },
      proxyTable: {
        '/api':{//此处并非一定和url一致。
          target:'https://zsj.itdos.com',
          changeOrigin:true,//允许跨域
          pathRewrite:{
            '^/api': ''
          }
        }
      },      
      methods:{
        submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
            //提交成功做的动作
            dialogFormVisible = false;
              /* alert('submit!') ; */
              this.$message({
                type: 'success',
                message: '提交成功' 
              });
            } else {
              console.log('error submit!!');
              return false;
            }
          }); 
        },

        resetForm(formName) {
          this.$refs[formName].resetFields();
        },

        changeWidthL(key, keyPath) {
          console.log(key, keyPath);
        },

        handleOpen(key, keyPath) {
          console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
          console.log(key, keyPath);
        },
        sendGetByStr(){
          axios.get(`api/v1/wx/getlistarticles?page=1`)//1.get通过直接发字符串拼接
            // .then(function (response) {
            //   console.log(response);
            //   console.log(response.data.info);
            // }
              .then(response => (this.info = response.data.info)
              )
            .catch(function (error) {
              console.log(error);
            });
        },
        sendGetByObj(){
          axios.get(`/v1/wx/getlistarticles`,{//2.get通过params选项
            params:{
                page:1
            }
          })
          // axios({
          //  headers: {
          //   'X-Requested-With': 'XMLHttpRequest',
          //   'Content-Type': 'application/json; charset=UTF-8',
          //   'Access-Control-Allow-Origin': '*'
          //   },//设置跨域请求头
          //   method: "GET",//请求方式
          //   url: "https://zsj.itdos.com/v1/wx/getlistarticles",//请求地址
          //   params:{
          //       page:1
          //   }
            // data: {
            //   "menu_id":1,
            //   "thirdapp_id":1//请求参数
            // }
          // })
            .then(response => (this.posts = response.data.articles))
            .catch(function (error) {
              console.log(error);
          });
        },
        //html剔除富文本标签，留下纯文本
        getSimpleText(html){
          var re1 = new RegExp("<.+?>","g");//匹配html标签的正则表达式，"g"是搜索匹配多个符合的内容
          var msg = html.replace(re1,'');//执行替换成空字符
          return '<br>'+msg.substring(0,20);
        },
        even: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers % 2 === 0
          // })
        },
        even1: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers === 0
          // })
        },
        changeCollapse:function(isCollapse){
          if (isCollapse==true) {
            this.isCollapse=false
          }else{
            this.isCollapse=true
          }
        },
        skip(a){
          this.$router.push(a)
        },
        jump(select){
          console.log(select);
          let routeUrl = this.$router.resolve({
            path: "/onlyoffice",
            // query: {id:96}
          });
          switch (select) {
            case ("onlyoffice"):
              routeUrl = this.$router.resolve({
                path: "/onlyoffice",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            case ("project"):
              routeUrl = this.$router.resolve({
                path: "/project",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            case ("design"):
              routeUrl = this.$router.resolve({
                path: "/design",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            default:
              routeUrl = this.$router.resolve({
                path: "/index",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
          }
          //this.$router.push({path: '/cart?goodsId=12'})
          //this.$router.go(-2)
          //后退两步
          // let routeUrl = this.$router.resolve({
          //   path: "/onlyoffice",
          //   query: {id:96}
          // });
          // window.open(routeUrl .href, '_blank');
          // window.open(routeUrl .href);
        }

      }
  });

</script>
</body>

<!-- 注：对于路径的写法： ./ 当前目录 ../ 父级目录 / 根目录
<el-carousel trigger="click" :height="bannerH +'px'">
   <el-carousel-item v-for="(item,index) in bannerImgLst" :key="index">
       <img :src="'https://mirror198829.github.io/static/cloud/'+item" class="bannerImg"/>
   </el-carousel-item>
</el-carous>

export default {
  name: 'homePage',
  data () {
    return {
      bannerH:200,
      bannerImgLst:['navBg1.png','navBg2.png','navBg3.jpg']
    }
  },
  methods:{
    setBannerH(){
      this.bannerH = document.body.clientWidth / 4
    }
  },
  mounted(){
    this.setBannerH()
    window.addEventListener('resize', () => {
      this.setBannerH()
    }, false)
  },
  created(){}
} -->