<template>
  <div style="width: 100%">
    <div v-if="!props.disabled">
      <el-select
        v-model="pvalue"
        clearable
        :multiple="props.multiple"
        filterable
        remote
        reserve-keyword
        :placeholder="props.placeholder"
        :remote-show-suffix="props.selectable"
        :remote-method="remoteMethod"
        :loading="loading"
        style="width: 100%"
        :disabled="props.disabled"
        @change="handleChange"
      >
        <el-option
          v-for="item in options"
          :key="item[filterInfo.key]"
          :label="item[filterInfo.label]"
          :value="item[filterInfo.value]"
        >
          <span v-for="filter in filterInfo.columns.filter(f => f.isShow).sort((a, b) => a.sort - b.sort)" :key="filter.columnName" style="margin-right: 10px;">
            {{ item[filter.columnName] }}
          </span>
        </el-option>
      </el-select>
      <span v-if="selectedOption?.note && !props.notNote">{{ selectedOption?.note }}</span>
    </div>

    <span v-else>
      <span v-for="filter in filterInfo?.columns?.filter(f => f.isShow).sort((a, b) => a.sort - b.sort)" :key="filter.columnName" style="margin-right: 10px;">
        {{ findLabel(props.value)[filter.columnName] }}
      </span>
    </span>
  </div>
</template>

<script setup>
import { ref, defineProps, defineEmits, watch } from 'vue'
import { requestWithCache, requestFilterData } from '@/utils/sysDataFilterCache'

const props = defineProps({
  filterId: {
    type: Number,
    default: 0
  },
  multiple: {
    type: Boolean,
    default: false
  },
  default: {
    type: Array,
    default: () => []
  },
  value: {
    type: [Object, Array, String, Number],
    default: () => null
  },
  disabled: {
    type: Boolean,
    default: false
  },
  selectable: {
    type: Boolean,
    default: true
  },
  change: {
    type: Function,
    default: () => {}
  },
  placeholder: {
    type: String,
    default: 'Please Select'
  },
  filter: {
    type: Function,
    default: null
  },
  notNote: {
    type: Boolean,
    default: false
  }

})

const emit = defineEmits(['update:value', 'change'])

const pvalue = ref(props.value)
const loading = ref(false)
const filterInfo = ref({})
const options = ref([])

// 监听外部值变化
watch(() => props.value, (newVal) => {
  if (JSON.stringify(newVal) !== JSON.stringify(pvalue.value)) {
    pvalue.value = newVal
    // 当外部值变化时，更新 selectedOption
    if (newVal && options.value.length > 0) {
      selectedOption.value = options.value.find(item => item[filterInfo.value.key] === newVal)
    }
  }
}, { deep: true })

const init = async() => {
  try {
    const res = await requestWithCache(props.filterId)
    filterInfo.value = res.filter
    options.value = res.data
    // 初始化时如果已经有值，设置 selectedOption
    if (pvalue.value && options.value.length > 0) {
      selectedOption.value = options.value.find(item => item[filterInfo.value.key] === pvalue.value)
    }
  } catch (error) {
    console.error('Failed to initialize filter:', error)
  }
}

init()

const remoteMethod = async(query) => {
  loading.value = true
  try {
    console.log('props.default', props.default)
    const querys = [query, ...props.default].filter(item => item)
    console.log('querys', querys)
    const res = await requestFilterData(querys, props.filterId)
    if (props.filter) {
      options.value = res.data.filter(props.filter)
    } else {
      options.value = res.data
    }
  } catch (error) {
    console.error('Failed to load filter data:', error)
  } finally {
    loading.value = false
  }
}

const findLabel = (value) => {
  const item = options.value.find(item => item[filterInfo.value.key] === value)
  return item || {}
}

const selectedOption = ref(null)
const handleChange = (value) => {
  emit('update:value', value)
  emit('change', value)
  selectedOption.value = options.value.find(item => item[filterInfo.value.key] === value)
  props.change(value, selectedOption.value)
}

</script>
