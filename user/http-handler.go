package user

import (
	"encoding/json"
	"net/http"

	"github.com/dzemildupljak/web_app_with_unitest/common"
	"github.com/go-playground/validator/v10"
)

type UserHttpHdl struct {
	service  IUserService
	validate *validator.Validate
}

func NewUserHttpHdl() *UserHttpHdl {
	srv := NewUserService()
	return &UserHttpHdl{
		validate: common.Validate,
		service:  srv,
	}
}
func (handler *UserHttpHdl) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := handler.service.ListUsers(ctx)
	if err != nil {
		common.AuthCredsErrorRes(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
