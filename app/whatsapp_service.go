package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type WhatsappService struct {
	uri string
}

const (
	clientSecret string = "9ca19353e1fb861f6d3aed8af6803c0b"
	accessToken  string = "d42cd44868d2fe953329677aa1c71e1d80f0f4d7"
	siteHost     string = "http://coop.apps.knpuniversity.com"
	userID       string = "460"
)

func (ss *WhatsappService) CheckPhoneNumber() {
	client := &http.Client{}

	data := url.Values{}
	data.Set("client_id", `Lazy Test`)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/token", siteHost), bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(f))
}

func NewWhatsappService() *WhatsappService {
	var whatsapp WhatsappService
	whatsapp.uri = Env.Get("whatsapp_url_host")
	return &whatsapp
}
