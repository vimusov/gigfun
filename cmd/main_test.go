package main

import "testing"

func TestGetFanSpeed(t *testing.T) {
	tests := []struct {
		name     string
		temp     int
		expected int
	}{
		{"TempBelow45_1", 0, 0},
		{"TempBelow45_2", 44, 0},
		{"Temp45", 45, 30},
		{"Temp75", 75, 100},
		{"TempAbove75_1", 76, 100},
		{"TempAbove75_2", 100, 100},

		{"Temp46_RoundDown", 46, 32},
		{"Temp51_RoundDown", 51, 44},

		{"Temp47_RoundUp", 47, 35},
		{"Temp57_RoundUp", 57, 58},

		{"Temp60", 60, 65},
		{"Temp74", 74, 98},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := getFanSpeed(tt.temp); actual != tt.expected {
				t.Errorf("getFanSpeed(%d) = %d, expected %d", tt.temp, actual, tt.expected)
			}
		})
	}
}
