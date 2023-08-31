package user

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id                       uuid.UUID  `json:"id" column_name:"id"`
	Name                     string     `json:"name" column_name:"name"`
	Username                 string     `json:"username" column_name:"username"`
	Email                    string     `json:"email" column_name:"email"`
	Password                 string     `json:"-" column_name:"password"`
	Address                  string     `json:"address" column_name:"address"`
	Picture                  string     `json:"picture" column_name:"picture"`
	Isverified               bool       `json:"isverified" column_name:"isverified"`
	Tokenhash                []byte     `json:"-" column_name:"tokenhash"`
	Role                     string     `json:"role" column_name:"role"`
	GoogleId                 string     `json:"-" column_name:"google_id"`
	EmailVerifyCode          []byte     `json:"-" column_name:"email_verify_code"`
	EmailVerifyCodeCreatedAt time.Time  `json:"-" column_name:"email_verify_code_created_at"`
	IsMaintaining            bool       `json:"-" column_name:"is_maintaining"`
	CreatedAt                time.Time  `json:"-" column_name:"created_at"`
	UpdatedAt                time.Time  `json:"-" column_name:"updated_at"`
	DeletedAt                *time.Time `json:"-" column_name:"deleted_at"`
}

func (u *User) PtrFields() []any {
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

type PartialUser struct {
	Id       uuid.UUID `json:"id" column_name:"id"`
	Name     string    `json:"name" column_name:"name"`
	Username string    `json:"username" column_name:"username"`
	Password string    `json:"-" column_name:"password"`
	Email    string    `json:"email" column_name:"email"`
}

func (u *PartialUser) PtrFields() []any {
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

type SignInUser struct {
	Password string `json:"-" column_name:"password"`
	Email    string `json:"email" column_name:"email"`
}

func (u *SignInUser) PtrFields() []any {
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
