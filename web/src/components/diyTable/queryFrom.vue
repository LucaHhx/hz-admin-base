<template>
  <el-tooltip
    :visible="visible"
    trigger="manual"
    placement="top"
    :enterable="true"
    :hide-after="0"
    popper-class="search-tooltip"
  >
    <template #default>
      <el-icon
        size="15"
        :color="isConfirm?'#2885f2':''"
        @click.stop="handleIconClick"
      >
        <Search />
      </el-icon>
    </template>
    <template #content>
      <div
        class="search-container"
        @click.stop
      >
        <template v-if="props.column.filterList">
          <!-- 非布尔类型才显示查询方式选择 -->
          <el-select
            v-if="props.column.type !== 'boolean' && props.column.type !== 'bool' "
            v-model="data.option"
            :placeholder="$t('common.queryMethod')"
            clearable
            class="search-select"
            @change="handleOptionChange"
          >
            <el-option
              v-for="item in props.column.filterList"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>

          <!-- 布尔类型 -->
          <template v-if="props.column.type === 'boolean' || props.column.type === 'bool'">
            <el-select
              v-model="data.value"
              :placeholder="$t('common.pleaseSelect')"
              clearable
              class="search-select"
              @change="handleBooleanChange"
            >
              <el-option
                v-for="item in booleanOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </template>

          <!-- 枚举类型 -->
          <template v-else-if="props.column.type === 'enum' || props.column.type === 'protoEnum'">
            <template v-if="data.option === 'in' || data.option === 'not in'">
              <el-select
                v-model="data.value"
                value-key=""
                placeholder=""
                clearable
                filterable
                multiple
              >
                <el-option
                  v-for="item in props.column.enum || []"
                  :key="item"
                  :label="$t('business.'+props.column.structName+'.' + 'enums.'+props.column.jsonName+'.'+item)"
                  :value="item"
                />
              </el-select>
            </template>
            <template v-else>
              <el-select
                v-model="data.value"
                value-key=""
                placeholder=""
                clearable
                filterable
              >
                <el-option
                  v-for="item in props.column.enum || []"
                  :key="item"
                  :label="$t('business.'+props.column.structName+'.' + 'enums.'+props.column.jsonName+'.'+item)"
                  :value="item"
                />
              </el-select>
            </template>
          </template>

          <!-- 日期时间类型 -->
          <template v-else-if="props.column.type === 'date' || props.column.type === 'datetime'">
            <template v-if="data.option === 'between' || data.option === 'not between'">
              <el-date-picker
                v-model="data.value"
                :type="props.column.type + 'range'"
                range-separator="至"
                :start-placeholder="$t('common.startDate')"
                :end-placeholder="$t('common.endDate')"
                value-format="YYYY-MM-DD HH:mm:ss"
                class="search-select"
              />
            </template>
            <template v-else>
              <el-date-picker
                v-model="data.value"
                :type="props.column.type"
                :placeholder="$t('common.pleaseSelectDate')"
                value-format="YYYY-MM-DD HH:mm:ss"
                class="search-select"
              />
            </template>
          </template>

          <!-- 字符串类型 -->
          <template v-else>
            <template v-if="data.option === 'between' || data.option === 'not between'">
              <el-input
                v-model="data.value[0]"
                :placeholder="$t('common.pleaseInput')"
                clearable
                class="search-select"
              />
              <el-input
                v-model="data.value[1]"
                :placeholder="$t('common.pleaseInput')"
                clearable
                class="search-select"
              />
            </template>
            <template v-else-if="data.option === 'in' || data.option === 'not in'">
              <el-select
                v-model="data.value"
                multiple
                filterable
                allow-create
                default-first-option
                :reserve-keyword="false"
                :placeholder="$t('common.pleaseInput')"
                class="search-select"
              />
            </template>
            <template v-else>
              <el-input
                v-model="data.value"
                :placeholder="$t('common.pleaseInput')"
                clearable
                class="search-select"
              />
            </template>
          </template>
          <div class="search-buttons">
            <el-button
              type="primary"
              class="search-button"
              @click="handleSearch"
            >
              {{ $t('common.search') }}
            </el-button>
            <el-button
              class="search-button"
              @click="handleReset"
            >
              {{ $t('common.reset') }}
            </el-button>
          </div>
        </template>
      </div>
    </template>
  </el-tooltip>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { convertDateTime, convertDate } from '@/utils/time'

