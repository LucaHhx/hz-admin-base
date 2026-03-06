import { useRoute } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'
import { reactive } from 'vue'
export const useColsAuth = () => {
  const route = useRoute()
  // console.log('route', route)
  return route.meta.cols || reactive({})
}

export const useColsForRoute = (name) => {
  const routerStore = useRouterStore()
  // console.log('route', routerStore.routeMap[name])
  return routerStore.routeMap[name]?.meta?.cols || reactive({})
}
