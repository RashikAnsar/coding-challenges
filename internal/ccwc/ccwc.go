package ccwc

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GetBytes(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func GetLines(file *os.File) int {
	file.Seek(0, io.SeekStart)
	lines := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines++
	}

	return lines
}

func GetWords(file *os.File) int {
	file.Seek(0, io.SeekStart)
	words := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words++
	}

	return words
}

func GetChars(file *os.File) int {
	file.Seek(0, io.SeekStart)
	chars := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		chars++
	}
	return chars
}

func DefaultOutput(file *os.File) {
	bytesCount, _ := GetBytes(file)
	wordsCount := GetWords(file)
	linesCount := GetLines(file)
	fmt.Printf("\t%d\t%d\t%d\t%s\n", linesCount, wordsCount, bytesCount, file.Name())
}
