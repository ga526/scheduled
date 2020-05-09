package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func testAes() {
// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	result, err := AesEncrypt("www.baidu.com", "beijingtiananmen")
	if err != nil {
	panic(err)
	}
	fmt.Println(result) //zero UR5c4C1iW5mIdxrv5rxo4w==,pkcs jE7BUAKWpdJWb2ulcFWd/g==  和pthon,js相同
	origData, err := AesDecrypt(result, "beijingtiananmen")
	if err != nil {
	panic(err)
	}
	fmt.Println(string(origData))
}

func AesEncrypt(origDataStr string, keystr string) (string, error) {
	key := []byte(keystr)
	origData := []byte(origDataStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	//origData = PKCS5Padding(origData, blockSize)
	origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //iv=key
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(cryptedStr string, keystr string) ([]byte, error) {
	key := []byte(keystr)
	crypted := []byte(cryptedStr)
	block, err := aes.NewCipher(key)
	if err != nil {
	return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]

}
