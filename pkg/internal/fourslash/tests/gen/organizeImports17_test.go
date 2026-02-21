package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/core"
	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/ls/lsutil"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestOrganizeImports17(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import { Both } from "module-specifiers-unsorted";
import { aa, CaseInsensitively, sorted } from "aardvark";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import { aa, CaseInsensitively, sorted } from "aardvark";
import { Both } from "module-specifiers-unsorted";
`,
		lsproto.CodeActionKindSourceSortImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSUnknown,
		},
	)
}
