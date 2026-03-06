package code

var (
	// 通用业务错误
	Err_Business_Create = NewError(CodeBusinessError, "create")
	Err_Business_Import = NewError(CodeBusinessError, "import")
	Err_Business_Delete = NewError(CodeBusinessError, "delete")
	Err_Business_Update = NewError(CodeBusinessError, "update")
	Err_Business_Get    = NewError(CodeBusinessError, "get")
	Err_Business_List   = NewError(CodeBusinessError, "list")
	Err_Business_Export = NewError(CodeBusinessError, "export")
)
