<style type="text/css">
	.el-input > input{
		max-width: 193px
	}
</style>
<div id="packing">
	<el-row>
		<el-col>
			<el-collapse v-model="activeNames">
				<el-collapse-item title="打包" name="1">
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
										<el-checkbox v-for="channel in channels" :label="channel.channelId" :key="channel.channelId">{{channel.channel}}</el-checkbox>
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
				<el-collapse-item title="历史记录(最近20条)" name="2">
					<el-table :data="jobs">
						<el-table-column prop="createTime" label="日期">
							<template scope="scope">
								{{scope.row.createTime|dateformat}}
							</template>
						</el-table-column>
						<el-table-column prop="creatorId" label="操作人">
							<template scope="scope">
								{{scope.row.username}}
							</template>
						</el-table-column>
						<el-table-column prop="appId" label="应用">
							<template scope="scope">
								<span v-if="scope.row.appId==3">ICY</span>
								<span v-if="scope.row.appId==1">穿衣助手</span>
							</template>
						</el-table-column>
						<el-table-column prop="releaseType" label="发布类型">
							<template scope="scope">
								<span v-if="scope.row.releaseType==0">默认</span>
								<span v-if="scope.row.releaseType==1">首发</span>
								<span v-if="scope.row.releaseType==2">活动</span>
							</template>
						</el-table-column>
						<el-table-column prop="apkVersion" label="版本"></el-table-column>
						<el-table-column prop="appName" label="应用名"></el-table-column>
						<el-table-column prop="status" label="进度">
							<template scope="scope">
								<el-progress :text-inside="true" :stroke-width="24" :percentage="42"></el-progress>
 							</template>
						</el-table-column>
						<el-table-column prop="downloadUrl" label="下载"></el-table-column>
					</el-table>
				</el-collapse-item>
			</el-collapse>
		</el-col>
	</el-row>
</div>
