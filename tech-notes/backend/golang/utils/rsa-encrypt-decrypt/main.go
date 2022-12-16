// WARNING: use of rsa.EncryptPKCS1v15 function to encrypt plaintexts other than session keys is dangerous. Use RSA OAEP in new protocols.

package main

import (
	"crypto/rand"
	"crypto/rsa"
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

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(plaintext))
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

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func main() {
	publicKey, err := os.ReadFile("public.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	privateKey, err := os.ReadFile("private.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	rsaUtil := NewRSAUtil(publicKey, privateKey)

	encrypted, err := rsaUtil.Encrypt("test")
	if err != nil {
		fmt.Println(err)
		return
	}
	decrypted, err := rsaUtil.Decrypt(encrypted)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)
	fmt.Println(decrypted)
}
