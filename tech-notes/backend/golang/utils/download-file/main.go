// Reference: https://golangbyexample.com/download-image-file-url-golang/

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	filename := "cat.jpeg"
	url := "https://images.pexels.com/photos/1741205/pexels-photo-1741205.jpeg?cs=srgb&dl=pexels-lina-kivaka-1741205.jpg&fm=jpg&w=640&h=960"

	err := downloadFile(url, filename)
	if err != nil {
		fmt.Printf("downloadFile: %s", err)
		return
	}
}

func downloadFile(URL, fileName string) error {
	// Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("error download file")
	}
	// Create an empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
