<template>
  <el-dialog
    v-model="visible"
    :title="$t('common.' + formType)"
    width="50%"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    @close="handleClose"
  >
    <div class="gva-form-box">
      <el-form :model="form" ref="formRef" label-position="right" :rules="rules" label-width="auto">
{{- range .Fields}}
{{- if .Form}}
        <el-form-item :label="$t('business.{{$.PackageName}}.columns.{{.FieldJson}}') + ':'" prop="{{.FieldJson}}">
{{- if .CheckDataSource}}
          <el-select {{if eq .DataSource.Association 2}} multiple {{ end }} v-model="form.{{.FieldJson}}" placeholder="please select {{.FieldJson}}" style="width:100%" :clearable="{{.Clearable}}" :disabled="isView">
            <el-option v-for="(item,key) in dataSource.{{.FieldJson}}" :key="key" :label="item.label" :value="item.value" />
          </el-select>
{{- else }}
          {{- if eq .FieldType "bool" }}
          <el-switch v-model="form.{{.FieldJson}}" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" :disabled="isView" />
          {{- else if eq .FieldType "string" }}
          <el-input v-model="form.{{.FieldJson}}" :clearable="{{.Clearable}}" placeholder="please enter {{.FieldJson}}" :disabled="isView" />
          {{- else if eq .FieldType "richtext" }}
          <el-input type="textarea" v-model="form.{{.FieldJson}}" autosize :disabled="isView" clearable />
          {{- else if eq .FieldType "int" }}
          <el-input-number v-model="form.{{.FieldJson}}" :disabled="isView" :clearable="{{.Clearable}}" />
          {{- else if eq .FieldType "int32" }}
          <el-input-number v-model="form.{{.FieldJson}}" :disabled="isView" :clearable="{{.Clearable}}" />
          {{- else if eq .FieldType "int64" }}
          <el-input-number v-model="form.{{.FieldJson}}" :disabled="isView" :clearable="{{.Clearable}}" />
          {{- else if eq .FieldType "date" }}
          <el-date-picker v-model="form.{{.FieldJson}}" type="date" value-format="YYYY-MM-DD" :disabled="isView" placeholder="Select Date" :clearable="{{.Clearable}}"/>
          {{- else if eq .FieldType "datetime" }}
          <el-date-picker v-model="form.{{.FieldJson}}" type="datetime" value-format="YYYY-MM-DD HH:mm:ss" :disabled="isView" placeholder="Selection time" :clearable="{{.Clearable}}"/>
          {{- else if eq .FieldType "float64" }}
          <el-input-number v-model="form.{{.FieldJson}}" :precision="2" :disabled="isView" :clearable="{{.Clearable}}" />
          {{- else if eq .FieldType "enum" }}
          <el-select v-model="form.{{.FieldJson}}" placeholder="please select" style="width:100%" :clearable="{{.Clearable}}" :disabled="isView">
            <el-option v-for="item in 3" :key="item" :label="$t('business.{{$.PackageName}}.enums.{{.FieldJson}}')+item" :value="item" />
          </el-select>
          {{- else if eq .FieldType "json" }}
          <el-input type="textarea" v-model="form.{{.FieldJson}}" autosize :disabled="isView" clearable />
          {{- else if eq .FieldType "array" }}
          <el-input type="textarea" v-model="form.{{.FieldJson}}" autosize :disabled="isView" clearable />
          {{- else if eq .FieldType "binary" }}
          <el-input type="textarea" v-model="form.{{.FieldJson}}" autosize :disabled="isView" clearable />
          {{- end }}
{{- end }}
        </el-form-item>
{{- end }}
{{- end }}
      </el-form>
    </div>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">{{ "{{" }} $t('common.cancel') {{ "}}" }}</el-button>
        <el-button v-if="!isView" type="primary" :loading="loading" @click="handleSubmit">{{ "{{" }} $t('common.confirm') {{ "}}" }}</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>

