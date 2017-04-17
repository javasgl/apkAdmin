  <div id="app">
    <el-form label-position="left" label-width="0px" class="login-container">
      <h3 class="title">apk packaging system</h3>
      <el-form-item prop="account">
        <el-input type="text" v-model="account" auto-complete="off" placeholder="账号"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="password" auto-complete="off" placeholder="密码"></el-input>
      </el-form-item>
      <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
      <el-form-item style="width:100%;">
        <el-button type="primary" style="width:100%;" class="login-btn">已有账号->>登录</el-button>
        <el-button type="primary" style="width:100%;" class="regist-btn">没有账号->>注册</el-button>
      </el-form-item>
    </el-form>
  </div>
