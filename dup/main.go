// Prints both count and text of lines that came up multiple times
package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	// method1()
	// method2()
	method3()
}

// func method1() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)

// 	for input.Scan() {
// 		counts[input.Text()]++
// 		if input.Text() == "q" {
// 			break
// 		}
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func method2() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]

// 	for _, arg := range files {
// 		f, err := os.Open(arg)

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
// 			continue
// 		}

// 		input := bufio.NewScanner(f)
// 		// Open a stream
// 		for input.Scan() {
// 			counts[input.Text()]++
// 		}

// 		f.Close()
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

func method3() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// Load into memory
		data, err := os.ReadFile(filename) // data is a slice of bytes

		// The default type is a decimal of a single byte (uint8)
		fmt.Println(reflect.TypeOf(data[0]))

		fmt.Println("Decimal", data)

		fmt.Printf("Octal: %o\n", data)

		fmt.Printf("Hexadecimal: %x\n", data)

		fmt.Printf("Binary: %b\n", data)

		fmt.Printf("String: %s\n", data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
