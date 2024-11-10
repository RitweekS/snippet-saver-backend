package response

type GetSnippetResponse struct {
	Id          int      `json:"id"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Note        string   `json:"note"`
	CodeSnippet string   `json:"snippet"`
	Language    string   `json:"language"`
}