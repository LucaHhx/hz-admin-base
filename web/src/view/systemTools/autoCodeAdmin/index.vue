<template>
  <div>
    <div class="hab-table-box">
      <div class="hab-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="goAutoCode(null)"
        >
          {{ $t('common.add') }}
        </el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="id"
          width="60"
          prop="ID"
        />
        <el-table-column
          align="left"
          :label="$t('autoCode.admin.date')"
          width="180"
        >
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="$t('autoCode.admin.structName')"
          min-width="150"
          prop="structName"
        />
        <el-table-column
          align="left"
          :label="$t('autoCode.admin.structDesc')"
          min-width="150"
          prop="description"
        />
        <el-table-column
          align="left"
          :label="$t('autoCode.admin.tableName')"
          min-width="150"
          prop="tableName"
        />
        <el-table-column
          align="left"
          :label="$t('autoCode.admin.rollbackMark')"
          min-width="150"
          prop="flag"
        >
          <template #default="scope">
            <el-tag
              v-if="scope.row.flag"
              type="danger"
              effect="dark"
            >
              {{ $t('autoCode.admin.hasRollback') }}
            </el-tag>
            <el-tag
              v-else
              type="success"
              effect="dark"
            >
              {{ $t('autoCode.admin.notRollback') }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          min-width="240"
        >
          <template #default="scope">
            <div>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="addFuncBtn(scope.row)"
              >
                {{ $t('autoCode.admin.operations.addMethod') }}
              </el-button>
              <el-button
                type="primary"
                link
                @click="goAutoCode(scope.row, 1)"
              >
                {{ $t('autoCode.admin.operations.addField') }}
              </el-button>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="openDialog(scope.row)"
              >
                {{ $t('autoCode.admin.operations.rollback') }}
              </el-button>
              <el-button
                type="primary"
                link
                @click="goAutoCode(scope.row)"
              >
                {{ $t('autoCode.admin.operations.reuse') }}
              </el-button>
              <el-button
                type="primary"
                link
                @click="deleteRow(scope.row)"
              >
                {{ $t('autoCode.admin.operations.delete') }}
              </el-button>
            </div>
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
    <el-dialog
      v-model="dialogFormVisible"
      :title="dialogFormTitle"
      :before-close="closeDialog"
      width="600px"
    >
      <el-form
        :inline="true"
        :model="formData"
        label-width="80px"
      >
        <el-form-item :label="$t('autoCode.admin.rollbackDialog.options')">
          <el-checkbox
            v-model="formData.deleteApi"
            :label="$t('autoCode.admin.rollbackDialog.deleteApi')"
          />
          <el-checkbox
            v-model="formData.deleteMenu"
            :label="$t('autoCode.admin.rollbackDialog.deleteMenu')"
          />
          <el-checkbox
            v-model="formData.deleteTable"
            :label="$t('autoCode.admin.rollbackDialog.deleteTable')"
            @change="deleteTableCheck"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">
            {{ $t('common.cancel') }}
          </el-button>
          <el-popconfirm
            :title="$t('autoCode.admin.rollbackDialog.confirmDelete')"
            @confirm="enterDialog"
          >
            <template #reference>
              <el-button type="primary">
                {{ $t('common.confirm') }}
              </el-button>
            </template>
          </el-popconfirm>
        </div>
      </template>
    </el-dialog>

    <el-drawer
      v-model="funcFlag"
      size="60%"
      :show-close="false"
      :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ $t('autoCode.admin.functionDrawer.operationBar') }}</span>
          <div>
            <el-button
              type="primary"
              :loading="aiLoading"
              @click="runFunc"
            >
              {{ $t('autoCode.admin.functionDrawer.generate') }}
            </el-button>
            <el-button
              type="primary"
              :loading="aiLoading"
              @click="closeFunc"
            >
              {{ $t('autoCode.admin.functionDrawer.cancel') }}
            </el-button>
          </div>
        </div>
      </template>
      <div class="">
        <el-form
          v-loading="aiLoading"
          label-position="top"
          :element-loading-text="$t('autoCode.admin.functionDrawer.aiThinking')"
          :model="autoFunc"
          label-width="80px"
        >
          <el-row :gutter="12">
            <el-col :span="8">
              <el-form-item :label="$t('autoCode.admin.functionDrawer.packageName')">
                <el-input
                  v-model="autoFunc.package"
                  :placeholder="$t('autoCode.admin.functionDrawer.packageNamePlaceholder')"
                  disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item :label="$t('autoCode.admin.functionDrawer.structName')">
                <el-input
                  v-model="autoFunc.structName"
                  :placeholder="$t('autoCode.admin.functionDrawer.structNamePlaceholder')"
                  disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item :label="$t('autoCode.admin.functionDrawer.frontendFileName')">
                <el-input
                  v-model="autoFunc.packageName"
                  :placeholder="$t('autoCode.admin.functionDrawer.frontendFileNamePlaceholder')"
                  disabled
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="12">
            <el-col :span="8">
              <el-form-item :label="$t('autoCode.admin.functionDrawer.backendFileName')">
                <el-input
                  v-model="autoFunc.humpPackageName"
                  :placeholder="$t('autoCode.admin.functionDrawer.backendFileNamePlaceholder')"
                  disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item :label="$t('autoCode.admin.functionDrawer.description')">
                <el-input
                  v-model="autoFunc.description"
                  :placeholder="$t('autoCode.admin.functionDrawer.descriptionPlaceholder')"
                  disabled
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item :label="$t('autoCode.admin.functionDrawer.abbreviation')">
                <el-input
                  v-model="autoFunc.abbreviation"
                  :placeholder="$t('autoCode.admin.functionDrawer.abbreviationPlaceholder')"
                  disabled
                />
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item :label="$t('autoCode.admin.functionDrawer.useAI')">
            <el-switch v-model="autoFunc.isAi" />
            <span class="text-sm text-red-600 p-2">{{ $t('autoCode.admin.functionDrawer.aiWarning') }}</span>
          </el-form-item>
          <template v-if="autoFunc.isAi">
            <el-form-item :label="$t('autoCode.admin.functionDrawer.aiHelp')">
              <div class="relative w-full">
                <el-input
                  v-model="autoFunc.prompt"
                  type="textarea"
                  :placeholder="$t('autoCode.admin.functionDrawer.aiHelpPlaceholder')"
                  :rows="5"
                  @input="autoFunc.router = autoFunc.router.replace(/\//g, '')"
                />
                <el-button
                  type="primary"
                  class="absolute right-2 bottom-2"
                  @click="aiAddFunc"
                >
                  <ai-hab />{{ $t('autoCode.admin.functionDrawer.aiWrite') }}
                </el-button>
              </div>
            </el-form-item>
            <el-form-item :label="$t('autoCode.admin.functionDrawer.apiMethod')">
              <v-ace-editor
                v-model:value="autoFunc.apiFunc"
                lang="golang"
                theme="github_dark"
                class="h-80 w-full"
              />
            </el-form-item>
            <el-form-item :label="$t('autoCode.admin.functionDrawer.serverMethod')">
              <v-ace-editor
                v-model:value="autoFunc.serverFunc"
                lang="golang"
                theme="github_dark"
                class="h-80 w-full"
              />
            </el-form-item>
            <el-form-item :label="$t('autoCode.admin.functionDrawer.frontendJSAPI')">
              <v-ace-editor
                v-model:value="autoFunc.jsFunc"
                lang="javascript"
                theme="github_dark"
                class="h-80 w-full"
              />
            </el-form-item>
          </template>

          <el-form-item :label="$t('autoCode.admin.functionDrawer.methodIntro')">
            <div class="flex w-full gap-2">
              <el-input
                v-model="autoFunc.funcDesc"
                class="flex-1"
                :placeholder="$t('autoCode.admin.functionDrawer.methodIntroPlaceholder')"
              />
              <el-button
                type="primary"
                @click="autoComplete"
              >
                <ai-hab />{{ $t('autoCode.admin.functionDrawer.complete') }}
              </el-button>
            </div>
          </el-form-item>
          <el-form-item :label="$t('autoCode.admin.functionDrawer.methodName')">
            <el-input
              v-model="autoFunc.funcName"
              :placeholder="$t('autoCode.admin.functionDrawer.methodNamePlaceholder')"
              @blur="autoFunc.funcName = toUpperCase(autoFunc.funcName)"
            />
          </el-form-item>
          <el-form-item :label="$t('autoCode.admin.functionDrawer.method')">
            <el-select
              v-model="autoFunc.method"
              :placeholder="$t('autoCode.admin.functionDrawer.methodPlaceholder')"
            >
              <el-option
                v-for="item in ['GET', 'POST', 'PUT', 'DELETE']"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('autoCode.admin.functionDrawer.needAuth')">
            <el-switch
              v-model="autoFunc.isAuth"
              :active-text="$t('autoCode.admin.functionDrawer.yes')"
              :inactive-text="$t('autoCode.admin.functionDrawer.no')"
            />
          </el-form-item>
          <el-form-item :label="$t('autoCode.admin.functionDrawer.routerPath')">
            <el-input
              v-model="autoFunc.router"
              :placeholder="$t('autoCode.admin.functionDrawer.routerPathPlaceholder')"
              @input="autoFunc.router = autoFunc.router.replace(/\//g, '')"
            />
            <div>
              {{ $t('autoCode.admin.functionDrawer.apiPath') }}: [{{ autoFunc.method }}] /{{ autoFunc.abbreviation }}/{{
                autoFunc.router
              }}
            </div>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getSysHistory,
  rollback,
  delSysHistory,
  addFunc,
  butler
} from '@/api/autoCode.js'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { formatDate } from '@/utils/format'
import { toUpperCase } from '@/utils/stringFun'
import { useI18n } from 'vue-i18n'

