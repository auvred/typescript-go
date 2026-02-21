package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestIntellisenseInObjectLiteral(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var x = 3;

class Foo {
    static something() {
        return { "prop": /**/x };
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "", "var x: number", "")
}
