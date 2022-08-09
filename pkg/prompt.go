package pkg

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/StiviiK/azctx/utils"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/manifoldco/promptui"
)

// BuildPrompt builds a prompt for the user to select a subscription
func BuildPrompt(subscriptions []Subscription) promptui.Select {
	// Sort the subscriptions by name
	sort.Sort(subscriptionSorter(subscriptions))

	// Build the prompt
	var subscriptionNames utils.StringSlice = GetAzureSubscriptionNames(subscriptions)
	maxContextLength := subscriptionNames.LongestStringLength()

	return promptui.Select{
		Label: fmt.Sprint("Name" + strings.Repeat(" ", maxContextLength-4) + " | " + "SubscriptionId" + strings.Repeat(" ", 36-14) + " | " + "TenantId" + strings.Repeat(" ", 36-8)),
		Items: subscriptions,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ \" \" | repeat 4 }}{{ . }} |",
			Inactive: builItemTemplate(maxContextLength, ""),
			Active:   "▸ " + builItemTemplate(maxContextLength, "bold")[2:],
			FuncMap:  newTemplateFuncMap(),
		},
		HideSelected: true,
		Searcher: func(input string, index int) bool {
			return fuzzy.MatchNormalized(strings.ToLower(input), strings.ToLower(subscriptionNames[index]))
		},
		Size:   int(utils.Min(len(subscriptions), 10)),
		Stdout: utils.NoBellStdout,
	}
}

// buildItemTemplate builds the item template
func builItemTemplate(maxContextLength int, additionalStyle string) string {
	return fmt.Sprintf("  {{ repeat %[1]d \" \" | print .Name | trunc %[1]d | green | %[2]s }} | {{ repeat 36 \" \" | print .ID | trunc 36 | cyan | %[2]s }} | {{ repeat 36 \" \" | print .Tenant | trunc 36 | faint | %[2]s }} |", maxContextLength, additionalStyle)
}

func newTemplateFuncMap() template.FuncMap {
	ret := sprig.TxtFuncMap()
	ret["green"] = promptui.Styler(promptui.FGGreen)
	ret["cyan"] = promptui.Styler(promptui.FGCyan)
	ret["bold"] = promptui.Styler(promptui.FGBold)
	ret["faint"] = promptui.Styler(promptui.FGFaint)
	return ret
}
