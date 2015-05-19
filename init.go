package got

import (
	"fmt"
	"os/exec"
)

// Init will try to initialise a new Git repository at GotClient.Path.
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
