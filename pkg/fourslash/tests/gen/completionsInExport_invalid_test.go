package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	. "github.com/microsoft/typescript-go/pkg/fourslash/tests/util"
	"github.com/microsoft/typescript-go/pkg/ls"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestCompletionsInExport_invalid(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function topLevel() {}
if (!!true) {
  const blockScoped = 0;
  export { /**/ };
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"topLevel",
				&lsproto.CompletionItem{
					Label:    "type",
					SortText: new(string(ls.SortTextGlobalsOrKeywords)),
				},
			},
		},
	})
}
