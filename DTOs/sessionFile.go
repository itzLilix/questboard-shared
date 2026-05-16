package dtos

import "time"

type SessionFile struct {
	ID         string    `json:"id"`
	SessionID  string    `json:"sessionId"`
	UploaderID string    `json:"uploaderId"`
	Filename   string    `json:"filename"`
	URL        string    `json:"url"`
	MimeType   string    `json:"mimeType"`
	SizeBytes  int64     `json:"sizeBytes"`
	UploadedAt time.Time `json:"uploadedAt"`
}
