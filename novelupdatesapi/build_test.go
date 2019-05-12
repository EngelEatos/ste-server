package novelupdatesapi

import "testing"

func TestParseSQLFile(t *testing.T) {
	queries := ParseSQLFile("../ste_files/schema-ste.sql")
	if len(queries) != 50 {
		t.Errorf("expected %d, got %d", 50, len(queries))
	}
}

func TestStart(t *testing.T) {
	Start()

}
