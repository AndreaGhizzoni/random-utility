package out

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIsDirectory(t *testing.T) {

}

func TestArgumentsIsDirectory(t *testing.T) {

}

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
	if _, err := openIfCanRW(noRead); err == nil {
		t.Fatal("With no-read permission openIfCanRW() must return an error")
	}

	noWrite := "no-write.txt"
	nw, err := os.Create(noWrite)
	if err != nil {
		t.Fatal(err)
	}
	nw.Chmod(0444) // this sets permissions to -r--r--r--
	defer os.Remove(noWrite)
	if _, err := openIfCanRW(noWrite); err == nil {
		t.Fatal("With no-write permission openIfCanRW() must return an error")
	}

	dirPath, err := ioutil.TempDir("", "aDir")
	defer os.Remove(dirPath)
	if _, err := openIfCanRW(dirPath); err == nil {
		t.Fatal("With a directory openIfCanRW() must return an error")
	}
}
