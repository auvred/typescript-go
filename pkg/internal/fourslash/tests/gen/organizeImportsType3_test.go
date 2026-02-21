package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/core"
	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/ls/lsutil"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestOrganizeImportsType3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    d, 
    type d as D,
    type c,
    c as C,
    b,
    b as B,
    type A,
    a
} from './foo';
console.log(A, a, B, b, c, C, d, D);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    type A,
    b as B,
    c as C,
    type d as D,
    a,
    b,
    type c,
    d
} from './foo';
console.log(A, a, B, b, c, C, d, D);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSFalse,
			OrganizeImportsTypeOrder:  lsutil.OrganizeImportsTypeOrderInline,
		},
	)
}
