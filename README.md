# toast(土司)

## 起因
   * 觉得Go的加密解密封装程度感觉略低，用起来不是很顺手，这里把之前项目里用到的加解密汇总了下，希望对大家有所帮助。

## 进度
   目前支持AES加密解密，RSA 使用PKCS1/PKCS8秘钥加密/解密/签名/验签。

### aes使用


*   aes包定义了Cipher接口，其中声明了Encrypt和Decrypt两个方法，前者对数据进行加密，后者对数据进行解密。您可以调用NewDefault(key string) (Cipher, error) 获得一个使用CBC工作模式，PKCS7Padding填充方式的Cipher实例，或者使用New(key []byte, cipherType ct.Type, paddingType pt.Type)(Cipher,  error)创建一个自定义的Cipher实例，cipherType和paddingType是分别定义在 toast/aes/cipher和 toast/aes/padding中的常量
 
    默认Cipher加密例子：
 
        func main(){
            client, err := aes.NewDefault([]byte("3|$asKitICV.eua*h)W1[7kG"))
            if err != nil {
	        	   fmt.Println(err)
	        }
            cp := client.Encrypt([]byte("b,c"))
            cpStr := base64.URLEncoding.EncodeToString(cp)
            fmt.Println(cpStr)
        } 
 
    默认Cipher解密例子：
 
        func main(){
            client, err := aes.NewDefault([]byte("3|$asKitICV.eua*h)W1[7kG"))
            ppBy, err := base64.URLEncoding.DecodeString("BMLdfSUpfeR00CjzVaDSww==")
	        if err != nil {
	        	   fmt.Println(err)
	        }
	        pp := client.Decrypt(ppBy)
            fmt.Println(string(pp))
        } 
 
    自定义Cipher加密例子：
 
        func main(){
            client, err := aes.New([]byte("3|$asKitICV.eua*h)W1[7kG"), cipher.CFB, padding.PKCS7)
            if err != nil {
	        	   fmt.Println(err)
	        }
            cp := client.Encrypt([]byte("b,c"))
            cpStr := base64.URLEncoding.EncodeToString(cp)
            fmt.Println(cpStr)
        } 
 
    自定义Cipher解密例子：
 
        func main(){
            client, err := aes.New([]byte("3|$asKitICV.eua*h)W1[7kG"), cipher.CFB, padding.PKCS7)
            ppBy, err := base64.URLEncoding.DecodeString("BMLdfSUpfeR00CjzVaDSww==")
	        if err != nil {
	        	   fmt.Println(err)
	        }
	        pp := client.Decrypt(ppBy)
            fmt.Println(string(pp))
        } 

### rsa使用

