package openapix

type Spec struct {
	OpenAPI string          `json:"openapi"`
	Info    Info            `json:"info"`
	Paths   map[string]Path `json:"paths"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type Path struct {
	Get  *Operation `json:"get,omitempty"`
	Post *Operation `json:"post,omitempty"`
}

type Operation struct {
	Summary string `json:"summary"`
}

func TodoSpec() Spec {
	return Spec{
		OpenAPI: "3.1.0",
		Info: Info{
			Title:   "Todo Service",
			Version: "1.0.0",
		},
		Paths: map[string]Path{
			"/tasks": {
				Get:  &Operation{Summary: "List tasks"},
				Post: &Operation{Summary: "Create a task"},
			},
			"/tasks/complete": {
				Post: &Operation{Summary: "Complete a task"},
			},
		},
	}
}
