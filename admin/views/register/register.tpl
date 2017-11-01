<div id="register">
	<el-form label-position="left" label-width="0px" class="register-container">
		<h3 class="title">apk packing system</h3>
		<el-form-item prop="username">
			<el-input type="text" v-model="username" auto-complete="off" placeholder="请填写企业邮箱">
				<template slot="append"><?.registerMail?></template>
			</el-input>
		</el-form-item>
		<el-form-item prop="password">
			<el-input type="password" v-model="password" auto-complete="off" placeholder="密码"></el-input>
		</el-form-item>
		<el-form-item prop="confirmedPassword">
			<el-input type="password" v-model="confirmedPassword" auto-complete="off" placeholder="确认密码"></el-input>
		</el-form-item>
		<el-form-item style="width:100%;">
			<el-button type="primary" style="width:100%;" class="login-btn" @click="doRegister">注册</el-button>

			<el-button type="primary" style="width:100%;" class="regist-btn" @click="goLogin">已有账号？->>去登录</el-button>

		</el-form-item>
	</el-form>
</div>
