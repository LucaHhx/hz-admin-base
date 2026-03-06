<template>
  <el-dialog
    v-model="visible"
    title="图片预览"
    :style="dialogStyle"
    :close-on-click-modal="false"
    :destroy-on-close="true"
    class="image-dialog"
    modal-class="dialog_modal"
    draggable
    :modal="false"
  >
    <div class="flex h-full overflow-hidden">
      <!-- 图片区域 -->
      <div
        ref="container"
        class="flex-1 overflow-scroll border border-gray-200 relative"
        @mousedown="startDrag"
        @mouseup="stopDrag"
        @mouseleave="stopDrag"
        @mousemove="onDrag"
      >
        <img
          ref="image"
          :src="imgSrc"
          alt="预览图"
          class="select-none"
          draggable="false"
          @click="onClick"
        >
      </div>

      <!-- 坐标列表 -->
      <div class="w-64 p-2 border-l border-gray-300 bg-gray-50 overflow-y-auto">
        <h4 class="text-base font-semibold mb-2">
          点击坐标列表
        </h4>
        <el-scrollbar height="80%">
          <ul class="space-y-1">
            <li v-for="(pt, index) in points" :key="index" class="text-sm">
              #{{ index + 1 }}: ({{ pt.x }}, {{ pt.y }})
            </li>
          </ul>
        </el-scrollbar>
      </div>
    </div>

    <!-- info 显示区域 -->
    <div class="mt-4 px-4">
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

    <!-- OCR结果展示区域 -->
    <div v-if="currentUuid" class="mt-4 px-4">
      <div v-if="ocrData" class="ocr-result-container">
        <h4 class="text-base font-semibold mb-2">
          AI识别结果
        </h4>
        <div class="ocr-result-item">
          <span class="ocr-label">识别类型:</span>
          <span class="ocr-value">{{ ocrData.ocrType || '-' }}</span>
        </div>
        <div class="ocr-result-item">
          <span class="ocr-label">识别结果:</span>
          <span class="ocr-value">{{ ocrData.ocrResult || '-' }}</span>
        </div>
        <div class="ocr-result-item">
          <span class="ocr-label">识别内容:</span>
          <span class="ocr-value">{{ ocrData.ocrContent || '-' }}</span>
        </div>
      </div>
      <div v-else class="ocr-loading">
        <el-icon class="is-loading">
          <Loading />
        </el-icon>
        <span class="ml-2">AI识别中，请稍候...</span>
      </div>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref, nextTick, watch, onBeforeUnmount } from 'vue'
import { findBOcrLogByUuid } from '@/api/business/bOcrLog'
import { Loading } from '@element-plus/icons-vue'

const visible = ref(false)
const imgSrc = ref('')
const dialogStyle = ref({})
const points = ref([])
const info = ref('')
const currentUuid = ref('')
const ocrData = ref(null)
const pollingTimer = ref(null)

const container = ref(null)
const image = ref(null)

let dragging = false
let lastX = 0
let lastY = 0

// 清理轮询定时器
const clearPolling = () => {
  if (pollingTimer.value) {
    clearInterval(pollingTimer.value)
    pollingTimer.value = null
  }
}

// 轮询查询OCR结果
const pollOcrResult = async() => {
  if (!currentUuid.value) {
    return
  }

  try {
    const res = await findBOcrLogByUuid({ uuid: currentUuid.value })
    if (res.code === 0 && res.data) {
      ocrData.value = res.data
      // 如果获取到结果，停止轮询
      clearPolling()
    }
  } catch (error) {
    console.error('Poll OCR result failed:', error)
  }
}

function open(dataUrl, infostr, uuid) {
  console.log('uuid', uuid)

  imgSrc.value = dataUrl
  points.value = []
  visible.value = true
  info.value = infostr
  currentUuid.value = uuid || ''
  ocrData.value = null

  // 清除之前的轮询
  clearPolling()

  // 如果有UUID，开始轮询
  if (currentUuid.value) {
    // 立即查询一次
    pollOcrResult()
    // 每5秒轮询一次
    pollingTimer.value = setInterval(pollOcrResult, 5000)
  }

  const img = new Image()
  img.src = dataUrl
  img.onload = () => {
    const maxW = window.innerWidth * 0.9
    const maxH = window.innerHeight * 0.8
    const scale = Math.min(maxW / img.width, maxH / img.height, 1)
    const w = Math.round(img.width * scale)
    const h = Math.round(img.height * scale)

    dialogStyle.value = {
      width: `${w + 260}px`,
      height: `${h + 300}px`, // 增加 info 区域高度
    }
  }
}

function startDrag(e) {
  dragging = true
  lastX = e.clientX
  lastY = e.clientY
}
function stopDrag() {
  dragging = false
}
function onDrag(e) {
  if (!dragging) return
  const dx = e.clientX - lastX
  const dy = e.clientY - lastY
  container.value.scrollLeft -= dx
  container.value.scrollTop -= dy
  lastX = e.clientX
  lastY = e.clientY
}

function onClick(e) {
  const rect = image.value.getBoundingClientRect()
  const x = Math.round(e.clientX - rect.left)
  const y = Math.round(e.clientY - rect.top)
  points.value.push({ x, y })
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
  if (val) {
    makeDialogDraggable()
  } else {
    // 页面关闭时清理轮询
    clearPolling()
  }
})

// 组件卸载时清理轮询
onBeforeUnmount(() => {
  clearPolling()
})

defineExpose({ open })
</script>

  <style lang="scss" scoped>
  .text-sm {
    color: #070a11;
  }

  .ocr-result-container {
    padding: 16px;
    background: #f5f7fa;
    border-radius: 4px;
    border: 1px solid #e4e7ed;
  }

  .ocr-result-item {
    margin-bottom: 12px;
    display: flex;
    align-items: flex-start;

    &:last-child {
      margin-bottom: 0;
    }
  }

  .ocr-label {
    font-weight: 600;
    color: #606266;
    min-width: 80px;
    margin-right: 8px;
  }

  .ocr-value {
    color: #303133;
    flex: 1;
    word-break: break-all;
  }

  .ocr-loading {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 24px;
    background: #f5f7fa;
    border-radius: 4px;
    border: 1px solid #e4e7ed;
    color: #909399;
  }
  </style>

  <style scoped>
  img {
    display: block;
    max-width: none;
    cursor: url('data:image/svg+xml;utf8,\
  <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64">\
  <line x1="0" y1="32" x2="64" y2="32" stroke="red" stroke-width="1"/>\
  <line x1="32" y1="0" x2="32" y2="64" stroke="red" stroke-width="1"/>\
  </svg>') 32 32, crosshair;
  }
  </style>
