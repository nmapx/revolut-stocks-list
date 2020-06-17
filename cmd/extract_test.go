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

	expected = `SCCO,Southern Copper
HEI,Heico
SIRI,Sirius XM
HES,Hess Corporation
COLD,Americold Realty Trus
NRG,NRG Energy
HFC,HollyFrontier
SGMO,Sangamo Therapeutic
AZO,AutoZone
HGV,Hilton Grand Vacation
WEX,WEX
NUE,Nucor
SENS,Senseonics Holdings
DAN,Dana Holding
STLD,Steel Dynamics
AMTD,TD Ameritrade
DNKN,Dunkin' Brands Group
EQNR,Equinor
NWL,Newell Rubbermaid
FHN,First Horizon National...Corporation
URBN,Urban Outfitters
PSTG,Pure Storage
SBH,Sally Beauty
`
	csv, _ = extract(in, out, langs, whT, whN)
	if csv.String() != expected {
		t.Errorf("Output CSV are not equal")
	}
}
