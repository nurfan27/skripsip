package app

import (
	"encoding/json"
	"log"

	"github.com/parnurzeal/gorequest"
)

type SiakadService struct {
	uri string
}

func (ss *SiakadService) CheckPhoneNumber(phoneNumber string) BrivaResponse {
	var response BrivaResponse

	uri := ss.uri + "/payment/briva"
	log.Println("URL : ", uri)

	log.Println("ap_key")
	log.Println(Env.Get("siakad_app_key"))

	_, resp, errorResponse := gorequest.New().
		Get(uri).
		Set("app-key", Env.Get("siakad_app_key")).
		Set("client-key", phoneNumber).
		End()

	log.Println("Response : ")
	log.Println(resp)

	if errorResponse != nil {
		log.Println(errorResponse)
	}

	json.Unmarshal([]byte(resp), &response)
	return response
}

func (ss *SiakadService) GetBriva(phoneNumber string) BrivaResponse {
	var response BrivaResponse

	log.Println("phone number : ")
	log.Println(phoneNumber)

	uri := ss.uri + "/payment/briva"
	log.Println("URL : ", uri)

	log.Println("ap_key")
	log.Println(Env.Get("siakad_app_key"))

	_, resp, errorResponse := gorequest.New().
		Get(uri).
		Set("app-key", Env.Get("siakad_app_key")).
		Set("client-key", phoneNumber).
		End()

	log.Println("Response : ")
	log.Println(resp)

	if errorResponse != nil {
		log.Println(errorResponse)
	}

	json.Unmarshal([]byte(resp), &response)
	return response
}

func (ss *SiakadService) GetDospem(phoneNumber string) DospemResponse {
	var response DospemResponse

	log.Println("phone number : ")
	log.Println(phoneNumber)

	uri := ss.uri + "/counselors"
	log.Println("URL : ", uri)

	log.Println("ap_key")
	log.Println(Env.Get("siakad_app_key"))

	_, resp, errorResponse := gorequest.New().
		Get(uri).
		Set("app-key", Env.Get("siakad_app_key")).
		Set("client-key", phoneNumber).
		End()

	log.Println("Response : ")
	log.Println(resp)

	if errorResponse != nil {
		log.Println(errorResponse)
	}

	json.Unmarshal([]byte(resp), &response)
	return response
}

func (ss *SiakadService) GetSpp(phoneNumber string) SppResponse {
	var response SppResponse

	log.Println("phone number : ")
	log.Println(phoneNumber)

	uri := ss.uri + "/payment/status"
	log.Println("URL : ", uri)

	log.Println("ap_key")
	log.Println(Env.Get("siakad_app_key"))

	_, resp, errorResponse := gorequest.New().
		Get(uri).
		Set("app-key", Env.Get("siakad_app_key")).
		Set("client-key", phoneNumber).
		End()

	log.Println("Response : ")
	log.Println(resp)

	if errorResponse != nil {
		log.Println(errorResponse)
	}

	json.Unmarshal([]byte(resp), &response)
	return response
}

func NewSiakadService() *SiakadService {
	var service SiakadService
	service.uri = Env.Get("siakad_url_host")
	return &service
}
