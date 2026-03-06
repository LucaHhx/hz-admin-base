import service from '@/utils/request'

// 获取所有时区列表
export const getTimezoneList = () => {
  return service({
    url: '/timezone/list',
    method: 'get'
  })
}

// 根据地区获取时区
export const getTimezonesByRegion = (region) => {
  return service({
    url: '/timezone/region',
    method: 'get',
    params: {
      region: region
    }
  })
}

// 获取当前服务器时区
export const getServerTimezone = () => {
  return service({
    url: '/timezone/server',
    method: 'get'
  })
}

// 时区转换
export const convertTimezone = (data) => {
  return service({
    url: '/timezone/convert',
    method: 'post',
    data
  })
}