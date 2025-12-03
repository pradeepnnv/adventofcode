package day1

import "testing"

func TestSecretEntranceTest1(t *testing.T) {
	input := "../inputs/day1_test.txt"
	want := 3
	got := SecretEntrancePart1(input, false)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSecretEntranceTest2(t *testing.T) {
	input := "../inputs/day1_test.txt"
	want := 6
	got := SecretEntrancePart2(input)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
