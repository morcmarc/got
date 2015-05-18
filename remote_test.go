package got

import (
	"os/exec"
	"testing"
)

func TestSetRemote(t *testing.T) {
	gc, path := setup()
	defer teardown(path)

	gc.Init()

	gc.changeDirToPath()
	cmd := exec.Command("git", "remote", "add", "origin", "git@example.com:user/repo1")
	_, err := cmd.CombinedOutput()
	gc.restorePath()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	gc.SetRemote("origin", "git@example.com:user/repo2")

	gc.changeDirToPath()
	cmd = exec.Command("git", "remote", "-v")
	out, err := cmd.CombinedOutput()
	gc.restorePath()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	expected := "origin\tgit@example.com:user/repo2 (fetch)\n" +
		"origin\tgit@example.com:user/repo2 (push)\n"
	if string(out) != expected {
		t.Errorf("Was expecting:\n%s\ngot:\n%s", expected, out)
	}
}
