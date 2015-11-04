package aes_test

import (
	"encoding/base64"
	"fmt"
	"testing"
	"toast/aes"
	"toast/aes/cipher"
	"toast/aes/padding"
)

func Test_DEFAULT(t *testing.T) {

	client, err := aes.NewDefault([]byte("3|$asKitICV.eua*h)W1[7kG"))

	if err != nil {
		t.Error(err)
	}

	cp := client.Encrypt([]byte("b,c"))

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

	cp := client.Encrypt([]byte("b,c"))

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

	cp := client.Encrypt([]byte("b,c"))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := client.Decrypt(ppBy)
	fmt.Println(string(pp))
}
