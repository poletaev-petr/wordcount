package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 && os.Args[0] != " " && os.Args[0] != "" &&
		os.Args[1] != " " && os.Args[1] != "" {
		v := strings.Split(os.Args[1], " ")
		fmt.Println(len(v))

	} else {
		fmt.Println("0")
	}

}
