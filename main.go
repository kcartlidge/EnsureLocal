// AGPL - See LICENSE.txt

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	buff  []byte
	debug bool = false
)

func main() {
	fmt.Println()
	fmt.Println(" _____                        __                _ ")
	fmt.Println("|   __|___ ___ _ _ ___ ___   |  |   ___ ___ ___| |")
	fmt.Println("|   __|   |_ -| | |  _| -_|  |  |__| . |  _| .'| |")
	fmt.Println("|_____|_|_|___|___|_| |___|  |_____|___|___|__,|_|")
	fmt.Println()
	fmt.Println("Scans a nested file structure and reads byte 1 of each file.")
	fmt.Println("Cloud files will therefore be synced locally.")
	fmt.Println()
	fmt.Println("File read failures are retried twice (cloud sync may be catching up).")
	fmt.Println("The first retry is after 2 seconds, the second after 10 seconds.")
	fmt.Println()
	fmt.Println("Folders/files with names starting '.' are skipped.")
	fmt.Println()
	fmt.Println("Usage:  EnsureLocal [folder]")
	fmt.Println("        (If no folder is provided you will be prompted)")
	fmt.Println()

	folder := ""
	if len(os.Args) != 2 {
		cwd, _ := os.Getwd()
		fmt.Println("The current working folder (.) is:")
		fmt.Println(cwd)
		fmt.Println()
		fmt.Println("Folder (absolute, or relative to the above): ")
		n, err := fmt.Scanln(&folder)
		fmt.Println()
		if n != 1 || err != nil {
			os.Exit(2)
		}
	} else {
		folder = os.Args[1]
	}
	absFolder, err := filepath.Abs(folder)
	check(err)

	fmt.Println()
	fmt.Println("SCANNING:")
	fmt.Println(absFolder)
	fmt.Println()
	fmt.Println("Stop the run at any time using Ctrl+C.")
	fmt.Println()
	fmt.Println("Cloud/local speed/volumes will impact time taken.")
	fmt.Println("Press Enter/Return to continue (or Ctrl+C)")
	fmt.Println()
	fmt.Scanln()

	buff = make([]byte, 1)
	err = filepath.WalkDir(folder, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			if d == nil {
				// Errors reading the root.
				return err
			} else {
				showIssue(path, err.Error())
			}
		}

		// Skip anything whose path inludes a dotted portion.
		segments := strings.Split(path, string(filepath.Separator))
		include := true
		for _, segment := range segments {
			if segment == "." || strings.HasPrefix(segment, ".") {
				include = false
				break
			}
		}
		if include {
			if d.IsDir() {
				fmt.Println(path)
			} else {
				fi, err := d.Info()
				if err != nil {
					showIssue(path, err.Error())
				} else if fi.Size() > 0 {
					// Make 3 attempts, with a 2 second and then a 10 second back-off.
					err = readFirstByte(path)
					if err != nil {
						time.Sleep(time.Second * 2)
						err = readFirstByte(path)
						if err != nil {
							time.Sleep(time.Second * 10)
							err = readFirstByte(path)
							showIssue(path, err.Error())
						}
					}
				}
				if debug {
					fmt.Printf("  %x  %s\n", buff, path)
				}
			}
		}
		return nil
	})
	check(err)

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}

func showIssue(path string, errorMessage string) {
	fmt.Printf("* %s -- %s", path, errorMessage)
}

func readFirstByte(path string) error {
	f, err := os.Open(path)
	if err == nil {
		n, err := f.Read(buff)
		if err == nil {
			if n != 1 {
				return errors.New("could not read first byte of file: " + path)
			}
		}
	}
	return err
}

func check(err error) {
	if err != nil {
		fmt.Println()
		fmt.Println("ERROR")
		fmt.Println(err.Error())
		fmt.Println()
		os.Exit(1)
	}
}
