<script type="text/javascript">
	new Vue({
		el:'#app',
		methods:{
			doLogin:function(){
				axios.post('/api/login',{username:this.username,password:this.password}).then(function(res){
					console.log(res);
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
