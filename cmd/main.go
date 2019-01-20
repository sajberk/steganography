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
		fmt.Fprintf(os.Stderr, "Operation must be specified")
		return
	}
	if *carrierFile == "" {
		fmt.Fprintf(os.Stderr, "Carrier file must be specified")
		return
	}
	if *dataFile == "" && *operation == "encode" {
		fmt.Fprintf(os.Stderr, "Data file must be specified")
		return
	}

	switch *operation {
	case "encode":
		err := stego.Encode(&stego.StegoFile{CarrierFileName: *carrierFile, FileName: *dataFile})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case "decode":
		err := stego.Decode(&stego.StegoFile{CarrierFileName: *carrierFile})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unsupported operation: %q", *operation)
	}

}
