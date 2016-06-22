package crypto_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/89hmdys/toast/cipher"
	"github.com/89hmdys/toast/crypto"
)

var key = "|-8xrgPfS)Aa4xtAIL^k5qX)$Y5Rim9Z"

func Test_AES_ECB(t *testing.T) {

	cipher, err := crypto.NewAES([]byte(key))
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_AES_CBC ok")
}

func Test_AES_CBC(t *testing.T) {

	mode := cipher.NewCBCMode()                         //加密工作模式，支持 CBC ECB CFB CTR 四种工作模式
	cipher, err := crypto.NewAESWith([]byte(key), mode) //创建一个AES 加密的builder
	if err != nil {
		t.Error(err)
		return
	}

	planttext := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

	ciphertext := cipher.Encrypt([]byte(planttext))

	ciphertextWithBase64 := base64.URLEncoding.EncodeToString(ciphertext)

	fmt.Println(ciphertextWithBase64)

	ciphertext, err = base64.URLEncoding.DecodeString(ciphertextWithBase64)
	if err != nil {
		t.Error(err)
	}
	planttextBytes := cipher.Decrypt(ciphertext)

	fmt.Println(string(planttextBytes))
}

func Test_AES_CFB(t *testing.T) {

	mode := cipher.NewCFBMode()

	cipher, err := crypto.NewAESWith([]byte(key), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_AES_CFB ok")
}

func Test_AES_OFB(t *testing.T) {

	mode := cipher.NewOFBMode()

	cipher, err := crypto.NewAESWith([]byte(key), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_AES_OFB ok")
}

func Test_AES_CTR(t *testing.T) {

	mode := cipher.NewCTRMode()

	cipher, err := crypto.NewAESWith([]byte(key), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_AES_CTR ok")
}