import { VAceEditor } from 'vue3-ace-editor'
import 'ace-builds/src-noconflict/mode-javascript'
import 'ace-builds/src-noconflict/mode-golang'
import 'ace-builds/src-noconflict/theme-github_dark'

defineOptions({
  name: 'AutoCodeAdmin'
})

const aiLoading = ref(false)

const formData = ref({
  id: undefined,
  deleteApi: true,
  deleteMenu: true,
  deleteTable: false
})

const router = useRouter()
const dialogFormVisible = ref(false)
const dialogFormTitle = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const activeInfo = ref('')

const autoFunc = ref({
  package: '',
  funcName: '',
  structName: '',
  packageName: '',
  description: '',
  abbreviation: '',
  humpPackageName: '',
  businessDB: '',
  method: '',
  funcDesc: '',
  isAuth: false,
  isAi: false,
  apiFunc: '',
  serverFunc: '',
  jsFunc: ''
})

const { t } = useI18n()

const addFuncBtn = (row) => {
  const req = JSON.parse(row.request)
  activeInfo.value = row.request
  autoFunc.value.package = req.package
  autoFunc.value.structName = req.structName
  autoFunc.value.packageName = req.packageName
  autoFunc.value.description = req.description
  autoFunc.value.abbreviation = req.abbreviation
  autoFunc.value.humpPackageName = req.humpPackageName
  autoFunc.value.businessDB = req.businessDB
  autoFunc.value.method = ''
  autoFunc.value.funcName = ''
  autoFunc.value.router = ''
  autoFunc.value.funcDesc = ''
  autoFunc.value.isAuth = false
  autoFunc.value.isAi = false
  autoFunc.value.apiFunc = ''
  autoFunc.value.serverFunc = ''
  autoFunc.value.jsFunc = ''
  funcFlag.value = true
}

