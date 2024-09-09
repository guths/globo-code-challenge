package morse

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func HandleFile(filepath string, writer io.Writer) error {
	if writer.(*os.File) == nil {
		writer = os.Stdout
	}

	file, err := os.Open(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	Handle(file, writer)

	return nil
}

func HandleStdin(text string, writer io.Writer) error {
	if writer.(*os.File) == nil {
		writer = os.Stdout
	}

	Handle(strings.NewReader(text), writer)
	return nil
}

func customSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	delimiter := []byte("   ")

	if i := bytes.Index(data, delimiter); i >= 0 {
		return i + len(delimiter), data[:i], nil
	}

	if atEOF {
		if len(data) > 0 {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	return 0, nil, nil
}

func Handle(reader io.Reader, w io.Writer) error {
	m := NewMorser()

	scanner := bufio.NewScanner(reader)

	scanner.Split(customSplitter)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 0 {
			word := m.Decode(word)
			w.Write([]byte(word))
			w.Write([]byte(" "))
		}
	}

	return nil
}
