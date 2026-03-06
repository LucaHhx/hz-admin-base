import service from '@/utils/request'

// 根据结构体名称获取表格列
export const getStructNameColumns = (structName) => {
  return service({
    url: '/sysTableColumns/getStructNameColumns',
    method: 'get',
    params: {
      structName
    }
  })
}
