<script type="text/javascript">
	new Vue({
		el:'#packing',
		methods:{
			handleCheckAllChange:function(event){
				this.form.checkedChannels = event.target.checked ? this.channels : [];
				this.isIndeterminate = false;	
			},
			handleCheckedChannelsChange(value) {
				let checkedCount = value.length;
				this.checkAll = checkedCount === this.channels.length;
				this.isIndeterminate = checkedCount > 0 && checkedCount < this.channels.length;
			},
			packing:function(){
				if(!this.form.apkVersion){
					this.$notify({
						title: 'Tips',
						message: '请输入客户端app版本号',
						type:'error'
					});
					return;
				}
				if(!this.form.checkedChannels.length){
					this.$notify({
						title: 'Tips',
						message: '请选择渠道',
						type:'error'
					});
					return;
				}
				axios.post('/dashboard/addJob/',this.form).then(function(res){

				});

			},
			navi:function(index){

					console.log(index)
					location.href=index
			}
		},
		mounted:function(){
			axios.get('/dashboard/getJobs').then((res)=>{
				this.jobs = res.data.jobs
			});
			axios.get('/dashboard/getChannels').then((res)=>{
				for (index in res.data.channels){
					Vue.set(this.channels,index,res.data.channels[index].channel)
				}
			});
		},
		data:{
			activeNames:['1','2'],
			channels:[],
			checkAll:true,
			isIndeterminate:true,
			icon:'#',
			form:{
				checkedChannels:[],
				apkVersion:''
			},
			jobs:[],
		}
	});
</script>
