package maintenance

import (
	"context"
	"fmt"
	"strings"

	"github.com/dzemildupljak/web_app_with_unitest/generic/common"
	"github.com/dzemildupljak/web_app_with_unitest/generic/db"
)

const userTable = "public.user_maintenance_tokens"
const getUsersQuery = "SELECT %s FROM " + userTable

type umtrepoType interface {
	UserMTokens
}

type UmtRepo[T umtrepoType, PT common.Ptr[T]] struct {
}

func NewUmtRepo[T umtrepoType, PT common.Ptr[T]]() UmtRepo[T, PT] {
	ur := UmtRepo[T, PT]{}

	return ur
}

func (ur *UmtRepo[T, PT]) GetUMToken(qo db.QueryOptions) (T, error) {
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
	} else {
		return user, db.ErrIvalidFilters
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
