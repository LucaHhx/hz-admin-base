package enum

type BankType string

const (
	BankCode_BCA     BankType = "bca" // 农业银行
	BankCode_BRI     BankType = "bri"
	BankCode_BNI     BankType = "bni"
	BankCode_MANDIRI BankType = "mandiri"
	BankCode_CIMB    BankType = "cimb"
	BankCode_DANA    BankType = "dana"
	BankCode_BCA_P   BankType = "bca_p"
)

type AccType string

const (
	AccTypeCommon AccType = "common"
	AccTypeMake   AccType = "make"
	AccTypeAudit  AccType = "audit"
)

type ImportType string

const (
	ImportType_Full   ImportType = "full"   // 全量导入
	ImportType_Append ImportType = "append" // 增量导入
)

type GrpcType string

const (
	GrpcType_Unknown GrpcType = "unknown" // 未知
	GrpcType_OpenVc  GrpcType = "openvc"  // OpenVC
	GrpcType_Serial  GrpcType = "serial"  // 串口
	GrpcType_Chrome  GrpcType = "chrome"  // Chrome
	GrpcType_Serve   GrpcType = "serve"   // Serve
)

type OrderStatus string

const (
	OrderStatus_Unknown    OrderStatus = "unknown"    // 未知
	OrderStatus_Created    OrderStatus = "created"    // 创建
	OrderStatus_Submit     OrderStatus = "submit"     // 提交
	OrderStatus_Success    OrderStatus = "success"    // 成功
	OrderStatus_Fail       OrderStatus = "fail"       // 失败
	OrderStatus_Processing OrderStatus = "processing" // 处理中
	OrderStatus_Timeout    OrderStatus = "timeout"    // 超时
	OrderStatus_Checked    OrderStatus = "checked"    // 已核对
	OrderStatus_PinUse     OrderStatus = "pinUse"     // PIN已使用
)

type Status string

const (
	Status_Unknown    Status = "unknown"    // 未知
	Status_Success    Status = "success"    // 成功
	Status_Fail       Status = "fail"       // 失败
	Status_Processing Status = "processing" // 处理中
	Status_Timeout    Status = "timeout"    // 超时
)

type CallBackStatus string

const (
	CallBackStatus_Unknown CallBackStatus = "unknown" // 未知
	CallBackStatus_Success CallBackStatus = "success" // 成功
	CallBackStatus_Fail    CallBackStatus = "fail"    // 失败
)

type AccStatus string

const (
	AccStatus_Unknown    AccStatus = "unknown"    // 未知
	AccStatus_Offline    AccStatus = "offline"    // 离线
	AccStatus_Online     AccStatus = "online"     // 在线
	AccStatus_Vpn_Online AccStatus = "vpn_online" // 在线
)

type UseStatus string

const (
	UseStatus_Normal   UseStatus = "normal"   // 正常
	UseStatus_Disabled UseStatus = "disabled" // 停用
)

type TranStatus string

const (
	TranStatus_Unknown    TranStatus = "unknown"    // 未知
	TranStatus_Success    TranStatus = "success"    // 成功
	TranStatus_Fail       TranStatus = "fail"       // 失败
	TranStatus_Processing TranStatus = "processing" // 处理中
	TranStatus_Timeout    TranStatus = "timeout"    // 超时
	TranStatus_Refund     TranStatus = "refund"     // 退款
)

type WriteStatus string

const (
	WriteStatus_All     WriteStatus = "all"     // 未知
	WriteStatus_Success WriteStatus = "success" // 成功
	WriteStatus_Fail    WriteStatus = "fail"    // 失败
)

type TranType string

const (
	TranType_CR TranType = "CR" // 入账（Credit）
	TranType_DB TranType = "DB" // 出账（Debit）
	TranType_BP TranType = "BP" // 批次（BatchPay）
)

type MatchType string

const (
	MatchType_Unknown MatchType = "unknown"      // 未知
	MatchType_Amount  MatchType = "match_amount" // 金额匹配
	MatchType_Note    MatchType = "match_note"   // 备注匹配
)

type UserAccType string

const (
	UserAccType_Identical UserAccType = "identical" // 同名账户
	UserAccType_Different UserAccType = "different" // 异名账户
)

type PayoutOptType string

const (
	PayoutOptType_Preserve PayoutOptType = "preserve" // 保留
	PayoutOptType_End      PayoutOptType = "end"      // 结束
	PayoutOptType_Retry    PayoutOptType = "retry"    // 重试
	PayoutOptType_Success  PayoutOptType = "success"  // 成功
)

type TimelyType string

const (
	TimelyType_AccStatusChange TimelyType = "acc_status_change" // 账号状态变更
	TimelyType_AccPayout       TimelyType = "acc_payout"        // 账号出账
	TimelyType_AccBalance      TimelyType = "acc_balance"       // 账号余额
	TimelyType_AccPayinMonitor TimelyType = "acc_payin_monitor" // 代收监测下发
)

type TimelyExecuteType string

const (
	TimelyExecuteType_Once   TimelyExecuteType = "once"   // 单次执行
	TimelyExecuteType_Day    TimelyExecuteType = "day"    // 循环
	TimelyExecuteType_Minute TimelyExecuteType = "minute" // 循环
)

