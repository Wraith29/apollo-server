package query

import (
	"fmt"
	"strings"
)

type F struct {
	table, column string
}

func (f *F) string() string {
	if f.table != "" {
		return fmt.Sprintf(`%s."%s"`, f.table, f.column)
	}

	return fmt.Sprintf(`"%s"`, f.column)
}

type queryBuilder struct {
	buf strings.Builder
}

func NewQueryBuilder() *queryBuilder {
	return &queryBuilder{
		buf: strings.Builder{},
	}
}

func (b *queryBuilder) String() string {
	return b.buf.String()
}

func (b *queryBuilder) Select(fields ...F) *queryBuilder {
	query := strings.Builder{}

	for idx, field := range fields {
		query.WriteString(field.string())

		if idx != len(fields)-1 {
			query.WriteString(", ")
		}
	}

	b.buf.WriteString(fmt.Sprintf("SELECT %s ", query.String()))

	return b
}
