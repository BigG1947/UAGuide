package router

import (
	"UAGuide/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func callCityApi(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
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

	//var response []string
	//
	//for _, city := range cities {
	//	if city["NP"] == "М" {
	//		response = append(response, city["NU"])
	//	}
	//}

	jsonResponse, err := json.Marshal(cities)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	_, err = writer.Write(jsonResponse)
	if err != nil {
		log.Printf("%s\n", err)
		return
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

func apiUploadProfileImages(writer http.ResponseWriter, request *http.Request) {
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

	src, err := uploadProfileImages(body)
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}

	responseMap["result"] = true
	responseMap["src"] = "/" + src
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

func apiUploadPlacesImages(writer http.ResponseWriter, request *http.Request) {
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
	photo := request.FormValue("photo")
	cropPhoto := request.FormValue("crop_photo")

	photoSrc, err := uploadPlacesImages([]byte(photo))
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}

	cropPhotoSrc, err := uploadPlacesImages([]byte(cropPhoto))
	if err != nil {
		log.Printf("%s\n", err)
		writer.WriteHeader(500)
		return
	}

	responseMap["result"] = true
	responseMap["photo_src"] = "/" + photoSrc
	responseMap["crop_photo_src"] = "/" + cropPhotoSrc
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
