package controllers

import (
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
)

type ImageResource struct {
	Service *models.ImageService
}

func (rs ImageResource) Routes() chi.Router {
	r := chi.NewRouter()
	//r.Post("/", rs.Upload) // GET /posts - Read a list of measure.

	return r
}

//func (rs ImageResource) Upload(w http.ResponseWriter, r *http.Request) {
//	// helped by: https://freshman.tech/file-upload-golang/
//
//	fmt.Println("File Upload Endpoint Hit")
//	var maxUploadSize int64 = 10 << 20 // 10 binary shifted 20 times. 10 * 2^20 = 10 MB
//
//	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
//	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
//		http.Error(w, "The uploaded image is too big. Please use an image less than 10MB in size", http.StatusBadRequest)
//		return
//	}
//
//	file, fileHeader, err := r.FormFile("image")
//	if err != nil {
//		fmt.Println("Error Retrieving the file")
//		fmt.Println(err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	defer file.Close()
//	fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
//	fmt.Printf("File Size: %+v\n", fileHeader.Size)
//
//	errCode, err := rs.Service.ValidateImage(file, fileHeader)
//	if err != nil {
//		http.Error(w, err.Error(), errCode)
//		return
//	}
//
//	if err != rs.Service.UploadImage(file, fileHeader.Filename) {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Fprintf(w, "Successfully Uploaded File\n")
//}
