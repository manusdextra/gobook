package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func printargs(method string) {
	switch method {
	case "loop":
		for i := 1; i < len(os.Args); i++ {
			fmt.Printf("%s\n", os.Args[i])
		}
	case "range":
		for _, value := range os.Args[1:] {
			fmt.Printf("%s\n", value)
		}
	case "join":
		fmt.Println(strings.Join(os.Args[1:], "\n"))
	}
}

func main() {
	var start time.Time
	timetaken := make(map[string][]float64)
	iterations := 10000

	for _, method := range []string{"range", "join", "loop"} {
		timetaken[method] = make([]float64, iterations)
		for i := 0; i < iterations; i++ {
			start = time.Now()
			printargs(method)
			timetaken[method][i] = time.Since(start).Seconds()
		}
	}

	for k, v := range timetaken {
		fmt.Printf("Method %q ", k)
		var sum float64
		for _, time := range v {
			sum = sum + time
		}
		fmt.Printf("takes %fs on average\n", sum/float64(iterations))
	}
}
