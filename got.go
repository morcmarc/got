// Package got provides basic Git commands for managing Git repositories.
package got

import (
	"os"
	"os/exec"
)

type GotClient struct {
	Path    string
	Verbose bool
}

func IsGitInstalled() bool {
	cmd := exec.Command("git", "--version")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	} else {
		return true
	}
}

func NewGotClient(path string) *GotClient {
	return &GotClient{
		Verbose: false,
		Path:    path,
	}
}

var (
	workDir string
)

func (g *GotClient) changeDirToPath() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	workDir = wd

	if err := os.Chdir(g.Path); err != nil {
		panic(err)
	}
}

func (g *GotClient) restorePath() {
	if err := os.Chdir(workDir); err != nil {
		panic(err)
	}
}
