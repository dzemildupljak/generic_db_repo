package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/joho/godotenv"

	"github.com/dzemildupljak/web_app_with_unitest/generic/config"
	"github.com/dzemildupljak/web_app_with_unitest/generic/db"
	"github.com/dzemildupljak/web_app_with_unitest/generic/maintenance-tokens"
	"github.com/dzemildupljak/web_app_with_unitest/generic/user"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	config.Init()
	db.Init()
}

func main() {
	start := time.Now()

	ur := user.NewUserRepo[user.User]()
	qo := db.QueryOptions{
		Filter: map[string]any{
			// "id": "36189251-b335-4ba2-9cbe-37e2f7dd3192",
		},
		OrderBy:  "id",
		OrderDir: db.DESC,
	}
	users, err := ur.GetUsers(qo)
	if err != nil {
		fmt.Println("Error getUsers:", err)
	}

	// ur := user.New[user.PartialUser]()
	// users, _ := ur.GetUser(qo)

	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling struct:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	qo = db.QueryOptions{
		Filter: map[string]interface{}{
			"user_id":    "e2e39a1d-719d-4bc4-a6c7-708f5d0795fd",
			"token_type": "forgot_password",
		},
	}
	umtr := maintenance.NewUmtRepo[maintenance.UserMTokens]()
	token, err := umtr.GetUMToken(qo)
	fmt.Println(err)

	jsonData, err = json.MarshalIndent(token, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling struct:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	elapsed := time.Since(start)
	fmt.Printf("Function took %s\n", elapsed)
}
