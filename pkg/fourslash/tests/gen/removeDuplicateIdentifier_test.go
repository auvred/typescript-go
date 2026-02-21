package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestRemoveDuplicateIdentifier(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class foo{}
function foo() { return null; }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToBOF(t)
	f.DeleteAtCaret(t, 11)
}
