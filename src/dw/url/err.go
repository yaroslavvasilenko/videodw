package url

import "fmt"

type DownloadErrorInterface interface {
	Error() string
}

type DownloadError struct {
	InterruptedDownload string
	Message             string
	Err                 error
}

func (r *DownloadError) Error() string {
	return fmt.Sprintf("Downloading error occurred: %s", r.InterruptedDownload)
}
