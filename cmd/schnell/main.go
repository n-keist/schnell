package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/n-keist/schnell/internal/model"
	"github.com/n-keist/schnell/internal/runner"
)

func main() {
	workspaceFile, err := os.ReadFile("workspace.json")
	if err != nil {
		slog.Error("workspace.json error", "error", err.Error())
		return
	}

	var workspace model.Workspace
	if err := json.Unmarshal(workspaceFile, &workspace); err != nil {
		slog.Error("could not decode workspace config", "error", err.Error())
		return
	}

	var wg sync.WaitGroup
	for _, p := range workspace.Programs {
		wg.Add(1)
		go runner.RunProgram(p, &wg)
	}

	wg.Wait()
	fmt.Println("All programs exited.")
}
