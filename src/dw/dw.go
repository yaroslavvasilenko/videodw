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
	"video-dw/src/dw/url"
)

func DownloadFile(filepath string, urls []string) {
	wg := new(sync.WaitGroup)
	var nameFile int

	for _, urldw := range urls {
		nameFile++
		log.Printf("STARTED DOWNLOADING URL: %s\n", urldw)
		wg.Add(1)
		// ToDo: channel error collection
		err := downloadSingleFile(urldw, filepath+strconv.Itoa(nameFile)+".mp4", wg)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 3)
	}
	wg.Wait()
}

// ToDo: trace all uncompleted downloads with error channels
// ToDo: filepath -> dirpath. Сделать директорию и названия видео числовое 12345

func downloadSingleFile(urldw string, filepath string, wg *sync.WaitGroup) error {
	go func() error {
		defer wg.Done()
		out, err := os.Create(filepath)
		if err != nil {
			return &url.DownloadError{
				InterruptedDownload: urldw,
				Message:             fmt.Sprintf("Failed to create file: %s", filepath),
				Err:                 err}
		}
		defer out.Close()

		// Get the data
		resp, err := http.Get(urldw)
		if err != nil {
			return &url.DownloadError{
				InterruptedDownload: urldw,
				Message:             "Get request failed",
				Err:                 err}
		}
		defer resp.Body.Close()

		// Check server response
		if resp.StatusCode != http.StatusOK {
			return &url.DownloadError{
				InterruptedDownload: urldw,
				Message:             "Status code doesn't equal 200",
				Err:                 err}
		}

		// Writer the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return &url.DownloadError{
				InterruptedDownload: urldw,
				Message:             "Failed to write data",
				Err:                 err}
		}

		log.Println("Один файл загружен")
		return nil
	}()
	return nil
}
