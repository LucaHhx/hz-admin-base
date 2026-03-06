<template>
  <div class="mt-2">
    <div class="flex flex-col md:flex-row gap-4">
      <div class="w-full md:w-1/2">
        <el-card class="min-w-96">
          <template #header>
            <el-divider>{{ $t('about.projectTitle') }}</el-divider>
          </template>
          <div>
            <div class="w-full flex items-center justify-center">
              <span class="text-gray-500 text-xl font-bold">HAB Admin</span>
            </div>
          </div>
        </el-card>
        <el-card class="min-w-96 mt-5">
          <template #header>
            <div>{{ $t('about.team') }}</div>
          </template>
          <div>
            <div class="w-full flex items-center justify-center">
              <span class="text-gray-500">HAB Team</span>
            </div>
          </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
defineOptions({
  name: 'About'
})
</script>
