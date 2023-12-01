package day2

import "testing"

func TestValidatePasswords(t *testing.T) {
	want := 1
	got := ValidatePasswords([]string{"a"})
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
