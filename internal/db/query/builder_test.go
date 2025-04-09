package query

import "testing"

func Test_SimpleSelect_WithTableAlias(t *testing.T) {
	b := NewQueryBuilder().
		Select(F{"T", "Col1"}, F{"T", "Col2"})

	expected := `SELECT T."Col1", T."Col2" FROM "MyTable" T`
	query := b.String()

	if query != expected {
		t.Errorf(`Expected "%s" Got "%s"\n`, expected, query)
	}
}
