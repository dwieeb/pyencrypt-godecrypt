package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"log"
)

func Decrypt(text []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte("example key 1234"))
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Decrypt: aes.NewCipher: %v", err)
	}
	if len(text) < aes.BlockSize {
		return nil, fmt.Errorf("[ERROR] Decrypt: ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	log.Println(len(text))
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	return text, nil
}

func main() {
	b, err := ioutil.ReadFile("msg")
	if err != nil {
		panic(err)
	}

	d, err := Decrypt(b)
	if err != nil {
		panic(err)
	}

	log.Printf("%s", d)
	
	if string(d) != "my small message" {
		log.Printf("Well, that didn't work.")
	}
}
