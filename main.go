package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/dzemildupljak/web_app_with_unitest/common"
	"github.com/dzemildupljak/web_app_with_unitest/config"
	db "github.com/dzemildupljak/web_app_with_unitest/db/pg"
	"github.com/dzemildupljak/web_app_with_unitest/user"
	"github.com/dzemildupljak/web_app_with_unitest/utils"
)

func init() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	config.InitDbConfig()
	db.Init()
	common.InitValidator()
}

func main() {
	mainr := mux.NewRouter()

	mainr.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("auth healthcheck")
	})

	// v1 sub router
	r := mainr.PathPrefix("/api/v1").Subrouter()
	r.Use(utils.Logger)

	user.UserRouter(r)

	appport := os.Getenv("APP_PORT")

	fmt.Println("ListenAndServe on port: " + appport)
	http.ListenAndServe(":"+appport, config.CorsCofing(mainr))
}

// func main() {
// 	ctx := context.Background()
// 	start := time.Now()
// 	// ur := user.NewUserRepo[user.User, *user.User](ctx)
// 	usrv := user.NewUserService(ctx)

// 	res, _ := usrv.LoginUser()
// 	jsonData, err := json.MarshalIndent(res, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshaling struct:", err)
// 		return
// 	}

// 	// qf := db.QueryFilter{
// 	// 	"id": "52128ce9-f47b-4464-b98a-ac697f812d6a",
// 	// }
// 	// rnum, err := ur.DeleteUser(qf)

// 	// fmt.Println(rnum)
// 	// fmt.Println(err)

// 	// qo := db.QueryOptions{
// 	// 	Limit: 200,
// 	// }
// 	// users, err := ur.GetUsers(qo)
// 	// if err != nil {
// 	// 	fmt.Println("Error getUsers:", err)
// 	// }

// 	// jsonData, err := json.MarshalIndent(users, "", "  ")
// 	// if err != nil {
// 	// 	fmt.Println("Error marshaling struct:", err)
// 	// 	return
// 	// }

// 	// for i := 70; i < 75; i++ {
// 	// 	err := ur.CreateUser(user.User{
// 	// 		Name:          "John Doe",
// 	// 		Username:      "johndoe123",
// 	// 		Email:         fmt.Sprintf("johndoe%d@example.com", i),
// 	// 		Password:      "P@ssw0rd",
// 	// 		Address:       sql.NullString{String: "123 Main St, Cityville, USA"},
// 	// 		Picture:       sql.NullString{String: "https://example.com/johndoe.jpg"},
// 	// 		Role:          "user",
// 	// 		Tokenhash:     []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJ"),
// 	// 		IsMaintaining: false,
// 	// 	})

// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// }

// 	elapsed := time.Since(start)

// 	// Print the JSON data
// 	fmt.Println(string(jsonData))
// 	fmt.Printf("Function took %s\n", elapsed)

// }
