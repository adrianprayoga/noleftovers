package controllers

import (
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

type ImageResource struct {
	Service *models.ImageService
}

func (rs ImageResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(r chi.Router) {
		r.Use(GetCtx)      // Use the same middleware
		r.Get("/", rs.Get) // GET /posts/{id} - Read a single post by :id.
	})

	return r
}

func (rs ImageResource) Get(w http.ResponseWriter, r *http.Request) {
	filename, ok := r.Context().Value("id").(string)
	if !ok {
		http.Error(w, "invalid filename", http.StatusBadRequest)
	}

	basePath := viper.GetString("imageLocation")


	fmt.Println("filename", fmt.Sprintf("%s/%s", basePath, filename))
	buf, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", basePath, filename))

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/*")
	w.Write(buf)

	w.WriteHeader(http.StatusOK)
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
