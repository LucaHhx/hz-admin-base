<template>
  <div
    class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 py-2 gap-4 md:gap-2 hab-container2"
  >
    <hab-card custom-class="col-span-1 lg:col-span-2 ">
      <hab-chart
        :type="1"
        title="访问人数"
      />
    </hab-card>
    <hab-card custom-class="col-span-1 lg:col-span-2  ">
      <hab-chart
        :type="2"
        title="新增客户"
      />
    </hab-card>
    <hab-card custom-class="col-span-1 lg:col-span-2 ">
      <hab-chart
        :type="3"
        title="解决数量"
      />
    </hab-card>
    <hab-card
      title="快捷功能"
      show-action
      custom-class="col-start-1 md:col-start-3 lg:col-start-7 row-span-2 "
    >
      <hab-quick-link />
    </hab-card>
    <hab-card
      title="内容数据"
      custom-class="col-span-1 md:col-span-2 md:row-start-2 lg:col-span-6 col-start-1 row-span-2"
    >
      <hab-chart :type="4" />
    </hab-card>
    <hab-card
      title="文档"
      show-action
      custom-class="md:row-start-8 md:col-start-3 lg:row-start-3 lg:col-start-7"
    >
      <hab-wiki />
    </hab-card>

    <hab-card
      title="最新更新"
      custom-class="col-span-1 md:col-span-3 row-span-2"
    >
      <hab-table />
    </hab-card>
    <hab-card
      title="最新插件"
      custom-class="col-span-1 md:col-span-3 row-span-2"
    >
      <hab-plugin-table />
    </hab-card>

    <hab-card
      title="公告"
      show-action
      custom-class="col-span-1 lg:col-start-7"
    >
      <hab-notice />
    </hab-card>

    <hab-card
      without-padding
      custom-class="overflow-hidden lg:h-40 col-span-1 md:col-start-2 md:col-span-1 lg:col-start-7"
    >
      <hab-banner />
    </hab-card>
  </div>
</template>

<script setup>
  import {
    HabPluginTable,
    HabTable,
    HabChart,
    HabWiki,
    HabNotice,
    HabQuickLink,
    HabCard,
    HabBanner
  } from './components'
  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>
