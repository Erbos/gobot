package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Yellouuu!!!")
}

func VerificationEndpoint(w http.ResponseWriter, r *http.Request)  {
  challange := r.URL.Query().Get("hub.challenge")
  token := r.URL.Query().Get("hub.verify_token")

  if token == os.Getenv("VERIFY_TOKEN"){
    w.WriteHeader(200)
    w.Write(([]byte(challange)))
  } else {
    w.WriteHeader(404)
    w.Write([]byte("Error, wrong validation token"))
  }
}

func MessagesEndpoint(w http.ResponseWriter, r *http.Request)  {
  var callback Callback
  json.NewDecoder(r.Body).Decode(&amp;callback)
  if callback.Object == "page"
  for _, entry := range callback.Entry{
    for _, event := range entry.Messaging{
      ProccessMessage(event)
    }
  }
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeEndPoint)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
