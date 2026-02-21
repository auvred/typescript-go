package printer

import (
	"github.com/microsoft/typescript-go/pkg/ast"
	"github.com/microsoft/typescript-go/pkg/core"
	"github.com/microsoft/typescript-go/pkg/tsoptions"
	"github.com/microsoft/typescript-go/pkg/tspath"
)

// NOTE: EmitHost operations must be thread-safe
type EmitHost interface {
	Options() *core.CompilerOptions
	SourceFiles() []*ast.SourceFile
	UseCaseSensitiveFileNames() bool
	GetCurrentDirectory() string
	CommonSourceDirectory() string
	IsEmitBlocked(file string) bool
	WriteFile(fileName string, text string, writeByteOrderMark bool) error
	GetEmitModuleFormatOfFile(file ast.HasFileName) core.ModuleKind
	GetEmitResolver() EmitResolver
	GetProjectReferenceFromSource(path tspath.Path) *tsoptions.SourceOutputAndProjectReference
	IsSourceFileFromExternalLibrary(file *ast.SourceFile) bool
}
