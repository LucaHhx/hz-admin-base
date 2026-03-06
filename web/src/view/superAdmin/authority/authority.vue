<template>
  <div class="authority">
    <warning-bar :title="$t('superAdmin.authority.warningMessage')" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addAuthority(0)"
        >
          {{ $t('superAdmin.authority.addRole') }}
        </el-button>
      </div>
      <el-table
        :data="tableData"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        row-key="authorityId"
        style="width: 100%"
      >
        <el-table-column
          :label="$t('superAdmin.authority.roleId')"
          min-width="180"
          prop="authorityId"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.authority.roleName')"
          min-width="180"
          prop="authorityName"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.authority.operations')"
          width="460"
        >
          <template #default="scope">
            <el-button
              icon="setting"
              type="primary"
              link
              @click="openDrawer(scope.row)"
            >
              {{ $t('superAdmin.authority.setPermissions') }}
            </el-button>
            <!-- <el-button
              icon="plus"
              type="primary"
              link
              @click="addAuthority(scope.row.authorityId)"
            >{{ $t('superAdmin.authority.addSubRole') }}</el-button> -->
            <el-button
              icon="copy-document"
              type="primary"
              link
              @click="copyAuthorityFunc(scope.row)"
            >
              {{ $t('superAdmin.authority.copy') }}
            </el-button>
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editAuthority(scope.row)"
            >
              {{ $t('common.edit') }}
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteAuth(scope.row)"
            >
              {{ $t('common.delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增角色弹窗 -->
    <el-drawer
      v-model="authorityFormVisible"
      :size="appStore.drawerSize"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ authorityTitleForm }}</span>
          <div>
            <el-button @click="closeAuthorityForm">
              {{ $t('common.cancel') }}
            </el-button>
            <el-button
              type="primary"
              @click="submitAuthorityForm"
            >
              {{ $t('common.confirm') }}
            </el-button>
          </div>
        </div>
      </template>
      <el-form
        ref="authorityForm"
        :model="form"
        :rules="rules"
        label-width="auto"
      >
        <!-- <el-form-item :label="$t('superAdmin.authority.parentRole')" prop="parentId">
          <el-cascader
            v-model="form.parentId"
            style="width: 100%"
            :disabled="dialogType === 'add'"
            :options="AuthorityOption"
            :props="{
              checkStrictly: true,
              label: 'authorityName',
              value: 'authorityId',
              disabled: 'disabled',
              emitPath: false
            }"
            :show-all-levels="false"
            filterable
          />
        </el-form-item> -->
        <!-- <el-form-item :label="$t('superAdmin.authority.roleId')" prop="authorityId">
          <el-input
            v-model="form.authorityId"
            :disabled="dialogType === 'edit'"
            autocomplete="off"
            maxlength="15"
          />
        </el-form-item> -->
        <el-form-item
          :label="$t('superAdmin.authority.roleName')"
          prop="authorityName"
        >
          <el-input
            v-model="form.authorityName"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      v-if="drawer"
      v-model="drawer"
      :size="appStore.drawerSize"
      :title="$t('superAdmin.authority.roleConfig')"
    >
      <el-tabs
        :before-leave="autoEnter"
        type="border-card"
      >
        <el-tab-pane
          v-if="userInfo.type ===0"
          :label="$t('superAdmin.authority.roleMenu')"
        >
          <Menus
            ref="menus"
            :row="activeRow"
            @change-row="changeRow"
          />
        </el-tab-pane>
        <el-tab-pane
          v-if="userInfo.type ===0"
          :label="$t('superAdmin.authority.roleApi')"
        >
          <Apis
            ref="apis"
            :row="activeRow"
            @change-row="changeRow"
          />
        </el-tab-pane>
        <el-tab-pane
          v-if="userInfo.type !==0"
          :label="$t('superAdmin.authority.roleMenu')"
        >
          <MerMenus
            ref="merMenus"
            :row="activeRow"
            @change-row="changeRow"
          />
        </el-tab-pane>
        <!-- <el-tab-pane :label="$t('superAdmin.authority.resourcePermissions')">
          <Datas
            ref="datas"
            :authority="tableData"
            :row="activeRow"
            @changeRow="changeRow"
          />
        </el-tab-pane> -->
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getAuthorityList,
  deleteAuthority,
  createAuthority,
  updateAuthority,
  copyAuthority
} from '@/api/authority'

import Menus from '@/view/superAdmin/authority/components/menus.vue'
import Apis from '@/view/superAdmin/authority/components/apis.vue'
import MerMenus from '@/view/superAdmin/authority/components/mer_menus.vue'
// import Datas from '@/view/superAdmin/authority/components/datas.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { useI18n } from 'vue-i18n'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore, useUserStore } from '@/pinia'

const { t } = useI18n()

const userInfo = useUserStore().userInfo
// console.log(userInfo)

const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error(t('superAdmin.authority.positiveIntegerRequired')))
  }
  return callback()
}

const AuthorityOption = ref([
  {
    authorityId: 0,
    authorityName: t('superAdmin.authority.rootRole')
  }
])
const drawer = ref(false)
const dialogType = ref('add')
const activeRow = ref({})
const appStore = useAppStore()

