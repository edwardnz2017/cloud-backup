package files

import (
	"fmt"
	"os"
)

func ReadLocalFile(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(file))
}
