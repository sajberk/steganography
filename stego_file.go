/*
@Time 2019-01-19 15:17
@Author HANG

*/
package stego

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type StegoFile struct {
	CarrierFileName string
	FileName        string
}

func (s *StegoFile) GetCarrierFileName() string {
	return s.CarrierFileName
}

func (s *StegoFile) InputData() ([]byte, error) {
	_, file := path.Split(s.FileName)
	if len(file) > 255 {
		return nil, fmt.Errorf("file name too long")
	}
	data := []byte{uint8(len(file))}
	data = append(data, file...)
	//fmt.Println("file name:", string(data[1:data[0]+1]))

	bytes, err := ioutil.ReadFile(s.FileName)
	if err != nil {
		return nil, fmt.Errorf("read file err:%s", err.Error())
	}
	data = append(data, bytes...)
	return data, nil
}

func (s *StegoFile) OutputData(data []byte) error {
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
