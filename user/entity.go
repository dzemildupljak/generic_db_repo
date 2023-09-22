package user

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id                       uuid.UUID      `json:"id" column_name:"id" validate:"required,uuid"`
	Name                     string         `json:"name" column_name:"name" validate:"required"`
	Username                 string         `json:"username" column_name:"username" validate:"required"`
	Email                    string         `json:"email" column_name:"email" validate:"required,email"`
	Password                 string         `json:"-" column_name:"password" validate:"required"`
	Address                  sql.NullString `json:"address" column_name:"address"`
	Picture                  sql.NullString `json:"picture" column_name:"picture"`
	Isverified               bool           `json:"isverified" column_name:"isverified" `
	Role                     string         `json:"role" column_name:"role" validate:"required,oneof=admin user moderator"`
	Tokenhash                []byte         `json:"-" column_name:"tokenhash" validate:"required"`
	GoogleId                 sql.NullString `json:"-" column_name:"google_id"`
	EmailVerifyCode          sql.NullByte   `json:"-" column_name:"email_verify_code"`
	EmailVerifyCodeCreatedAt sql.NullTime   `json:"-" column_name:"email_verify_code_created_at"`
	IsMaintaining            bool           `json:"-" column_name:"is_maintaining" validate:"required"`
	CreatedAt                time.Time      `json:"-" column_name:"created_at"`
	UpdatedAt                sql.NullTime   `json:"-" column_name:"updated_at"`
	DeletedAt                sql.NullTime   `json:"-" column_name:"deleted_at"`
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
