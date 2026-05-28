package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flutterffi/PFGolang/pkg/cachex"
	"github.com/flutterffi/PFGolang/pkg/flagx"
	"github.com/flutterffi/PFGolang/pkg/openapix"
	"github.com/flutterffi/PFGolang/pkg/queuex"
	"github.com/flutterffi/PFGolang/pkg/todoservice"
)

func perform(handler http.Handler, method, path string) {
	request := httptest.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	fmt.Printf("%s %s -> status=%d cache=%s\n", method, path, recorder.Code, recorder.Header().Get("X-Cache"))
	if recorder.Body.Len() > 0 {
		fmt.Printf("body=%s\n", recorder.Body.String())
	}
}

func main() {
	fmt.Println("Lesson 48: Production-Style Todo Service")

	cache := cachex.New()
	flags := flagx.New(map[string]bool{
		"publish_task_events": true,
	})
	publisher := &queuex.MemoryPublisher{}
	service := todoservice.New(cache, flags, publisher)
	handler := service.Routes()

	perform(handler, http.MethodGet, "/tasks")
	perform(handler, http.MethodGet, "/tasks")
	perform(handler, http.MethodPost, "/tasks?title=Ship%20a%20practice%20service")
	perform(handler, http.MethodPost, "/tasks/complete?id=1")
	perform(handler, http.MethodGet, "/tasks")

	spec := openapix.TodoSpec()
	fmt.Printf("openapi version: %s\n", spec.OpenAPI)
	fmt.Printf("published events: %v\n", publisher.Messages)
}
