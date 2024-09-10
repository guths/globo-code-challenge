package morse_test

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/guths/globo-code-challenge/pkg/morse"
)

func captureOutput(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out), err
}

func TestHandleFileWithInvalidPathFile(t *testing.T) {
	err := morse.HandleFile("", os.Stdout)

	if err == nil {
		t.Error("test should fail with nil path")
	}
}

func TestHandleFileWithValidPathFile(t *testing.T) {
	output, err := captureOutput(func() error {
		f, _ := os.Create("test2.txt")

		defer os.Remove("test2.txt")

		f.WriteString(".... . .-.. .-.. ---")

		var nf *os.File

		err := morse.HandleFile("test2.txt", nf)

		return err
	})

	if err != nil {
		t.Fail()
	}

	if output != "HELLO " {
		t.Errorf("invalid output" + output)
	}
}

func TestHandleFileWithValidPathFileAndOutputfile(t *testing.T) {
	f, _ := os.Create("test.txt")

	defer os.Remove("test.txt")

	f.WriteString(".... . .-.. .-.. ---   .-- --- .-. .-.. -..")

	outputfile, _ := os.Create("outputfile1.txt")

	defer os.Remove("outputfile1.txt")

	err := morse.HandleFile("test.txt", outputfile)

	if err != nil {
		t.Error(err)
	}

	var content string

	openFile, _ := os.Open("outputfile1.txt")

	scanner := bufio.NewScanner(openFile)

	for scanner.Scan() {
		content += scanner.Text()
	}

	if content != "HELLO WORLD " {
		t.Error(content)
	}
}

func TestHandleStdinWithStdout(t *testing.T) {
	output, err := captureOutput(func() error {
		var nf *os.File
		return morse.HandleStdin(".... . .-.. .-.. ---   .-- --- .-. .-.. -..", nf)
	})

	if err != nil {
		t.Fail()
	}

	if output != "HELLO WORLD " {
		t.Errorf("invalid output" + output)
	}
}

func TestHandleStdinWithOutputFile(t *testing.T) {
	outputfile, _ := os.Create("outputfile.txt")

	err := morse.HandleStdin(".... . .-.. .-.. ---   .-- --- .-. .-.. -..", outputfile)

	if err != nil {
		t.Fail()
	}

	var content string

	openFile, _ := os.Open("outputfile.txt")

	defer os.Remove("outputfile.txt")
	scanner := bufio.NewScanner(openFile)

	for scanner.Scan() {
		content += scanner.Text()
	}

	if content != "HELLO WORLD " {
		t.Fail()
	}
}
