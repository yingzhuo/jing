/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *      _ _
 *     | (_)_ __   __ _
 *  _  | | | '_ \ / _` |
 * | |_| | | | | | (_| |
 *  \___/|_|_| |_|\__, |                        https://github.com/yingzhuo/jing
 *                |___/
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package io

import (
	"bufio"
	"os"
	"strings"
)

func NewFileLines(filename string) (*FileLines, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	result := &FileLines{
		Count: 0,
		Lines: make([]string, 0),
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result.Lines = append(result.Lines, scanner.Text())
		result.Count++
	}
	return result, nil
}

type FileLines struct {
	Count int
	Lines []string
}

func (fl *FileLines) Map(fn func(line string) string) *FileLines {

	result := &FileLines{
		Count: fl.Count,
		Lines: make([]string, 0),
	}

	for _, line := range fl.Lines {
		newLine := fn(line)
		result.Lines = append(result.Lines, newLine)
	}

	return result
}

func (fl *FileLines) TrimSpace() *FileLines {
	return fl.Map(func(line string) string {
		return strings.TrimSpace(line)
	})
}

func (fl *FileLines) Filter(fn func(line string) bool) *FileLines {
	result := &FileLines{
		Count: 0,
		Lines: make([]string, 0),
	}
	for _, line := range fl.Lines {
		if fn(line) {
			result.Lines = append(result.Lines, line)
			result.Count++
		}
	}
	return result
}

func (fl *FileLines) String() string {
	build := strings.Builder{}
	for _, line := range fl.Lines {
		build.WriteString(line)
	}
	return build.String()
}
