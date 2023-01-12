package main

import (
	"fmt"
	mask "go-masker/mask"
)

type ExampleJson struct {
	Name          string   `json:"name" mask:""`
	SSN           *int     `json:"ssn" mask:""`
	StudentNumber *[]int32 `json:"studentNumber" mask:""`
	AccountNumber []int64  `json:"accountNumber" mask:""`
	Percentage    float32  `json:"percentage" mask:""`
}

type ExampleXml struct {
	Name          string   `xml:"name" mask:""`
	SSN           *int     `xml:"ssn" mask:""`
	StudentNumber *[]int32 `xml:"studentNumber" mask:""`
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

	maskedJson := mask.Mask(jsonInput, "mask")
	maskedXml := mask.MaskToXml(xmlInput, "mask")

	fmt.Println(maskedJson)
	fmt.Println(maskedXml)
}
