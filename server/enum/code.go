package enum

import "fmt"

type Code int

const (
	// 通用状态码
	Code_Success Code = 0   // 成功
	Code_Failed  Code = 7   // 失败
	Code_NotFind Code = 404 // 未找到资源

	// 商户相关
	Code_MerchantNotFound     Code = 101 // 商户未找到
	Code_MerchantInsufficient Code = 102 // 余额不足
	Code_MerchantSignError    Code = 103 // 签名错误

	// 数据库错误
	Code_DBError            Code = 1001 // 数据库操作失败
	Code_DBRecordNotFound   Code = 1002 // 数据库记录未找到
	Code_DBConnectionFailed Code = 1003 // 数据库连接失败
	Code_DBQueryError       Code = 1004 // 数据库查询错误
	Code_DBInsertError      Code = 1005 // 数据库插入错误
	Code_DBUpdateError      Code = 1006 // 数据库更新错误
	Code_DBDeleteError      Code = 1007 // 数据库删除错误

	// 缓存错误
	Code_CacheError            Code = 2001 // 缓存操作失败
	Code_CacheMiss             Code = 2002 // 缓存未命中
	Code_CacheConnectionFailed Code = 2003 // 缓存连接失败
	Code_CacheSetError         Code = 2004 // 缓存写入失败
	Code_CacheGetError         Code = 2005 // 缓存读取失败

	// 权限相关错误
	Code_Unauthorized Code = 3001 // 未授权
	Code_Forbidden    Code = 3002 // 无权限
	Code_TokenExpired Code = 3003 // Token 过期
	Code_TokenInvalid Code = 3004 // Token 无效

	// 参数校验错误
	Code_InvalidParams Code = 4001 // 参数无效
	Code_MissingParams Code = 4002 // 缺少必要参数
	Code_InvalidFormat Code = 4003 // 格式错误

	// 网络相关错误
	Code_NetworkError   Code = 5001 // 网络错误
	Code_RequestTimeout Code = 5002 // 请求超时
	Code_ServiceDown    Code = 5003 // 服务不可用

	// 文件相关错误
	Code_FileNotFound    Code = 6001 // 文件未找到
	Code_FileUploadError Code = 6002 // 文件上传失败
	Code_FileReadError   Code = 6003 // 文件读取失败
	Code_FileWriteError  Code = 6004 // 文件写入失败

	// 系统内部错误
	Code_InternalError Code = 7001 // 系统内部错误
	Code_ServiceError  Code = 7002 // 服务异常
	Code_ConfigError   Code = 7003 // 配置错误

	// 任务相关错误
	Code_TaskError  Code = 8001 // 任务执行错误
	Code_TaskStop   Code = 8002 // 任务停止
	Code_TaskCancel Code = 8007 // 任务取消

	//Api相关错误
	Code_HeadError Code = 9001 // 请求头错误

	//订单相关
	Code_OrderError       Code = 10001 // 订单错误
	Code_OrderNotFound    Code = 10002 // 订单未找到
	Code_OrderStatusError Code = 10003 // 订单状态错误
	Code_Order_Over       Code = 10004 // 订单已结束

)

type CodeMsg struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeMsg(code Code, format string, a ...any) CodeMsg {
	return CodeMsg{Code: code, Msg: fmt.Sprintf(format, a...)}
}

func (msg CodeMsg) Error() string {
	return msg.Msg
}

