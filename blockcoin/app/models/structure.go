package models

import (
	"time"
)


const (
	BTCTerminal   =  time.Second * 5
	EOSTerminal   =  time.Second * 5
	XRPTerminal   =  time.Second * 5
	ETHTerminal   =  time.Second * 5
	USDTTerminal  =  time.Second * 5
)

type CronModel struct {
	Id    				 int			`orm:"column(id)"`
	SwitchBTC            int            `orm:"column(switch_btc)"`
	SwitchEOS            int            `orm:"column(switch_eos)"`
	SwitchUSDT           int            `orm:"column(switch_usdt)"`
	SwitchETH            int            `orm:"column(switch_eth)"`
	SwitchXRP            int            `orm:"column(switch_xrp)"`
	Phone 				 string			`orm:"column(phone)"`
	Email 		   	  	 string			`orm:"column(email);unique"`
	AddressBTC			 string			`orm:"column(address_btc)"`
	AccountEOS			 string			`orm:"column(account_eos)"`
	AddressXRP			 string			`orm:"column(address_xrp)"`
	AddressETH			 string			`orm:"column(address_eth)"`
	BtcIdentification    int			`orm:"column(btc_identification)"`
	EthIdentification	 string			`orm:"column(eth_identification)"`
	UsdtIdentification	 string			`orm:"column(usdt_identification)"`
	EosIdentification	 int			`orm:"column(eos_identification)"`
	XrpIdentification	 string			`orm:"column(xrp_identification)"`
	CreatedAt			 time.Time		`orm:"auto_now_add;type(datetime)"`
	UpdatedAt            time.Time      `orm:"auto_now;type(datetime)"`
}

func (u *CronModel) TableName() string {
	return "cron_job"
}

//历史流水相关
type BtcData struct {
	TotalCount    int		`json:"total_count"`
	Page          int		`json:"page"`
	PageSize      int		`json:"pagesize"`
	List         []BtcList	`json:"list"`
}

type BtcList struct {
	CreatedAt    int64			`json:"created_at"`
	Fee    		  float64		`json:"fee"`
	Hash          string		`json:"hash"`
	BalanceDiff  float64		`json:"balance_diff"`
}

type BtcCoin struct {
	Data    BtcData				`json:"data"`
	ErrMsg  string				`json:"err_msg"`
	ErrNo   int					`json:"err_no"`
}

type EthCoin struct {
	Timestamp   int64			`json:"timestamp"`
	From		string			`json:"from"`
	To			string			`json:"to"`
	Hash		string			`json:"hash"`
	Value		float64			`json:"value"`
}

type EosTraceList struct {
	TrxId		string			`json:"trx_id"`
	Timestamp 	string			`json:"timestamp"`
	Receiver	string			`json:"receiver"`
	Sender		string			`json:"sender"`
	Quantity	string			`json:"quantity"`
	Symbol		string			`json:"symbol"`
}

type EosData struct {
	TraceCount   int				`json:"trace_count"`
	List 	     []EosTraceList 	`json:"trace_list"`
}

type EosCoin struct {
	Errno    int				`json:"errno"`
	Errmsg   string				`json:"errmsg"`
	Data     EosData			`json:"data"`
}

type UsdtTransactions struct {
	Amount  	  	    string		`json:"amount"`
	Fee					string		`json:"fee"`
	Txid				string		`json:"txid"`
	Sendingaddress		string		`json:"sendingaddress"`
	Referenceaddress	string		`json:"referenceaddress"`
}

type UsdtCoin struct {
	Address  		string	      	`json:"address"`
	CurrentPage		int			 `json:"current_page"`
	Pages			int				 `json:"pages"`
	Transactions	[]UsdtTransactions     `json:"transactions"`
}

type Tx struct {
	TransactionType  string
	Amount			 string
	Fee			   	 string
	Account			 string
	Destination		 string
}

type XrpTransactions struct {
	Hash  			string
	ledgerIndex 	int
	Date            time.Time
	Tx              Tx
}

type XrpCoin struct {
	Result    string
	Count     int
	Transactions  []XrpTransactions
}
/////


///Email
type EmailData struct {
	ReceiverAddress   string
	NickName          string
	Subject			  string
	Body			  string
}

type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	DevInfo  string `json:"dev_info"`
	MoreInfo string `json:"more_info"`
}
////



type EOSQuantity struct {
	RamQuantity   string    `json:"ram_quantity"`
	CpuQuantity   string    `json:"cpu_quantity"`
	NetQuantity   string    `json:"net_quantity"`
}

type CreateEOSAccount struct {
	CreatorAccount    	string    	 	 `json:"creator_account"`
	CreatorPriKey     	string   		 `json:"creator_priKey"`
	NewAccount        	string   		 `json:"new_account"`
	NewAccountPriKey    string   		 `json:"new_account_prikey"`
	NewAccountPubKey    string   		 `json:"new_account_pubkey"`
	Quantity		 	EOSQuantity   	 `json:"quantity"`
}

