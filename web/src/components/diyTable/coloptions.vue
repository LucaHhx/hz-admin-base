<template>
  <el-button
    type="primary" icon="Setting" @click="console.log(props.columns);
                                          dialogVisible = true"
  />
  <el-dialog
    v-model="dialogVisible" :title="$t('common.columnConfig')" width="1200px" max-height="500px"
    style="overflow-y: auto;" draggable
  >
    <el-table :data="props.columns" border style="width: 100%" :auto-resize="true" :max-height="500">
      <el-table-column :label="$t('common.columnName')" prop="jsonName" width="120px" fixed="left">
        <template #default="scope">
          <span>{{ $t('business.' + props.structName + '.columns.' + scope.row.jsonName) }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('common.columnName')" prop="jsonName" width="140px" fixed="left" />
      <el-table-column :label="$t('business.sysTableColumns.columns.isAddSearch')" width="80px">
        <template #default="scope">
          <el-switch v-model="scope.row.isAddSearch" size="small" />
        </template>
      </el-table-column>

      <el-table-column :label="$t('business.sysTableColumns.columns.isShow')" width="120px">
        <template #default="scope">
          <el-switch v-model="scope.row.isShow" size="small" />
        </template>
      </el-table-column>

      <el-table-column :label="$t('business.sysTableColumns.columns.isSum')" width="120px">
        <template #default="scope">
          <el-switch v-model="scope.row.isSum" size="small" />
        </template>
      </el-table-column>

      <el-table-column :label="$t('business.sysTableColumns.columns.searchWidth')" width="120px">
        <template #default="scope">
          <el-input-number v-model="scope.row.searchWidth" :min="120" :step="10" size="small" style="width: 100%" />
        </template>
      </el-table-column>

      <el-table-column :label="$t('common.width')" width="120px">
        <template #default="scope">
          <el-input-number v-model="scope.row.with" :min="100" :step="10" size="small" style="width: 100%" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('common.fixed')" width="120px">
        <template #default="scope">
          <el-select v-model="scope.row.fixed" size="small" style="width: 100%">
            <el-option :label="$t('common.notFixed')" value="none" />
            <el-option :label="$t('common.fixedLeft')" value="left" />
            <el-option :label="$t('common.fixedRight')" value="right" />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column :label="$t('common.sort')" width="120px">
        <template #default="scope">
          <el-input-number v-model="scope.row.sort" :min="0" :step="1" size="small" style="width: 100%" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.format')" width="120px">
        <template #default="scope">
          <el-input v-model="scope.row.format" size="small" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.type')" width="120px">
        <template #default="scope">
          <el-select v-model="scope.row.type" size="small" style="width: 100%">
            <el-option v-for="item in typeOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </template>
      </el-table-column>

      <el-table-column :label="$t('business.sysTableColumns.columns.filterId')" width="130px">
        <template #default="scope">
          <SelectTool v-if="scope.row.type === 'table'" v-model:value="scope.row.filterId" :filter-id="1" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.editorFilterId')" width="130px">
        <template #default="scope">
          <SelectTool v-if="scope.row.type === 'table'" v-model:value="scope.row.editorFilterId" :filter-id="1" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.sortable')" width="80px">
        <template #default="scope">
          <el-switch v-model="scope.row.sortable" size="small" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.defaultFilter')" width="120px">
        <template #default="scope">
          <el-select v-model="scope.row.defaultFilter" size="small" style="width: 100%">
            <el-option v-for="item in filterList" :key="item" :label="item" :value="item" />
          </el-select>
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.note')" width="120px">
        <template #default="scope">
          <el-input v-model="scope.row.note" size="small" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.enum')" width="500px">
        <template #default="scope">
          <el-select v-model="scope.row.enum" size="small" multiple allow-create filterable style="width: 100%" />
        </template>
      </el-table-column>
      <el-table-column :label="$t('business.sysTableColumns.columns.isAdditional')" width="120px">
        <template #default="scope">
          <el-switch v-model="scope.row.isAdditional" size="small" />
        </template>
      </el-table-column>
    </el-table>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">{{ $t('common.cancel') }}</el-button>
        <el-button style="margin-left: 10px;" type="primary" @click="handleConfirm">{{ $t('common.confirm') }}</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { updateSysTableColumnsInfo } from '@/api/system/sysTableColumns'
import { ref, defineProps } from 'vue'
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
    label: '二进制',
    value: 'binary'
  },
  {
    label: 'protoEnum',
    value: 'protoEnum'
  },
  {
    label: 'image',
    value: 'image'
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
