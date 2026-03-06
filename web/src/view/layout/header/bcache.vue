<template>
  <div>
    <el-dialog v-model="dialogVisible" title="缓存刷新结果" width="500px" :close-on-click-modal="false">
      <div v-loading="loading" element-loading-text="正在刷新缓存...">
        <div v-if="refreshData">
          <h4 style="margin: 0 0 8px 0;">gRPC 服务缓存</h4>
          <el-table :data="refreshData.grpcResults" border size="small" style="margin-bottom: 16px;">
            <el-table-column prop="name" label="服务名称" />
            <el-table-column label="状态" width="80" align="center">
              <template #default="{ row }">
                <el-tag :type="row.success ? 'success' : 'danger'" size="small">
                  {{ row.success ? '成功' : '失败' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="error" label="错误信息">
              <template #default="{ row }">
                <span style="color: #f56c6c;">{{ row.error || '-' }}</span>
              </template>
            </el-table-column>
          </el-table>
          <h4 style="margin: 0 0 8px 0;">服务本地缓存</h4>
          <el-table :data="serviceResultList" border size="small">
            <el-table-column prop="mode" label="模式" />
            <el-table-column label="状态" width="80" align="center">
              <template #default="{ row }">
                <el-tag :type="row.success ? 'success' : 'danger'" size="small">
                  {{ row.success ? '已确认' : '未确认' }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div v-else-if="!loading" style="text-align: center; color: #909399; padding: 20px 0;">
          请求失败，请重试
        </div>
      </div>
      <template #footer>
        <el-button type="primary" :disabled="loading" @click="dialogVisible = false">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { refreshBCache } from '@/api/menu'
import { ElMessageBox } from 'element-plus'

defineOptions({
  name: 'BCache',
})

const dialogVisible = ref(false)
const refreshData = ref(null)
const loading = ref(false)

const serviceResultList = computed(() => {
  if (!refreshData.value?.serviceResults) return []
  return Object.entries(refreshData.value.serviceResults).map(([mode, success]) => ({
    mode,
    success,
  }))
})

const open = () => {
  ElMessageBox.confirm('确认刷新缓存操作?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    refreshData.value = null
    loading.value = true
    dialogVisible.value = true
    try {
      const res = await refreshBCache({})
      if (res.code === 0) {
        refreshData.value = res.data
      }
    } finally {
      loading.value = false
    }
  })
}

defineExpose({ open })
</script>