*   rsa包定义了Cipher接口，其中声明了Encrypt、Decrypt、Sign、Verify四个方法，分别对数据进行加密、解密、签名和验签。您可以调用NewDefault(privateKey, publicKey string) (Cipher, error) 方法，传入PEM编码的公钥/PKCS8格式的私钥来创建一个Cipher实例，或者调用New(privateKey, publicKey []byte, privateKeyType privatekey.Type) (Cipher, error)方法，传入解码后的公钥/私钥来创建一个Cipher实例。privateKeyType是定义在toast/rsa/privatekey中的常量。 
关于pkcs pem这些概念，参见：http://my.oschina.net/u/1023800/blog/526647


    默认Cipher加密例子：
 
        func main(){
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
                                          -----END PRIVATE KEY-----`, 
                                          `-----BEGIN PUBLIC KEY-----
                                          MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                          hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                          3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                          z+voPw1y4Urun6IR+QIDAQAB
                                          -----END PUBLIC KEY-----`)

            if err != nil {
                fmt.Println(err)
            }
            cp,err := client.Encrypt([]byte("b,c"))
            if err != nil {
                fmt.Println(err)
            }
            cpStr := base64.URLEncoding.EncodeToString(cp)
            fmt.Println(cpStr)
        } 
 
    默认Cipher解密例子：
 
        func main(){
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
                                          -----END PRIVATE KEY-----`, 
                                          `-----BEGIN PUBLIC KEY-----
                                          MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                          hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                          3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                          z+voPw1y4Urun6IR+QIDAQAB
                                          -----END PUBLIC KEY-----`)

            if err != nil {
                fmt.Println(err)
            }
            ppBy, err := base64.URLEncoding.DecodeString("密文")
            if err != nil {
                   fmt.Println(err)
            }
            pp := client.Decrypt(ppBy)
            fmt.Println(string(pp))
        } 
 
    默认Cipher签名例子：
 
        func main(){
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
                                          -----END PRIVATE KEY-----`, 
                                          `-----BEGIN PUBLIC KEY-----
                                          MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                          hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                          3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                          z+voPw1y4Urun6IR+QIDAQAB
                                          -----END PUBLIC KEY-----`)

            signBytes, err := client.Sign([]byte(src), crypto.SHA256)
            if err != nil {
                fmt.Println(err)
            }
            sign := hex.EncodeToString(signBytes)
            fmt.Println(sign)
        } 
 
    默认Cipher验签例子：
 
        func main(){
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
                                          -----END PRIVATE KEY-----`, 
                                          `-----BEGIN PUBLIC KEY-----
                                          MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                          hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                          3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                          z+voPw1y4Urun6IR+QIDAQAB
                                          -----END PUBLIC KEY-----`)

            signB, err := hex.DecodeString("你的签名值")
            if err != nil {
                fmt.Println(err)
            }
            errV := cipher.Verify([]byte(src), signB, crypto.SHA256)
            if errV != nil {
                fmt.Println(errV)
            }
            fmt.Println("verify success")
        }

    自定义Cipher加密例子：
 
        func main(){
            prk, err := base64.StdEncoding.DecodeString(`MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMb5ObcZNOKOooaQ
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
                                 VrUyzNXDbQRNyQ==`)
            if err != nil {
              fmt.Println(err)
            }
          
            puk, err := base64.StdEncoding.DecodeString(`MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                                           hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                                           3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                                           z+voPw1y4Urun6IR+QIDAQAB`)
            if err != nil {
              fmt.Println(err)
            }
          
            client, err := rsa.New(prk, puk, privatekey.PKCS8)

            if err != nil {
                fmt.Println(err)
            }
            cp,err := client.Encrypt([]byte("b,c"))
            if err != nil {
                fmt.Println(err)
            }
            cpStr := base64.URLEncoding.EncodeToString(cp)
            fmt.Println(cpStr)
        } 
 
    默认Cipher解密例子：
 
        func main(){
            prk, err := base64.StdEncoding.DecodeString(`MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMb5ObcZNOKOooaQ
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
                                 VrUyzNXDbQRNyQ==`)
            if err != nil {
              fmt.Println(err)
            }
          
            puk, err := base64.StdEncoding.DecodeString(`MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                                           hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                                           3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                                           z+voPw1y4Urun6IR+QIDAQAB`)
            if err != nil {
              fmt.Println(err)
            }
          
            client, err := rsa.New(prk, puk, privatekey.PKCS8)
            if err!=nil{
                fmt.Println(err)
            }

            ppBy, err := base64.URLEncoding.DecodeString("密文")
            if err != nil {
                   fmt.Println(err)
            }
            pp := client.Decrypt(ppBy)
            fmt.Println(string(pp))
        } 
 
    默认Cipher签名例子：
 
        func main(){
            prk, err := base64.StdEncoding.DecodeString(`MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMb5ObcZNOKOooaQ
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
                                 VrUyzNXDbQRNyQ==`)
            if err != nil {
              fmt.Println(err)
            }
          
            puk, err := base64.StdEncoding.DecodeString(`MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                                           hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                                           3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                                           z+voPw1y4Urun6IR+QIDAQAB`)
            if err != nil {
              fmt.Println(err)
            }
          
            client, err := rsa.New(prk, puk, privatekey.PKCS8)
            if err != nil{
                fmt.Println(err)
            }

            signBytes, err := client.Sign([]byte(src), crypto.SHA256)
            if err != nil {
                fmt.Println(err)
            }
            sign := hex.EncodeToString(signBytes)
            fmt.Println(sign)
        } 
 
    默认Cipher验签例子：
 
        func main(){
            prk, err := base64.StdEncoding.DecodeString(`MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMb5ObcZNOKOooaQ
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
                                 VrUyzNXDbQRNyQ==`)
            if err != nil {
              fmt.Println(err)
            }
          
            puk, err := base64.StdEncoding.DecodeString(`MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG+Tm3GTTijqKGkPk3VgshiTJD
                                                           hvZEIdecaoFVs0aiNQw3MfoObDmtWnPwDTaSUfpFn5GoeaPQu8ZTErqSkqjm6NoN
                                                           3AIoIi1Wm3QWVn+UsTf2h6OFBNCtDYNlKwHtWBJHvjGEH3mogmq4A+heqMUkQbqx
                                                           z+voPw1y4Urun6IR+QIDAQAB`)
            if err != nil {
              fmt.Println(err)
            }
          
            client, err := rsa.New(prk, puk, privatekey.PKCS8)
            if err != nil{
                fmt.Println(err)
            }

            errV := cipher.Verify([]byte(src), signB, crypto.SHA256)
            if errV != nil {
                fmt.Println(errV)
            }
            fmt.Println("verify success")
        }