package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	. "github.com/microsoft/typescript-go/pkg/fourslash/tests/util"
	"github.com/microsoft/typescript-go/pkg/ls"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestJavaScriptModules14(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: myMod.js
if (true) {
    exports.b = true;
} else {
    exports.n = 3;
}
function fn() {
    exports.s = 'foo';
}
var invisible = true;
// @Filename: isGlobal.js
var y = 10;
// @Filename: consumer.js
var x = require('myMod');
/**/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:    "y",
					SortText: new(string(ls.SortTextGlobalsOrKeywords)),
				},
			},
			Excludes: []string{
				"invisible",
			},
		},
	})
}
