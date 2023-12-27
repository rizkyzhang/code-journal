---
tags:
  - golang-util
---
```go
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

type AesEncryptUtil interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
}

type baseAesEncrypt struct {
	key []byte
}

func NewAesEncrypt(_key string) AesEncryptUtil {
	key, _ := hex.DecodeString(_key)

	return &baseAesEncrypt{
		key: key,
	}
}

func (b *baseAesEncrypt) Encrypt(plaintext string) (string, error) {
	c, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	byt, err := gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(byt), nil
}

func (b *baseAesEncrypt) Decrypt(ciphertextStr string) (string, error) {
	ciphertextByt, err := hex.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(b.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertextByt) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertextByt[:nonceSize], ciphertextByt[nonceSize:]

	byt, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}

func main() {
	// The key has to be 32 bytes long!
	key := "5908b285ca7fbe782691319026c52080ef8f43912e2f8651a22554877ea0ffc1"
	aes := NewAesEncrypt(key)

	ciphertext, err := aes.Encrypt("test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ciphertext)

	plaintext, err := aes.Decrypt(ciphertext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(plaintext)
}
```
