package main

import (
	"debug/pe"
	"fmt"
	"os"
	"strings"
)

// variables
var pfile pe.File
var dataDir [16]pe.DataDirectory

// var sizeOptHeader32 = uint16(binary.Size(pe.OptionalHeader32{}))
// var sizeOptHeader64 = uint16(binary.Size(pe.OptionalHeader64{}))

// GET - PE Information Functions
// Thanks to https://gist.github.com/nokute78/46c1eb6a2f6050db4c5a87845dbdf87c
// func getOptionalHeader(f *pe.File) {
// 	fmt.Print("Optional Header Size: ")
// 	fmt.Println(sizeOptHeader32, " ", sizeOptHeader64)
// 	switch f.OptionalHeader.(type) {
// 	case *pe.OptionalHeader32:
// 		fmt.Fprintf(os.Stdout, "Header32\n")
// 		o := f.OptionalHeader.(*pe.OptionalHeader32)
// 		dataDir = o.DataDirectory
// 	case *pe.OptionalHeader64:
// 		fmt.Fprintf(os.Stdout, "Header64\n")
// 		o := f.OptionalHeader.(*pe.OptionalHeader64)
// 		dataDir = o.DataDirectory
// 	default:
// 		fmt.Fprintf(os.Stderr, "INVALID FORMAT\n")
// 		os.Exit(1)
// 	}
// }

func getSections(f *pe.File) {
	fmt.Println("Sections: ")
	fmt.Println(f.Sections)
}

func getSymbols(f *pe.File) {
	fmt.Println("Symbols:")
	fmt.Println(f.Symbols)
}

func getStrings(f *pe.File) {
	a := string(f.StringTable) // String extracts string from COFF string table st at offset start.
	b := strings.Replace(a, ".", "\n", -1)
	fmt.Println(b)
}

func main() {
	// Open opens the named file using os.Open and prepares it for use as a PE binary.
	pfile, err := pe.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "pe.Open:%s\n", err)
		os.Exit(1)
	}
	defer pfile.Close()

	//getOptionalHeader(pfile)                  // Optional Header Characteristics
	//fmt.Printf("DataDirectory:%v\n", dataDir) // Print Data Directory Information
	//getSections(pfile)                        // PE Sections
	//getSymbols(pfile)                         // PE Symbols
	getStrings(pfile) // PE Strings

}
