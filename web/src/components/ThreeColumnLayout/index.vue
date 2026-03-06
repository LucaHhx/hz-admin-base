<template>
  <div
    class="three-column-container"
    :style="containerStyle"
  >
    <div
      class="section left-section"
      :style="{ gap: itemGap + 'px' }"
    >
      <slot name="left" />
    </div>
    <div
      class="section center-section"
      :style="{ gap: itemGap + 'px' }"
    >
      <slot name="center" />
    </div>
    <div
      class="section right-section"
      :style="{ gap: itemGap + 'px' }"
    >
      <slot name="right" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

defineOptions({
  name: 'ThreeColumnLayout'
})

const props = defineProps({
  // 区块之间的间距
  sectionGap: {
    type: Number,
    default: 16
  },
  // 每个区块内部项目的间距
  itemGap: {
    type: Number,
    default: 8
  }
})

const containerStyle = computed(() => ({
  gap: props.sectionGap + 'px'
}))
</script>

<style scoped>
.three-column-container {
  display: flex;
  align-items: center;
  width: 100%;
  min-height: 40px;
}

.section {
  display: flex;
  align-items: center;
}

.left-section {
  flex: 0 0 auto;
  /* margin-right: 16px; */
}

.center-section {
  flex: 1 1 auto;
  justify-content: center;
}

.right-section {
  flex: 0 0 auto;
  /* margin-left: 16px; */
}

:deep(.el-button + .el-button) {
  margin-left: 0;
}
</style>
