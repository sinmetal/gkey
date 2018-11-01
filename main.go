package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	encKey := make([]byte, 32)
	_, err := rand.Read(encKey)
	if err != nil {
		panic(err)
	}

	str := base64.StdEncoding.EncodeToString(encKey)
	fmt.Println(str)
}
