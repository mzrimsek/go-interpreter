package parser

import (
	"monkey/ast"
	"monkey/lexer"
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

func TestAddAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"x += 5;", "x", 5},
		{"y += 2.5;", "y", 2.5},
		{"z += y;", "z", "y"},
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
		testAddAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
	}
}

func TestSubAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"x -= 5;", "x", 5},
		{"y -= 2.5;", "y", 2.5},
		{"z -= y;", "z", "y"},
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
		testSubAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
	}
}

func TestMultAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"x *= 5;", "x", 5},
		{"y *= 2.5;", "y", 2.5},
		{"z *= y;", "z", "y"},
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
		testMultAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
	}
}

func TestDivAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"x /= 5;", "x", 5},
		{"y /= 2.5;", "y", 2.5},
		{"z /= y;", "z", "y"},
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
		testDivAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
	}
}

func TestModAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"x %= 5;", "x", 5},
		{"y %= 2.5;", "y", 2.5},
		{"z %= y;", "z", "y"},
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
		testModAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
	}
}

func TestPowAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
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
		testPowAssignStatement(t, stmt, tt.expectedName, tt.expectedValue)
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

func testAddAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	addAssignStmt, ok := stmt.(*ast.AddAssignStatement)
	if !ok {
		t.Errorf("stmt not ast.AddAssignStatement. got=%T", stmt)
		return false
	}

	if addAssignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, addAssignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, addAssignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, addAssignStmt.Value, float64(v))
	case string:
		testIdentifier(t, addAssignStmt.Value, v)
	}

	return true
}

func testSubAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	subAssignStmt, ok := stmt.(*ast.SubAssignStatement)
	if !ok {
		t.Errorf("stmt not ast.SubAssignStatement. got=%T", stmt)
		return false
	}

	if subAssignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, subAssignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, subAssignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, subAssignStmt.Value, float64(v))
	case string:
		testIdentifier(t, subAssignStmt.Value, v)
	}

	return true
}

func testMultAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	multAssignStmt, ok := stmt.(*ast.MultAssignStatement)
	if !ok {
		t.Errorf("stmt not ast.MultAssignStatement. got=%T", stmt)
		return false
	}

	if multAssignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, multAssignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, multAssignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, multAssignStmt.Value, float64(v))
	case string:
		testIdentifier(t, multAssignStmt.Value, v)
	}

	return true
}

func testDivAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	divAssignStmt, ok := stmt.(*ast.DivAssignStatement)
	if !ok {
		t.Errorf("stmt not ast.DivAssignStatement. got=%T", stmt)
		return false
	}

	if divAssignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, divAssignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, divAssignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, divAssignStmt.Value, float64(v))
	case string:
		testIdentifier(t, divAssignStmt.Value, v)
	}

	return true
}

func testModAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	modAssignStmt, ok := stmt.(*ast.ModAssignStatement)
	if !ok {
		t.Errorf("stmt not ast.ModAssignStatement. got=%T", stmt)
		return false
	}

	if modAssignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, modAssignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, modAssignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, modAssignStmt.Value, float64(v))
	case string:
		testIdentifier(t, modAssignStmt.Value, v)
	}

	return true
}

func testPowAssignStatement(t *testing.T, stmt ast.Statement, name string, value interface{}) bool {
	powAssignStmt, ok := stmt.(*ast.PowAssignStatement)
	if !ok {
		t.Errorf("stmt not ast.PowAssignStatement. got=%T", stmt)
		return false
	}

	if powAssignStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Target not '%s'. got=%s", name, powAssignStmt.Name)
		return false
	}

	switch v := value.(type) {
	case int64:
		testIntegerLiteral(t, powAssignStmt.Value, int64(v))
	case float64:
		testFloatLiteral(t, powAssignStmt.Value, float64(v))
	case string:
		testIdentifier(t, powAssignStmt.Value, v)
	}

	return true
}
