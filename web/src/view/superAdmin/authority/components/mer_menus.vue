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
          check-strictly
          :filter-node-method="filterNode"
          @check="nodeChange"
        >
          <template #default="{ node, data }">
            <span class="custom-tree-node">
              <span>{{ node.label }}</span>
            </span>
          </template>
        </el-tree>
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup>
import {
  getBaseMenuTree,
  getMenuAuthority,
  addMenuAuthority,
} from '@/api/menu'
import { updateAuthority } from '@/api/authority'
import { getAuthorityBtnAllApi } from '@/api/authorityBtn'
import { nextTick, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n'
import { useUserStore } from '@/pinia'
import { FormItemInputContext } from 'ant-design-vue/es/form/FormItemContext'

const { t } = useI18n()
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
    return data?.meta?.title ? t('menu.' + data.meta.title) : t('common.' + data.name)
  },
  disabled: function(data) {
    return data.name === 'userInfo'
  }
})

const authButtons = ref({})

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

  AddButtonNodes(menuTreeData.value)
  const res2 = await getAuthorityBtnAllApi({
    authorityId: props.row.authorityId
  })
  res2.data?.selected?.forEach((item) => {
    arr.push(item + 15023)
  })
  arr.push(53)
  menuTreeIds.value = arr
}

const AddButtonNodes = (items) => {
  items.forEach((item) => {
    if (item.menuBtn) {
      if (item.children) {
        AddButtonNodes(item.children)
      }
      if (item.menuBtn.length > 0) {
        item.children = []
        item.menuBtn.forEach((item2) => {
          item.children.push({
            ID: item2.ID + 15023,
            name: item2.name,
            parentId: item.ID,
            children: []
          })
        })
      }
    }
  })
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

const handleNodeClick = (data, node) => {
  menuTree.value.setChecked(node, !node.checked)
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

const filterNode = (value, data) => {
  if (!value) return true
  return data.meta.title.indexOf(value) !== -1
}

watch(filterText, (val) => {
  menuTree.value.filter(val)
})

defineExpose({ enterAndNext, needConfirm })
</script>

<style lang="scss" scoped>
.custom-tree-node {
  span + span {
    @apply ml-3;
  }
}
</style>
