package main

import (
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ImageInfo struct {
    Width int
    Height int
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        imageUrl := r.URL.Query().Get("url")

        response, err := http.Get(imageUrl)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer response.Body.Close()

        // Create a temporary file to store the image
        tempFile, err := os.Create(filepath.Base(imageUrl))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer tempFile.Close()

        // Copy the image data to the temporary file
        _, err = io.Copy(tempFile, response.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Open the temporary file and get the image dimensions
        file, err := os.Open(tempFile.Name())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer file.Close()

        image, _, err := image.DecodeConfig(file)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Create the JSON response with the image dimensions
        imageInfo := ImageInfo{
            Width: image.Width,
            Height: image.Height,
        }
        jsonResponse, err := json.Marshal(imageInfo)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Write the JSON response to the response writer
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonResponse)
    })

    http.ListenAndServe(":8080", nil)
}
