package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//PKCS7 填充模式
func PKCS7Padding(ciphertext string, blockSize int) string {
	ciphertextByte := []byte(ciphertext)
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtextByte := bytes.Repeat([]byte{byte(padding)}, padding)
	return string(append(ciphertextByte, padtextByte...))
}

//填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

/*
   AES  CBC 加密
   key:加密key
   plaintext：加密明文
   ciphertext:解密返回字节字符串[ 整型以十六进制方式显示]

*/
func AESCBCEncrypt(key, plaintext string) (ciphertext string) {
	plaintext = PKCS7Padding(plaintext, 16)
	plainbyte := []byte(plaintext)
	keybyte := []byte(key)
	if len(plainbyte)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err)
	}

	cipherbyte := make([]byte, aes.BlockSize+len(plainbyte))
	iv := cipherbyte[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherbyte[aes.BlockSize:], plainbyte)

	ciphertext = fmt.Sprintf("%x\n", cipherbyte)
	return
}

/*
   AES  CBC 解码
   key:解密key
   ciphertext:加密返回的串
   plaintext：解密后的字符串
*/
func AESCBCDecrypter(key, ciphertext string) (plaintext string) {
	cipherbyte, _ := hex.DecodeString(ciphertext)
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err)
	}
	if len(cipherbyte) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := cipherbyte[:aes.BlockSize]
	cipherbyte = cipherbyte[aes.BlockSize:]
	if len(cipherbyte)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherbyte, cipherbyte)
	cipherbyte, _ = PKCS7UnPadding(cipherbyte)
	//fmt.Printf("%s\n", ciphertext)
	plaintext = string(cipherbyte[:])
	return
}

/*
   AES  GCM 加密
   key:加密key
   plaintext：加密明文
   ciphertext:解密返回字节字符串[ 整型以十六进制方式显示]

*/
func AESGCMEncrypt(key, plaintext string) (ciphertext, noncetext string) {
	plainbyte := []byte(plaintext)
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err.Error())
	}

	// 由于存在重复的风险，请勿使用给定密钥使用超过2^32个随机值。
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherbyte := aesgcm.Seal(nil, nonce, plainbyte, nil)
	ciphertext = fmt.Sprintf("%x\n", cipherbyte)
	noncetext = fmt.Sprintf("%x\n", nonce)
	return
}

/*
   AES  CBC 解码
   key:解密key
   ciphertext:加密返回的串
   plaintext：解密后的字符串
*/
func AESGCMDecrypter(key, ciphertext, noncetext string) (plaintext string) {

	cipherbyte, _ := hex.DecodeString(ciphertext)
	nonce, _ := hex.DecodeString(noncetext)
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainbyte, err := aesgcm.Open(nil, nonce, cipherbyte, nil)
	if err != nil {
		panic(err.Error())
	}

	//fmt.Printf("%s\n", ciphertext)
	plaintext = string(plainbyte[:])
	return
}
