package tests

import (
	"fmt"
	"strconv"
	"testing"
)

var testSuite = []struct {
	description string
	expected    bool
	input       string
}{
	{
		"WithLengthLessOne",
		true,
		"",
	},
	{
		"WithLengthEqualOne",
		true,
		"a",
	},
	{
		"WithLengthTwoAndTHeSameCharacter",
		true,
		"aa",
	},
	{
		"WithFirstCharacterAndLastAreDifferent",
		false,
		"ab",
	},
	{
		"WithPalindromeText",
		true,
		"kayak",
	},
	{
		"WithNotPalindromeText",
		false,
		"palindrome",
	},
}

//func TestIsPalindrome(t *testing.T) {
//	for _, ts := range testSuite {
//		actual := IsPalindrome(ts.input)
//		if ts.expected != actual {
//			t.Errorf("[%s] Expected %v but got %v", ts.description, ts.expected, actual)
//		}
//	}
//}

func TestIsPalindrome_InParallel(t *testing.T) {
	testMap := map[string]struct{}{}

	for _, ts := range testSuite {
		ts := ts
		t.Run(ts.description, func(t *testing.T) {
			t.Parallel()

			testMap["a"] = struct{}{}
			//actual := IsPalindrome(ts.input)
			//if ts.expected != actual {
			//	t.Errorf("[%s] Expected %v but got %v", ts.description, ts.expected, actual)
			//}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	// 161685669 - 168006567
	// 6.862 ns/op
	for i := 0; i < b.N; i++ {
		IsPalindrome("kayak")
	}
}

func BenchmarkSlice(b *testing.B) {
	b.Skip()
	var testSlice []interface{}
	for i := 0; i < b.N; i++ {
		testSlice = append(testSlice, i)
	}
}

func BenchmarkSprintF(b *testing.B) {
	b.Skip()
	for i := 0; i < b.N; i++ {
		if v := fmt.Sprintf("%d", 42); v != "42" {
			b.Fatalf("Unexpected string: %v", v)
		}
	}
}

func BenchmarkSet(b *testing.B) {
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		set := Set{set: map[string]struct{}{}}

		b.StartTimer()
		for _, elem := range testSet {
			set.Add(elem)
		}
		b.StopTimer()

		for _, elem := range testSet {
			set.Delete(elem)
		}
	}
}

// "a", "b", "1", ""
//func TestIsPalindrome_WithLengthLessOrEqualOne(t *testing.T) {
//	expected := true
//
//	actual := IsPalindrome("a")
//
//	if expected != actual {
//		t.Errorf("Expected %v but got %v", expected, actual)
//	}
//
//	actual = IsPalindrome("")
//
//	if expected != actual {
//		t.Errorf("Expected %v but got %v", expected, actual)
//	}
//}

// "aa" "bb" "cc"
//func TestIsPalindrome_WithLengthTwoAndTHeSameCharacter(t *testing.T) {
//	expected := true
//
//	actual := IsPalindrome("aa")
//
//	if expected != actual {
//		t.Errorf("Expected %v but got %v", expected, actual)
//	}
//}

// "abc" "cbd"
//func TestIsPalindrome_WithFirstCharacterAndLastAreDifferent(t *testing.T) {
//	expected := false
//
//	actual := IsPalindrome("abc")
//
//	if expected != actual {
//		t.Errorf("Expected %v but got %v", expected, actual)
//	}
//}

// kayak
//func TestIsPalindrome_WithPalindromeText(t *testing.T) {
//	expected := true
//
//	actual := IsPalindrome("kayak")
//
//	if expected != actual {
//		t.Errorf("Expected %v but got %v", expected, actual)
//	}
//}

//func TestIsPalindrome_WithNotPalindromeText(t *testing.T) {
//	expected := false
//
//	actual := IsPalindrome("palindrome")
//
//	if expected != actual {
//		t.Errorf("Expected %v but got %v", expected, actual)
//	}
//}
