package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func updateItem(user *User, commands []string) {

	//Check commands
	commandsTrue := []string{}
	for _, val := range commands {
		command := strings.Split(val, " ")
		if len(command) == 3 {
			commandsTrue = append(commandsTrue, val)
		}
	}

	commandTrue := []string{}

	for _, val := range commandsTrue {
		commandTrue = strings.Split(val, " ")
		switch commandTrue[0] {
		case "update":
			switch commandTrue[1] {
			case "name":
				user.Name = commandTrue[2]
			case "email":
				user.Email = commandTrue[2]
			case "age":
				user.Age, _ = strconv.Atoi(commandTrue[2])
			}

		case "delete":
			switch commandTrue[1] {
			case "name":
				user.Name = ""
			case "age":
				user.Age = 0
			case "email":
				user.Email = ""
			}

		}
	}
}

func main() {
	user := User{
		Name:  "Kasper",
		Age:   34,
		Email: "ksaper@gmail.com",
	}

	commandsUpdate := []string{
		"update name Rat",
		"update email emil@gmail.com dfupdate email emil@gmail.com df",
		"asd sda",
		"update age 30",
		"update email emil@gmail.com df",
		"update email ratmir@gmail.com",
	}

	commandsDelete := []string{
		"delete name Email",
		"delete age 30",
		"delete email emil@gmail.comlete email emil@gmail.com",
		"delete email emil@gmail.com",
		"sdafgssssr",
	}

	fmt.Println("main: ", user)

	updateItem(&user, commandsUpdate)
	fmt.Println("update: ", user)

	updateItem(&user, commandsDelete)
	fmt.Println("delete: ", user)

	updateItem(&user, commandsUpdate)
	fmt.Println("update: ", user)

	commandsDelete = []string{
		"delete name -",
	}
	updateItem(&user, commandsDelete)
	fmt.Println("delete name: ", user)
}
