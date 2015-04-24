package got

import (
	"fmt"
	"os/exec"
)

// Init will initialise a new Git repository at the current location. The
// location be set using Go's Chdir method.
func (g *GotClient) Init() error {
	g.changeDirToPath()

	cmd := exec.Command("git", "init")
	out, err := cmd.CombinedOutput()
	if g.Verbose {
		fmt.Println(string(out))
	}

	g.restorePath()

	return err
}
