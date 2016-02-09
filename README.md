# toast(土司)

## 起因
   * 项目从JAVA迁移到Go后，由于之前JAVA项目AES加密时使用的是默认ECB在Go中不被支持，只能自己动手实现一个，最后越做越多，就把Go AES/DES加密解密都做了个简单的封装，方便日后使用。

## 进度
   目前支持AES/DES加密解密

### aes使用


*   cipher包定义了Cipher接口，其中声明了Encrypt和Decrypt两个方法，前者对数据进行加密，后者对数据进行解密。
*   调用NewAES()会生成一个使用AES加密算法、ECB工作模式、PKCS57填充模式的默认cipher.
*   使用NewAESWith(mode CipherMode) Cipher方法指定不同的工作模式，来构建cipher.
