package query

import (
	"fmt"
	"strings"
	"time"
)

type Select []string
type Where map[interface{}]interface{}
type Limit int

type Model struct {
	_connection string
	_select     Select
	_where      Where
	_limit      Limit
}

func Test() {

	user := &User{
		Model: &Model{},
		Email: "jeo@doe.com",
	}

	user.Select(
		[]string{"id", "created_at", "email"},
	).Where(Where{
		"email":  "jeo@doe.com",
		"height": Where{">": 5.7},
	}).Limit(50).Get()

	fmt.Println("")
}

type Entity interface{}
type Query map[string]interface{}

type Pagination struct {
	Total   int64    `json:"total"`
	Page    int64    `json:"page"`
	PerPage int      `json:"per_page"`
	Items   []Entity `json:"items"`
	Query   Query
}

type Database interface {
	Connection(name string) Database
	Select(name string) Database
	Where(name string) Database
	Limit(name string) Database
	Get() []Entity
	Pagination(name string) *Pagination
}

type User struct {
	*Model
	ID        string    `orm_name:"id" orm_type:"primary_key"`
	CreatedAt time.Time `orm_name:"created_at" orm_type:"datetime"`
	Email     string    `orm_name:"email" orm_type:"string"`
	Height    float32   `orm_name:"height" orm_type:"float"`
	Password  string    `orm_name:"password" orm_type:"string"`
}

// Comment
func (ctx *Model) Connection(name string) *Model {
	return ctx
}

// Comment
func (ctx *Model) Select(s Select) *Model {
	ctx._select = s

	return ctx
}

// Comment
func (ctx *Model) Where(w Where) *Model {
	ctx._where = w

	return ctx
}

// Comment
func (ctx *Model) Limit(l Limit) *Model {
	ctx._limit = l

	return ctx
}

type QueryBuilder struct {
	*Model
	values []interface{}
}

const (
	SPACE = "  "
)

// Comment
func (ctx *QueryBuilder) SelectStatement() (string, error) {
	if len(ctx._select) == 0 {
		ctx._select = []string{"*"}
	}

	return strings.Join([]string{
		"SELECT", SPACE + strings.Join(ctx._select, ", "), "FROM",
	}, "\r\n"), nil
}

func (ctx QueryBuilder) whereStatementBuilder() (string, error) {
	// where := []string{}

	return "", nil
}

// Comment
func (ctx *QueryBuilder) WhereStatement() (string, error) {
	where := []string{}

	for key, value := range ctx._where {

		switch value.(type) {
		case int, int8, int16, int32, int64, string, float32, float64:
			k, ok := key.(string)

			if !ok {
				return "", fmt.Errorf("Where key (%v) is not type of string", key)
			}

			where = append(where, strings.Join([]string{k, "?"}, "="))

			ctx.values = append(ctx.values, value)
			break

		case Where:

			break

		case func() Where:

			break

		default:
			return "", nil //fmt.Errorf("Error where statement")
		}

	}

	return strings.Join(where, "\r\n"), nil
}

// Comment
func (ctx *QueryBuilder) Build() (string, error) {
	query := []string{}

	_select, err := ctx.SelectStatement()

	if err != nil {
		return "", err
	}

	query = append(query, _select)
	query = append(query, SPACE+"users")

	_where, err := ctx.WhereStatement()

	if err != nil {
		return "", err
	}

	query = append(query, _where)

	return strings.Join(query, "\r\n"), nil
}

// Comment
func (ctx *Model) Get() (interface{}, error) {
	builder := &QueryBuilder{Model: ctx}

	sql, err := builder.Build()

	if err != nil {
		return nil, err
	}

	fmt.Printf("QUERY\r\n\r\n\r\n%s\r\n\r\n", sql)

	return nil, nil
}

type Values map[string]interface{}

// Comment
func (ctx *Model) Create(values Values) (*Model, error) {

	return nil, nil
}
