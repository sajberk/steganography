package main

import (
	"flag"
	"fmt"
	"github.com/zhcppy/steganography"
	"os"
)

func main() {
	var (
		operation   = flag.String("op", "", "operation (one of the following: encode, decode)")
		carrierFile = flag.String("carrier", "", "carrier file in which the data is encoded")
		dataFile    = flag.String("data", "", "data file which is being encoded in carrier")
	)
	flag.Parse()

	if *operation == "" {
		fmt.Fprintln(os.Stderr, "Operation must be specified")
		return
	}
	if *carrierFile == "" {
		fmt.Fprintln(os.Stderr, "Carrier file must be specified")
		return
	}
	if *dataFile == "" && *operation == "encode" {
		fmt.Fprintln(os.Stderr, "Data file must be specified")
		return
	}

	switch *operation {
	case "encode":
		err := stego.Encode(&stego.FileCarrier{CarrierFileName: *carrierFile, ResourceFileName: *dataFile})
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	case "decode":
		err := stego.Decode(&stego.FileCarrier{CarrierFileName: *carrierFile})
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	default:
		fmt.Fprintf(os.Stderr, "Unsupported operation: %s", *operation)
	}

}
