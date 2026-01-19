package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"k6clone/internal/generator"
	"k6clone/internal/model"
	"k6clone/internal/service"
)

type ScriptHandler struct {
	service *service.ScriptService
	k6Gen   generator.K6JSGenerator
}

func NewScriptHandler(
	s *service.ScriptService,
	gen generator.K6JSGenerator,
) *ScriptHandler {
	return &ScriptHandler{
		service: s,
		k6Gen:   gen,
	}
}

func (h *ScriptHandler) HandleScripts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAllScripts(w, r)
	case http.MethodPost:
		h.CreateScript(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ScriptHandler) CreateScript(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	script, err := h.service.CreateFromURL(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(script)
}

func (h *ScriptHandler) GetAllScripts(w http.ResponseWriter, r *http.Request) {
	scripts, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if scripts == nil {
		scripts = []*model.Script{}
	}

	json.NewEncoder(w).Encode(scripts)
}

func (h *ScriptHandler) GetScriptByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/scripts/")
	if id == "" {
		http.Error(w, "script id required", http.StatusBadRequest)
		return
	}

	script, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "script not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(script)
}

func (h *ScriptHandler) GetK6Script(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "script id required", http.StatusBadRequest)
		return
	}

	script, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "script not found", http.StatusNotFound)
		return
	}

	code, err := h.k6Gen.Generate(&generator.K6JSInput{
		Script: script,
		Config: model.TestConfig{
			VUs:      10,
			Duration: 30,
		},
	})
	if err != nil {
		http.Error(w, "generation failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(code))
}

