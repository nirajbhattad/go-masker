package main

import (
	"fmt"
	mask "go-masker/mask"
)

// Mask Is the struct tag to select fields which are to be masked
type ExampleJson struct {
	Name          string   `json:"name"`
	SSN           *int     `json:"ssn" mask:""`
	StudentNumber *[]int32 `json:"studentNumber"`
	AccountNumber []int64  `json:"accountNumber" mask:""`
	Percentage    float32  `json:"percentage" mask:""`
}

type ExampleXml struct {
	Name          string   `xml:"name"`
	SSN           *int     `xml:"ssn" mask:""`
	StudentNumber *[]int32 `xml:"studentNumber"`
	AccountNumber []int64  `xml:"accountNumber" mask:""`
	Percentage    float32  `xml:"percentage" mask:""`
}

func main() {
	ssn := 1234567
	ssnarray := []int32{1234}
	jsonInput := ExampleJson{
		Name:          "niraj",
		SSN:           &ssn,
		StudentNumber: &ssnarray,
		AccountNumber: []int64{888888, 123456},
		Percentage:    23.056,
	}

	xmlInput := ExampleXml{
		Name:          "niraj",
		SSN:           &ssn,
		StudentNumber: &ssnarray,
		AccountNumber: []int64{888888, 123456},
		Percentage:    23.056,
	}

	// Prints Masked Json
	maskedJson := mask.Mask(jsonInput, "mask")
	// {"name":"niraj","ssn":0,"studentNumber":[1234],"accountNumber":[0,0],"percentage":0}

	// Prints Masked XML
	maskedXml := mask.MaskToXml(xmlInput, "mask")
	// <ExampleXml><name>niraj</name><ssn>0</ssn><studentNumber>1234</studentNumber><accountNumber>0</accountNumber><accountNumber>0</accountNumber><percentage>0</percentage></ExampleXml>

	fmt.Println(maskedJson)
	fmt.Println(maskedXml)
}
