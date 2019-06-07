package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weeny/application"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	r      *mux.Router
	server *http.Server
	app    *application.Application
}

type response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func respondError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	respond(w, message, "")
}
func respond(w io.Writer, msg, data string) {
	response := response{
		Message: msg,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true

}

func (api *ApiServer) shotern(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&payload)
	if valid := isValidURL(payload.URL); !valid {
		respondError(w, "Not a valid URL")
		return
	}

	urlHash, err := api.app.Save(payload.URL)
	if err != nil {
		respondError(w, "Failure")
		return
	}

	respond(w, "Success", urlHash)

}

func (api *ApiServer) redirect(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	url, err := api.app.Get(hash)
	if err != nil {
		fmt.Printf("Error : %v \n", err)
		respondError(w, "Failure")
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	respond(w, "Success", url)

}

func (api *ApiServer) lookup(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	url, err := api.app.Get(hash)
	if err != nil {
		fmt.Printf("Error : %v \n", err)
		respondError(w, "Failure")
		return
	}
	respond(w, "Success", url)
}

func NewServer(app *application.Application) *ApiServer {
	return &ApiServer{
		r:   mux.NewRouter(),
		app: app,
	}
}

func (api *ApiServer) Start(host string, port int) error {
	api.r.HandleFunc("/ping", ping).Methods("GET")
	api.r.HandleFunc("/shortern", api.shotern).Methods("POST")
	api.r.HandleFunc("/{hash}", api.redirect).Methods("GET")
	api.r.HandleFunc("/lookup/{hash}", api.lookup).Methods("GET")
	fmt.Println("Starting the server... ")
	addr := fmt.Sprintf("%s:%d", host, port)
	api.server = &http.Server{
		Addr:    addr,
		Handler: api.r,
	}
	return api.server.ListenAndServe()
}
