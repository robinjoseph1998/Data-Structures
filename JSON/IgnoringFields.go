package main

import (
	"encoding/json"
	"fmt"
)

//Ingnoring Fields When Serializing (encoding) or deserilizing(decoding) by using "-"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` //password field ignored (kept private) when encoding or decoding
}

func main() {
	Data := User{Name: "Robin Joseph",
		Email:    "robin@gmail.com",
		Password: "robin1234"}

	jsonData, _ := json.Marshal(Data)
	fmt.Println(string(jsonData))
}
