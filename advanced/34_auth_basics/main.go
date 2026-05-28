package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flutterffi/PFGolang/pkg/authx"
)

func main() {
	fmt.Println("Lesson 34: Auth Basics")

	request := httptest.NewRequest(http.MethodGet, "/private", nil)
	request.Header.Set("Authorization", "Bearer training-token")

	if err := authx.ValidateStaticToken(request, "training-token"); err != nil {
		fmt.Printf("auth failed: %v\n", err)
		return
	}

	fmt.Println("request authorized")

	badRequest := httptest.NewRequest(http.MethodGet, "/private", nil)
	if err := authx.ValidateStaticToken(badRequest, "training-token"); err != nil {
		fmt.Printf("second request rejected as expected: %v\n", err)
	}
}
