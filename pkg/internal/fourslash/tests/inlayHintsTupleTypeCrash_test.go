package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/ls/lsutil"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestInlayHintsTupleTypeCrash(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function iterateTuples(tuples: [string][]): void {
  tuples.forEach((l) => {})
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayFunctionParameterTypeHints: true,
		},
	})
}
