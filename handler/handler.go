package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ggerritsen/building-blocks-app/model"
)

type docService interface {
	Store(name string) (*model.Document, error)
	Retrieve(id int) (*model.Document, error)
}

type handler struct {
	s docService
}

// NewHandler creates a new http handler to serve model.Document objects
func NewHandler(s docService) *handler {
	return &handler{s}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.readDoc(w, r)
		return
	}

	if r.Method == http.MethodPost {
		h.saveDoc(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

var p = regexp.MustCompile("/documents/([0-9]+)")

func (h *handler) readDoc(w http.ResponseWriter, r *http.Request) {
	m := p.FindStringSubmatch(r.URL.Path)

	// no match found (m[0] matches the whole string)
	if len(m) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i, err := strconv.ParseInt(m[1], 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := int(i)

	doc, err := h.s.Retrieve(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if doc == nil {
		http.NotFound(w, r)
		return
	}

	if err := json.NewEncoder(w).Encode(doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) saveDoc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "No name provided", http.StatusBadRequest)
		return
	}

	doc, err := h.s.Store(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
