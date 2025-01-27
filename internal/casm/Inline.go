package casm

import (
	"github.com/aleferri/casmeleon/pkg/parser"
)

// Inline is a temporary structure
type Inline struct {
	name   string   //name
	params []string //parameters name
	types  []uint32 //param types
}

// PruneToInline remove the header from the inline CST and return Inline and Body CST
func PruneToInline(lang *Language, op parser.CSTNode) (Inline, parser.CSTNode, error) {
	toks := op.Symbols()
	name := toks[1]
	children := op.Children()

	argsLUT, err := extractTypes(lang, children[0].Children())
	if err != nil {
		return Inline{}, nil, err
	}

	params := []string{}
	types := []uint32{}
	for k, v := range argsLUT {
		params = append(params, k)
		types = append(types, v)
	}

	body := children[1]
	return Inline{name: name.Value(), params: params, types: types}, body, nil
}
