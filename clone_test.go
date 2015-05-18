package got

import (
	"os/exec"
	"testing"
)

func TestCloneWithInvalidUrl(t *testing.T) {
	gc, path := setup()
	defer teardown(path)

	err := gc.Clone("invalid url")
	if err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestCloneWithValidUrl(t *testing.T) {
	gc, path := setup()
	// Remove tmp dir immediately to prevent git error during clone
	teardown(path)

	err := gc.Clone("git@github.com:morcmarc/got.git")
	if err != nil {
		t.Fatalf("Wasn't expecting error %s", err)
	}

	cmd := exec.Command("git", "status")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Was expecting a git repo got error %s", err)
	}

	defer teardown(path)
}
