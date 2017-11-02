<style type="text/css">
	.el-input > input{
		max-width: 193px
	}
</style>

<div id="channel">
	<el-row>
		<el-col>
			<el-collapse v-model="activeNames">
				<el-collapse-item title="添加渠道" name="1">
					<el-row>
						<el-col :span="8">
							<div>
								当前时间:{{new Date().toLocaleString()}}
							</div>
						</el-col>
						<el-col :span="8">
							<el-form label-width="80px">
								<el-form-item label="渠道">
									<el-input v-model="form.channel" placeholder="eg: xiaomi"></el-input>
								</el-form-item>
								<el-form-item label="备注">
									<el-input v-model="form.remark" placeholder="eg: 小米渠道"></el-input>
								</el-form-item>
								<el-form-item label="">
									<el-button type="primary" size="small" icon="plus" @click="addChannel">添加渠道</el-button>
								</el-form-item>
							</el-form>
						</el-col>
						<el-col :span="8"></el-col>
					</el-row>
				</el-collapse-item>
				<el-collapse-item title="渠道列表" name="2">
					<el-table :data="channels">
						<el-table-column prop="channelId" label="channelId"></el-table-column>
						<el-table-column prop="channel" label="渠道"></el-table-column>
						<el-table-column prop="remark" label="备注"></el-table-column>
						<el-table-column prop="createTime" label="日期"></el-table-column>
						<el-table-column label="操作">
							<template scope="scope">
								<el-button size="small" type="danger" @click="delChannel(scope.$index,scope.row)">删除</el-button>
							</template>
						</el-table-column>
					</el-table>
				</el-collapse-item>
			</el-collapse>
		</el-col>
	</el-row>

</div>