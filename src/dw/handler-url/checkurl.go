package handler_url

import "regexp"

func MinimizeUrls(filePtr string) []string {
	re, _ := regexp.Compile(`https://events-storage\.webinar\.ru/api-storage/files/[a-z0-9]+/[0-9]{4}/[0-9]{2}/[0-9]{2}/[a-z0-9]+\.mp4`)
	foundUrls := re.FindAllString(filePtr, -1)
	// ToDo: add logging with basic info
	// e.g. log.Println("found 10 urls")
	var minimizedUrls []string
	for _, foundUrl := range foundUrls {
		urlAlreadyAdded := false
		for _, addedUrl := range minimizedUrls {
			if foundUrl == addedUrl {
				urlAlreadyAdded = true
				break
			}
		}
		if !urlAlreadyAdded {
			minimizedUrls = append(minimizedUrls, foundUrl)
		}
	}
	return minimizedUrls
}
