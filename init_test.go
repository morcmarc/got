package got

import (
	"os/exec"
	"testing"
)

func TestInit(t *testing.T) {
	gc, path := setup()
	defer teardown(path)

	err := gc.Init()
	if err != nil {
		t.Fatalf("Wasn't expecting error %s", err)
	}

	cmd := exec.Command("git", "status")
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Was expecting a git repo got error %s", err)
	}
}
