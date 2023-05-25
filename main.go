// AGPL - See LICENSE.txt

package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("ENSURE LOCAL")
	fmt.Println("Scans the complete file structure starting")
	fmt.Println("where it is run from. Reads the first byte")
	fmt.Println("of each file found. This should ensure any")
	fmt.Println("cloud service syncs files locally.")
	fmt.Println()
}
