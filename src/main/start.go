package start

import (
	"video-dw/src/dw"
	handler_url "video-dw/src/dw/handler-url"
)

func start(url string) {
	urlAll := handler_url.MakeRequest(url)
	videoname := handler_url.GetVideoName(urlAll)
	uniqueUrls := handler_url.MinimizeUrls(string(urlAll))
	dw.DownloadFile(
		videoname,
		uniqueUrls,
	)
}
