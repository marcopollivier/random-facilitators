package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:users`
}

type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Present bool   `json:"present"`
}

func main() {

	parseUsers()

}

func parseUsers() {
	var jsonFile, err = os.Open("assets/in/users.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	var jsonBytes, _ = ioutil.ReadAll(jsonFile)

	var users Users
	json.Unmarshal(jsonBytes, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Id: " + strconv.FormatInt(users.Users[i].Id, 10))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("User Present: " + users.Users[i].Name)
	}
}
