package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func parseDateTime(datetimeStr string) (time.Time, error) {
	if unixTime, err := strconv.ParseInt(datetimeStr, 10, 64); err == nil {
		return time.Unix(unixTime, 0), nil
	}

	// List of datetime formats to try
	formats := []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.DateTime,
	}

	// Try parsing the datetime string using the predefined formats
	for _, format := range formats {
		if parsedTime, err := time.Parse(format, datetimeStr); err == nil {
			return parsedTime, nil
		}
	}

	// If none of the formats work, return an error
	return time.Time{}, fmt.Errorf("unable to parse datetime: %s", datetimeStr)
}

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
			parsedTime, err := parseDateTime(dTime)
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
