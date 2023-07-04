package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func sessionsEndpoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		reqBody, _ := ioutil.ReadAll(r.Body)
		var session Session
		json.Unmarshal(reqBody, &session)
		storeSession(session)
		fmt.Printf("%+v\n", session)
	case "GET":
		sessions := []Session{}
		sessions, err := GetSessions()
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(sessions)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"not found"}`))
	}
}

func main() {
	pr := mux.NewRouter()

	pr.Handle("/swagger.yml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "swagger.yml"}
	sh := middleware.SwaggerUI(opts, nil)
	pr.Handle("/docs", sh)

	pr.Handle("/api/v1/sessions", BasicAuth(sessionsEndpoints))
	log.Fatal(http.ListenAndServe(":8010", pr))

	defer disconnectFromMongo()
}
