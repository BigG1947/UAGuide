package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func callCityApi(writer http.ResponseWriter, request *http.Request) {
	var cities []map[string]string
	file, err := os.Open("./static/data/city.json")
	if err != nil{
		log.Printf("Error in opening file with cuty: %s\n", err)
		return
	}
	data, err := ioutil.ReadAll(file)
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
	if err != nil{
		log.Printf("Error in reading data: %s\n", err)
		return
	}
	err = json.Unmarshal(data, &cities)
	if err != nil{
		log.Printf("Error in unmarshaling data in map: %s\n", err)
		return
	}

	for _, city := range cities {
		if city["NP"] == "М"{
			fmt.Fprintf(writer, "City: %s\n", city["NU"])
		}
	}
}
