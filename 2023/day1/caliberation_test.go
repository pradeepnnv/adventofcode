package day1

import "testing"

func TestCaliberationValuePart1(t *testing.T) {
	want := 142
	list := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	got := CaliberationValuePart1(list)
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestCaliberationValuePart2(t *testing.T) {
	want := 281
	list := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	got := CaliberationValuePart2(list)
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}
