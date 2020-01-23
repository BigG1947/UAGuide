package router

import (
	"UAGuide/models"
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
	if err != nil {
		log.Printf("Error in opening file with cuty: %s\n", err)
		return
	}
	data, err := ioutil.ReadAll(file)
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
	if err != nil {
		log.Printf("Error in reading data: %s\n", err)
		return
	}
	err = json.Unmarshal(data, &cities)
	if err != nil {
		log.Printf("Error in unmarshaling data in map: %s\n", err)
		return
	}

	for _, city := range cities {
		if city["NP"] == "М" {
			fmt.Fprintf(writer, "City: %s\n", city["NU"])
		}
	}
}

func checkPhone(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	phone, err := ioutil.ReadAll(request.Body)
	responseMap := make(map[string]string)
	if err != nil {
		writer.WriteHeader(500)
		responseMap["error"] = err.Error()
		responseMap["ok"] = "false"
		response, err := json.Marshal(responseMap)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(response)
		return
	}
	if models.CheckUserPhoneExist(string(phone), db) {
		writer.WriteHeader(200)
		responseMap["ok"] = "true"
		response, err := json.Marshal(responseMap)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(response)
		return
	} else {
		writer.WriteHeader(200)
		responseMap["ok"] = "false"
		response, err := json.Marshal(responseMap)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(response)
		return
	}
}

func checkEmail(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	email, err := ioutil.ReadAll(request.Body)
	responseMap := make(map[string]string)
	if err != nil {
		writer.WriteHeader(500)
		responseMap["error"] = err.Error()
		responseMap["ok"] = "false"
		response, err := json.Marshal(responseMap)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(response)
		return
	}
	if models.CheckUserEmailExist(string(email), db) {
		writer.WriteHeader(200)
		responseMap["ok"] = "true"
		response, err := json.Marshal(responseMap)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(response)
		return
	} else {
		writer.WriteHeader(200)
		responseMap["ok"] = "false"
		response, err := json.Marshal(responseMap)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write(response)
		return
	}
}

func apiUploadImages(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}

	if !checkLoginUser(session) {
		log.Printf("No auth user try load images\n")
		writer.WriteHeader(500)
		return
	}

	var responseMap = make(map[string]interface{})
	responseMap["result"] = false
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}

	tempFile, err := ioutil.TempFile("upload/profile-photo/", "upload-photo-*.jpg")
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}
	defer tempFile.Close()

	_, err = tempFile.Write(body)
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}

	responseMap["result"] = true
	responseMap["src"] = "/" + tempFile.Name()
	responseJson, err := json.Marshal(responseMap)
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}

	if _, err := writer.Write(responseJson); err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}
	return
}
