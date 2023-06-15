package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	s, spec := "", " "
	for _, arg := range os.Args {
		s += arg + spec
		spec = " "
	}
	fmt.Println(s)
	fmt.Printf("Use for to print, time: %.2f \n", time.Since(start1).Seconds())
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("Use for to print, time: %.2f \n", time.Since(start2).Seconds())
}
