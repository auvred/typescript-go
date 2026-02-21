package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestSignatureHelpJSX(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: test.tsx
//@jsx: react
declare var React: any;
const z = <div>{[].map(x => </**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.VerifyNoSignatureHelpWithContext(t, &lsproto.SignatureHelpContext{TriggerKind: lsproto.SignatureHelpTriggerKindTriggerCharacter, TriggerCharacter: new("<"), IsRetrigger: false})
}
