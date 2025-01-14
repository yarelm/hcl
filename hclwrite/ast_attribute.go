package hclwrite

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Attribute struct {
	inTree

	LeadComments *node
	name         *node
	expr         *node
	LineComments *node
}

func newAttribute() *Attribute {
	return &Attribute{
		inTree: newInTree(),
	}
}

func (a *Attribute) init(name string, expr *Expression) {
	expr.assertUnattached()

	nameTok := newIdentToken(name)
	nameObj := newIdentifier(nameTok)
	a.LeadComments = a.children.Append(newComments(nil))
	a.name = a.children.Append(nameObj)
	a.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte{'='},
		},
	})
	a.expr = a.children.Append(expr)
	a.expr.list = a.children
	a.LineComments = a.children.Append(newComments(nil))
	a.children.AppendUnstructuredTokens(Tokens{
		{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte{'\n'},
		},
	})
}

func (a *Attribute) Expr() *Expression {
	return a.expr.content.(*Expression)
}
