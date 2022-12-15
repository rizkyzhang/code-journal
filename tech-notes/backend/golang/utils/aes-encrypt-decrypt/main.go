// Reference

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	return gcm.Open(nil, nonce, ciphertext, nil)
}

func main() {
	// The key has to be 32 bytes long!
	key, err := hex.DecodeString("5908b285ca7fbe782691319026c52080ef8f43912e2f8651a22554877ea0ffc1")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := []byte("Golang is cool!")
	fmt.Printf("key: %v\n", key)

	ciphertext, err := Encrypt(text, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s => %x\n", text, ciphertext)

	plaintext, err := Decrypt(ciphertext, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%x => %s\n", ciphertext, plaintext)
}
