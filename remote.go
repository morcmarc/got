package got

import (
	"fmt"
	"os/exec"
	"strings"
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
