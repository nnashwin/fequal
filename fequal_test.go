package fequal_test

import (
	f "github.com/ru-lai/fequal"
	"strings"
	"testing"
	"time"
)

func appendTestDir(fName string) string {
	return strings.Join([]string{"testdata", fName}, "/")
}

func TestDoesEqual(t *testing.T) {
	uu := map[string]struct {
		fileName1, fileName2 string
		expected             bool
	}{
		"same file returns true": {
			fileName1: "forgottenRealms.json",
			fileName2: "forgottenRealms.json",
			expected:  true,
		},
		"different files with same contents return true": {
			fileName1: "forgottenRealms.json",
			fileName2: "forgottenRealmsCopy.json",
			expected:  true,
		},
	}

	for k, u := range uu {
		actual, err := f.AreEqual(appendTestDir(u.fileName1), appendTestDir(u.fileName2))
		if err != nil {
			t.Fatal(err)
		}

		if u.expected != actual {
			t.Fatalf("test case %s failed with the following values:\nactual: %t, expected: %t\n", k, actual, u.expected)
		}
	}
}

func TestDoesNotEqualUntimed(t *testing.T) {
	uu := map[string]struct {
		fileName1, fileName2 string
		expected             bool
	}{
		"different length files return false": {
			fileName1: "forgottenRealms.json",
			fileName2: "forgottenRealms2.json",
			expected:  false,
		},
		"files of the same length that contain different content return false": {
			fileName1: "forgottenRealms.json",
			fileName2: "forgottenRealmsMispell.json",
			expected:  false,
		},
	}

	for k, u := range uu {
		actual, err := f.AreEqual(appendTestDir(u.fileName1), appendTestDir(u.fileName2))
		if err != nil {
			t.Fatal(err)
		}

		if u.expected != actual {
			t.Fatalf("test case %s failed with the following values:\nactual: %t, expected: %t\n", k, actual, u.expected)
		}
	}
}

func TestDoesNotEqualTimed(t *testing.T) {
	uu := map[string]struct {
		fileName1, fileName2 string
		time                 time.Duration
		expected             bool
	}{
		"if a file can not be read in the appropriate time, it will return an error": {
			fileName1: "bible.txt",
			fileName2: "bible-modified.txt",
			time:      time.Millisecond * 0,
			expected:  false,
		},
	}

	for k, u := range uu {
		actual, err := f.AreEqualTimed(appendTestDir(u.fileName1), appendTestDir(u.fileName2), u.time)
		if err == nil {
			t.Fatal("The test case was supposed to fail with a time out error")
		}

		if u.expected != actual {
			t.Fatalf("test case %s failed with the following values:\nactual: %t, expected: %t\n", k, actual, u.expected)
		}
	}
}
