<template>
  <div>
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-search-box">
      <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item :label="$t('superAdmin.user.username')">
          <el-input
            v-model="searchInfo.username"
            :placeholder="$t('superAdmin.user.username')"
          />
        </el-form-item>
        <el-form-item :label="$t('superAdmin.user.nickName')">
          <el-input
            v-model="searchInfo.nickname"
            :placeholder="$t('superAdmin.user.nickName')"
          />
        </el-form-item>
        <el-form-item :label="$t('superAdmin.user.phone')">
          <el-input
            v-model="searchInfo.phone"
            :placeholder="$t('superAdmin.user.phone')"
          />
        </el-form-item>
        <el-form-item :label="$t('superAdmin.user.email')">
          <el-input
            v-model="searchInfo.email"
            :placeholder="$t('superAdmin.user.email')"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >
            {{ $t('common.search') }}
          </el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >
            {{ $t('common.reset') }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addUser"
        >
          {{ $t('superAdmin.user.addUser') }}
        </el-button>
      </div>
      <el-table
        :data="tableData"
        row-key="ID"
        border
      >
        <el-table-column
          align="left"
          label="ID"
          min-width="50"
          prop="ID"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.user.headerImg')"
          min-width="75"
        >
          <template #default="scope">
            <CustomPic
              style="margin-top: 8px"
              :pic-src="scope.row.headerImg"
            />
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          :label="$t('superAdmin.user.username')"
          min-width="150"
          prop="userName"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.user.nickName')"
          min-width="150"
          prop="nickName"
        />

        <el-table-column
          align="left"
          :label="$t('superAdmin.user.type')"
          min-width="75"
          prop="type"
        >
          <template #default="scope">
            <el-tag
              v-if="scope.row.type === 0"
              type="success"
              effect="dark"
            >
              Admin
            </el-tag>
            <el-tag
              v-if="scope.row.type === 2"
              type="warning"
              effect="dark"
            >
              Merchant
            </el-tag>
            <el-tag
              v-if="scope.row.type === 1"
              type="danger"
              effect="dark"
            >
              Agent
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="$t('superAdmin.user.parameter')"
          prop="parameter"
          min-width="75"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.user.phone')"
          min-width="180"
          prop="phone"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.user.email')"
          min-width="180"
          prop="email"
        />
        <el-table-column
          align="left"
          :label="$t('superAdmin.user.authorityName')"
          min-width="200"
        >
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{
                multiple: true,
                checkStrictly: true,
                label: 'authorityName',
                value: 'authorityId',
                disabled: 'disabled',
                emitPath: false
              }"
              :clearable="false"
              @visible-change="
                (flag) => {
                  changeAuthority(scope.row, flag, 0)
                }
              "
              @remove-tag="
                (removeAuth) => {
                  changeAuthority(scope.row, false, removeAuth)
                }
              "
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="$t('superAdmin.user.enabled')"
          min-width="150"
        >
          <template #default="scope">
            <el-switch
              v-model="scope.row.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="
                () => {
                  switchEnable(scope.row)
                }
              "
            />
          </template>
        </el-table-column>

        <el-table-column
          :label="$t('common.operations')"
          :min-width="appStore.operateMinWith"
          fixed="right"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteUserFunc(scope.row)"
            >
              {{ $t('common.delete') }}
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="openEdit(scope.row)"
            >
              {{ $t('common.edit') }}
            </el-button>
            <el-button
              type="primary"
              link
              icon="magic-stick"
              @click="resetPasswordFunc(scope.row)"
            >
              {{ $t('superAdmin.user.resetPassword') }}
            </el-button>
            <el-button
              type="warning"
              link
              icon="unlock"
              @click="unbindSecurityFunc(scope.row)"
            >
              {{ $t('superAdmin.user.unbind') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
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
      v-model="addUserDialog"
      :size="appStore.drawerSize"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">
            {{ $t('superAdmin.user.title') }}
          </span>
          <div>
            <el-button @click="closeAddUserDialog">
              {{ $t('common.cancel') }}
            </el-button>
            <el-button
              type="primary"
              @click="enterAddUserDialog"
            >
              {{ $t('common.confirm') }}
            </el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="userForm"
        :rules="rules"
        :model="userInfo"
        label-width="80px"
      >
        <el-form-item
          v-if="dialogFlag === 'add'"
          :label="$t('superAdmin.user.username')"
          prop="userName"
        >
          <el-input v-model="userInfo.userName" />
        </el-form-item>
        <el-form-item
          v-if="dialogFlag === 'add'"
          :label="$t('superAdmin.user.password')"
          prop="password"
        >
          <el-input v-model="userInfo.password">
            <template #append>
              <el-button
                icon="refresh"
                type="primary"
                size="default"
                @click="userInfo.password =generateRandomString(15)"
              />
            </template>
          </el-input>
        </el-form-item>

        <el-form-item
          :label="$t('superAdmin.user.nickName')"
          prop="nickName"
        >
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.user.phone')"
          prop="phone"
        >
          <el-input v-model="userInfo.phone" />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.user.email')"
          prop="email"
        >
          <el-input v-model="userInfo.email" />
        </el-form-item>
        <el-form-item
          v-if="userStore.userInfo.type === 0"
          :label="$t('superAdmin.user.type')"
          prop="type"
        >
          <el-select v-model="userInfo.type">
            <el-option
              label="Admin"
              :value="0"
            />
            <el-option
              label="Agent"
              :value="1"
            />
            <el-option
              label="Merchant"
              :value="2"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          v-if="userStore.userInfo.type === 0"
          :label="$t('superAdmin.user.parameter')"
          prop="parameter"
        >
          <el-input v-model="userInfo.parameter" />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.user.authorityName')"
          prop="authorityId"
        >
          <el-cascader
            v-model="userInfo.authorityIds"
            style="width: 100%"
            :options="authOptions"
            :show-all-levels="false"
            :props="{
              multiple: true,
              checkStrictly: true,
              label: 'authorityName',
              value: 'authorityId',
              disabled: 'disabled',
              emitPath: false
            }"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.user.enabled')"
          prop="disabled"
        >
          <el-switch
            v-model="userInfo.enable"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
        <el-form-item
          :label="$t('superAdmin.user.headerImg')"
          label-width="80px"
        >
          <SelectImage v-model="userInfo.headerImg" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword, unbindSecurity } from '@/api/user.js'
import { useI18n } from 'vue-i18n'
import { generateRandomString } from '@/utils/rand'
import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import SelectImage from '@/components/selectImage/selectImage.vue'
import { useAppStore, useUserStore } from '@/pinia'

const { t } = useI18n()

const appStore = useAppStore()

const userStore = useUserStore()
const searchInfo = ref({
  username: '',
  nickname: '',
  phone: '',
  email: ''
})

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    username: '',
    nickname: '',
    phone: '',
    email: ''
  }
  getTableData()
}
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
      AuthorityData.forEach((item) => {
        if (item.children && item.children.length) {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            children: []
          }
          setAuthorityOptions(item.children, option.children)
          optionsData.push(option)
        } else {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName
          }
          optionsData.push(option)
        }
      })
}

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
  tableData.value = []
  const table = await getUserList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(
  () => tableData.value,
  () => {
    setAuthorityIds()
  }
)

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList()
  setOptions(res.data)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm('是否将此用户密码重置为123456?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await resetPassword({
      ID: row.ID
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg
      })
    }
  })
}

