<template>
  <div>
    <div class="sticky top-0.5 z-10">
      <el-input
        v-model="filterText"
        class="w-3/5"
        :placeholder="$t('common.filter')"
      />
      <el-button
        class="float-right"
        type="primary"
        @click="relation"
      >
        {{ $t('common.confirm') }}
      </el-button>
    </div>
    <div class="tree-content clear-both">
      <el-scrollbar>
        <el-tree
          ref="menuTree"
          :data="menuTreeData"
          :default-checked-keys="menuTreeIds"
          :props="menuDefaultProps"
          default-expand-all
          highlight-current
          node-key="ID"
          show-checkbox
          :filter-node-method="filterNode"
          @check="nodeChange"
        >
          <template #default="{ node, data }">
            <!-- {{data}} -->
            <span class="custom-tree-node">
              <span>{{ t('menu.'+node.label) }}</span>
              <span>
                <el-button
                  v-if="userStore.userInfo.type === 0"
                  type="primary"
                  link
                  :style="{
                    color:
                      row.defaultRouter === data.name ? '#E6A23C' : '#85ce61'
                  }"
                  @click.stop="() => setDefault(data)"
                >
                  {{ row.defaultRouter === data.name ? $t('superAdmin.authority.components.menus.homepage') : $t('superAdmin.authority.components.menus.setAsHomepage') }}
                </el-button>
              </span>
              <span v-if="data.menuBtn && data.menuBtn.length">
                <el-button
                  type="primary"
                  link
                  @click.stop="() => OpenBtn(data)"
                >
                  {{ $t('superAdmin.authority.components.menus.assignButtons') }}
                </el-button>
              </span>
              <span v-if="data.tableColumns && data.tableColumns.length">
                <el-button
                  type="primary"
                  link
                  @click.stop="() => OpenColumns(data)"
                >
                  {{ $t('superAdmin.authority.components.menus.assignColumns') }}
                </el-button>
              </span>
            </span>
          </template>
        </el-tree>
      </el-scrollbar>
    </div>
    <el-dialog
      v-model="btnVisible"
      :title="$t('superAdmin.authority.components.menus.assignButtons')"
      destroy-on-close
    >
      <el-table
        ref="btnTableRef"
        :data="btnData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          :label="$t('superAdmin.authority.components.menus.buttonName')"
          prop="name"
        />
        <el-table-column
          :label="$t('superAdmin.authority.components.menus.buttonRemark')"
          prop="desc"
        >
          <template #default="{ row }">
            {{ $t('common.'+row.name) }}
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <div class="dialog-footer">
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
      </template>
    </el-dialog>

    <el-dialog
      v-model="colVisible"
      :title="$t('superAdmin.authority.components.menus.assignColumns')"
      destroy-on-close
    >
      <el-table
        ref="colTableRef"
        :data="colData"
        row-key="ID"
        @selection-change="handleColSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          :label="$t('superAdmin.authority.components.menus.columnName')"
          prop="jsonName"
        />
        <el-table-column
          :label="$t('superAdmin.authority.components.menus.columnDesc')"
          prop="columnName"
        >
          <template #default="{ row }">
            {{ $t('business.'+row.structName+'.columns.'+row.jsonName) }}
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeColDialog">
            {{ $t('common.cancel') }}
          </el-button>
          <el-button
            type="primary"
            @click="enterColDialog"
          >
            {{ $t('common.confirm') }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getBaseMenuTree,
  getMenuAuthority,
  addMenuAuthority
} from '@/api/menu'
import { updateAuthority } from '@/api/authority'
import { getAuthorityBtnApi, setAuthorityBtnApi } from '@/api/authorityBtn'
import { getAuthorityColApi, setAuthorityColApi, getMenuColumnsApi } from '@/api/authorityCol'
import { getStructNameColumns } from '@/api/systableColumns'
import { nextTick, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

import { useUserStore } from '@/pinia'

const userStore = useUserStore()

defineOptions({
  name: 'Menus'
})

const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const emit = defineEmits(['changeRow'])
const filterText = ref('')
const menuTreeData = ref([])
const menuTreeIds = ref([])
const needConfirm = ref(false)
const menuDefaultProps = ref({
  children: 'children',
  label: function(data) {
    return data.meta.title
  },
  disabled: function(data) {
    // return props.row.defaultRouter === data.name
  }
})

const init = async() => {
  // 获取所有菜单树
  const res = await getBaseMenuTree()
  menuTreeData.value = res.data.menus
  const res1 = await getMenuAuthority({ authorityId: props.row.authorityId })
  const menus = res1.data.menus
  const arr = []
  menus.forEach((item) => {
    // 防止直接选中父级造成全选
    if (!menus.some((same) => same.parentId === item.menuId)) {
      arr.push(Number(item.menuId))
    }
  })
  menuTreeIds.value = arr
}

init()

const setDefault = async(data) => {
  const res = await updateAuthority({
    authorityId: props.row.authorityId,
    AuthorityName: props.row.authorityName,
    parentId: props.row.parentId,
    defaultRouter: data.name
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: t('superAdmin.authority.components.menus.settingSuccess') })
    emit('changeRow', 'defaultRouter', res.data.authority.defaultRouter)
  }
}
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  relation()
}
// 关联树 确认方法
const menuTree = ref(null)
const relation = async() => {
  const checkArr = menuTree.value.getCheckedNodes(false, true)
  const res = await addMenuAuthority({
    menus: checkArr,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('superAdmin.authority.components.menus.menuSettingSuccess')
    })
  }
}

