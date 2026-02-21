package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestFindAllRefsGlobalThisKeywordInModule(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noLib: true
/*1*/this;
export const c = 1;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t, "1")
}
