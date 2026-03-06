<template>
  <el-dialog
    v-model="visible"
    title="点击图片选择坐标"
    :close-on-click-modal="false"
    :destroy-on-close="true"
    :style="dialogStyle"
    width="auto"
    @closed="onClosed"
  >
    <div
      ref="container"
      class="overflow-auto"
      style="display: flex; justify-content: center; align-items: center;"
      @click="handleClick"
    >
      <img
        ref="image"
        :src="imgSrc"
        :style="imgStyle"
        class="select-none"
        draggable="false"
      >
    </div>
  </el-dialog>
</template>

<script setup>
import { ref } from 'vue'

const visible = ref(false)
const imgSrc = ref('')
const imgStyle = ref({})
const dialogStyle = ref({})
const image = ref(null)

let resolveClick = null

function open(imageDataUrl) {
  return new Promise((resolve) => {
    imgSrc.value = imageDataUrl
    visible.value = true
    resolveClick = resolve

    const img = new Image()
    img.src = imageDataUrl
    img.onload = () => {
      const maxW = window.innerWidth * 0.9
      const maxH = window.innerHeight * 0.8
      const scale = Math.min(maxW / img.width, maxH / img.height, 1)
      const w = Math.round(img.width * scale)
      const h = Math.round(img.height * scale)

      imgStyle.value = {
        width: `${w}px`,
        height: `${h}px`,
      }
      dialogStyle.value = {
        width: `${w + 40}px`,
      }
    }
  })
}

function handleClick(e) {
  const rect = image.value.getBoundingClientRect()
  const clickX = e.clientX - rect.left
  const clickY = e.clientY - rect.top

  const naturalWidth = image.value.naturalWidth
  const naturalHeight = image.value.naturalHeight
  const displayWidth = rect.width
  const displayHeight = rect.height

  const x = Math.round((clickX / displayWidth) * naturalWidth)
  const y = Math.round((clickY / displayHeight) * naturalHeight)

  visible.value = false
  resolveClick?.({ x, y })
  resolveClick = null
}

function onClosed() {
  imgSrc.value = ''
  resolveClick = null
}

defineExpose({ open })
</script>

<style lang="scss" scoped>
.text-sm {
  color: #070a11;
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
