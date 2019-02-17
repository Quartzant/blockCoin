package controllers

import (
	"blockcoin/app/models"
	"blockcoin/app/models/admin"
	"blockcoin/app/models/apiService"
	common2 "blockcoin/common"

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"math"
	"math/big"
	"strconv"
)

const erc20ABI string = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"INITIAL_SUPPLY","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`

type WalletController struct {
	BaseController
}


func (c *WalletController) GetBtcFee()  {
	as := apiService.NewAPIService("https://bitcoinfees.earn.com", nil)
	fee, err := as.BtcFee()
	if err != nil {
		c.RetError(errSystem)
	}
	sucGetInfo.Data = fee
	c.RetSuccess(sucGetInfo)

	c.ServeJSON()
}

func (c *WalletController) GetEthFee()  {
	as := apiService.NewAPIService("https://api.infura.io", nil)
	fee, _ := as.EthFee()
	//fmt.Println(fee)
	//fmt.Println(err)
	//if err != nil {
	//	c.RetError(errSystem)
	//}
	sucGetInfo.Data = fee
	c.RetSuccess(sucGetInfo)

	c.ServeJSON()
}

func (c *WalletController) GetBtcBalance()  {
	form := models.BtcBalanceForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		c.RetError(errInputData)
		return
	}
	as := apiService.NewAPIService("https://blockchain.info", nil)
	balance, err := as.BtcBalance(form.Address)
	if err != nil {
		c.RetError(errSystem)
	}
	sucGetInfo.Data = balance
	c.RetSuccess(sucGetInfo)

	c.ServeJSON()
}

//获取余额
func (c *WalletController) GetEthBalance()  {

	address := c.GetString("address")

	if len(address) != 42 {
		c.Data["json"] = common2.ErrAddress1
		c.ServeJSON()
	}
	as := apiService.NewAPIService("http://api.etherscan.io", nil)
	fmt.Println(as)
	balance, err := as.EthBalance(address)

	fmt.Println(balance)
	fmt.Println(err)
	fmt.Println(err)
	if err != nil {
		c.Data["json"] = common2.ErrAddress2
		c.ServeJSON()
	}
	sucGetInfo.Data = balance
	c.RetSuccess(sucGetInfo)

	c.ServeJSON()
}

func (c *WalletController) GetErcBalance()  {
	form := models.ErcBalanceForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		beego.Debug(err)
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		beego.Debug(err)
		c.RetError(errInputData)
		return
	}
	balanceOutput := models.GetErc20Balance(form.Address, form.ContractAddress)
	decimalsOutput := models.GetErc20Decimals(form.ContractAddress)

	output := models.ErcBalanceOutput{
		Balance: balanceOutput,
		Decimals: decimalsOutput,
	}
	sucGetInfo.Data = output
	c.RetSuccess(sucGetInfo)

	c.ServeJSON()
}

func (c *WalletController) GetErcDecimals() {
	form := models.ErcDecimalsForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		beego.Debug(err)
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		beego.Debug(err)
		c.RetError(errInputData)
		return
	}
	output := models.GetErc20Decimals(form.ContractAddress)
	sucGetInfo.Data = output
	c.RetSuccess(sucGetInfo)

	c.ServeJSON()
}

//获得交易详情
func (c *WalletController) GetEthTx()  {
	//form := models.EthBalanceForm{}
	address := c.GetString("address")
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
	//	fmt.Println(len("0x0722d55922cfa6f3bbe80e21363108dbe91ac3c4"))
	//	fmt.Println(form.Address)
	//	c.RetError(errInputData)
	//	return
	//}
	//if err := c.VerifyForm(&form); err != nil {
	//	c.RetError(errInputData)
	//	return
	//}
	as := apiService.NewAPIService("http://api.etherscan.io", nil)
	tx, _ := as.EthTx(address)
	//fmt.Println(tx)
	//fmt.Println(err)
	//if err != nil {
	//	c.RetError(errSystem)
	//}
	sucGetInfo.Data = tx
	c.RetSuccess(sucGetInfo)

	c.ServeJSON()
}

