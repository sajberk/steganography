/*
@Time 2019-01-17 14:14
@Author HANG

*/
package stego

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {

	err := Encode(&StegoFile{CarrierFileName: "images/carrier.png", FileName: "images/ponyo.jpg"})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDecode(t *testing.T) {

	err := Decode(&StegoFile{CarrierFileName: "images/new_carrier.png"})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEncode_Decode_text(t *testing.T) {
	err := Encode(&StegoText{CarrierFileName: "images/carrier.png", TextContent: "I like it"})
	if err != nil {
		t.Fatal(err.Error())
	}
	err = Decode(&StegoText{CarrierFileName: "images/new_carrier.png"})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_getImageAsRGBA(t *testing.T) {
	_, imageFormat, err := getImageAsRGBA("images/carrier.png")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("image format >", imageFormat)

	_, imageFormat, err = getImageAsRGBA("images/ponyo.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("image format >", imageFormat)
}
