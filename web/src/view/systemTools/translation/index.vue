<template>
  <div class="translation-container">
    <el-container>
      <el-aside width="300px">
        <div class="aside-header">
          <span>语言包管理</span>
          <el-button type="primary" link @click="newLanguageVisible = true">
            <el-icon><Plus /></el-icon>新增
          </el-button>
        </div>
        <el-tree
          :data="treeData"
          :highlight-current="true"
          node-key="path"
          default-expand-all
          @node-click="handleNodeClick"
        >
          <template #default="{ node, data }">
            <div class="tree-node">
              <span>{{ data.title }}</span>
              <div class="node-actions">
                <el-button
                  v-if="!data.isLeaf && data.path !== 'zh-CN' && data.path !== 'en-US' && !data.path.includes('/')"
                  type="danger"
                  link
                  @click.stop="handleDeleteLanguage(data.title)"
                >
                  <el-icon><Delete /></el-icon>
                </el-button>
                <el-button
                  v-if="!data.isLeaf && !data.path.includes('/')"
                  type="primary"
                  link
                  @click.stop="handleDownload(data.title)"
                >
                  <el-icon><Download /></el-icon>
                </el-button>
                <el-button
                  v-if="!data.isLeaf && !data.path.includes('/')"
                  type="success"
                  link
                  @click.stop="handleUpload(data.title)"
                >
                  <el-icon><Upload /></el-icon>
                </el-button>
              </div>
            </div>
          </template>
        </el-tree>
      </el-aside>
      <el-main>
        <div v-if="currentFile" class="editor-box">
          <div class="editor-header">
            <span>{{ currentFile }}</span>
            <div class="editor-actions">
              <el-button type="info" size="small" @click="handleUndo">撤回</el-button>
              <el-button type="primary" size="small" @click="handleSave">保存</el-button>
            </div>
          </div>
          <el-input
            v-model="jsonContent"
            type="textarea"
            :autosize="{ minRows: 25 }"
            :placeholder="'请输入JSON内容'"
            tabindex="0"
            @input="handleJsonInput"
            @keydown.tab.prevent="handleTab"
            @keydown.shift.tab.prevent="handleShiftTab"
          />
        </div>
        <el-empty v-else description="请选择要编辑的文件" />
      </el-main>
    </el-container>

    <!-- 新增语言对话框 -->
    <el-dialog
      v-model="newLanguageVisible"
      title="新增语言"
      width="400px"
    >
      <el-form
        ref="newLanguageFormRef"
        :model="newLanguageForm"
        :rules="newLanguageRules"
        label-width="80px"
      >
        <el-form-item label="源语言" prop="fromLang">
          <el-select v-model="newLanguageForm.fromLang" placeholder="请选择源语言">
            <el-option
              v-for="lang in languages"
              :key="lang"
              :label="lang"
              :value="lang"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="目标语言" prop="toLang">
          <el-input v-model="newLanguageForm.toLang" placeholder="请输入目标语言代码" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="newLanguageVisible = false">取消</el-button>
          <el-button type="primary" @click="handleAddLanguage">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import {
  getFileTree,
  getFileContent,
  updateTranslation,
  copyLanguage,
  deleteLanguage,
  downloadLanguage,
  uploadLanguage
} from '@/api/system/translation'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Delete, Download, Upload } from '@element-plus/icons-vue'
import JSZip from 'jszip'

// 文件树数据
const treeData = ref([])
const currentFile = ref('')
const jsonContent = ref('')
const languages = ref([])

// 新增语言
const newLanguageVisible = ref(false)
const newLanguageForm = ref({
  fromLang: '',
  toLang: ''
})
const newLanguageRules = {
  fromLang: [{ required: true, message: '请选择源语言', trigger: 'change' }],
  toLang: [{ required: true, message: '请输入目标语言代码', trigger: 'blur' }]
}
const newLanguageFormRef = ref(null)

