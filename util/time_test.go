package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Error("Year is not 2019")
	}

	if convertedTime.Month() != 02 {
		t.Error("Year is not 02")
	}

	if convertedTime.Hour() != 10 {
		t.Error("Year is not 10")
	}

}