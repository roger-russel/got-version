package reader

import (
	"bufio"
	"os"
)

// Pipe Reader
func Pipe() (*bufio.Reader, error) {
	info, err := os.Stdin.Stat()

	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		return nil, ErrDataIsNotFromAPIPE
	}

	return bufio.NewReader(os.Stdin), nil
}
