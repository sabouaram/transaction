package main

import (
	"encoding/json"
	"fmt"
	"github.com/sabouaram/transaction/controller"
	"github.com/sabouaram/transaction/models"
	"github.com/sabouaram/transaction/sql"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

func init() {
	_, err := sql.Clientdb.Exec("DELETE FROM users")
	if err != nil {
		log.Fatalf("Unable to clean the users table")
	}
	resp, err := http.Get("https://git.io/Jm76h")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
	var User models.User
	for _, user := range users {
		balance, _ := strconv.ParseFloat(user.Balance, 64)
		User.Id = user.Id
		User.Name = user.Name
		User.Balance = balance
		User.CreateUser()
	}
}
func main() {
	router := controller.NewRouter()
	server := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}
	log.Fatal(server.ListenAndServe())
}
