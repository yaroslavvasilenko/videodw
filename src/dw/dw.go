package dw

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"video-dw/src/handler-err"
)

func DownloadFile(filepath string, urls []string) {
	var nameFile int
	for _, url := range urls {
		nameFile++
		log.Printf("STARTED DOWNLOADING URL: %s\n", url)
		err := downloadSingleFile(url, filepath+strconv.Itoa(nameFile)+".mp4")
		if err != nil {
			log.Println(err)
			// ToDo: add comments to the error
			// e.g. "Problem with download has occurred"
			// count append name video
			// оставать имя из json (используй unicode decoder!!!)
		}
	}
}

func downloadSingleFile(url string, filepath string) *handler_err.DownloadError {
	out, err := os.Create(filepath)
	if err != nil {
		return &handler_err.DownloadError{
			InterruptedDownload: url,
			Message:             fmt.Sprintf("Failed to create file: %s", filepath),
			Err:                 err}
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return &handler_err.DownloadError{
			InterruptedDownload: url,
			Message:             "Get request failed",
			Err:                 err}
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return &handler_err.DownloadError{
			InterruptedDownload: url,
			Message:             "Status code doesn't equal 200",
			Err:                 err}
	}

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		filepath,
	)

	// Writer the body to file
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return &handler_err.DownloadError{
			InterruptedDownload: url,
			Message:             "Failed to write data",
			Err:                 err}
	}
	return nil
}
