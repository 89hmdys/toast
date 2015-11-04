# toast(土司)

## 起因
   * 觉得Go的加密解密封装程度感觉略低，用起来不是很顺手，这里把之前项目里用到的加解密汇总了下，希望对大家有所帮助。

## 进度
   目前仅支持AES加密解密，RSA加密解密验签努力开发ing

### aes使用


*   aes包定义了Cipher接口，其中声明了Encrypt和Decrypt两个方法，前者对数据进行加密，后者对数据进行解密。您可以调用NewDefault(key string) (Cipher, error) 获得一个使用CBC工作模式，PKCS7Padding填充方式的Cipher实例，或者使用New(key []byte, cipherType, paddingType int64)(Cipher,  error)创建一个自定义的Cipher实例，cipherType和paddingType是分别定义在 toast/aes/cipher和 toast/aes/padding中的常量
 
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
	        	   t.Error(err)
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
	        	   t.Error(err)
	        }
	        pp := client.Decrypt(ppBy)
            fmt.Println(string(pp))
        } 
