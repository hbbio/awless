package ast

import (
	"reflect"
	"testing"

	"github.com/hbbio/awless/template/env"
	"github.com/hbbio/awless/template/params"
)

func TestCloneAST(t *testing.T) {
	tree := &AST{}

	cmd := new(fakeCmd)

	tree.Statements = append(tree.Statements, &Statement{Node: &DeclarationNode{
		Ident: "myvar",
		Expr: &CommandNode{
			Action: "create", Entity: "vpc",
			ParamNodes: map[string]interface{}{"count": InterfaceNode{i: 1}, "myname": RefNode{key: "name"}},
			Refs:       map[string]interface{}{"here": "there"},
		}}}, &Statement{Node: &DeclarationNode{
		Ident: "myothervar",
		Expr: &CommandNode{
			Command: cmd,
			Action:  "create", Entity: "subnet",
			ParamNodes: map[string]interface{}{"vpc": HoleNode{key: "myvar"}},
			Refs:       map[string]interface{}{},
		}}}, &Statement{Node: &CommandNode{
		Action: "create", Entity: "instance",
		ParamNodes: map[string]interface{}{"subnet": HoleNode{key: "myothervar"}},
		Refs:       map[string]interface{}{"donald": "duck"},
	}},
	)

	clone := tree.Clone()

	if got, want := clone, tree; !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot %#v\n\nwant %#v", got, want)
	}

	clone.Statements[0].Node.(*DeclarationNode).Expr.(*CommandNode).ParamNodes["new"] = InterfaceNode{i: "mynode"}

	if got, want := clone.Statements, tree.Statements; reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot %s\n\nwant %s", got, want)
	}
}

func TestIsQuoted(t *testing.T) {
	tcases := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"'", false},
		{"\"", false},
		{"''", true},
		{"\"\"", true},
		{"\"'", false},
		{"'\"", false},
		{"'test\"", false},
		{"\"test'", false},
		{"\"test\"", true},
		{"'test'", true},
	}
	for i, tcase := range tcases {
		if got, want := isQuoted(tcase.in), tcase.out; got != want {
			t.Fatalf("%d: got %t, want %t", i+1, got, want)
		}
	}
}

type fakeCmd struct{}

func (*fakeCmd) ParamsSpec() params.Spec                                      { return nil }
func (*fakeCmd) Run(env.Running, map[string]interface{}) (interface{}, error) { return nil, nil }
