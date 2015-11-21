package aes_test

import (
	"encoding/base64"
	"fmt"
	"github.com/89hmdys/toast/aes"
	"github.com/89hmdys/toast/aes/cipher"
	"github.com/89hmdys/toast/aes/padding"
	"testing"
)

func Test_DEFAULT(t *testing.T) {

	client, err := aes.NewDefault([]byte("3|$asKitICV.eua*h)W1[7kG"))

	if err != nil {
		t.Error(err)
	}
	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := client.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := client.Decrypt(ppBy)
	fmt.Println(string(pp))
}

func Test_CFBandPKCS7(t *testing.T) {

	client, err := aes.New([]byte("3|$asKitICV.eua*h)W1[7kG"), cipher.CFB, padding.PKCS7)

	if err != nil {
		t.Error(err)
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := client.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := client.Decrypt(ppBy)
	fmt.Println(string(pp))
}

func Test_OFBandPKCS7(t *testing.T) {

	client, err := aes.New([]byte("3|$asKitICV.eua*h)W1[7kG"), cipher.OFB, padding.PKCS7)

	if err != nil {
		t.Error(err)
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := client.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := client.Decrypt(ppBy)
	fmt.Println(string(pp))
}
