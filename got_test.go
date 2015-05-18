package got

import (
	"io/ioutil"
	"os"
	"testing"
)

func setup() (*GotClient, string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path, err := ioutil.TempDir(wd, "gottest")
	if err != nil {
		panic(err)
	}

	return NewGotClient(path), path
}

func teardown(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

func TestChangeDirToPath(t *testing.T) {
	gc, path := setup()
	defer teardown(path)

	owd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	gc.changeDirToPath()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if wd != path {
		t.Fatalf("Was expecting working dir to be %s, got %s", path, wd)
	}

	gc.restorePath()

	wd, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	if owd != wd {
		t.Fatalf("Was expecting working dir to be %s, got %s", owd, wd)
	}
}
