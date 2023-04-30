package aws

import (
	"fmt"
	"io"
	"os"
)

type localStorageClient struct{}

func (*localStorageClient) Upload(file io.Reader, name string, contentType string) (string, error) {
	fileData, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll("./static/temp/local_s3", 0777)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(fmt.Sprintf("./static/temp/local_s3/%s", name), fileData, 0666)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/temp/local_s3/%s", name), nil
}
