package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestGetOutliningSpansForUnbalancedEndRegion(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// bottom-heavy region balance
[|// #region matched

// #endregion matched|]

// #endregion unmatched`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOutliningSpans(t, lsproto.FoldingRangeKindRegion)
}
