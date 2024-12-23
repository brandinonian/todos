package main

import (
	"os"
	"testing"
)

func TestSearchDir(t *testing.T) {
	if err := os.MkdirAll("test/test/test", 0755); err != nil {
		t.Fatalf("Error creating test dir: %v", err)
	}

	pass1 := "test/test/pass1.txt"
	pass2 := "test/test/test/pass2.txt"
	fail1 := "test/test/test/fail1.txt"

	if err := os.WriteFile(pass1, []byte("TODO"), 0644); err != nil {
		t.Fatalf("Error creating pass1.txt: %v", err)
	}

	if err := os.WriteFile(pass2, []byte("TODO"), 0644); err != nil {
		t.Fatalf("Error creating pass1.txt: %v", err)
	}

	if err := os.WriteFile(fail1, []byte("FAIL"), 0644); err != nil {
		t.Fatalf("Error creating pass1.txt: %v", err)
	}

	defer t.Cleanup(func() {
		if err := os.RemoveAll("test"); err != nil {
			t.Fatalf("Error deleting test files: %v", err)
		}
	})

	err, results := SearchDir("test")
	if err != nil {
		t.Fatalf("Error calling SearchDir(): %v", err)
	}

	first := results[0]
	second := results[1]

	if len(results) != 2 || first != "pass1.txt" || second != "pass2.txt" {
		t.Fatalf("SearchDir() results: %v", results)
	}
}
