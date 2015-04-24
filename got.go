// Package got provides basic Git commands for managing Git repositories.
package got

import (
	"os"
	"os/exec"
)

type GotClient struct {
	Path    string // Path to the repository
	Verbose bool   // Verbose output flag
}

// IsGitInstalled will run `git --version` to determine whether Git is installed
// on the target system and return true or false.
func IsGitInstalled() bool {
	cmd := exec.Command("git", "--version")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	} else {
		return true
	}
}

// NewGotClient returns a instance of GotClient initialised for the given path.
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
