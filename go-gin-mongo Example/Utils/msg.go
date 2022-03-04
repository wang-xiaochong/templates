package utils

// MsgFlags 信息反馈
var MsgFlags = map[int]string{
	SUCCESS: "OK",
	ERROR:   "fail",
	EXCEPTION:  "出现异常",
	PARA_MISS:  "缺少必要参数",
	PARA_ERROR: "参数格式错误",

	// USER_NOT_FOUND:          "用户不存在",
	LOGIN_FAIL:              "登录失败",
	HAS_EXIST:               "已存在",
	NOT_EXIST:               "不存在",
	// HTTP_METHOD_ERROR:       "HTTP请求类型错误",
	// NO_PERMISSION:           "当前用户没有权限操作",
	// UNIONID_MISS_PERMISSION: "用户没有授权获取unionId",
	// SEND_SUCCESS:            "发送成功，请等待wamp回复",
	// NONE_FOUND:              "查询为空",
	SYS_BUSY:                "系统繁忙，请稍后再试",
	// CODE_INVALID:            "code无效",
	// DECYPT_FAIL:             "解密失败",
	// APPLY_FAIL:              "申请失败",
	// NAME_UNVERIFIED:         "未实名",

	// API_UNDEFINED:        "接口未定义",
	// NO_DEVICE_PERMISSION: "无设备权限",
	TOKEN_MISSING:        "token缺失",
	// TOKEN_NOT_FOUND:      "查无此token",
	TOKEN_INVALID:        "token不匹配",
	DBError:              "数据库操作错误",

	TOKEN_EXPIRE:			"Token过期",
	Token_Faild:		"Token格式错误",
	/*
		InformError:    "放入表单错误",
		ArgumentsError: "输入参数错误",*/
}

// GetMsg 通过状态码拿到信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
