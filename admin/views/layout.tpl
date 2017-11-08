<!DOCTYPE html>
<html>
<head>
	<title>apkAdmin</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link rel="stylesheet" type="text/css" href="/static/css/element-ui/index.css">
	<?template "header.tpl"?>
</head>

<body>
	<div id="navi-bar">
		<?template "navibar.tpl" .?>
	</div>
	<div>
		<?.LayoutContent?>
	</div>
</body>
<script type="text/javascript" src="/static/js/vuejs/vue.js"></script>
<script type="text/javascript" src="/static/js/element-ui/index.js"></script>
<script type="text/javascript" src="/static/js/axios/axios.min.js"></script>
<script type="text/javascript">
	new Vue({
		el:"#navi-bar",
		methods:{
			navi:function(index){
				console.log(index)
				location.href=index
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
		data:{}
	});
</script>
<?.HtmlScripts?>
</html>
