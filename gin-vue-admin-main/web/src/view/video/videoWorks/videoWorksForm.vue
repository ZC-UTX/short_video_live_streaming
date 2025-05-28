
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入标题" />
</el-form-item>
        <el-form-item label="描述:" prop="desc">
    <el-input v-model="formData.desc" :clearable="true" placeholder="请输入描述" />
</el-form-item>
        <el-form-item label="选择音乐:" prop="musicId">
    <el-input v-model.number="formData.musicId" :clearable="true" placeholder="请输入选择音乐" />
</el-form-item>
        <el-form-item label="作品类型:" prop="workType">
    <el-select v-model="formData.workType" placeholder="请选择作品类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in worksTypeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="审核状态:" prop="checkStatus">
    <el-select v-model="formData.checkStatus" placeholder="请选择审核状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in ReviewTypeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="审核人:" prop="checkUser">
    <el-input v-model.number="formData.checkUser" :clearable="true" placeholder="请输入审核人" />
</el-form-item>
        <el-form-item label="IP地址:" prop="ipAddress">
    <el-input v-model="formData.ipAddress" :clearable="true" placeholder="请输入IP地址" />
</el-form-item>
        <el-form-item label="作品权限:" prop="workPermission">
    <el-select v-model="formData.workPermission" placeholder="请选择作品权限" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in WorkRightsOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="喜欢数量:" prop="likeCount">
    <el-input v-model.number="formData.likeCount" :clearable="true" placeholder="请输入喜欢数量" />
</el-form-item>
        <el-form-item label="评论数:" prop="commentCount">
    <el-input v-model.number="formData.commentCount" :clearable="true" placeholder="请输入评论数" />
</el-form-item>
        <el-form-item label="分享数:" prop="shareCount">
    <el-input v-model.number="formData.shareCount" :clearable="true" placeholder="请输入分享数" />
</el-form-item>
        <el-form-item label="收藏数:" prop="collectCount">
    <el-input v-model.number="formData.collectCount" :clearable="true" placeholder="请输入收藏数" />
</el-form-item>
        <el-form-item label="浏览量:" prop="browseCount">
    <el-input v-model.number="formData.browseCount" :clearable="true" placeholder="请输入浏览量" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createVideoWorks,
  updateVideoWorks,
  findVideoWorks
} from '@/api/video/videoWorks'

defineOptions({
    name: 'VideoWorksForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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
               }],
               checkStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               workPermission : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVideoWorks({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ReviewTypeOptions.value = await getDictFunc('ReviewType')
    WorkRightsOptions.value = await getDictFunc('WorkRights')
    worksTypeOptions.value = await getDictFunc('worksType')
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
