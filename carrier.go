/*
@Time 2019-01-19 15:17
@Author zhcppy

*/
package stego

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type FileCarrier struct {
	CarrierFileName  string
	ResourceFileName string
}

func (s *FileCarrier) GetCarrierFileName() string {
	return s.CarrierFileName
}

func (s *FileCarrier) InputData() ([]byte, error) {
	_, file := path.Split(s.ResourceFileName)
	if len(file) > 255 {
		return nil, fmt.Errorf("file name too long")
	}
	data := []byte{uint8(len(file))}
	data = append(data, file...)
	//fmt.Println("file name:", string(data[1:data[0]+1]))

	bytes, err := ioutil.ReadFile(s.ResourceFileName)
	if err != nil {
		return nil, fmt.Errorf("read file err:%s", err.Error())
	}
	data = append(data, bytes...)
	return data, nil
}

func (s *FileCarrier) OutputData(data []byte) error {
	if len(data) < 0 || uint(len(data)) < uint(data[0]) {
		return fmt.Errorf("data format error")
	}
	//fmt.Println(data[1:data[0]+1])

	fileName := fmt.Sprintf("%s/new_%s", path.Dir(s.CarrierFileName), data[1:data[0]+1])
	resultFile, err := os.Create(fileName)
	defer resultFile.Close()
	if err != nil {
		return fmt.Errorf("error creating result file: %v", err)
	}
	_, err = resultFile.Write(data[data[0]+1:])
	return err
}

type TextCarrier struct {
	CarrierFileName string
	TextContent         string
}

func (s *TextCarrier) GetCarrierFileName() string {
	return s.CarrierFileName
}

func (s *TextCarrier) InputData() ([]byte, error) {
	return []byte(s.TextContent), nil
}

func (s *TextCarrier) OutputData(data []byte) error {
	fmt.Println("data info:", string(data))
	return nil
}
