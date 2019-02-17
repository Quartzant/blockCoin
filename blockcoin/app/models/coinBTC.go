package models

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)


//写入数据库（各地址）
func InsertIdentification(address RegisterAddress) AddressIdentifications {

	var addrinfications  AddressIdentifications

	if address.AddressBTC != "" {
		dataBTC, _ := GetBTC(address.AddressBTC, "1", "5")
		addrinfications.BTCIdentification  = dataBTC.TotalCount
		dataUSDT,_ := GetUSDT(address.AddressBTC)
		fmt.Println(dataUSDT)
		addrinfications.USDTIdentification = dataUSDT[0].Txid
	}

	if address.AddressETH != "" {
		dataETH, _ := GetETH(address.AddressETH, "5")
		addrinfications.ETHIdentification = dataETH[0].Hash
	}

	if address.AccountEOS != "" {
		dataEOS, _ := GetEOS(address.AccountEOS, "1", "5")
		addrinfications.EOSIdentification = dataEOS.TraceCount
	}

	if address.AddressXRP != "" {
		dataXRP, _ := GetXRP(address.AddressXRP, "5")
		addrinfications.XRPIdentification = dataXRP[0].Hash
	}

	return addrinfications
}

func GenerateBTC() (string, string, error) {
	//先椭圆曲线生成私钥，打包成wif格式私钥，序列化生成公钥
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
	if err != nil {
		return "", "", err
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()

	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", err
	}

	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
}

func GenerateBTCTest() (string, string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.TestNet3Params, false)
	if err != nil {
		return "", "", err
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()

	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.TestNet3Params)
	if err != nil {
		return "", "", err
	}

	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
}