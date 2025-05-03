package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f1Name, f2Name := os.Args[1], os.Args[2]
	f1ContentRaw, err := os.ReadFile(f1Name)
	if err != nil {
		panic(err)
	}

	f2ContentRaw, err := os.ReadFile(f2Name)
	if err != nil {
		panic(err)
	}

	f1Content := string(f1ContentRaw)
	f2Content := string(f2ContentRaw)
	f1Arr := strings.Split(f1Content, "\n")
	f2Arr := strings.Split(f2Content, "\n")

	diffResult := diff(f1Arr, f2Arr)
	fmt.Print(diffResult.String())
}
