package got

import (
	"fmt"
	"os/exec"
)

func (this *GotClient) SetRemote(name, uri string) error {
	this.changeDirToPath()

	cmd := exec.Command("git", "remote", "set-url", name, uri)
	out, err := cmd.CombinedOutput()

	if this.Verbose {
		fmt.Println(string(out))
	}

	this.restorePath()

	return err
}
