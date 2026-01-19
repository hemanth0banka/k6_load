package handlers

import (
	"encoding/json"
	"net/http"

	"k6clone/internal/repository"
)

type HistoryHandler struct {
	repo repository.TestResultRepository
}

func NewHistoryHandler(r repository.TestResultRepository) *HistoryHandler {
	return &HistoryHandler{repo: r}
}

func (h *HistoryHandler) GetHistory(w http.ResponseWriter, _ *http.Request) {
	results := h.repo.FindAll()

	json.NewEncoder(w).Encode(results)
}
