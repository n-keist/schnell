package runner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/n-keist/schnell/internal/model"
)

func RunProgram(p model.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	progColor := color.New(p.Color).SprintFunc()

	args := strings.Split(p.Cmd, " ")
	if len(args) < 1 {
		fmt.Printf("App %s has no command\n", progColor(p.Label))
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = p.Path
	cmd.Env = append(os.Environ(), "FORCE_COLOR=1")
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	reader := io.MultiReader(stdoutPipe, stderrPipe)

	if err := cmd.Start(); err != nil {
		fmt.Printf("%s failed to start. \n\t%s\n", progColor(p.Label), err.Error())
		return
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf(
			"%s %s \n",
			progColor(fmt.Sprintf("%s >>", p.Label)),
			line,
		)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("%s scan error \n\t%s\n", progColor(p.Label), err.Error())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "%s exited \n\t%s\n", progColor(p.Label), err)
	} else {
		fmt.Printf("goodbye %s", progColor(p.Label))
	}
}