import { create{{.StructName}}, update{{.StructName}}, find{{.StructName}} } from '@/api/{{.Package}}/{{.PackageName}}'
import { ElMessage } from 'element-plus'
import { ref, computed } from 'vue'
import { convertDate, convertDateTime } from '@/utils/time'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const props = defineProps({
  callback: {
    type: Function,
    default: null
  }
})

defineOptions({
  name: '{{.StructName}}Form'
})

// 计算属性
const isView = computed(() => formType.value === 'view')
// 状态
const visible = ref(false)
// 加载状态
const loading = ref(false)
// 表单引用
const formRef = ref()
// 表单数据
const form = ref(getInitialForm())
// 表单类型
const formType = ref('create')

{{- if .HasDataSource }}
// 数据源
const dataSource = ref([])
const getDataSourceFunc = async()=>{
  const res = await get{{.StructName}}DataSource()
  if (res.code === 0) {
    dataSource.value = res.data
  }
}
getDataSourceFunc()
{{- end }}

// 表单校验规则
const rules = {
{{- range .Fields }}
  {{- if .Require }}
  {{.FieldJson}}: [{ required: true, message: '{{ .ErrorText }}', trigger: 'blur' }],
  {{- end }}
{{- end }}
}

// 获取初始表单数据
function getInitialForm() {
  return {
{{- range .Fields}}
  {{- if .Form}}
    {{- if eq .FieldType "bool" }}
    {{.FieldJson}}: false,
    {{- else if eq .FieldType "string" }}
    {{.FieldJson}}: '',
    {{- else if eq .FieldType "richtext" }}
    {{.FieldJson}}: '',
    {{- else if eq .FieldType "int" }}
    {{.FieldJson}}: 0,
    {{- else if eq .FieldType "int32" }}
    {{.FieldJson}}: 0,
    {{- else if eq .FieldType "int64" }}
    {{.FieldJson}}: 0,
    {{- else if eq .FieldType "date" }}
    {{.FieldJson}}: null,
    {{- else if eq .FieldType "datetime" }}
    {{.FieldJson}}: null,
    {{- else if eq .FieldType "float64" }}
    {{.FieldJson}}: 0,
    {{- else if eq .FieldType "enum" }}
    {{.FieldJson}}: '',
    {{- else if eq .FieldType "json" }}
    {{.FieldJson}}: {},
    {{- else if eq .FieldType "array" }}
    {{.FieldJson}}: [],
    {{- else if eq .FieldType "binary" }}
    {{.FieldJson}}: '',
    {{- end }}
  {{- end }}
{{- end }}
  }
}

// 打开弹窗
async function openDialog(type, data) {
  visible.value = true
  formType.value = type
  if (formType.value === 'create') {
    form.value = getInitialForm()
  } else if (formType.value === 'copy') {
    const copyData = { ...data }
    delete copyData.ID
    delete copyData.createdAt
    delete copyData.updatedAt
    delete copyData.deletedAt
    form.value = copyData
  } else {
    try {
      const res = await find{{.StructName}}({ ID: data.id })
      if (res.code === 0) {
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
  {{- range .Fields}}
  {{- if eq .FieldType "date" }}
  item.{{.FieldJson}} = convertDate(item.{{.FieldJson}})
  {{- else if eq .FieldType "datetime" }}
  item.{{.FieldJson}} = convertDateTime(item.{{.FieldJson}})
  {{- end }}
  {{- end }}
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
        api = create{{.StructName}}
        break
      case 'edit':
        api = update{{.StructName}}
        break
      case 'view':
        api = null
        break
      case 'copy':
        api = create{{.StructName}}
        break
    }

    if (api) {
      const res = await api(itemFormat(form.value))
      if (res.code === 0) {
        ElMessage.success('operation is successful')
        props.onSuccess?.(form.value)
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
  form.value = getInitialForm()
  formRef.value?.resetFields()
}

// 暴露方法
defineExpose({
  openDialog
})
</script>

<style lang="scss" scoped>

</style>
