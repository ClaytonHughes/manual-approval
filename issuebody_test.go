package main

import (
	"reflect"
	"testing"
)

func TestIssueBody(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
	}{
		{
			name:     "with_no_extra_body",
			input:    []string{""},
		},
		{
			name:     "with_extra_body",
			input:    []string{"line 1\nline 2"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := deduplicateUsers(testCase.input)
			if !reflect.DeepEqual(testCase.expected, actual) {
				t.Fatalf(
					"unequal depulicated: expected %v actual %v",
					testCase.expected,
					actual,
				)
			}
		})
	}
}