const funcFlag = ref(false)

const closeFunc = () => {
  funcFlag.value = false
}

const runFunc = async() => {
  if (!autoFunc.value.funcName) {
    ElMessage.error(t('autoCode.admin.messages.enterMethodName'))
    return
  }
  if (!autoFunc.value.method) {
    ElMessage.error(t('autoCode.admin.messages.selectMethod'))
    return
  }
  if (!autoFunc.value.router) {
    ElMessage.error(t('autoCode.admin.messages.enterRouter'))
    return
  }
  if (!autoFunc.value.funcDesc) {
    ElMessage.error(t('autoCode.admin.messages.enterMethodIntro'))
    return
  }

  if (autoFunc.value.isAi) {
    if (
      !autoFunc.value.apiFunc ||
        !autoFunc.value.serverFunc ||
        !autoFunc.value.jsFunc
    ) {
      ElMessage.error(t('autoCode.admin.messages.completeAICode'))
      return
    }
  }

  const res = await addFunc(autoFunc.value)
  if (res.code === 0) {
    ElMessage.success(t('autoCode.admin.messages.addMethodSuccess'))
    closeFunc()
  }
}

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
  tableData.value = []
  const table = await getSysHistory({
    page: page.value,
    pageSize: pageSize.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const deleteRow = async(row) => {
  ElMessageBox.confirm(t('autoCode.admin.messages.deleteConfirm'), t('common.warning'), {
    confirmButtonText: t('common.confirm'),
    cancelButtonText: t('common.cancel'),
    type: 'warning'
  }).then(async() => {
    const res = await delSysHistory({ id: Number(row.ID) })
    if (res.code === 0) {
      ElMessage.success(t('autoCode.admin.messages.deleteSuccess'))
      getTableData()
    }
  })
}

// 打开弹窗
const openDialog = (row) => {
  dialogFormTitle.value = t('autoCode.admin.rollbackDialog.title') + row.structName
  formData.value.id = row.ID
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    id: undefined,
    deleteApi: true,
    deleteMenu: true,
    deleteTable: false
  }
}

// 确认删除表
const deleteTableCheck = (flag) => {
  if (flag) {
    ElMessageBox.confirm(
      t('autoCode.admin.rollbackDialog.deleteTableWarning'),
      t('common.warning'),
      {
        closeOnClickModal: false,
        distinguishCancelAndClose: true,
        confirmButtonText: t('common.confirm'),
        cancelButtonText: t('common.cancel'),
        type: 'warning'
      }
    )
      .then(() => {
        ElMessageBox.confirm(
          t('autoCode.admin.rollbackDialog.deleteTableConfirm'),
          t('common.warning'),
          {
            closeOnClickModal: false,
            distinguishCancelAndClose: true,
            confirmButtonText: t('common.confirm'),
            cancelButtonText: t('common.cancel'),
            type: 'warning'
          }
        ).catch(() => {
          formData.value.deleteTable = false
        })
      })
      .catch(() => {
        formData.value.deleteTable = false
      })
  }
}

const enterDialog = async() => {
  const res = await rollback(formData.value)
  if (res.code === 0) {
    ElMessage.success('回滚成功')
    getTableData()
  }
}

const goAutoCode = (row, isAdd) => {
  if (row) {
    router.push({
      name: 'autoCodeEdit',
      params: {
        id: row.ID
      },
      query: {
        isAdd: isAdd
      }
    })
  } else {
    router.push({ name: 'autoCode' })
  }
}

const aiAddFunc = async() => {
  aiLoading.value = true
  autoFunc.value.apiFunc = ''
  autoFunc.value.serverFunc = ''
  autoFunc.value.jsFunc = ''

  if (!autoFunc.value.prompt) {
    ElMessage.error(t('autoCode.admin.messages.enterPrompt'))
    return
  }

  const res = await addFunc({ ...autoFunc.value, isPreview: true })
  if (res.code !== 0) {
    aiLoading.value = false
    ElMessage.error(res.msg)
    return
  }

  const aiRes = await butler({
    structInfo: activeInfo.value,
    template: JSON.stringify(res.data),
    prompt: autoFunc.value.prompt,
    command: 'addFunc'
  })
  aiLoading.value = false
  if (aiRes.code === 0) {
    try {
      const aiData = JSON.parse(aiRes.data)
      autoFunc.value.apiFunc = aiData.api
      autoFunc.value.serverFunc = aiData.server
      autoFunc.value.jsFunc = aiData.js
      autoFunc.value.method = aiData.method
      autoFunc.value.funcName = aiData.funcName
      const routerArr = aiData.router.split('/')
      autoFunc.value.router = routerArr[routerArr.length - 1]
      autoFunc.value.funcDesc = autoFunc.value.prompt
    } catch {
      ElMessage.error(t('autoCode.admin.messages.aiError'))
    }
  }
}

const autoComplete = async() => {
  aiLoading.value = true
  const aiRes = await butler({
    prompt: autoFunc.value.funcDesc,
    command: 'autoCompleteFunc'
  })
  aiLoading.value = false
  if (aiRes.code === 0) {
    try {
      const aiData = JSON.parse(aiRes.data)
      autoFunc.value.method = aiData.method
      autoFunc.value.funcName = aiData.funcName
      autoFunc.value.router = aiData.router
      autoFunc.value.prompt = autoFunc.value.funcDesc
    } catch {
      ElMessage.error(t('autoCode.admin.messages.aiError2'))
    }
  }
}
</script>
