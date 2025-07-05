package main

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/n-keist/schnell/internal/model"
	"github.com/n-keist/schnell/internal/runner"
)

func main() {
	progs := []model.Program{
		{
			Label: "frontend",
			Path:  "../ablage/vue-ablage",
			Cmd:   "npm run dev",
			Color: color.FgGreen,
		},
		{
			Label: "backend",
			Path:  "../ablage/go-ablage/cmd/api-server",
			Cmd:   "go run .",
			Color: color.FgBlue,
		},
	}

	var wg sync.WaitGroup
	for _, p := range progs {
		wg.Add(1)
		go runner.RunProgram(p, &wg)
	}

	wg.Wait()
	fmt.Println("All programs exited.")
}
