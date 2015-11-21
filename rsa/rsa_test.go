package rsa_test

import (
	"encoding/base64"
	"fmt"
	"github.com/89hmdys/toast/rsa"
	"testing"
)

var cipher rsa.Cipher

func init() {
	client, err := rsa.NewDefault(`-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMb5ObcZNOKOooaQ
+TdWCyGJMkOG9kQh15xqgVWzRqI1DDcx+g5sOa1ac/ANNpJR+kWfkah5o9C7xlMS
upKSqObo2g3cAigiLVabdBZWf5SxN/aHo4UE0K0Ng2UrAe1YEke+MYQfeaiCargD
6F6oxSRBurHP6+g/DXLhSu6fohH5AgMBAAECgYBesPrzlU5ix4i3S8UevZcchaNj
GZaXhBeNO+6DL4dc6KwAlIsxU/X69wIX6uHerp7RhYgXSpRYYRSYMmCSIrO4IE+F
PFQRa4epEQN5r7MOYWf7YUoiVpLqRb5jElX09Gb9/Cese/5pmqV2k+0NEHfufYTD
iAZIHC7EYqk0nKrgdQJBAPxNsWoO07z24RbgwZEiFsR2shAE6YOWU43vdrXnCZrg
4OoQYUrhoLCBibuDOxNUIPJqLY+Ibf4M+6gRvXOEO/MCQQDJ44JokPVmJdU54MjU
uXqBqXMUEYXWiCWA2ldmNyfYpBsI2BqiI9dRerk9zTNMQCSIPkcpM9QxkEP+Zl3m
HVFjAkEA5RTh/otB43RBgdVGy7Eo/O9M09Cx+aFXcis4HQeOqAqBDOUcgbIFhd3I
IfKQhAdB9vlDLuzP+fjt0ndxDd7F0wJAbI7RiLipvAuL5FtiokA6B46+OoRRm1IK
GIdPh78QxgU4JEFP0O/E0CNViE3Wz6GOA1S5nwYni58vcJRK2XnaUQJAY+dVekQj
dsjvoscCgLi2q6LVxOIlKYRgeDB7SJ460h/fWllMxo0Y2WnzktcqM3Y4yonJSVHG
VrUyzNXDbQRNyQ==
-----END PRIVATE KEY-----`, `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
z+voPw1y4Urun6IR+QIDAQAB
-----END PUBLIC KEY-----`)

	if err != nil {
		fmt.Println(err)
	}

	cipher = client
}

func Test_DefaultClient(t *testing.T) {

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp, err := cipher.Encrypt([]byte(plant))
	if err != nil {
		t.Error(err)
		return
	}
	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
		return
	}
	pp, err := cipher.Decrypt(ppBy)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(pp))
}

//
//func Test_Sign_DefaultClient(t *testing.T) {
//
//	src := "王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔王翔"
//
//	signBytes, err := cipher.Sign([]byte(src), crypto.SHA256)
//	if err != nil {
//		t.Error(err)
//	}
//	sign := hex.EncodeToString(signBytes)
//	fmt.Println(sign)
//
//	signB, err := hex.DecodeString(sign)
//	if err != nil {
//		t.Error(err)
//	}
//	errV := cipher.Verify([]byte(src), signB, crypto.SHA256)
//	if errV != nil {
//		t.Error(errV)
//	}
//	fmt.Println("verify success")
//}
