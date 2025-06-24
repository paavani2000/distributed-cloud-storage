package controllers

import (
	"context"
	"file_storage/utils"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	// Safely get the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Failed to read uploaded file: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file found in request"})
		return
	}

	// Safely open the file
	src, err := file.Open()
	if err != nil {
		log.Printf("Failed to open file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open uploaded file"})
		return
	}
	defer src.Close()

	// Read file content into buffer
	buf := make([]byte, file.Size)
	_, err = io.ReadFull(src, buf)
	if err != nil {
		log.Printf("Failed to read file content: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read uploaded file"})
		return
	}

	// Generate IDs
	fileID := uuid.New().String()
	chunkID := uuid.New().String()
	s3Key := "uploads/" + chunkID + "_" + file.Filename

	// Upload to S3
	err = utils.UploadToS3(s3Key, buf)
	if err != nil {
		log.Printf("S3 upload failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "S3 upload failed"})
		return
	}

	// Insert metadata into `files` table
	_, err = utils.DB.Exec(context.Background(),
		`INSERT INTO files (id, filename, user_id, uploaded_at) VALUES ($1, $2, $3, $4)`,
		fileID, file.Filename, "user123", time.Now())
	if err != nil {
		log.Printf("Failed to insert into files table: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
		return
	}

	// Insert metadata into `file_chunks` table
	checksum := fmt.Sprintf("%x", crc32.ChecksumIEEE(buf))
	_, err = utils.DB.Exec(context.Background(),
		`INSERT INTO file_chunks (id, file_id, chunk_index, s3_key, size, checksum, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		chunkID, fileID, 0, s3Key, len(buf), checksum, time.Now())
	if err != nil {
		log.Printf("Failed to insert into file_chunks table: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_id": fileID})
}

func DownloadFile(c *gin.Context) {
	fileID := c.Param("id")
	var s3Key, filename string

	err := utils.DB.QueryRow(context.Background(),
		`SELECT c.s3_key, f.filename FROM file_chunks c
		 JOIN files f ON c.file_id = f.id
		 WHERE f.id = $1 AND c.chunk_index = 0`, fileID).
		Scan(&s3Key, &filename)

	if err != nil {
		log.Printf("❌ Download failed for file ID %s: %v\n", fileID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Download from S3
	content, err := utils.DownloadFromS3(s3Key)
	if err != nil {
		log.Printf("❌ S3 download failed for key %s: %v\n", s3Key, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download file"})
		return
	}

	// Send file as download response
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/octet-stream", content)
}
