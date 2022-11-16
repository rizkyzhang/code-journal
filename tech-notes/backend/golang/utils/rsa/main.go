package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type RSAUtil interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(encrypted string) (string, error)
}

type baseRsaUtil struct {
	PublicKey  []byte
	PrivateKey []byte
}

func NewRSAUtil(publicKey []byte, privateKey []byte) RSAUtil {
	return &baseRsaUtil{PublicKey: publicKey, PrivateKey: privateKey}
}

func (b *baseRsaUtil) Encrypt(plaintext string) (string, error) {
	block, _ := pem.Decode(b.PublicKey)
	if block == nil {
		return "", errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, []byte(plaintext), nil)
	if err != nil {
		return "", err
	}
	encoding := base64.StdEncoding.EncodeToString([]byte(encrypted))

	return encoding, nil
}

func (b *baseRsaUtil) Decrypt(encrypted string) (string, error) {
	block, _ := pem.Decode(b.PrivateKey)
	if block == nil {
		return "", errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func main() {
	publicKey, err := os.ReadFile("public.pem")
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err := os.ReadFile("private.pem")
	if err != nil {
		fmt.Println(err)
	}

	rsaUtil := NewRSAUtil(publicKey, privateKey)
	fmt.Println(rsaUtil)
}
