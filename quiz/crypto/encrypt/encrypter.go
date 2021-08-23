package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

var keyText = "this is a NOT a good secret key." // this should be of 32 size.
func main() {
	fmt.Println("How To Encrypt Message In GO")

	text := []byte("This is a Secure Message.")
	key := []byte(keyText)

	// generate a new aes cipher using our 32 byte long key
	cipherBlock, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm is a symmetric key cryptographic block ciphers
	// https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	bytes := gcm.Seal(nonce, nonce, text, nil)
	fmt.Println(bytes)
	err = ioutil.WriteFile("crypto/encrypted-message.data", bytes, 0777)
	if err != nil {
		fmt.Printf("error writing to file : %v\n", err)
	}

}