// 加载文件树
const loadFileTree = async() => {
  try {
    const treeRes = await getFileTree()
    treeData.value = treeRes.data
    // 获取语言列表
    languages.value = treeData.value.map(item => item.title)
  } catch (error) {
    console.error('加载文件树失败:', error)
    ElMessage.error('加载文件树失败')
  }
}

// 处理节点点击
const handleNodeClick = async(data) => {
  if (data.isLeaf) {
    try {
      currentFile.value = data.path
      const [lang, ...filePath] = data.path.split('/')
      const file = filePath.join('/')

      // 获取文件内容
      const res = await getFileContent(lang, file)
      if (res.data) {
        jsonContent.value = res.data
      } else {
        throw new Error('找不到文件内容')
      }
    } catch (error) {
      console.error('内容解析失败:', error)
      ElMessage.error(error.message || '文件内容格式错误')
      currentFile.value = ''
      jsonContent.value = ''
    }
  }
}

// 撤回编辑
const handleUndo = async() => {
  try {
    if (!currentFile.value) return
    const [lang, ...filePath] = currentFile.value.split('/')
    const file = filePath.join('/')

    // 重新获取文件内容
    const res = await getFileContent(lang, file)
    if (res.data) {
      jsonContent.value = res.data
    }
  } catch (error) {
    console.error('撤回编辑失败:', error)
    ElMessage.error('撤回编辑失败')
  }
}

// 保存翻译
const handleSave = async() => {
  try {
    if (!jsonContent.value) {
      throw new Error('没有可保存的内容')
    }

    // 验证 JSON 格式
    try {
      JSON.parse(jsonContent.value)
    } catch {
      throw new Error('JSON 格式不正确，请检查后重试')
    }

    const [lang, ...filePath] = currentFile.value.split('/')
    await updateTranslation({
      language: lang,
      file: filePath.join('/'),
      content: jsonContent.value
    })
    ElMessage.success('保存成功')
    loadFileTree()
  } catch (error) {
    console.error('保存失败:', error)
    ElMessage.error(error.message || '保存失败')
  }
}

// 新增语言
const handleAddLanguage = async() => {
  try {
    await newLanguageFormRef.value.validate()
    await copyLanguage(newLanguageForm.value)
    ElMessage.success('新增语言成功')
    newLanguageVisible.value = false
    loadFileTree()
  } catch (error) {
    if (error) {
      console.error('新增语言失败:', error)
      ElMessage.error('新增语言失败')
    }
  }
}

