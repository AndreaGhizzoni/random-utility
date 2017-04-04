package out

import (
	"os"
	"testing"
)

func TestOpenIfCanRW(t *testing.T) {

}

func TestArgumentsOpenIfCanRW(t *testing.T) {
	noRead := "no-read.txt"
	nr, err := os.Create(noRead)
	if err != nil {
		t.Fatal(err)
	}
	nr.Chmod(0222) // this sets permissions to --w--w--w-
	defer os.Remove(noRead)
	if _, err := openFileIfCanRW("", noRead); err == nil {
		t.Fatal("With no-read permission, openFileIfCanRW() must return an error")
	}

	noWrite := "no-write.txt"
	nw, err := os.Create(noWrite)
	if err != nil {
		t.Fatal(err)
	}
	nw.Chmod(0444) // this sets permissions to -r--r--r--
	defer os.Remove(noWrite)
	if _, err := openFileIfCanRW("", noWrite); err == nil {
		t.Fatal("With no-write permission, openFileIfCanRW() must return an error")
	}

	aDir := "aDir"
	if err := os.MkdirAll(aDir, 0777); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(aDir)
	if _, err := openFileIfCanRW("", aDir); err == nil {
		t.Fatal("With a directory, openFileIfCanRW() must return an error")
	}
}
