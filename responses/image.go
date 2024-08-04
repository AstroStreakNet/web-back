package responses

type ImageGet struct {
	ID         string   `json:"id"`
	User       string   `json:"user"`
	UploadDate string   `json:"upload_date"`
	URL        string   `json:"url"`
	Tags       []string `json:"tags"`
}

type ImagePost struct {
	Success bool `json:"success"`
}

type ImageUpdate ImagePost
type ImageDelete ImagePost
