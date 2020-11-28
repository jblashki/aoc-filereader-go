package filereader

import (
	"regexp"
	"testing"
)

func TestReadAllInts(t *testing.T) {
	const inputFile = "./test_input/ReadAllInts.tst"
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 999999999, -1, -10000, -999999999}

	results, err := ReadAllInts(inputFile)
	if err != nil {
		t.Fatalf(`ReadAllInts(%q) returned %v`, inputFile, err)
	}

	resultsLen := len(results)
	wantLen := len(want)

	if resultsLen != wantLen {
		t.Fatalf(`ReadAllInts(%q) returned slice with length %v, want slice with length %v`,
			inputFile, resultsLen, wantLen)
	}

	for i := 0; i < len(results); i++ {
		if results[i] != want[i] {
			t.Fatalf(`ReadAllInts(%q) returned %v @ position %v, want %v`,
				inputFile, results[i], i, want[i])
		}
	}
}

func TestReadAllIntsInvalidInput(t *testing.T) {
	const inputFile = "./test_input/ReadAllIntsInvalidInput.tst"
	_, err := ReadAllInts(inputFile)
	if err == nil {
		t.Fatalf(`ReadAllInts(%q) failed to return error on invalid input`, inputFile)
	}

	want := regexp.MustCompile(`parsing \"1a\": invalid syntax`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`ReadAllInts(%q) error: %q, want match for %#q`,
			inputFile, err.Error(), want)
	}

}

func TestReadAllIntsLargeInput(t *testing.T) {
	const inputFile = "./test_input/ReadAllIntsLargeInput.tst"
	_, err := ReadAllInts(inputFile)
	if err == nil {
		t.Fatalf(`ReadAllInts(%q) failed to return error on invalid input`, inputFile)
	}

	want := regexp.MustCompile(`parsing \"9999999999999999999\": value out of range`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`ReadAllInts(%q) error: %q, want match for %#q`,
			inputFile, err.Error(), want)
	}

}

func TestReadIntoString(t *testing.T) {
	const inputFile = "./test_input/ReadIntoString.tst"
	want := "line1\n\tline2\ntesting 123\n"

	results, err := ReadIntoString(inputFile)
	if err != nil {
		t.Fatalf(`ReadIntoString(%q) returned %v`, inputFile, err)
	}

	resultsLen := len(results)
	wantLen := len(want)

	if resultsLen != wantLen {
		t.Fatalf(`ReadIntoString(%q) returned string with length %v, want string with length %v`,
			inputFile, resultsLen, wantLen)
	}

	for i := 0; i < len(results); i++ {
		if results[i] != want[i] {
			t.Fatalf(`ReadIntoString(%q) returned '%c' @ position %v, want '%c'`,
				inputFile, results[i], i, want[i])
		}
	}
}

func TestReadLines(t *testing.T) {
	const inputFile = "./test_input/ReadLines.tst"
	want := []string{"line1",
		"\tline2",
		"testing 123",
	}

	results, err := ReadLines(inputFile)
	if err != nil {
		t.Fatalf(`ReadLines(%q) returned %v`, inputFile, err)
	}

	resultsLen := len(results)
	wantLen := len(want)

	if resultsLen != wantLen {
		t.Fatalf(`ReadLines(%q) returned slice with length %v, want slice with length %v`,
			inputFile, resultsLen, wantLen)
	}

	for i := 0; i < len(results); i++ {
		if results[i] != want[i] {
			t.Fatalf(`ReadLines(%q) returned %q @ position %v, want %q`,
				inputFile, results[i], i, want[i])
		}
	}
}

func TestReadCSVInts(t *testing.T) {
	const inputFile = "./test_input/ReadCSVInts.tst"
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 999999999, -1, -10000, -999999999}

	results, err := ReadCSVInts(inputFile)
	if err != nil {
		t.Fatalf(`ReadCSVInts(%q) returned %v`, inputFile, err)
	}

	resultsLen := len(results)
	wantLen := len(want)

	if resultsLen != wantLen {
		t.Fatalf(`ReadCSVInts(%q) returned slice with length %v, want slice with length %v`,
			inputFile, resultsLen, wantLen)
	}

	for i := 0; i < len(results); i++ {
		if results[i] != want[i] {
			t.Fatalf(`ReadCSVInts(%q) returned %v @ position %v, want %v`,
				inputFile, results[i], i, want[i])
		}
	}
}

func TestReadCSVIntsInvalidInput(t *testing.T) {
	const inputFile = "./test_input/ReadCSVIntsInvalidInput.tst"
	_, err := ReadCSVInts(inputFile)
	if err == nil {
		t.Fatalf(`ReadCSVInts(%q) failed to return error on invalid input`, inputFile)
	}

	want := regexp.MustCompile(`parsing \"1a\": invalid syntax`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`ReadCSVInts(%q) error: %q, want match for %#q`,
			inputFile, err.Error(), want)
	}

}

func TestReadCSVIntsLargeInput(t *testing.T) {
	const inputFile = "./test_input/ReadCSVIntsLargeInput.tst"
	_, err := ReadCSVInts(inputFile)
	if err == nil {
		t.Fatalf(`ReadCSVInts(%q) failed to return error on invalid input`, inputFile)
	}

	want := regexp.MustCompile(`parsing \"9999999999999999999\": value out of range`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`ReadCSVInts(%q) error: %q, want match for %#q`,
			inputFile, err.Error(), want)
	}

}

func TestReadCSVIntsBareQuote(t *testing.T) {
	const inputFile = "./test_input/ReadCSVIntsBareQuote.tst"
	_, err := ReadCSVInts(inputFile)
	if err == nil {
		t.Fatalf(`ReadCSVInts(%q) failed to return error on invalid csv input`, inputFile)
	}

	want := regexp.MustCompile(`bare \" in non-quoted-field`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`ReadCSVInts(%q) error: %q, want match for %#q`,
			inputFile, err.Error(), want)
	}

}
