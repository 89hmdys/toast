# toast(土司)

## 起因
   * 项目从JAVA迁移到Go后，由于之前JAVA项目AES加密时使用的是默认ECB在Go中不被支持，只能自己动手实现一个，最后越做越多，就把Go AES/DES加密解密都做了个简单的封装，方便日后使用。

## 进度
   目前支持AES/DES加密解密

### aes使用


*   cipher包定义了Cipher接口，其中声明了Encrypt和Decrypt两个方法，前者对数据进行加密，后者对数据进行解密。

    DES 加密/解密的例子：
 
        func main(){
            builder := crypto.DES([]byte("Z'{ru/^e")) //创建一个des 加密的builder

            cipher := builder.ECB() //选定工作模式 ，现在支持CBC、CFB、OFB、CTR、ECB 5种工作模式，目前默认使用PKCS7Padding

            plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

            cp := cipher.Encrypt([]byte(plant))

            cpStr := base64.URLEncoding.EncodeToString(cp)

            fmt.Println(cpStr)

            ppBy, err := base64.URLEncoding.DecodeString(cpStr)
            if err != nil {
            	fmt.Println(err)
            }
            pp := cipher.Decrypt(ppBy)

            fmt.Println(string(pp))
        } 
 
    AES 加密/解密的例子：
 
        func main(){
             builder := crypto.AES([]byte("Z'{ru/^e")) //创建一个aes 加密的builder，和DES相比，就这么点区别...

             cipher := builder.ECB() //选定工作模式 ，现在支持CBC、CFB、OFB、CTR、ECB 5种工作模式，目前默认使用PKCS7Padding

             plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

             cp := cipher.Encrypt([]byte(plant))

             cpStr := base64.URLEncoding.EncodeToString(cp)

             fmt.Println(cpStr)

             ppBy, err := base64.URLEncoding.DecodeString(cpStr)
             if err != nil {
             	fmt.Println(err)
             }
             pp := cipher.Decrypt(ppBy)

             fmt.Println(string(pp))
        }
