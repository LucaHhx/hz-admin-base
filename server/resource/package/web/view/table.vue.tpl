<template>
  <div>
   <DivSearch :columns="columns" ref="searchRef" >
        <template #elFormItem>
          <el-form-item >
            <el-button type="primary" :icon="Search" @click="getTableData()">{{ "{{" }} $t('common.search') {{ "}}" }}</el-button>
            <el-button :icon="Refresh" @click="onReset">{{ "{{" }} $t('common.reset') {{ "}}" }}</el-button>
          </el-form-item>
        </template>
      </DivSearch>
    <div class="gva-table-box">
    <ThreeColumnLayout :section-gap="5" :item-gap="10">
      <template #left>
       <el-button v-auth="btnAuth.add" type="primary" :icon="Plus" @click="openDialog()">{{ "{{" }} $t('common.add') {{ "}}" }}</el-button>
       <el-button v-auth="btnAuth.batchDelete" :icon="Delete" type="danger" plain
                 :disabled="!tableRef?.getSelectedData().length" @click="onDelete">{{ "{{" }} $t('common.batchDelete') {{ "}}" }}</el-button>

      </template>
      <template #center>
        <!-- 中间部分可自定义布局 -->
      </template>
      <template #right>
        <ColOptions v-if="btnAuth.columnConfig" :columns="columns" :structName="structName"/>
        <CustomExport
          :fetch-function="get{{.StructName}}List"
          :search-info="getFullSearchInfo()"
          :columns="columns"
          :file-name="structName"
          :batch-size="1000"
        />
        <CustomImport
          :import-function="import{{.StructName}}"
          :columns="columns"
          @import-success="getTableData"
          style="margin-top: 10px;"
        />
      </template>
    </ThreeColumnLayout>
    </div>

    <div class="gva-table-box">
      <Table
        :tableData="tableData"
        :columns="columns"
        rowKey="ID"
        :show-selection="true"
        @selection-change="handleSelectionChange"
        :total="total"
        :getTableData="getTableData"
        :structName="structName"
        ref="tableRef">
        <template v-if="btnAuth.info || btnAuth.edit || btnAuth.delete" #operate="scope">
          <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)">
            <el-icon class="button-icon"><InfoFilled /></el-icon>{{ "{{" }} $t('common.view') {{ "}}" }}
          </el-button>
          <el-button v-auth="btnAuth.edit" type="primary" link class="table-button" @click="update{{.StructName}}Func(scope.row)">
            <el-icon class="button-icon"><Edit /></el-icon> {{ "{{" }} $t('common.edit') {{ "}}" }}
          </el-button>
          <el-button v-auth="btnAuth.delete" type="danger" link class="table-button" @click="delete{{.StructName}}Func(scope.row)">
            <el-icon class="button-icon"><Delete /></el-icon> {{ "{{" }} $t('common.delete') {{ "}}" }}
          </el-button>
          <el-button v-auth="btnAuth.copy" type="primary" link class="table-button" @click="copy{{.StructName}}Func(scope.row)">
            <el-icon class="button-icon"><CopyDocument /></el-icon>  {{ "{{" }} $t('common.copy') {{ "}}" }}
          </el-button>
        </template>
      </Table>
    </div>
     <DiyFrom
      ref="{{.PackageName}}FormRef"
      v-model:form-data="formData"
      :init-form-data="initFormData"
      :callback="getTableData"
      :struct-name="structName"
      :columns="columns"
      :find-api="find{{.StructName}}"
      :create-api="create{{.StructName}}"
      :update-api="update{{.StructName}}"
    >
      <template #header-right>
        <FromColOptions v-if="btnAuth.columnConfig" :columns="columns" :struct-name="structName" />
      </template>
    </DiyFrom>
  </div>
</template>

