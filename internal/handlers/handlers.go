package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerGetHTML(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "file not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	r.ParseMultipartForm(20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "cannot read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "cannot read data", http.StatusInternalServerError)
		return
	}

	convertData, err := service.CheckStr(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	localFile, err := os.Create(time.Now().UTC().String())
	if err != nil {
		http.Error(w, "cannot create file", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()
	_, err = localFile.Write([]byte(convertData))
	if err != nil {
		http.Error(w, "cannot write file", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertData))

}
