package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Pompom   bool `json:"pompom"`   // don't turn of pompom, he's protecting your passwords :(
	Register bool `json:"register"` // aka directly to login
	Auth     bool `json:"auth"`     // aka single password per user (recommended)
}

func Setup() (bool, Config) {
	_, auth_err := os.Stat("./auth")
	auth_status := os.IsNotExist(auth_err)
	_, store_err := os.Stat("./store")
	os.IsNotExist(store_err)
	store_status := os.IsNotExist(store_err)

	if auth_status && store_status {
		os.Mkdir("./store", 0755)
		os.Mkdir("./auth", 0755)
		fmt.Println("\033[32mAuthentication and storage successfully generated.\033[0m")
	} else if auth_status && !store_status {
		os.Mkdir("./auth", 0755)
		fmt.Println("\033[31mStorage exists but authentification doesn't.\nOld users can currently not log in if authentification is enabled.\nCheck out auth_recovery.txt for more information.\033[0m")
		fmt.Println("\033[32mNew authentication successfully generated.\033[0m")
	} else if !auth_status && store_status {
		os.Mkdir("./store", 0755)
		fmt.Println("\033[33mNo storage directory found.\033[0m")
		fmt.Println("\033[33m(If you have created users or keys before they seem to have been deleted or moved.)\033[0m")
		fmt.Println("\033[32mNew storage successfully generated.\033[0m")
	}
	if _, err := os.Stat("./config"); os.IsNotExist(err) {
		os.Mkdir("./config", 0755)
		os.WriteFile("./config/config.json", []byte(`{"pompom":true,"register":true,"auth":true}`), 0644)

		fmt.Println("\033[32mConfig successfully generated.\033[0m")
		return true, Config{Pompom: true, Register: true, Auth: true}
	}

	if _, err := os.Stat("./config"); !os.IsNotExist(err) {
		var config Config
		read_config, _ := os.ReadFile("./config/config.json")

		err := json.Unmarshal(read_config, &config)

		if err != nil {
			os.WriteFile("./config/config.json", []byte(`{"pompom":true,"register":true,"auth":true}`), 0644)
			fmt.Println("\033[31mConfig was incomplete and therefore regenerated.\033[0m")
			return true, Config{Pompom: true, Register: true, Auth: true}
		}
		return true, config
	}

	fmt.Println("\033[31mConfig failed.\033[0m")
	return false, Config{}
}
