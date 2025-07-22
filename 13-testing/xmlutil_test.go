package myutil

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define struct for XML comparison
type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func TestCompareXMLUsingStruct(t *testing.T) {
	xml1 := `<person><name>John Doe</name><age>30</age></person>`
	xml2 := `<person> <name>John Doe</name> <age>30</age> </person>`

	var p1, p2 Person

	// Unmarshal both XMLs into structs
	err1 := xml.Unmarshal([]byte(xml1), &p1)
	err2 := xml.Unmarshal([]byte(xml2), &p2)

	// Ensure unmarshaling didn't fail
	assert.NoError(t, err1, "Unmarshaling first XML should not return an error")
	assert.NoError(t, err2, "Unmarshaling second XML should not return an error")

	// Compare the structs instead of raw XML strings
	assert.Equal(t, p1, p2, "The two XMLs should be equivalent")
}
