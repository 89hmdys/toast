package crypto_test

import (
	. "crypto"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/89hmdys/toast/crypto"
	"github.com/89hmdys/toast/rsa"
)

func Test_LoadFromPEMFile(t *testing.T) {

	plant := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

	key, err := rsa.LoadKeyFromPEMFile(
		`/Users/alex/Documents/go/src/toast/crypto/rsa_public_key.pem`,
		`/Users/alex/Documents/go/src/toast/crypto/rsa_private_key.pem`,
		rsa.ParsePKCS8Key)
	if err != nil {
		t.Error(err)
		return
	}
	//
	cipher, err := crypto.NewRSA(key)
	if err != nil {
		t.Error(err)
		return
	}

	enT, err := cipher.Encrypt([]byte(plant))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(enT))

	deT, err := cipher.Decrypt(enT)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(deT))

	signBytes, err := cipher.Sign([]byte(plant), SHA1)
	if err != nil {
		t.Error(err)
		return
	}

	sign := base64.StdEncoding.EncodeToString(signBytes)

	fmt.Println(sign)

	errV := cipher.Verify([]byte(plant), signBytes, SHA1)
	if errV != nil {
		t.Error(errV)
		return
	}

	fmt.Println("verify success")
}
