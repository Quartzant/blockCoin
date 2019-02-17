package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"regexp"
	"strings"
)

// Predefined const error strings.
const (
	ErrInputData    = "数据输入错误"
	ErrDatabase     = "数据库操作错误"
	ErrDupUser      = "用户信息已存在"
	ErrNoUser       = "用户信息不存在"
	ErrPass         = "密码不正确"
	ErrNoUserPass   = "用户信息不存在或密码不正确"
	ErrNoUserChange = "用户信息不存在或数据未改变"
	ErrInvalidUser  = "用户信息不正确"
	ErrOpenFile     = "打开文件出错"
	ErrWriteFile    = "写文件出错"
	ErrSystem       = "操作系统错误"
)

var captchaType = map[int]string{
	1: "register",
	2: "bind",
	3: "set",
	4: "resetLog",
	5: "resetAss",
	6: "withdraw",
}

var emailCaptchaSubject= map[int]string{
	1: "【Coin4OTC】验证码-",
	2: "【Coin4OTC】驗證碼-",
	3: "【Coin4OTC】Verification code-",
}

var emailCaptchaContent = map[int]string{
	1: "您的邮箱验证码是：",
	2: "您的郵箱驗證碼是：",
	3: "Your email verification code is: ",
}

// ControllerError is controller error info structer.
type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Msg  	 string `json:"msg"`
	DevInfo  string `json:"dev_info"`
}

type ControllerSuccess struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Msg  	 string `json:"msg"`
	Data     interface{} `json:"data"`
}

