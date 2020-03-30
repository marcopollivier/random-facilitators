package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Present bool   `json:"present"`
}

type Roles struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	fmt.Println("")

	var users = returnAllUsers()
	for i := 0; i < len(users.Users); i++ {
		fmt.Print("User Id: " + strconv.Itoa(users.Users[i].Id) + " | ")
		fmt.Print("User Name: " + users.Users[i].Name + " | ")
		fmt.Println("User Present: " + strconv.FormatBool(users.Users[i].Present))
	}

	fmt.Println("")

	var roles = returnAllRoles()
	for i := 0; i < len(roles.Roles); i++ {
		fmt.Print("User Id: " + strconv.Itoa(roles.Roles[i].Id) + " | ")
		fmt.Println("User Name: " + roles.Roles[i].Name)
	}

}

func returnAllUsers() Users {
	var users Users
	unmarshalFromFile("users.json", &users)
	return users
}

func returnAllRoles() Roles {
	var roles Roles
	unmarshalFromFile("roles.json", &roles)
	return roles
}

func unmarshalFromFile(assetFilename string, dataStructure interface{}) error {
	var jsonFile, _ = os.Open("assets/in/" + assetFilename)
	defer jsonFile.Close()

	var jsonBytes, _ = ioutil.ReadAll(jsonFile)

	return json.Unmarshal(jsonBytes, dataStructure)
}
