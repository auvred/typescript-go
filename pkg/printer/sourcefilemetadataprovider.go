package printer

import (
	"github.com/microsoft/typescript-go/pkg/ast"
	"github.com/microsoft/typescript-go/pkg/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
}
