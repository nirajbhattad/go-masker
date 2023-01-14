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

// TestMaskNested tests Masking on a nested struct with json.
func TestMaskNestedJson(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `json:"userName" `
		Password  string   `json:"password" mask:""`
		DbSecrets []string `json:"dbSecrets" mask:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `json:"secret" mask:""`
		Keys     []string `json:"keys" mask:""`
		UserInfo []User   `json:"userInfo"`
	}

	users := &Users{
		Secret: "secret_sshhh",
		Keys:   []string{"key_1", "key_2", "key_3"},
		UserInfo: []User{
			{
				Username:  "Masking Test",
				Password:  "Masking_Password",
				DbSecrets: []string{"Masking_db_secret_1", "Masking_db_secret_2"},
			},
			{
				Username:  "Masking Test 2",
				Password:  "Masking_Password",
				DbSecrets: []string{"Masking_db_secret_1", "Masking_db_secret_2"},
			},
		},
	}

	usersMasked := &Users{
		Secret: "********",
		Keys:   []string{"********", "********", "********"},
		UserInfo: []User{
			{
				Username:  "Masking Test",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Masking Test 2",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
		},
	}

	validateJsonMasking(t, users, usersMasked)
}

// TestMaskNestedXml tests Masking on a nested struct with xml.
func TestMaskNestedXml(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `xml:"userName" `
		Password  string   `xml:"password" mask:""`
		DbSecrets []string `xml:"dbSecrets" mask:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `xml:"secret" mask:""`
		Keys     []string `xml:"keys" mask:""`
		UserInfo []User   `xml:"userInfo"`
	}

	users := &Users{
		Secret: "secret_sshhh",
		Keys:   []string{"key_1", "key_2", "key_3"},
		UserInfo: []User{
			{
				Username:  "Masking Test",
				Password:  "Masking_Password",
				DbSecrets: []string{"Masking_db_secret_1", "Masking_db_secret_2"},
			},
			{
				Username:  "Masking Test 2",
				Password:  "Masking_Password",
				DbSecrets: []string{"Masking_db_secret_1", "Masking_db_secret_2"},
			},
		},
	}

	usersMasked := &Users{
		Secret: "********",
		Keys:   []string{"********", "********", "********"},
		UserInfo: []User{
			{
				Username:  "Masking Test",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Masking Test 2",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
		},
	}

	validateXmlMasking(t, users, usersMasked)
}

// TestMaskEmptyPointer tests Masking on a empty pointer.
func TestMaskEmptyFields(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `json:"userName" mask:""`
		Password  string   `json:"password" mask:""`
		DbSecrets []string `json:"dbSecrets" mask:""`
	}

	user := &User{
		Username:  "",
		Password:  "Maskingpassword",
		DbSecrets: []string{},
	}

	userMasked := &User{
		Username:  "",
		Password:  "********",
		DbSecrets: []string{},
	}

	// Validate input with empty fields
	validateJsonMasking(t, user, userMasked)
}

func TestMaskEmptyPointer(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `json:"userName" mask:""`
		Password  string   `json:"password" mask:""`
		DbSecrets []string `json:"dbSecrets" mask:""`
	}

	// Validate empty pointer input
	var userEmpty *User
	validateJsonMasking(t, userEmpty, userEmpty)
}

// TestMaskNil tests redacting on a empty or nil input with json tags.
func TestMaskNil(t *testing.T) {

	t.Helper()

	got := Mask(nil, "mask")

	assert.Equal(t, "", got,
		"JSON representation mismatch after redacting fields")
}

// TestMaskNilXml tests redacting on a empty or nil input with xml tags.
func TestMaskNilXml(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `xml:"userName" mask:""`
		Password  string   `xml:"password" mask:""`
		DbSecrets []string `xml:"dbSecrets" mask:""`
	}

	user := &User{
		Username:  "",
		Password:  "Maskingpassword",
		DbSecrets: []string{},
	}

	userMasked := &User{
		Username:  "",
		Password:  "********",
		DbSecrets: []string{},
	}

	// Validate input with empty fields
	validateXmlMasking(t, user, userMasked)

	//  Validate nil input
	validateXmlMasking(t, nil, nil)
}

// TestMaskNestedNil tests redacting on a nested complex struct with
// some nil, empty and specified sensitive fields.
func TestMaskNestedNil(t *testing.T) {
	// Simple struct
	type User struct {
		Username  string   `json:"userName" `
		Password  string   `json:"password" mask:""`
		DbSecrets []string `json:"dbSecrets" mask:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `json:"secret" mask:""`
		Keys     []string `json:"keys" mask:""`
		UserInfo []User   `json:"userInfo"`
	}

	users := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Masking 1",
				Password:  "",
				DbSecrets: []string{"Masking_db_secret_1", "Masking_db_secret_2"},
			},
			{
				Username:  "Masking 2",
				Password:  "Masking_Password",
				DbSecrets: []string{},
			},
		},
	}

	userMasked := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Masking 1",
				Password:  "",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Masking 2",
				Password:  "********",
				DbSecrets: []string{},
			},
		},
	}

	validateJsonMasking(t, users, userMasked)
}

// TestMaskNestedXmlNil tests redacting on a nested complex struct with
// some nil, empty and specified sensitive fields.
func TestMaskNestedXmlNil(t *testing.T) {
	// Simple struct
	type User struct {
		Username  string   `xml:"userName" `
		Password  string   `xml:"password" mask:""`
		DbSecrets []string `xml:"dbSecrets" mask:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `xml:"secret" mask:""`
		Keys     []string `xml:"keys" mask:""`
		UserInfo []User   `xml:"userInfo"`
	}

	users := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Masking 1",
				Password:  "",
				DbSecrets: []string{"Masking_db_secret_1", "Masking_db_secret_2"},
			},
			{
				Username:  "Masking 2",
				Password:  "Masking_Password",
				DbSecrets: []string{},
			},
		},
	}

	userMasked := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Masking 1",
				Password:  "",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Masking 2",
				Password:  "********",
				DbSecrets: []string{},
			},
		},
	}

	validateXmlMasking(t, users, userMasked)
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
