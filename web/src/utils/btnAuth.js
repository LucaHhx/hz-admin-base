import { useRoute } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import { reactive } from 'vue'

export const useBtnAuth = () => {
  const route = useRoute()
  return route.meta.btns || reactive({})
}

export const useBtnAuthForRoute = (name) => {
  const routerStore = useRouterStore()
  return routerStore.routeMap[name]?.meta?.btns || reactive({})
}
