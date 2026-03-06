<template>
  <el-drawer
    v-model="drawer"
    :title="$t('settings.title')"
    direction="rtl"
    :size="width"
    :show-close="false"
  >
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg">{{ $t('settings.title') }}</span>
        <el-button
          type="primary"
          @click="saveConfig"
        >
          {{ $t('common.save') }}
        </el-button>
      </div>
    </template>
    <div class="flex flex-col">
      <div class="mb-8">
        <Title :title="$t('settings.theme.title')" />
        <div class="mt-2 text-sm p-2 flex items-center justify-center gap-2">
          <el-segmented
            v-model="config.darkMode"
            :options="options"
            size="default"
            @change="appStore.toggleDarkMode"
          />
        </div>
      </div>
      <div class="mb-8">
        <Title :title="$t('settings.themeColor')" />
        <div class="mt-2 text-sm p-2 flex items-center gap-2 justify-center">
          <div
            v-for="item in colors"
            :key="item"
            class="w-5 h-5 rounded cursor-pointer flex items-center justify-center"
            :style="`background:${item}`"
            @click="appStore.togglePrimaryColor(item)"
          >
            <el-icon v-if="config.primaryColor === item">
              <Select />
            </el-icon>
          </div>
          <el-color-picker
            v-model="customColor"
            @change="appStore.togglePrimaryColor"
          />
        </div>
      </div>
      <div class="mb-8">
        <Title :title="$t('settings.display.title')" />
        <div class="mt-2 text-md p-2 flex flex-col gap-2">
          <div class="flex items-center justify-between">
            <div>{{ $t('settings.display.watermark') }}</div>
            <el-switch
              v-model="config.show_watermark"
              @change="appStore.toggleConfigWatermark"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ $t('settings.display.greyMode') }}</div>
            <el-switch
              v-model="config.grey"
              @change="appStore.toggleGrey"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ $t('settings.display.weaknessMode') }}</div>
            <el-switch
              v-model="config.weakness"
              @change="appStore.toggleWeakness"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ $t('settings.display.language') }}</div>
            <el-select
              v-model="config.language"
              class="w-40"
              @change="appStore.toggleLanguage"
            >
              <el-option
                v-for="item in LanguageList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </div>
          <div class="flex items-center justify-between">
            <div>{{ $t('settings.display.menuMode') }}</div>
            <el-segmented
              v-model="config.side_mode"
              :options="sideModes"
              size="default"
              @change="appStore.toggleSideMode"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ $t('settings.display.showTabs') }}</div>
            <el-switch
              v-model="config.showTabs"
              @change="appStore.toggleTabs"
            />
          </div>
          <div class="flex items-center justify-between gap-2">
            <div class="flex-shrink-0">
              {{ $t('settings.display.animation') }}
            </div>
            <el-select
              v-model="config.transition_type"
              class="w-40"
              @change="appStore.toggleTransition"
            >
              <el-option
                value="fade"
                :label="$t('settings.animations.fade')"
              />
              <el-option
                value="slide"
                :label="$t('settings.animations.slide')"
              />
              <el-option
                value="zoom"
                :label="$t('settings.animations.zoom')"
              />
              <el-option
                value="none"
                :label="$t('settings.animations.none')"
              />
            </el-select>
          </div>
        </div>
      </div>

      <div class="mb-8">
        <Title :title="$t('settings.layout.title')" />
        <div class="mt-2 text-md p-2 flex flex-col gap-2">
          <div class="flex items-center justify-between mb-2">
            <div>{{ $t('settings.layout.sideWidth') }}</div>
            <el-input-number
              v-model="config.layout_side_width"
              :min="150"
              :max="400"
              :step="10"
            />
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>{{ $t('settings.layout.sideCollapsedWidth') }}</div>
            <el-input-number
              v-model="config.layout_side_collapsed_width"
              :min="60"
              :max="100"
            />
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>{{ $t('settings.layout.sideItemHeight') }}</div>
            <el-input-number
              v-model="config.layout_side_item_height"
              :min="30"
              :max="50"
            />
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>{{ $t('settings.layout.tableHeightEnable') }}</div>
            <el-switch
              v-model="config.table_hight_enable"
            />
          </div>
          <div v-if="config.table_hight_enable" class="flex items-center justify-between mb-2">
            <div>{{ $t('settings.layout.tableHeight') }}</div>
            <el-input-number
              v-model="config.table_hight"
              :min="300"
              :max="1000"
            />
          </div>
        </div>
      </div>

      <!--      <el-alert type="warning" :closable="false">
        请注意，所有配置请保存到本地文件的
        <el-tag>config.json</el-tag> 文件中，否则刷新页面后会丢失配置
      </el-alert>-->
    </div>
  </el-drawer>
</template>

<script setup>
import { useAppStore } from '@/pinia'
import { storeToRefs } from 'pinia'
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { setSelfSetting } from '@/api/user'
import Title from './title.vue'
import { useI18n } from 'vue-i18n'
import { getDict } from '@/utils/dictionary'

const LanguageList = ref([])
getDict('Language').then((res) => {
  LanguageList.value = res
})

const { t } = useI18n()
const appStore = useAppStore()
const { config, device } = storeToRefs(appStore)

defineOptions({
  name: 'HabSetting'
})

const width = computed(() => {
  return device.value === 'mobile' ? '100%' : '500px'
})

const colors = [
  '#EB2F96',
  '#3b82f6',
  '#2FEB54',
  '#EBEB2F',
  '#EB2F2F',
  '#2FEBEB'
]

const drawer = defineModel('drawer', {
  default: true,
  type: Boolean
})

const options = ref([
  {
    value: 'dark',
    label: t('settings.theme.dark')
  },
  {
    value: 'light',
    label: t('settings.theme.light')
  },
  {
    value: 'auto',
    label: t('settings.theme.auto')
  }
])

const sideModes = ref([
  {
    label: t('settings.menuModes.normal'),
    value: 'normal'
  },
  {
    label: t('settings.menuModes.head'),
    value: 'head'
  },
  {
    label: t('settings.menuModes.combination'),
    value: 'combination'
  }
])

const saveConfig = async() => {
  /* const input = document.createElement("textarea");
  input.value = JSON.stringify(config.value);
  // 添加回车
  input.value = input.value.replace(/,/g, ",\n");
  document.body.appendChild(input);
  input.select();
  document.execCommand("copy");
  document.body.removeChild(input);
  ElMessage.success("复制成功, 请自行保存到本地文件中");*/
  const newConfig = { ...config.value }
  newConfig.messages = {}
  const res = await setSelfSetting(newConfig)
  if (res.code === 0) {
    localStorage.setItem('originSetting', JSON.stringify(config.value))
    ElMessage.success(t('common.save'))
    drawer.value = false
  }
}

const customColor = ref('')
</script>

<style lang="scss" scoped>
  ::v-deep(.el-drawer__header) {
    @apply border-gray-400 dark:border-gray-600;
  }
</style>
