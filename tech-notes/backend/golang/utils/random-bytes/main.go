package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func EncodeBytesToString(bytes []byte) string {
	str := hex.EncodeToString(bytes)

	return str
}

func DecodeStringToBytes(str string) ([]byte, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func main() {
	randomBytes, err := GenerateRandomBytes(32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("random bytes: %v\n", randomBytes)

	str := EncodeBytesToString(randomBytes)
	fmt.Printf("encoded: %s\n", str)

	byt, err := DecodeStringToBytes(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("decoded: %v\n", byt)
}
