package types

type StdResponse struct {
	// Response status code
	StatusCode int `json:"statusCode"`

	// Response message
	Message string `json:"message"`

	// Response result object
	Result interface{} `json:"result"`
}

type ResultSiteMeta struct {
	// Page title
	Title string `json:"title"`

	// Page description
	Description string `json:"desc"`

	// Default icon
	Icon string `json:"icon"`

	// Page open-graph thumb
	Thumb string `json:"thumb"`
}
