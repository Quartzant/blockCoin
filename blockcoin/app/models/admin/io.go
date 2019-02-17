package admin

import (
	"time"
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
)

// 0 已到期未发放 1 已到期已发放

// Admin Input
type RegisterAdminForm struct {
	Username string `form:"username"    valid:"Required;MinSize(2);MaxSize(30)"`
	Password string `form:"password"    valid:"Required;MinSize(32)"`
	GenerateAccount 	bool	`form:"generateAccount"    valid:"Required"`
	OperateAccount 		bool	`form:"operateAccount"    valid:"Required"`
	CreateArticle		bool	`form:"createArticle"    valid:"Required"`
	ReadArticle			bool	`form:"readArticle"    valid:"Required"`
	DeleteArticle		bool	`form:"deleteArticle"    valid:"Required"`
	CreateFinancial		bool	`form:"createFinancial"    valid:"Required"`
	ReadFinancial		bool	`form:"readFinancial"    valid:"Required"`
	DeleteFinancial		bool	`form:"deleteFinancial"    valid:"Required"`
	CreateSys			bool	`form:"createSys"    valid:"Required"`
	ReadSys				bool	`form:"readSys"    valid:"Required"`
	DeleteSys			bool	`form:"deleteSys"    valid:"Required"`
	ReadRewards			bool	`form:"readRewards"    valid:"Required"`
	ReadFeedback		bool	`form:"readFeedback"    valid:"Required"`
	IssueIncome			bool	`form:"issueIncome"    valid:"Required"`
	ReadIco				bool	`form:"readIco"    valid:"Required"`
}

type LoginAdminForm struct {
	Username string `form:"username"    valid:"Required"`
	Password string `form:"password"    valid:"Required;MinSize(32);MaxSize(32)"`
}

type Condition struct {
	Limit	int	`form:"limit"    valid:"Required"`
	Offset	int	`form:"offset"    valid:"Required"`
}

type FinancialManForm struct {
	Token		string	`form:"token"    valid:"Required"`
	Term		int		`form:"term"     valid:"Required"`
	Min			string	`form:"min"      valid:"Required"`
	IncomeRate	string	`form:"incomeRate"      valid:"Required"`
	ReceiptAddress	string	`form:"receiptAddress"    valid:"Required"`
}

type QueryFinancialManForm struct {
	Id		int	`form:"id"    valid:"Required"`
}

type QueryAdminForm struct {
	Id		int	`form:"id"    valid:"Required"`
}

type QueryArticleForm struct {
	Id		int	`form:"id"    valid:"Required"`
}

type FeedbackForm struct {
	Content		string	`form:"content"    valid:"Required;MinSize(5);MaxSize(200)"`
}

type WalletForm struct {
	Password 	string	`form:"password"    valid:"Required;MinSize(6);MaxSize(20)"`
}

type MnemonicForm struct {
	Mnemonic 	string	`form:"mnemonic"    valid:"Required"`
	Password 	string	`form:"password"    valid:"Required;MinSize(6);MaxSize(20);"`
}

type KeyStoreForm struct {
	KeyStore 	string	`form:"keyStore"    valid:"Required"`
	Password 	string	`form:"password"    valid:"Required"`
}

type TransferEthForm struct {
	FromAddress		string	`form:"fromAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
	ToAddress		string	`form:"toAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
	Private 		string	`form:"private"    valid:"Required;MinSize(64);MaxSize(64)"`
	Amount			string	`form:"amount"    valid:"Required"`
	GasPrice		int64	`form:"gasPrice"    valid:"Required"`
	GasLimit		uint64	`form:"gasLimit"    valid:"Required"`
}

