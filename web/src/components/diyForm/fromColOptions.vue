<template>
  <el-button type="primary" icon="Setting" @click="dialogVisible = true" />
  <el-dialog
    v-model="dialogVisible" :title="$t('common.columnConfig')" width="900px" max-height="500px"
    style="overflow-y: auto;"
    :append-to-body="true"
    :close-on-click-modal="true"
    draggable
  >
    <el-table :data="columns" border style="width: 100%" :auto-resize="true" :max-height="500">
      <el-table-column :label="$t('common.columnName')" prop="jsonName" width="120px" fixed="left">
        <template #default="scope">
          <span>{{ $t('business.' + props.structName + '.columns.' + scope.row.jsonName) }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('common.columnName')" prop="jsonName" width="120px" fixed="left">
        <template #default="scope">
          <span>{{ scope.row.jsonName }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.formWith')" prop="formWith" width="120px">
        <template #default="scope">
          <el-input-number v-model="scope.row.formWith" :min="30" :max="100" :step="5" size="small" style="width: 100%" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.formDisabled')" prop="formDisabled" width="120px">
        <template #default="scope">
          <el-switch v-model="scope.row.formDisabled" size="small" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.formHidden')" prop="formHidden" width="120px">
        <template #default="scope">
          <el-switch v-model="scope.row.formHidden" size="small" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.formMust')" prop="formMust" width="120px">
        <template #default="scope">
          <el-switch v-model="scope.row.formMust" size="small" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.formOrder')" prop="formOrder" width="120px">
        <template #default="scope">
          <el-input-number v-model="scope.row.formOrder" :step="1" size="small" style="width: 100%" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.type')" width="120px">
        <template #default="scope">
          <el-select v-model="scope.row.type" size="small" style="width: 100%">
            <el-option v-for="item in typeOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </template>
      </el-table-column>

      <!-- <el-table-column :label="$t('business.sysTableColumns.columns.filterId')" width="130px">
        <template #default="scope">
          <SelectTool v-if="scope.row.type === 'table'" v-model:value="scope.row.filterId" :filter-id="1" />
        </template>
      </el-table-column> -->

      <el-table-column :label="$t('business.sysTableColumns.columns.editorFilterId')" width="250px">
        <template #default="scope">
          <SelectTool v-if="scope.row.type === 'table'" v-model:value="scope.row.editorFilterId" :filter-id="1" />
        </template>
      </el-table-column>
    </el-table>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" style="margin-left: 10px;" @click="handleConfirm">{{ $t('common.confirm') }}</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { updateSysTableColumnsInfo } from '@/api/system/sysTableColumns'
import { ref, defineProps, computed } from 'vue'
import { ElMessage } from 'element-plus'
import SelectTool from '@/components/filterPop/selectTool.vue'
const props = defineProps({
  columns: {
    type: Array,
    required: true
  },
  structName: {
    type: String,
    required: true
  },
})
const columns = computed(() => props.columns.filter(item => true).sort((a, b) => {
  return a.ID - b.ID
}))

const dialogVisible = ref(false)
const filterList = ['=', '>', '>=', '<', '<=', 'like', 'in', 'between']
const typeOptions = [
  {
    label: '字符串',
    value: 'string'
  },
  {
    label: '富文本',
    value: 'richtext'
  },
  {
    label: '布尔值',
    value: 'bool'
  },
  {
    label: '枚举',
    value: 'enum'
  },
  {
    label: '关联表',
    value: 'table'
  },
  {
    label: '整型',
    value: 'int32'
  },
  {
    label: '大整数',
    value: 'int64'
  },
  {
    label: '浮点型',
    value: 'float64'
  },
  {
    label: '整型',
    value: 'int'
  },
  {
    label: '时间',
    value: 'datetime'
  },
  {
    label: '日期',
    value: 'date'
  },
  {
    label: '无符号日期',
    value: 'uintDate'
  },
  {
    label: 'JSON',
    value: 'json'
  },
  {
    label: 'JSON字符',
    value: 'jsonStr'
  },
  {
    label: '数组',
    value: 'array'
  },
  {
    label: '对象',
    value: 'object'
  },
]

const handleConfirm = () => {
  updateSysTableColumnsInfo(props.columns).then(res => {
    if (res.code === 0) {
      ElMessage.success('更新成功')
    } else {
      ElMessage.error(res.msg)
    }
  })
  dialogVisible.value = false
}
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}

:deep(.el-table .el-table__cell) {
  padding: 4px 0;
}

:deep(.el-input-number),
:deep(.el-select),
:deep(.el-input) {
  width: 100%;
}

:deep(.el-table) {
  width: 100% !important;
}

:deep(.el-table__body) {
  width: 100% !important;
}

:deep(.el-table__header) {
  width: 100% !important;
}
</style>
