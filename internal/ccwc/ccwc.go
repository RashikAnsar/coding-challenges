package ccwc

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getBytes(file *os.File) int {
	file.Seek(0, io.SeekStart)
	bytes := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		bytes++
	}
	return bytes
}

func getLines(file *os.File) int {
	file.Seek(0, io.SeekStart)
	lines := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines++
	}

	return lines
}

func getWords(file *os.File) int {
	file.Seek(0, io.SeekStart)
	words := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words++
	}

	return words
}

func getChars(file *os.File) int {
	file.Seek(0, io.SeekStart)
	chars := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		chars++
	}
	return chars
}

func defaultOutput(file *os.File, filename string) {
	linesCount := getLines(file)
	bytesCount := getBytes(file)
	wordsCount := getWords(file)
	fmt.Printf("\t%d\t%d\t%d\t%s\n", linesCount, wordsCount, bytesCount, filename)
}

func Run(file *os.File, lines, chars, words, byteSize bool, filename string) {
	var result int

	if lines {
		result = getLines(file)
	} else if chars {
		result = getChars(file)
	} else if words {
		result = getWords(file)
	} else if byteSize {
		result = getBytes(file)
	} else {
		defaultOutput(file, filename)
		return
	}

	fmt.Println(result)
}
