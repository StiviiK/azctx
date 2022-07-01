package pkg

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

const (
	maxContextLength = 25
)

func BuildPrompt() promptui.Select {
	return promptui.Select{
		Label:     "Select the Azure Subscription you want to use",
		Items:     []string{"Subscription 1", "Subscription 2", "Subscription 3"},
		Templates: buildTemplate(),
	}
}

func buildTemplate() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		//Inactive: fmt.Sprintf(`  {{ repeat %[1]d " " | print .Context | trunc %[1]d | %[2]s }} | {{ repeat %[1]d " " | print .Cluster | trunc %[1]d | %[2]s }} | {{ repeat %[1]d  " " | print .File | trunc %[1]d | %[2]s }} |`, maxContextLength, ""),
		//Active:   fmt.Sprintf(`▸ {{ repeat %[1]d " " | print .Context | trunc %[1]d | %[2]s }} | {{ repeat %[1]d " " | print .Cluster | trunc %[1]d | %[2]s }} | {{ repeat %[1]d  " " | print .File | trunc %[1]d | %[2]s }} |`, maxContextLength, "bold | cyan"),
		Label: fmt.Sprint("  Name" + strings.Repeat(" ", maxContextLength-7) + " | " + "ID" + strings.Repeat(" ", maxContextLength-7) + " | " + "Tenant" + strings.Repeat(" ", maxContextLength-4) + " "),
	}
}
