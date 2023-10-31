package main

import (
	"flag"
	"fmt"
	"os"
)

var ENCRYPT_CHUNK = 16
var key = "this_must_be_of_32_byte_length!!"


func main() {
	in := flag.String("inputFilePath", "inputfile.txt", "string")
	out := flag.String("outputFilePath", "outputfile.txt", "string")
	enc := flag.Bool("enc", false, "a bool")
	flag.Parse()
	fmt.Println(*enc)

	input, err := os.Open(*in)
		if err != nil {
			fmt.Println(err.Error())
		}
		output, err := os.Create(*out)
		if err != nil {
			fmt.Println(err.Error())
		}

	if *enc == true {	
		EncryptFile(input, output)
	} else {
		DecryptFile(input, output)
	}

	defer input.Close()
	defer output.Close()
}