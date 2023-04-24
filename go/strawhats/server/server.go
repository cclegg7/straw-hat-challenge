package server

import (
	"fmt"
	"net/http"

	"github.com/cclegg7/straw-hat-challenge/clients/aws"

	"github.com/cclegg7/straw-hat-challenge/db"
)

type Server struct {
	database *db.Database
	s3       *aws.S3Client
}

func NewServer(database *db.Database, s3 *aws.S3Client) *Server {
	return &Server{
		database: database,
		s3:       s3,
	}
}

func (s *Server) Serve() error {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/users", s.getUsersHandler)
	http.HandleFunc("/climb", s.postClimbHandler)
	http.HandleFunc("/run", s.postRunHandler)
	http.HandleFunc("/scores", s.getScoresHandler)
	http.HandleFunc("/runs", s.listUserRunsHandler)
	http.HandleFunc("/climbs", s.listUserClimbsHandler)
	http.HandleFunc("/upload-file", s.uploadFileHandler)
	fmt.Println("Serving!")
	return http.ListenAndServe(":81", nil)
}
