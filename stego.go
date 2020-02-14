package stego

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
	"path"
)

/*
 * 图像由图像源组成，图像源是有三原色组成，在计算机中一般表示为这样子：
 * R = 11011010
 * G = 10010111
 * B = 10010100
 *
 * LSB（最低有效位）Least Significant Bit 是指二进制数据中的最低位，权值为2^0
 *
 * LSB-Steganography 隐写技术 可以在图像中隐藏数据 通过将需要隐藏的数据替换最低有效位图像来达到隐藏的效果
 * 原因就是修改三原色的最低有效位，不会造成图片发生明显变化，想一想484很有意思，很神奇有木有
 *
 * 这里我通过修改三原色中的后两位，来达到在相同大小的图片中保存更多的数据
 * 图片中的前62位（即8个字节）用于保存需要隐藏数据的大小
 * 后面存放的数据可以自行实现接口定义存储结构，当然还包括读取了
 *
 * 这里面涉及到了一些位运算，自己动手写完后，感觉很有趣，印象深刻不少，修炼才刚刚开始。。。
 */

type Steganography interface {
	GetCarrierFileName() string // get image carrier
	InputData() ([]byte, error)
	OutputData(data []byte) error
}

func Encode(s Steganography) (err error) {
	rgba, imageFormat, err := getImageAsRGBA(s.GetCarrierFileName())
	if err != nil {
		return err
	}

	bytes, err := s.InputData()
	if err != nil {
		return err
	}
	if err = upColorSegment(rgba, readData(bytes)); err != nil {
		return err
	}

	dir, fileName := path.Split(s.GetCarrierFileName())
	resultFile, err := os.Create(path.Join(dir, "new_"+fileName))
	if err != nil {
		return fmt.Errorf("error creating result file: %v", err)
	}
	defer resultFile.Close()
	switch imageFormat {
	case "png", "jpeg":
		return png.Encode(resultFile, rgba)
	//case "jpeg":
	//	return jpeg.Encode(resultFile, rgba, nil)
	default:
		return fmt.Errorf("unsupported carrier format:%s", imageFormat)
	}
}

func Decode(s Steganography) (err error) {
	rgba, _, err := getImageAsRGBA(s.GetCarrierFileName())
	if err != nil {
		return err
	}

	var data []byte
	var dataSize, count uint64
loop:
	for x := 0; x < rgba.Bounds().Dx(); x++ {
		for y := 0; y < rgba.Bounds().Dy(); y++ {
			c := rgba.RGBAAt(x, y)
			if dataSize > 0 && count%4 == 0 && (count*3/4) >= (dataSize+8) {
				break loop
			}
			data = append(data, c.R&(3), c.G&(3), c.B&(3))
			count++
			if count%4 == 0 {
				data[len(data)-12] = data[len(data)-12]<<6 | data[len(data)-11]<<4 | data[len(data)-10]<<2 | data[len(data)-9]
				data[len(data)-11] = data[len(data)-8]<<6 | data[len(data)-7]<<4 | data[len(data)-6]<<2 | data[len(data)-5]
				data[len(data)-10] = data[len(data)-4]<<6 | data[len(data)-3]<<4 | data[len(data)-2]<<2 | data[len(data)-1]
				data = data[:len(data)-9]
				if len(data) == 9 {
					dataSize = binary.LittleEndian.Uint64(data[:8])
					data = data[8:]
				}
			}
		}
	}
	fmt.Println("data size:", dataSize, "data len:", len(data))
	return s.OutputData(data[:dataSize])
}

func upColorSegment(rgba *image.RGBA, data <-chan byte) error {
	var isReadFinish = false
	for x := 0; x < rgba.Bounds().Dx(); x++ {
		for y := 0; y < rgba.Bounds().Dy(); y++ {
			color := rgba.RGBAAt(x, y)
			for i := range []byte{color.R, color.G, color.B} {
				bt, ok := <-data
				if !ok && i == 0 {
					isReadFinish = true
					return nil
				}
				switch i {
				case 0:
					color.R = color.R&252 | bt
				case 1:
					color.G = color.G&252 | bt
				case 2:
					color.B = color.B&252 | bt
				}
			}
			rgba.SetRGBA(x, y, color)
		}
	}
	if !isReadFinish {
		return fmt.Errorf("data file too large for this carrier")
	}
	return nil
}

func getImageAsRGBA(carrierFileName string) (rgba *image.RGBA, imageFormat string, err error) {

	carrier, err := os.Open(carrierFileName)
	if err != nil {
		return
	}
	defer carrier.Close()

	var img image.Image
	img, imageFormat, err = image.Decode(carrier)
	if err != nil {
		return
	}
	if imageFormat != "png" && imageFormat != "jpeg" {
		err = fmt.Errorf("unsupported carrier format :%s", imageFormat)
		return
	}

	rgba = image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)
	return
}

func readData(data []byte) <-chan byte {
	out := make(chan byte, 1024)

	//fmt.Println(data[0], len(data))
	go func() {
		dataSize := make([]byte, 8)
		binary.LittleEndian.PutUint64(dataSize, uint64(len(data)))
		for _, bt := range append(dataSize, data...) {

			bts := []byte{
				bt & 192 >> 6,
				bt & 48 >> 4,
				bt & 12 >> 2,
				bt & 3,
			}
			for _, v := range bts {
				out <- v
			}
		}
		close(out)
	}()

	return out
}
