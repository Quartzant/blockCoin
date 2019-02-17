package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

const (
	Success		= 0
	ErrDatabase = 1
	ErrSystem   = 2
	ErrWrite	= 3
	ErrDupRows  = 4
	ErrNotFound = 5
	ErrInput    = 6
	ErrNotExist = 7
	ErrNotMatch = 8
	ErrValidate = 9
	ErrExist	= 10
	ErrExpired	= 11
)

// Validate Input
type GoogleAuthCode struct {
	GoogleAuthCode string `json:"code"`
}

type SendPhoneCaptchaForm struct {
	Type		int	   `form:"type"         valid:"Required;Range(1, 5)"`
	AreaCode	string `form:"areaCode"     valid:"Required;MinSize(2);MaxSize(3);"`
	Phone 		string `form:"phone"        valid:"Required"`
}

type SendEmailCaptchaForm struct {
	Type		int	   `form:"type"         valid:"Required;Range(1, 5)"`
	Email 		string `form:"email"        valid:"Required;Email"`
}

type SendCaptchaForm struct {
	Type		int	    `form:"type"         valid:"Required;Range(1, 6)"`
	To			string	`form:"to"          valid:"Required;"`
	Lang		int		`form:"lang"          valid:"Required;Range(1, 3)"`
	GeetestChallenge	string	`form:"geetestChallenge"          valid:"Required;"`
	GeetestValidate 	string	`form:"geetestValidate"          valid:"Required;"`
	GeetestSeccode		string	`form:"geetestSeccode"          valid:"Required;"`
	GtServerStatus		int 	`form:"gtServerStatus"          valid:"Required;"`
	GeetestUserId		string	`form:"geetestUserId"          valid:"Required;"`
}

// Users Input
type RegisterPhoneForm struct {
	AreaCode	 string `form:"areaCode"        valid:"Required;MinSize(2);MaxSize(3);"`
	Phone        string `form:"phone"        valid:"Required"`
	Captcha	     string `form:"captcha"      valid:"Required;MinSize(6);MaxSize(6)"`
	Username     string `form:"username"     valid:"Required;MinSize(3);MaxSize(16);"`
	Password     string `form:"password"     valid:"Required;MinSize(32);MaxSize(32);"`
	ConfirmPass  string `form:"confirmPass"      valid:"Required;MinSize(32);MaxSize(32);"`
}

type CheckPhone struct {
	Phone        string `form:"phone"        valid:"Required"`
}

type CheckEmail struct {
	Email        string `form:"email"        valid:"Required;Email"`
}

type CheckUsername struct {
	Username     string `form:"username"     valid:"Required;MinSize(3);MaxSize(16);"`
}

type RegisterEmailForm struct {
	Email        string `form:"email"        valid:"Required;Email"`
	Captcha	     string `form:"captcha"      valid:"Required;MinSize(6);MaxSize(6);"`
	Username     string `form:"username"     valid:"Required;MinSize(3);MaxSize(16);"`
	Password     string `form:"password"     valid:"Required;MinSize(32);MaxSize(32);"`
	ConfirmPass  string `form:"confirmPass"      valid:"Required;MinSize(32);MaxSize(32);"`
}

type LoginPhoneForm struct {
	Phone    	string `form:"phone"    valid:"Required"`
	Password 	string `form:"password" valid:"Required;MinSize(32);MaxSize(32);"`
}

type LoginEmailForm struct {
	Email    string `form:"email"    valid:"Required;Email"`
	Password string `form:"password" valid:"Required;MinSize(32);MaxSize(32);"`
}

type LoginInfo struct {
	Code     int    `json:"code"`
	Token 	 string `json:"token"`
}

type ModifyPassForm struct {
	OldPassword string  `form:"oldPassword" valid:"Required;MinSize(32);MaxSize(32);"`
	NewPassword string  `form:"newPassword" valid:"Required;MinSize(32);MaxSize(32);"`
}

type ResetLogPassForm struct {
	Key			string `form:"key"    valid:"Required;"`
	Password	string	`form:"password" valid:"Required;MinSize(32);MaxSize(32);"`
	Captcha		string `form:"captcha"      valid:"Required;MinSize(6);MaxSize(6);"`
}

type ResetAssPassForm struct {
	Key			string `form:"key"    valid:"Required;"`
	Password	string	`form:"password"  valid:"Required;MinSize(32);MaxSize(32);"`
	Captcha		string  `form:"captcha"   valid:"Required;MinSize(6);MaxSize(6);"`
	Code		string	`form:"code"      valid:"Required;MinSize(6);MaxSize(6);"`
}

type AssPassForm struct {
	Key			string `form:"key"    valid:"Required;"`
	Password	string `form:"password"     valid:"Required;MinSize(32);MaxSize(32);"`
	Captcha		string `form:"captcha"      valid:"Required;MinSize(6);MaxSize(6);"`
}

