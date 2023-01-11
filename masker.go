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
