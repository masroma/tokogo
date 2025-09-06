package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadFile handles file upload and returns the file path
func UploadFile(file *multipart.FileHeader, uploadDir string) (string, error) {
	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Validate file type
	allowedTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/gif"}
	contentType := file.Header.Get("Content-Type")
	if !isAllowedType(contentType, allowedTypes) {
		return "", fmt.Errorf("file type not allowed. Allowed types: %v", allowedTypes)
	}

	// Validate file size (max 5MB)
	maxSize := int64(5 * 1024 * 1024) // 5MB
	if file.Size > maxSize {
		return "", fmt.Errorf("file size too large. Maximum size: 5MB")
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), generateRandomString(10), ext)
	filePath := filepath.Join(uploadDir, filename)

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	// Return relative path for database storage
	return filePath, nil
}

// DeleteFile deletes a file from the filesystem
func DeleteFile(filePath string) error {
	if filePath == "" {
		return nil
	}

	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

// isAllowedType checks if the content type is in the allowed list
func isAllowedType(contentType string, allowedTypes []string) bool {
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			return true
		}
	}
	return false
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
