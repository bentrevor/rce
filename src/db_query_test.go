package rce_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/bentrevor/rce/src"

	_ "github.com/lib/pq"
)

func TestDB_CanBuildQuery(t *testing.T) {
	institution := HedgeFund{Name: "hedge fund name", Dollars: 100}
	statement := CreateInstitutionStatement(institution)

	assert(t, strings.Index(statement, "insert into hedge_funds") != -1, fmt.Sprintf("should have found insert clause in '%s'", statement))
	assert(t, strings.Index(statement, "(name,dollars) values ('hedge fund name',100)") != -1, fmt.Sprintf("should have found values clause in '%s'", statement))
}
