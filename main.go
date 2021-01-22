package main

import (
	"debug/pe"
	"fmt"
	"os"
)

// variables
// var pfile pe.File

// var dataDir [16]pe.DataDirectory

// var sizeOptHeader32 = uint16(binary.Size(pe.OptionalHeader32{}))
// var sizeOptHeader64 = uint16(binary.Size(pe.OptionalHeader64{}))

// GET - PE Information Functions
// Thanks to https://gist.github.com/nokute78/46c1eb6a2f6050db4c5a87845dbdf87c
// func printOptionalHeader(f *pe.File) {
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

// Thanks to https://gist.github.com/josephspurrier/120d6f1c87fde4d9769e
// func printSections(f *pe.File) {
// 	fmt.Println("Sections:\n")
// 	for _, s := range f.Sections {
// 		fmt.Printf("%s\n", s.Name)
// 		fmt.Printf("%0#8x %s\n", s.VirtualSize, "Virtual Size")
// 		fmt.Printf("%0#8x %s\n", s.VirtualAddress, "Virtual Address")
// 		fmt.Printf("%0#8x %s\n", s.Size, "Size")
// 		fmt.Printf("%0#8x %s\n", s.Offset, "Offset")
// 		fmt.Printf("%0#8x %s\n", s.PointerToRelocations, "Pointer To Relocations")
// 		fmt.Printf("%0#8x %s\n", s.PointerToLineNumbers, "Pointer to Line Numbers")
// 		fmt.Printf("%0#8x %s\n", s.NumberOfRelocations, "Number of Relocations")
// 		fmt.Printf("%0#8x %s\n", s.NumberOfLineNumbers, "Number of Line Numbers")
// 		fmt.Printf("%0#8x %s\n", s.Characteristics, "Characteristics")
// 		fmt.Println()
// 	}
// }

func printSymbols(f *pe.File) {
	fmt.Println(":")
	fmt.Println(f.Symbols)

	fmt.Println("Symbols:\n")
	for _, s := range f.Symbols {
		fmt.Printf("%s\n", s.Name)
		fmt.Printf("%0#8x %s\n", s.Value, "Value")
		fmt.Printf("%0#8x %s\n", s.SectionNumber, "SectionNumber")
		fmt.Printf("%0#8x %s\n", s.Type, "Type")
		fmt.Printf("%0#8x %s\n", s.StorageClass, "StorageClass")
		fmt.Println()
	}

}

// func printStringTable(f *pe.File) {
// 	a := string(f.StringTable)
// 	b := strings.Replace(a, ".", "\n", -1)
// 	fmt.Println(b)
// }

// func printFileHeader(f *pe.File) {
// 	a := f.FileHeader.Characteristics
// 	fmt.Print("File Header: ")
// 	fmt.Println(a)
// }

func main() {
	// Open opens the named file using os.Open and prepares it for use as a PE binary.
	pfile, err := pe.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "pe.Open:%s\n", err)
		os.Exit(1)
	}
	defer pfile.Close()
	//printStringTable(pfile)
	//printFileHeader(pfile)

	//printSections(pfile)
	printSymbols(pfile)
	//printOptionalHeader(pfile)

}
