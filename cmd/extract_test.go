package cmd

import (
	"bytes"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestUnify(t *testing.T) {
	var slice, expected []string
	slice = unify(append(slice, "a", ",", "b", "", "c"))
	expected = append(expected, "a", "b", "c")
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Slices are not equal")
	}
}

func TestExtract(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	cacheDir := filepath.Dir(filename) + "/../cache"
	rootDir := filepath.Dir(filename) + "/.."

	var csv *bytes.Buffer
	var expected string
	var in, langs []string
	in = append(in, rootDir+"/example-input.jpg")
	langs = append(langs, "eng")
	out := cacheDir + "/test_output"
	whT := "QWERTYUIOPASDFGHJKLZXCVBNM"
	whN := ""

	expected = `WAB,WABTEC Corpor
WBA,Walgreens Boot:
WMT,Walmart
DIS,Walt Disney
WM,Waste Managen
W,Wayfair
WB,Weibo
WFC,Wells Fargo
WELL,Welltower
WEN,Wendy's Compal
WDC,Western Digital
WU,Western Union
WRK,WestRock
WEX,WEX
WY,Weyerhaeuser C
WLL,Whiting PetroleL
`
	csv = extract(in, out, langs, whT, whN)
	if csv.String() != expected {
		t.Errorf("Output CSV are not equal")
	}
}