// 删除语言包
const handleDeleteLanguage = async(lang) => {
  try {
    await ElMessageBox.confirm(
      '确定要删除该语言包吗？此操作将无法恢复！',
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    await deleteLanguage({ language: lang })
    ElMessage.success('删除成功')
    loadFileTree()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 下载语言包
const handleDownload = async(lang) => {
  try {
    // 获取该语言下的所有文件
    const files = treeData.value.find(item => item.title === lang)?.children || []

    // 递归获取所有文件路径
    const getAllFiles = (nodes) => {
      const paths = []
      nodes.forEach(node => {
        if (node.isLeaf) {
          paths.push(node.path)
        } else if (node.children) {
          paths.push(...getAllFiles(node.children))
        }
      })
      return paths
    }

    const filePaths = getAllFiles(files)
    const zip = new JSZip()

    // 逐个获取文件内容并添加到 zip
    for (const path of filePaths) {
      const [language, ...filePath] = path.split('/')
      const file = filePath.join('/')

      try {
        const res = await getFileContent(language, file)
        if (res.code === 0 && res.data) {
          zip.file(file, res.data)
        }
      } catch (error) {
        console.error(`获取文件 ${file} 内容失败:`, error)
        continue
      }
    }

    // 生成 zip 文件
    const content = await zip.generateAsync({ type: 'blob' })

    // 创建下载链接
    const url = window.URL.createObjectURL(content)
    const link = document.createElement('a')
    link.href = url
    link.download = `${lang}.zip`
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)
  } catch (error) {
    console.error('下载失败:', error)
    ElMessage.error('下载失败')
  }
}

// 上传语言包
const handleUpload = async(lang) => {
  try {
    // 创建文件输入元素
    const input = document.createElement('input')
    input.type = 'file'
    input.accept = '.zip'

    input.onchange = async(e) => {
      const file = e.target.files[0]
      if (!file) return

      try {
        await ElMessageBox.confirm(
          '上传将覆盖现有文件，是否继续？',
          '警告',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        )

        const zip = new JSZip()
        const zipContent = await zip.loadAsync(file)

        // 获取压缩包的根目录名称
        const rootDir = Object.keys(zipContent.files)[0].split('/')[0]

        // 遍历 zip 中的文件
        for (const [path, content] of Object.entries(zipContent.files)) {
          if (content.dir) continue // 跳过目录

          // 去掉压缩包名称，获取实际文件路径
          const actualPath = path.replace(`${rootDir}/`, '')

          // 获取文件内容
          const fileContent = await content.async('string')

          // 调用保存接口
          try {
            const res = await updateTranslation({
              language: lang,
              file: actualPath,
              content: fileContent
            })

            if (res.code === 0) {
              console.log(`文件 ${actualPath} 上传成功`)
            } else {
              console.error(`文件 ${actualPath} 上传失败:`, res.msg)
            }
          } catch (error) {
            console.error(`文件 ${actualPath} 上传失败:`, error)
          }
        }

        ElMessage.success('上传完成')
        loadFileTree()
      } catch (error) {
        if (error !== 'cancel') {
          console.error('上传失败:', error)
          ElMessage.error('上传失败')
        }
      }
    }

    input.click()
  } catch (error) {
    console.error('上传失败:', error)
    ElMessage.error('上传失败')
  }
}

// 在 script setup 部分添加以下函数
const handleTab = (e) => {
  const textarea = e.target
  const start = textarea.selectionStart
  const end = textarea.selectionEnd

  jsonContent.value = jsonContent.value.substring(0, start) + '    ' + jsonContent.value.substring(end)

  // 设置光标位置到插入的制表符后面
  setTimeout(() => {
    textarea.selectionStart = textarea.selectionEnd = start + 4
  }, 0)
}

const handleShiftTab = (e) => {
  const textarea = e.target
  const start = textarea.selectionStart
  const end = textarea.selectionEnd

  // 获取当前行的开始位置
  const beforeCursor = jsonContent.value.substring(0, start)
  const lines = beforeCursor.split('\n')
  const currentLineStart = start - lines[lines.length - 1].length

  // 检查当前行是否有4个空格
  const currentLine = jsonContent.value.substring(currentLineStart, start)
  if (currentLine.startsWith('    ')) {
    jsonContent.value = jsonContent.value.substring(0, currentLineStart) + jsonContent.value.substring(currentLineStart + 4)

    // 设置光标位置
    setTimeout(() => {
      textarea.selectionStart = textarea.selectionEnd = start - 4
    }, 0)
  }
}

// 组件挂载和销毁
onMounted(() => {
  loadFileTree()
})

onBeforeUnmount(() => {
  // 清理工作
})
</script>

<style scoped>
.translation-container {
  height: calc(100vh - 120px);
  padding: 20px;
}

.el-container {
  height: 100%;
  background-color: var(--el-bg-color);
  border-radius: 4px;
}

.el-aside {
  padding: 20px;
  border-right: 1px solid var(--el-border-color-light);
}

.aside-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--el-border-color-light);
}

.el-main {
  padding: 20px;
}

.editor-box {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--el-border-color-light);
}

.editor-actions {
  display: flex;
  gap: 10px;
}

.el-tree {
  background: none;
}

:deep(.el-textarea__inner) {
  font-family: monospace;
  font-size: 14px;
  line-height: 1.6;
  tab-size: 2;
}

.tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 8px;
}

.tree-node:hover .node-actions {
  opacity: 1;
}

:deep(.el-tree-node__content) {
  height: 32px;
}

:deep(.el-tree-node__content:hover) {
  background-color: var(--el-tree-node-hover-bg-color);
}
</style>
