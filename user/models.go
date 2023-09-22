package user

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type PartialUser struct {
	Id       uuid.UUID `json:"id" column_name:"id" validate:"required,uuid"`
	Name     string    `json:"name" column_name:"name" validate:"required"`
	Username string    `json:"username" column_name:"username" validate:"required"`
	Password string    `json:"-" column_name:"password" validate:"required"`
	Email    string    `json:"email" column_name:"email" validate:"required"`
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
	Password string `json:"-" column_name:"password" validate:"required"`
	Email    string `json:"email" column_name:"email" validate:"required"`
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

type CreateUser struct {
	Name          string         `json:"name" column_name:"name" validate:"required,uuid"`
	Username      string         `json:"username" column_name:"username" validate:"required"`
	Email         string         `json:"email" column_name:"email" validate:"required,email"`
	Password      string         `json:"-" column_name:"password" validate:"required"`
	Address       sql.NullString `json:"address" column_name:"address"`
	Picture       sql.NullString `json:"picture" column_name:"picture"`
	Role          string         `json:"role" column_name:"role" validate:"required,oneof=admin user moderator"`
	Tokenhash     []byte         `json:"-" column_name:"tokenhash" validate:"required"`
	IsMaintaining bool           `json:"-" column_name:"is_maintaining" validate:"required"`
	CreatedAt     time.Time      `json:"-" column_name:"created_at"`
	UpdatedAt     sql.NullTime   `json:"-" column_name:"updated_at"`
}

func (u *CreateUser) PtrFields() []any {
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

// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

// CREATE TABLE IF NOT EXISTS users (
//         id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
//         name VARCHAR NOT NULL CHECK (name <> ''),
//         username VARCHAR NOT NULL,
//         email VARCHAR NOT NULL UNIQUE CHECK (name <> ''),
//         password VARCHAR NOT NULL CHECK (name <> ''),
//         address VARCHAR,
//         picture VARCHAR,
//         isverified BOOLEAN DEFAULT false,
//         tokenhash BYTEA NOT NULL,
//         role VARCHAR NOT NULL,
//         google_id VARCHAR,
//         email_verify_code BYTEA,
//         email_verify_code_created_at TIMESTAMP,
//         is_maintaining BOOLEAN DEFAULT false,
//         created_at TIMESTAMP NOT NULL,
//         updated_at TIMESTAMP,
//         deleted_at TIMESTAMP
//     );

// CREATE OR REPLACE FUNCTION set_created_at()
// RETURNS TRIGGER AS $$
// BEGIN
//     NEW.created_at = NOW();
//     RETURN NEW;
// END;
// $$ LANGUAGE plpgsql;

// CREATE TRIGGER set_created_at_trigger
// BEFORE INSERT ON users
// FOR EACH ROW
// EXECUTE FUNCTION set_created_at();

// CREATE OR REPLACE FUNCTION set_updated_at()
// RETURNS TRIGGER AS $$
// BEGIN
//     NEW.updated_at = NOW();
//     RETURN NEW;
// END;
// $$ LANGUAGE plpgsql;

// CREATE TRIGGER set_updated_at_trigger
// BEFORE UPDATE ON users
// FOR EACH ROW
// EXECUTE FUNCTION set_updated_at();
