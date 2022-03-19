package handler_url

import (
	"encoding/json"
	"log"
	"strings"
)

func GetVideoName(jsonString []byte) string {
	type myStruct struct {
		Name string `json:"name"`
	}
	var s myStruct
	_ = json.Unmarshal(jsonString, &s)
	res := strings.Trim(s.Name, `{}`)
	log.Println("Found video name: ", res)
	return res
}
