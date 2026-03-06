<template>
  <div>
    <el-table
      border
      style="width: 100%"
      tooltip-effect="dark"
      :data="props.tableData"
      :row-key="props.rowKey"
      :show-summary="props.columns?.filter(column => column.isSum).length > 0"
      :summary-method="getSummaries"
      :height="tableHight"
      @selection-change="handleSelectionChange"
      @cell-dblclick="(row,col,cell)=>{
        copyToClipboard(row[col.property])
      }"
    >
      <el-table-column
        v-if="props.showSelection"
        type="selection"
        width="55"
      />
      <el-table-column
        v-for="column in props.columns?.filter(column => column.isShow)"
        :key="column.columnName"
        show-overflow-tooltip
        :label="$t('business.'+props.structName+'.' + column.jsonName)"
        :prop="column.jsonName"
        :fixed="column.fixed === 'left' ? 'left' : column.fixed === 'right' ? 'right' : null"
        :width="column.with < 120 ? '120px' : column.with"
      >
        <template #header="">
          <ThreeColumnLayout
            :section-gap="3"
            :item-gap="10"
          >
            <template #left />
            <template #center>
              <el-tooltip
                placement="top"
                :show-after="500"
              >
                <template #content>
                  <span>{{ column.jsonName }}</span>
                  <br>
                  <span>{{ column.note ? column.note : '' }}</span>
                </template>
                <span>{{ $t('business.'+props.structName+'.columns.' + column.jsonName) }}</span>
              </el-tooltip>
            </template>
            <template #right>
              <el-icon
                v-if="column.sortable"
                style="width: 14px; flex-shrink: 0; margin-left: 4px;"
                @click.stop="handleSort(column.columnName)"
              >
                <Sort v-if="!sortInfo[column.columnName]" />
                <CaretTop v-if="sortInfo[column.columnName]?.type === 'asc'" />
                <CaretBottom v-if="sortInfo[column.columnName]?.type === 'desc'" />
              </el-icon>
            </template>
          </ThreeColumnLayout>
        </template>
        <template #default="scope">
          <span v-if="column.type === 'boolean'">{{ scope.row[column.jsonName] ? '是' : '否' }}</span>
          <span v-else-if="column.type === 'date'">{{ scope.row[column.jsonName]?.split(" ")[0] }}</span>
          <span v-else-if="column.type === 'datetime'">{{ scope.row[column.jsonName] }}</span>
          <span v-else-if="column.type === 'uintDate'">{{ $serverDate(scope.row[column.jsonName],'date') }}</span>
          <span v-else-if="column.type === 'tag'">
            <el-tag
              v-if="column.tagData[scope.row[column.jsonName]]"
              :type="column.tagData[scope.row[column.jsonName]].type"
              size="default"
              :effect="column.tagData[scope.row[column.jsonName]].effect"
            >
              {{ column.tagData[scope.row[column.jsonName]].label }}
            </el-tag>
            <span v-else>{{ scope.row[column.jsonName] }}</span>
          </span>
          <span v-else-if="column.type === 'enum' || column.type === 'protoEnum'">
            <el-tag
              v-if="column.enum.indexOf(scope.row[column.jsonName]+'')%5===0"
              size="default"
              effect="dark"
              type="primary"
            >
              {{ $t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+(scope.row[column.jsonName]+'').replaceAll('.', '')) }}
            </el-tag>
            <el-tag
              v-else-if="column.enum.indexOf(scope.row[column.jsonName]+'')%5===1"
              size="default"
              effect="dark"
              type="success"
            >
              {{ $t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+(scope.row[column.jsonName]+'').replaceAll('.', '')) }}
            </el-tag>
            <el-tag
              v-else-if="column.enum.indexOf(scope.row[column.jsonName]+'')%5===2"
              size="default"
              effect="dark"
              type="warning"
            >
              {{ $t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+(scope.row[column.jsonName]+'').replaceAll('.', '')) }}
            </el-tag>
            <el-tag
              v-else-if="column.enum.indexOf(scope.row[column.jsonName]+'')%5===3"
              size="default"
              effect="dark"
              type="danger"
            >
              {{ $t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+(scope.row[column.jsonName]+'').replaceAll('.', '')) }}
            </el-tag>
            <el-tag
              v-else-if="column.enum.indexOf(scope.row[column.jsonName]+'')%5===4"
              size="default"
              effect="dark"
              type="info"
            >
              {{ $t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+(scope.row[column.jsonName]+'').replaceAll('.', '')) }}
            </el-tag>
            <el-tag
              v-else
              size="default"
              effect="plain"
              type="info"
            >
              {{ (scope.row[column.jsonName]+'').replaceAll('.', '') }}
            </el-tag>
          </span>
          <SelectTool v-else-if="column.type === 'table'" v-model:value="scope.row[column.jsonName]" :selectable="false" :disabled="true" :filter-id="column.filterId" />
          <el-image v-else-if="column.type === 'image'" style="width: 100px; height: 100px" :src="'data:image/png;base64,'+scope.row[column.jsonName]" fit="fill" />
          <span v-else>{{ format(scope.row[column.jsonName], column.format) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        v-if="$slots.operate"
        align="left"
        :label="$t('common.operate')"
        fixed="right"
        :min-width="appStore.operateMinWith"
      >
        <template #default="scope">
          <slot
            name="operate"
            :row="scope.row"
          />
        </template>
      </el-table-column>
    </el-table>
    <div class="gva-pagination">
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="pageInfo.page"
        :page-size="pageInfo.pageSize"
        :page-sizes="[5,10, 30, 50, 100]"
        :total="props.total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
    <!-- <ColOptions :columns="props.columns" /> -->
  </div>
</template>

<script setup>
import { defineProps, ref, defineExpose, onMounted, computed } from 'vue'
import { Sort, CaretTop, CaretBottom } from '@element-plus/icons-vue'
import queryFrom from './queryFrom.vue'
import { useAppStore } from '@/pinia'
import ThreeColumnLayout from '@/components/ThreeColumnLayout/index.vue'
import { formatString, formatNumber } from '@/utils/format'
import SelectTool from '@/components/filterPop/selectTool.vue'
import { copyToClipboard } from '@/utils/clipboard'
import { ElMessage } from 'element-plus'
// import ColOptions from './coloptions.vue'
const format = (value, format) => {
  if (typeof value === 'number') {
    return formatNumber(value, format)
  }
  if (typeof value === 'string') {
    return formatString(value, format)
  }
  return value
}

const appStore = useAppStore()

const props = defineProps({
  structName: {
    type: String,
    required: true
  },
  tableData: {
    type: Array,
    required: true
  },
  columns: {
    type: Array,
    required: true
  },
  rowKey: {
    type: String,
    required: true,
    default: 'ID'
  },
  showSelection: {
    type: Boolean,
    default: false
  },
  total: {
    type: Number,
    required: true,
    default: 0
  },
  getTableData: {
    type: Function,
    required: true
  },
  defaultSort: {
    type: Object,
    default: () => ({
    })
  },
  sumData: {
    type: Object,
    default: () => ({
    })
  },
  height: {
    type: [String, Number],
    required: false,
    default: ''
  },
})

const pageInfo = ref({
  page: 1,
  pageSize: 10,
  total: 0
})

const searchInfo = ref({})
const sortInfo = ref({ ...props.defaultSort })
const selectedData = ref([])

// const getLocalStorage = () => {
//   const tablePageSize = localStorage.getItem('tablePageSize')
//   if (tablePageSize) {
//     pageInfo.value.pageSize = Number(tablePageSize)
//   }
// }
// getLocalStorage()

const tableHight = computed(() => {
  if (props.height) {
    return props.height
  }
  if (appStore.config.table_hight_enable) {
    return appStore.config.table_hight
  }
  return 'auto'
})

const handleSort = (key) => {
  if (sortInfo.value[key]) {
    if (sortInfo.value[key].type === 'asc') {
      sortInfo.value[key].type = 'desc'
    } else {
      delete sortInfo.value[key]
    }
  } else {
    sortInfo.value[key] = {
      index: Object.keys(sortInfo.value).length,
      type: 'asc'
    }
  }
  props.getTableData()
}

// 分页
const handleSizeChange = (val) => {
  localStorage.setItem('tablePageSize', val)
  pageInfo.value.pageSize = val
  props.getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  pageInfo.value.page = val
  props.getTableData()
}

// 获取选中的数据
const getSelectedData = () => {
  return selectedData.value
}

const clearSearch = () => {
  searchInfo.value = {}
  sortInfo.value = { ...props.defaultSort }
  // props.getTableData()
}

const getQueryInfo = () => {
  return {
    pageInfo: pageInfo.value,
    searchInfo: searchInfo.value,
    sortInfo: sortInfo.value
  }
}

const handleSelectionChange = (selection) => {
  console.log('selection', selection)
  selectedData.value = selection
}

// 获取合计数据
const getSummaries = (param) => {
  const columns = props.columns.filter(column => column.isSum)
  if (columns.length === 0) {
    return []
  }
  const sumRes = {}
  const sumWait = []
  const sums = []
  param.columns.forEach((paramColumn, index) => {
    if (index === 0) {
      sums[index] = 'Sum'
      return
    }
    const column = columns.find(item => item.jsonName === paramColumn.property)
    if (!column) {
      return
    }

    const sumDatafn = props.sumData[column.jsonName]
    // 是否有自定义的求和方法 如果有 则等待计算完成后再赋值
    if (sumDatafn) {
      sumWait.push({ index, column: column })
      sums[index] = null
      return
    }

    const values = param.data.map((item) => Number(item[paramColumn.property]))
    if (values.every((value) => Number.isNaN(value))) {
      sums[index] = 'N/A'
      return
    }

    const sum = values.reduce((prev, curr) => { return prev + curr }, 0)
    // 保存格式化前的数据到结果
    sumRes[column.jsonName] = sum
    sums[index] = format(sum, column.format)
  })

  // 最后在获取自定义求和方法的结果
  sumWait.forEach((v) => {
    const sumDatafn = props.sumData[v.column.jsonName]
    const params = {
      res: sumRes,
      data: param.data,
      column: v.column
    }
    const sum = sumDatafn(params)
    sums[v.index] = format(sum, v.column.format)
  })

  return sums
}

onMounted(() => {
  // console.log('props.columns', props.columns)
})
// 暴露给父组件的方法
defineExpose({
  clearSearch,
  getSelectedData,
  getQueryInfo,
})
</script>
