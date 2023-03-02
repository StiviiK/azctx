package prompt

// promptTemplate represents the json template for the prompt (see templates/*.json)
type promptTemplate struct {
	Label    string `json:"label"`
	Active   string `json:"active"`
	Inactive string `json:"inactive"`
}
