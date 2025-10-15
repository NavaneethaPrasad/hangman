package main

import (
	"fmt"
)

func getSecretWord(wordFileName string) string {
	return "navaneetha"

}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
