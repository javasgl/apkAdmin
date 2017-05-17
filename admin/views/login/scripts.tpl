<script type="text/javascript">
	new Vue({
		el:'#login',
		methods:{
			doLogin:function(){
				var that = this;
				axios.post('/api/login',{username:this.username,password:this.password}).then(function(res){
					if(res.data){
						location.href="/dashboard/channel";
					}else{
						that.$message.error('用户名密码不正确~');
					}
				});
			}
		},
		data:{
			checked: true,
			username:'songgl',
			password:'123456'
		}
	});
</script>
