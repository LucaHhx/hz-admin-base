<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div
    class="flex justify-between fixed top-0 left-0 right-0 z-10 h-16 bg-white text-slate-700 dark:text-slate-300 dark:bg-slate-900 shadow dark:shadow-gray-700 items-center px-2"
  >
    <div class="flex items-center cursor-pointer flex-1">
      <div
        class="flex items-center cursor-pointer"
        :class="isMobile ? '' : 'min-w-48'"
        @click="router.push({ path: '/' })"
      >
        <img
          alt
          class="h-12 bg-white rounded-full"
          :src="$GIN_VUE_ADMIN.appLogo"
        >
        <div
          v-if="!isMobile"
          class="inline-flex font-bold text-2xl ml-2"
          :class="
            (config.side_mode === 'head' ||
              config.side_mode === 'combination') &&
              'min-w-fit'
          "
        >
          {{ getPageTitle() }}
        </div>
      </div>

      <el-breadcrumb
        v-show="!isMobile"
        v-if="config.side_mode !== 'head' && config.side_mode !== 'combination'"
        class="ml-4"
      >
        <!-- <el-breadcrumb-item
          v-for="item in matched.slice(1, matched.length)"
          :key="item.path"
        >
          {{ fmtTitle(item.meta.title, route) }}
        </el-breadcrumb-item> -->
      </el-breadcrumb>
      <gva-aside
        v-if="config.side_mode === 'head' && !isMobile"
        class="flex-1"
      />
      <gva-aside
        v-if="config.side_mode === 'combination' && !isMobile"
        mode="head"
        class="flex-1"
      />
    </div>

    <div class="ml-2 flex items-center">
      <div class="mr-4 text-sm">
        {{ $serverDate(currentTime,'long') }}
      </div>
      <tools />
      <el-dropdown>
        <div class="flex justify-center items-center h-full w-full">
          <span
            class="cursor-pointer flex justify-center items-center text-black dark:text-gray-100"
          >
            <CustomPic />
            <span
              v-show="!isMobile"
              class="w-16"
            >
              {{
                userStore.userInfo.nickName
              }}
            </span>
            <el-icon>
              <arrow-down />
            </el-icon>
          </span>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <span class="font-bold">
                {{ $t('user.currentRole') }}: {{ userStore.userInfo.authority.authorityName }}
              </span>
            </el-dropdown-item>
            <template v-if="userStore.userInfo.authorities">
              <el-dropdown-item
                v-for="item in userStore.userInfo.authorities.filter(
                  (i) => i.authorityId !== userStore.userInfo.authorityId
                )"
                :key="item.authorityId"
                @click="changeUserAuth(item.authorityId)"
              >
                <span> {{ $t('user.switchTo') }}: {{ item.authorityName }} </span>
              </el-dropdown-item>
            </template>
            <el-dropdown-item
              icon="avatar"
              @click="toPerson"
            >
              {{ $t('user.personalSettings') }}
            </el-dropdown-item>
            <el-dropdown-item
              icon="reading-lamp"
              @click="userStore.LoginOut"
            >
              {{ $t('user.logout') }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import tools from './tools.vue'
import CustomPic from '@/components/customPic/index.vue'
import { useUserStore } from '@/pinia/modules/user'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/pinia'
import { storeToRefs } from 'pinia'
import { computed, ref, onMounted, onUnmounted } from 'vue'
import { setUserAuthority } from '@/api/user'
import { fmtTitle } from '@/utils/fmtRouterTitle'
import gvaAside from '@/view/layout/aside/index.vue'
import getPageTitle from '@/utils/page'

const userStore = useUserStore()
const router = useRouter()
const route = useRoute()
const appStore = useAppStore()
const { device, config } = storeToRefs(appStore)
const isMobile = computed(() => {
  return device.value === 'mobile'
})
const toPerson = () => {
  router.push({ name: 'person' })
}
const matched = computed(() => route.meta.matched)

const currentTime = ref('')

const updateTime = () => {
  const now = new Date()
  currentTime.value = now.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: false
  })
}

let timer = null

onMounted(() => {
  updateTime()
  timer = setInterval(updateTime, 1000)
})

onUnmounted(() => {
  if (timer) {
    clearInterval(timer)
  }
})

const changeUserAuth = async(id) => {
  const res = await setUserAuthority({
    authorityId: id
  })
  if (res.code === 0) {
    window.sessionStorage.setItem('needCloseAll', 'true')
    window.sessionStorage.setItem('needToHome', 'true')
    window.location.reload()
  }
}
</script>

<style scoped lang="scss"></style>
