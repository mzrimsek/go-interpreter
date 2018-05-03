package evaluator

import (
	"zip/lexer"
	"zip/object"
	"zip/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"3 % 2", 1},
		{"10 % 4 * 2 + 6", 10},
		{"++5", 6},
		{"1 + ++(3 / 3 + 1)", 4},
		{"let a = 1; let b = ++a; b;", 2},
		{"--5", 4},
		{"1 + --(3 / 3 + 1)", 2},
		{"let a = 1; let b = --a; b;", 0},
		{"2 ** 2", 4},
		{"2 ** 5", 32},
		{"5++", 5},
		{"let a = 2; let b = a++; a;", 3},
		{"5--", 5},
		{"let a = 2; let b = a--; a;", 1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.5", 5.5},
		{"10.3", 10.3},
		{"-5.5", -5.5},
		{"-10.3", -10.3},
		{"5 + 5.5 + 5 + 5.5 - 10.5", 10.5},
		{"2 * 2.5 * 2 * 2.5 * 0.5", 12.5},
		{"-5.5 + 0.3 + -5.5", -10.7},
		{"5.2 * 2 + 10", 20.4},
		{"5.5 + 2 * 10", 25.5},
		{"20 + 2.5 * -10", -5.0},
		{"50.0 / 4 * 2 + 10.5", 35.5},
		{"2 * (5.3 + 10)", 30.6},
		{"3 * 3 * 3 + 10.5", 37.5},
		{"3 * (3.5 * 3) + 10", 41.5},
		{"(5 + 10 * 2 + 15.0 / 2) * 2 + -10", 55.0},
		{"3.5 % 2", 1.5},
		{".5 * 0.5", .25},
		{"++2.5", 3.5},
		{"let a = 1.5; let b = ++a; b;", 2.5},
		{"--2.5", 1.5},
		{"let a = 1.5; let b = --a; b;", 0.5},
		{"4 ** 0.5", 2},
		{"2.5++", 2.5},
		{"let a = 2.5; let b = a++; a;", 3.5},
		{"2.5--", 2.5},
		{"let a = 2.5; let b = a--; a;", 1.5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 > 2 && 2 == 2", false},
		{"1 > 2 || 2 == 2", true},
		{"1 <= 2", true},
		{"1 >= 2", false},
		{"1 <= 1", true},
		{"1 >= 1", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{`"hello" == "hello"`, true},
		{`"hello" == "goodbye"`, false},
		{`"hello" != "hello"`, false},
		{`"hello" != "goodbye"`, true},
		{"true && true", true},
		{"true && false", false},
		{"false && true", false},
		{"false && false", false},
		{"true || true", true},
		{"true || false", true},
		{"false || true", true},
		{"false || false", false},
		{"1.5 < 2", true},
		{"1.5 > 2", false},
		{"1.5 < 1.5", false},
		{"1.5 > 1.5", false},
		{"1.5 == 1.5", true},
		{"1.5 != 1.5", false},
		{"1.5 == 2", false},
		{"1.5 != 2", true},
		{"1.5 > 2 && 2 == 2", false},
		{"1.5 > 2 || 2 == 2", true},
		{"1.5 <= 2", true},
		{"1.5 >= 2", false},
		{"1.5 <= 1.5", true},
		{"1.5 >= 1.5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{
			`if (10 > 1) { 
				if (10 > 1) { 
					return 10; 
				} 
				return 1; 
			 }`, 10},
	}

	for _, tt := range tests {
		evalutated := testEval(tt.input)
		testIntegerObject(t, evalutated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5 + true;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + true; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; true + false; 5", "unknown operator: BOOLEAN + BOOLEAN"},
		{"if (10 > 1) { true + false; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{
			`if (10 > 1) { 
				if (10 > 1) { 
					return true + false; 
				} 
				return 1;
			}`, "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "identifier not found: foobar"},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
		{`{"name": "Monkey"}[fn(x) { x }];`, "unusable as hash key: FUNCTION"},
		{"2 && false", "type mismatch: INTEGER && BOOLEAN"},
		{`true && "hello"`, "type mismatch: BOOLEAN && STRING"},
		{"2 || false", "type mismatch: INTEGER || BOOLEAN"},
		{`true || "hello"`, "type mismatch: BOOLEAN || STRING"},
		{`"hello" + false`, "type mismatch: STRING + BOOLEAN"},
		{`"hello" - 3`, "unknown operator: STRING - INTEGER"},
		{`++"hello"`, "unknown operator: ++STRING"},
		{`--"hello"`, "unknown operator: --STRING"},
		{"++false", "unknown operator: ++BOOLEAN"},
		{`3.5 * "hello"`, "unknown operator: FLOAT * STRING"},
		{`"hello" * 3.5`, "unknown operator: STRING * FLOAT"},
		{`"hello" ** 4`, "unknown operator: STRING ** INTEGER"},
		{`"hello"++`, "unknown operator: STRING++"},
		{`"hello"--`, "unknown operator: STRING--"},
		{"false++", "unknown operator: BOOLEAN++"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T (%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `let newAdder = fn(x) {
				fn(y) { x + y };
			};
			
			let addTwo = newAdder(2);
			addTwo(2);`

	testIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Fatalf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"Hello" + " " + "World!"`, "Hello World!"},
		{`"Hello" + 3`, "Hello3"},
		{`3 + "Hello"`, "3Hello"},
		{`"Hello" + 3.5`, "Hello3.5"},
		{`3.5 + "Hello"`, "3.5Hello"},
		{`(12 * .5) + "Hello"`, "6Hello"},
		{`6.0 + "Hello"`, "6Hello"},
		{`"Hello" * 3`, "HelloHelloHello"},
		{`3 * "Hello"`, "HelloHelloHello"},
		{`'H' + "ello"`, "Hello"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
		}

		if str.Value != tt.expected {
			t.Errorf("String has wrong value. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to 'len' not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{"len([])", 0},
		{"len([1 , 2])", 2},
		{`let myArray = [1, 2, 3]; len(myArray);`, 3},
		{"len([1, 2, 3], [4, 5, 6])", "wrong number of arguments. got=2, want=1"},
		{"first([])", NULL},
		{"first([1, 2])", 1},
		{`let myArray = [1, 2, 4]; first(myArray);`, 1},
		{`first(["one", "two"])`, "one"},
		{"first([1, 2], [3, 4])", "wrong number of arguments. got=2, want=1"},
		{"first(1)", "argument to 'first' must be ARRAY, got INTEGER"},
		{"last([])", NULL},
		{"last([1, 2])", 2},
		{`let myArray = [1, 2, 4]; last(myArray);`, 4},
		{`last(["one", "two"])`, "two"},
		{"last([1, 2], [3, 4])", "wrong number of arguments. got=2, want=1"},
		{"last(1)", "argument to 'last' must be ARRAY, got INTEGER"},
		{"tail([])", NULL},
		{"tail([1, 2, 3])", []int{2, 3}},
		{"tail([1, 2], [3, 4])", "wrong number of arguments. got=2, want=1"},
		{"tail(1)", "argument to 'tail' must be ARRAY, got INTEGER"},
		{"push([], 1)", []int{1}},
		{"push([1, 2], 3)", []int{1, 2, 3}},
		{"push([1, 2])", "wrong number of arguments. got=1, want=2"},
		{"push(1, 1)", "argument to 'push' must be ARRAY, got INTEGER"},
		{"type(1)", "INTEGER"},
		{"type(1 - 5)", "INTEGER"},
		{"type([])", "ARRAY"},
		{`type("hello")`, "STRING"},
		{`substr("hello", 0, 2)`, "he"},
		{`substr("hello", 0, 5)`, "hello"},
		{"substr(1, 0, 2)", "first argument to 'substr' must be STRING, got INTEGER"},
		{`substr("hello", 1)`, "wrong number of arguments. got=2, want=3"},
		{`substr("hello", 1, 6)`, "substring indices must be within string length"},
		{`substr("hello", -1, 5)`, "substring indices must be within string length"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				testStringObject(t, evaluated, string(expected))
				continue
			}

			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		case []int:
			testArrayObject(t, evaluated, expected)
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := `[1, 2 * 2, 3 + 3, "hello"]`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 4 {
		t.Fatalf("array has not enough elements. want 4, got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
	testStringObject(t, result.Elements[3], "hello")
}

func TestHashLiterals(t *testing.T) {
	input := `let two = "two";
			  {
				  "one": 10 - 9,
				  two: 1 + 1,
				  "thr" + "ee": 6 / 2,
				  4: 4,
				  true: 5,
				  false: 6,
				  'a': 7
			  }`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("object is not Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
		(&object.Character{Value: 'a'}).HashKey():  7,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong number of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue int64
	}{
		{"let a = 5; a = 6; a;", 6},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expectedValue)
	}
}

func TestShortcutAssignStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue int64
	}{
		{"let a = 5; a += 2; a;", 7},
		{"let a = 5; a -= 2; a;", 3},
		{"let a = 5; a *= 2; a;", 10},
		{"let a = 6; a /= 2; a;", 3},
		{"let a = 5; a %= 2; a;", 1},
		{"let a = 2; a **= 4; a;", 16},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expectedValue)
	}
}

func TestCharacterLiteral(t *testing.T) {
	input := "'a'"

	evaluated := testEval(input)
	char, ok := evaluated.(*object.Character)
	if !ok {
		t.Fatalf("object is not Character. got=%T (%+v)", evaluated, evaluated)
	}

	if char.Value != 'a' {
		t.Fatalf("Character has wrong value. got=%q", char.Value)
	}
}

func testEval(input string) object.Object {
	ShouldPrint = true

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T, (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q", result.Value, expected)
		return false
	}

	return true
}

func testArrayObject(t *testing.T, obj object.Object, expected []int) bool {
	result, ok := obj.(*object.Array)
	if !ok {
		t.Errorf("object. is not Array. got=%T, (%+v)", obj, obj)
		return false
	}

	for index, val := range result.Elements {
		integer, ok := val.(*object.Integer)
		if !ok {
			t.Errorf("array value not Integer. got=%T (%+v)", val, val)
			return false
		}

		if integer.Value != int64(expected[index]) {
			t.Errorf("array has wrong value at index. got=%d, want=%d", integer.Value, int64(expected[index]))
			return false
		}
	}

	return true
}
