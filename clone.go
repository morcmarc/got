package got

import (
	"fmt"
	"os/exec"
)

// Clone will `git clone --recursive` the repository given by its uri.
func (g *GotClient) Clone(uri string) error {
	cmd := exec.Command("git", "clone", "--recursive", uri, g.Path)
	out, err := cmd.CombinedOutput()

	if g.Verbose {
		fmt.Println(string(out))
	}

	return err
}
