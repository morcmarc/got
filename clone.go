package got

import (
	"fmt"
	"os/exec"
)

// Clone will recursively `git clone` the repository at the given uri into
// the destination path.
func (g *GotClient) Clone(uri string) error {
	cmd := exec.Command("git", "clone", "--recursive", uri, g.Path)
	out, err := cmd.CombinedOutput()

	if g.Verbose {
		fmt.Println(string(out))
	}

	return err
}
