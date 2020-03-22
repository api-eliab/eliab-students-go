package main

import (
	"mime/multipart"
)

// FileRequest doc ...
type FileRequest struct {
	File FileInfo `json:"file"`
}

// FileInfo doc ...
type FileInfo struct {
	Extension   string
	Description string
	File        multipart.File
	FileHeader  *multipart.FileHeader
}
