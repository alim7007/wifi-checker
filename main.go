package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func removeNonNumbers(strs []string) []string {
	var result []string
	var removed []string
	for _, str := range strs {
		if len(str) >= 8 {
			_, err := strconv.Atoi(str)
			if err == nil {
				if !strings.Contains(str, "199") {
					if !strings.Contains(str, "198") {
						if !strings.Contains(str, "197") {
							if !strings.Contains(str, "196") {
								removed = append(removed, str)
							}
						}
					}
				}
			}
		}
	}
	result = removeDuplicateStr(removed)
	return result
}

func main() {
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var passwords []string

	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	file.Close()
	shellFormat := []string{}
	for _, each := range removeNonNumbers(passwords) {
		w1 := `"`
		w2 := `"`
		word := w1 + each + w2
		shellFormat = append(shellFormat, word)
	}
	fmt.Println(shellFormat)
	fmt.Println(len(shellFormat))
}

// func readLines(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	return lines, scanner.Err()
// }

// func init() {
// 	lines, err := readLines("inputs_copy.txt")
// 	if err != nil {
// 		log.Fatalf("readLines: %s", err)
// 	}
// 	fmt.Println(lines[0])

// 	splited := strings.Split(lines[1], `" "`)
// 	for _, each := range splited {
// 		fmt.Println(each)
// 	}
// }