// Predefined controller error values.
var (
	errNotFound     = &ControllerError{404, 404, "page not found", "page not found"}
	errInputData    = &ControllerError{200, 10001, "数据输入错误", "客户端参数错误"}
	errDiffPassInput= &ControllerError{200, 10001, "数据输入错误", "两次密码不一致"}
	errSendPhone    = &ControllerError{200, 10001, "验证码发送失败", "手机验证码发送失败"}
	errSendEmail    = &ControllerError{200, 10001, "验证码发送失败", "邮箱验证码发送失败"}
	errSendOften    = &ControllerError{200, 10001, "验证码发送失败", "不能频繁发送验证码"}
	errPhoneExpired = &ControllerError{200, 10017, "验证码失效", "手机验证码已失效"}
	errPhoneCaptcha = &ControllerError{200, 10018, "验证码不正确", "手机验证码不正确"}
	errEmailExpired = &ControllerError{200, 10018, "验证码失效", "邮箱验证码已失效"}
	errEmailCaptcha = &ControllerError{200, 10018, "验证码不正确", "邮箱验证码不正确"}
	errDatabase     = &ControllerError{500, 10002, "服务器错误", "数据库操作错误"}
	errDupUser      = &ControllerError{200, 10003, "用户信息已存在", "数据库记录重复"}
	errNoUser       = &ControllerError{200, 10004, "用户信息不存在", "数据库记录不存在"}
	errExistPhone   = &ControllerError{200, 10004, "该手机已被注册", "该手机已被注册"}
	errExistEmail   = &ControllerError{200, 10004, "该邮箱已被注册", "该邮箱已被注册"}
	errExistUsername= &ControllerError{200, 10004, "该用户名已被注册", "该用户名已被注册"}
	errPass         = &ControllerError{200, 10005, "用户信息不存在或密码不正确", "密码不正确"}
	errNoUserPass   = &ControllerError{200, 10006, "用户信息不存在或密码不正确", "数据库记录不存在或密码不正确"}
	errOpenFile     = &ControllerError{500, 10009, "服务器错误", "打开文件出错"}
	errWriteFile    = &ControllerError{500, 10010, "服务器错误", "写文件出错"}
	errNotMatch     = &ControllerError{500, 10019, "图片类型和大小不符合", "图片类型和大小不符合"}
	errSystem       = &ControllerError{500, 10011, "服务器错误", "操作系统错误"}
	errLogExpired   = &ControllerError{200, 10012, "登录已过期", "验证token过期"}
	errPermission   = &ControllerError{200, 10013, "没有权限", "没有操作权限"}
	errGoogleAuth   = &ControllerError{200, 10014, "验证失败", "谷歌两步验证失败"}
	errExistBank	= &ControllerError{200, 10050, "绑定失败", "已经绑定过该银行卡"}
	errBindBank		= &ControllerError{200, 10014, "绑定失败", "绑定银行卡数量超过限制"}
	errBindPhone    = &ControllerError{200, 10014, "绑定失败", "你当前为手机登录,不能绑定手机"}
	errBindEmail   	= &ControllerError{200, 10014, "绑定失败", "你当前为邮箱登录,不能绑定邮箱"}
	errPassKyc   	= &ControllerError{200, 10014, "未通过KYC,不能绑定银行卡", "未通过KYC,不能绑定银行卡"}
	errGeetest  	= &ControllerError{400, 10015, "验证失败", "极验验证失败"}
	errNotExistRecord	= &ControllerError{400, 10045, "暂无记录", "找不到"}
	errNotEnoughToken	= &ControllerError{400, 10030, "该Token余额不足", "该Token余额不足"}
	errOutWithLimit		= &ControllerError{400, 10031, "超出提币限额", "超出提币限额"}
	errExistRechAddr	= &ControllerError{400, 10032, "该地址已被绑定,请选择其他地址", "该地址已被绑定"}
	errExistWithAddr	= &ControllerError{400, 10033, "用户该Token提币地址已存在，只能设置一个提币地址", "用户该Token提币地址已存在"}
	errNotExistWithAddr	= &ControllerError{400, 10034, "用户该Token提币地址不存在，请新建", "用户该Token提币地址已存在"}
	errNotMatchReceAddr = &ControllerError{400, 10035, "与官方收币地址不匹配", "与官方收币地址不匹配"}

	errGetQuotes 		= &ControllerError{400, 10018, "获取行情失败", "爬虫炸了"}
	errGetLive		 	= &ControllerError{400, 10018, "获取快讯失败", "爬虫炸了"}
	errGetSysnotify		= &ControllerError{400, 10018, "获取系统通知失败", ""}
	errGetFinancialMan 	= &ControllerError{400, 10018, "获取理财产品列表失败", "数据库查询错误"}

	errGetAddressTx 	= &ControllerError{400, 10018, "获取该地址相关TxHash失败", "爬虫炸了"}

	errExistFinancial	= &ControllerError{400, 10018, "当前已存在该类型的理财产品", "杰哥牛逼"}
	errExistSysNotify	= &ControllerError{400, 10018, "当前已存在该系统通知", "杰哥对不起"}
)
// Predefined controller success values.
var (
	sucRegister      	= &ControllerSuccess{200, 10000, "注册成功", ""}
	sucLogin      	    = &ControllerSuccess{200, 10000, "登录成功", ""}
	sucNotExist			= &ControllerSuccess{200, 10000, "用户名或手机或邮箱有效", ""}
	sucSendPhone      	= &ControllerSuccess{200, 10000, "手机验证码发送成功", ""}
	sucSendEmail      	= &ControllerSuccess{200, 10000, "邮箱验证码发送成功", ""}
	sucSendCaptcha      = &ControllerSuccess{200, 10000, "验证码发送成功", ""}
	sucGoogleAuth      	= &ControllerSuccess{200, 10000, "谷歌两步验证成功", ""}
	sucGeetest      	= &ControllerSuccess{200, 10000, "极验验证成功", ""}
	sucSetAssPass		= &ControllerSuccess{200, 10000, "设置资金密码成功", ""}
	sucModifyLogPass	= &ControllerSuccess{200, 10000, "修改登录密码成功", ""}
	sucModifyAssPass	= &ControllerSuccess{200, 10000, "修改资金密码成功", ""}
	sucResetLogPass		= &ControllerSuccess{200, 10000, "重置登录密码成功", ""}
	sucResetAssPass		= &ControllerSuccess{200, 10000, "重置资金密码成功", ""}
	sucBindPhone      	= &ControllerSuccess{200, 10000, "手机绑定成功", ""}
	sucBindEmail      	= &ControllerSuccess{200, 10000, "邮箱绑定成功", ""}
	sucBindGoogle      	= &ControllerSuccess{200, 10000, "谷歌两步验证绑定成功", ""}
	sucBindRechAddr     = &ControllerSuccess{200, 10000, "充币地址绑定成功", ""}
	sucUploadKyc	    = &ControllerSuccess{200, 10000, "上传实名信息成功", ""}
	sucAddBank	   	    = &ControllerSuccess{200, 10000, "添加银行信息成功", ""}
	sucDeleteBank	   	= &ControllerSuccess{200, 10000, "添加银行信息成功", ""}
	sucSubWitAdd	    = &ControllerSuccess{200, 10000, "添加用户提币地址成功", ""}
	sucSubRechargeOrder		= &ControllerSuccess{200, 10000, "提交用户充币工单成功", ""}
	sucSubWithdrawOrder		= &ControllerSuccess{200, 10000, "提交用户提币工单成功", ""}
	sucGetUserInfo			= &ControllerSuccess{200, 10000, "获取用户信息成功", ""}
	sucGetRechargeRecord	= &ControllerSuccess{200, 10000, "获取用户充币记录成功", ""}
	sucGetWithdrawRecord	= &ControllerSuccess{200, 10000, "获取用户提币记录成功", ""}
	sucGetReceiptAddress	= &ControllerSuccess{200, 10000, "获取收币地址成功", ""}
	sucGetRechargeAddress	= &ControllerSuccess{200, 10000, "获取充币地址成功", ""}
	sucGetWithdrawAddress	= &ControllerSuccess{200, 10000, "获取用户提币地址成功", ""}
	sucGetWithdrawInfo		= &ControllerSuccess{200, 10000, "获取提币限额和手续费成功", ""}
	sucGetUserTxList    	= &ControllerSuccess{200, 10000, "获取用户交易记录成功", ""}
	sucGetUserAdList    	= &ControllerSuccess{200, 10000, "获取用户广告记录成功", ""}

	sucCancelTx			= &ControllerSuccess{200, 10000, "取消交易成功", ""}
	sucPayTx			= &ControllerSuccess{200, 10000, "确认支付成功", ""}
	sucPublishSellAd	= &ControllerSuccess{200, 10000, "发布出售代币广告成功", ""}
	sucPublishBuyAd		= &ControllerSuccess{200, 10000, "发布购买代币广告成功", ""}

	sucPassUserKyc	    = &ControllerSuccess{200, 10000, "该用户KYC信息已通过", ""}
	sucRechargeArrival	= &ControllerSuccess{200, 10000, "该用户充币已到账", ""}
	sucWithdrawArrival  = &ControllerSuccess{200, 10000, "该用户提币已到账", ""}
	sucGetInfo			= &ControllerSuccess{200, 10000, "获取信息成功", ""}
	sucUpdateInfo		= &ControllerSuccess{200, 10000, "更新信息成功", ""}
	sucGetQuotes		= &ControllerSuccess{200, 10000, "获取行情信息成功", ""}
	sucGetLive			= &ControllerSuccess{200, 10000, "获取快讯信息成功", ""}
	sucGetFinancialMan  = &ControllerSuccess{200, 10000, "获取理财产品列表成功", ""}
	sucGetAppVersion    = &ControllerSuccess{200, 10000, "获取App版本号成功", ""}
	sucIssueFeedback	= &ControllerSuccess{200, 10000, "提交反馈列表成功", ""}
	sucGetIco		    = &ControllerSuccess{200, 10000, "获取Ico列表成功", ""}
	sucRegisterAdmin	= &ControllerSuccess{200, 10000, "创建管理员账户成功", ""}
	sucLoginAdmin		= &ControllerSuccess{200, 10000, "后台管理登录成功", ""}
	sucModifyAdPass		= &ControllerSuccess{200, 10000, "修改管理员密码成功", ""}
	sucGetAdmin			= &ControllerSuccess{200, 10000, "分页获取管理员账号成功", ""}
	sucDeleteAdmin		= &ControllerSuccess{200, 10000, "删除该管理员账户成功", ""}
	sucGetFeedback		= &ControllerSuccess{200, 10000, "获取反馈列表成功", ""}
	sucCreateSysNotify	= &ControllerSuccess{200, 10000, "新建系统通知成功", ""}
	sucGetSysNotify		= &ControllerSuccess{200, 10000, "获取系统通知列表成功", ""}
	sucDeleteSysNotify	= &ControllerSuccess{200, 10000, "删除该系统通知成功", ""}
	sucCreateFinancial	= &ControllerSuccess{200, 10000, "新建理财产品成功", ""}
	sucGetFinancial		= &ControllerSuccess{200, 10000, "获取理财产品列表成功", ""}
	sucDeleteFinancial	= &ControllerSuccess{200, 10000, "删除该理财产品成功", ""}
	sucUploadArticlePic	= &ControllerSuccess{200, 10000, "上传文章图片成功", ""}
	sucCreateArticle	= &ControllerSuccess{200, 10000, "新建文章成功", ""}
	sucGetArticle		= &ControllerSuccess{200, 10000, "获取文章列表成功", ""}
	sucDeleteArticle	= &ControllerSuccess{200, 10000, "删除该文章成功", ""}
	sucGetInvest	    = &ControllerSuccess{200, 10000, "获取该地址相关投资列表成功", ""}
	sucGetRewards		= &ControllerSuccess{200, 10000, "获取投资奖励处理列表成功", ""}
	sucUpdateRewards	= &ControllerSuccess{200, 10000, "更新该投资奖励处理成功", ""}
	sucCreataWallet		= &ControllerSuccess{200, 10000, "创建钱包成功", ""}
	sucImportMnemonic	= &ControllerSuccess{200, 10000, "导入助记词成功", ""}
	sucImportBtcPrivate	= &ControllerSuccess{200, 10000, "导入BTC私钥成功", ""}
	sucImportEthPrivate	= &ControllerSuccess{200, 10000, "导入ETh私钥成功", ""}
	sucImportEthKeysotre= &ControllerSuccess{200, 10000, "导入KeySotre成功", ""}
	sucTransferEth		= &ControllerSuccess{200, 10000, "打包发送交易成功", ""}
	sucTransferBtc		= &ControllerSuccess{200, 10000, "打包发送交易成功", ""}
)

