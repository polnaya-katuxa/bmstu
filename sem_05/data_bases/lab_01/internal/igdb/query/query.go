package query

import (
	"fmt"
	"strings"
)

type Query struct {
	fields []string
	limit  int
	wheres []Condition
}

type Condition struct {
	field  string
	symbol string
	value  interface{}
}

func New() *Query {
	return &Query{}
}

func (q *Query) Fields(fields ...string) *Query {
	q.fields = append(q.fields, fields...)

	return q
}

func (q *Query) Limit(limit int) *Query {
	q.limit = limit

	return q
}

func (q *Query) Where(conds ...Condition) *Query {
	q.wheres = append(q.wheres, conds...)

	return q
}

func Equal(field string, value interface{}) Condition {
	return Condition{
		field:  field,
		symbol: "=",
		value:  value,
	}
}

func NotEqual(field string, value interface{}) Condition {
	return Condition{
		field:  field,
		symbol: "!=",
		value:  value,
	}
}

func More(field string, value interface{}) Condition {
	return Condition{
		field:  field,
		symbol: ">",
		value:  value,
	}
}

func Less(field string, value interface{}) Condition {
	return Condition{
		field:  field,
		symbol: "<",
		value:  value,
	}
}

func (q *Query) fieldsToString() string {
	if len(q.fields) == 0 {
		return ""
	}

	format := "fields %s;\n"

	return fmt.Sprintf(format, strings.Join(q.fields, ","))
}

func (q *Query) limitToString() string {
	if q.limit == 0 {
		return ""
	}

	format := "limit %d;\n"

	return fmt.Sprintf(format, q.limit)
}

func (c *Condition) String() string {
	format := "%s %s %s"

	strValue := fmt.Sprint(c.value)
	if c.value == nil {
		strValue = "null"
	}

	return fmt.Sprintf(format, c.field, c.symbol, strValue)
}

func (q *Query) wheresToString() string {
	if len(q.wheres) == 0 {
		return ""
	}

	wheresStr := make([]string, len(q.wheres))
	for i, cond := range q.wheres {
		wheresStr[i] = cond.String()
	}

	format := "where %s;\n"

	return fmt.Sprintf(format, strings.Join(wheresStr, " & "))
}

func (q *Query) String() string {
	return q.fieldsToString() + q.wheresToString() + q.limitToString()
}

func IDsToString(ids []int) string {
	str := "("

	for i, id := range ids {
		if i == 0 {
			str += fmt.Sprintf("%d", id)
			continue
		}
		str += fmt.Sprintf(",%d", id)
	}

	str += ")"

	return str
}