func (c *WalletController) Create() {
	//form := admin.WalletForm{}
	password := c.GetString("password")
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
	//	c.RetError(errInputData)
	//	return
	//}
	//if err := c.VerifyForm(&form); err != nil {
	//	c.RetError(errInputData)
	//	return
	//}

	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)

	btcAccount, _ := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/0'/0'/0/0"), false)
	ethAccount, _ := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0"), false)
	xrpAccount, _ := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/144'/0'/0/0"), false)

	btcPrivateKeyHex, _ := wallet.PrivateKeyHex(btcAccount)
	ethPrivateKeyHex, _ := wallet.PrivateKeyHex(ethAccount)
	xrpPrivateKeyHex, _ := wallet.PrivateKeyHex(xrpAccount)
	beego.Debug(xrpPrivateKeyHex)

	/*btcPrivateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), []byte(btcPrivateKeyHex))
	pubKeySerial := btcPrivateKey.PubKey().SerializeCompressed()
	pubKeyAddress, _ := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)*/
	btcPrivateWif := models.PrivateKeyWifFromPK(btcPrivateKeyHex)
	wif, _ := btcutil.DecodeWIF(btcPrivateWif)
	pubKeySerial := wif.PrivKey.PubKey().SerializeCompressed()
	pubKeyAddress, _ := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)

	ethPrivateKey, _ := models.PrivateKeyFromHex(ethPrivateKeyHex)
	address, keyStore, _ := models.ImportPrivateKey(ethPrivateKey, password, 65536, 1)

	output := models.WalletOutput{
		Mnemonic: mnemonic,
		BtcPrivate: btcPrivateWif,
		BtcAddress: pubKeyAddress.EncodeAddress(),
		EthPrivate: ethPrivateKeyHex,
		EthAddress: address.String(),
		EthKeyStore: string(keyStore),
	}
	sucCreataWallet.Data = output
	c.RetSuccess(sucCreataWallet)

	c.ServeJSON()
}

func (c *WalletController) ImportMnemonic() {
	form := admin.MnemonicForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		c.RetError(errInputData)
		return
	}
	wallet, _ := hdwallet.NewFromMnemonic(form.Mnemonic)
	btcAccount, _ := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/0'/0'/0/0"), false)
	ethAccount, _ := wallet.Derive(hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0"), false)

	btcPrivateKeyHex, _ := wallet.PrivateKeyHex(btcAccount)
	ethPrivateKeyHex, _ := wallet.PrivateKeyHex(ethAccount)

	/*btcPrivateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), []byte(btcPrivateKeyHex))
	pubKeySerial := btcPrivateKey.PubKey().SerializeCompressed()
	pubKeyAddress, _ := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)*/
	btcPrivateWif := models.PrivateKeyWifFromPK(btcPrivateKeyHex)
	wif, _ := btcutil.DecodeWIF(btcPrivateWif)
	pubKeySerial := wif.PrivKey.PubKey().SerializeCompressed()
	pubKeyAddress, _ := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)

	ethPrivateKey, _ := models.PrivateKeyFromHex(ethPrivateKeyHex)
	address, keyStore, _ := models.ImportPrivateKey(ethPrivateKey, form.Password, 65536, 1)

	output := models.WalletOutput{
		Mnemonic: form.Mnemonic,
		BtcPrivate: btcPrivateWif,
		BtcAddress: pubKeyAddress.EncodeAddress(),
		EthPrivate: ethPrivateKeyHex,
		EthAddress: address.String(),
		EthKeyStore: string(keyStore),
	}
	sucImportMnemonic.Data = output
	c.RetSuccess(sucImportMnemonic)

	c.ServeJSON()
}

func (c *WalletController) ImportBtcPrivate() {
	form := admin.BtcPrivateForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		c.RetError(errInputData)
		return
	}
	wif, _ := btcutil.DecodeWIF(form.Private)
	pubKeySerial := wif.PrivKey.PubKey().SerializeCompressed()
	pubKeyAddress, _ := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)
	output := models.ImportBtcPrivateOutput{
		BtcPrivate: form.Private,
		BtcAddress: pubKeyAddress.EncodeAddress(),
	}
	sucImportBtcPrivate.Data = output
	c.RetSuccess(sucImportBtcPrivate)

	c.ServeJSON()
}

