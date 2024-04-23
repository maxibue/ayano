package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func randomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789^°!§³$%&/{([)]=}?+*~#'<>|,;.:-_@€µ£}"
	const length = 50

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}

func GenerateAuth(username, password string) bool {
	os.Mkdir("./auth/"+username, 0755)
	cipher, nonce, salt, err := Encrypt(randomString(), password)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("./auth/%s/%s", username, username), cipher, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("./auth/%s/%s_nonce", username, username), nonce, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("./auth/%s/%s_salt", username, username), salt, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
