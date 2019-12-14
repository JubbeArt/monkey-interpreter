package parser

import (
	"testing"

	"../ast"
	"../lexer"
)

func TestStatements(t *testing.T) {
	input := `
		x = 5
		y = 10.123
		z = "heelo world"

		a = func () 
			b = 234
			return b
		end
	`

	lex := lexer.New(input)
	pars := New(lex)
	program := pars.ParseProgram()

	if pars.HasErrors() {
		t.Fatalf("Parser found some errors:" + pars.errors[0])
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
		expectedValue      string
	}{
		{"x", "5"},
		{"y", "10"},
		{"foobar", "838383"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAssignmentStatement(t, stmt, tt.expectedIdentifier, tt.expectedValue) {
			return
		}
	}
}
func testAssignmentStatement(t *testing.T, s ast.Statement, name string, value string) bool {
	assignStmt, ok := s.(ast.AssignmentStatement)
	if !ok {
		t.Errorf("statment not *ast.LetStatement. got=%T", s)
		return false
	}
	if assignStmt.Name != name {
		t.Errorf("assignStatemnt.Name not '%s'. got=%s", name, assignStmt.Name)
		return false
	}
	if assignStmt.Value.String() != value {
		t.Errorf("assignStatment.value not '%s'. got=%s", value, assignStmt.Value.String())
		return false
	}
	return true
}
