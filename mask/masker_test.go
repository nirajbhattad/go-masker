package masker

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMaskSimple tests masking on a simple struct with json
func TestMaskSimple(t *testing.T) {

	// Simple struct
	type User struct {
		Username      string   `json:"userName" `
		Password      string   `json:"password" mask:""`
		DbSecrets     []string `json:"dbSecrets"`
		SIN           *int     `json:"sin" mask:""`
		SSN           int      `json:"ssn" mask:""`
		AccountNum    int32    `json:"accountNum" mask:""`
		AccountNumber []int64  `json:"accountNumber" mask:""`
		AccountPer    float32  `json:"accountPer" mask:""`
		AccountsPer   float64  `json:"accountsPer" mask:""`
	}
	sin := 1234568
	expectedSin := 0

	user := User{
		Username:      "MaskingJson",
		Password:      "codepassword",
		DbSecrets:     []string{"db_secret_1", "db_secret_2"},
		SIN:           &sin,
		SSN:           1234567,
		AccountNum:    1234,
		AccountNumber: []int64{888888, 123456},
		AccountPer:    23.056,
		AccountsPer:   12345.5678,
	}

	userMasked := &User{
		Username:      "MaskingJson",
		Password:      "********",
		DbSecrets:     []string{"db_secret_1", "db_secret_2"},
		SIN:           &expectedSin,
		SSN:           0,
		AccountNum:    0,
		AccountNumber: []int64{0, 0},
		AccountPer:    0.0,
		AccountsPer:   0.0,
	}

	validateJsonMasking(t, user, userMasked)
}

// TestMaskSimpleXml tests masking on a simple struct with xml
func TestMaskSimpleXml(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `xml:"userName" `
		Password  string   `xml:"password" mask:""`
		DbSecrets []string `xml:"dbSecrets"`
	}

	user := &User{
		Username:  "MaskXml",
		Password:  "codepassword",
		DbSecrets: []string{"db_secret_1", "db_secret_2"},
	}

	userMasked := &User{
		Username:  "MaskXml",
		Password:  "********",
		DbSecrets: []string{"db_secret_1", "db_secret_2"},
	}

	validateXmlMasking(t, user, userMasked)
}

// validateJsonMasking is a helper function to validate masking functionality on a struct with json tags.
func validateJsonMasking(t *testing.T, msg, maskedMsg interface{}) {
	t.Helper()

	// Get the masked string from Mask Library.
	got := Mask(msg, "mask")

	// Compare it against the given masked representaation.
	var b []byte
	b, _ = json.Marshal(maskedMsg)
	want := string(b)

	assert.Equal(t, want, got,
		"JSON representation mismatch after masking fields")
}

// validateXmlMasking is a helper function to validate masking functionality on a struct with xml tags.
func validateXmlMasking(t *testing.T, msg, maskedMsg interface{}) {
	t.Helper()

	// Get the masked string from Mask Library.
	got := MaskToXml(msg, "mask")

	// Compare it against the given masked representaation.
	var b []byte
	b, _ = xml.Marshal(maskedMsg)
	want := string(b)

	assert.Equal(t, want, got,
		"XML representation mismatch after masking fields")
}
