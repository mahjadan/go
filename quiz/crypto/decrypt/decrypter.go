package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

var KeyText = "this is a NOT a good secret key." // this should be of 32 size.
func main() {

	cipherBytes, err := ioutil.ReadFile("crypto/encrypted-message.data")
	if err != nil {
		fmt.Printf("error while reading file : %v\n", err)
	}
	println("size of bytes from file: ", len(cipherBytes))

	key := []byte(KeyText)
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error creating cipher :%v\n", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Printf("error creating GCM :%v\n", err)
	}
	noneSize := gcm.NonceSize()

	fmt.Println("nonceSize is : ", noneSize)

	if len(cipherBytes) < noneSize {
		fmt.Printf("error size of text is less than nonceSize")
	}

	nonce, cipherBytes := splitCipher(cipherBytes, noneSize)

	//dstBytes:= []byte("testing ")
	deCipheredBytes, err := decryptText(gcm, nonce, cipherBytes)
	if err != nil {
		fmt.Printf("error on decrypting text : %v\n", err)
	}
	fmt.Println(string(deCipheredBytes))
}

func decryptText(gcm cipher.AEAD, nonce []byte, cipherBytes []byte) ([]byte, error) {
	return gcm.Open(nil, nonce, cipherBytes, nil)
}

func splitCipher(cipherBytes []byte, nonceSize int) ([]byte, []byte) {
	return cipherBytes[:nonceSize], cipherBytes[nonceSize:]
}