const props = defineProps({
  column: {
    type: Object,
    default: () => ({})
  },
  isConfirm: {
    type: Boolean,
    default: false
  },
  confirmed: {
    type: Function,
    default: () => {}
  },
  cleared: {
    type: Function,
    default: () => {}
  }
})

// 布尔值选项
const booleanOptions = [
  { label: 'true', value: true },
  { label: 'false', value: false }
]

const visible = ref(false)
const data = ref({
  option: props.column.defaultFilter || null,
  value: null
})

// 处理图标点击
const handleIconClick = () => {
  visible.value = !visible.value
  if (visible.value && props.column.type === 'boolean' && !data.value?.option) {
    data.value.option = '='
  }
}

// 处理选项变化
const handleOptionChange = (value) => {
  if (value === 'between' || value === 'not between') {
    data.value.value = [null, null]
  } else if (value === 'in' || value === 'not in') {
    data.value.value = []
  } else {
    data.value.value = null
  }
}

// 处理布尔值变化
const handleBooleanChange = (val) => {
  data.value.option = val !== null ? '=' : null
}

import { unformatNumber } from '@/utils/format'
// 处理搜索
const handleSearch = () => {
  if (!data.value.option && (props.column.type !== 'boolean' || props.column.type !== 'bool')) {
    return
  }

  const confirmed = {
    option: data.value.option,
    type: props.column.type,
    value: convertToString(data.value.value, props.column.type, props.column.format)
  }

  props.confirmed(confirmed)
  visible.value = false
}

// 转换函数
const convertToString = (value, type) => {
  if (value === null || value === undefined) {
    return ''
  }

  // 如果是数组，用逗号连接
  if (Array.isArray(value)) {
    if (type === 'int64' || type === 'int32' || type === 'amount') {
      return value.map(item => unformatNumber(item, format)).join(',')
    }
    return value.join(',')
  }

  // 如果是布尔值
  if (typeof value === 'boolean') {
    return value ? '1' : '0'
  }

  // if (type === 'date' || type === 'datetime') {
  //   return value.toISOString()
  // }
  if (props.column.type === 'int64' || props.column.type === 'int32' || props.column.type === 'amount') {
    value = unformatNumber(value, props.column.format)
  }
  // 其他类型直接转字符串
  return String(value)
}

// 处理重置
const handleReset = () => {
  data.value = {
    option: props.column.defaultFilter,
    value: null
  }
  props.cleared(props.column.prop)
  visible.value = false
}

// 处理点击外部
const handleClickOutside = (e) => {
  // 优化点击检测逻辑
  const tooltips = document.querySelectorAll('.search-tooltip')
  const isClickInside = Array.from(tooltips).some(tooltip => tooltip.contains(e.target))
  const isClickOnIcon = e.target.closest('.el-icon')
  const isClickOnPicker = e.target.closest('.el-picker__popper') ||
          e.target.closest('.el-date-picker') ||
          e.target.closest('.el-time-picker') ||
          e.target.closest('.el-select-dropdown')
  if (!isClickInside && !isClickOnIcon && !isClickOnPicker) {
    visible.value = false
  }
}

// 监听组件挂载
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

// 监听组件卸载
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.search-container {
    min-width: 240px;
    padding: 8px;
}

.search-select {
    width: 100%;
    margin-bottom: 8px;
}

.search-buttons {
    display: flex;
    justify-content: space-between;
    gap: 8px;
    margin-top: 16px;
}

.search-button {
    flex: 1;
}
</style>
