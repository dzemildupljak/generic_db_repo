package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/dzemildupljak/web_app_with_unitest/common"
	db "github.com/dzemildupljak/web_app_with_unitest/db/pg"
	"github.com/go-playground/validator/v10"
)

const userTable = "public.users"
const createUsersQuery = "INSERT INTO " + userTable + " (%s)"
const getUsersQuery = "SELECT %s FROM " + userTable
const deleteUsersQuery = "DELETE FROM " + userTable

type urepoType interface {
	User | SignInUser | PartialUser | CreateUser
}

type UserRepo[T urepoType, PT common.Ptr[T]] struct {
	validate *validator.Validate
	ctx      context.Context
}

func NewUserRepo[T urepoType, PT common.Ptr[T]](ctx context.Context) UserRepo[T, PT] {
	ur := UserRepo[T, PT]{
		validate: common.Validate,
		ctx:      ctx,
	}

	return ur
}

func (ur *UserRepo[T, PT]) CreateUser(user T) error {
	if err := ur.validate.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field %s failed validation with tag '%s'\n", err.Field(), err.Tag())
		}
	}

	columnsName, columnsValue := db.GetColumnsAndValues[T, PT](user)
	query := fmt.Sprintf(createUsersQuery, strings.Join(columnsName, ", "))

	var colmnsArg []string
	for i := 0; i < len(columnsName); i++ {
		colmnsArg = append(colmnsArg, fmt.Sprintf("$%d", i+1))
	}

	query += fmt.Sprintf(" VALUES (%s)", strings.Join(colmnsArg, ", "))
	err := db.Create(ur.ctx, query, columnsValue...)
	if err != nil {
		// TODO: add persistent logger for db errors
		return err
	}

	return nil
}

func (ur *UserRepo[T, PT]) GetUser(qo db.QueryFilter) (T, error) {
	var user T

	validColumns, columnsNames := db.GetValidColumns[T, PT]()
	for k := range qo {
		if !validColumns[k] {
			return user, fmt.Errorf("invalid filter name: %s", k)
		}
	}

	query := fmt.Sprintf(getUsersQuery, strings.Join(columnsNames, ", "))

	values := []interface{}{}

	whereClauses, whereValues := db.BuildWhereClauses(qo, len(values))
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
		values = append(values, whereValues...)
	}

	user, err := db.FindRow[T, PT](ur.ctx, query, values...)
	if err != nil {
		// TODO: add persistent logger for db errors
		return user, err
	}

	return user, nil
}

func (ur *UserRepo[T, PT]) GetUsers(qo db.QueryOptions) ([]T, error) {
	var users []T

	validColumns, columnsNames := db.GetValidColumns[T, PT]()
	for k := range qo.Filter {
		if !validColumns[k] {
			return users, fmt.Errorf("invalid filter name: %s", k)
		}
	}

	query := fmt.Sprintf(getUsersQuery, strings.Join(columnsNames, ", "))

	values := []interface{}{}

	whereClauses, whereValues := db.BuildWhereClauses(qo.Filter, len(values))
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
		values = append(values, whereValues...)
	}

	query += qo.CreateQueryFromOptions()

	users, err := db.FindRows[T, PT](ur.ctx, query, values...)
	if err != nil {
		// TODO: add persistent logger for db errors
		return users, err
	}

	return users, nil
}

func (ur *UserRepo[T, PT]) DeleteUser(qf db.QueryFilter) (int, error) {
	validColumns, _ := db.GetValidColumns[T, PT]()
	for k := range qf {
		if !validColumns[k] {
			return 0, fmt.Errorf("invalid filter name: %s", k)
		}
	}

	query := deleteUsersQuery
	values := []interface{}{}

	whereClauses, whereValues := db.BuildWhereClauses(qf, len(values))
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
		values = append(values, whereValues...)
	}

	rnum, err := db.Delete(ur.ctx, query, values...)
	if err != nil {
		// TODO: add persistent logger for db errors
		return int(rnum), err
	}

	return int(rnum), nil
}
