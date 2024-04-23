package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Menu(password string, path string) {
	while := true
	for while {
		fmt.Println("\n\033[32m[A] Add a new password\033[0m \033[36m[R] Read a password\033[0m \033[31m[D] Delete a password\033[0m \033[35m[Q] Quit\033[0m")
		var command string
		fmt.Scanln(&command)
		switch strings.ToLower(command) {
		case "a":
			var new_password_name string
			var new_password string
			fmt.Println("\nPlease provide a name for the new password:")
			fmt.Scanln(&new_password_name)
			fmt.Println("\nPlease enter the actual password:")
			fmt.Scanln(&new_password)
			cipher, nonce, salt, err := Encrypt(new_password, password)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile(fmt.Sprintf("%s/%s", path, new_password_name), cipher, 0644)
			os.WriteFile(fmt.Sprintf("%s/%s_nonce", path, new_password_name), nonce, 0644)
			os.WriteFile(fmt.Sprintf("%s/%s_salt", path, new_password_name), salt, 0644)
			fmt.Println("\n\033[32mNew password has been added.\033[0m")
		case "r":
			var requested_password_name string
			fmt.Println("\nPlease provide the name of the password you want to read:")
			fmt.Scanln(&requested_password_name)
			cipher, err := os.ReadFile(fmt.Sprintf("%s/%s", path, requested_password_name))
			if err != nil {
				fmt.Println("\n\033[31mPassword not found.\033[0m")
				break
			}
			nonce, _ := os.ReadFile(fmt.Sprintf("%s/%s_nonce", path, requested_password_name))
			salt, _ := os.ReadFile(fmt.Sprintf("%s/%s_salt", path, requested_password_name))

			requested_password, decrypt_err := Decrypt([]byte(password), salt, nonce, cipher)
			if decrypt_err != nil {
				fmt.Println("\n\033[31mPassword decryption failed.\033[0m")
				break
			}
			if decrypt_err == nil {
				fmt.Println("\n\033[32mThe password for \"" + requested_password_name + "\" is: " + string(requested_password) + "\033[0m")
			}
		case "d":
			var requested_password_name string
			fmt.Println("\nPlease provide the name of the password you want to delete:")
			fmt.Scanln(&requested_password_name)
			fmt.Println("\n\n\033[31mAre you sure you want to delete the password: " + requested_password_name + " [Y/N]\033[0m")
			var delete_y_n string
			fmt.Scanln(&delete_y_n)
			switch strings.ToLower(delete_y_n) {
			case "y":
				os.Remove(fmt.Sprintf("%s/%s", path, requested_password_name))
				os.Remove(fmt.Sprintf("%s/%s_nonce", path, requested_password_name))
				os.Remove(fmt.Sprintf("%s/%s_salt", path, requested_password_name))
				fmt.Println("\n\n\033[32mSuccessfully deleted password: " + requested_password_name + "\033[0m")
			case "n":
				fmt.Println("\n\033[31mPassword deletion cancelled.\033[0m")
			}
		case "q":
			fmt.Println("\n\033[31mQuitting Ayano...\033[0m")
			while = false
		}
	}
}
