package start

import (
	"github.com/yaroslavvasilenko/videodw/src/dw"
	handler "github.com/yaroslavvasilenko/videodw/src/dw/url"
)

func Start(url string) {
	urlAll := handler.MakeRequest(url)
	videoname := handler.GetVideoName(urlAll)
	uniqueUrls := handler.MinimizeUrls(string(urlAll))
	dw.DownloadFile(
		videoname,
		uniqueUrls,
	)
}
