package filereader

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Function to read file line by line into slice of ints
func ReadAllInts(path string) ([]int, error) {
	var retArray []int

	if path == "" {
		return retArray, errors.New("empty path")
	}

	// Open file here
	buf, err := os.Open(path)
	if err != nil {
		return retArray, err
	}

	defer func() {
		if err = buf.Close(); err != nil {
			return
		}
	}()

	// Read file line by line
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		var number int
		number, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return retArray, err
		}
		retArray = append(retArray, number)
	}
	err = scanner.Err()

	return retArray, err
}

// Function to read file into string
func ReadIntoString(path string) (string, error) {
	if path == "" {
		return "", errors.New("empty path")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), err
}

// Function to read file line by line into slice of strings
func ReadLines(path string) ([]string, error) {
	var retArray []string

	if path == "" {
		return retArray, errors.New("empty path")
	}

	// Open file here
	buf, err := os.Open(path)
	if err != nil {
		return retArray, err
	}

	defer func() {
		if err = buf.Close(); err != nil {
			return
		}
	}()

	// Read file line by line
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		retArray = append(retArray, scanner.Text())
	}
	err = scanner.Err()

	return retArray, err
}

// Function to read csv file into int slice
func ReadCSVInts(path string) ([]int, error) {
	var retArray []int

	if path == "" {
		return retArray, errors.New("empty path")
	}

	// Open file here
	csvFile, err := os.Open(path)
	if err != nil {
		return retArray, err
	}

	defer func() {
		if err = csvFile.Close(); err != nil {
			return
		}
	}()

	reader := csv.NewReader(csvFile)

	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			if perr, ok := err.(*csv.ParseError); ok {
				// Only care about ParseErrors if it's
				// Error with field count
				if perr.Err != csv.ErrFieldCount {
					return retArray, err
				}
			} else {
				return retArray, err
			}
		}

		for value := range record {
			var number int
			number, err = strconv.Atoi(strings.TrimSpace(record[value]))
			if err != nil {
				return retArray, err
			}
			retArray = append(retArray, number)
		}
	}

	return retArray, nil
}

// Function to read csv file into int slice
func ReadCSVStrings(path string) ([]string, error) {
	var retArray []string

	if path == "" {
		return retArray, errors.New("empty path")
	}

	// Open file here
	csvFile, err := os.Open(path)
	if err != nil {
		return retArray, err
	}

	defer func() {
		if err = csvFile.Close(); err != nil {
			return
		}
	}()

	reader := csv.NewReader(csvFile)

	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			if perr, ok := err.(*csv.ParseError); ok {
				// Only care about ParseErrors if it's
				// Error with field count
				if perr.Err != csv.ErrFieldCount {
					return retArray, err
				}
			} else {
				return retArray, err
			}
		}

		for value := range record {
			s := strings.TrimSpace(record[value])
			retArray = append(retArray, s)
		}
	}

	return retArray, nil
}

// Function to read csv file into slice of string slices for each line
func ReadCSVStringsPerLine(path string) ([][]string, error) {
	var retArray [][]string

	if path == "" {
		return retArray, errors.New("empty path")
	}

	// Open file here
	csvFile, err := os.Open(path)
	if err != nil {
		return retArray, err
	}

	defer func() {
		if err = csvFile.Close(); err != nil {
			return
		}
	}()

	reader := csv.NewReader(csvFile)

	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			if perr, ok := err.(*csv.ParseError); ok {
				// Only care about ParseErrors if it's
				// Error with field count
				if perr.Err != csv.ErrFieldCount {
					return retArray, err
				}
			} else {
				return retArray, err
			}
		}

		line := make([]string, len(record))
		copy(line, record)
		retArray = append(retArray, line)
	}

	return retArray, nil
}