const unbindSecurityFunc = (row) => {
  ElMessageBox.confirm('是否解绑此用户的安全验证（谷歌验证器和通行密钥）?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await unbindSecurity({
      ID: row.ID
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '解绑成功'
      })
      await getTableData() // 刷新表格数据
    } else {
      ElMessage({
        type: 'error',
        message: res.msg || '解绑失败'
      })
    }
  })
}

const setAuthorityIds = () => {
  tableData.value &&
      tableData.value.forEach((user) => {
        user.authorityIds =
          user.authorities &&
          user.authorities.map((i) => {
            return i.authorityId
          })
      })
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const deleteUserFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteUser({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      await getTableData()
    }
  })
}

// 弹窗相关
const userInfo = ref({
  userName: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1
})

const rules = ref({
  userName: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 4, message: '最低4位字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  nickName: [{ required: true, message: '请输入用户昵称', trigger: 'blur' }],
  phone: [
    {
      pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/,
      message: '请输入合法手机号',
      trigger: 'blur'
    }
  ],
  email: [
    {
      pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g,
      message: '请输入正确的邮箱',
      trigger: 'blur'
    }
  ],
  authorityId: [
    { required: true, message: '请选择用户角色', trigger: 'blur' }
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async() => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async(valid) => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const tempAuth = {}
const changeAuthority = async(row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authorityIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '角色设置成功' })
  } else {
    if (!removeAuth) {
      row.authorityIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authorityIds = [removeAuth, ...row.authorityIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async(row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: `${req.enable === 2 ? '禁用' : '启用'}成功`
    })
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
  }
}
</script>

<style lang="scss">
  .header-img-box {
    @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
  }
</style>