// BaseController definiton.
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok"}
	c.ServeJSON()
}

// RetError return error information in JSON.
func (base *BaseController) RetError(e *ControllerError) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.DevInfo = ""
	}

	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()

	base.StopRun()
}

// RetSuccess return error information in JSON.
func (base *BaseController) RetSuccess(e *ControllerSuccess) {
	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()

	base.StopRun()
}

var sqlOp = map[string]string{
	"eq": "=",
	"ne": "<>",
	"gt": ">",
	"ge": ">=",
	"lt": "<",
	"le": "<=",
}

// ParseQueryParm parse query parameters.
//   query=col1:op1:val1,col2:op2:val2,...
//   op: one of eq, ne, gt, ge, lt, le
func (base *BaseController) ParseQueryParm() (v map[string]string, o map[string]string, err error) {
	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
	queryVal := make(map[string]string)
	queryOp := make(map[string]string)

	query := base.GetString("query")
	if query == "" {
		return queryVal, queryOp, nil
	}

	for _, cond := range strings.Split(query, ",") {
		kov := strings.Split(cond, ":")
		if len(kov) != 3 {
			return queryVal, queryOp, errors.New("Query format != k:o:v")
		}

		var key string
		var value string
		var operator string
		if !nameRule.MatchString(kov[0]) {
			return queryVal, queryOp, errors.New("Query key format is wrong")
		}
		key = kov[0]
		if op, ok := sqlOp[kov[1]]; ok {
			operator = op
		} else {
			return queryVal, queryOp, errors.New("Query operator is wrong")
		}
		value = strings.Replace(kov[2], "'", "\\'", -1)

		queryVal[key] = value
		queryOp[key] = operator
	}

	return queryVal, queryOp, nil
}

