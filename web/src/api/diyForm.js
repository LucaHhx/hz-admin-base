import service from '@/utils/request'

// 根据类型获取表单配置
export const getFormConfig = (formType) => {
  return service({
    url: '/form/config',
    method: 'get',
    params: {
      type: formType
    }
  })
}

// 提交表单数据
export const submitFormData = (formType, data) => {
  return service({
    url: '/form/submit',
    method: 'post',
    data: {
      type: formType,
      data: data
    }
  })
}

// 获取表单初始数据
export const getFormInitData = (formType) => {
  return service({
    url: '/form/init',
    method: 'get',
    params: {
      type: formType
    }
  })
}