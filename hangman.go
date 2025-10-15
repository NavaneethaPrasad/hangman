package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSecretWord(wordFileName string) string {
	file, err := os.Open(wordFileName)
	if err != nil {
		fmt.Errorf("The file could not open", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return "navaneetha"
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
