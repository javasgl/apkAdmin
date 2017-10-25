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
