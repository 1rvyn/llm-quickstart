package models

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Labels struct {
	Labels []Label `json:"labels"`
}
