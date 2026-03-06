<template>
  <div class="hab-table-box">
    <el-form
      ref="form"
      :rules="rules"
      label-width="80px"
      :inline="true"
      size="default"
    >
      <el-form-item
        v-for="column in needSearchColumns"
        :key="column.columnName"
        :style="{width:column.searchWidth<=0?'220px':column.searchWidth+'px'}"
      >
        <!-- 布尔类型 -->
        <template v-if="column.type === 'boolean' || column.type === 'bool'">
          <el-select
            v-model="formData[column.columnName]"
            :placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)"
            clearable
          >
            <el-option
              v-for="item in booleanOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </template>

        <!-- 关联表 -->
        <template v-else-if="column.type === 'table'">
          <SelectTool v-model:value="formData[column.columnName]" :filter-id="column.filterId" :placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)" :not-note="true" />
        </template>

        <!-- 枚举类型 -->
        <template v-else-if="column.type === 'enum'">
          <!-- 枚举多选 -->
          <template v-if="column.defaultFilter === 'in' || column.defaultFilter === 'not in'">
            <el-select
              v-model="formData[column.columnName]"
              :placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)"
              multiple
              clearable
            >
              <el-option
                v-for="item in column.enum || []"
                :key="item"
                :label="$t('business.'+column.structName+'.' + 'enums.'+column.jsonName+'.'+(item+'').replaceAll('.', ''))"
                :value="item"
              />
            </el-select>
          </template>
          <!-- 枚举单选 -->
          <template v-else>
            <el-select
              v-model="formData[column.columnName]"
              :placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)"
              clearable
            >
              <el-option
                v-for="item in column.enum || []"
                :key="item"
                :label="$t('business.'+column.structName+'.' + 'enums.'+column.jsonName+'.'+(item+'').replaceAll('.', ''))"
                :value="item"
              />
            </el-select>
          </template>
        </template>

        <!-- 日期时间类型 -->
        <template v-else-if="column.type === 'date' || column.type === 'datetime' || column.type === 'uintDate'">
          <!-- 日期范围 -->
          <template v-if="column.defaultFilter === 'between' || column.defaultFilter === 'not between'">
            <el-date-picker
              v-model="formData[column.columnName]"
              :type="column.type + 'range'"
              range-separator="至"
              :start-placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)+' ('+ $t('common.include') +')'"
              :end-placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)+' ('+ $t('common.exclude') +')'"
              :value-format="column.type === 'date' ? 'YYYY-MM-DD' : column.type === 'uintDate' ? 'YYYYMMDD' : 'YYYY-MM-DD HH:mm:ss'"
            />
          </template>
          <!-- 单个日期 -->
          <template v-else>
            <el-date-picker
              v-model="formData[column.columnName]"
              :type="column.type"
              :placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)"
              :value-format="column.type === 'date' ? 'YYYY-MM-DD' : column.type === 'uintDate' ? 'YYYYMMDD' : 'YYYY-MM-DD HH:mm:ss'"
            />
          </template>
        </template>

        <!-- 其他类型 -->
        <template v-else>
          <!-- 范围查询 -->
          <template v-if="column.defaultFilter === 'between' || column.defaultFilter === 'not between'">
            <div class="between-inputs">
              <el-input
                v-model="formData[column.columnName][0]"
                :placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)"
                class="between-input"
              />
              <span class="between-separator">
                -
              </span>
              <el-input
                v-model="formData[column.columnName][1]"
                :placeholder="$t('common.max')"
                class="between-input"
              />
            </div>
          </template>
          <!-- 多值查询 -->
          <template v-else-if="column.defaultFilter === 'in' || column.defaultFilter === 'not in'">
            <el-select
              v-model="formData[column.columnName]"
              multiple
              filterable
              allow-create
              default-first-option
              :reserve-keyword="false"
              :placeholder="$t('common.pleaseInput')"
            />
          </template>
          <!-- 普通输入 -->
          <template v-else>
            <el-input
              v-model="formData[column.columnName]"
              :placeholder="$t('business.'+column.structName+'.columns.'+column.jsonName)"
            />
          </template>
        </template>
      </el-form-item>
      <slot name="elFormItem" />
    </el-form>
    <!-- <el-button type="primary" @click="getQueryInfo()">{{ $t('common.search') }}</el-button> -->
  </div>
</template>

