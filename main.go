package main

import (
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/verocity-gaming/rcon"
)

type handler struct {
	*rcon.Conn
}

func authorized(header string) bool {
	return true
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !authorized(r.Header.Get("Auth")) {
		w.WriteHeader(http.StatusUnauthorized)

		_, err := w.Write([]byte("invalid auth header provided"))
		if err != nil {
			log.Error().Err(err).Msg("failed to send unauthorized response")
		}

		return
	}

	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.post(w, r)
	}
}

// clean will verify the url path starts with "/", and return a clean url path.
func clean(r *http.Request) string {
	p := r.URL.Path
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
		r.URL.Path = p
	}

	return path.Clean(p)
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	p := clean(r)

	w.WriteHeader(http.StatusOK)

	response, err := h.Send("get", strings.Trim(p, "/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = err.Error()
	}

	_, err = w.Write([]byte(response))
	if err != nil {
		log.Error().Err(err).Str("route", p).Msg("failed to send response")
	}
}

func (h *handler) post(w http.ResponseWriter, r *http.Request) {
	switch clean(r) {

	}
}

func main() {
	err := parse()
	if err != nil {
		panic(err)
	}

	h := &handler{}

	h.Conn, err = rcon.New(addr, password)
	if err != nil {
		panic(err)
	}

	if certificate == "" && key == "" {
		err := http.ListenAndServe(":"+port, h)
		if err != nil {
			shutdown(err)
		}
	} else {
		err := http.ListenAndServeTLS(":"+port, certificate, key, h)
		if err != nil {
			shutdown(err)
		}
	}
}

func shutdown(err error) {
	log.Error().Err(err).Msg("shutting down")
	time.Sleep(1000)
}
