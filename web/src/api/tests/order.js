import service from '@/utils/request'
// @Tags Order
// @Summary 创建订单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "创建订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /order/createOrder [post]
export const createOrder = (data) => {
  return service({
    url: '/order/createOrder',
    method: 'post',
    data
  })
}

// @Tags Order
// @Summary 删除订单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "删除订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /order/deleteOrder [delete]
export const deleteOrder = (params) => {
  return service({
    url: '/order/deleteOrder',
    method: 'delete',
    params
  })
}

// @Tags Order
// @Summary 批量删除订单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /order/deleteOrder [delete]
export const deleteOrderByIds = (params) => {
  return service({
    url: '/order/deleteOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags Order
// @Summary 更新订单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "更新订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /order/updateOrder [put]
export const updateOrder = (data) => {
  return service({
    url: '/order/updateOrder',
    method: 'put',
    data
  })
}

// @Tags Order
// @Summary 用id查询订单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Order true "用id查询订单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /order/findOrder [get]
export const findOrder = (params) => {
  return service({
    url: '/order/findOrder',
    method: 'get',
    params
  })
}

// @Tags Order
// @Summary 分页获取订单管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Success"}"
// @Router /order/getOrderList [get]
export const getOrderList = (data) => {
  return service({
    url: '/order/getOrderList',
    method: 'post',
    data
  })
}

// @Tags Order
// @Summary 不需要鉴权的订单管理接口
// @Accept application/json
// @Produce application/json
// @Param data query testsReq.OrderSearch true "分页获取订单管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "Success"
// @Router /order/getOrderPublic [get]
export const getOrderPublic = () => {
  return service({
    url: '/order/getOrderPublic',
    method: 'get',
  })
}

// @Tags Order
// @Summary 导入订单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.ImportOrderReq true "导入Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
export const importOrder = (data) => {
  return service({
    url: '/order/importOrder',
    method: 'post',
    data
  })
}