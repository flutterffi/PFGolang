package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

type healthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		payload := healthResponse{
			Status: "ok",
			Time:   time.Now().Format(time.RFC3339),
		}

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return mux
}

func main() {
	fmt.Println("Lesson 22: Real HTTP Server")

	mux := newMux()

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)

	fmt.Printf("Self-check status: %d\n", recorder.Code)
	fmt.Printf("Self-check body: %s", recorder.Body.String())
	fmt.Println("Starting local server at http://127.0.0.1:8080")

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
		close(errCh)
	}()

	time.Sleep(80 * time.Millisecond)

	response, err := http.Get("http://127.0.0.1:8080/health")
	if err != nil {
		fmt.Printf("Live request could not run in this environment: %v\n", err)
		fmt.Println("The handler still passed the in-memory self-check above.")
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read response error: %v\n", err)
		return
	}

	fmt.Printf("Live request status: %d\n", response.StatusCode)
	fmt.Printf("Live request body: %s", string(body))

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("shutdown error: %v\n", err)
		return
	}

	if err, ok := <-errCh; ok && err != nil {
		fmt.Printf("server error: %v\n", err)
		return
	}

	fmt.Println("Server shutdown complete.")
}
