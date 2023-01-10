package masker

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var Tag string

// Returns a pointer to the given input struct
func getStructPtr(obj interface{}) interface{} {

	// Check if the value is pointer
	reqAddrValue := reflect.ValueOf(obj)
	if reqAddrValue.Kind() == reflect.Ptr {
		return obj
	} else {
		// Create a new instance of the underlying type
		vp := reflect.New(reflect.TypeOf(obj))
		vp.Elem().Set(reflect.ValueOf(obj))
		return vp.Interface()
	}
}

func Masker(req interface{}, tag string) string {
	if req == nil {
		return ""
	}

	// Set the input tag
	Tag = tag

	out := getStructPtr(req)

	// Declare original values slice
	originalValues := make([]interface{}, 0)

	// Redact the json
	masker(out, &originalValues, true, false)

	// Create a json redaction copy
	var jsonRedaction []byte
	jsonRedaction, _ = json.Marshal(out)
	fmt.Println(string(jsonRedaction))

	masker(out, &originalValues, false, false)

	return string(jsonRedaction)
}

// masker is a recursive function which takes in the input object to mask and plays with values of it's fields
// and decides whether to mask the fields based on the input struct tag
func masker(req interface{}, originalValues *[]interface{}, save bool, isRedact bool) {

}
