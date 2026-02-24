package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/ls/lsutil"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestInlayHintsEnumMemberValue(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `enum E {
    A,
    AA,
    B = 10,
    BB,
    C = 'C',
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{InlayHints: lsutil.InlayHintsPreferences{IncludeInlayEnumMemberValueHints: true}})
}
