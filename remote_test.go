package got

import (
	"testing"
)

func TestRemote(t *testing.T) {
	gc, path := setup()
	defer teardown(path)

	gc.Init()

	err := gc.AddRemote("origin", "git@example.com:user/repo1")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	expected := "git@example.com:user/repo2"
	err = gc.SetRemote("origin", expected)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	rem, err := gc.GetRemote("origin")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if rem != expected {
		t.Errorf("Was expecting:\n%s\ngot:\n%s", expected, rem)
	}
}
