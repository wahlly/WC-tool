package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetFileByteSize(t *testing.T) {
	type paramsI struct{
		fileName	string
		input		string
		expected	int
	}

	testCases := []paramsI {
		{
			fileName:     "empty file",
			input:    "",
			expected: 0,
		},
		{
			fileName:     "single line",
			input:    "hello",
			expected: 7, // 5 bytes for "hello" + 2 bytes for newline
		},
		{
			fileName:     "multiple lines",
			input:    "hello\nworld",
			expected: 14, // 5 bytes for "hello" + 2 bytes for newline + 5 bytes for "world" + 2 bytes for newline
		},
		{
			fileName:     "lines with different lengths",
			input:    "short\nveryverylongline",
			expected: 25, // 5 bytes for "short" + 2 bytes for newline + 16 bytes for "veryverylongline" + 2 bytes for newline
		},
	}

	for _, test := range testCases {
		reader := strings.NewReader(test.input)
		actual, err := getFileByteSize(reader)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}

	}
}

func TestCountWordsInFile(t *testing.T) {
	type testsI struct {
		name		string
		content	string
		expected 	int
	}

	tests := []testsI {
		{
			name:     "empty file",
			content:  "",
			expected: 0,
		},
		{
			name:     "single word",
			content:  "hello",
			expected: 1,
		},
		{
			name:     "multiple words",
			content:  "hello world",
			expected: 2,
		},
		{
			name:     "multiple lines and words",
			content:  "hello\nworld\nthis is a test",
			expected: 6,
		},
		{
			name:     "words with different whitespace",
			content:  "hello\tworld\nthis  is\ta\ttest",
			expected: 6,
		},
	}

	for _, test := range tests {
		tempFile, err := os.CreateTemp("", "testFile")
		if err != nil {
			t.Fatalf("failed to create temporary file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		_, writeErr := tempFile.WriteString(test.content)
		if writeErr != nil {
			t.Fatalf("failed to write to temporary file: %v", writeErr)
		}

		//close the file to flush the content
		if err := tempFile.Close(); err != nil {
			t.Fatalf("failed to close temporary file: %v", err)
		}

		actual, err := countWordsInFile(tempFile.Name())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}