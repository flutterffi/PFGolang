package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flutterffi/PFGolang/pkg/deploytodo"
	"github.com/flutterffi/PFGolang/pkg/jobx"
	"github.com/flutterffi/PFGolang/pkg/lockx"
)

func perform(handler http.Handler, method, path, tenant string) {
	request := httptest.NewRequest(method, path, nil)
	if tenant != "" {
		request.Header.Set("X-Tenant-ID", tenant)
	}
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	fmt.Printf("%s %s tenant=%s -> status=%d\n", method, path, tenant, recorder.Code)
	if recorder.Body.Len() > 0 {
		fmt.Printf("body=%s\n", recorder.Body.String())
	}
}

func main() {
	fmt.Println("Lesson 54: Deployment-Ready Todo Service")

	service := deploytodo.New(jobx.NewRunner(), lockx.New())
	handler := service.Routes()

	perform(handler, http.MethodGet, "/tasks?page=1&page_size=2", "team-a")
	perform(handler, http.MethodPost, "/tasks?title=Ship%20tenant-aware%20service", "team-a")
	perform(handler, http.MethodGet, "/tasks?page=1&page_size=5", "team-a")
	perform(handler, http.MethodPost, "/maintenance/run", "")

	fmt.Println("Next upgrades: add health endpoints, env config, auth, and persistent storage.")
}