type TransferEOS struct {
	FromAccount     string   	 `json:"from_account"`
	FromPriKey      string   	 `json:"from_priKey"`
	ToAccount       string   	 `json:"to_account"`
	Quantity        string   	 `json:"quantity"`
	Memo    		string   	 `json:"memo"`
	Code            string   	 `json:"code"`
}


type BTCUnspentOutputs struct {
	UnspentOutputs     []BTCUnspentOutputsList    `json:"unspent_outputs"`
}

type BTCUnspentOutputsList struct {
	TxHash     				string        `json:"tx_hash" `
	TxHashBigEndian  	    string        `json:"tx_hash_big_endian"`
	TxIndex   			    int64         `json:"tx_index"`
	TxOutputN   		    int           `json:"tx_output_n"`
	Script    			    string        `json:"script"`
	Value     		  	    int64         `json:"value"`
	ValueHex     			string        `json:"value_hex"`
	Confirmations     		int64        `json:"confirmations"`
}

type BTCTransferReceive struct {
	FromAddress 	 string 	    `json:"from_address"`
	Private  		 string  		`json:"private"`
	ToAddress 	 	 string  		`json:"to_address"`
	Amount  		 int64  		`json:"amount"`
	Fee  			 int64  		`json:"fee"`
}

type USDTTransferReceive struct {
	FromAddress 	 string 	    `json:"from_address"`
	Private  		 string  		`json:"private"`
	ToAddress 	 	 string  		`json:"to_address"`
	Amount  		 float64  		`json:"amount"`
	Fee  			 int64  		`json:"fee"`
}


type AddressIdentifications struct {
	BTCIdentification    int		`orm:"column(btc_identification)"`
	ETHIdentification    string		`orm:"column(eth_identification)"`
	USDTIdentification   string		`orm:"column(usdt_identification)"`
	EOSIdentification    int		`orm:"column(eos_identification)"`
	XRPIdentification    string		`orm:"column(xrp_identification)"`

}

type RegisterAddress struct {
	AddressBTC	string		   `json:"address_btc"`
	AccountEOS	string		   `json:"account_eos"`
	AddressXRP	string		   `json:"address_xrp"`
	AddressETH	string		   `json:"address_eth"`
}

type Form struct {
	Phone    	string		     `json:"phone"`
	Email    	string	         `json:"email"`
	Address     RegisterAddress  `json:"address"`
}

type RegisterData struct {
	Form	   Form		    `json:"form"`
}


type EOSBuyRam struct {
	FromAccount   string  	 `json:"from_account"  valid:"Required"`
	PrivateKey    string     `json:"private_key"   valid:"Required"`
	ToAccount     string  	 `json:"to_account"    valid:"Required"`
	NumBytes      int  	 	 `json:"num_bytes"     valid:"Required"`
}

type EOSSellRam struct {
	Account 	  string  	 `json:"account"       valid:"Required"`
	PrivateKey    string     `json:"private_key"   valid:"Required"`
	NumBytes      int32  	 `json:"num_bytes"     valid:"Required"`
}


type EOSDelegateBW struct {
	FromAccount   		string     `json:"from_account"             valid:"Required"`
	PrivateKey   	    string     `json:"private_key"              valid:"Required"`
	ToAccount    	    string     `json:"to_account"               valid:"Required"`
	SkateNetQuantity    string	   `json:"stake_net_quantity"       valid:"Required"`
	SkateCpuQuantity    string	   `json:"stake_cpu_quantity"       valid:"Required"`
	Transfer            bool       `json:"transfer"                 valid:"Required"`
}

type EOSUnDelegateBW struct {
	FromAccount   		string     `json:"from_account"             valid:"Required"`
	PrivateKey   	    string     `json:"private_key"              valid:"Required"`
	ToAccount    	    string     `json:"to_account"               valid:"Required"`
	SkateNetQuantity    string	   `json:"stake_net_quantity"       valid:"Required"`
	SkateCpuQuantity    string	   `json:"stake_cpu_quantity"       valid:"Required"`
}


type DelegateResult struct {
	TransactionId    string    		`json:"transaction_id"`
}

type AddPermission struct {
	Account    			 	 string     `json:"account"                   valid:"Required"`
	PrivateKey   			 string     `json:"private_key"               valid:"Required"`
	ParentPermissionName     string     `json:"parent_permission_name"    valid:"Required"`
	NewPublicKey   			 string     `json:"new_public_key"            valid:"Required"`
	NewPermissionName      	 string     `json:"new_permission_name"       valid:"Required"`
	Threshold				 int		`json:"threshold"                 valid:"Required"`
	Weight  				 int		`json:"weight"                    valid:"Required"`
}


type DeletePermission struct {
	Account    			 	 string     `json:"account"                   valid:"Required"`
	PrivateKey   			 string     `json:"private_key"               valid:"Required"`
	PermissionName   	     string     `json:"permission_name"           valid:"Required"`
}