defineExpose({ enterAndNext, needConfirm })

const btnVisible = ref(false)

const btnData = ref([])
const multipleSelection = ref([])
const btnTableRef = ref()
let menuID = ''
const OpenBtn = async(data) => {
  menuID = data.ID
  const res = await getAuthorityBtnApi({
    menuID: menuID,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    openDialog(data)
    await nextTick()
    if (res.data.selected) {
      res.data.selected.forEach((id) => {
        btnData.value.some((item) => {
          if (item.ID === id) {
            btnTableRef.value.toggleRowSelection(item, true)
          }
        })
      })
    }
  }
}

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const openDialog = (data) => {
  btnVisible.value = true
  btnData.value = data.menuBtn
}

const closeDialog = () => {
  btnVisible.value = false
}
const enterDialog = async() => {
  const selected = multipleSelection.value.map((item) => item.ID)
  const res = await setAuthorityBtnApi({
    menuID,
    selected,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: t('superAdmin.authority.components.menus.settingSuccess') })
    btnVisible.value = false
  }
}

const filterNode = (value, data) => {
  if (!value) return true
  // console.log(data.mate.title)
  return data.meta.title.indexOf(value) !== -1
}

watch(filterText, (val) => {
  menuTree.value.filter(val)
})

const colVisible = ref(false)

const colData = ref([])
const multipleColSelection = ref([])
const colTableRef = ref()
let menuIDCol = ''
const OpenColumns = async(data) => {
  menuIDCol = data.ID
  const res = await getAuthorityColApi({
    menuId: menuIDCol,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    openColDialog(data)
    await nextTick()
    if (res.data.selected) {
      res.data.selected.forEach((id) => {
        colData.value.some((item) => {
          if (item.ID === id) {
            colTableRef.value.toggleRowSelection(item, true)
          }
        })
      })
    }
  }
}

const handleColSelectionChange = (val) => {
  multipleColSelection.value = val
}

const openColDialog = async(data) => {
  colVisible.value = true

  try {
    // 首先检查菜单中是否已有TableColumns数据
    if (data.tableColumns && data.tableColumns.length > 0) {
      colData.value = data.tableColumns

      // 如果存在columns映射，需要标记已选中的列
      if (data.columns && Object.keys(data.columns).length > 0) {
        await nextTick()
        colData.value.forEach((item) => {
          if (data.columns[item.jsonName]) {
            colTableRef.value.toggleRowSelection(item, true)
          }
        })
      }
      return
    }

    // 如果菜单中没有TableColumns数据，尝试通过API获取
    if (data.name) {
      // 使用新的API直接获取菜单对应的列
      const res = await getMenuColumnsApi(data.name)
      if (res.code === 0 && res.data && res.data.length > 0) {
        colData.value = res.data
        return
      }

      // 如果通过菜单名直接获取失败，尝试使用首字母大写的结构体名称
      const structName = data.name.charAt(0).toUpperCase() + data.name.slice(1)
      const res2 = await getStructNameColumns(structName)
      if (res2.code === 0 && res2.data && res2.data.length > 0) {
        colData.value = res2.data
        return
      }
    }
  } catch (error) {
    console.error('获取表格列失败:', error)
  }

  // 如果无法获取具体列数据，则设置为空数组
  colData.value = []
}

const closeColDialog = () => {
  colVisible.value = false
}
const enterColDialog = async() => {
  const selected = multipleColSelection.value.map((item) => item.ID)
  const res = await setAuthorityColApi({
    menuId: menuIDCol,
    selected,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: t('superAdmin.authority.components.menus.settingSuccess') })
    colVisible.value = false
  }
}
</script>

<style lang="scss" scoped>
  .custom-tree-node {
    span + span {
      @apply ml-3;
    }
  }
</style>
