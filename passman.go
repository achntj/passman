package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var passDB string = filepath.Join(os.Getenv("GOPATH"), "/password/pass.db")

func add(platform string, username string, password string) {
	if !get(platform, username, false) {
		data := platform + "," + username + "," + password + "\n"
		f, err := os.OpenFile(passDB, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString(data)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Printf("Already exists\n")
	}
}

func get(platform string, username string, out bool) bool {
	f, err := os.Open(passDB)
	if err != nil {
		fmt.Println(err)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		entry := strings.Split(input.Text(), ",")
		if entry[0] == platform && entry[1] == username {
			if out {
				fmt.Println(entry[1], entry[2])
			}
			return true
		}
	}
	if out {
		fmt.Printf("USER - '%s' @ PLATFORM '%s' doesn't exist\n", username, platform)
	}
	return false
}

func main() {
	var userPass string
	fmt.Printf("Enter the master pass: ")
	fmt.Scanln(&userPass)
	if userPass == os.Getenv("PASSMAN_KEY") {
		var args []string
		args = os.Args
		if args[1] == "add" {
			var platform string
			var username string
			var pass string
			fmt.Printf("Enter the platform: ")
			fmt.Scanln(&platform)
			fmt.Printf("Enter the username: ")
			fmt.Scanln(&username)
			fmt.Printf("Enter the password: ")
			fmt.Scanln(&pass)
			fmt.Println("-------------------------")
			add(platform, username, pass)
			fmt.Println("-------------------------")
		} else if args[1] == "get" {
			var platform string
			var username string
			fmt.Printf("Enter the platform: ")
			fmt.Scanln(&platform)
			fmt.Printf("Enter the username: ")
			fmt.Scanln(&username)
			fmt.Println("-------------------------")
			get(platform, username, true)
			fmt.Println("-------------------------")
		} else {
			fmt.Println("-------------------------")
			fmt.Println("Invalid operation ", args[1])
			fmt.Println("-------------------------")
		}
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Wrong MASTER PASSWORD :/")
		fmt.Println("-------------------------")
	}
}