<script setup>
import Table from '@/components/diyTable/table.vue'
import DivSearch from '@/components/diySearch/search.vue'
import { getColumns } from '@/utils/columns'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useBtnAuth } from '@/utils/btnAuth'
import { Search, Refresh, Delete, Edit, InfoFilled, Plus } from '@element-plus/icons-vue'
import ThreeColumnLayout from '@/components/ThreeColumnLayout/index.vue'
import {{.StructName}}Form from './{{.PackageName}}Form.vue'
import CustomExport from '@/components/exportExcel/customExport.vue'
import CustomImport from '@/components/exportExcel/customImport.vue'
import ColOptions from '@/components/diyTable/coloptions.vue'
import FromColOptions from '@/components/diyForm/fromColOptions.vue'
import DiyFrom from '@/components/diyForm/diyForm.vue'
import {
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  get{{.StructName}}List,
  import{{.StructName}},
  find{{.StructName}},
  create{{.StructName}},
  update{{.StructName}}
} from '@/api/{{.Package}}/{{.PackageName}}'

const { t } = useI18n()

const structName = '{{.PackageName}}'
defineOptions({
  name: '{{.StructName}}'
})

// 获取初始表单数据
function initFormData() {
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

const {{.PackageName}}FormRef = ref(null)
const columns = ref([])

import { useColsForRoute } from '@/utils/colAuth'
const colsAuth = useColsForRoute(structName)

getColumns(structName).then((res) => {
  columns.value = res.filter(item => colsAuth[item.jsonName]).sort((a, b) => a.sort - b.sort)
})

// 按钮权限实例化
const btnAuth = useBtnAuth()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const selectedData = ref([])
const formData = ref(initFormData())

// 处理选择变化
const handleSelectionChange = (selection) => {
  selectedData.value = selection
}

// 重置
const onReset = () => {
  tableRef?.value?.clearSearch()
  searchRef.value?.clearSearch()
  getTableData()
}

const tableRef = ref(null)
const searchRef = ref(null)

// 获取完整的搜索信息（包含搜索条件、分页信息和排序信息）
const getFullSearchInfo = () => {
  const tableQueryInfo = tableRef.value?.getQueryInfo()
  return {
    searchInfo: searchRef.value?.getQueryInfo() || {},
    pageInfo: tableQueryInfo?.pageInfo || { page: 1, pageSize: 10 },
    sortInfo: tableQueryInfo?.sortInfo || {}
  }
}

// 查询
const getTableData = async() => {
  var searchInfo = tableRef?.value?.getQueryInfo()
  if (!searchInfo) {
    searchInfo = {
      pageInfo: {
        page: 1,
        pageSize: 10
      },
      searchInfo: {},
      sortInfo: {}
    }
  }
   searchInfo.searchInfo = searchRef.value?.getQueryInfo()
      if (!searchInfo.searchInfo) {
        searchInfo.searchInfo = {}
      }
      tableData.value = []
  const table = await get{{.StructName}}List(searchInfo)
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 获取选中数据
const getSelectedRows = () => {
  return tableRef.value?.getSelectedData() || []
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('Are you sure you want to delete?','Prompt', {
    confirmButtonText: 'Confirm',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(async() => {
    const IDs = []
    const multipleSelection = getSelectedRows()
    if (multipleSelection.length === 0) {
      ElMessage({
        type: 'warning',
        message: 'Please select the data'
      })
      return
    }
    multipleSelection &&
        multipleSelection.map(item => {
          IDs.push(item.ID)
        })

    const res = await delete{{.StructName}}ByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Deletion successful'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 更新行
const update{{.StructName}}Func = async(row) => {
  {{.PackageName}}FormRef?.value?.openDialog('edit', { ID: row.ID })
}

const copy{{.StructName}}Func = async(row) => {
  {{.PackageName}}FormRef?.value?.openDialog('copy', { ID: row.ID })
}

// 删除行
const delete{{.StructName}}Func = async(row) => {
  ElMessageBox.confirm('Are you sure you want to delete?','Prompt', {
    confirmButtonText: 'Confirm',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(async() => {
    const res = await delete{{.StructName}}({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Deletion successful'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 打开弹窗
const openDialog = () => {
  {{.PackageName}}FormRef?.value?.openDialog('create', null)
}

// 打开详情
const getDetails = async(row) => {
  {{.PackageName}}FormRef?.value?.openDialog('view', row)
}

// 初始化
getTableData()
</script>
