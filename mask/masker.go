package masker

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
)

var Tag string

// Masks the input struct/interface to a json string
func Mask(req interface{}, tag string) string {
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

	masker(out, &originalValues, false, false)

	return string(jsonRedaction)
}

// Masks the input struct/interface to a xml
func MaskToXml(req interface{}, tag string) string {
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

	// Create a xml redaction copy
	var xmlRedaction []byte
	xmlRedaction, _ = xml.Marshal(out)

	masker(out, &originalValues, false, false)

	return string(xmlRedaction)
}

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

/*
This Go function is designed to take in an input in the form of an interface, as well as a flag to
either save or restore the original values, and a flag to determine whether the values should be masked.

It uses the Go "reflect" package to recursively iterate through the input and its fields,
which may be either structs or slices.

The first step in the function is to check that the input is a pointer,
and if not, the function immediately exits.

It then checks if the input is a struct or a slice, and if so,
it recursively iterates through the fields or elements of the struct or slice.

If the input is not a struct or a slice, the function checks if it is a string, int or int32,
if yes and the save flag is true, it saves the input value in the originalValues slice, and
depending on the isRedact flag, it will either be set to a default value or not.
If the save flag is false, it will restore the original value.

The function uses a Tag variable which is used to check whether the current field should be masked or not.
The tag is used to check for a specific metadata on the field and determine whether the field should be
masked or not.
*/
func masker(req interface{}, originalValues *[]interface{}, save bool, isRedact bool) {
	// if target is not pointer, then immediately return
	// modifying struct's field requires addressable object
	reqAddrValue := reflect.ValueOf(req)
	if reqAddrValue.Kind() != reflect.Ptr {
		return
	}

	inputValue := reqAddrValue.Elem()
	if !inputValue.IsValid() {
		return
	}

	requestType := inputValue.Type()

	// If the input is passed by pointer, dereference it to get the underlying value
	if inputValue.Kind() == reflect.Ptr && !inputValue.IsNil() {
		inputValue = inputValue.Elem()
		if !inputValue.IsValid() {
			return
		}

		requestType = inputValue.Type()
	}

	// Validate if the input is struct type
	if requestType.Kind() == reflect.Struct {
		// If target is a struct then recurse on each of its field.
		for i := 0; i < requestType.NumField(); i++ {
			fieldType := requestType.Field(i)
			fValue := inputValue.Field(i)

			_, shouldRedact := fieldType.Tag.Lookup(Tag)
			if !fValue.IsValid() {
				continue
			}

			if !fValue.CanAddr() {
				// Cannot take pointer of this field, so can't scrub it.
				continue
			}

			if !fValue.Addr().CanInterface() {
				continue
			}
			masker(fValue.Addr().Interface(), originalValues, save, shouldRedact)
		}
		return
	}

	// Validate if the input is slice type
	if requestType.Kind() == reflect.Array || requestType.Kind() == reflect.Slice {
		for i := 0; i < inputValue.Len(); i++ {
			arrValue := inputValue.Index(i)
			if !arrValue.IsValid() {
				continue
			}

			if !arrValue.CanAddr() {
				continue
			}

			if !arrValue.Addr().CanInterface() {
				continue
			}
			masker(arrValue.Addr().Interface(), originalValues, save, isRedact)
		}

		return
	}

	// Base Condition To Return From Recursive Function
	if inputValue.CanSet() && inputValue.Kind() == reflect.String && !inputValue.IsZero() {
		if save {
			*originalValues = append(*originalValues, inputValue.String())
			if isRedact {
				inputValue.SetString("********")
			}
		} else {
			inputValue.SetString((*originalValues)[0].(string))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanInt() && inputValue.Kind() == reflect.Int {
		if save {
			*originalValues = append(*originalValues, inputValue.Int())
			if isRedact {
				inputValue.SetInt(0)
			}
		} else {
			temp := int((*originalValues)[0].(int64))
			inputValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanInt() && inputValue.Kind() == reflect.Int32 {
		if save {
			*originalValues = append(*originalValues, inputValue.Int())
			if isRedact {
				inputValue.SetInt(0)
			}
		} else {
			temp := int32((*originalValues)[0].(int64))
			inputValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanInt() && inputValue.Kind() == reflect.Int64 {
		if save {
			*originalValues = append(*originalValues, inputValue.Int())
			if isRedact {
				inputValue.SetInt(0)
			}
		} else {
			inputValue.SetInt((*originalValues)[0].(int64))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanFloat() && inputValue.Kind() == reflect.Float64 {
		if save {
			*originalValues = append(*originalValues, inputValue.Float())
			if isRedact {
				inputValue.SetFloat(0.0)
			}
		} else {
			inputValue.SetFloat((*originalValues)[0].(float64))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanFloat() && inputValue.Kind() == reflect.Float32 {
		if save {
			*originalValues = append(*originalValues, inputValue.Float())
			if isRedact {
				inputValue.SetFloat(0.0)
			}
		} else {
			temp := float32((*originalValues)[0].(float64))
			inputValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	}
}
