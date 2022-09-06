package uploadresponse

import "time"

type UploadResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Reason  string      `json:"reason"`
	Data    interface{} `json:"data"`
}
type File struct {
	FileName string
	FileSize int64
	FileType string
	FileUrl  string
}

type UploadFileUrl struct {
	UploadUrl string
	Expires   time.Time
	FileUrl   string
}
