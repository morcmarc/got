package got

import (
	"fmt"
	"os/exec"
	"strings"
)

// SetRemote will attempt to update the remote's url with a new one. Returns an
// error if the remote doesn't exist.
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

// GetRemote will return the url of the remote server. Returns an error if the
// remote does not exist.
func (this *GotClient) GetRemote(name string) (string, error) {
	this.changeDirToPath()

	cmd := exec.Command("git", "config", "--get", fmt.Sprintf("remote.%s.url", name))
	out, err := cmd.CombinedOutput()

	if this.Verbose {
		fmt.Println(string(out))
	}

	this.restorePath()

	return strings.TrimSpace(string(out)), err
}

// AddRemote will add a new remote server to the repository. Use SetRemote if
// you'd like to update a remote; AddRemote will return an error if the remote
// is already set for the repository.
func (this *GotClient) AddRemote(name, uri string) error {
	this.changeDirToPath()

	cmd := exec.Command("git", "remote", "add", name, uri)
	out, err := cmd.CombinedOutput()

	if this.Verbose {
		fmt.Println(string(out))
	}

	this.restorePath()

	return err
}
