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

	expected = `XRX,Xerox
XLNX,Xilinx
AUY,Yamana Gold
YPF,YPF
YUM,Yum!
YUMC,Yum!
YY,YY
ZAYO,Zayo Group Ho!
ZBRA,Zebra Technolc
ZEN,Zendesk
Z,Zillow Group
ZBH,Zimmer Biome
ZION,Zions Bancorpc
ZTS,Zoetis
ZM,Zoom
ZS,Zscaler
ZTO,ZTO Express
ZNGA,Zynga
`
	csv = extract(in, out, langs, whT, whN)
	if csv.String() != expected {
		t.Errorf("Output CSV are not equal")
	}
}
