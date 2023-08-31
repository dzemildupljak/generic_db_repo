package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/dzemildupljak/web_app_with_unitest/generic/common"
	"github.com/dzemildupljak/web_app_with_unitest/generic/db"
)

const userTable = "public.users"
const getUsersQuery = "SELECT %s FROM " + userTable

type urepoType interface {
	User | SignInUser | PartialUser
}

type UserRepo[T urepoType, PT common.Ptr[T]] struct {
}

func NewUserRepo[T urepoType, PT common.Ptr[T]]() UserRepo[T, PT] {
	ur := UserRepo[T, PT]{}

	return ur
}

func (ur *UserRepo[T, PT]) GetUser(qo db.QueryOptions) (T, error) {
	var user T
	ctx := context.Background()

	validColumns, _ := db.GetValidColumns[T, PT]()
	_, columnsNames := db.GetValidColumns[T, PT]()
	for k := range qo.Filter {
		if !validColumns[k] {
			return user, fmt.Errorf("invalid filter name: %s", k)
		}
	}

	query := fmt.Sprintf(getUsersQuery, strings.Join(columnsNames, ", "))

	values := []interface{}{}

	whereClauses, whereValues := db.BuildWhereClauses(qo.Filter, len(values))
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
		values = append(values, whereValues...)
	}

	fmt.Println(query)
	fmt.Println(values...)

	user, err := db.GetRow[T, PT](ctx, query, values...)
	if err != nil {
		// TODO: add persistent logger for db errors
		return user, err
	}

	return user, nil
}

func (ur *UserRepo[T, PT]) GetUsers(qo db.QueryOptions) ([]T, error) {
	var users []T
	ctx := context.Background()

	validColumns, _ := db.GetValidColumns[T, PT]()
	_, columnsNames := db.GetValidColumns[T, PT]()
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

	fmt.Println(query)

	users, err := db.FindRows[T, PT](ctx, query, values...)
	if err != nil {
		// TODO: add persistent logger for db errors
		return users, err
	}

	return users, nil
}
