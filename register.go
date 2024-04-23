package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

func Register() (string, string) {
	var username string
	var password string
	var path string
	username_while := true
	for username_while {
		fmt.Println("\n\033[33mSelect a username:\033[0m")
		fmt.Scanln(&username)
		path = "./store/" + username
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			fmt.Println("\n\033[31mUsername already exists.\033[0m")
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			register_while := true
			fmt.Println("\n\033[32mRegistering new user...\033[0m")
			var confirm_password string
			var byte_password []byte
			var byte_confirm_password []byte

			for register_while {
				fmt.Println("\n\033[33mPlease set a password for user: " + username + "\033[0m")
				fmt.Println("\n\033[36mEnter your new password:\033[0m")
				byte_password, err = term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					log.Fatal(err)
				}
				password = string(byte_password)
				fmt.Println("\n\033[36mConfirm your new password:\033[0m")
				byte_confirm_password, err = term.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					log.Fatal(err)
				}
				confirm_password = string(byte_confirm_password)
				if password != confirm_password {
					fmt.Println("\033[31mPasswords do not match.\033[0m")
				}
				if password == confirm_password {
					register_while = false
				}
			}

			os.Mkdir(path, 0755)
			if GenerateAuth(username, password) {
				fmt.Println("\n\033[32mUser \"" + username + "\" has been registered successfully.\033[0m")

				fmt.Println("\n\033[31mThis is the only time you can ever see the password you just set.\033[0m")
				fmt.Println("\033[36mDo you want to see the password you just set one more time? [Y/N]\033[0m")
				var password_y_n string
				fmt.Scanln(&password_y_n)
				switch strings.ToLower(password_y_n) {
				case "y":
					fmt.Println("\n\033[31mYour password is: " + password + "\033[0m")
				case "n":
					fmt.Println("\n\033[32mYour password will not be shown again.\033[0m")
				}
				username_while = false
			} else {
				os.Exit(1)
			}
		}
	}
	return username, password
}
