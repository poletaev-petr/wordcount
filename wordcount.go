package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	v := strings.Split(os.Args[1], " ")
	fmt.Println(len(v))
}
