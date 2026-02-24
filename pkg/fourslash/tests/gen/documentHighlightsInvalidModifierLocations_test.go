package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	. "github.com/microsoft/typescript-go/pkg/fourslash/tests/util"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestDocumentHighlightsInvalidModifierLocations(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C {
    m([|readonly|] p) {}
}
function f([|readonly|] p) {}

class D {
    m([|public|] p) {}
}
function g([|public|] p) {}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
