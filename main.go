// dino entry point
package main

import (
	"Dino/process"
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🦖 Welcome to Dino")
	fmt.Println("Type commands to run processes. Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		//launch the process concurrently
		wg.Add(1)
		go process.RunProcess(input, &wg)
	}

	wg.Wait()
	fmt.Println("All processes finished! Dino says goodbye 🦖")
}
