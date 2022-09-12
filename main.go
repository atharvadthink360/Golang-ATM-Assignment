package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	username string
	password int
	balance float32
	statement []string
}
 
func main() {

	var usernameVal string
	var passwordVal int
	var balanceVal float32
	var statementVal []string

	fmt.Println("Enter a Username:");
	fmt.Scanln(&usernameVal);
	fmt.Println("Enter a Password:");
	fmt.Scanln(&passwordVal);
	fmt.Println("Enter Initial Balance:");
	fmt.Scan(&balanceVal);

	data := User{
		usernameVal,
		passwordVal,
		balanceVal,
		statementVal,
    }

	fmt.Println(data);
 
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("test.json", file, 0644)
}