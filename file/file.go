package file

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

// String -> []String
// Function accepts a string containing a filename and returns the file's contents
// in a slice of strings, one line per string in slice
func FileLines(fileToken string) (a []string, err error) {
	// Open file and create Reader
	file, err := os.Open(fileToken)
	if err != nil {
		return nil, err
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// Attach file reader to buffered reader
	r := bufio.NewReader(file)
	// Read the first line into a string
	line, lineErr 	:= r.ReadString('\n')
	line 			= strings.Split(line,"\n")[0]
	line 			= strings.Split(line,"\r")[0]
	// Loop through, appending each line onto the slice until
	// ReadString returns an error (Probable EOF)
	for lineErr == nil {
		a 				= append(a, line)
		line, lineErr 	= r.ReadString('\n')
		line 			= strings.Split(line,"\n")[0]
		line 			= strings.Split(line,"\r")[0]
	}
	if len(line) > 0 {
		a = append(a, line)
	}
	// Return the slice a
	return
}

// String -> [][]int
// Function accepts a string containing a filename, and parses the file
// for integers. Returns a slice of slices of integers, one slice
// per line
func IntLines(fileToken string) [][]int {
	// read the file in as strings
	fileLines, _ 	:= FileLines(fileToken)
	// Create a slice of slices, one slice per line
	intSlice	:= make([][]int, len(fileLines))
	// Parse each string in fileLines for integers and add the returned
	// slice to slice of slices
	for i, v := range fileLines {
		intSlice[i] = intParse(v)
	}
	return intSlice
}

// String -> []int
// Function accepts a string and parses it for integers. Will not behave well
// if non-integer values are contained in string
func intParse(line string) []int {
	// Removes everything on or after \n character
	sliceLine	:= strings.Split(line, "\n")
	// Removes everything on or after \r character
	sliceLine	=  strings.Split(sliceLine[0], "\r")
	// Cuts string into a slice of strings by separating wherever there is a space
	sliceLine	=  strings.Split(sliceLine[0], " ")
	// Make slice of ints appropriately sized to hold the conversions
	ints 		:= make([]int, len(sliceLine))
	// Convert each string in the slice to integers and add them to the int slice
	for i, v := range sliceLine {
		ints[i], _ = strconv.Atoi(v)
	}
	return ints 
}

