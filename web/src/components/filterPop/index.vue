<template>
  <div>
    <el-dialog
      v-model="visible"
      :title="title"
      width="60%"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      @close="handleClose"
    >
      <div class="filter-pop-container">
        <el-form :model="data">
          <el-form-item label="">
            <el-input v-model="filters[0]" style="width: 100px;" />
            <el-input v-model="filters[1]" style="width: 100px;" />
            <el-input v-model="filters[2]" style="width: 100px;" />
            <el-button type="primary" @click="handleFilter">
              过滤
            </el-button>
          </el-form-item>
        </el-form>
        <el-table
          ref="tableRef"
          style="--el-table-current-row-bg-color: #262727"
          highlight-current-row
          :data="tableData"
          border
          max-height="500"
          @selection-change="handleSelectionChange"
          @row-click="handleRowClick"
        >
          <el-table-column
            v-if="multiple"
            type="selection"
            width="55"
          />
          <el-table-column
            v-for="column in data.columns"
            :key="column.columnName"
            :prop="column.columnName"
            :label="column.label"
            width="100"
          />
        </el-table>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="handleClose">
            {{ $t('common.cancel') }}
          </el-button>
          <el-button type="primary" @click="handleConfirm">
            {{ $t('common.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { filterData, findSysDataFilter } from '@/api/system/sysDataFilter'
import { ref, defineProps } from 'vue'
const props = defineProps({
  filterId: {
    type: Number,
    default: 0
  }
})

const visible = ref(false)
const tableRef = ref(null)
const tableData = ref([])
const selectedData = ref([])
const currentData = ref([])
const data = ref({})
const columns = ref([])
const filters = ref([])
const title = ref('')
const multiple = ref(false)
const confirmed = ref(null)
// 打开弹窗
const open = (id, mpe = false, confirmFunc = null) => {
  filters.value = ['', '', '']
  multiple.value = mpe
  if (confirmFunc) {
    confirmed.value = confirmFunc
  }
  findSysDataFilter({ ID: id }).then(res => {
    if (res.code === 0) {
      title.value = res.data.name
      data.value = res.data
      console.log(data.value.columns)
      columns.value = data.value.columns
    }
  }
  )

  visible.value = true
  filterData({ filters: [], id: id }).then(res => {
    if (res.code === 0) {
      tableData.value = res.data
    }
  })
}

// 关闭弹窗
const handleClose = () => {
  visible.value = false
  selectedData.value = []
  // emit('close')
}

// 选择变化
const handleSelectionChange = (selection) => {
  if (props.multiple) {
    selectedData.value = selection
  } else {
    selectedData.value = selection[0]
  }
}

// 过滤
const handleFilter = () => {
  filterData({ filters: filters.value, id: props.filterId }).then(res => {
    if (res.code === 0) {
      tableData.value = res.data
    }
  })
}

// 行点击
const handleRowClick = (row) => {
  currentData.value = row
}

// 确认选择
const handleConfirm = () => {
  if (confirmed.value) {
    if (props.multiple) {
      confirmed.value(selectedData.value)
    } else {
      confirmed.value(currentData.value)
    }
  }
  tableRef.value.clearSelection()
  visible.value = false
}

defineExpose({
  open
})
</script>

<style lang="scss">
.filter-pop-container {
  padding: 20px;
}
</style>
