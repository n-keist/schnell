package model

type (
	Workspace struct {
		Name     string    `json:"name"`
		Programs []Program `json:"programs"`
	}
)
