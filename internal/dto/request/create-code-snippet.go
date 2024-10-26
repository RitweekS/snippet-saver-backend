package request

type CreateSnippetRequest struct {
	Tags []string `json:"tags"`
	Title string `json:"title"`
	Note string `json:"note"`
	CodeSnippet string `json:"snippet"`
	Language string `json:"language"`
}