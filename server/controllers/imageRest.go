package controllers

import (
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type ImageResource struct {
	Service *models.ImageService
}

func (rs ImageResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rs.Upload) // GET /posts - Read a list of measure.

	return r
}

func (rs ImageResource) Upload(w http.ResponseWriter, r *http.Request) {
	// helped by: https://freshman.tech/file-upload-golang/

	fmt.Println("File Upload Endpoint Hit")
	var maxUploadSize int64 = 10 << 20 // 10 binary shifted 20 times. 10 * 2^20 = 10 MB

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "The uploaded image is too big. Please use an image less than 1MB in size", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the file")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
	fmt.Printf("File Size: %+v\n", fileHeader.Size)

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileType := http.DetectContentType(buff)
	if fileType != "image/jpeg" &&
		fileType != "image/png" &&
		fileType != "image/jpg" {
		http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
		return
	}

	// point back the pointer of the file
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.MkdirAll("/Users/apr/uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: change location
	dst, err := os.Create(fmt.Sprintf("/Users/apr/uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
