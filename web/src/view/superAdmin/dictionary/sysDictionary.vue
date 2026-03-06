<template>
  <div>
    <warning-bar
      :title="$t('superAdmin.dictionary.warningMessage')"
    />
    <div class="flex gap-4 p-2">
      <div
        class="flex-none w-52 bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900 rounded p-4"
      >
        <div class="flex justify-between items-center">
          <span class="text font-bold">{{ $t('superAdmin.dictionary.list') }}</span>
          <el-button
            type="primary"
            @click="openDrawer"
          >
            {{ $t('common.add') }}
          </el-button>
        </div>
        <el-scrollbar
          class="mt-4"
          style="height: calc(100vh - 300px)"
        >
          <div
            v-for="dictionary in dictionaryData"
            :key="dictionary.ID"
            class="rounded flex justify-between items-center px-2 py-4 cursor-pointer mt-2 hover:bg-blue-50 dark:hover:bg-blue-900 bg-gray-50 dark:bg-gray-800 gap-4"
            :class="
              selectID === dictionary.ID
                ? 'text-active'
                : 'text-slate-700 dark:text-slate-50'
            "
            @click="toDetail(dictionary)"
          >
            <span class="max-w-[160px] truncate">{{ dictionary.name }}</span>
            <div class="min-w-[40px]">
              <el-icon
                class="text-blue-500"
                @click.stop="updateSysDictionaryFunc(dictionary)"
              >
                <Edit />
              </el-icon>
              <el-icon
                class="ml-2 text-red-500"
                @click="deleteSysDictionaryFunc(dictionary)"
              >
                <Delete />
              </el-icon>
            </div>
          </div>
        </el-scrollbar>
      </div>
      <div
        class="flex-1 bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900"
      >
        <sysDictionaryDetail :sys-dictionary-i-d="selectID" />
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
            type === 'create' ? $t('superAdmin.dictionary.add') : $t('superAdmin.dictionary.edit')
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
          :label="$t('superAdmin.dictionary.nameZh')"
          prop="name"
        >
          <el-input
            v-model="formData.name"
            :placeholder="$t('superAdmin.dictionary.nameZhPlaceholder')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.dictionary.nameEn')"
          prop="type"
        >
          <el-input
            v-model="formData.type"
            :placeholder="$t('superAdmin.dictionary.nameEnPlaceholder')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.dictionary.status')"
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
          :label="$t('superAdmin.dictionary.description')"
          prop="desc"
        >
          <el-input
            v-model="formData.desc"
            :placeholder="$t('superAdmin.dictionary.descriptionPlaceholder')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createSysDictionary,
    deleteSysDictionary,
    updateSysDictionary,
    findSysDictionary,
    getSysDictionaryList
  } from '@/api/sysDictionary'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { useI18n } from 'vue-i18n'
  import sysDictionaryDetail from './sysDictionaryDetail.vue'
  import { Edit } from '@element-plus/icons-vue'
  import { useAppStore } from "@/pinia"

  const { t } = useI18n()
  const appStore = useAppStore()

  const selectID = ref(0)

  const formData = ref({
    name: null,
    type: null,
    status: true,
    desc: null
  })
  const rules = ref({
    name: [
      {
        required: true,
        message: t('superAdmin.dictionary.nameZhRequired'),
        trigger: 'blur'
      }
    ],
    type: [
      {
        required: true,
        message: t('superAdmin.dictionary.nameEnRequired'),
        trigger: 'blur'
      }
    ],
    desc: [
      {
        required: true,
        message: t('superAdmin.dictionary.descriptionRequired'),
        trigger: 'blur'
      }
    ]
  })

  const dictionaryData = ref([])

  // 查询
  const getTableData = async () => {
    const res = await getSysDictionaryList()
    if (res.code === 0) {
      dictionaryData.value = res.data
      selectID.value = res.data[0].ID
    }
  }

  getTableData()

  const toDetail = (row) => {
    selectID.value = row.ID
  }

  const drawerFormVisible = ref(false)
  const type = ref('')
  const updateSysDictionaryFunc = async (row) => {
    const res = await findSysDictionary({ ID: row.ID, status: row.status })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data.resysDictionary
      drawerFormVisible.value = true
    }
  }
  const closeDrawer = () => {
    drawerFormVisible.value = false
    formData.value = {
      name: null,
      type: null,
      status: true,
      desc: null
    }
  }
  const deleteSysDictionaryFunc = async (row) => {
    ElMessageBox.confirm(t('superAdmin.dictionary.deleteConfirm'), t('common.warning'), {
      confirmButtonText: t('common.confirm'),
      cancelButtonText: t('common.cancel'),
      type: 'warning'
    }).then(async () => {
      const res = await deleteSysDictionary({ ID: row.ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('superAdmin.dictionary.deleteSuccess')
        })
        getTableData()
      }
    })
  }

  const drawerForm = ref(null)
  const enterDrawer = async () => {
    drawerForm.value.validate(async (valid) => {
      if (!valid) return
      let res
      switch (type.value) {
        case 'create':
          res = await createSysDictionary(formData.value)
          break
        case 'update':
          res = await updateSysDictionary(formData.value)
          break
        default:
          res = await createSysDictionary(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage.success(t('superAdmin.dictionary.operationSuccess'))
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
</script>

<style>
  .dict-box {
    height: calc(100vh - 240px);
  }

  .active {
    background-color: var(--el-color-primary) !important;
    color: #fff;
  }
</style>
