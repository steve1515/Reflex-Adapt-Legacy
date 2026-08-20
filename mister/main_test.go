package main

import "testing"

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