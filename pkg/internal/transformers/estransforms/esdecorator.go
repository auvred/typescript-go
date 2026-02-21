package estransforms

import (
	"github.com/microsoft/typescript-go/pkg/ast"
	"github.com/microsoft/typescript-go/pkg/transformers"
)

type esDecoratorTransformer struct {
	transformers.Transformer
}

func (ch *esDecoratorTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newESDecoratorTransformer(opts *transformers.TransformOptions) *transformers.Transformer {
	tx := &esDecoratorTransformer{}
	return tx.NewTransformer(tx.visit, opts.Context)
}
