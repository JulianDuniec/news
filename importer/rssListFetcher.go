package importer

import (
	"bufio"
	"os"
	"io"
	"strings"
)

func fetchRssList(path string) ([]string, error) {
	file,  err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	r := bufio.NewReader(file)
	for  {
		s, err := r.ReadString(10)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		lines = append(lines, strings.TrimRight(s, "\n"))
	}
	return lines, nil
}