// ParseOrderParm parse order parameters.
//   order=col1:asc|desc,col2:asc|esc,...
func (base *BaseController) ParseOrderParm() (o map[string]string, err error) {
	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
	order := make(map[string]string)

	v := base.GetString("order")
	if v == "" {
		return order, nil
	}

	for _, cond := range strings.Split(v, ",") {
		kv := strings.Split(cond, ":")
		if len(kv) != 2 {
			return order, errors.New("Order format != k:v")
		}
		if !nameRule.MatchString(kv[0]) {
			return order, errors.New("Order key format is wrong")
		}
		if kv[1] != "asc" && kv[1] != "desc" {
			return order, errors.New("Order val isn't asc/desc")
		}

		order[kv[0]] = kv[1]
	}

	return order, nil
}

// ParseLimitParm parse limit parameter.
//   limit=n
func (base *BaseController) ParseLimitParm() (l int64, err error) {
	if v, err := base.GetInt64("limit"); err != nil {
		return 10, err
	} else if v > 0 {
		return v, nil
	} else {
		return 10, nil
	}
}

// ParseOffsetParm parse offset parameter.
//   offset=n
func (base *BaseController) ParseOffsetParm() (o int64, err error) {
	if v, err := base.GetInt64("offset"); err != nil {
		return 0, err
	} else if v > 0 {
		return v, nil
	} else {
		return 0, nil
	}
}

// VerifyForm use validation to verify input parameters.
func (base *BaseController) VerifyForm(obj interface{}) (err error) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		return err
	}
	if !ok {
		str := ""
		for _, err := range valid.Errors {
			str += err.Key + ":" + err.Message + ";"
		}
		return errors.New(str)
	}

	return nil
}

