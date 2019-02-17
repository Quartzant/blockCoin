package models

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/daisuke310vvv/infura-go"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PrivateKey btcec.PrivateKey
type PublicKey btcec.PublicKey

type Transaction struct {
	TxId               string `json:"txid"`
	SourceAddress      string `json:"source_address"`
	DestinationAddress string `json:"destination_address"`
	Amount             int64  `json:"amount"`
	UnsignedTx         string `json:"unsignedtx"`
	SignedTx           string `json:"signedtx"`
}

type Utxo struct {
	UnspentOutputs 		[]UtxoItem	`json:"unspent_outputs"`
}

type UtxoItem struct {
	TxHash				string	 `json:"tx_hash"`
	TxHashBigEndian		string	 `json:"tx_hash_big_endian"`
	TxIndex				int	 `json:"tx_index"`
	TxOutputN			int	 `json:"tx_output_n"`
	Script				string	 `json:"script"`
	Value				int	 `json:"value"`
	ValueHex			string	 `json:"value_hex"`
	Confirmations		int	 `json:"confirmations"`
}

func PrivateKeyWifFromPK(pk string) (string) {
	key := "80" + pk + "01"
	bytes, _ := hex.DecodeString(key)
	one := sha256.New()
	one.Write(bytes)
	two := sha256.New()
	two.Write(one.Sum(nil))
	checkSum := hex.EncodeToString(two.Sum(nil))[0:8]
	key += checkSum
	newKey, _ := hex.DecodeString(key)
	return base58.Encode(newKey)
}

func ImportPrivateKey(priv_key *ecdsa.PrivateKey, passphrase string, scryptN, scryptP int) (common.Address, []byte, error) {
	nkey := newKeyFromECDSA(priv_key)
	keyjson, err := keystore.EncryptKey(nkey, passphrase, scryptN, scryptP)
	if err != nil {
		return common.Address{}, nil, err
	}
	return crypto.PubkeyToAddress(nkey.PrivateKey.PublicKey), keyjson, nil
}

func ExportPrivateKey(keyjson []byte, auth string) (common.Address, *ecdsa.PrivateKey, error) {
	nkey, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return common.Address{}, nil, err
	}
	return nkey.Address, nkey.PrivateKey, nil
}

func PrivateKeyToHex(priv_key *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(priv_key))
}

func newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *keystore.Key {
	key := &keystore.Key{
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key
}

func PrivateKeyFromHex(str string) (*ecdsa.PrivateKey, error) {
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.ToECDSA(data)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func EncodeTx(tx *types.Transaction) (string, error) {
	txb, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}
	txHex := hexutil.Encode(txb)
	return txHex, nil
}

func GetNonce(address string) (output int64, code int) {
	config := infura.NewConfig("5e700452a94b430a9da21d488d337965", infura.Mainnet)
	infuraClient, _ := infura.New(config)

	input := &infura.EthGetTransactionCountInput{
		Address:        address,
		BlockParameter: infura.NewBlockParameter("latest"),
	}
	req, res := infuraClient.EthGetTransactionCountRequest(input)
	_ = req.Call()
	nonce, err := strconv.ParseInt(res.Result[2:len(res.Result)], 16, 32)
	if err != nil {
		return 0, ErrSystem
	}
	return nonce, Success
}

func PushEthTx(hex string) (output *infura.EthSendRawTransactionOutput) {
	config := infura.NewConfig("5e700452a94b430a9da21d488d337965", infura.Mainnet)
	infuraClient, _ := infura.New(config)

	input := &infura.EthSendRawTransactionInput{
		Data:        hex,
	}
	req, res := infuraClient.EthSendRawTransactionRequest(input)
	_ = req.Call()

	return res
}

func GetErc20Decimals(contractAddress string) (output *infura.EthCallOutput) {
	config := infura.NewConfig("5e700452a94b430a9da21d488d337965", infura.Mainnet)
	infuraClient, _ := infura.New(config)

	input := &infura.EthCallInput{
		Transaction: infura.Transaction{
			To: &contractAddress,
			Data: "0x313ce567",
		},
		BlockParameter: infura.NewBlockParameter("latest"),
	}
	req, res := infuraClient.EthCallRequest(input)
	_ = req.Call()

	return res
}

func GetErc20Balance(address string, contractAddress string) (output *infura.EthCallOutput) {
	config := infura.NewConfig("5e700452a94b430a9da21d488d337965", infura.Mainnet)
	infuraClient, _ := infura.New(config)

	input := &infura.EthCallInput{
		Transaction: infura.Transaction{
			To: &contractAddress,
			Data: "0x70a08231000000000000000000000000" + address[2:],
		},
		BlockParameter: infura.NewBlockParameter("latest"),
	}
	req, res := infuraClient.EthCallRequest(input)
	_ = req.Call()

	return res
}

func GetUtxos(address string) (txHash []string, total int64, code int) {
	form := Utxo{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://blockchain.info/unspent?active=" + address, nil)
	if err != nil {
		beego.Debug(err)
		return txHash, 0, ErrSystem
	}
	res, err := client.Do(req)
	if err != nil {
		beego.Debug(err)
		return txHash, 0, ErrSystem
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Debug(err)
		return txHash, 0, ErrSystem
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return txHash, 0, ErrSystem
	}

	if err := json.Unmarshal(textRes, &form); err != nil {
		beego.Debug(err)
		return txHash, 0, ErrSystem
	}
	for _, value := range form.UnspentOutputs {
		txHash = append(txHash, value.TxHashBigEndian)
		total += int64(value.Value)
	}
	return txHash, total, Success
}

func CreateTransaction(secret string, destination string, amount int64, txHash []string, fee int64, total int64) (Transaction, error) {
	var transaction Transaction
	wif, err := btcutil.DecodeWIF(secret)
	if err != nil {
		return Transaction{}, err
	}
	addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.MainNetParams)
	sourceTx := wire.NewMsgTx(wire.TxVersion)

	destinationAddress, err := btcutil.DecodeAddress(destination, &chaincfg.MainNetParams)
	sourceAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &chaincfg.MainNetParams)
	if err != nil {
		return Transaction{}, err
	}
	for _, value := range txHash {
		sourceUtxoHash, _ := chainhash.NewHashFromStr(value)
		sourceUtxo := wire.NewOutPoint(sourceUtxoHash, 0)
		sourceTxIn := wire.NewTxIn(sourceUtxo, nil, nil)
		sourceTx.AddTxIn(sourceTxIn)
	}
	destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress)
	sourcePkScript, _ := txscript.PayToAddrScript(sourceAddress)
	beego.Debug(total - amount - fee)
	sourceTxOut := wire.NewTxOut(total - amount - fee, sourcePkScript)
	sourceTx.AddTxOut(sourceTxOut)
	sourceTxHash := sourceTx.TxHash()
	redeemTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(&sourceTxHash, 0)
	redeemTxIn := wire.NewTxIn(prevOut, nil, nil)
	redeemTx.AddTxIn(redeemTxIn)
	redeemTxOut := wire.NewTxOut(amount, destinationPkScript)
	redeemTx.AddTxOut(redeemTxOut)
	sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTx.TxOut[0].PkScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return Transaction{}, err
	}
	redeemTx.TxIn[0].SignatureScript = sigScript
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(sourceTx.TxOut[0].PkScript, redeemTx, 0, flags, nil, nil, amount)
	if err != nil {
		return Transaction{}, err
	}
	if err := vm.Execute(); err != nil {
		return Transaction{}, err
	}
	var unsignedTx bytes.Buffer
	var signedTx bytes.Buffer
	sourceTx.Serialize(&unsignedTx)
	redeemTx.Serialize(&signedTx)
	transaction.TxId = sourceTxHash.String()
	transaction.UnsignedTx = hex.EncodeToString(unsignedTx.Bytes())
	transaction.Amount = amount
	transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())
	transaction.SourceAddress = sourceAddress.EncodeAddress()
	transaction.DestinationAddress = destinationAddress.EncodeAddress()
	return transaction, nil
}