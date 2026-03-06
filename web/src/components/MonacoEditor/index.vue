<template>
  <el-input
    v-model="inputValue"
    type="textarea"
    :rows="10"
    :autosize="{ minRows: 10, maxRows: 20 }"
    :placeholder="language === 'json' ? '请输入 JSON 内容' : '请输入内容'"
    @input="handleInput"
  />
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: [Object, String],
    default: null
  },
  language: {
    type: String,
    default: 'json'
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

const inputValue = ref('')

// 监听值变化
watch(() => props.modelValue, (newVal) => {
  if (newVal !== undefined && newVal !== null) {
    inputValue.value = typeof newVal === 'string' ? newVal : JSON.stringify(newVal, null, 2)
  }
}, { immediate: true, deep: true })

const handleInput = (value) => {
  try {
    if (props.language === 'json') {
      const parsed = JSON.parse(value)
      emit('update:modelValue', parsed)
      emit('change', parsed)
    } else {
      emit('update:modelValue', value)
      emit('change', value)
    }
  } catch {
    // 如果 JSON 格式不正确，仍然更新原始值
    emit('update:modelValue', value)
    emit('change', value)
  }
}

// 暴露方法
defineExpose({
  refresh: () => {
    // 重置值
    if (props.modelValue !== undefined && props.modelValue !== null) {
      inputValue.value = typeof props.modelValue === 'string'
        ? props.modelValue
        : JSON.stringify(props.modelValue, null, 2)
    }
  }
})
</script>

<style lang="scss" scoped>
:deep(.el-textarea__inner) {
  font-family: 'Menlo', 'Monaco', 'Courier New', monospace;
  font-size: 14px;
  line-height: 1.5;
  padding: 12px;
  border-radius: 4px;
  resize: vertical;
  @apply dark:bg-slate-800 dark:text-gray-200 dark:border-gray-600;
}
</style>
