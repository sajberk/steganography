/*
@Time 2019-01-19 15:17
@Author HANG

*/
package stego

import "fmt"

type StegoText struct {
	CarrierFileName string
	TextContent     string
}

func (s *StegoText) GetCarrierFileName() string {
	return s.CarrierFileName
}

func (s *StegoText) InputData() ([]byte, error) {
	return []byte(s.TextContent), nil
}

func (s *StegoText) OutputData(data []byte) error {
	fmt.Println("data info:", string(data))
	return nil
}
