package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestOrganizeImports19(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const a = 1;
export { a };

const b = 1;
export { b };

const c = 1;
export { c };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`const a = 1;
export { a };

const b = 1;
export { b };

const c = 1;
export { c };
`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