type TransferErc20Form struct {
	FromAddress		string	`form:"formAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
	ToAddress		string	`form:"toAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
	ContractAddress	string	`form:"contractAddress"    valid:"Required;MinSize(42);MaxSize(42)"`
	Private 		string	`form:"private"    valid:"Required;MinSize(64);MaxSize(64)"`
	Amount			string	`form:"amount"    valid:"Required"`
	GasPrice		int64	`form:"gasPrice"    valid:"Required"`
	GasLimit		uint64	`form:"gasLimit"    valid:"Required"`
}

type TransferBtcForm struct {
	FromAddress		string	`form:"formAddress" valid:"Required;MinSize(34);MaxSize(35)"`
	ToAddress		string	`form:"toAddress"   valid:"Required;MinSize(34);MaxSize(35)"`
	Private 		string	`form:"private"     valid:"Required;MinSize(52);MaxSize(52)"`
	Amount			string	`form:"amount"      valid:"Required"`
	Fee				int64	`form:"fee"         valid:"Required"`
}

type EthPrivateForm struct {
	Private 	string	`form:"private"    valid:"Required;MinSize(64);MaxSize(64)"`
	Password 	string	`form:"password"    valid:"Required;MinSize(6);MaxSize(20);"`
}

type BtcPrivateForm struct {
	Private 	string	`form:"private"    valid:"Required;MinSize(52);MaxSize(52)"`
}

type SysNotificationForm struct {
	Title		string	`form:"title"    valid:"Required"`
	Content		string	`form:"content"    valid:"Required"`
}

type ArticleForm struct {
	Title		string	`form:"title"      valid:"Required"`
	Desc 		string	`form:"desc"       valid:"Required"`
	Content		string	`form:"content"    valid:"Required"`
}

type QuerySysNotification struct {
	Id		int	`form:"id"    valid:"Required"`
}

type QueryAddressInvestForm struct {
	Id		int `form:"id"    valid:"Required"`
	Limit	int	`form:"limit"    valid:"Required"`
	Offset	int	`form:"offset"    valid:"Required"`
}

type QueryAddressRewardsForm struct {
	Id		int `form:"id"    valid:"Required"`
	Limit	int	`form:"limit"    valid:"Required"`
	Offset	int	`form:"offset"    valid:"Required"`
}

type UpdateAddressRewardsForm struct {
	Id		int 	`form:"id"    valid:"Required"`
}

// Admin Output
type AdminOutput struct {
	Total int64
	Lists []*Admin
}

type FeedbackOutput struct {
	Total int64
	Lists []*Feedback
}

type FinancialManOutput struct {
	Total int64
	Lists []*FinancialMan
}

type SysNotificationOutput struct {
	Total int64
	Lists []*SystemNotification
}

type ArticleOutput struct {
	Total int64
	Lists []*Article
}

type FinancialTokenAndAddressOutput struct {
	Token		string
	Address		string
}

type InvestmentOutput []struct{
	Id				int
	Token		 	string
	Term		 	int
	Sum				int
	Tx				string
	IncomeRate	 	float32
	IncomeAddress	string
	ReceiptAddress	string
	StratTime	 	time.Time
	EndTime	 		time.Time
}

type AllInvestmentOutput struct {
	Total int64
	Lists []*InvestmentList
}

type AllIcoOutput struct {
	Total int64
	Lists []*Ico
}

type InvestmentRewardsOutput []struct{
	Id				int
	Token		 	string
	Term		 	int
	Sum				int
	Tx				string
	IncomeRate	 	float32
	IncomeAddress	string
	ReceiptAddress	string
	StratTime	 	time.Time
	EndTime	 		time.Time
}

type AllInvestmentRewardsOutput struct {
	Total int64
	Lists []*InvestmentList
}

type AllArticleOutput []struct {
	Id           int
	Title		 string
	Desc		 string
	Created 	 int64
}

type AssetsOutput struct {
	Id			int
	Fcc			string
	Eth			string
	Btc			string
	Hkdc		string
	FccFreeze		string
	EthFreeze		string
	BtcFreeze		string
	HkdcFreeze		string
}