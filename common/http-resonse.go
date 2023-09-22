package common

import (
	"encoding/json"
	"net/http"
)

// func packTokensInHttpOnlyCookie(w http.ResponseWriter, tokens interface{}) {
// 	acc_cookie := &http.Cookie{
// 		Name:     "access_jwt_token",
// 		Value:    tokens.Access_token,
// 		HttpOnly: true,
// 		Expires:  time.Now().Add(24 * time.Hour), // Set expiration time
// 		Path:     "/",
// 	}

// 	ref_cookie := &http.Cookie{
// 		Name:     "refresh_jwt_token",
// 		Value:    tokens.Refresh_token,
// 		HttpOnly: true,
// 		Expires:  time.Now().Add(24 * time.Hour), // Set expiration time
// 		Path:     "/",
// 	}

// 	http.SetCookie(w, acc_cookie)
// 	http.SetCookie(w, ref_cookie)
// }

func AuthTokenErrorRes(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("authentication failed. token not provided or malformed")
}

func AuthCredsErrorRes(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("wrong credentials")
}
