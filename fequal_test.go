package fequal_test

import (
	f "github.com/ru-lai/fequal"
	"strings"
	"testing"
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
