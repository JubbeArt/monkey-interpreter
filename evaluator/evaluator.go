package evaluator

import (
	"fmt"

	"../ast"
	"../object"
	"../tokens"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Statements
	case ast.Program:
		return evalProgram(node.Body, env)
	case ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case ast.AssignmentStatement:
		value := Eval(node.Value, env)
		if isError(value) {
			return value
		}
		env.Set(node.Name, value)
		return nil
	case ast.IfStatement:
		return evalIfStatement(node, env)
	case ast.BlockStatement:
		return evalBlockStatements(node.Statements, env)
	case ast.ReturnStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		return object.ReturnValue{Value: val}

	// Expressions
	case ast.IdentifierExpression:
		return evalIdentifier(node, env)
	case ast.InfixExpression:
		left := Eval(node.LeftSide, env)
		if isError(left) {
			return left
		}
		right := Eval(node.RightSide, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case ast.PrefixExpression:
		right := Eval(node.RightSide, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case ast.NumberExpression:
		return object.Number(node.Value)
	case ast.TextExpression:
		return object.String(node.Value)
	case ast.BooleanExpression:
		return object.Boolean(node.Value)
	case ast.FunctionExpression:
		return object.Function{Parameters: node.Parameters, Body: &node.Body, Env: env}

	case ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args, err := evalExpressions(node.Arguments, env)

		if err != nil {
			return err
		}
		return applyFunction(function, args)
	default:
		fmt.Printf("unhandled ast node: %T\n", node)
	}

	return nil
}

func evalIdentifier(identifier ast.IdentifierExpression, env *object.Environment) object.Object {
	if value, ok := env.Get(identifier.Name); ok {
		return value
	}

	if builtin, ok := builtins[identifier.Name]; ok {
		return builtin
	}

	return newErrorF("could not find identifier %q", identifier.Name)
}

func evalIfStatement(node ast.IfStatement, env *object.Environment) object.Object {
	condition := Eval(node.Conditions, env)
	if isError(condition) {
		return condition
	}

	if condition.Bool() {
		Eval(node.Consequences, env)
	}

	// TODO: Loop
	//else if node.ElseBlock != nil {
	//	return Eval(node.ElseBlock, env)
	//} else {
	//	return NULL
	//}
	return nil
}

func evalExpressions(expressions []ast.Expression, env *object.Environment) ([]object.Object, object.Object) {
	var result []object.Object

	for _, exp := range expressions {
		value := Eval(exp, env)

		if isError(value) {
			return nil, value
		}

		result = append(result, value)
	}

	return result, nil
}

func evalProgram(statements []ast.Statement, env *object.Environment) object.Object {
	for _, statement := range statements {
		result := Eval(statement, env)

		switch result := result.(type) {
		case object.ReturnValue:
			return result.Value
		case object.Error:
			return result
		}
	}

	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case object.Function:
		extendedEnv := extendedFunctionEnv(fn, args)
		return unwrapReturnValue(Eval(fn.Body, extendedEnv))
	case object.BuiltinFunction:
		return fn(args...)
	default:
		return newErrorF("cannot call type %s as a function", fn.Type())
	}
}

func extendedFunctionEnv(function object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(function.Env)

	for i, param := range function.Parameters {
		env.Set(param, args[i])
	}

	return env
}

func unwrapReturnValue(value object.Object) object.Object {
	if returnValue, ok := value.(object.ReturnValue); ok {
		return returnValue.Value
	}

	return value
}

func evalBlockStatements(statements []ast.Statement, env *object.Environment) object.Object {
	for _, statement := range statements {
		result := Eval(statement, env)

		if result == nil {
			continue
		}

		if result.Type() == object.RETURN || result.Type() == object.ERROR {
			return result
		}
	}

	return nil
}

func newErrorF(format string, args ...interface{}) object.Error {
	return object.Error(fmt.Sprintf(format, args...))
}

func isError(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ERROR
}

func evalInfixExpression(operator tokens.TokenType, left object.Object, right object.Object) object.Object {
	switch operator {
	case tokens.EQ:
		return left.Equal(right)
	case tokens.NOT_EQ:
		return !left.Equal(right)
	case tokens.AND:
		if left.Bool() {
			return right
		} else {
			return left
		}
	case tokens.OR:
		if left.Bool() {
			return left
		} else {
			return right
		}
	}

	switch {
	case left.Type() == object.NUMBER && right.Type() == object.NUMBER:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.STRING && right.Type() == object.STRING:
		return evalStringInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newErrorF("type mismatch %v %v %v", left.Type(), operator, right.Type())
	default:
		return newErrorF("bad operator for type %v %v %v", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator tokens.TokenType, left object.Object, right object.Object) object.Object {
	leftVal := left.(object.Number)
	rightVal := right.(object.Number)

	switch operator {
	case tokens.ADD:
		return leftVal.Add(rightVal)
	case tokens.SUB:
		return leftVal.Sub(rightVal)
	case tokens.MUL:
		return leftVal.Mul(rightVal)
	case tokens.DIV:
		return leftVal.Div(rightVal)
	case tokens.LESS:
		return leftVal.Less(rightVal)
	case tokens.LESS_EQ:
		return leftVal.LessEq(rightVal)
	case tokens.GREATER:
		return leftVal.Greater(rightVal)
	case tokens.GREATER_EQ:
		return leftVal.GreaterEq(rightVal)
	default:
		return newErrorF("invalid operator for numbers: %v", operator)
	}
}

func evalStringInfixExpression(operator tokens.TokenType, left object.Object, right object.Object) object.Object {
	leftVal := left.(object.String)
	rightVal := right.(object.String)

	switch operator {
	case tokens.ADD:
		return leftVal.Add(rightVal)
	default:
		return newErrorF("invalid operator for text: %v", operator)
	}
}

func evalPrefixExpression(operator tokens.TokenType, right object.Object) object.Object {
	switch operator {
	case tokens.NOT:
		return object.Boolean(!right.Bool())
	case tokens.SUB:
		if right.Type() == object.NUMBER {
			return right.(object.Number).Negate()
		}
	}
	return newErrorF("invalid operator, %q, for type: %v", operator, right.Type())
}
