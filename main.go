package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"skripsip/app"
)

var (
	Env = app.Env
)

func main() {
	http.HandleFunc("/chatbot", index)

	fmt.Println("starting web server at" + Env.Get("listen_port"))
	http.ListenAndServe(":"+Env.Get("listen_port"), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var request app.WhatsappChatRequest

	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Sorry, only member , contact administrator.")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			log.Printf("ParseForm() err: %v", err)
			return
		}

		log.Printf("Post from webhook! r.PostFrom = %v\n", r.PostForm)

		reqFromWebhook := r.FormValue("data")

		json.Unmarshal([]byte(reqFromWebhook), &request)

		msg := app.NewUseCase().Handle(request)

		jsonInBytes, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("response : ")
		log.Println(string(jsonInBytes))

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)

		return

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
