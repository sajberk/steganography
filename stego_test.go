package stego

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {

	err := Encode(&FileCarrier{CarrierFileName: "images/carrier.png", ResourceFileName: "images/ponyo.jpg"})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDecode(t *testing.T) {

	err := Decode(&FileCarrier{CarrierFileName: "images/new_carrier.png"})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEncode_Decode_text(t *testing.T) {
	err := Encode(&TextCarrier{CarrierFileName: "images/carrier.png", TextContent: "I like it"})
	if err != nil {
		t.Fatal(err.Error())
	}
	err = Decode(&TextCarrier{CarrierFileName: "images/new_carrier.png"})
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
