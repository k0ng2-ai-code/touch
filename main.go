package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Define flags
	var atimeOnly bool
	var cSuppressCreation bool
	var dTime string
	var mtimeOnly bool
	var rReference string

	flag.BoolVar(&atimeOnly, "a", false, "change only the access time")
	flag.BoolVar(&cSuppressCreation, "c", false, "do not create any files")
	flag.StringVar(&dTime, "d", "", "set the access and modification times using the specified string")
	flag.BoolVar(&mtimeOnly, "m", false, "change only the modification time")
	flag.StringVar(&rReference, "r", "", "use this file's times instead of the current time")

	// Parse command-line arguments
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("usage: touch [OPTION]... FILE...")
		os.Exit(1)
	}

	// Process each filename provided
	for _, filename := range flag.Args() {
		// Check if the file exists
		fileInfo, err := os.Stat(filename)
		if os.IsNotExist(err) {
			// File does not exist, handle -c flag
			if cSuppressCreation {
				continue
			}
			// Create the file
			file, err := os.Create(filename)
			if err != nil {
				fmt.Printf("error creating file: %v\n", err)
				os.Exit(1)
			}
			file.Close()
			// Get file info for the newly created file
			fileInfo, err = os.Stat(filename)
			if err != nil {
				fmt.Printf("error stating file: %v\n", err)
				os.Exit(1)
			}
		} else if err != nil {
			fmt.Printf("error checking file: %v\n", err)
			os.Exit(1)
		}

		// Determine the time to set
		var atime, mtime time.Time
		if rReference != "" {
			refInfo, err := os.Stat(rReference)
			if err != nil {
				fmt.Printf("error stating reference file: %v\n", err)
				os.Exit(1)
			}
			atime = refInfo.ModTime()
			mtime = refInfo.ModTime()
		} else if dTime != "" {
			parsedTime, err := time.Parse(time.RFC3339, dTime)
			if err != nil {
				fmt.Printf("error parsing time: %v\n", err)
				os.Exit(1)
			}
			atime = parsedTime
			mtime = parsedTime
		} else {
			currentTime := time.Now().Local()
			atime = currentTime
			mtime = currentTime
		}

		// Adjust times based on flags
		if atimeOnly {
			mtime = fileInfo.ModTime()
		}
		if mtimeOnly {
			atime = fileInfo.ModTime()
		}

		// Update the file's timestamps
		err = os.Chtimes(filename, atime, mtime)
		if err != nil {
			fmt.Printf("error updating timestamps: %v\n", err)
			os.Exit(1)
		}
	}
}
