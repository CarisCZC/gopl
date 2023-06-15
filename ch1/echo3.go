package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	sec := time.Since(start).Seconds()
	fmt.Printf("%.2f Args: %s", sec, os.Args[1:])
}
