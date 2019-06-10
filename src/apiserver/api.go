package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weeny/application"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ApiServer struct {
	r      *mux.Router
	server *http.Server
	app    *application.Application
	logger zerolog.Logger
}

type response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func NewServer(app *application.Application) *ApiServer {
	return &ApiServer{
		r:      mux.NewRouter(),
		app:    app,
		logger: log.With().Str("component", "server").Logger(),
	}
}

func (api *ApiServer) Start(port int) error {

	api.r.HandleFunc("/ping", api.ping).Methods("GET")
	api.r.HandleFunc("/shortern", api.shotern).Methods("POST")
	api.r.HandleFunc("/{hash}", api.redirect).Methods("GET")
	api.r.HandleFunc("/lookup/{hash}", api.lookup).Methods("GET")

	addr := fmt.Sprintf(":%d", port)
	api.logger.Info().Msgf("Starting the server, binding to: `%s`", addr)

	api.server = &http.Server{
		Addr:    addr,
		Handler: api.r,
	}
	return api.server.ListenAndServe()
}

func (api *ApiServer) Shutdown() error {
	// TODO: use shutdown and a context
	return api.server.Close()
}

func (api *ApiServer) ping(w http.ResponseWriter, r *http.Request) {
	api.logger.Debug().Msg("ping received")
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
	return err == nil
}

func (api *ApiServer) shotern(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		URL string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		api.logger.Warn().Msgf("decoding payload: %#v, err: %s", payload, err)
		respondError(w, "Invalid payload")
		return
	}

	if valid := isValidURL(payload.URL); !valid {
		api.logger.Warn().Msgf("validating URL: %s", payload.URL)
		respondError(w, "Not a valid URL")
		return
	}

	urlHash, err := api.app.Save(payload.URL)
	if err != nil {
		api.logger.Error().Msgf("saving hash: %s, URL: %s, err: %s", urlHash, payload.URL, err)
		respondError(w, "Failure")
		return
	}

	respond(w, "Success", urlHash)

}

func (api *ApiServer) redirect(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	url, err := api.app.Get(hash)
	if err != nil {
		api.logger.Error().Msgf("redirect hash: %s, err: %s", hash, err)
		respondError(w, "Failure")
		return
	}

	api.logger.Debug().Msgf("redirecting hash: %s, url: %s", hash, url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	respond(w, "Success", url)

}

func (api *ApiServer) lookup(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	url, err := api.app.Get(hash)
	if err != nil {
		api.logger.Error().Msgf("lookup hash: %s, err: %s", hash, err)
		respondError(w, "Failure")
		return
	}
	respond(w, "Success", url)
}
