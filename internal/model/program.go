package model

import "github.com/fatih/color"

type (
	Program struct {
		Label string `json:"label"`
		Path  string `json:"path"`
		Cmd   string `json:"cmd"`
		Color color.Attribute
	}
)
