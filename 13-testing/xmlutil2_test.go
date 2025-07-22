package myutil

import (
	"encoding/xml"
	"reflect"
	"testing"
)

// normalizeXML parses an XML string into a structured format to enable reliable comparison.
func normalizeXML(input string) (interface{}, error) {
	var result interface{}
	if err := xml.Unmarshal([]byte(input), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func TestCompareXMLStrings(t *testing.T) {
	xmlStr1 := `<root><item id="1">Hello</item><item id="2">World</item></root>`
	xmlStr2 := `<root> <item id="1">Hello</item> <item id="2">World</item> </root>`

	norm1, err1 := normalizeXML(xmlStr1)
	if err1 != nil {
		t.Fatalf("Failed to parse first XML: %v", err1)
	}

	norm2, err2 := normalizeXML(xmlStr2)
	if err2 != nil {
		t.Fatalf("Failed to parse second XML: %v", err2)
	}

	if !reflect.DeepEqual(norm1, norm2) {
		t.Errorf("XML strings do not match\nExpected: %s\nGot: %s", xmlStr1, xmlStr2)
	}
}
