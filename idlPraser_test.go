package idl

import (
	"strings"
	"testing"
)

func TestParaserNormalData(t *testing.T) {
	data := `asd:
        - go
        - a`
	sr := strings.NewReader(data)
	parser := NewSimpleIdlParser(sr)
	infors := parser.ParseInfomations()
	if len(infors) == 1 {
		if infors[0].Name != "asd" {
			t.Error("infos[0].Name:" + infors[0].Name)
		}
	} else {
		t.Error("info len:", len(infors))
	}
}

func TestRawString(t *testing.T) {
	data := `" asd ":
        - go
        - "  a "`
	infos := parserString(data)
	if infos[0].Name != " asd " {
		t.Error("info name:", infos[0].Name)
	}
	attrs := infos[0].Attrs
	if attrs[1].GetValue() != "  a " {
		t.Error("attr 2:", attrs[1].GetValue())
	}
}

func parserString(data string) []*Information {
	sr := strings.NewReader(data)
	parser := NewSimpleIdlParser(sr)
	infors := parser.ParseInfomations()
	return infors
}
