package main

import "testing"

func TestStringCenter(t *testing.T) {
	var expected, actual string

	expected = "  test  "
	actual = StringCenter("test", 8, " ")
	if expected != actual {
		t.Fatalf("Basic center failed (expected \"%s\" actual \"%s\")", expected, actual)
	}

	expected = "  abc  "
	actual = StringCenter("abc", 7, " ")
	if expected != actual {
		t.Fatalf("Odd string length center failed (expected \"%s\" actual \"%s\")", expected, actual)
	}

	expected = " test  "
	actual = StringCenter("test", 7, " ")
	if expected != actual {
		t.Fatalf("Odd space width center failed (expected \"%s\" actual \"%s\")", expected, actual)
	}

	expected = " abc  "
	actual = StringCenter("abc", 6, " ")
	if expected != actual {
		t.Fatalf("Odd string and space length center failed (expected \"%s\" actual \"%s\")", expected, actual)
	}

	expected = " 天地無用 "
	actual = StringCenter("天地無用", 10, " ")
	if expected != actual {
		t.Fatalf("Unicode center failed (expected \"%s\" actual \"%s\")", expected, actual)
	}
}
