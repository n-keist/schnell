package model

import "github.com/fatih/color"

type (
	Program struct {
		Label string
		Path  string
		Cmd   string
		Color color.Attribute
	}
)
