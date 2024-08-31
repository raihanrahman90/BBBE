package utils

import (
	"encoding/base64"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func SaveBase64Image(base64String string) (string, error) {
	// Extract the base64 data from the string (skip "data:image/jpeg;base64," prefix if present)
	base64Data := strings.Split(base64String, ",")[1]

	// Decode base64 data
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	// Generate a unique file name
	filename := "image_" + RandomString(10) + ".jpg" // Use a random string generator for unique filenames

	// Save decoded data to a file
	filePath := filepath.Join(os.Getenv("PATH_STATIC"), filename)
	err = ioutil.WriteFile(filePath, decoded, 0644)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func GetImageUrl(filename string) string{
	path := filepath.Join("/static/", filename)
	baseURL := &url.URL{
		Scheme: "http",
		Host:   os.Getenv("DOMAIN"),
		Path:   path,
	}
	return baseURL.String()
}