<div id="packing">
	<el-row>
		<el-col :span="4">
			<el-menu mode="horizontal" theme="dark">
				<el-menu-item index="-">
					apksys
				</el-menu-item>
			</el-menu>
		</el-col>
		<el-col :span="18">
			<el-menu mode="horizontal" theme="dark" default-active="/dashboard/packing" @select="navi">
				<el-menu-item index="/dashboard/packing">packing center</el-menu-item>
				<el-menu-item index="/dashboard/channel">channel manage</el-menu-item>
			</el-menu>
		</el-col>
		<el-col :span="2">
			<el-menu mode="horizontal" theme="dark">
				<el-submenu index="2">
					<template slot="title">我的工作台</template>
					<el-menu-item index="2-3">退出登录</el-menu-item>
				</el-submenu>
			</el-menu>
		</el-col>
	</el-row>
	<el-row>
		<el-col>
			<el-collapse v-model="activeNames">
				<el-collapse-item title="packing" name="1">
					<el-row>
						<el-col :span="8">
							<div>
								当前时间:{{new Date().toLocaleString()}}
							</div>
						</el-col>
						<el-col :span="8">
							<el-form label-width="80px">
								<el-form-item label="选择应用">
									<el-radio-group v-model="form.appId">
										<el-radio v-for="app in apps" :label="app.appId">{{app.app}}</el-radio>
									</el-radio-group>
								</el-form-item>
								<el-form-item label="发布类型">
									<el-radio-group v-model="form.releaseType">
										<el-radio v-for="type in releaseTypes" :label="type.typeId">{{type.type}}</el-radio>
									</el-radio-group>
								</el-form-item>
								<el-form-item label="版本号">
									<el-input v-model="form.apkVersion" placeholder="例如: 8.2.6"></el-input>
								</el-form-item>
								<el-form-item label="应用名">
									<el-input v-model="form.appName" placeholder="例如: icy"></el-input>
								</el-form-item>
								<el-form-item label="渠道">
									<el-checkbox :indeterminate="isIndeterminate" v-model="checkAll" @change="handleCheckAllChange">全选</el-checkbox>
									<el-checkbox-group v-model="form.checkedChannels" @change="handleCheckedChannelsChange">
										<el-checkbox v-for="channel in channels" :label="channel" :key="channel">{{channel}}</el-checkbox>
									</el-checkbox-group>
								</el-form-item>
								<el-form-item label="">
									<el-button type="primary" size="small" icon="menu" @click="packing">start packing ! </el-button>
								</el-form-item>
							</el-form>
						</el-col>
						<el-col :span="8"></el-col>
					</el-row>
				</el-collapse-item>
				<el-collapse-item title="history" name="2">
					<el-table :data="jobs">
						<el-table-column prop="createTime" label="日期"></el-table-column>
						<el-table-column prop="creatorId" label="操作人"></el-table-column>
						<el-table-column prop="apkVersion" label="版本"></el-table-column>
						<el-table-column prop="status" label="进度"></el-table-column>
						<el-table-column prop="downloadUrl" label="下载"></el-table-column>
					</el-table>
				</el-collapse-item>
			</el-collapse>
		</el-col>
	</el-row>
</div>
