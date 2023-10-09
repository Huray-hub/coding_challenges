package testhelper

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// WIP: This helper function aims to automate hackerrank problem's local testing
func RunTestCases(t *testing.T, main func()) {
	t.Helper()

	files, err := os.ReadDir("./input")
	if err != nil {
		t.Fatalf("setup phase - could not read directory: %s", err)
	}

	outputDir := t.TempDir()

	for _, file := range files {
		runTestCase(t, file.Name(), outputDir, main)
	}
}

// TODO: Maybe export this too
func runTestCase(t *testing.T, fileName, outputDir string, main func()) {
	testCase := extractTestCase(fileName)

	t.Run(fmt.Sprintf("Test case %s", testCase), func(t *testing.T) {
		// Arrange
		inputFilePath := fmt.Sprintf("./input/input%s.txt", testCase)
		resultFilePath := fmt.Sprintf("%s/output%s.txt", outputDir, testCase)

		inputFile, err := os.Open(inputFilePath)
		if err != nil {
			t.Fatalf("could not open input file: %s", err)
		}
		defer inputFile.Close()

		stdin := os.Stdin
		os.Stdin = inputFile
		defer func() { os.Stdin = stdin }()

		os.Setenv("OUTPUT_PATH", resultFilePath)

		expectedOutput, err := expectedOutput(testCase)
		if err != nil {
			t.Fatalf("could not read output file: %s", err)
		}

		resultFile, err := os.Create(resultFilePath)
		if err != nil {
			t.Fatalf("could not open result file: %s", err)
		}
		defer resultFile.Close()

		stdout := os.Stdout
		os.Stdout = resultFile
		defer func() { os.Stdout = stdout }()

		// Act
		main()

		// Assert
		actualOutput, err := actualOutput(resultFilePath)
		if err != nil {
			t.Fatalf("could not read result file: %s", err)
		}

		if expectedOutput != actualOutput {
			t.Fatalf("\nexpected:\n%q\nactual:\n%q", expectedOutput, actualOutput)
		}
	})
}

// TODO: what if...
// func inputFile(testCase string) (inputFile *os.File, restoreStdin func(), err error) {
// 	inputFilePath := fmt.Sprintf("./input/input%s.txt", testCase)
// 	inputFile, err = os.Open(inputFilePath)
// 	if err != nil {
// 		return
// 		// t.Fatalf("could not open input file: %s", err)
// 	}
//
// 	stdin := os.Stdin
// 	os.Stdin = inputFile
// 	restoreStdin = func() { os.Stdin = stdin }
// 	return
// }

// extractTestCase function extracts the test case number from the file name
func extractTestCase(fileName string) string {
	testCase := strings.TrimSuffix(fileName, ".txt")
	testCase = strings.TrimPrefix(testCase, "input")
	return testCase
}

func expectedOutput(testCase string) (string, error) {
	outputFilePath := fmt.Sprintf("./output/output%s.txt", testCase)
	contents, err := os.ReadFile(outputFilePath)
	if err != nil {
		return "", err
	}
	// expectedOutput := strings.TrimRight(string(expectedOutputB), "\r\n")
	expectedOutput := strings.TrimRight(string(contents), "\n")
	return expectedOutput, nil
}

func actualOutput(resultFilePath string) (string, error) {
	// TODO: use readOpenFile instead of opening temp file twice
	// contents, err := readOpenFile(resultFile)
	contents, err := os.ReadFile(resultFilePath)
	if err != nil {
		return "", err
	}

	actualOutput := strings.TrimSuffix(string(contents), "\n")
	actualOutput = strings.TrimSuffix(actualOutput, "\n")
	return actualOutput, nil
}

// TODO: remove this if not needed in the end
func readOpenFile(f *os.File) ([]byte, error) {
	var size int
	if info, err := f.Stat(); err == nil {
		size64 := info.Size()
		if int64(int(size64)) == size64 {
			size = int(size64)
		}
	}
	size++ // one byte for final read at EOF

	// If a file claims a small size, read at least 512 bytes.
	// In particular, files in Linux's /proc claim size 0 but
	// then do not work right if read in small pieces,
	// so an initial read of 1 byte would not work correctly.
	if size < 512 {
		size = 512
	}

	data := make([]byte, 0, size)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}
}
