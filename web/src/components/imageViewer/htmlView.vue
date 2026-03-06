<template>
  <el-dialog
    v-model="visible"
    title="HTML预览"
    :style="dialogStyle"
    :close-on-click-modal="false"
    :destroy-on-close="true"
    class="html-dialog"
    modal-class="dialog_modal"
    draggable
    :modal="false"
  >
    <div class="flex flex-col h-full overflow-hidden">
      <!-- 切换按钮 -->
      <div class="mb-2 flex justify-end">
        <el-button
          :type="showSource ? 'primary' : 'default'"
          size="small"
          @click="toggleView"
        >
          {{ showSource ? '视图模式' : '源码模式' }}
        </el-button>
      </div>

      <!-- HTML内容区域 -->
      <div class="flex-1 overflow-hidden border border-gray-200">
        <!-- 源码模式 -->
        <el-input
          v-if="showSource"
          v-model="htmlContent"
          type="textarea"
          readonly
          class="w-full h-full"
          :autosize="{ minRows: 20 }"
          resize="none"
        />
        <!-- 视图模式 -->
        <iframe
          v-else
          ref="htmlFrame"
          :srcdoc="htmlContent"
          class="w-full h-full border-0"
          sandbox="allow-scripts allow-same-origin"
        />
      </div>

      <!-- info 显示区域 -->
      <div v-if="info" class="mt-4 px-4">
        <el-input
          v-model="info"
          type="textarea"
          readonly
          autosize
          placeholder="附加信息"
          :disabled="true"
          resize="none"
        />
      </div>
    </div>
  </el-dialog>
</template>

<script setup name="HtmlViewer">
import { ref, nextTick, watch } from 'vue'

const visible = ref(false)
const htmlContent = ref('')
const info = ref('')
const showSource = ref(true)
const dialogStyle = ref({
  width: '90vw',
  height: '80vh'
})

function open(htmlData, infostr) {
  visible.value = true
  htmlContent.value = htmlData
  info.value = infostr || ''
  showSource.value = true // 默认显示视图模式
}

function toggleView() {
  showSource.value = !showSource.value
}

function makeDialogDraggable() {
  nextTick(() => {
    const header = document.querySelector('.el-dialog__header')
    const dialog = document.querySelector('.el-dialog')
    if (!header || !dialog) return

    header.style.cursor = 'move'
    let offsetX = 0
    let offsetY = 0
    let dragging = false

    header.onmousedown = e => {
      dragging = true
      offsetX = e.clientX - dialog.offsetLeft
      offsetY = e.clientY - dialog.offsetTop

      document.onmousemove = e => {
        if (!dragging) return
        dialog.style.margin = '0'
        dialog.style.left = e.clientX - offsetX + 'px'
        dialog.style.top = e.clientY - offsetY + 'px'
        dialog.style.position = 'absolute'
      }

      document.onmouseup = () => {
        dragging = false
        document.onmousemove = null
        document.onmouseup = null
      }
    }
  })
}

watch(visible, val => {
  if (val) makeDialogDraggable()
})

defineExpose({ open })
</script>

<style scoped>
.html-dialog {
  --el-dialog-margin-top: 5vh;
}
</style>
