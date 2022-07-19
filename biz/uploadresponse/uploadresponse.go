package uploadresponse

type UploadResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Data    File   `json:"data"`
}
type File struct {
	FileName string
	FileSize int64
	FileType string
	FileUrl  string
}