<script setup>
import { ref, defineProps, computed, watch } from 'vue'
import { unformatNumber } from '@/utils/format'
import SelectTool from '@/components/filterPop/selectTool.vue'
const props = defineProps({
  columns: {
    type: Array,
    default: () => []
  },
  defaultFilter: {
    type: Object,
    default: () => {}
  }
})

// 布尔值选项
const booleanOptions = [
  { label: 'true', value: true },
  { label: 'false', value: false }
]

const needSearchColumns = computed(() => {
  return props.columns.filter(item => item.isAddSearch)
})

const needSearchColumnsMap = computed(() => {
  return props.columns.reduce((acc, item) => {
    if (item.isAddSearch) {
      acc[item.columnName] = item
    }
    return acc
  }, {})
})

const formData = ref({

})

// 监听 needSearchColumns 变化，初始化 formData
watch(needSearchColumns, (columns) => {
  columns.forEach(column => {
    // 根据不同的筛选条件初始化数据
    if (props?.defaultFilter && props?.defaultFilter[column?.columnName]) {
      formData.value[column.columnName] = props.defaultFilter[column.columnName]
    } else if (column.defaultFilter === 'between' || column.defaultFilter === 'not between') {
      formData.value[column.columnName] = [null, null]
    } else if (column.defaultFilter === 'in' || column.defaultFilter === 'not in') {
      formData.value[column.columnName] = []
    } else {
      formData.value[column.columnName] = null
    }
  })
}, { immediate: true })

const rules = ref({})

const getQueryInfo = (def) => {
  if (def) {
    formData.value = { ...formData.value, ...def }
  }

  const searchInfo = {}
  for (const key in formData.value) {
    const column = needSearchColumnsMap.value[key]
    if (formData.value[key] === null || formData.value[key] === '' ||
        (Array.isArray(formData.value[key]) && formData.value[key].length === 0)) {
      continue
    }
    searchInfo[key] = {
      option: column?.defaultFilter,
      type: column?.type,
      value: convertToString(formData.value[key], column)
    }
  }
  return searchInfo
}

const convertToString = (value, column) => {
  if (value === null || value === undefined || !column) {
    return ''
  }

  // 如果是数组，用逗号连接
  if (Array.isArray(value)) {
    if (column.type === 'int64' || column.type === 'int32' || column.type === 'amount' || column.type === 'enum') {
      return value.map(item => unformatNumber(item, column.format)).join(',')
    }
    const str = value.join(',')
    return str === ',' ? '' : str
  }

  // 如果是布尔值
  if (typeof value === 'boolean') {
    return value ? '1' : '0'
  }

  console.log(column)

  if (column.type === 'int64' || column.type === 'int32' || column.type === 'amount') {
    value = unformatNumber(value, column.format)
  }
  return String(value)
}

const clearSearch = () => {
  formData.value = {}
}

// 用于外部更改搜索条件
const setSearchInfo = (searchInfo) => {
  // 根据传入的searchInfo更新formData
  for (const key in searchInfo) {
    const column = needSearchColumnsMap.value[key]
    if (column) {
      const searchItem = searchInfo[key]
      // 根据不同的筛选条件和数据类型处理值
      if (column.defaultFilter === 'between' || column.defaultFilter === 'not between') {
        // 范围查询，期望值是数组格式
        if (typeof searchItem.value === 'string' && searchItem.value.includes(',')) {
          formData.value[key] = searchItem.value.split(',')
        } else {
          formData.value[key] = [searchItem.value, null]
        }
      } else if (column.defaultFilter === 'in' || column.defaultFilter === 'not in') {
        // 多值查询，期望值是数组格式
        if (typeof searchItem.value === 'string') {
          formData.value[key] = searchItem.value.split(',')
        } else {
          formData.value[key] = Array.isArray(searchItem.value) ? searchItem.value : [searchItem.value]
        }
      } else {
        // 普通查询
        if (column.type === 'boolean' || column.type === 'bool') {
          formData.value[key] = searchItem.value === '1' || searchItem.value === true
        } else {
          formData.value[key] = searchItem.value
        }
      }
    }
  }
}

// 暴露给父组件的方法
defineExpose({
  clearSearch,
  getQueryInfo,
  setSearchInfo,
})
</script>

<style scoped>
.hab-table-box {
  padding: 12px;
}

.el-form-item {
  margin-right: 10px;
}

.el-date-picker,
.el-select {
  width: 100%;
}

.between-inputs {
  display: flex;
  align-items: center;
  gap: 8px;
}

.between-input {
  flex: 1;
}

.between-separator {
  color: #909399;
}
</style>
