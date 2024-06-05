package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(response)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Fatalln(err.Error())
	}
}
