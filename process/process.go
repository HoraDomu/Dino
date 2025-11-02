//process handling

package process

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// RunProcess launches a command as a new process
func RunProcess(command string, wg *sync.WaitGroup) {
	defer wg.Done()

	parts := strings.Fields(command)
	if len(parts) == 0 {
		fmt.Println("[Dino] No command provided")
		return
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("[Dino] Starting:", command)
	if err := cmd.Run(); err != nil {
		fmt.Println("[Dino] Error running", command, ":", err)
	} else {
		fmt.Println("[Dino] Finished:", command)
	}
}
