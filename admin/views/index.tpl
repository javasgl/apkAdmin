<!DOCTYPE html>
<html>
<head>
  <title>apkAdmin</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="static/css/element-ui/index.css">
</head>

<body>
  <div id="app">
    <el-form label-position="left" label-width="0px" class="login-container">
      <h3 class="title">apk package system</h3>
      <el-form-item prop="account">
        <el-input type="text" v-model="account" auto-complete="off" placeholder="账号"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" auto-complete="off" placeholder="密码"></el-input>
      </el-form-item>
      <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
      <el-form-item style="width:100%;">
        <el-button type="primary" style="width:100%;" class="login-btn">已有账号->>登录</el-button>
        <el-button type="primary" style="width:100%;" class="regist-btn">没有账号->>注册</el-button>
      </el-form-item>
    </el-form>
  </div>
</body>
<script type="text/javascript" src="static/js/vuejs/vue.js"></script>
<script type="text/javascript" src="static/js/element-ui/index.js"></script>
<script type="text/javascript">
  new Vue({
    el:'#app',
    data:{
      checked: true
    }
  });
</script>
<style type="text/css">
   .login-container {
    -webkit-border-radius: 5px;
    border-radius: 5px;
    -moz-border-radius: 5px;
    background-clip: padding-box;
    margin: 180px auto;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
  } 
  .title{
    margin: 0px auto 40px auto;
    text-align: center;
    color: #505458;

  }
  .remember {
    margin: 0px 0px 35px 0px;
  }
  .login-btn{
    margin-bottom: 5px;
  }
  .regist-btn{
    margin-left: 0 !important;
  }
</style>
</html>
