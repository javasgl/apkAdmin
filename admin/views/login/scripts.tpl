<script type="text/javascript">
	new Vue({
		el:'#app',
		methods:{
			doLogin:function(){
				axios.get('/api/login').then(function(res){
					console.log(res);
				});
			}
		},
		data:{
			checked: true,
			account:'songgl',
			password:'123456'
		}
	});
</script>
