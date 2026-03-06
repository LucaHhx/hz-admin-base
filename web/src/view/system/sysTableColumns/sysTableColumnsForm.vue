<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item
          :label="$t('business.sysTableColumns.columns.structName') + ':'"
          prop="structName"
        >
          <el-input
            v-model="formData.structName"
            :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.structName') })"
          />
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.jsonName') + ':'"
          prop="jsonName"
        >
          <el-input
            v-model="formData.jsonName"
            :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.jsonName') })"
          />
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.columnName') + ':'"
          prop="columnName"
        >
          <el-input
            v-model="formData.columnName"
            :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.columnName') })"
          />
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.sort') + ':'"
          prop="sort"
        >
          <el-input
            v-model.number="formData.sort"
            :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.sort') })"
          />
        </el-form-item>

        <el-form-item
          :label="$t('business.sysTableColumns.columns.with') + ':'"
          prop="with"
        />
        <el-form-item
          :label="$t('business.sysTableColumns.columns.type') + ':'"
          prop="type"
        >
          <el-select
            v-model="formData.type"
            value-key=""
            :placeholder="$t('common.selectPlaceholder')"
            clearable
            filterable
          >
            <el-option
              v-for="item in typeOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.sortable') + ':'"
          prop="sortable"
        >
          <el-switch
            v-model="formData.sortable"
            active-color="#13ce66"
            inactive-color="#ff4949"
            :active-text="$t('common.yes')"
            :inactive-text="$t('common.no')"
            clearable
          />
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.filter') + ':'"
          prop="filter"
        >
          <el-switch
            v-model="formData.filter"
            active-color="#13ce66"
            inactive-color="#ff4949"
            :active-text="$t('common.yes')"
            :inactive-text="$t('common.no')"
            clearable
          />
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.defaultFilter') + ':'"
          prop="defaultFilter"
        >
          <el-input
            v-model="formData.defaultFilter"
            :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.defaultFilter') })"
          />
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.filterList') + ':'"
          prop="filterList"
        >
          <el-input
            v-model="formData.filterList"
            :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.filterList') })"
          />
        </el-form-item>
        <el-form-item
          :label="$t('business.sysTableColumns.columns.note') + ':'"
          prop="note"
        >
          <el-input
            v-model="formData.note"
            type="textarea"
            autosize
            :clearable="true"
            :placeholder="$t('common.inputPlaceholder', { field: $t('business.sysTableColumns.columns.note') })"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            :loading="btnLoading"
            type="primary"
            @click="save"
          >
            {{ $t('common.save') }}
          </el-button>
          <el-button
            type="primary"
            @click="back"
          >
            {{ $t('common.back') }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSysTableColumns,
  updateSysTableColumns,
  findSysTableColumns
} from '@/api/system/sysTableColumns'

import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

// const typeOptions = ref([
//   { label: '字符串', value: 'string' },
//   { label: '整数', value: 'integer' },
//   { label: '浮点数', value: 'float' },
//   { label: '布尔值', value: 'boolean' },
// ])
const type = ref('')
const formData = ref({
  structName: '',
  jsonName: '',
  columnName: '',
  type: '',
  sortable: false,
  filter: false,
  defaultFilter: '',
  filterList: [],
})
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findSysTableColumns({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data
      formData.value
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
  btnLoading.value = true
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    btnLoading.value = false
    let res
    const data = formData.value.map(item => {
      if (item.filterList) {
        item.filterList = item.filterList.join(',')
      }
      return item
    })
    switch (type.value) {
      case 'create':
        res = await createSysTableColumns(data)
        break
      case 'update':
        res = await updateSysTableColumns(data)
        break
      default:
        res = await createSysTableColumns(data)
        break
    }
    btnLoading.value = false
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
