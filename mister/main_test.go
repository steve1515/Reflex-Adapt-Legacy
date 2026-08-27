package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRequiredUbootChanges(t *testing.T) {
	tests := []struct {
		name        string
		quirks      []string
		fastUSB     bool
		wantQuirk   bool
		wantFastUSB bool
	}{
		{name: "nothing configured", wantQuirk: true, wantFastUSB: true},
		{name: "quirk only", quirks: []string{adaptQuirks}, wantFastUSB: true},
		{name: "polling only", fastUSB: true, wantQuirk: true},
		{name: "fully configured", quirks: []string{adaptQuirks}, fastUSB: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotQuirk, gotFastUSB := requiredUbootChanges(test.quirks, test.fastUSB)
			if gotQuirk != test.wantQuirk || gotFastUSB != test.wantFastUSB {
				t.Fatalf("requiredUbootChanges() = (%v, %v), want (%v, %v)", gotQuirk, gotFastUSB, test.wantQuirk, test.wantFastUSB)
			}
		})
	}
}
func TestExtractAndCleanupUpdater(t *testing.T) {
	dir, err := extractUpdater()
	if err != nil {
		t.Fatalf("extractUpdater() error = %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "placeholder.txt")); err != nil {
		t.Fatalf("embedded placeholder was not extracted: %v", err)
	}
	if err := cleanupUpdater(dir); err != nil {
		t.Fatalf("cleanupUpdater() error = %v", err)
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Fatalf("temporary directory still exists after cleanup: %v", err)
	}
}
