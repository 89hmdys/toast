package crypto_test

import (
	"testing"
	"encoding/base64"
	. "crypto"
	"fmt"

	"github.com/89hmdys/toast/crypto"
	"github.com/89hmdys/toast/rsa"
)

func Test_DefaultClient(t *testing.T) {

	plant := `本来觉得大厅环境挺好的，可惜被包场了。然后今天的音乐可能为了开同学会的，实在太吵。就餐环境影响了美食的心情 。双人套餐的量很足，两个人吃不完呢,海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好,海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好,海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好海鲜炒饭的大虾不新鲜，其他都还好.`

	//plant:=`G|U09`

	key, err := rsa.ParsePKCS8KeyWithPEM(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCeZ6HFJXiXORcv5ljI27J8ZUb/
YIXDzRIpVN53NOgZ0NZ4OplXPumZBxR/gksskd79sPMcy9Rvpz8ZiPUKTTUuTmUM
jtL9f/E1XafVcjvUUrUILv+aJb65OiR9YHqbGSqj8B9qR5pmtyP8TAuBA2CRooBF
01WrYRHXxYv328aDWwIDAQAB
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAJ5nocUleJc5Fy/m
WMjbsnxlRv9ghcPNEilU3nc06BnQ1ng6mVc+6ZkHFH+CSyyR3v2w8xzL1G+nPxmI
9QpNNS5OZQyO0v1/8TVdp9VyO9RStQgu/5olvrk6JH1gepsZKqPwH2pHmma3I/xM
C4EDYJGigEXTVathEdfFi/fbxoNbAgMBAAECgYEAl8vxny4oYKpKCRHxlRHL+h9H
qSSDKz6Sn97/jTa7EToqvG5TUeMtEgNR5lsi1OQ4z93JK5g8zH52Hm87exK/2U0E
/o7PGAWbxV3Lyzq0FniVtBdBWyfukRj5Ig3ABUkUMcCYrpGmMCdL0TjHLF79YuVT
A6pc8asazBi70Y3QrOECQQDQQf9cTPDjK9PLEnpTpmbT4JcqPymHq3cheHYtIDnD
Ty7qJs+kxFTAS6xzaoghm97O8MAD3d2+S1E5dBsQ2oaRAkEAwrfq4Vvm0qKhnbs1
MS6qP7/VVb+zT8zj1Mb3xs581lzf0lXrsun0cjuaVkgEDeDZeXKV5MrZLOvgFW8r
lXHZKwJACk1Zfo1n1TUT0xXk60JuD8kqcTKSsV1wFT3KSs0vTlQadAbbesEjmCem
Lkd02ITHbuFF/mr5TzKWoAr4U8sboQJAVl7aUug+9MOqyJpXt98pKWngKU8FLKqH
jMRM9+Rzv2om5dey2wOnqFwD063SDo3kKVjIYFoSBzkBhsBvJrT/TQJBAIxk1xBL
Ef/7gujmusVwgiNwvJ9ipXkLvs6ec4X10HH+il3kilmiN8Ja+vieZ7LNxsExMZr1
4U0FuAJ6PsFV0HA=
-----END PRIVATE KEY-----`)
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
}

func Test_Sign_DefaultClient(t *testing.T) {

	plant := `dsfdasfrewqr234`

	key, err := rsa.ParsePKCS8KeyWithPEM(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCeZ6HFJXiXORcv5ljI27J8ZUb/
YIXDzRIpVN53NOgZ0NZ4OplXPumZBxR/gksskd79sPMcy9Rvpz8ZiPUKTTUuTmUM
jtL9f/E1XafVcjvUUrUILv+aJb65OiR9YHqbGSqj8B9qR5pmtyP8TAuBA2CRooBF
01WrYRHXxYv328aDWwIDAQAB
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAJ5nocUleJc5Fy/m
WMjbsnxlRv9ghcPNEilU3nc06BnQ1ng6mVc+6ZkHFH+CSyyR3v2w8xzL1G+nPxmI
9QpNNS5OZQyO0v1/8TVdp9VyO9RStQgu/5olvrk6JH1gepsZKqPwH2pHmma3I/xM
C4EDYJGigEXTVathEdfFi/fbxoNbAgMBAAECgYEAl8vxny4oYKpKCRHxlRHL+h9H
qSSDKz6Sn97/jTa7EToqvG5TUeMtEgNR5lsi1OQ4z93JK5g8zH52Hm87exK/2U0E
/o7PGAWbxV3Lyzq0FniVtBdBWyfukRj5Ig3ABUkUMcCYrpGmMCdL0TjHLF79YuVT
A6pc8asazBi70Y3QrOECQQDQQf9cTPDjK9PLEnpTpmbT4JcqPymHq3cheHYtIDnD
Ty7qJs+kxFTAS6xzaoghm97O8MAD3d2+S1E5dBsQ2oaRAkEAwrfq4Vvm0qKhnbs1
MS6qP7/VVb+zT8zj1Mb3xs581lzf0lXrsun0cjuaVkgEDeDZeXKV5MrZLOvgFW8r
lXHZKwJACk1Zfo1n1TUT0xXk60JuD8kqcTKSsV1wFT3KSs0vTlQadAbbesEjmCem
Lkd02ITHbuFF/mr5TzKWoAr4U8sboQJAVl7aUug+9MOqyJpXt98pKWngKU8FLKqH
jMRM9+Rzv2om5dey2wOnqFwD063SDo3kKVjIYFoSBzkBhsBvJrT/TQJBAIxk1xBL
Ef/7gujmusVwgiNwvJ9ipXkLvs6ec4X10HH+il3kilmiN8Ja+vieZ7LNxsExMZr1
4U0FuAJ6PsFV0HA=
-----END PRIVATE KEY-----`)
	if err != nil {
		t.Error(err)
		return
	}

	cipher, err := crypto.NewRSA(key)
	if err != nil {
		t.Error(err)
		return
	}

	signBytes, err := cipher.Sign([]byte(plant), SHA1)
	if err != nil {
		t.Error(err)
		return
	}

	sign := base64.StdEncoding.EncodeToString(signBytes)

	fmt.Println(sign)

	errV := cipher.Verify([]byte("dsfdasfrewqr234"), signBytes, SHA1)
	if errV != nil {
		t.Error(errV)
		return
	}
	fmt.Println("verify success")
}
