package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/core"
	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/ls/lsutil"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestOrganizeImportsUnicode1(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    Ab,
    _aB,
    aB,
    _Ab,
} from './foo';

console.log(_aB, _Ab, aB, Ab);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    Ab,
    _Ab,
    _aB,
    aB,
} from './foo';

console.log(_aB, _Ab, aB, Ab);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSFalse,
			OrganizeImportsCollation:  lsutil.OrganizeImportsCollationOrdinal,
		},
	)
	f.VerifyOrganizeImports(t,
		`import {
    _aB,
    _Ab,
    aB,
    Ab,
} from './foo';

console.log(_aB, _Ab, aB, Ab);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSFalse,
			OrganizeImportsCollation:  lsutil.OrganizeImportsCollationUnicode,
		},
	)
}
