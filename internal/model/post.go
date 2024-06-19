package model

// Post struct definitions
type Post struct {
	Owner       UserPreview `json:"owner"`
	Text        string      `json:"text"`
	Like        int         `json:"likes"`
	Tags        []string    `json:"tags"`
	PublishDate string      `json:"publishDate"`
}

// APIResponsePost struct to match the JSON structure for posts
type APIResponsePost struct {
	Data  []Post `json:"data"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
