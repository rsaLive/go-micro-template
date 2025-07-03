package msg

const (
	SERVER_CONFIG         = "config.yaml"
	SERVER_DUBBOGO_CONFIG = "dubbogo.yaml"
	MODE_ENV              = "MODE_ENV"
)

const (
	Success = "操作成功"
	Failed  = "操作失败"
)

const (
	Http = 200
)

const (
	ErrorLogin        = "账号或密码错误"
	ErrorJSONParse    = "json解析失败"
	ErrorJSONMarshal  = "json序列化失败"
	ErrorForUpdate    = "锁定错误"
	ErrorInsert       = "插入异常"
	ErrorDelete       = "删除异常"
	ErrorUpdate       = "更新异常"
	ErrorSelect       = "查询异常"
	ErrorEmptyParam   = "值为空"
	ErrorCopierStruct = "拷贝结构体错误"
	ErrorNoAction     = "无需操作"
	ErrorInvalidParam = "参数不合法"
	ErrorNoData       = "没有数据"
	ErrorDatetime     = "时间格式错误"
	ErrorUserExists   = "账号已存在"
	ErrorCanNotEdit   = "不可以编辑"
	ErrorCanNotDel    = "不可以删除"

	ErrorCreateRoom = "创建房间错误"

	ErrorSha256Write     = "sha256加密错误"
	ErrorInvalidStatus   = "不合法的状态"
	ErrorUnKnowAction    = "未知操作"
	ErrorApprovalIDEmpty = "审批单号为空"
	ErrorCanNotUnBind    = "无法解绑"
	ErrorNoBind          = "未绑定"
)
