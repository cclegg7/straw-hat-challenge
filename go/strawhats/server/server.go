package server

import (
	"fmt"
	"github.com/cclegg7/straw-hat-challenge/configs"
	"github.com/cclegg7/straw-hat-challenge/database"
	"log"
	"net/http"

	"github.com/cclegg7/straw-hat-challenge/clients/aws"
)

type Server struct {
	database *database.Database
	s3       aws.S3Client
	configs  *configs.Server
}

func New(database *database.Database, s3 aws.S3Client, configs *configs.Server) *Server {
	return &Server{
		database: database,
		s3:       s3,
		configs:  configs,
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
	if s.configs.UseHTTPS {
		go func() {
			if err := http.ListenAndServe(":80", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				redirectURL := fmt.Sprintf("https://%s:%d%s", s.configs.Hostname, s.configs.Port, request.RequestURI)
				http.Redirect(writer, request, redirectURL, http.StatusMovedPermanently)
			})); err != nil {
				log.Fatalf("error redirecting http: %v", err)
			}
		}()
		return http.ListenAndServeTLS(fmt.Sprintf(":%d", s.configs.Port), s.configs.CertFilePath, s.configs.CertKeyPath, nil)
	} else {
		return http.ListenAndServe(fmt.Sprintf(":%d", s.configs.Port), nil)
	}
}
