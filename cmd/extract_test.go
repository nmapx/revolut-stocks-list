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
	rootDir := filepath.Dir(filename) + "/.."

	var csv *bytes.Buffer
	var expected string
	var in, langs []string
	in = append(in, rootDir+"/example-input.jpg")
	langs = append(langs, "eng")
	out := rootDir + "/output_test.csv"
	whT := "QWERTYUIOPASDFGHJKLZXCVBNM."
	whN := ""

	expected = `JMIA,Jumia
SPG,Simon Property
COTY,Coty
CLNY,Colony Capital
`
	csv, _ = extract(in, out, langs, whT, whN)
	if csv.String() != expected {
		t.Errorf("Output CSV are not equal")
	}
}
