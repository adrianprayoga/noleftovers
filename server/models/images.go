package models

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ImageService struct {
	DB *sql.DB
}

func (service *ImageService) ValidateImage(file multipart.File, fileHeader *multipart.FileHeader) (int, error) {
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("unable to read file: %w", err)
	}

	fileType := http.DetectContentType(buff)
	if fileType != "image/jpeg" &&
		fileType != "image/png" &&
		fileType != "image/jpg" {
		return http.StatusBadRequest, fmt.Errorf("provided file format is not allowed. Please upload a JPEG or PNG image: %w", err)
	}

	// point back the pointer of the file
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error editing file pointer: %w", err)
	}

	return http.StatusOK, nil
}

func (service *ImageService) UploadImage(file io.Reader, fileName string) error {
	// TODO: change location
	basePath := "/Users/apr/noleftovers/client/public/images/recipe"
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s", basePath, fileName))
	if err != nil {
		return fmt.Errorf("error creating a file: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return fmt.Errorf("error in writing image: %w", err)
	}

	return nil
}

func (service *ImageService) RemoveImage(fileName string) error {
	// TODO: change location
	basePath := "/Users/apr/noleftovers/client/public/images/recipe"
	err := os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.Remove(fmt.Sprintf("%s/%s", basePath, fileName))
	if err != nil {
		return fmt.Errorf("error removing file: %w", err)
	}

	return nil
}
