package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func Login() (string, string) {
	var username string
	var password string
	var path string
	username_while := true
	for username_while {
		fmt.Println("\n\033[36mEnter your username:\033[0m")
		fmt.Scanln(&username)
		path = "./store/" + username
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("\n\033[31mUsername doesn't exist.\033[0m")
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			username_while = false
			password_while := true
			var byte_password []byte

			for password_while {
				fmt.Println("\n\033[36mEnter your password:\033[0m")
				byte_password, err = term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					log.Fatal(err)
				}
				password = string(byte_password)

				cipher, _ := os.ReadFile(fmt.Sprintf("./auth/%s/%s", username, username))
				nonce, _ := os.ReadFile(fmt.Sprintf("./auth/%s/%s_nonce", username, username))
				salt, _ := os.ReadFile(fmt.Sprintf("./auth/%s/%s_salt", username, username))

				_, err := Decrypt(byte_password, salt, nonce, cipher)
				if err != nil {
					fmt.Println("\n\033[31mIncorrect password.\033[0m")
				}
				if err == nil {
					password_while = false
				}
			}
		}
	}
	return username, password
}
