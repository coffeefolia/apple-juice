package main

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"os"
)

// variables
var dataDir [16]pe.DataDirectory
var sizeOptHeader32 = uint16(binary.Size(pe.OptionalHeader32{}))
var sizeOptHeader64 = uint16(binary.Size(pe.OptionalHeader64{}))
var file pe.File

func main() {
	// Open opens the named file using os.Open and prepares it for use as a PE binary.
	file, err := pe.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "pe.Open:%s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	getOptionalHeader(file)                   // Optional Header Characteristics
	fmt.Printf("DataDirectory:%v\n", dataDir) // Print Data Directory Information
	getSections(file)                         // PE Sections
	getSymbols(file)                          // PE Symbols
	getStrings(file)                          // PE Strings
}

// GET - PE Information Functions
func getOptionalHeader(f *pe.File) {
	fmt.Print("Optional Header Size: ")
	fmt.Println(sizeOptHeader32, " ", sizeOptHeader64)
	switch f.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		fmt.Fprintf(os.Stdout, "Header32\n")
		o := f.OptionalHeader.(*pe.OptionalHeader32)
		dataDir = o.DataDirectory
	case *pe.OptionalHeader64:
		fmt.Fprintf(os.Stdout, "Header64\n")
		o := f.OptionalHeader.(*pe.OptionalHeader64)
		dataDir = o.DataDirectory
	default:
		fmt.Fprintf(os.Stderr, "INVALID FORMAT\n")
		os.Exit(1)
	}
}

func getSections(f *pe.File) {
	fmt.Println("Sections: ")
	fmt.Println(f.Sections)
}

func getSymbols(f *pe.File) {
	fmt.Println("Symbols:")
	fmt.Println(f.Symbols)
}

func getStrings(f *pe.File) {
	fmt.Println("Strings:\n", f.StringTable)
	fmt.Println()
}
