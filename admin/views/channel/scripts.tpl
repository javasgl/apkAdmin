<script type="text/javascript">
	new Vue({
		el:'#channel',
		methods:{
			navi:function(index){
					location.href=index
			},
			addChannel:function(){
				if(!this.form.channel){
					this.$notify({
						title: 'Tips',
						message: '请输入渠道标识',
						type:'error'
					});
					return;
				}
				if(!this.form.remark){
					this.$notify({
						title: 'Tips',
						message: '请填写渠道备注',
						type:'error'
					});
					return;
				}
				axios.post('/dashboard/addChannel/',this.form).then(function(res){
					location.reload(true)
				});
			},
			delChannel:function(index,channel){
				this.$confirm('确认要删除 '+channel.channel+' 渠道?', '提示', {
					confirmButtonText: '确定',
					cancelButtonText: '取消',
					type: 'warning'
				}).then(() => {
					axios.post('/dashboard/delChannel',{channelId:channel.channelId}).then((res)=>{
						this.channels.splice(index,1)
						this.$message({
							type: 'success',
							message: '删除成功!'
						});
					});
				}).catch(() => {
				});

			},
			loginout:function(){
				axios.post('/dashboard/loginout').then((res)=>{
					if(res){
						location.href="/"
						return;
					}
				});
			}
		},
		mounted:function(){
			axios.get('/dashboard/getChannels').then((res)=>{
				this.channels = res.data.channels
			});
		},
		data:{
			activeNames:["1","2"],
			channels:[],
			form:{
				channel:'',
				remark:''
			}
		}
	});
</script>
