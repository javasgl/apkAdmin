<script type="text/javascript">
	new Vue({
		el:'#channel',
		methods:{
			handleCheckAllChange:function(event){
				// debugger;
				this.form.checkedChannels = event.target.checked ? this.channels : [];
				this.isIndeterminate = false;	
			},
			handleCheckedChannelsChange(value) {
				let checkedCount = value.length;
				this.checkAll = checkedCount === this.channels.length;
				this.isIndeterminate = checkedCount > 0 && checkedCount < this.channels.length;
			}
		},
		data:{
			activeNames:['1','2'],
			channels:[
			'a','b','c'
			],
			checkAll:true,
			isIndeterminate:true,
			icon:'#',
			form:{
				checkedChannels:['a','c']
			}
		}
	});
</script>
