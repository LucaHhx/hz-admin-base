<template>
  <div>
    <el-dialog
      v-model="visible"
      title="Sync Table Columns"
      width="50%"
    >
      <h2>新增列表</h2>
      <el-table :data="addList" border stripe>
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
          sortable align="left" :label="$t('business.sysTableColumns.columns.note')" prop="note"
          width="120"
        />
        <el-table-column>
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="plus"
              @click="()=>{
                create(scope.row)}"
            >
              {{ $t('common.add') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <h2>删除列表</h2>
      <el-table :data="delList" border stripe>
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
          sortable align="left" :label="$t('business.sysTableColumns.columns.note')" prop="note"
          width="120"
        />
        <el-table-column>
          <template #default="scope">
            <el-button
              type="danger"
              link
              icon="delete"
              @click="()=>{
                deleteSysTableColumns({ ID: scope.row.ID }).then((res) => {
                  if (res.code === 0) {
                    ElMessage.success('Column deleted successfully')
                    sync()
                  } else {
                    ElMessage.error(res.msg || 'Failed to delete column')
                  }
                }).catch((error) => {
                  console.error(error)
                  ElMessage.error('An error occurred during deletion')
                })
              }"
            >
              {{ $t('common.delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <template #footer>
        <span>
          <el-button @click=" ()=>{visible= false;addList=[];delList=[]}">Close</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  syncSysTableColumnsInfo,
  createSysTableColumns,
  deleteSysTableColumns
} from '@/api/system/sysTableColumns'
import { ref, defineExpose } from 'vue'
import { ElMessage } from 'element-plus'
const visible = ref(false)
const addList = ref([])
const delList = ref([])
const name = ref('')

const sync = async() => {
  await syncSysTableColumnsInfo({ name: name.value }).then((res) => {
    if (res.code === 0) {
      addList.value = res.data.addList || []
      delList.value = res.data.delList || []
    } else {
      ElMessage.error(res.msg || 'Failed to retrieve sync data')
    }
  }).catch((error) => {
    console.error(error)
    ElMessage.error('An error occurred during sync')
  })
}

const open = async(tbname) => {
  name.value = tbname
  await sync()
  visible.value = true
}

defineExpose({
  open,
})

const create = async(row) => {
  await createSysTableColumns(row).then((res) => {
    console.log(res)
    if (res.code === 0) {
      ElMessage.success('Column added successfully')
    } else {
      ElMessage.error(res.msg || 'Failed to add column')
    }
  })
  addList.value = []
  delList.value = []
  await sync()
}
</script>
