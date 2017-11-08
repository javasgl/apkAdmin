<el-row>
	<el-col :span="4">
		<el-menu mode="horizontal" theme="dark">
			<el-menu-item index="-">
				apksys
			</el-menu-item>
		</el-menu>
	</el-col>
	<el-col :span="18">
		<el-menu mode="horizontal" theme="dark" :default-active="activedMenu" @select="navi">
			<el-menu-item index="/dashboard/packing">packing center</el-menu-item>
			<el-menu-item index="/dashboard/channel">channel manage</el-menu-item>
		</el-menu>
	</el-col>
	<el-col :span="2">
		<el-menu mode="horizontal" theme="dark">
			<el-submenu index="2">
				<template slot="title"><?.loginUser.Username?></template>
				<el-menu-item index="--"><?.loginUser.Email?></el-menu-item>
				<el-menu-item index="-" @click="loginout">退出系统</el-menu-item>
			</el-submenu>
		</el-menu>
	</el-col>
</el-row>
