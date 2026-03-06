<template>
  <el-dialog
    v-model="visible" :title="$t('common.' + formType)" width="50%" :close-on-click-modal="false"
    :close-on-press-escape="false" @close="handleClose"
  >
    <div class="hab-form-box">
      <el-form ref="formRef" :model="form" label-position="right" :rules="rules" label-width="auto">
        <el-container>
          <el-aside width="50%">
            <el-form-item :label="$t('business.sysDataFilter.columns.name') + ':'" prop="name">
              <el-input v-model="form.name" :clearable="true" placeholder="please enter name" :disabled="isView" />
            </el-form-item>
            <el-form-item :label="$t('business.sysDataFilter.columns.sql') + ':'" prop="sql">
              <el-input v-model="form.sql" type="textarea" autosize :disabled="isView" clearable />
              <el-button type="primary" size="default" @click="generateSql">
                生成
              </el-button>
            </el-form-item>
            <el-form-item :label="$t('business.sysDataFilter.columns.note') + ':'" prop="note">
              <el-input v-model="form.note" type="textarea" autosize :disabled="isView" clearable />
            </el-form-item>
            <el-form-item :label="$t('business.sysDataFilter.columns.label') + ':'" prop="label">
              <!-- <el-input v-model="form.label" type="textarea" autosize :disabled="isView" clearable /> -->
              <el-select v-model="form.label" placeholder="Please enter a keyword" style="width: 240px">
                <el-option v-for="item in form.columns" :key="item.columnName" :label="item.label || item.columnName" :value="item.columnName" />
              </el-select>
            </el-form-item>
            <el-form-item :label="$t('business.sysDataFilter.columns.key') + ':'" prop="key">
              <!-- <el-input v-model="form.key" type="textarea" autosize :disabled="isView" clearable /> -->
              <el-select v-model="form.key" placeholder="Please enter a keyword" style="width: 240px">
                <el-option v-for="item in form.columns" :key="item.columnName" :label="item.label || item.columnName" :value="item.columnName" />
              </el-select>
            </el-form-item>
            <el-form-item :label="$t('business.sysDataFilter.columns.value') + ':'" prop="value">
              <!-- <el-input v-model="form.value" type="textarea" autosize :disabled="isView" clearable /> -->
              <el-select v-model="form.value" placeholder="Please enter a keyword" style="width: 240px">
                <el-option v-for="item in form.columns" :key="item.columnName" :label="item.label || item.columnName" :value="item.columnName" />
              </el-select>
            </el-form-item>
          </el-aside>
          <el-container>
            <el-table :data="form.columns" border stripe max-height="500">
              <el-table-column type="index" width="50" />
              <el-table-column prop="columnName" label="Name" width="100" />
              <el-table-column prop="label" label="Label" width="100">
                <template #default="scope">
                  <el-input v-model="scope.row.label" />
                </template>
              </el-table-column>
              <el-table-column prop="isShow" label="Show" width="100">
                <template #default="scope">
                  <el-switch v-model="scope.row.isShow" />
                </template>
              </el-table-column>
              <el-table-column prop="sort" label="Sort" width="100">
                <template #default="scope">
                  <el-input v-model.number="scope.row.sort" />
                </template>
              </el-table-column>
              <el-table-column prop="filter" label="Filter" width="100">
                <template #default="scope">
                  <el-switch v-model="scope.row.filter" />
                </template>
              </el-table-column>
            </el-table>
          </el-container>
        </el-container>
      </el-form>
    </div>
    <template #footer>
      <div class="dialog-footer">
        <el-button type="danger" @click="filterPopRef.open(form.ID)">预览</el-button>
        <el-button @click="handleClose">
          {{ $t('common.cancel') }}
        </el-button>
        <el-button v-if="!isView" type="primary" :loading="loading" @click="handleSubmit">
          {{ $t('common.confirm') }}
        </el-button>
      </div>
    </template>
  </el-dialog>
  <FilterPop ref="filterPopRef" :confirmed="handleFilterConfirmed" />
</template>

<script setup>

import {
  createSysDataFilter,
  updateSysDataFilter,
  findSysDataFilter,
  executeSql,
} from '@/api/system/sysDataFilter'
import FilterPop from '@/components/filterPop/index.vue'
import { ElMessage } from 'element-plus'
import { ref, computed } from 'vue'
import { convertDate, convertDateTime } from '@/utils/time'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const props = defineProps({
  callback: {
    type: Function,
    default: null
  }
})

