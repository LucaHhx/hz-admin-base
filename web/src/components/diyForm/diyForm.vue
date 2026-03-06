<template>
  <el-dialog
    v-model="visible"
    :title="$t('common.' + formType)"
    :width="width"
    :style="{ '--dialog-height': height }"
    :close-on-click-modal="true"
    :close-on-press-escape="false"
    :rules="rules"
    draggable
    @close="handleClose"
  >
    <template #header>
      <ThreeColumnLayout :section-gap="5" :item-gap="10">
        <template #left>
          <span>{{ $t('menu.'+ props.structName) +'-'+ $t('common.'+ formType) }}</span>
          <slot name="header-left" />
        </template>
        <template #center>
          <slot name="header-center" />
        </template>
        <template #right>
          <slot name="header-right" />
        </template>
      </ThreeColumnLayout>
    </template>

    <div class="hab-form-box">
      <el-form ref="formRef" :model="form" label-position="right" :rules="rules" inline label-width="auto">
        <el-form-item
          v-for="column in visibleColumns"
          :key="column.columnName" :label="$t('business.'+props.structName+'.columns.' + column.jsonName)" :prop="column.jsonName"
          :style="{width: column.formWith+'%' }"
        >
          <template v-if="$slots['column-'+column.jsonName] ">
            <slot :name="'column-'+column.jsonName" :column="column" :row="form" :view-type="formType" :is-disabled="isDisabled(column)" />
          </template>
          <template v-else>
            <template v-if="column.type === 'boolean' || column.type === 'bool'">
              <el-switch v-model="form[column.jsonName]" :disabled="isDisabled(column)" />
            </template>
            <template v-else-if="column.type === 'table'">
              <SelectTool
                v-model:value="form[column.jsonName]"
                :filter-id="column.editorFilterId"
                :disabled="isDisabled(column)"
              />
            </template>

            <template v-else-if="column.type === 'enum' ">
              <el-select
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseSelect')"
                :disabled="isDisabled(column)"
                clearable
                style="width: 100%"
              >
                <el-option
                  v-for="item in column.enum || []"
                  :key="item"
                  :label="$t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+item.replaceAll('.', ''))"
                  :value="item"
                />
              </el-select>
            </template>

            <template v-else-if="column.type === 'protoEnum' ">
              <el-select
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseSelect')"
                :disabled="isDisabled(column)"
                clearable
                style="width: 100%"
              >
                <el-option
                  v-for="item in column.enum || []"
                  :key="item"
                  :label="$t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+item.replaceAll('.', ''))"
                  :value="Number(item)"
                />
              </el-select>
            </template>

            <template v-else-if="column.type === 'enums'">
              <el-select
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseSelect')"
                :disabled="isDisabled(column)"
                clearable
                style="width: 100%"
              >
                <el-option
                  v-for="item in column.enum || []"
                  :key="item"
                  :label="$t('business.'+props.structName+'.' + 'enums.'+column.jsonName+'.'+item.replaceAll('.', ''))"
                  :value="item"
                />
              </el-select>
            </template>

            <template v-else-if="column.type === 'date' || column.type === 'datetime' || column.type === 'uintDate'">
              <el-date-picker
                v-model="form[column.jsonName]"
                :type="column.type"
                :placeholder="$t('common.pleaseSelect')"
                :disabled="isDisabled(column)"
                style="width: 100%"
                format="YYYY-MM-DD HH:mm:ss"
                value-format="YYYY-MM-DD HH:mm:ss"
                @change="(value)=>{handleDateChange(value,column.jsonName)}"
              />
              <!-- <span>服务器时间：{{ $serverDate(form[column.jsonName],'long') }}</span> -->
              <!-- <TimezoneDatePicker
                v-model="form[column.jsonName]"
                :type="column.type"
                :placeholder="$t('common.pleaseSelect')"
                :disabled="isDisabled(column)"
              /> -->
            </template>

            <template v-else-if="column.type === 'int64' || column.type === 'int32' || column.type === 'amount' || column.type === 'number'">
              <el-input-number
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseInput')"
                :disabled="isDisabled(column)"
                style="width: 100%"
              />
            </template>
            <template v-else-if="column.type === 'float' || column.type === 'double' || column.type === 'float64'">
              <el-input-number
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseInput')"
                :disabled="isDisabled(column)"
                style="width: 100%"
                :precision="4"
              />
            </template>

            <template v-else-if="column.type === 'textarea'">
              <el-input
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseInput')"
                :disabled="isDisabled(column)"
                type="textarea"
                autosize
                clearable
                style="width: 100%"
              />
            </template>
            <template v-else-if="column.type === 'string'">
              <el-input
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseInput')"
                :disabled="isDisabled(column)"
                clearable
                style="width: 100%"
              />
            </template>
            <template v-else-if="column.type === 'object'">
              <span>{{ form[column.jsonName] }}</span>
            </template>
            <template v-else>
              <el-input
                v-model="form[column.jsonName]"
                :placeholder="$t('common.pleaseInput')"
                :disabled="isDisabled(column)"
                type="textarea"
                autosize
                clearable
                style="width: 100%"
              />
            </template>
          </template>
        </el-form-item>
        <slot
          name="additional" :columns="Object.fromEntries(visibleColumns.map(item => [item.jsonName, item]))" :row="form"
          :is-disabled-map="Object.fromEntries(allColumns.map((k, i) => [k.jsonName, isDisabled(k)]))"
        />
      </el-form>
    </div>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">{{ $t('common.cancel') }}</el-button>
        <el-button v-if="!isView" type="primary" :loading="loading" @click="handleSubmit">{{ $t('common.confirm') }}</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>

