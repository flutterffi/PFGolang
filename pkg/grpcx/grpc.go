package grpcx

type TaskRequest struct {
	TenantID string
	Title    string
}

type TaskResponse struct {
	ID      int
	Message string
}

func SimulateCreate(req TaskRequest) TaskResponse {
	return TaskResponse{
		ID:      101,
		Message: "task created for tenant " + req.TenantID + ": " + req.Title,
	}
}
