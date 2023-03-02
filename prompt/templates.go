package prompt

var (
	template_Long = promptTemplate{
		Label:    "{{ repeat 4 \" \" }}{{ repeat %[1]d \" \" | print \"Name\" | trunc %[1]d }} {{ repeat 38 \" \" | print \"SubscriptionId\" | trunc 38 }} {{ repeat %[2]d \" \" | print \"Tenant\" | trunc %[2]d }}",
		Active:   "▸ {{ repeat %[1]d \" \" | print .Name | trunc %[1]d | green | %[3]s }} | {{ repeat 36 \" \" | print .Id | trunc 36 | cyan | %[3]s }} | {{ repeat %[2]d \" \" | print \")\" | print .Tenant | print \" (\" | print .TenantName | trunc %[2]d | faint | %[3]s }}",
		Inactive: "{{ repeat 2 \" \" }}{{ repeat %[1]d \" \" | print .Name | trunc %[1]d | green | %[3]s }} | {{ repeat 36 \" \" | print .Id | trunc 36 | cyan | %[3]s }} | {{ repeat %[2]d \" \" | print \")\" | print .Tenant | print \" (\" | print .TenantName | trunc %[2]d | faint | %[3]s }}",
	}
	template_Short = promptTemplate{
		Label:    "{{ repeat 4 \" \" }}{{ repeat %[1]d \" \" | print \"Name\" | trunc %[1]d }} {{ repeat %[2]d \" \" | print \"Tenant\" | trunc %[2]d }}",
		Active:   "▸ {{ repeat %[1]d \" \" | print .Name | trunc %[1]d | green | %[3]s }} | {{ repeat %[2]d \" \" | print .TenantName | trunc %[2]d | cyan | %[3]s }}",
		Inactive: "{{ repeat 2 \" \" }}{{ repeat %[1]d \" \" | print .Name | trunc %[1]d | green | %[3]s }} | {{ repeat %[2]d \" \" | print .TenantName | trunc %[2]d | cyan | %[3]s }}",
	}
)

// templateName returns the template to use
func template() promptTemplate {
	if ShortPrompt {
		return template_Short
	}

	return template_Long
}
