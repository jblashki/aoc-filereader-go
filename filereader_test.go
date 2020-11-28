package filereader

import (
	"regexp"
	"testing"
)

func TestReadAllInts(t *testing.T) {
	const inputFile = "./test_input/ReadAllInts"
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
	const inputFile = "./test_input/ReadAllIntsInvalidInput"
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
	const inputFile = "./test_input/ReadAllIntsLargeInput"
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
	const inputFile = "./test_input/ReadIntoString"
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
	const inputFile = "./test_input/ReadLines"
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
	const inputFile = "./test_input/ReadCSVInts"
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
	const inputFile = "./test_input/ReadCSVIntsInvalidInput"
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
	const inputFile = "./test_input/ReadCSVIntsLargeInput"
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
	const inputFile = "./test_input/ReadCSVIntsBareQuote"
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

func TestReadCSVStringsPerLine(t *testing.T) {
	const inputFile = "./test_input/ReadCSVStringsPerLine"
	want := make([][]string, 3)
	want[0] = []string{"l1f1", "l1f2", "l1f3", "l1f4"}
	want[1] = []string{"l2f1", "l2f2"}
	want[2] = []string{"l3f1", "l3f2", "l3f3", "l3f4", "l3f5", "l3f6"}

	results, err := ReadCSVStringsPerLine(inputFile)
	if err != nil {
		t.Fatalf(`ReadCSVStringsPerLine(%q) returned %v`, inputFile, err)
	}

	resultsLen := len(results)
	wantLen := len(want)

	if resultsLen != wantLen {
		t.Fatalf(`ReadCSVStringsPerLine(%q) returned slice with length %v, want slice with length %v`,
			inputFile, resultsLen, wantLen)
	}

	for i := 0; i < len(results); i++ {
		resultsLineLen := len(results[i])
		wantLineLen := len(want[i])

		if resultsLineLen != wantLineLen {
			t.Fatalf(`ReadCSVStringsPerLine(%q) returned slice element %v with length %v, want slice element with length %v`,
				inputFile, i, resultsLineLen, wantLineLen)
		}

		for j := 0; j < len(results[i]); j++ {
			if results[i][j] != want[i][j] {
				t.Fatalf(`ReadCSVStringsPerLine(%q) returned %q @ position [%v][%v], want %q`,
					inputFile, results[i][j], i, j, want[i][j])
			}
		}
	}
}

func TestReadCSVStringsPerLineBareQuote(t *testing.T) {
	const inputFile = "./test_input/ReadCSVStringsPerLineBareQuote"
	_, err := ReadCSVStringsPerLine(inputFile)
	if err == nil {
		t.Fatalf(`ReadCSVStringsPerLine(%q) failed to return error on invalid csv input`, inputFile)
	}

	want := regexp.MustCompile(`bare \" in non-quoted-field`)

	if !want.MatchString(err.Error()) {
		t.Fatalf(`ReadCSVStringsPerLine(%q) error: %q, want match for %#q`,
			inputFile, err.Error(), want)
	}

}