import { ElMessage } from 'element-plus'
import { ref, computed } from 'vue'
import { convertDate, convertDateTime } from '@/utils/time'
import ThreeColumnLayout from '@/components/ThreeColumnLayout/index.vue'
import SelectTool from '@/components/filterPop/selectTool.vue'
import { formatTimeToStr } from '@/utils/date'
// import TimezoneDatePicker from './TimezoneDatePicker.vue'
import {} from '@/i18n'
const props = defineProps({
  callback: {
    type: Function,
    default: null
  },
  structName: {
    type: String,
    required: true
  },
  columns: {
    type: Array,
    required: true
  },
  formatItem: {
    type: Function,
    default: null
  },
  findApi: {
    type: Function,
    default: null
  },
  createApi: {
    type: Function,
    default: null
  },
  updateApi: {
    type: Function,
    default: null
  },
  formData: {
    type: Object,
    default: null
  },
  defaultRules: {
    type: Object,
    default: () => {
      return {}
    }
  },
  initFormData: {
    type: Function,
    default: () => {
      return {}
    }
  },
  visibleColumnsFunc: {
    type: Function,
    default: null
  },
  width: {
    type: String,
    default: '50%'
  },
  height: {
    type: String,
    default: '50%'
  }
})

const rules = computed(() => {
  var newRules = {}
  props.columns?.forEach(column => {
    if (column.formMust) {
      newRules[column.jsonName] = [{ required: true, message: 'required' }]
    }
  })
  return { ...newRules, ...props.defaultRules }
})

defineOptions({
  name: 'DiyForm'
})

// 计算属性
const isView = computed(() => formType.value === 'view')

// 过滤出可见的列
const visibleColumns = computed(() => {
  return props.columns.filter(column => {
    if (props.visibleColumnsFunc) {
      return props.visibleColumnsFunc(column, form.value) && !column.formHidden
    }
    return !column.formHidden
  }).sort((a, b) => a.formOrder - b.formOrder)
})

const allColumns = computed(() => {
  return props.columns.filter(a => { return true }).sort((a, b) => a.formOrder - b.formOrder)
})

// 判断字段是否禁用
const isDisabled = (column) => {
  return isView.value || (column.formDisabled && (formType.value === 'update' || formType.value === 'edit')) || false
}

// 状态
const visible = ref(false)
// 加载状态
const loading = ref(false)
// 表单引用
const formRef = ref()
// 表单数据
const form = ref(props.formData)
// 表单类型
const formType = ref('create')

