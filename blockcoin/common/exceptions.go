package common


type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	DevInfo  string `json:"dev_info"`
	MoreInfo string `json:"more_info"`
}

var (
	//000, 系统错误
	ErrInSystem     = &ControllerError{500, 000000, "系统错误", "操作错误", "系统错误，请联系管理员"}

	//100开头，与用户相关
	Err404          = &ControllerError{404, 404, "page not found", "page not found", ""}
	ErrInputData    = &ControllerError{400, 10001, "数据输入错误", "客户端参数错误", ""}
	ErrDatabase     = &ControllerError{500, 10002, "服务器错误", "数据库操作错误", ""}
	ErrDupUser      = &ControllerError{400, 10003, "用户信息已存在", "数据库记录重复", ""}
	ErrNoUser       = &ControllerError{400, 10004, "用户信息不存在", "数据库记录不存在", ""}
	ErrPass         = &ControllerError{400, 10005, "用户信息不存在或密码不正确", "密码不正确", ""}
	ErrNoUserPass   = &ControllerError{400, 10006, "用户信息不存在或密码不正确", "数据库记录不存在或密码不正确", ""}
	ErrNoUserChange = &ControllerError{400, 10007, "用户信息不存在或数据未改变", "数据库记录不存在或数据未改变", ""}
	ErrInvalidUser  = &ControllerError{400, 10008, "用户信息不正确", "Session信息不正确", ""}
	ErrOpenFile     = &ControllerError{500, 10009, "服务器错误", "打开文件出错", ""}
	ErrWriteFile    = &ControllerError{500, 10010, "服务器错误", "写文件出错", ""}
	ErrSystem       = &ControllerError{500, 10011, "服务器错误", "操作系统错误", ""}
	ErrExpired      = &ControllerError{400, 10012, "登录已过期", "验证token过期", ""}
	ErrPermission   = &ControllerError{400, 10013, "没有权限", "没有操作权限", ""}
	Actionsuccess   = &ControllerError{200, 90000, "操作成功", "操作成功", ""}


	//200开头，与文章相关
	MissArticle     = &ControllerError{404, 20001, "查找的文章不存在", "操作失败", "你查找的文章不存在或者已经被删除"}



	//300开头，与coin相关
	ErrGetMessage     = &ControllerError{500, 30001, "获取交易流水错误", "操作失败", "输入地址或其它参数错误"}
	ErrCron           = &ControllerError{500, 30002, "配置监听错误", "操作失败", "输入地址或其它参数错误"}
	ErrCloseCron      = &ControllerError{500, 30003, "关闭监听错误", "操作失败", "没有找到可以关闭的数据"}
	ErrCloseCron2     = &ControllerError{500, 30004, "关闭监听错误", "操作失败", "查找监听数据失败，可能已经关闭了监听"}
	ErrOpenCron       = &ControllerError{500, 30005, "打开监听错误", "操作失败", "没有找到可以监听的数据"}
	MissAddress       = &ControllerError{500, 30006, "打开监听错误", "操作失败", "没有配置地址或者地址错误"}
	ErrInsert         = &ControllerError{500, 30007, "写入数据库错误", "操作失败", "可能邮箱已经存在或者其它字段不符合要求"}
	MissEmail		  = &ControllerError{500, 30008, "获取邮件信息错误", "操作失败", "可能邮箱错误或者其它字段不符合要求"}
	MissCode		  = &ControllerError{500, 30009, "获取code码错误", "操作失败", "可能code码已经过期错或者邮箱错误"}

	ErrPriKey		  = &ControllerError{500, 30010, "私钥形式错误", "操作失败", "请输入正确的私钥"}
	ErrPubKey		  = &ControllerError{500, 30011, "公钥形式错误", "操作失败", "请输入正确的公钥"}
	ErrSellRam        = &ControllerError{500, 30012, "出售ram错误", "操作失败", "请检查账户是否正确或余额是否足够出售"}
	ErrGetKeys        = &ControllerError{500, 30013, "获取密钥错误", "操作失败", "密钥获取错误，请重试"}

	NotEnough          = &ControllerError{500, 30014, "BTC账户余额不够", "操作失败", "请检查比特币账户余额再重试"}
	ErrInputCoinInfo   = &ControllerError{500, 30015, "用户输入错误", "操作失败", "请检查提交数据是否有效"}
	ErrGetNounce       = &ControllerError{500, 30016, "无法获得nounce", "操作失败", "请检查提交数据是否有效"}
	ErrAddress1        = &ControllerError{500, 30017, "地址形式错误", "操作失败", "地址解析错误，请重新检查再写入"}
	ErrAddress2        = &ControllerError{500, 30018, "地址形式错误", "操作失败", "地址解析错误，请重新检查再写入"}
	ErrAccountOrSymbol = &ControllerError{500, 30019, "账户形式或者代币错误", "操作失败", "账户错误或者代币错误，请重新检查再写入"}
	MissParameter      = &ControllerError{500, 30020, "参数错误", "操作失败", "参数错误，请重新检查再写入"}
	ErrDelegateBW      = &ControllerError{500, 30021, "抵押/赎回资源错误", "操作失败", "请重资源是否足够后新检查再写入"}
	ErrBuyRam          = &ControllerError{500, 30022, "购买ram错误", "操作失败", "请检查账户是否正确或余额是否足够购买"}
	ErrAccount         = &ControllerError{500, 30023, "账户形式错误", "操作失败", "请检查账户是否正确"}
	ErrAddPermission   = &ControllerError{500, 30024, "添加权限错误", "操作失败", "请检添加权限参数是否正确或者网络情况、参数等"}
	ErrDelPermission   = &ControllerError{500, 30024, "删除权限错误", "操作失败", "请检删除权限的私钥是否有足够的权限或者网络情况、参数等"}

	//350开头为成功
	SuccessOpenCron    = &ControllerError{200, 35001, "打开监听成功", "操作成功", "正在监听数据"}
	SuccessCloseCron   = &ControllerError{200, 35002, "关闭监听成功", "操作成功", "已经关闭监听"}

	//宝聪
	SucTransferBtc     = &ControllerError{200, 80001, "BTC推送成功", "操作成功", "BTC推送成功"}
)