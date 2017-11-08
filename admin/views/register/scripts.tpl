<script type="text/javascript">
	new Vue({
		el:'#register',
		methods:{
			doRegister:function(){

				if(this.username==''){
					this.$notify.error('请输入企业邮箱');
					return;
				}
				if(this.username.indexOf('@')!=-1){
					this.$notify.error('邮箱不需要输入@后缀');
					return;
				}
				if (this.password==''){
					this.$notify.error('请输入密码');
					return;
				}
				if(this.confirmedPassword==''){
					this.$notify.error('请输入确认密码');
					return;
				}
				if(this.confirmedPassword!=this.password){
					this.$notify.error('两次密码不一致');
					return;
				}
				axios.post('/api/register',{username:this.username,password:this.password}).then((res)=>{

					if(res.data){
						this.$notify.info('注册完成,3s后将前往登录页面')
						setTimeout(function(){
							location.href="/dashboard/packing";
						},3000);
					}else{
						this.$notify.error('注册异常~');
					}
				});
			},
			goLogin:function(){
				location.href="/"
			}
		},
		data:{
			username:'',
			password:'',
			confirmedPassword:''
		}
	});
</script>
