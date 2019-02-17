package admin

import (
	"html/template"
)

type Admin struct {
	Id       			int        `orm:"pk;auto"`
	Username 			string	   `orm:"size(50)"`
	Password 			string	   `orm:"size(128)"`
	ReadIco				bool	`orm:"default(true)"`
	GenerateAccount 	bool	`orm:"default(true)"`
	OperateAccount		bool	`orm:"default(true)"`
	CreateArticle		bool	`orm:"default(true)"`
	ReadArticle			bool	`orm:"default(true)"`
	DeleteArticle		bool	`orm:"default(true)"`
	CreateFinancial		bool	`orm:"default(true)"`
	ReadFinancial		bool	`orm:"default(true)"`
	DeleteFinancial		bool	`orm:"default(true)"`
	CreateSys			bool	`orm:"default(true)"`
	ReadSys				bool	`orm:"default(true)"`
	DeleteSys			bool	`orm:"default(true)"`
	ReadInvest			bool	`orm:"default(true)"`
	ReadRewards			bool	`orm:"default(true)"`
	ReadFeedback		bool	`orm:"default(true)"`
	IssueIncome			bool	`orm:"default(true)"`
	Created 	 		int		`orm:"size(10)"`
}

type Ico struct {
	Id       			int         `orm:"pk;auto"`
	Tx					string		`orm:"size(66)"`
	Address				string		`orm:"size(42)"`
	Sum					float64 	`orm:"digits(16);decimals(8)"`
	Created 	 		int			`orm:"size(10)"`
}

type FinancialMan struct {
	Id           	int    		`orm:"pk;auto"`
	Token		 	string		`orm:"size(4)"`
	Term		 	int			`orm:"size(5)"`
	Min				float32		`orm:"digits(8);decimals(4)"`
	IncomeRate	 	float32 	`orm:"digits(8);decimals(4)"`
	ReceiptAddress	string		`orm:"size(42)"`
	Created 	 	int			`orm:"size(10)"`
}

type InvestmentList struct {
	Id           	int    		`orm:"pk;auto"`
	Token		 	string		`orm:"size(4)"`
	Term		 	int			`orm:"size(5)"`
	Tx				string		`orm:"size(66)"`
	Sum				float64 	`orm:"digits(16);decimals(4)"`
	IncomeRate	 	float32 	`orm:"digits(8);decimals(4)"`
	IncomeAddress	string		`orm:"size(42)"`
	ReceiptAddress	string		`orm:"size(42)"`
	StartTime	 	int			`orm:"size(10)"`
	EndTime	 		int			`orm:"size(10)"`
	State			int			`orm:"null;size(1);default(0)"`
	Created 	 	int			`orm:"size(10)"`
}

type Feedback struct {
	Id       int       	`orm:"pk;auto"`
	Content	 string		`orm:"size(500)"`
	Created  int			`orm:"size(10)"`
}

type AppVersion struct {
	Id       int       	`orm:"pk;auto"`
	Version		string		`orm:"default(1.0.2)"`
}

type SystemNotification struct {
	Id      	 int       	`orm:"pk;auto"`
	Title		 string		`orm:"size(100)"`
	Content		 string		`orm:"size(1000)"`
	Created 	 int		`orm:"size(10)"`
}

type AddressIncome struct {
	Id      	 int       	`orm:"pk;auto"`
	Token		 string		`orm:"size(4)"`
	Address		 string		`orm:"size(42)"`
	InvestSum	 float64	`orm:"digits(16);decimals(8)"`
	IncomeSum	 float64	`orm:"digits(16);decimals(8)"`
	State		 int		`orm:"size(4)"`
	TxHash		 string		`orm:"size(66)"`
	Created 	 int		`orm:"size(10)"`
}

type Article struct {
	Id           int    		`orm:"pk;auto"`
	Title		 string			`orm:"size(50)"`
	Desc		 string			`orm:"size(100)"`
	Content		 template.HTML	`orm:"size(9999)"`
	Created 	 int			`orm:"size(10)"`
}