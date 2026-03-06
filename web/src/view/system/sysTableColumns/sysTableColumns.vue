<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <!-- <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
<el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
  :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
—
<el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
  :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
</el-form-item> -->

        <el-form-item :label="$t('business.sysTableColumns.columns.menuId')" prop="menuId">
          <el-input v-model="searchInfo.menuId" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.structName')" prop="structName">
          <el-input v-model="searchInfo.structName" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.jsonName')" prop="jsonName">
          <el-input v-model="searchInfo.jsonName" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.columnName')" prop="columnName">
          <el-input v-model="searchInfo.columnName" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.with')" prop="with">
          <el-input v-model="searchInfo.with" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.type')" prop="type">
          <el-input v-model="searchInfo.type" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.sortable')" prop="sortable">
          <el-select v-model="searchInfo.sortable" clearable :placeholder="$t('common.selectPlaceholder')">
            <el-option key="true" :label="$t('common.yes')" value="true" />
            <el-option key="false" :label="$t('common.no')" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.filter')" prop="filter">
          <el-select v-model="searchInfo.filter" clearable :placeholder="$t('common.selectPlaceholder')">
            <el-option key="true" :label="$t('common.yes')" value="true" />
            <el-option key="false" :label="$t('common.no')" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.defaultFilter')" prop="defaultFilter">
          <el-input v-model="searchInfo.defaultFilter" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.filterList')" prop="filterList">
          <el-input v-model="searchInfo.filterList" :placeholder="$t('common.searchPlaceholder')" />
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset">
            重置
          </el-button>
          <el-button v-if="!showAllQuery" link type="primary" icon="arrow-down" @click="showAllQuery = true">
            展开
          </el-button>
          <el-button v-else link type="primary" icon="arrow-up" @click="showAllQuery = false">
            收起
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">
          新增
        </el-button>
        <el-button
          v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;"
          :disabled="!multipleSelection.length" @click="onDelete"
        >
          删除
        </el-button>
      </div>
      <el-table
        ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange" @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />

        <!-- <el-table-column align="left" label="日期" prop="createdAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.menuId')" prop="menuId"
          width="120"
        />
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.structName')"
          prop="structName" width="120"
        />
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.jsonName')" prop="jsonName"
          width="120"
        />
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.columnName')"
          prop="columnName" width="120"
        />
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.type')" prop="type"
          width="120"
        />
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.with')" prop="with"
          width="120"
        />
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.fixed')" prop="fixed"
          width="120"
        />

        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.sort')" prop="sort"
          width="120"
        />
        <!-- <el-table-column sortable align="left" :label="$t('business.sysTableColumns.columns.isShow')" prop="isShow" width="120"></el-table-column> -->
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.sortable')" prop="sortable"
          width="120"
        >
          <template #default="scope">
            {{ formatBoolean(scope.row.sortable) }}
          </template>
        </el-table-column>
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.filter')" prop="filter"
          width="120"
        >
          <template #default="scope">
            {{ formatBoolean(scope.row.filter) }}
          </template>
        </el-table-column>
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.defaultFilter')"
          prop="defaultFilter" width="120"
        />
        <!-- <el-table-column sortable align="left" :label="$t('business.sysTableColumns.columns.filterList')" prop="filterList" width="120" /> -->
        <el-table-column
          sortable align="left" :label="$t('business.sysTableColumns.columns.note')" prop="note"
          width="120"
        />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>{{ $t('common.view') }}
            </el-button>
            <el-button
              type="primary" link icon="edit" class="table-button"
              @click="updateSysTableColumnsFunc(scope.row)"
            >
              {{ $t('common.edit') }}
            </el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">
              {{ $t('common.delete') }}
            </el-button>
            <el-button type="primary" link icon="copyDocument" @click="copyRow(scope.row)">
              {{ $t('common.copy') }}
            </el-button>

            <el-button
              v-if="scope.row.structName[0] == 'b' || scope.row.structName[0] == 'v'" type="primary" link icon="refresh"
              @click="()=>{openSync(scope.row.structName)
              }"
            >
              {{ $t('common.sync') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer
      v-model="dialogFormVisible" destroy-on-close :size="appStore.drawerSize" :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">
              确 定
            </el-button>
            <el-button @click="closeDialog">
              取 消
            </el-button>
          </div>
        </div>
      </template>

      <el-form ref="elFormRef" :model="formData" label-position="top" :rules="rule" label-width="80px">
        <el-form-item :label="$t('business.sysTableColumns.columns.menuId') + ':'" prop="menuId">
          <el-input
            v-model.number="formData.menuId" :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.menuId') })"
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.structName') + ':'" prop="structName">
          <el-input
            v-model="formData.structName" :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.structName') })"
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.jsonName') + ':'" prop="jsonName">
          <el-input
            v-model="formData.jsonName" :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.jsonName') })"
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.columnName') + ':'" prop="columnName">
          <el-input
            v-model="formData.columnName" :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.columnName') })"
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.note') + ':'" prop="note">
          <el-input
            v-model="formData.note" type="textarea" autosize :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.note') })"
          />
        </el-form-item>
        <!-- <el-form-item :label="$t('business.sysTableColumns.columns.isShow') + ':'"  prop="isShow" >
              <el-switch v-model="formData.isShow" active-color="#13ce66" inactive-color="#ff4949" :active-text="$t('common.yes')" :inactive-text="$t('common.no')" clearable ></el-switch>
            </el-form-item> -->
        <el-form-item :label="$t('business.sysTableColumns.columns.sort') + ':'" prop="sort">
          <el-input
            v-model.number="formData.sort" :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.sort') })"
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.with') + ':'" prop="with">
          <el-input
            v-model="formData.with" :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.with') })"
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.fixed') + ':'" prop="fixed">
          <el-select v-model="formData.fixed" :clearable="true" :placeholder="$t('common.selectPlaceholder')">
            <el-option label="left" value="left" />
            <el-option label="right" value="right" />
            <el-option label="none" value="none" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.type') + ':'" prop="type">
          <el-select v-model="formData.type" :clearable="true" :placeholder="$t('common.selectPlaceholder')">
            <el-option v-for="item in typeOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.sortable') + ':'" prop="sortable">
          <el-switch
            v-model="formData.sortable" active-color="#13ce66" inactive-color="#ff4949"
            :active-text="$t('common.yes')" :inactive-text="$t('common.no')" clearable
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.filter') + ':'" prop="filter">
          <el-switch
            v-model="formData.filter" active-color="#13ce66" inactive-color="#ff4949"
            :active-text="$t('common.yes')" :inactive-text="$t('common.no')" clearable
          />
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.defaultFilter') + ':'" prop="defaultFilter">
          <el-select v-model="formData.defaultFilter" :clearable="true" :placeholder="$t('common.selectPlaceholder')">
            <el-option v-for="item in filterList" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.filterList') + ':'" prop="filterList">
          <el-select
            v-model="formData.filterList" :clearable="true" :placeholder="$t('common.selectPlaceholder')"
            multiple
          >
            <el-option v-for="item in filterList" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('business.sysTableColumns.columns.enum') + ':'" prop="enum" allow-create>
          <el-select
            v-model="formData.enum" :clearable="true" :placeholder="$t('common.selectPlaceholder')" multiple
            allow-create filterable
          >
            <!-- <el-option v-for="item in 100" :key="item" :label="item" :value="item" /> -->
          </el-select>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      v-model="detailShow" destroy-on-close :size="appStore.drawerSize" :show-close="true"
      :before-close="closeDetailShow" :title="$t('common.view')"
    >
      <el-descriptions :column="1" border>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.structName')">
          {{ detailFrom.structName }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.jsonName')">
          {{ detailFrom.jsonName }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.columnName')">
          {{ detailFrom.columnName }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.with')">
          {{ detailFrom.with }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.type')">
          {{ detailFrom.type }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.sortable')">
          {{ detailFrom.sortable }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.filter')">
          {{ detailFrom.filter }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.defaultFilter')">
          {{ detailFrom.defaultFilter }}
        </el-descriptions-item>
        <el-descriptions-item :label="$t('business.sysTableColumns.columns.filterList')">
          {{ detailFrom.filterList }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
    <TcSync ref="tcSyncRef" />
  </div>
</template>

<script setup>
import {
  createSysTableColumns,
  deleteSysTableColumns,
  deleteSysTableColumnsByIds,
  updateSysTableColumns,
  findSysTableColumns,
  getSysTableColumnsList
} from '@/api/system/sysTableColumns'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from '@/pinia'
import { useI18n } from 'vue-i18n'
import TcSync from './tcSync.vue'

const tcSyncRef = ref(null)

const { t } = useI18n()
const filterList = ['=', '!=', '>', '>=', '<', '<=', 'like', 'in', 'not in', 'between', 'not between']

const typeOptions = ref([
  {
    label: '字符串',
    value: 'string'
  },
  {
    label: '富文本',
    value: 'richtext'
  },
  {
    label: '布尔值',
    value: 'bool'
  },
  {
    label: '枚举',
    value: 'enum'
  },
  {
    label: '整型',
    value: 'int32'
  },
  {
    label: '大整数',
    value: 'int64'
  },
  {
    label: '浮点型',
    value: 'float64'
  },
  {
    label: '整型',
    value: 'int'
  },
  {
    label: '时间',
    value: 'datetime'
  },
  {
    label: '日期',
    value: 'date'
  },
  {
    label: 'JSON',
    value: 'json'
  },
  {
    label: '数组',
    value: 'array'
  },
  {
    label: '二进制',
    value: 'binary'
  },
  {
    label: '金额',
    value: 'amount'
  },

])
defineOptions({
  name: 'SysTableColumns'
})
// 按钮权限实例化
const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  structName: '',
  jsonName: '',
  columnName: '',
  type: '',
  sortable: false,
  filter: false,
  defaultFilter: '',
  filterList: '',
})

// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    structName: 'struct_name',
    jsonName: 'json_name',
    columnName: 'column_name',
    with: 'with',
    type: 'type',
    sortable: 'sortable',
    filter: 'filter',
    defaultFilter: 'default_filter',
    filterList: 'filter_list',
    sort: 'sort',
    menuId: 'menu_id',
  }

  let sort = sortMap[prop]
  if (!sort) {
    sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.sortable === '') {
      searchInfo.value.sortable = null
    }
    if (searchInfo.value.filter === '') {
      searchInfo.value.filter = null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  tableData.value = []
  const table = await getSysTableColumnsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm(t('common.confirmDelete'), t('common.hint'), {
    confirmButtonText: t('common.confirm'),
    cancelButtonText: t('common.cancel'),
    type: 'warning'
  }).then(() => {
    deleteSysTableColumnsFunc(row)
  })
}

// 复制行
const copyRow = (row) => {
  formData.value = row
  type.value = 'copy'
  dialogFormVisible.value = true
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('Are you sure you want to delete it??', t('common.hint'), {
    confirmButtonText: t('common.confirm'),
    cancelButtonText: t('common.cancel'),
    type: 'warning'
  }).then(async() => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: 'Please select the data to be deleted.'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deleteSysTableColumnsByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'deleted successfully'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateSysTableColumnsFunc = async(row) => {
  const res = await findSysTableColumns({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteSysTableColumnsFunc = async(row) => {
  const res = await deleteSysTableColumns({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'deleted successfully'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    structName: '',
    jsonName: '',
    columnName: '',
    type: '',
    sortable: false,
    filter: false,
    defaultFilter: '',
    filterList: '',
    sort: 0,
    menuId: '',
    with: 120,
    fixed: 'none',
  }
}
// 弹窗确定
const enterDialog = async() => {
  btnLoading.value = true
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    btnLoading.value = false
    let res
    switch (type.value) {
      case 'create':
        res = await createSysTableColumns(formData.value)
        break
      case 'copy':
        formData.value.ID = undefined
        formData.value.createdAt = undefined
        formData.value.updatedAt = undefined
        res = await createSysTableColumns(formData.value)
        break
      case 'update':
        res = await updateSysTableColumns(formData.value)
        break
      default:
        res = await createSysTableColumns(formData.value)
        break
    }
    btnLoading.value = false
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Creation/Change Successful'
      })
      closeDialog()
      getTableData()
    }
  })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findSysTableColumns({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    if (res.data.filterList) {
      detailFrom.value.filterList = res.data.filterList.split(',')
    }
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

const openSync = (name) => {
  console.log(tcSyncRef.value)

  tcSyncRef?.value?.open(name)
}

</script>

<style></style>
