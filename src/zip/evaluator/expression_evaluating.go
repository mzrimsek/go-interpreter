package evaluator

import (
	"math"
	"strings"

	"github.com/mzrimsek/zip-lang/src/zip/ast"
	"github.com/mzrimsek/zip-lang/src/zip/object"
)

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	case "++":
		return evalIncrementPrefixOperatorExpression(right)
	case "--":
		return evalDecrementPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.INTEGER_OBJ:
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}
	case object.FLOAT_OBJ:
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}
}

func evalIncrementPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.INTEGER_OBJ:
		rightObj := right.(*object.Integer)
		rightObj.Value = rightObj.Value + 1
		return &object.Integer{Value: rightObj.Value}
	case object.FLOAT_OBJ:
		rightObj := right.(*object.Float)
		rightObj.Value = rightObj.Value + 1
		return &object.Float{Value: rightObj.Value}
	default:
		return newError("unknown operator: ++%s", right.Type())
	}
}

func evalDecrementPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.INTEGER_OBJ:
		rightObj := right.(*object.Integer)
		rightObj.Value = rightObj.Value - 1
		return &object.Integer{Value: rightObj.Value}
	case object.FLOAT_OBJ:
		rightObj := right.(*object.Float)
		rightObj.Value = rightObj.Value - 1
		return &object.Float{Value: rightObj.Value}
	default:
		return newError("unknown operator: --%s", right.Type())
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	_, leftIsNum := left.(object.Number)
	_, rightIsNum := right.(object.Number)
	hasNumArg := leftIsNum || rightIsNum

	_, leftIsChar := left.(*object.Character)
	_, rightIsChar := right.(*object.Character)
	hasCharArg := leftIsChar || rightIsChar

	switch {
	case leftIsNum && rightIsNum:
		return evalNumberInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ || right.Type() == object.STRING_OBJ) && (hasNumArg || hasCharArg):
		return evalMixedTypeInfixExpression(operator, left, right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalNumberInfixExpression(operator string, left, right object.Object) object.Object {
	var leftVal float64
	var rightVal float64
	isInt := left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ

	if left.Type() == object.INTEGER_OBJ {
		leftVal = float64(left.(*object.Integer).Value)
	} else {
		leftVal = left.(*object.Float).Value
	}

	if right.Type() == object.INTEGER_OBJ {
		rightVal = float64(right.(*object.Integer).Value)
	} else {
		rightVal = right.(*object.Float).Value
	}

	switch operator {
	case "+":
		val := leftVal + rightVal
		if isInt {
			return &object.Integer{Value: int64(val)}
		}
		return &object.Float{Value: val}
	case "-":
		val := leftVal - rightVal
		if isInt {
			return &object.Integer{Value: int64(val)}
		}
		return &object.Float{Value: val}
	case "*":
		val := leftVal * rightVal
		if isInt {
			return &object.Integer{Value: int64(val)}
		}
		return &object.Float{Value: val}
	case "/":
		val := leftVal / rightVal
		if isInt {
			return &object.Integer{Value: int64(val)}
		}
		return &object.Float{Value: val}
	case "%":
		if isInt {
			return &object.Integer{Value: int64(leftVal) % int64(rightVal)}
		}
		return &object.Float{Value: math.Mod(leftVal, rightVal)}
	case "**":
		val := math.Pow(leftVal, rightVal)
		if isInt {
			return &object.Integer{Value: int64(val)}
		}
		return &object.Float{Value: val}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalMixedTypeInfixExpression(operator string, left, right object.Object) object.Object {
	switch operator {
	case "+":
		return &object.String{Value: left.Inspect() + right.Inspect()}
	case "*":
		if left.Type() == object.INTEGER_OBJ {
			integer := left.(*object.Integer).Value
			return &object.String{Value: strings.Repeat(right.Inspect(), int(integer))}
		}
		if right.Type() == object.INTEGER_OBJ {
			integer := right.(*object.Integer).Value
			return &object.String{Value: strings.Repeat(left.Inspect(), int(integer))}
		}
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value

	switch operator {
	case "&&":
		return nativeBoolToBooleanObject(leftVal && rightVal)
	case "||":
		return nativeBoolToBooleanObject(leftVal || rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(left, index object.Object) object.Object {
	arrayObject := left.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.Object {
	condition := Eval(we.Condition, env)
	if isError(condition) {
		return condition
	}

	var result object.Object
	for isTruthy(condition) {
		result = Eval(we.Block, env)

		condition = Eval(we.Condition, env)
		if isError(condition) {
			return condition
		}
	}
	return result
}

func evalPostfixExpression(left object.Object, operator string) object.Object {
	switch operator {
	case "++":
		return evalIncrementPostfixOperatorExpression(left)
	case "--":
		return evalDecrementPostfixOperatorExpression(left)
	default:
		return newError("unknown operator: %s%s", left.Type(), operator)
	}
}

func evalIncrementPostfixOperatorExpression(left object.Object) object.Object {
	switch left.Type() {
	case object.INTEGER_OBJ:
		leftObj := left.(*object.Integer)
		returnVal := &object.Integer{Value: leftObj.Value}
		leftObj.Value = leftObj.Value + 1
		return returnVal
	case object.FLOAT_OBJ:
		leftObj := left.(*object.Float)
		returnVal := &object.Float{Value: leftObj.Value}
		leftObj.Value = leftObj.Value + 1
		return returnVal
	default:
		return newError("unknown operator: %s++", left.Type())
	}
}

func evalDecrementPostfixOperatorExpression(left object.Object) object.Object {
	switch left.Type() {
	case object.INTEGER_OBJ:
		leftObj := left.(*object.Integer)
		returnVal := &object.Integer{Value: leftObj.Value}
		leftObj.Value = leftObj.Value - 1
		return returnVal
	case object.FLOAT_OBJ:
		leftObj := left.(*object.Float)
		returnVal := &object.Float{Value: leftObj.Value}
		leftObj.Value = leftObj.Value - 1
		return returnVal
	default:
		return newError("unknown operator: %s--", left.Type())
	}
}
