package easy

import (
	"reflect"
	"testing"
)

func TestIsPalindrome_1(t *testing.T) {
	input := 121
	result := isPalindrome(input)

	expected := true
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("isPalindrome wrong: %v; expected %v", result, expected)
	}
}

func TestIsPalindrome_2(t *testing.T) {
	input := -121
	result := isPalindrome(input)

	expected := false
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("isPalindrome wrong: %v; expected %v", result, expected)
	}
}

func TestIsPalindrome_3(t *testing.T) {
	input := 10
	result := isPalindrome(input)

	expected := false
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("isPalindrome wrong: %v; expected %v", result, expected)
	}
}

// get ginkgo?
// add benchmarking
