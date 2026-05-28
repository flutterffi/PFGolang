package main

import (
	"context"
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/metricsx"
)

type traceKey string

const requestIDKey traceKey = "request_id"

func handleRequest(ctx context.Context, counters *metricsx.Counter) {
	requestID, _ := ctx.Value(requestIDKey).(string)
	counters.Add("requests_total", 1)
	counters.Add("requests_success", 1)
	fmt.Printf("trace request_id=%s event=handle_request status=success\n", requestID)
}

func main() {
	fmt.Println("Lesson 35: Metrics and Tracing Shape")

	counters := metricsx.NewCounter()
	ctx := context.WithValue(context.Background(), requestIDKey, "req-001")

	handleRequest(ctx, counters)
	handleRequest(context.WithValue(context.Background(), requestIDKey, "req-002"), counters)

	fmt.Printf("metrics snapshot: %+v\n", counters.Snapshot())
	fmt.Println("This is not a full tracing system, but it shows the data shape and flow.")
}
