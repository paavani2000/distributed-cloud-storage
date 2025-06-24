package models

import "time"

type FileMeta struct {
	ID         string    `json:"id"`
	Filename   string    `json:"filename"`
	UserID     string    `json:"user_id"`
	UploadedAt time.Time `json:"uploaded_at"`
}