type AccountAttr string

const (
	AccountAttrIn  = "in"  // 代收
	AccountAttrOut = "out" // 代付
)

type PayStatus string

const (
	PayStatusReady      PayStatus = "ready"      // 就绪
	PayStatusProcessing PayStatus = "processing" // 处理中
	PayStatusErr        PayStatus = "error"      // 异常
)

type BatchPayStatus string

const (
	BatchPayStatusPending    BatchPayStatus = "pending"    //待处理
	BatchPayStatusProcessing BatchPayStatus = "processing" //处理中：已经提交发货批次
	BatchPayStatusMake       BatchPayStatus = "make_order" //已制单：已经在对应业务账户系统内实现制单并提交
	BatchPayStatusWaitMatch  BatchPayStatus = "wait_match" //待匹配：已经在对应业务账户系统内实现订单审核
	BatchPayStatusSuccess    BatchPayStatus = "success"    //成功：批次内任意订单关联流水成功
	BatchPayStatusFail       BatchPayStatus = "fail"       //失败
	BatchPayStatusTimeout    BatchPayStatus = "timeout"    //超时
)

type BatchPayRuleStatus string

const (
	BatchPayRuleStatusAuto   BatchPayRuleStatus = "auto"
	BatchPayRuleStatusManual BatchPayRuleStatus = "manual"
	BatchPayRuleStatusClose  BatchPayRuleStatus = "close"
)

type BatchPayOrderType string

const (
	BatchPayOrderTypeInHouse BatchPayOrderType = "in-house" // 同行
	BatchPayOrderTypeBiFast  BatchPayOrderType = "bi-fast"  // bi-fast跨行
)

type BatchPayOrderStatus string

const (
	BatchPayOrderStatusInit       BatchPayOrderStatus = "init"       //初始：还没有进入到预分配池子的订单
	BatchPayOrderStatusChecked    BatchPayOrderStatus = "checked"    //已验证
	BatchPayOrderStatusPending    BatchPayOrderStatus = "pending"    //待处理：进入到池子，但是还没有成功提交到发货批次的订单
	BatchPayOrderStatusProcessing BatchPayOrderStatus = "processing" //处理中：已经提交发货批次的订单
	BatchPayOrderStatusMake       BatchPayOrderStatus = "make_order" //已制单：已经在对应业务账户系统内实现订单提交
	BatchPayOrderStatusWaitMatch  BatchPayOrderStatus = "wait_match" //待匹配：已经在对应业务账户系统内实现订单审核
	BatchPayOrderStatusSuccess    BatchPayOrderStatus = "success"    //成功：批次内任意订单关联流水成功
	BatchPayOrderStatusFail       BatchPayOrderStatus = "fail"       //失败
	BatchPayOrderStatusTimeout    BatchPayOrderStatus = "timeout"    //超时
)

type AccCheckStatus string

const (
	AccCheckStatusNone AccCheckStatus = "0" // 未知
	AccCheckStatusYes  AccCheckStatus = "1" // 正确
	AccCheckStatusNo   AccCheckStatus = "2" // 失败
)

type RelayType string

const (
	RelayType_Old RelayType = "1" // 串口
	RelayType_New RelayType = "2" // 网络
)

type AdbType string

const (
	AdbType_Tap   AdbType = "tap"   // 点击
	AdbType_Swd   AdbType = "text"  // 滑动
	AdbType_Start AdbType = "start" // 启动
	AdbType_Stop  AdbType = "stop"  // 停止
	AdbType_Key   AdbType = "key"   // 按键
)

type RecStatus string

const (
	RecStatus_Normal     RecStatus = "normal"     // 正常
	RecStatus_Backward   RecStatus = "backward"   // 正常
	RecStatus_Additional RecStatus = "additional" // 追加
	RecStatus_Pend       RecStatus = "pend"       // 追加
	RecStatus_Timeout    RecStatus = "timeout"
	RecStatus_Confirm    RecStatus = "confirm"
)

type TranSnaType string

const (
	TranSnaType_Tran      TranSnaType = "tran"
	TranSnaType_Yesterday TranSnaType = "yesterday"
	TranSnaType_Today     TranSnaType = "today"
	TranSnaType_Month     TranSnaType = "month"
)

type BcaTranType string

const (
	BcaTranType_OtherBca BcaTranType = "Other BCA Account"
	BcaTranType_Domestic BcaTranType = "Domestic Account"
)

type VoucherType string

const (
	VoucherType_Html VoucherType = "html"
	VoucherType_Png  VoucherType = "png"
)

type OrderType string

const (
	OrderType_Payin        OrderType = "payin"        // 代收
	OrderType_Payout       OrderType = "payout"       //下发
	OrderType_Batch        OrderType = "batch"        // 代付
	OrderType_Fee          OrderType = "fee"          // 手续费
	OrderType_Inter_payout OrderType = "inter_payout" // 手续费
	OrderType_Inter_payin  OrderType = "inter_payin"  // 手续费
	OrderType_Test         OrderType = "test"         // 手续费
)

type UsageType string

const (
	UsageType_Domestic UsageType = "domestic" // 跨行
	Usagetype_InHouse  UsageType = "inhouse"  // 同行
	UsageType_RTGS     UsageType = "rtgs"     //大额
)