var (
	// 通用状态码
	Msg_Success = CodeMsg{Code: Code_Success, Msg: "Operation successful"} // 操作成功
	Msg_Failed  = CodeMsg{Code: Code_Failed, Msg: "Operation failed"}      // 操作失败
	Msg_NotFind = CodeMsg{Code: Code_NotFind, Msg: "Resource not found"}   // 未找到资源

	// 商户相关
	Msg_MerchantNotFound     = CodeMsg{Code: Code_MerchantNotFound, Msg: "Merchant not found"}       // 商户未找到
	Msg_MerchantInsufficient = CodeMsg{Code: Code_MerchantInsufficient, Msg: "Insufficient balance"} // 余额不足
	Msg_SignError            = CodeMsg{Code: Code_MerchantSignError, Msg: "Signature error"}         // 签名错误

	// 数据库错误
	Msg_DBError            = CodeMsg{Code: Code_DBError, Msg: "Database operation failed"}                 // 数据库操作失败
	Msg_DBRecordNotFound   = CodeMsg{Code: Code_DBRecordNotFound, Msg: "Record not found in database"}     // 数据库记录未找到
	Msg_DBConnectionFailed = CodeMsg{Code: Code_DBConnectionFailed, Msg: "Failed to connect to database"}  // 数据库连接失败
	Msg_DBQueryError       = CodeMsg{Code: Code_DBQueryError, Msg: "Error executing database query"}       // 数据库查询错误
	Msg_DBInsertError      = CodeMsg{Code: Code_DBInsertError, Msg: "Failed to insert data into database"} // 数据库插入错误
	Msg_DBUpdateError      = CodeMsg{Code: Code_DBUpdateError, Msg: "Failed to update database record"}    // 数据库更新错误
	Msg_DBDeleteError      = CodeMsg{Code: Code_DBDeleteError, Msg: "Failed to delete database record"}    // 数据库删除错误

	// 缓存错误
	Msg_CacheError            = CodeMsg{Code: Code_CacheError, Msg: "Cache operation failed"}                // 缓存操作失败
	Msg_CacheMiss             = CodeMsg{Code: Code_CacheMiss, Msg: "Cache miss"}                             // 缓存未命中
	Msg_CacheConnectionFailed = CodeMsg{Code: Code_CacheConnectionFailed, Msg: "Failed to connect to cache"} // 缓存连接失败
	Msg_CacheSetError         = CodeMsg{Code: Code_CacheSetError, Msg: "Failed to write data to cache"}      // 缓存写入失败
	Msg_CacheGetError         = CodeMsg{Code: Code_CacheGetError, Msg: "Failed to read data from cache"}     // 缓存读取失败

	// 权限相关错误
	Msg_Unauthorized = CodeMsg{Code: Code_Unauthorized, Msg: "Unauthorized access"} // 未授权
	Msg_Forbidden    = CodeMsg{Code: Code_Forbidden, Msg: "Access forbidden"}       // 无权限
	Msg_TokenExpired = CodeMsg{Code: Code_TokenExpired, Msg: "Token has expired"}   // Token 过期
	Msg_TokenInvalid = CodeMsg{Code: Code_TokenInvalid, Msg: "Invalid token"}       // Token 无效

	// 参数校验错误
	Msg_InvalidParams = CodeMsg{Code: Code_InvalidParams, Msg: "Invalid parameters"}          // 参数无效
	Msg_MissingParams = CodeMsg{Code: Code_MissingParams, Msg: "Missing required parameters"} // 缺少必要参数
	Msg_InvalidFormat = CodeMsg{Code: Code_InvalidFormat, Msg: "Invalid data format"}         // 格式错误

	// 网络相关错误
	Msg_NetworkError   = CodeMsg{Code: Code_NetworkError, Msg: "Network error occurred"} // 网络错误
	Msg_RequestTimeout = CodeMsg{Code: Code_RequestTimeout, Msg: "Request timed out"}    // 请求超时
	Msg_ServiceDown    = CodeMsg{Code: Code_ServiceDown, Msg: "Service is unavailable"}  // 服务不可用

	// 文件相关错误
	Msg_FileNotFound    = CodeMsg{Code: Code_FileNotFound, Msg: "File not found"}           // 文件未找到
	Msg_FileUploadError = CodeMsg{Code: Code_FileUploadError, Msg: "Failed to upload file"} // 文件上传失败
	Msg_FileReadError   = CodeMsg{Code: Code_FileReadError, Msg: "Failed to read file"}     // 文件读取失败
	Msg_FileWriteError  = CodeMsg{Code: Code_FileWriteError, Msg: "Failed to write file"}   // 文件写入失败

	// 系统内部错误
	Msg_InternalError = CodeMsg{Code: Code_InternalError, Msg: "Internal server error"}                  // 系统内部错误
	Msg_ServiceError  = CodeMsg{Code: Code_ServiceError, Msg: "Service encountered an error"}            // 服务异常
	Msg_ConfigError   = CodeMsg{Code: Code_ConfigError, Msg: "Configuration error"}                      // 配置错误
	Msg_Sys_Busy      = CodeMsg{Code: Code_InternalError, Msg: "System is busy, please try again later"} // 系统繁忙

	Msg_TaskError  = CodeMsg{Code: Code_TaskError, Msg: "Task execution error"} // 任务执行错误
	Msg_TaskStop   = CodeMsg{Code: Code_TaskStop, Msg: "Task stopped"}          // 任务停止
	Msg_TaskCancel = CodeMsg{Code: Code_TaskCancel, Msg: "Task canceled"}       // 任务取消

	Msg_HeadError = CodeMsg{Code: Code_HeadError, Msg: "Request header error"} // 请求头错误

	//订单相关
	Msg_OrderError       = CodeMsg{Code: Code_OrderError, Msg: "Order error"}              // 订单错误
	Msg_OrderNotFound    = CodeMsg{Code: Code_OrderNotFound, Msg: "Order not found"}       // 订单未找到
	Msg_OrderStatusError = CodeMsg{Code: Code_OrderStatusError, Msg: "Order status error"} // 订单状态错误
	Msg_Order_Over       = CodeMsg{Code: Code_Order_Over, Msg: "Order has ended"}          // 订单已结束
)
