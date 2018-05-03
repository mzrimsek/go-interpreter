package parser

import (
	"zip/ast"
	"zip/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not ast.ReturnStatement. got=%T", stmt)
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Fatalf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}

		if testLiteralExpression(t, returnStmt.ReturnValue, tt.expectedValue) {
			return
		}
	}
}

func TestAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"x = 5;", "x", 5},
		{"y = true;", "y", true},
		{"foobar = y;", "foobar", "y"},
		{"x = 5.2;", "x", 5.2},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		testAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
	}
}

func TestShortcutAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"x += 5;", "x", 5},
		{"y += 2.5;", "y", 2.5},
		{"z += y;", "z", "y"},
		{"x -= 5;", "x", 5},
		{"y -= 2.5;", "y", 2.5},
		{"z -= y;", "z", "y"},
		{"x *= 5;", "x", 5},
		{"y *= 2.5;", "y", 2.5},
		{"z *= y;", "z", "y"},
		{"x /= 5;", "x", 5},
		{"y /= 2.5;", "y", 2.5},
		{"z /= y;", "z", "y"},
		{"x %= 5;", "x", 5},
		{"y %= 2.5;", "y", 2.5},
		{"z %= y;", "z", "y"},
		{"x **= 5;", "x", 5},
		{"y **= 2.5;", "y", 2.5},
		{"z **= y;", "z", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		testShortcutAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func testAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	assignStmt, ok := stmt.(*ast.AssignStatement)
	if !ok {
		t.Errorf("stmt not ast.AssignStatement. got=%T", stmt)
		return false
	}

	if assignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, assignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, assignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, assignStmt.Value, float64(v))
	case string:
		testIdentifier(t, assignStmt.Value, v)
	case bool:
		testBooleanLiteral(t, assignStmt.Value, v)
	}

	return true
}

func testShortcutAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	shortcutAssignStmt, ok := stmt.(*ast.ShortcutAssignStatement)
	if !ok {
		t.Errorf("stmt not ast.ShortcutAssignStatement. got=%T", stmt)
		return false
	}

	if shortcutAssignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, shortcutAssignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, shortcutAssignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, shortcutAssignStmt.Value, float64(v))
	case string:
		testIdentifier(t, shortcutAssignStmt.Value, v)
	}

	return true
}
