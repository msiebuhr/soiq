package main

// SOIQ is short for "sort | uniq -c", using less memory and making things faster.
//
// After trying it out a few times, it does actually have a runtime similar to
// "sort | uniq -c", but uses about 1/3 of the memory. I've tried disabling GC
// (set GOGC=-1), but to no avail. Further investigation needed...

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer os.Stdin.Close()
	output := make(map[string]uint)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		b := scanner.Text()

		val, ok := output[b]

		if !ok {
			val = 0
		}
		output[b] = val + 1
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for key, count := range output {
		fmt.Printf("%d\t%s\n", count, key)
	}
}
