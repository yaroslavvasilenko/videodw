package dw

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	"video-dw/src/handler-err"
)

func DownloadFile(filepath string, urls []string) {
	wg := new(sync.WaitGroup)
	var nameFile int
	for _, url := range urls {
		nameFile++
		log.Printf("STARTED DOWNLOADING URL: %s\n", url)
		wg.Add(1)
		err := downloadSingleFile(url, filepath+strconv.Itoa(nameFile)+".mp4", wg)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 3)
	}
	wg.Wait()
}

func downloadSingleFile(url string, filepath string, wg *sync.WaitGroup) *handler_err.DownloadError {
	go func() *handler_err.DownloadError {
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

		// Writer the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return &handler_err.DownloadError{
				InterruptedDownload: url,
				Message:             "Failed to write data",
				Err:                 err}
		}
		wg.Done()
		log.Println("Один файл загружен")
		return nil
	}()
	return nil
}
