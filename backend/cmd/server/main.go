package main

import (
	"fmt"
	"net/http"
	"k6clone/internal/router"
	"k6clone/internal/engine"
	"k6clone/internal/generator"
	"k6clone/internal/middleware"
	"k6clone/internal/repository"
	"k6clone/internal/service"
)

func main() {
	fmt.Println(" Starting K6 Load Testing ...")

	httpGen := generator.NewHttpGenerator()
	k6JSGen := generator.NewK6JSGenerator()

	scriptRepo := repository.NewFileScriptRepository("./scripts")
	historyRepo := repository.NewFileTestResultRepository("./scripts/results")

	loadEngine := engine.NewLoadEngine()

	scriptService := service.NewScriptService(httpGen, scriptRepo)
	testService := service.NewTestService(scriptRepo, historyRepo, loadEngine)

	mux := router.NewRouter(
		scriptService,
		testService,
		historyRepo,
		*k6JSGen,
	)

	fmt.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", middleware.CORSMiddleware(mux))
}