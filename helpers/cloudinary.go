package helpers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"mini-challenge-pertemuan-sembilan/configs"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Validate
	err := validateImage(fileHeader)
	if err != nil {
		return "", err
	}

	// Add Cloudinary product environment credentials.
	cld, err := cloudinary.NewFromParams(configs.EnvCloudName(), configs.EnvCloudAPIKey(), configs.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	// Convert file
	fileReader, err := convertFile(fileHeader)
	if err != nil {
		return "", err
	}

	// Upload file
	uploadParam, err := cld.Upload.Upload(ctx, fileReader, uploader.UploadParams{
		PublicID: fileName,
		Folder:   configs.EnvCloudUploadFolder(),
	})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}

// Accept *bytes.Reader or *multipart.File
func convertFile(fileHeader *multipart.FileHeader) (*bytes.Reader, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content into an in-memory buffer
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	// Create a bytes.Reader from the buffer
	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil
}

// validasi extension yang diperbolehkan adalah PNG, JPG, JPEG dan dengan maksimum file size 5mb
func validateImage(fileHeader *multipart.FileHeader) (err error) {

	extension := filepath.Ext(fileHeader.Filename)
	size := fileHeader.Size
	fmt.Println(extension)
	fmt.Println(size)

	if strings.EqualFold(extension, ".PNG") || strings.EqualFold(extension, ".JPG") || strings.EqualFold(extension, ".JPEG") {
	} else {
		err = errors.New("Image must be PNG, JPG, or JPEG")
		return
	}

	if size > int64(5<<20) {
		fmt.Println(size)
		err = errors.New("Image file size max 5MB")
		return
	}

	return
}

func RemoveExtension(filename string) string {
	return path.Base(filename[:len(filename)-len(path.Ext(filename))])
}