// Uploads Input
type BindPhoneForm struct {
	AreaCode	string `form:"areaCode"     valid:"Required;MinSize(2);MaxSize(3);"`
	Phone 	string `form:"phone" valid:"Required"`
	Captcha string `form:"captcha" valid:"Required;MinSize(6);MaxSize(6);"`
}

type BindEmailForm struct {
	Email 	string `form:"email" valid:"Required;Email"`
	Captcha string `form:"captcha" valid:"Required;MinSize(6);MaxSize(6);"`
}

type BindGAForm struct {
	Code string `form:"code" valid:"Required;MinSize(6);MaxSize(6);"`
}

type BindRechargeAddressForm struct {
	Token 	string `form:"token" valid:"Required;MinSize(3);MaxSize(4)"`
	Address string `form:"address" valid:"Required;MinSize(34);MaxSize(42)"`
}

type UploadsForm struct {
	Phone string `form:"phone" valid:"Required;Mobile"`
}

type UploadIdCardForm struct {
	IdName       string `form:"id_name"    valid:"Required"`
	IdNumber     string `form:"id_number"  valid:"Required"`
	CardFront    string `form:"card_front"  valid:"Required"`
	CardBack     string `form:"card_back"  valid:"Required"`
}

type UploadPassportForm struct {
	IdName       string `form:"id_name"    valid:"Required"`
	IdNumber     string `form:"id_number"  valid:"Required"`
	Passport     string `form:"passport"   valid:"Required"`
}

type BtcBalanceForm struct {
	Address		string		`form:"address"    valid:"Required;MinSize(34);MaxSize(35)"`
}

type EthBalanceForm struct {
	Address		string		`form:"address"    valid:"Required;MinSize(42);MaxSize(42)"`
}

type ErcBalanceForm struct {
	Address		string		`form:"address"    valid:"Required;MinSize(42);MaxSize(42)"`
	ContractAddress		string		`form:"contractAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
}

type ErcBalanceOutput struct {
	Balance			interface{}
	Decimals		interface{}
}

type ErcDecimalsForm struct {
	ContractAddress		string		`form:"contractAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
}

type DeleteBankForm struct {
	Id		int		`form:"id"    valid:"Required"`
}

type UploadBankForm struct {
	BankNumber  string `form:"bankNumber"  valid:"Required"`
}

type UploadWePayForm struct {
	WePayName    string `form:"wechat_name"    valid:"Required"`
	WePayNumber  string `form:"wechat_number"  valid:"Required"`
	WePayFileUrl string `form:"wechat_file"  valid:"Required"`
}

type UploadAliPayForm struct {
	AliPayName    string `form:"alipay_name"    valid:"Required"`
	AliPayNumber  string `form:"alipay_number"  valid:"Required"`
	AliPayFileUrl string `form:"alipay_file"  valid:"Required"`
}

type ReceiptAddressForm struct {
	Token		string	`form:"token"    valid:"Required;MinSize(3);MaxSize(4)"`
}

type RechargeTokenForm struct {
	RechargeToken 		string	`form:"rechargeToken"    valid:"Required;MinSize(3);MaxSize(4)"`
	RechargeAddress 	string	`form:"rechargeAddress"    valid:"Required;MinSize(34);MaxSize(42)"`
	ReceiptAddress 		string	`form:"receiptAddress"    valid:"Required;MinSize(34);MaxSize(42)"`
	RechargeSum			string	`form:"rechargeSum"    valid:"Required"`
}

type QueryWithdrawAddressForm struct {
	Token 	string	`form:"token"    valid:"Required;MinSize(3);MaxSize(4)"`
}

type WithdrawAddressForm struct {
	Token 			string	`form:"token"    valid:"Required;MinSize(3);MaxSize(4)"`
	Address 		string	`form:"address"  valid:"Required;MinSize(34);MaxSize(42)"`
	Desc		 	string	`form:"desc"     valid:"Required;MinSize(1);MaxSize(20)"`
}

type WithdrawTokenForm struct {
	WithdrawToken 		string	`form:"withdrawToken"    valid:"Required;MinSize(3);MaxSize(4)"`
	WithdrawSum			string	`form:"withdrawSum"    valid:"Required"`
	WithdrawRemark 		string	`form:"withdrawRemark"  valid:"Required;MinSize(1);MaxSize(20)"`
	Captcha	     		string `form:"captcha"      valid:"Required;MinSize(6);MaxSize(6)"`
}

// Output
type LoginOutput struct {
	Jwt 	 	string
	Username 	string
}

type JwtOutput struct {
	Jwt 	 	string
}

type WalletOutput struct {
	Mnemonic 		string
	BtcPrivate		string
	BtcAddress		string
	EthPrivate		string
	EthAddress		string
	EthKeyStore		string
}

