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

func BuildPrompt(subscriptions []Subscription) promptui.Select {
	subscriptionNames := GetAzureSubscriptionNames(subscriptions)
	maxContextLength := utils.GetLongestStringLength(subscriptionNames)

	return promptui.Select{
		Label:        fmt.Sprint("  Name" + strings.Repeat(" ", maxContextLength-4) + " | " + "ID" + strings.Repeat(" ", 36-2) + " | " + "Tenant" + strings.Repeat(" ", 36-6) + " "),
		Items:        subscriptions,
		Templates:    buildTemplate(maxContextLength),
		HideSelected: true,
		Searcher: func(input string, index int) bool {
			return fuzzy.MatchNormalized(strings.ToLower(input), strings.ToLower(subscriptionNames[index]))
		},
		Size: int(utils.Min(len(subscriptions), 10)),
	}
}

func buildTemplate(maxContextLength int) *promptui.SelectTemplates {
	itemTemplate := fmt.Sprintf(`  {{ repeat %[1]d " " | print .Name | trunc %[1]d | green | %[2]s }} | {{ repeat 36 " " | print .ID | trunc 36 | cyan | %[2]s }} | {{ repeat 36 " " | print .Tenant | trunc 36 | faint | %[2]s }} |`, maxContextLength, "")
	return &promptui.SelectTemplates{
		Inactive: itemTemplate,
		Active:   "▸ " + itemTemplate[2:],
		FuncMap:  newTemplateFuncMap(),
	}
}

func newTemplateFuncMap() template.FuncMap {
	ret := sprig.TxtFuncMap()
	ret["green"] = promptui.Styler(promptui.FGGreen)
	ret["cyan"] = promptui.Styler(promptui.FGCyan)
	ret["bold"] = promptui.Styler(promptui.FGBold)
	ret["faint"] = promptui.Styler(promptui.FGFaint)
	return ret
}
