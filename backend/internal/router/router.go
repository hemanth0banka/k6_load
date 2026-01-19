package router

import (
	"net/http"

	"k6clone/internal/handlers"
	"k6clone/internal/generator"
	"k6clone/internal/repository"
	"k6clone/internal/service"
)

func NewRouter(
	scriptService *service.ScriptService,
	testService *service.TestService,
	historyRepo repository.TestResultRepository,
	k6Gen generator.K6JSGenerator,
) *http.ServeMux {

	mux := http.NewServeMux()

	scriptHandler := handlers.NewScriptHandler(scriptService, k6Gen)
	testHandler := handlers.NewTestHandler(testService)
	historyHandler := handlers.NewHistoryHandler(historyRepo)

	mux.HandleFunc("/scripts", scriptHandler.HandleScripts)
	mux.HandleFunc("/scripts/", scriptHandler.GetScriptByID)
	mux.HandleFunc("/scripts/k6", scriptHandler.GetK6Script)
	mux.HandleFunc("/tests/run", testHandler.RunTest)
	mux.HandleFunc("/history", historyHandler.GetHistory)
	mux.HandleFunc("/health", health)

	return mux
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
