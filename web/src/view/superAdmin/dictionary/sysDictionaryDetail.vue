<template>
  <div>
    <div class="hab-table-box">
      <div class="hab-btn-list justify-between">
        <span class="text font-bold">{{ $t('superAdmin.dictionary.detail.title') }}</span>
        <el-button
          type="primary"
          icon="plus"
          @click="openDrawer"
        >
          {{ $t('superAdmin.dictionary.detail.add') }}
        </el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.dictionary.detail.date')"
          width="180"
        >
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          :label="$t('superAdmin.dictionary.detail.label')"
          prop="label"
        />

        <el-table-column
          align="left"
          :label="$t('superAdmin.dictionary.detail.value')"
          prop="value"
        />

        <el-table-column
          align="left"
          :label="$t('superAdmin.dictionary.detail.extend')"
          prop="extend"
        />

        <el-table-column
          align="left"
          :label="$t('superAdmin.dictionary.detail.status')"
          prop="status"
          width="120"
        >
          <template #default="scope">
            {{ formatBoolean(scope.row.status) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          :label="$t('superAdmin.dictionary.detail.sort')"
          prop="sort"
          width="120"
        />

        <el-table-column
          align="left"
          :label="$t('superAdmin.dictionary.detail.operations')"
          :min-width="appStore.operateMinWith"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updateSysDictionaryDetailFunc(scope.row)"
            >
              {{ $t('common.edit') }}
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteSysDictionaryDetailFunc(scope.row)"
            >
              {{ $t('common.delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="hab-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer
      v-model="drawerFormVisible"
      :size="appStore.drawerSize"
      :show-close="false"
      :before-close="closeDrawer"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            type === 'create' ? $t('superAdmin.dictionary.detail.add') : $t('superAdmin.dictionary.detail.edit')
          }}</span>
          <div>
            <el-button @click="closeDrawer">
              {{ $t('common.cancel') }}
            </el-button>
            <el-button
              type="primary"
              @click="enterDrawer"
            >
              {{ $t('common.confirm') }}
            </el-button>
          </div>
        </div>
      </template>
      <el-form
        ref="drawerForm"
        :model="formData"
        :rules="rules"
        label-width="110px"
      >
        <el-form-item
          :label="$t('superAdmin.dictionary.detail.label')"
          prop="label"
        >
          <el-input
            v-model="formData.label"
            :placeholder="$t('superAdmin.dictionary.detail.labelPlaceholder')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.dictionary.detail.value')"
          prop="value"
        >
          <el-input
            v-model="formData.value"
            :placeholder="$t('superAdmin.dictionary.detail.valuePlaceholder')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.dictionary.detail.extend')"
          prop="extend"
        >
          <el-input
            v-model="formData.extend"
            :placeholder="$t('superAdmin.dictionary.detail.extendPlaceholder')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.dictionary.detail.status')"
          prop="status"
          required
        >
          <el-switch
            v-model="formData.status"
            :active-text="$t('superAdmin.dictionary.enabled')"
            :inactive-text="$t('superAdmin.dictionary.disabled')"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.dictionary.detail.sort')"
          prop="sort"
        >
          <el-input-number
            v-model.number="formData.sort"
            :placeholder="$t('superAdmin.dictionary.detail.sortPlaceholder')"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createSysDictionaryDetail,
  deleteSysDictionaryDetail,
  updateSysDictionaryDetail,
  findSysDictionaryDetail,
  getSysDictionaryDetailList
} from '@/api/sysDictionaryDetail' // 此处请自行替换地址
import { ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatBoolean, formatDate } from '@/utils/format'
import { useAppStore } from '@/pinia'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const appStore = useAppStore()

const props = defineProps({
  sysDictionaryID: {
    type: Number,
    default: 0
  }
})

const formData = ref({
  label: null,
  value: null,
  status: true,
  sort: null
})
const rules = ref({
  label: [
    {
      required: true,
      message: t('superAdmin.dictionary.detail.labelRequired'),
      trigger: 'blur'
    }
  ],
  value: [
    {
      required: true,
      message: t('superAdmin.dictionary.detail.valueRequired'),
      trigger: 'blur'
    }
  ],
  sort: [
    {
      required: true,
      message: t('superAdmin.dictionary.detail.sortRequired'),
      trigger: 'blur'
    }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  if (!props.sysDictionaryID) return
  tableData.value = []
  const table = await getSysDictionaryDetailList({
    page: page.value,
    pageSize: pageSize.value,
    sysDictionaryID: props.sysDictionaryID
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const type = ref('')
const drawerFormVisible = ref(false)
const updateSysDictionaryDetailFunc = async(row) => {
  drawerForm.value && drawerForm.value.clearValidate()
  const res = await findSysDictionaryDetail({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reSysDictionaryDetail
    drawerFormVisible.value = true
  }
}

const closeDrawer = () => {
  drawerFormVisible.value = false
  formData.value = {
    label: null,
    value: null,
    status: true,
    sort: null,
    sysDictionaryID: props.sysDictionaryID
  }
}
const deleteSysDictionaryDetailFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteSysDictionaryDetail({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

const drawerForm = ref(null)
const enterDrawer = async() => {
  drawerForm.value.validate(async(valid) => {
    formData.value.sysDictionaryID = props.sysDictionaryID
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysDictionaryDetail(formData.value)
        break
      case 'update':
        res = await updateSysDictionaryDetail(formData.value)
        break
      default:
        res = await createSysDictionaryDetail(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDrawer()
      getTableData()
    }
  })
}
const openDrawer = () => {
  type.value = 'create'
  drawerForm.value && drawerForm.value.clearValidate()
  drawerFormVisible.value = true
}

watch(
  () => props.sysDictionaryID,
  () => {
    getTableData()
  }
)
</script>

<style></style>
