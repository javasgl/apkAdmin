<div id="channel">
	<el-row>
		<el-col :span="4">
			<el-menu mode="horizontal" theme="dark">
				<el-menu-item index="-">
					apksys
				</el-menu-item>
			</el-menu>
		</el-col>
		<el-col :span="20">
			<el-menu mode="horizontal" theme="dark" default-active="1">
				<el-menu-item index="1">packing center</el-menu-item>
				<el-menu-item index="2">channel manage</el-menu-item>
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
								{{new Date()}}
							</div>
						</el-col>
						<el-col :span="8">
							<el-form label-width="80px">
								<el-form-item label="版本号">
									<el-input :value="22" placeholder="eg: 8.2.6"></el-input>
								</el-form-item>
								<el-form-item label="渠道">
									<el-checkbox :indeterminate="isIndeterminate" v-model="checkAll" @change="handleCheckAllChange">全选</el-checkbox>
									<el-checkbox-group v-model="form.checkedChannels" @change="handleCheckedChannelsChange">
										<el-checkbox v-for="channel in channels" :label="channel" :key="channel">{{channel}}</el-checkbox>
									</el-checkbox-group>
								</el-form-item>
								<el-form-item label="">
									<el-button type="primary" size="small" icon="menu">start packing ! </el-button>
								</el-form-item>
							</el-form>
						</el-col>
						<el-col :span="8"></el-col>
					</el-row>
				</el-collapse-item>
				<el-collapse-item title="history" name="2">
					<el-table>
						<el-table-column prop="date" label="日期"></el-table-column>
						<el-table-column prop="date" label="操作人"></el-table-column>
						<el-table-column prop="date" label="版本"></el-table-column>
						<el-table-column prop="date" label="进度"></el-table-column>
						<el-table-column prop="date" label="下载"></el-table-column>
					</el-table>
				</el-collapse-item>
			</el-collapse>
		</el-col>
	</el-row>
</div>