defineOptions({
  name: 'SysDataFilterForm'
})

// 计算属性
const isView = computed(() => formType.value === 'view')
// 状态
const visible = ref(false)
// 加载状态
const loading = ref(false)
// 表单引用
const formRef = ref()
// 表单数据
const form = ref(getInitialForm())
// 表单类型
const formType = ref('create')
const filterPopRef = ref(null)
const columns = ref([{
  'columnName': 'id',
  'label': '',
  'key': false,
  'isShow': false,
  'filter': false
}])
// 表单校验规则
const rules = {
}

// 获取初始表单数据
function getInitialForm() {
  return {
    name: '',
    sql: '',
    columns: [],
    orderBy: '',
    note: '',
  }
}

// 打开弹窗
async function openDialog(type, data) {
  visible.value = true
  formType.value = type
  if (formType.value === 'create') {
    form.value = getInitialForm()
  } else if (formType.value === 'copy') {
    const copyData = { ...data }
    delete copyData.ID
    delete copyData.createdAt
    delete copyData.updatedAt
    delete copyData.deletedAt
    form.value = copyData
  } else {
    try {
      const res = await findSysDataFilter({ ID: data.ID })
      if (res.code === 0) {
        form.value = res.data
      } else {
        ElMessage.error('Failed to obtain data')
      }
    } catch (error) {
      console.error('Failed to retrieve data:', error)
      ElMessage.error('Failed to obtain data')
    }
  }
}

const itemFormat = (item) => {
  return item
}

// 提交表单
async function handleSubmit() {
  if (!formRef.value) return
  try {
    await formRef.value.validate()
    loading.value = true
    let api = null
    switch (formType.value) {
      case 'create':
        api = createSysDataFilter
        break
      case 'edit':
        api = updateSysDataFilter
        break
      case 'view':
        api = null
        break
      case 'copy':
        api = createSysDataFilter
        break
    }

    if (api) {
      const res = await api(itemFormat(form.value))
      if (res.code === 0) {
        ElMessage.success('operation is successful')
        props.onSuccess?.(form.value)
        handleClose()
      } else {
        ElMessage.error(res.message || 'operation failed')
      }
    }
  } catch (error) {
    console.error('Submission failed:', error)
    ElMessage.error('Submission failed')
  } finally {
    if (props.callback) {
      props.callback()
    }
    loading.value = false
  }
}

// 关闭弹窗
function handleClose() {
  visible.value = false
  form.value = getInitialForm()
  formRef.value?.resetFields()
}

// 生成SQL
function generateSql() {
  executeSql({
    sql: form.value.sql
  }).then((res) => {
    if (res.code === 0) {
      const newColumns = res.data
      const existingColumns = form.value.columns || []

      // 创建映射以便快速查找现有列
      const existingMap = new Map()
      existingColumns.forEach(col => {
        existingMap.set(col.columnName, col)
      })

      // 处理新数据：保留现有属性，添加新列
      const updatedColumns = newColumns.map(newCol => {
        const existing = existingMap.get(newCol.columnName)
        if (existing) {
          // 如果列已存在，保留原有的label、isShow、sort、filter等属性
          return {
            ...newCol,
            label: existing.label || newCol.label || '',
            isShow: existing.isShow !== undefined ? existing.isShow : false,
            sort: existing.sort !== undefined ? existing.sort : 0,
            filter: existing.filter !== undefined ? existing.filter : false
          }
        } else {
          // 新列，使用默认值
          return {
            ...newCol,
            label: newCol.label || '',
            isShow: false,
            sort: 0,
            filter: false
          }
        }
      })

      form.value.columns = updatedColumns
    }
  })
}

// 过滤确认
function handleFilterConfirmed(data) {
  console.log('handleFilterConfirmed', data)
}

// 暴露方法
defineExpose({
  openDialog
})
</script>

<style lang="scss">
.filter-pop-container {
  padding: 20px;
}

:root {
  --el-table-current-row-bg-color: #f0f0f0; // 这里可以设置您想要的颜色值
}

.el-table__body .el-table__row.current-row {
  background-color: var(--el-table-current-row-bg-color) !important;
}
</style>
