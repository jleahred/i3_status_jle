package main

import "testing"

func TestReadFileSplitLines(t *testing.T) {
	lines := _ReadFileSplitLines("test_info/meminfo.test")

	nlines := 46
	if len(lines) != nlines {
		t.Errorf("Expected %d, got %d", nlines, len(lines))
	}
}

func TestParseFileCheckProperties(t *testing.T) {
	mapMemInfo := _MapFromMemInfoLines(_ReadFileSplitLines("test_info/meminfo.test"))

	mapCheck := map[string]int64{
		"MemTotal":     8086284,
		"MemFree":      2838076,
		"MemAvailable": 6397776,
		"SwapTotal":    5859324,
		"Cached":       363124,
	}

	for name, val := range mapCheck {
		_, ok := mapMemInfo[name]
		if !ok {
			t.Errorf("Looking for %s property", name)
		}
		if mapMemInfo[name] != val {
			t.Errorf("Checking %s, expected %d got %d", name, val, mapMemInfo[name])
		}
	}
}
