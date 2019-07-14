package main

import (
	"fmt"
	"net/http"
	"skripsi/app"
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
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Post from webhook! r.PostFrom = %v\n", r.PostForm)

		request.From = r.PostFormValue("form")
		request.To = r.PostFormValue("to")
		request.Event = r.PostFormValue("event")
		request.Text = r.PostFormValue("text")

		app.NewUseCase().Handle(request)
		return

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
