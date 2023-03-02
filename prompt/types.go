package prompt

// promptJsonTemplate represents the json template for the prompt (see templates/*.json)
type promptJsonTemplate struct {
	Label    string `json:"label"`
	Active   string `json:"active"`
	Inactive string `json:"inactive"`
}
