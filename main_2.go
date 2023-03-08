package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
)

func main() {
  // Use a goroutine to download the image concurrently
  go downloadImage("https://www.example.com/image.jpg")
}

func downloadImage(url string) {
  // Create an HTTP client
  client := &http.Client{}

  // Create a request to the specified URL
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    fmt.Println("Error creating request:", err)
    return
  }

  // Send the request and get the response
  res, err := client.Do(req)
  if err != nil {
    fmt.Println("Error downloading image:", err)
    return
  }
  defer res.Body.Close()

  // Create a local file to store the image
  file, err := os.Create("image.jpg")
  if err != nil {
    fmt.Println("Error creating local file:", err)
    return
  }
  defer file.Close()

  // Write the image data to the local file
  _, err = io.Copy(file, res.Body)
  if err != nil {
    fmt.Println("Error writing to local file:", err)
    return
  }

  // Get the size of the downloaded image
  fileInfo, err := file.Stat()
  if err != nil {
    fmt.Println("Error getting file info:", err)
    return
  }
  fmt.Println("Image size:", fileInfo.Size())
}
