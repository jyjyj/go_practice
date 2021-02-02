package easy

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := 0
	result := isPalindrome(input)

	expected := true
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("isPalindrome wrong: %v; expected %v", result, expected)
	}
}

// get ginkgo?