func (c *WalletController) ImportEthPrivate() {
	form := admin.EthPrivateForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		c.RetError(errInputData)
		return
	}
	pk, err := models.PrivateKeyFromHex(form.Private)
	if err != nil {

	}
	address, keyStore, _ := models.ImportPrivateKey(pk, form.Password, 65536, 1)
	output := models.ImportEthPrivateOrKeystoreOutput{
		EthPrivate: form.Private,
		EthKeyStore: string(keyStore),
		EthAddress: address.String(),
	}
	sucImportEthPrivate.Data = output
	c.RetSuccess(sucImportEthPrivate)

	c.ServeJSON()
}

func (c *WalletController) ImportEthKeyStore() {
	form := admin.KeyStoreForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		c.RetError(errInputData)
		return
	}
	address, private, err := models.ExportPrivateKey([]byte(form.KeyStore), form.Password)
	if err != nil {
		c.RetError(errInputData)
		return
	}
	privateKey := models.PrivateKeyToHex(private)
	output := models.ImportEthPrivateOrKeystoreOutput{
		EthKeyStore: form.KeyStore,
		EthPrivate: privateKey,
		EthAddress: address.String(),
	}
	sucImportEthKeysotre.Data = output
	c.RetSuccess(sucImportEthKeysotre)

	c.ServeJSON()
}

//交易eth
func (c *WalletController) TransferEth() {
	form := admin.TransferEthForm{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		beego.Debug(err)
		c.RetError(errInputData)
		return
	}

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug(err)
		c.RetError(errInputData)
		return
	}
	num, code := models.GetNonce(form.FromAddress)
	if code != models.Success {
		c.RetError(errInputData)
		return
	}
	amount, _ := strconv.ParseInt(form.Amount, 10, 64)
	beego.Debug(amount)
	nonce := uint64(num)
	value := big.NewInt(amount)
	toAddress := common.HexToAddress(form.ToAddress)
	gasLimit := uint64(form.GasLimit)
	gasPrice := big.NewInt(form.GasPrice)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, common.FromHex("0x"))
	esPrvKey, err := models.PrivateKeyFromHex(form.Private)
	if err != nil {
		c.RetError(errInputData)
		return
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), esPrvKey)
	if err != nil {
		c.RetError(errInputData)
		return
	}
	signTxHex, err := models.EncodeTx(signTx)
	if err != nil {
		c.RetError(errInputData)
		return
	}
	beego.Debug(signTxHex)
	res := models.PushEthTx(signTxHex)
	beego.Debug(res)
	sucTransferEth.Data = res
	c.RetSuccess(sucTransferEth)

	c.ServeJSON()
}

func (c *WalletController) TransferErc20() {
	form := admin.TransferErc20Form{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.RetError(errInputData)
		return
	}
	if err := c.VerifyForm(&form); err != nil {
		c.RetError(errInputData)
		return
	}
	num, code := models.GetNonce(form.FromAddress)
	if code != models.Success {
		c.RetError(errInputData)
		return
	}
	// decimals
	decimalOutput := models.GetErc20Decimals(form.ContractAddress)
	decimal , _ := strconv.ParseInt(decimalOutput.Result[2:],16,64)
	beego.Debug(decimal)

	amount, _ := strconv.ParseFloat(form.Amount, 64)
	beego.Debug(amount)

	nonce := uint64(num)
	value := big.NewInt(0)
	toAddress := common.HexToAddress(form.ContractAddress)
	gasLimit := uint64(form.GasLimit)
	gasPrice := big.NewInt(form.GasPrice)

	sum := new(big.Int)
	sum.SetString(strconv.FormatFloat(math.Pow(10, float64(decimal)) * amount, 'f', -1, 64), 10)
	paddedAmount := common.LeftPadBytes(sum.Bytes(), 32)
	beego.Debug(hexutil.Encode(paddedAmount)[2:])

	data := "0xa9059cbb000000000000000000000000" + form.ToAddress[2:] + hexutil.Encode(paddedAmount)[2:]

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, common.FromHex(data))
	esPrvKey, err := models.PrivateKeyFromHex(form.Private)
	if err != nil {
		c.RetError(errInputData)
		return
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), esPrvKey)
	if err != nil {
		c.RetError(errInputData)
		return
	}
	signTxHex, err := models.EncodeTx(signTx)
	if err != nil {
		c.RetError(errInputData)
		return
	}
	beego.Debug(signTxHex)
	res := models.PushEthTx(signTxHex)
	beego.Debug(res)
	c.Data["json"] = "success"
	//c.RetSuccess(sucTransferErc)
	//
	c.ServeJSON()
}