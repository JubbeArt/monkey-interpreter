package ast

type Node interface {
	String(indent int) string
	//StartPos() tokens.Pos
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

const INDENT = "\t"

// -------------------------------------------
// ---------------- PROGRAM ------------------
// -------------------------------------------
type Program struct {
	Body BlockStatement
}

func (p Program) String(indent int) string {
	return p.Body.String(indent)
}
