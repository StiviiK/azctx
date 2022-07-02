package pkg

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/StiviiK/azctx/utils"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/manifoldco/promptui"
)

// BuildPrompt builds a prompt for the user to select a subscription
func BuildPrompt(subscriptions []Subscription) promptui.Select {
	subscriptionNames := GetAzureSubscriptionNames(subscriptions)
	maxContextLength := utils.GetLongestStringLength(subscriptionNames)

	return promptui.Select{
		Label:        fmt.Sprint("  Name" + strings.Repeat(" ", maxContextLength-4) + " | " + "Id" + strings.Repeat(" ", 36-2) + " | " + "TenantId" + strings.Repeat(" ", 36-8) + " "),
		Items:        subscriptions,
		Templates:    buildTemplate(maxContextLength),
		HideSelected: true,
		Searcher: func(input string, index int) bool {
			return fuzzy.MatchNormalized(strings.ToLower(input), strings.ToLower(subscriptionNames[index]))
		},
		Size: int(utils.Min(len(subscriptions), 10)),
	}
}

// buildTemplate builds the template for the prompt
func buildTemplate(maxContextLength int) *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Inactive: builItemTemplate(maxContextLength, ""),
		Active:   "▸ " + builItemTemplate(maxContextLength, "bold")[2:],
		FuncMap:  newTemplateFuncMap(),
	}
}

// buildItemTemplate builds the item template
func builItemTemplate(maxContextLength int, additionalStyle string) string {
	return fmt.Sprintf(`  {{ repeat %[1]d " " | print .Name | trunc %[1]d | green | %[2]s }} | {{ repeat 36 " " | print .ID | trunc 36 | cyan | %[2]s }} | {{ repeat 36 " " | print .Tenant | trunc 36 | faint | %[2]s }} |`, maxContextLength, additionalStyle)
}

func newTemplateFuncMap() template.FuncMap {
	ret := sprig.TxtFuncMap()
	ret["green"] = promptui.Styler(promptui.FGGreen)
	ret["cyan"] = promptui.Styler(promptui.FGCyan)
	ret["bold"] = promptui.Styler(promptui.FGBold)
	ret["faint"] = promptui.Styler(promptui.FGFaint)
	return ret
}
