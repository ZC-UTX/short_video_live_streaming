
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="标题" prop="title">
  <el-input v-model="searchInfo.title" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="描述" prop="desc">
  <el-input v-model="searchInfo.desc" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="选择音乐" prop="musicId">
  <el-input v-model.number="searchInfo.musicId" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="作品类型" prop="workType">
  <el-select v-model="searchInfo.workType" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.workType=undefined}">
    <el-option v-for="(item,key) in worksTypeOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="审核状态" prop="checkStatus">
  <el-select v-model="searchInfo.checkStatus" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.checkStatus=undefined}">
    <el-option v-for="(item,key) in ReviewTypeOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="审核人" prop="checkUser">
  <el-input v-model.number="searchInfo.checkUser" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="IP地址" prop="ipAddress">
  <el-input v-model="searchInfo.ipAddress" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="作品权限" prop="workPermission">
  <el-select v-model="searchInfo.workPermission" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.workPermission=undefined}">
    <el-option v-for="(item,key) in WorkRightsOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="video_VideoWorks" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="video_VideoWorks" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="video_VideoWorks" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="标题" prop="title" width="120" />

            <el-table-column align="left" label="描述" prop="desc" width="120" />

            <el-table-column align="left" label="选择音乐" prop="musicId" width="120" />

            <el-table-column align="left" label="作品类型" prop="workType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.workType,worksTypeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="审核状态" prop="checkStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.checkStatus,ReviewTypeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="审核人" prop="checkUser" width="120" />

            <el-table-column align="left" label="IP地址" prop="ipAddress" width="120" />

            <el-table-column align="left" label="作品权限" prop="workPermission" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.workPermission,WorkRightsOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="喜欢数量" prop="likeCount" width="120" />

            <el-table-column align="left" label="评论数" prop="commentCount" width="120" />

            <el-table-column align="left" label="分享数" prop="shareCount" width="120" />

            <el-table-column align="left" label="收藏数" prop="collectCount" width="120" />

            <el-table-column align="left" label="浏览量" prop="browseCount" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateVideoWorksFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>


          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-row>
              <el-form-item label="标题:" prop="title">
                <el-input v-model="formData.title" :clearable="true" placeholder="请输入标题"/>
              </el-form-item>
              &emsp;
              &emsp;
              <el-form-item label="描述:" prop="desc">
                <el-input v-model="formData.desc" :clearable="true" placeholder="请输入描述" />
              </el-form-item>
              &emsp;
              &emsp;
              <el-form-item label="IP地址:" prop="ipAddress">
                <el-input v-model="formData.ipAddress" :clearable="true" placeholder="请输入IP地址" />
              </el-form-item>
            </el-row>

            <el-row>
              <el-col :span="10">
                <el-form-item label="作品类型:" prop="workType">
                  <el-select v-model="formData.workType" placeholder="请选择作品类型" style="width:100%" filterable :clearable="true">
                    <el-option v-for="(item,key) in worksTypeOptions" :key="key" :label="item.label" :value="item.value" />
                  </el-select>
                </el-form-item>
              </el-col>
              &emsp;
              &emsp;
              <el-col :span="10">
                <el-form-item label="作品权限:" prop="workPermission">
                  <el-select v-model="formData.workPermission" placeholder="请选择作品权限" style="width:100%" filterable :clearable="true">
                    <el-option v-for="(item,key) in WorkRightsOptions" :key="key" :label="item.label" :value="item.value" />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>

            <el-form-item label="选择音乐:" prop="musicId">
    <el-input v-model.number="formData.musicId" :clearable="true" placeholder="请输入选择音乐" />
</el-form-item>

<!--            <el-form-item label="审核状态:" prop="checkStatus">-->
<!--    <el-select v-model="formData.checkStatus" placeholder="请选择审核状态" style="width:100%" filterable :clearable="true">-->
<!--        <el-option v-for="(item,key) in ReviewTypeOptions" :key="key" :label="item.label" :value="item.value" />-->
<!--    </el-select>-->
<!--</el-form-item>-->
<!--            <el-form-item label="审核人:" prop="checkUser">-->
<!--    <el-input v-model.number="formData.checkUser" :clearable="true" placeholder="请输入审核人" />-->
<!--</el-form-item>-->