// 打开弹窗
async function openDialog(type, data) {
  visible.value = true
  formType.value = type
  if (formType.value === 'create') {
    form.value = props.initFormData()
  } else if (formType.value === 'external') {
    form.value = data
  } else if (formType.value === 'copy') {
    const res = await props.findApi(data)
    const copyData = { ...res.data }
    delete copyData.ID
    delete copyData.createdAt
    delete copyData.updatedAt
    delete copyData.deletedAt
    form.value = copyData
  } else {
    try {
      const res = await props.findApi({ key: data.key, ID: data.ID })
      if (res.code === 0) {
        // 处理json字符串
        res.data = toJsonStr(res.data)
        form.value = res.data
      } else {
        ElMessage.error('Failed to obtain data')
      }
    } catch (error) {
      console.error('Failed to retrieve data:', error)
      ElMessage.error('Failed to obtain data')
    }
  }
}

const itemFormat = (item) => {
  props.columns.forEach(column => {
  })
  item = strToJson(item)
  if (props.formatItem) {
    return props.formatItem(item)
  }
  return item
}

// 提交表单
async function handleSubmit() {
  if (!formRef.value) return
  try {
    await formRef.value.validate()
    loading.value = true
    let api = null
    switch (formType.value) {
      case 'create':
        api = props.createApi
        break
      case 'external':
        api = props.createApi
        break
      case 'edit':
        api = props.updateApi
        break
      case 'view':
        api = null
        break
      case 'copy':
        api = props.createApi
        break
    }

    if (api) {
      const res = await api(itemFormat({ ... form.value }))
      if (res.code === 0) {
        ElMessage.success('operation is successful')
        handleClose()
      } else {
        ElMessage.error(res.message || 'operation failed')
      }
    }
  } catch (error) {
    console.error('Submission failed:', error)
    ElMessage.error('Submission failed')
  } finally {
    if (props.callback) {
      props.callback()
    }
    loading.value = false
  }
}

// 关闭弹窗
function handleClose() {
  visible.value = false
  form.value = props.initFormData()
  formRef.value?.resetFields()
}

// 处理json字符类型的数据
const toJsonStr = (data) => {
  props.columns.filter(v => v.type === 'jsonStr').forEach(column => {
    const v = data[column.jsonName]
    if (v) {
      if (typeof v !== 'string') {
        data[column.jsonName] = JSON.stringify(v)
      }
    } else {
      data[column.jsonName] = '{}'
    }
  })
  return data
}

const strToJson = (data) => {
  props.columns.filter(v => v.type === 'jsonStr').forEach(column => {
    const v = data[column.jsonName]
    if (v) {
      if (typeof v === 'string') {
        data[column.jsonName] = JSON.parse(v)
      }
    } else {
      data[column.jsonName] = {}
    }
  })
  return data
}

// 暴露方法
defineExpose({
  openDialog
})

// import { useI18n } from 'vue-i18n'
// const in18 = useI18n()

// import { DateTime } from 'luxon'
// // const { DateTime } = require('luxon')

const handleDateChange = (value, data) => {
  // const date = new Date(value)
  // console.log(formatTimeToStr(date)) // 输出日期对象;
  // // 解析输入字符串，指定输入格式和时区
  // const dt = DateTime.fromFormat(formatTimeToStr(date), 'yyyy-MM-dd HH:mm:ss', { zone: in18.timeZone })
  // // 转换为 ISO8601 字符串，保留时区偏移
  // console.log(dt.toISO()) // 默认格式为：2025-08-15T16:52:01+08:00);
  // // data = dt.toISO() // 将转换后的值赋给 data
  // // form[data]
  // form.value[data] = dt
  // console.log(form.value[data]) // 输出转换后的值
}
</script>

<style lang="scss" scoped>
.hab-form-box {
  padding: 12px;
  max-height: var(--dialog-height, auto);
  overflow-y: auto;
}

.el-form-item {
  margin-bottom: 12px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}

</style>
