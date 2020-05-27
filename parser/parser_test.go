package parser

import (
	"fmt"
	"testing"

	"../lexer"
)

func TestAssignmentAndBasicTypes(t *testing.T) {
	testParser(t, `
		a = 1
		b = 10.123
		c = "heelo åöö world"
		abc = true
	`, []string{
		"a = 1",
		"b = 10.123",
		"c = \"heelo åöö world\"",
		"abc = true",
	})
}

func TestListType(t *testing.T) {
	testParser(t, `
		a = []
		b = [1]
		c = [1,2,3]
		d = [123.4, "hellö", true]
	`, []string{
		`a = []`,
		`b = [
	1
]`,
		`c = [
	1,
	2,
	3
]`,
		`d = [
	123.4,
	"hellö",
	true
]`,
	})
}

func TestPrefixExpression(t *testing.T) {
	testParser(t, `
		a =   a
		b =   1
		c =   true
		d =   "why"
		e = not    me
		f = - 10.13
		g = (23)
		h = func (a,b,c)
			i = 123
		end
	`, []string{
		"a = a",
		"b = 1",
		"c = true",
		`d = "why"`,
		"e = not me",
		"f = -10.13",
		"g = 23",
		`h = func (a, b, c)
	i = 123
end`,
	})
}

func TestInfixExpressions(t *testing.T) {
	testParser(t, `
		a = 2 + 5 * 10
		b = (  2 + 5  )*10
		c +=  123 / 3
		d *= -  3
		e = 123 --- 5
		f = true == 123 * 3 * 5
		g = 123.5 > 234
		h = 123.5 >= 234
		i = 123.5 <= 234
		j = 123.5 < 234
		k = "123" != 123
		l = asdf()
		m = print(1,"asdf", [], [1,2,3], false)
		n = true and 123 + 5
		o = 1 or 2 and 3
	`, []string{
		"a = (2 + (5 * 10))",
		"b = ((2 + 5) * 10)",
		"c += (123 / 3)",
		"d *= -3",
		"e = (123 - --5)",
		"f = (true == ((123 * 3) * 5))",
		"g = (123.5 > 234)",
		"h = (123.5 >= 234)",
		"i = (123.5 <= 234)",
		"j = (123.5 < 234)",
		`k = ("123" != 123)`,
		"l = asdf()",
		"m = print(1, \"asdf\", [], [\n\t1,\n\t2,\n\t3\n], false)",
		"n = (true and (123 + 5))",
		"o = (1 or (2 and 3))",
	})
}

func TestIfStatement(t *testing.T) {
	testParser(t, `
		if true then end

		if 1 + 2 == 3  then 
		  a = true
			b = 10 +5
		end

		if a then
			a = 1
		else
			a = 2
		end

		if b then
			a = 1
		elseif false then
			a = 2
		end

		if c then
			d = 1
		elseif false then
			d = 2
		else
			d = 3
		end

	`, []string{
		`if true then
end`,
		`if ((1 + 2) == 3) then
	a = true
	b = (10 + 5)
end`,
		`if a then
	a = 1
if true then
	a = 2
end`,
		`if b then
	a = 1
if false then
	a = 2
end`,
		`if c then
	d = 1
if false then
	d = 2
if true then
	d = 3
end`,
	})
}

func TestLoopStatement(t *testing.T) {
	testParser(t, `
		loop end
	
		loop
			a = 10
			if true then
				a = 2
			end
		end
	`, []string{
		"loop\nend",
		`loop
	a = 10
	if true then
		a = 2
	end
end`,
	})
}

func TestReturnStatement(t *testing.T) {
	testParser(t, `
		return "hello"
		return 5 + 5
		return (1 + 5) * 10 + 3
		
		if true then
			asdf = 123
			return
		end

		a = func (a,b,c)
			return
		end

		if asdf then
			return
		elseif false then
			return
		else
			return 123
		end

		`, []string{
		`return "hello"`,
		`return (5 + 5)`,
		`return (((1 + 5) * 10) + 3)`,
		`if true then
	asdf = 123
	return
end`,
		`func (a, b, c)
	return
end
`,
		`if asdf then
	return
if false then
	return
if true then
	return 123
end`,
	})
}

func testParser(t *testing.T, input string, expected []string) {
	pars := New(lexer.New(input))
	program := pars.ParseProgram()

	if pars.HasErrors() {
		t.Fatalf("parser found an error:" + pars.errors[0])
	}

	if len(program.Body.Statements) != len(expected) {
		for i, stmt := range program.Body.Statements {
			fmt.Println(i)
			fmt.Println(stmt.String(0))
		}

		t.Fatalf("expected %v statements. got=%d", len(expected), len(program.Body.Statements))
	}

	for i, stmt := range program.Body.Statements {
		if stmt.String(0) != expected[i] {
			t.Fatalf("expected statement to have string:\n%q\ninstead we got:\n%q", expected[i], stmt.String(0))
		}
	}
}
