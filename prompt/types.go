package prompt

// promptJsonTemplate represents the json template for the prompt (see templates/*.json)
type promptJsonTemplate struct {
	Label  string `json:"label"`
	Prompt string `json:"prompt"`
}