<!--            <el-form-item label="喜欢数量:" prop="likeCount">-->
<!--    <el-input v-model.number="formData.likeCount" :clearable="true" placeholder="请输入喜欢数量" />-->
<!--</el-form-item>-->
<!--            <el-form-item label="评论数:" prop="commentCount">-->
<!--    <el-input v-model.number="formData.commentCount" :clearable="true" placeholder="请输入评论数" />-->
<!--</el-form-item>-->
<!--            <el-form-item label="分享数:" prop="shareCount">-->
<!--    <el-input v-model.number="formData.shareCount" :clearable="true" placeholder="请输入分享数" />-->
<!--</el-form-item>-->
<!--            <el-form-item label="收藏数:" prop="collectCount">-->
<!--    <el-input v-model.number="formData.collectCount" :clearable="true" placeholder="请输入收藏数" />-->
<!--</el-form-item>-->
<!--            <el-form-item label="浏览量:" prop="browseCount">-->
<!--    <el-input v-model.number="formData.browseCount" :clearable="true" placeholder="请输入浏览量" />-->
<!--</el-form-item>-->
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="标题">
    {{ detailFrom.title }}
</el-descriptions-item>
                    <el-descriptions-item label="描述">
    {{ detailFrom.desc }}
</el-descriptions-item>
                    <el-descriptions-item label="选择音乐">
    {{ detailFrom.musicId }}
</el-descriptions-item>
                    <el-descriptions-item label="作品类型">
    {{ detailFrom.workType }}
</el-descriptions-item>
                    <el-descriptions-item label="审核状态">
    {{ detailFrom.checkStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="审核人">
    {{ detailFrom.checkUser }}
</el-descriptions-item>
                    <el-descriptions-item label="IP地址">
    {{ detailFrom.ipAddress }}
</el-descriptions-item>
                    <el-descriptions-item label="作品权限">
    {{ detailFrom.workPermission }}
</el-descriptions-item>
                    <el-descriptions-item label="喜欢数量">
    {{ detailFrom.likeCount }}
</el-descriptions-item>
                    <el-descriptions-item label="评论数">
    {{ detailFrom.commentCount }}
</el-descriptions-item>
                    <el-descriptions-item label="分享数">
    {{ detailFrom.shareCount }}
</el-descriptions-item>
                    <el-descriptions-item label="收藏数">
    {{ detailFrom.collectCount }}
</el-descriptions-item>
                    <el-descriptions-item label="浏览量">
    {{ detailFrom.browseCount }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createVideoWorks,
  deleteVideoWorks,
  deleteVideoWorksByIds,
  updateVideoWorks,
  findVideoWorks,
  getVideoWorksList
} from '@/api/video/videoWorks'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'VideoWorks'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const ReviewTypeOptions = ref([])
const WorkRightsOptions = ref([])
const worksTypeOptions = ref([])
const formData = ref({
            title: '',
            desc: '',
            musicId: undefined,
            workType: '',
            checkStatus: '',
            checkUser: undefined,
            ipAddress: '',
            workPermission: '',
            likeCount: undefined,
            commentCount: undefined,
            shareCount: undefined,
            collectCount: undefined,
            browseCount: undefined,
        })



// 验证规则
const rule = reactive({
               workType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               checkStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               workPermission : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getVideoWorksList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    ReviewTypeOptions.value = await getDictFunc('ReviewType')
    WorkRightsOptions.value = await getDictFunc('WorkRights')
    worksTypeOptions.value = await getDictFunc('worksType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteVideoWorksFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteVideoWorksByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateVideoWorksFunc = async(row) => {
    const res = await findVideoWorks({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteVideoWorksFunc = async (row) => {
    const res = await deleteVideoWorks({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        title: '',
        desc: '',
        musicId: undefined,
        workType: '',
        checkStatus: '',
        checkUser: undefined,
        ipAddress: '',
        workPermission: '',
        likeCount: undefined,
        commentCount: undefined,
        shareCount: undefined,
        collectCount: undefined,
        browseCount: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createVideoWorks(formData.value)
                  break
                case 'update':
                  res = await updateVideoWorks(formData.value)
                  break
                default:
                  res = await createVideoWorks(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findVideoWorks({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
