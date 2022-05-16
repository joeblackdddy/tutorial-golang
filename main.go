package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[0:], "\n"))
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	fmt.Println(time.Since(start).Seconds())
}
