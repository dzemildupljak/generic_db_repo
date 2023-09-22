package maintenance

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type UMTokenType string

const (
	ForgotPwdToken       UMTokenType = "forgot_password"
	EmailVerifyCodeToken UMTokenType = "email_verify"
)

type UserMTokens struct {
	Id        uuid.UUID   `json:"id" column_name:"id"`
	TokenType UMTokenType `json:"token_type" column_name:"token_type"`
	Token     string      `json:"token" column_name:"token"`
	UserId    uuid.UUID   `json:"user_id" column_name:"user_id"`
	CreatedAt time.Time   `json:"-" column_name:"created_at"`
	UpdatedAt time.Time   `json:"-" column_name:"updated_at"`
	DeletedAt *time.Time  `json:"-" column_name:"deleted_at"`
}

func (u *UserMTokens) PtrFields() []any {
	val := reflect.ValueOf(u).Elem()
	fieldsCount := val.NumField()

	pointers := make([]any, fieldsCount)
	for i := 0; i < fieldsCount; i++ {
		field := val.Field(i)
		if field.CanAddr() {
			pointers[i] = field.Addr().Interface()
		} else {
			pointers[i] = nil
		}
	}

	return pointers
}
