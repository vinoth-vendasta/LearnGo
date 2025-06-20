package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	contents, err := os.ReadFile("/Users/vnataraj/Desktop/LearnGo/File_Handling/text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))

	// **************************** using flag package to read file path ****************************
	fptr := flag.String("fpath", "text.txt", "File path to read")
	flag.Parse()
	filePath := *fptr
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// **************************** Writing to a file ****************************
	newFile, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err = newFile.WriteString("This is a new file created using os.Create()")
	if err != nil {
		panic(err)
	}

}
