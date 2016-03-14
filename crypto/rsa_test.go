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

	plant := `本来觉得大厅环境挺好的，可惜被包场了。
	然后今天的音乐可能为了开同学会的，实在太吵。就餐环境影响了美食的心情 。
	双人套餐的量很足，两个人吃不完呢,海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜
	其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，
	其他都还好海鲜炒饭的大虾不新鲜，其他都还好,海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好,
	海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，
	其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，
	其他都还好海鲜炒饭的大虾不新鲜，其他都还好.`

	key, err := rsa.LoadPKCS8KeyFromPEMFile(
		`/Users/alex/Documents/go/src/toast/crypto/rsa_public_key.pem`,
		`/Users/alex/Documents/go/src/toast/crypto/rsa_private_key.pem`)
	if err != nil {
		t.Error(err)
		return
	}

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
