package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
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
	println("string saved: ", string(cipherBytes))
	msg := struct {
		Message string
	}{}
	err = json.Unmarshal(cipherBytes, &msg)
	if err != nil {
		fmt.Printf("error marshalling: %v\n", err)
	}
	fmt.Println("msg received : ", msg)

	decodedBytes, err := base64.StdEncoding.DecodeString(msg.Message)
	if err != nil {
		fmt.Printf("Error decoding bytes: %v\n", err)
	}
	fmt.Printf("decoded bytes : %v \n", decodedBytes)
	fmt.Printf("source bytes got : %v\n", cipherBytes)

	cipherBytes = decodedBytes

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
