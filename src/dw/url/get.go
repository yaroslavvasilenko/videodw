package url

import (
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest(inputUrl string) []byte {
	resp, err := http.Get(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
