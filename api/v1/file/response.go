package file

// UploadResponse the response object of a successful upload request.
type UploadResponse struct {
	FileName  string `json:"file_name"`
	AssertURL string `json:"asset_url"`
}
