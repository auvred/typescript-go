package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/pkg/fourslash"
	. "github.com/microsoft/typescript-go/pkg/fourslash/tests/util"
	"github.com/microsoft/typescript-go/pkg/ls"
	"github.com/microsoft/typescript-go/pkg/lsp/lsproto"
	"github.com/microsoft/typescript-go/pkg/testutil"
)

func TestCompletionDetailSignature(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `

/*a*/

function foo(x: string): string;
function foo(x: number): number;
function foo(x: any): any {
    return x;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:    "foo",
					Kind:     new(lsproto.CompletionItemKindFunction),
					SortText: new(string(ls.SortTextLocationPriority)),
					Detail:   new("function foo(x: string): string\nfunction foo(x: number): number"),
				},
			},
		},
	})
}