type ImportBtcPrivateOutput struct {
	BtcPrivate		string
	BtcAddress		string
}

type ImportEthPrivateOrKeystoreOutput struct {
	EthPrivate		string
	EthAddress		string
	EthKeyStore		string
}

//type UserInfoOutput struct {
//	Username		string
//	AreCode			string
//	Phone			string
//	Email			string
//	GoogleOpen		bool
//	AssPass			bool
//	Kyc				int
//	Name			string
//	BankInfo		[]*UserBank
//	GoogleAuth	    string
//}

type ReceiptAddressOutput struct {
	ethAddress		string
	fccAddress		string
	hkdcAddress		string
	btcAddress		string
}

type UserAssetsOutput struct {
	Fcc		UserTokenInfo
	Btc		UserTokenInfo
	Eth		UserTokenInfo
	Hkdc	UserTokenInfo
}

type UserTokenInfo struct {
	Available	float64
	Freeze		float64
	Price		decimal.Decimal
}

type UserAllWithdrawAddressOutput struct {
	BTC		UserWithdrawAddressOutput
	ETH		UserWithdrawAddressOutput
	FCC		UserWithdrawAddressOutput
	HKDC	UserWithdrawAddressOutput
}

type UserWithdrawAddressOutput struct {
	Desc		string
	Address		string
}

type UserWithdrawFeeOutput struct {
	Fee			float32
}

type UserWithdrawLimitOutput struct {
	Max			float32
	Min			float32
}

type UserWithdrawInfoOutput struct {
	Fee			float32
	Max			float32
	Min			float32
}

type UserRechargeRecordOutput struct {
	State 		 int
	RechAddr	 string
	ReceAddr	 string
	RechSum		 float64
	Comfirm      int
}

type UserWithdrawRecordOutput struct {
	State 		 int
	SerialNum	 string
	Token		 string
	AddrDesc	 string
	WithAddr	 string
	WithSum		 float64
	Fee			 float32
	Remark	  	 string
	CompTime	 int
	TxHash		 string
}

type UserTokenAmount struct {
	Amount 	float64
}

type AccessTokenResult struct {
	Status	int
	Msg		string
	Code	int
	Data	AccessTokenForm
}

type JwtResult struct {
	Status	int
	Msg		string
	Code	int
	Data	string
}

type AccessTokenForm struct {
	AccessToken		string
	ExpiresIn		int
	OpenId			string
}

type Condition struct {
	Limit	int	`form:"limit"    valid:"Required"`
	Offset	int	`form:"offset"    valid:"Required"`
}

// Wallet
type ArticleForm struct {
	Id 		int		`form:"id"  valid:"Required"`
}

type GetInvestmentListForm struct {
	Address		string		`form:"address"  valid:"Required;MinSize(34);MaxSize(42)"`
}

type FinancialManOutput struct {
	Token		 	string
	Term		 	int
	IncomeRate	 	float32
	ReceiptAddress	string
}


// Admin Otc
type KycForm struct {
	Id		int		`form:"id"  valid:"Required;"`
	State	int		`form:"id"  valid:"Required;Range(0, 3)"`
}

type UpdateReceiptAddressForm struct {
	FccAddress string `form:"fccAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
	EthAddress string `form:"ethAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
	BtcAddress string `form:"btcAddress"    valid:"Required;MinSize(34);MaxSize(35)"`
	HkdcAddress string `form:"hkdcAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
}

type UpdateTokenAdFee struct {
	FccAd string `form:"fccAdFee"    valid:"Required;"`
	EthAd string `form:"ethAdFee"    valid:"Required;"`
	BtcAd string `form:"btcAdFee"    valid:"Required;"`
	HkdcAd string `form:"hkdcAdFee"    valid:"Required;"`
}

type UpdateTokenWdFee struct {
	FccWd string `form:"fccWdFee"    valid:"Required;"`
	EthWd string `form:"ethWdFee"    valid:"Required;"`
	BtcWd string `form:"btcWdFee"    valid:"Required;"`
	HkdcWd string `form:"hkdcWdFee"    valid:"Required;"`
}

type UpdateTokenAdLimit struct {
	FccAdMin string `form:"fccAdMin"    valid:"Required;"`
	EthAdMin string `form:"ethAdMin"    valid:"Required;"`
	BtcAdMin string `form:"btcAdMin"    valid:"Required;"`
	HkdcAdMin string `form:"hkdcAdMin"    valid:"Required;"`
	FccAdMax string `form:"fccAdMax"    valid:"Required;"`
	EthAdMax string `form:"ethAdMax"    valid:"Required;"`
	BtcAdMax string `form:"btcAdMax"    valid:"Required;"`
	HkdcAdMax string `form:"hkdcAdMax"    valid:"Required;"`
}

type AllUserAssetsOutput struct {
	Total int64
	Lists []orm.ParamsList
}