// generate json web token
//func (base *BaseController) GenToken(userKey string, level string) string {
//	key := []byte("fcc@163.com")
//	claims := &jwt.StandardClaims{
//		NotBefore: int64(time.Now().Unix()),
//		ExpiresAt: int64(time.Now().Unix() + 3600000),
//		Audience:  userKey,
//		Issuer:    "fcc",
//		Subject:   level,
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	ss, err := token.SignedString(key)
//	if err != nil {
//		logs.Error(err)
//		return ""
//	}
//	return ss
//}
//
//// check token
//func (base *BaseController) CheckToken(token string) bool {
//	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
//		return []byte("fcc@163.com") , nil
//	})
//	if err != nil {
//		fmt.Println("parase with claims failed.", err)
//		return false
//	}
//	return true
//}
//
//// ParseToken parse JWT token in http header.
//func (base *BaseController) ParseToken() (t *jwt.Token, e *ControllerError) {
//	authString := base.Ctx.Input.Header("Authorization")
//
//	kv := strings.Split(authString, " ")
//	if len(kv) != 2 || kv[0] != "Bearer" {
//		beego.Error("AuthString invalid:", authString)
//		return nil, errInputData
//	}
//	tokenString := kv[1]
//
//	// Parse token
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte("fcc@163.com"), nil
//	})
//	if err != nil {
//		beego.Error("Parse token:", err)
//		if ve, ok := err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				// That's not even a token
//				return nil, errInputData
//			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
//				// Token is either expired or not active yet
//				return nil, errLogExpired
//
//			} else {
//				// Couldn't handle this token
//				return nil, errInputData
//			}
//		} else {
//			// Couldn't handle this token
//			return nil, errInputData
//		}
//	}
//	if !token.Valid {
//		beego.Error("Token invalid:", tokenString)
//		return nil, errInputData
//	}
//
//	return token, nil
//}

//jwt认证
//func (base *BaseController) ValidateUserJwt() (key string, code int){
//	token, err := base.ParseToken()
//	if err != nil {
//		return "", models.ErrValidate
//	}
//	if !base.CheckToken(token.Raw) {
//		return "", models.ErrValidate
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return "", models.ErrValidate
//	}
//	if claims["sub"].(string) != "user" {
//		return "", models.ErrValidate
//	}
//	userKey := claims["aud"].(string)
//	return userKey, models.Success
//}

//func (base *BaseController) ValidateAdminJwt() (key string, code int){
//	token, err := base.ParseToken()
//	if err != nil {
//		return "", models.ErrValidate
//	}
//	if !base.CheckToken(token.Raw) {
//		return "", models.ErrValidate
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return "", models.ErrValidate
//	}
//	if claims["sub"].(string) != "admin" {
//		return "", models.ErrValidate
//	}
//	beego.Debug(claims["aud"])
//	userKey := claims["aud"].(string)
//	return userKey, models.Success
//}

/*func (base *BaseController) ValidateSuperJwt() (key string, code int){
	token, e := base.ParseToken()
	if e == nil {
		return "", models.ErrValidate
	}
	if !base.CheckToken(token.Raw) {
		return "", models.ErrValidate
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", models.ErrValidate
	}
	if claims["sub"].(string) == "super" {
		return "", models.ErrValidate
	}
	userKey := claims["aud"].(string)
	return userKey, models.Success
}*/


//上传文件
//func (base *BaseController) SaveFileToPath(formFile string, path string) (code int, p string) {
//	f, h, err := base.GetFile(formFile)
//	beego.Debug(h.Size)
//	beego.Debug(filepath.Ext(h.Filename))
//	if h.Size > 2076672 {
//		return models.ErrNotMatch, ""
//	}
//	if err != nil {
//		return models.ErrInput, ""
//	}
//	if strings.ToLower(filepath.Ext(h.Filename)) != ".png" || strings.ToLower(filepath.Ext(h.Filename)) != ".jpg" || strings.ToLower(filepath.Ext(h.Filename)) != ".jpeg" {
//		return models.ErrNotMatch, ""
//	}
//	defer f.Close()
//
//	hash := md5.New()
//	if _, err := io.Copy(hash, f); err != nil {
//		beego.Debug(err)
//		return models.ErrSystem, ""
//	}
//
//	hex := fmt.Sprintf("%x", hash.Sum(nil))
//	dst, err := os.Create(path + hex + filepath.Ext(h.Filename))
//	if err != nil {
//		beego.Debug(err)
//		return models.ErrSystem, ""
//	}
//	defer dst.Close()
//
//	f.Seek(0, 0)
//	if _, err := io.Copy(dst, f); err != nil {
//		return models.ErrWrite, ""
//	}
//	return models.Success, path + hex + filepath.Ext(h.Filename)
//}