const authorityTitleForm = ref(t('superAdmin.authority.addRole'))
const authorityFormVisible = ref(false)
const apiDialogFlag = ref(false)
const copyForm = ref({})

const form = ref({
  authorityId: 0,
  authorityName: '',
  parentId: 0
})
const rules = ref({
  // authorityId: [
  //   { required: true, message: t('superAdmin.authority.roleIdRequired'), trigger: 'blur' },
  //   { validator: mustUint, trigger: 'blur', message: t('superAdmin.authority.positiveIntegerRequired') }
  // ],
  authorityName: [
    { required: true, message: t('superAdmin.authority.roleNameRequired'), trigger: 'blur' }
  ],
  parentId: [{ required: true, message: t('superAdmin.authority.parentRoleRequired'), trigger: 'blur' }]
})

const tableData = ref([])

// 查询
const getTableData = async() => {
  tableData.value = []
  const table = await getAuthorityList()
  if (table.code === 0) {
    tableData.value = table.data
  }
}

getTableData()

const changeRow = (key, value) => {
  activeRow.value[key] = value
}
const menus = ref(null)
const apis = ref(null)
const datas = ref(null)
const autoEnter = (activeName, oldActiveName) => {
  const paneArr = [menus, apis, datas]
  if (oldActiveName) {
    if (paneArr[oldActiveName].value && paneArr[oldActiveName].value.needConfirm) {
      paneArr[oldActiveName].value.enterAndNext()
      paneArr[oldActiveName].value.needConfirm = false
    }
  }
}
// 拷贝角色
const copyAuthorityFunc = (row) => {
  setOptions()
  authorityTitleForm.value = t('superAdmin.authority.copyRole')
  dialogType.value = 'copy'
  for (const k in form.value) {
    form.value[k] = row[k]
  }
  copyForm.value = row
  authorityFormVisible.value = true
}
const openDrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}
// 删除角色
const deleteAuth = (row) => {
  ElMessageBox.confirm(t('superAdmin.authority.deleteConfirm'), t('common.warning'), {
    confirmButtonText: t('common.confirm'),
    cancelButtonText: t('common.cancel'),
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteAuthority({ authorityId: row.authorityId })
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
// 初始化表单
const authorityForm = ref(null)
const initForm = () => {
  if (authorityForm.value) {
    authorityForm.value.resetFields()
  }
  form.value = {
    authorityId: 0,
    authorityName: '',
    parentId: 0
  }
}
// 关闭窗口
const closeAuthorityForm = () => {
  initForm()
  authorityFormVisible.value = false
  apiDialogFlag.value = false
}
// 确定弹窗

const submitAuthorityForm = () => {
  authorityForm.value.validate(async(valid) => {
    if (valid) {
      form.value.authorityId = Number(form.value.authorityId)
      switch (dialogType.value) {
        case 'add':
          {
            const res = await createAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功!'
              })
              getTableData()
              closeAuthorityForm()
            }
          }
          break
        case 'edit':
          {
            const res = await updateAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功!'
              })
              getTableData()
              closeAuthorityForm()
            }
          }
          break
        case 'copy': {
          const data = {
            authority: {
              authorityId: 0,
              authorityName: '',
              datauthorityId: [],
              parentId: 0
            },
            oldAuthorityId: 0
          }
          data.authority.authorityId = form.value.authorityId
          data.authority.authorityName = form.value.authorityName
          data.authority.parentId = form.value.parentId
          data.authority.dataAuthorityId = copyForm.value.dataAuthorityId
          data.oldAuthorityId = copyForm.value.authorityId
          const res = await copyAuthority(data)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '复制成功！'
            })
            getTableData()
          }
        }
      }

      initForm()
      authorityFormVisible.value = false
    }
  })
}
const setOptions = () => {
  AuthorityOption.value = [
    {
      authorityId: 0,
      authorityName: t('superAdmin.authority.rootRole')
    }
  ]
  setAuthorityOptions(tableData.value, AuthorityOption.value, false)
}
const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
  AuthorityData &&
      AuthorityData.forEach((item) => {
        if (item.children && item.children.length) {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            disabled: disabled || item.authorityId === form.value.authorityId,
            children: []
          }
          setAuthorityOptions(
            item.children,
            option.children,
            disabled || item.authorityId === form.value.authorityId
          )
          optionsData.push(option)
        } else {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            disabled: disabled || item.authorityId === form.value.authorityId
          }
          optionsData.push(option)
        }
      })
}
// 增加角色
const addAuthority = (parentId) => {
  initForm()
  authorityTitleForm.value = t('superAdmin.authority.addRole')
  dialogType.value = 'add'
  form.value.parentId = parentId
  setOptions()
  authorityFormVisible.value = true
}
// 编辑角色
const editAuthority = (row) => {
  setOptions()
  authorityTitleForm.value = t('superAdmin.authority.editRole')
  dialogType.value = 'edit'
  for (const key in form.value) {
    form.value[key] = row[key]
  }
  setOptions()
  authorityForm.value && authorityForm.value.clearValidate()
  authorityFormVisible.value = true
}
</script>

<style lang="scss">
  .authority {
    .el-input-number {
      margin-left: 15px;
      span {
        display: none;
      }
    }
  }
  .tree-content {
    margin-top: 10px;
    height: calc(100vh - 158px);
    overflow: auto;
  }
</style>
