package sourcemap

import "github.com/microsoft/typescript-go/pkg/core"

type Source interface {
	Text() string
	FileName() string
	ECMALineMap() []core.TextPos
}
