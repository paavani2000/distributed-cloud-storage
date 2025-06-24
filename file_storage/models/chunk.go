package models

import "time"

type FileChunk struct {
	ID        string    `json:"id"`
	FileID    string    `json:"file_id"`
	Index     int       `json:"chunk_index"`
	S3Key     string    `json:"s3_key"`
	Size      int       `json:"size"`
	Checksum  string    `json:"checksum"`
	CreatedAt time.Time `json:"created_at"`
}
