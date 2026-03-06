<template>
  <div>
    <el-card shadow="always" :body-style="{ padding: '20px',with: '100%'}">
      <template #header>
        <!-- <el-button type="primary" icon="plus" size="default" @click="addItem" /> -->
        <slot name="buttons" />
        <!-- <el-button v-if="importData" type="primary" icon="plus" size="default" @click="importData" /> -->
        <!-- <el-button v-if="defaultSort" type="primary" size="default" @click="sortItems">排序</el-button> -->
      </template>
      <el-table ref="dataTable" :data="props.list" border stripe style="width: 100%" max-height="500">
        <el-table-column
          v-if="props.isMobile"
          align="left"
          type="index"
          width="30"
        >
          <template #default="scope">
            <el-row :gutter="20">
              <el-button v-if="scope.$index>0" link type="danger" circle @click="()=>{uprow(scope.$index)}">⬆️</el-button>
            </el-row>
            <el-row :gutter="20">
              <el-button
                v-if="scope.$index<props.list.length-1"
                link type="danger" circle @click="()=>{dowrow(scope.$index)}"
              >⬇️</el-button>
            </el-row>
          </template>
        </el-table-column>
        <!-- <el-table-column type="index" width="50" /> -->
        <el-table-column
          v-for="col in columns"
          :key="col.id"
          :prop="col.id"
          :label="col.label"
          :width="col.width??100"
        >
          <template #default="scope">
            <el-input v-if="col.type=='number'" v-model.number="scope.row[col.id]" :disabled="col.disabled" placeholder="" size="default" clearable :width="col.width">
              <template v-if="col.prepend" #prepend>{{ col.prepend }}</template>
              <template v-if="col.append" #append>{{ col.append }}</template>
            </el-input>
            <el-input-number v-else-if="col.type=='float'" v-model.number="scope.row[col.id]" :disabled="col.disabled" placeholder="" size="default" clearable :with="col.with">
              <template v-if="col.prepend" #prepend>{{ col.prepend }}</template>
              <template v-if="col.append" #append>{{ col.append }}</template>
            </el-input-number>
            <el-input v-else v-model="scope.row[col.id]" placeholder="" :disabled="col.disabled" size="default" clearable :with="col.with">
              <template v-if="col.prepend" #prepend>{{ col.prepend }}</template>
              <template v-if="col.append" #append>{{ col.append }}</template>
            </el-input>
          </template>
        </el-table-column>
        <el-table-column v-if="$slots.operation" label="Operation" width="100">
          <template #default="scope">
            <slot name="operation" :index="scope.$index" />
            <!-- <el-button type="danger" icon="delete" @click="delItem(scope.$index)" /> -->
          </template>
        </el-table-column>
      </el-table>
      <!-- card body -->
    </el-card>
  </div>
</template>

<script setup>
import { defineProps, defineEmits, ref, watch, computed } from 'vue'

// 定义 props 和 emits
const props = defineProps({
  list: {
    type: Array,
    default: () => []
  },
  columns: {
    type: Array,
    default: () => []
  },
  addItem: {
    type: Function,
    default: () => {}
  },
  defaultSort: {
    type: Object,
    default: () => {}
  },
  importData: {
    type: Function,
    default: null
  },
  isMobile: {
    type: Boolean,
    default: false
  }
})
const dataTable = ref(null)

// const emit = defineEmits(['update:list'])

// 创建本地响应式变量
// const localList = ref([...props.list])

// 监听 props.list 的变化，并同步到 localList
// watch(
//   () => props.list,
//   (newList) => {
//     localList.value = [...newList]
//   },
//   { deep: true, immediate: true } // `immediate` 确保初始值同步
// )

// // 监听 localList 的变化，并同步到父组件
// watch(
//   () => localList.value,
//   (newList) => {
//     emit('update:list', newList)
//   },
//   { deep: true }
// )

// 方法：删除项
const delItem = (index) => {
  dataTable.value.data.splice(index, 1)
}

// // 方法：添加项
// const addItem = () => {
//   localList.value.push(props.newRow)
// }

const dowrow = (index) => {
  console.log(dataTable.value)
  const dataList = dataTable.value.data
  if (index < dataList.length - 1) {
    const temp = dataList[index]
    dataList[index] = dataList[index + 1]
    dataList[index + 1] = temp
  }
  // dataTable.value.data = dataList
}

const uprow = (index) => {
  const dataList = dataTable.value.data
  if (index > 0) {
    const temp = dataList[index]
    dataList[index] = dataList[index - 1]
    dataList[index - 1] = temp
  }
  // dataTable.value.data = dataList
}

// const sortItems = () => {
//   dataTable.value.sort(props.defaultSort.prop, props.defaultSort.order)
//   dataTable.value.clearSort()
// }
</script>
