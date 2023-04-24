package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const megabyte = 1 << 20
const fileFormKey = "file"

type UploadFileResponse struct {
	FileToken string `json:"file_token"`
}

func (s *Server) uploadFileHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseMultipartForm(megabyte * 150); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	uploadedFileHeaders := req.MultipartForm.File[fileFormKey]
	if len(uploadedFileHeaders) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("expected exactly 1 file uploaded in multipart data"))
		return
	}

	fileHeader := uploadedFileHeaders[0]
	contentType := fileHeader.Header.Get("Content-Type")
	file, err := fileHeader.Open()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	token, err := uuid.NewUUID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to generate uuid: %v", err.Error())))
		return
	}

	fileUrl, err := s.s3.Upload(file, token.String(), contentType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to upload file to s3: %v", err.Error())))
		return
	}

	if err := s.database.InsertFile(token.String(), fileUrl, contentType); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to save file record to db: %v", err.Error())))
		return
	}

	response := UploadFileResponse{
		FileToken: token.String(),
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to marshal response: %v", err.Error())))
		return
	}

	w.Write(responseBytes)
}
