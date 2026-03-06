import service from '@/utils/request'

export const getAuthorityColApi = (data) => {
  return service({
    url: '/authorityCol/getAuthorityCol',
    method: 'post',
    data
  })
}

export const setAuthorityColApi = (data) => {
  return service({
    url: '/authorityCol/setAuthorityCol',
    method: 'post',
    data
  })
}

export const getMenuColumnsApi = (menuName) => {
  return service({
    url: '/authorityCol/getMenuColumns',
    method: 'get',
    params: {
      menuName
    }
  })
}
