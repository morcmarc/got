package got

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	path string
	gc   *GotClient
)

func setup() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path, err = ioutil.TempDir(wd, "gottest")
	if err != nil {
		panic(err)
	}

	gc = NewGotClient(path)
}

func teardown() {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func TestChangeDirToPath(t *testing.T) {
	setup()
	defer teardown()

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
