package utils

const (
	// SUCCESS 成功
	SUCCESS    = 1000
	// ERROR 异常
	ERROR      = 1002
	// EXCEPTION 出现异常
	EXCEPTION  = 1004 
	// PARA_MISS 缺少必要参数
	PARA_MISS  = 1006
	// PARA_ERROR 参数格式错误
	PARA_ERROR = 1008 

	TOKEN_MISSING           = 1010 //token缺失
	// TOKEN_NOT_FOUND         = 1012 //查无此token
	TOKEN_INVALID           = 1014 //token不匹配
	// USER_NOT_FOUND          = 1016 //用户不存在
	LOGIN_FAIL              = 1018 //登录失败
	HAS_EXIST               = 1020 //已存在
	NOT_EXIST               = 1022 //不存在
	// HTTP_METHOD_ERROR       = 1024 //HTTP请求类型错误
	// NO_PERMISSION           = 1026 //当前用户没有权限操作
	// UNIONID_MISS_PERMISSION = 1028 //用户没有授权获取unionId，需请求权限获得unionId
	// SEND_SUCCESS            = 1030 //发送成功，请等待wamp回复
	// NONE_FOUND              = 1034 //查询为空
	SYS_BUSY                = 1036 //系统繁忙，请稍后再试
	// CODE_INVALID            = 1038 //code无效
	// DECYPT_FAIL             = 1040 //解密失败
	// APPLY_FAIL              = 1042 //申请失败
	// NAME_UNVERIFIED         = 1044 //未实名
	// API_UNDEFINED           = 1100 // 接口未定义
	// NO_DEVICE_PERMISSION    = 1102 //无设备权限
	
	TOKEN_EXPIRE			= 1046 //Token过期
	Token_Faild				= 1048 //Token格式错误
	DBError = 3001 //数据库操作错误
	/*
		InformError    = 5001
		ArgumentsError = 5002 //参数错误*/

)
