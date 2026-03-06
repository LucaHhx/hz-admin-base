<template>
  <div>
    <div class="hab-table-box">
      <div class="hab-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addMenu(0)"
        >
          {{ $t('superAdmin.menu.addRoot') }}
        </el-button>
      </div>

      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <el-table-column
          align="left"
          label="ID"
          min-width="100"
          prop="ID"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.displayName')"
          min-width="120"
          prop="authorityName"
        >
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
            <br>
            <span>{{ $t('menu.'+scope.row.meta.title) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.icon')"
          min-width="140"
          prop="authorityName"
        >
          <template #default="scope">
            <div
              v-if="scope.row.meta.icon"
              class="icon-column"
            >
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.routeName')"
          show-overflow-tooltip
          min-width="160"
          prop="name"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.routePath')"
          show-overflow-tooltip
          min-width="160"
          prop="path"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.hidden')"
          min-width="100"
          prop="hidden"
        >
          <template #default="scope">
            <span>{{ scope.row.hidden ? $t('superAdmin.menu.yes') : $t('superAdmin.menu.no') }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.parentNode')"
          min-width="90"
          prop="parentId"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.sort')"
          min-width="70"
          prop="sort"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.menu.filePath')"
          min-width="360"
          prop="component"
        />
        <el-table-column
          align="left"
          fixed="right"
          :label="$t('superAdmin.menu.operations')"
          :min-width="appStore.operateMinWith"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >
              {{ $t('superAdmin.menu.addSubmenu') }}
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >
              {{ $t('common.edit') }}
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >
              {{ $t('common.delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-drawer
      v-model="dialogFormVisible"
      :size="appStore.drawerSize"
      :before-close="handleClose"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">
            {{ dialogTitle }}
          </span>
          <div>
            <el-button @click="closeDialog">
              {{ $t('common.cancel') }}
            </el-button>
            <el-button
              type="primary"
              @click="enterDialog"
            >
              {{ $t('common.confirm') }}
            </el-button>
          </div>
        </div>
      </template>

      <warning-bar :title="$t('superAdmin.menu.warningMessage')" />
      <el-form
        v-if="dialogFormVisible"
        ref="menuForm"
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
      >
        <el-row class="w-full">
          <el-col :span="16">
            <el-form-item
              :label="$t('superAdmin.menu.filePath')"
              prop="component"
            >
              <components-cascader
                :component="form.component"
                @change="fmtComponent"
              />
              <span style="font-size: 12px; margin-right: 12px">
                {{ $t('superAdmin.menu.routerViewHint') }}
              </span>
              <el-button
                style="margin-top: 4px"
                @click="form.component = 'view/routerHolder.vue'"
              >
                {{ $t('superAdmin.menu.setRouterView') }}
              </el-button>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item
              :label="$t('superAdmin.menu.displayName')"
              prop="meta.title"
            >
              <el-input
                v-model="form.meta.title"
                autocomplete="off"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="w-full">
          <el-col :span="8">
            <el-form-item
              :label="$t('superAdmin.menu.routeName')"
              prop="path"
            >
              <el-input
                v-model="form.name"
                autocomplete="off"
                :placeholder="$t('superAdmin.menu.routeNamePlaceholder')"
                @change="changeName"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item prop="path">
              <template #label>
                <span style="display: inline-flex; align-items: center">
                  <span>{{ $t('superAdmin.menu.routePath') }}</span>
                  <el-checkbox
                    v-model="checkFlag"
                    style="margin-left: 12px; height: auto"
                  >
                    {{ $t('superAdmin.menu.addParams') }}
                  </el-checkbox>
                </span>
              </template>

              <el-input
                v-model="form.path"
                :disabled="!checkFlag"
                autocomplete="off"
                :placeholder="$t('superAdmin.menu.routePathPlaceholder')"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item :label="$t('superAdmin.menu.hidden')">
              <el-select
                v-model="form.hidden"
                style="width: 100%"
                :placeholder="$t('superAdmin.menu.hiddenPlaceholder')"
              >
                <el-option
                  :value="false"
                  :label="$t('superAdmin.menu.no')"
                />
                <el-option
                  :value="true"
                  :label="$t('superAdmin.menu.yes')"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="w-full">
          <el-col :span="8">
            <el-form-item :label="$t('superAdmin.menu.parentId')">
              <el-cascader
                v-model="form.parentId"
                style="width: 100%"
                :disabled="!isEdit"
                :options="menuOption"
                :props="{
                  checkStrictly: true,
                  label: 'title',
                  value: 'ID',
                  disabled: 'disabled',
                  emitPath: false
                }"
                :show-all-levels="false"
                filterable
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item
              :label="$t('superAdmin.menu.icon')"
              prop="meta.icon"
            >
              <icon v-model="form.meta.icon" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item
              :label="$t('superAdmin.menu.sort')"
              prop="sort"
            >
              <el-input
                v-model.number="form.sort"
                autocomplete="off"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="w-full">
          <el-col :span="8">
            <el-form-item prop="meta.activeName">
              <template #label>
                <div>
                  <span>{{ $t('superAdmin.menu.activeMenu') }}</span>
                  <el-tooltip
                    :content="$t('superAdmin.menu.activeMenuHint')"
                    placement="top"
                    effect="light"
                  >
                    <el-icon><QuestionFilled /></el-icon>
                  </el-tooltip>
                </div>
              </template>
              <el-input
                v-model="form.meta.activeName"
                :placeholder="form.name"
                autocomplete="off"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item
              label="KeepAlive"
              prop="meta.keepAlive"
            >
              <el-select
                v-model="form.meta.keepAlive"
                style="width: 100%"
                placeholder="是否keepAlive缓存页面"
              >
                <el-option
                  :value="false"
                  label="否"
                />
                <el-option
                  :value="true"
                  label="是"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item
              label="CloseTab"
              prop="meta.closeTab"
            >
              <el-select
                v-model="form.meta.closeTab"
                style="width: 100%"
                placeholder="是否自动关闭tab"
              >
                <el-option
                  :value="false"
                  label="否"
                />
                <el-option
                  :value="true"
                  label="是"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="w-full">
          <el-col :span="8">
            <el-form-item>
              <template #label>
                <div>
                  <span> 是否为基础页面 </span>
                  <el-tooltip
                    content="此项选择为是，则不会展示左侧菜单以及顶部信息。"
                    placement="top"
                    effect="light"
                  >
                    <el-icon><QuestionFilled /></el-icon>
                  </el-tooltip>
                </div>
              </template>

              <el-select
                v-model="form.meta.defaultMenu"
                style="width: 100%"
                placeholder="是否为基础页面"
              >
                <el-option
                  :value="false"
                  label="否"
                />
                <el-option
                  :value="true"
                  label="是"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item>
              <template #label>
                <div>
                  <span> 路由切换动画 </span>
                  <el-tooltip
                    content="如果设置了路由切换动画，在本路由下的动画优先级高于全局动画切换优先级。"
                    placement="top"
                    effect="light"
                  >
                    <el-icon><QuestionFilled /></el-icon>
                  </el-tooltip>
                </div>
              </template>

              <el-select
                v-model="form.meta.transitionType"
                style="width: 100%"
                placeholder="跟随全局"
                clearable
              >
                <el-option
                  value="fade"
                  label="淡入淡出"
                />
                <el-option
                  value="slide"
                  label="滑动"
                />
                <el-option
                  value="zoom"
                  label="缩放"
                />
                <el-option
                  value="none"
                  label="无动画"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div>
        <div class="flex items-center gap-2">
          <el-button
            type="primary"
            icon="edit"
            @click="addParameter(form)"
          >
            {{ $t('superAdmin.menu.addMenuParam') }}
          </el-button>
        </div>
        <el-table
          :data="form.parameters"
          style="width: 100%; margin-top: 12px"
        >
          <el-table-column
            align="left"
            prop="type"
            :label="$t('superAdmin.menuTable.paramType')"
            width="180"
          >
            <template #default="scope">
              <el-select
                v-model="scope.row.type"
                :placeholder="$t('superAdmin.menuTable.selectType')"
              >
                <el-option
                  key="query"
                  value="query"
                  label="query"
                />
                <el-option
                  key="params"
                  value="params"
                  label="params"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="key"
            :label="$t('superAdmin.menuTable.paramKey')"
            width="180"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.key" />
              </div>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="value"
            :label="$t('superAdmin.menuTable.paramValue')"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.value" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  icon="delete"
                  @click="deleteParameter(form.parameters, scope.$index)"
                >
                  {{ $t('common.delete') }}
                </el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <div class="flex items-center gap-2 mt-3">
          <el-button
            type="primary"
            icon="edit"
            @click="addBtn(form)"
          >
            {{ $t('superAdmin.menu.addBtn') }}
          </el-button>
          <el-icon
            class="cursor-pointer"
            @click="
              toDoc('#')
            "
          >
            <QuestionFilled />
          </el-icon>
        </div>

        <el-table
          :data="form.menuBtn"
          style="width: 100%; margin-top: 12px"
        >
          <el-table-column
            align="left"
            prop="name"
            :label="$t('superAdmin.menuTable.buttonName')"
            width="180"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="name"
            :label="$t('superAdmin.menuTable.remark')"
            width="180"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.desc" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  icon="delete"
                  @click="deleteBtn(form.menuBtn, scope.$index)"
                >
                  {{ $t('common.delete') }}
                </el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  updateBaseMenu,
  getMenuList,
  addBaseMenu,
  deleteBaseMenu,
  getBaseMenuById
} from '@/api/menu'
import icon from '@/view/superAdmin/menu/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { QuestionFilled } from '@element-plus/icons-vue'
import { toDoc } from '@/utils/doc'
import { toLowerCase } from '@/utils/stringFun'
import ComponentsCascader from '@/view/superAdmin/menu/components/components-cascader.vue'

import pathInfo from '@/pathInfo.json'
import { useAppStore } from '@/pinia'
import { useI18n } from 'vue-i18n'

defineOptions({
  name: 'Menus'
})

const appStore = useAppStore()
useI18n()

const rules = reactive({
  path: [{ required: true, message: '请输入菜单name', trigger: 'blur' }],
  component: [{ required: true, message: '请输入文件路径', trigger: 'blur' }],
  'meta.title': [
    { required: true, message: '请输入菜单展示名称', trigger: 'blur' }
  ]
})

const tableData = ref([])
// 查询
const getTableData = async() => {
  tableData.value = []
  const table = await getMenuList()
  if (table.code === 0) {
    tableData.value = table.data
  }
}

getTableData()

// 新增参数
const addParameter = (form) => {
  if (!form.parameters) {
    form.parameters = []
  }
  form.parameters.push({
    type: 'query',
    key: '',
    value: ''
  })
}

const fmtComponent = (component) => {
  form.value.component = component.replace(/\\/g, '/')
  form.value.name = toLowerCase(pathInfo['/src/' + component])
  form.value.path = form.value.name
}

// 删除参数
const deleteParameter = (parameters, index) => {
  parameters.splice(index, 1)
}

// 新增可控按钮
const addBtn = (form) => {
  if (!form.menuBtn) {
    form.menuBtn = []
  }
  form.menuBtn.push({
    name: '',
    desc: ''
  })
}
// 删除可控按钮
const deleteBtn = async(btns, index) => {
  const btn = btns[index]
  if (btn.ID === 0) {
    btns.splice(index, 1)
    return
  }
  const res = await canRemoveAuthorityBtnApi({ id: btn.ID })
  if (res.code === 0) {
    btns.splice(index, 1)
  }
}

const form = ref({
  ID: 0,
  path: '',
  name: '',
  hidden: false,
  parentId: 0,
  component: '',
  meta: {
    activeName: '',
    title: '',
    icon: '',
    defaultMenu: false,
    closeTab: false,
    keepAlive: false
  },
  parameters: [],
  menuBtn: []
})
const changeName = () => {
  form.value.path = form.value.name
}

const handleClose = (done) => {
  initForm()
  done()
}
// 删除菜单
const deleteMenu = (ID) => {
  ElMessageBox.confirm(
    '此操作将永久删除所有角色下该菜单, 是否继续?',
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  )
    .then(async() => {
      const res = await deleteBaseMenu({ ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })

        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}
// 初始化弹窗内表格方法
const menuForm = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  menuForm.value.resetFields()
  form.value = {
    ID: 0,
    path: '',
    name: '',
    hidden: false,
    parentId: 0,
    component: '',
    meta: {
      title: '',
      icon: '',
      defaultMenu: false,
      closeTab: false,
      keepAlive: false
    }
  }
}
// 关闭弹窗

const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}
// 添加menu
const enterDialog = async() => {
  menuForm.value.validate(async(valid) => {
    if (valid) {
      let res
      if (isEdit.value) {
        res = await updateBaseMenu(form.value)
      } else {
        res = await addBaseMenu(form.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? '编辑成功' : '添加成功!'
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}

const menuOption = ref([
  {
    ID: '0',
    title: '根菜单'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: 0,
      title: '根目录'
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
      menuData.forEach((item) => {
        if (item.children && item.children.length) {
          const option = {
            title: item.meta.title,
            ID: item.ID,
            disabled: disabled || item.ID === form.value.ID,
            children: []
          }
          setMenuOptions(
            item.children,
            option.children,
            disabled || item.ID === form.value.ID
          )
          optionsData.push(option)
        } else {
          const option = {
            title: item.meta.title,
            ID: item.ID,
            disabled: disabled || item.ID === form.value.ID
          }
          optionsData.push(option)
        }
      })
}

// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref('新增菜单')
const addMenu = (id) => {
  dialogTitle.value = '新增菜单'
  form.value.parentId = id
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 修改菜单方法
const editMenu = async(id) => {
  dialogTitle.value = '编辑菜单'
  const res = await getBaseMenuById({ id })
  form.value = res.data.menu
  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}
</script>

<style scoped lang="scss">
  .warning {
    color: #dc143c;
  }
  .icon-column {
    display: flex;
    align-items: center;
    .el-icon {
      margin-right: 8px;
    }
  }
</style>
