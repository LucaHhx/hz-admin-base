<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-carousel class="-mt-2">
    <el-carousel-item
      v-for="(item, index) in banners"
      :key="index"
      class="cursor-pointer lg:h-40"
      @click="openLink(item.link)"
    >
      <el-image
        class="h-full w-full"
        :src="item.img"
        fit="fill"
      />
    </el-carousel-item>
  </el-carousel>
</template>

<script setup>
  import banner from '@/assets/banner.svg'

  const openLink = (link) => {
    window.open(link, '_blank')
  }

  const banners = [
    {
      img: banner,
      link: '#'
    }
  ]
</script>

<style scoped lang="scss"></style>
