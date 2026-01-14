package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	show := os.Args[1:]
	show1 := ""
	if len(show) == 0 {
		// Read all data from os.Stdin until EOF
		stdinData, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		// Convert the byte slice to a string and process it
		//fmt.Printf("Received %d bytes:\n%s", len(stdinData), string(stdinData))
		show1 = string(stdinData)
	} else {
		show1 = strings.Join(show, " ")
	}

	fmt.Println("")
	fmt.Println("----------------------------")
	fmt.Println(show1)
	fmt.Println("----------------------------")
	fmt.Println("     \\")
	fmt.Println("      \\")
	fmt.Println("       \\  ⠀⠀⠀⢀⣀⣤⣤⣤⣤⣄⡀⠀⠀⠀⠀")
	fmt.Println("          ⠀⢀⣤⣾⣿⣾⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀")
	fmt.Println("          ⢠⣾⣿⢛⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀")
	fmt.Println("          ⣾⣯⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧")
	fmt.Println("          ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿")
	fmt.Println("          ⣿⡿⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠻⢿⡵")
	fmt.Println("          ⢸⡇⠀⠀⠉⠛⠛⣿⣿⠛⠛⠉⠀⠀⣿⡇")
	fmt.Println("          ⢸⣿⣀⠀⢀⣠⣴⡇⠹⣦⣄⡀⠀⣠⣿⡇")
	fmt.Println("          ⠈⠻⠿⠿⣟⣿⣿⣦⣤⣼⣿⣿⠿⠿⠟⠀ ")
	fmt.Println("          ⠀⠀⠀⠀⠸⡿⣿⣿⢿⡿⢿⠇⠀⠀⠀ ")
	fmt.Println("          ⠀⠀⠀⠀⠸⡿⣿⣿⢿⡿⢿⠇⠀⠀⠀ ")
	fmt.Println("          ⠀⠀⠀⠀⠀⠀⠈⠁⠈⠁⠀⠀⠀⠀⠀⠀")
}
