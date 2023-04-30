package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Database struct {
	Hostname string
	User     string
	Password string
	Name     string
}

type Server struct {
	Hostname     string
	UseHTTPS     bool
	Port         int
	CertFilePath string
	CertKeyPath  string
}

type FileStorage struct {
	UseS3    bool
	S3Region string
	S3Bucket string
}

type Configs struct {
	Server      *Server
	Database    *Database
	FileStorage *FileStorage
}

func New() (*Configs, error) {
	viper.SetConfigName("configs")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/strawhats/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found, continuing")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	setDefaults()

	return &Configs{
		Database: &Database{
			Hostname: viper.GetString("Database.Hostname"),
			User:     viper.GetString("Database.User"),
			Password: viper.GetString("Database.Password"),
			Name:     viper.GetString("Database.Name"),
		},
		Server: &Server{
			Hostname:     viper.GetString("Server.Hostname"),
			UseHTTPS:     viper.GetBool("Server.UseHTTPS"),
			Port:         viper.GetInt("Server.Port"),
			CertFilePath: viper.GetString("Server.CertFilePath"),
			CertKeyPath:  viper.GetString("Server.CertKeyPath"),
		},
		FileStorage: &FileStorage{
			UseS3:    viper.GetBool("FileStorage.UseS3"),
			S3Region: viper.GetString("FileStorage.S3Region"),
			S3Bucket: viper.GetString("FileStorage.S3Bucket"),
		},
	}, nil
}

func setDefaults() {
	viper.SetDefault("Database.Hostname", "localhost")
	viper.SetDefault("Database.User", "root")
	viper.SetDefault("Database.Name", "straw_hat_challenge")

	viper.SetDefault("Server.Hostname", "localhost")
	viper.SetDefault("Server.UseHTTPS", false)
	viper.SetDefault("Server.Port", 80)

	viper.SetDefault("FileStorage.UseS3", false)
}
