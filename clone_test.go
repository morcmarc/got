package got

import (
	"testing"
)

func TestCloneWithInvalidUrl(t *testing.T) {
	setup()
	defer teardown()

	err := gc.Clone("invalid url")
	if err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestCloneWithValidUrl(t *testing.T) {
	setup()
	// Remove tmp dir immediately to prevent git error during clone
	teardown()

	err := gc.Clone("git@github.com:morcmarc/got.git")
	if err != nil {
		t.Fatalf("Wasn't expecting error %s", err)
	}
}
