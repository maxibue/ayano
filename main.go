package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	setup_status, config := Setup()
	if !setup_status {
		os.Exit(1)
	}

	if config.Pompom {
		fmt.Printf("\n\033[33m%s\033[0m\n", pompom)
	}

	fmt.Print("\n")
	fmt.Printf("\033[36m%s\033[0m\n", ayano)

	if config.Register {
		var register_or_login string
		fmt.Println("\n\033[32m[R] Register\033[0m \033[36m[L] Login\033[0m")
		fmt.Scanln(&register_or_login)
		switch strings.ToLower(register_or_login) {
		case "r":
			username, password := Register()
			path := fmt.Sprintf("./store/%s", username)
			Menu(password, path)
			os.Exit(0)
		case "l":
			username, password := Login()
			path := fmt.Sprintf("./store/%s", username)
			Menu(password, path)
			os.Exit(0)
		}
		os.Exit(0)
	} else {
		username, password := Login()
		path := fmt.Sprintf("./store/%s", username)
		Menu(password, path)
		os.Exit(0)
	}
}
