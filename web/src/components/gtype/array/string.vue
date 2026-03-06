<template>
  <div>
    <el-row v-for="(item, index) in localList" :key="index">
      <el-col :span="20">
        <el-input v-model="localList[index]" placeholder="Enter value" size="default" clearable />
      </el-col>
      <el-col :span="4">
        <el-button type="danger" icon="delete" @click="removeItem(index)">Delete</el-button>
      </el-col>
    </el-row>
    <el-button type="primary" icon="plus" @click="addItem">Add</el-button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

// 使用 v-model 语法糖
const props = defineProps({
  list: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['update:list'])

// 计算属性，确保 `localList` 和 `props.list` 同步
const localList = computed({
  get: () => props.list,
  set: (val) => emit('update:list', val)
})

// 方法：删除项
const removeItem = (index) => {
  localList.value.splice(index, 1)
}

// 方法：添加项
const addItem = () => {
  localList.value.push('')
}
</script>
