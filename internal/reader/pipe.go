package reader

import (
	"bufio"
	"errors"
	"os"
)

// Pipe Reader
func Pipe() (*bufio.Reader, error) {
	info, err := os.Stdin.Stat()

	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		return nil, errors.New("This version only accept data from pipe")
	}

	return bufio.NewReader(os.Stdin), nil
}
