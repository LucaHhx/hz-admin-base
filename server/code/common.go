package code

var (
	Err_Common_InvalidParams = NewError(CodeCommonError, "invalid_params")
	Err_Common_NoPermission  = NewError(CodeCommonError, "no_permission")
	Err_Common_NoData        = NewError(CodeCommonError, "no_data")
	Err_Common_Unknown       = NewError(CodeCommonError, "unknown_error")
	Err_Common_Err           = NewError(CodeCommonError, "error")
)
