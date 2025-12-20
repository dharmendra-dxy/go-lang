package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// OPEN THE FILE
	// f, err:= os.Open("example.txt")
	// if err != nil {
	// 	panic(err) // log|panic the error
	// }

	// GET INFO ABOUT THE FILE:

	// fileInfo, err:= f.Stat()
	// if err != nil {
	// 	panic(err) // log|panic the error
	// }

	// fmt.Println("filename: ", fileInfo.Name())
	// fmt.Println("is_folder: ", fileInfo.IsDir())
	// fmt.Println("file size: ", fileInfo.Size())
	// fmt.Println("file permission: ", fileInfo.Mode())
	// fmt.Println("file modified at: ", fileInfo.ModTime())

	// READ THE FILE:

	// 1. Open the file

	// Close the file (defer)
	// defer f.Close()

	// buff:= make([]byte, 20) // buff array

	// d, err:=f.Read(buff)
	// if err!= nil{
	// 	panic(err)
	// }

	// fmt.Println("d: ", d)
	// fmt.Println("DATA : buff: ");

	// for i:=0; i< len(buff); i++{
	// 	fmt.Printf(string(buff[i]), " ")
	// }

	/*
		----------------------------------
		ALTERNATIVE WAY TO READ THE FILE:
		1. No need to open/close the file
		2. Only good for small files, not fit for large files
		----------------------------------
	*/

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("DATA: ", string(data))

	/*
		-------------------------
			READING A FOLDER:
		-------------------------
	*/

	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1) // fileInfo is a slice :: -1 gives all of the files
	if err != nil {
		panic(err)
	}

	fmt.Println("FileINFO: ")
	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}

	/*
		-------------------------
			CREATING  A FILE:
		-------------------------
	*/
	newFile, err := os.Create("exmaple2.txt")
	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	// newFile.WriteString("Hello Go Lang walon !! KAISE HO ??" ) // it appends in the file, do not overwrites

	// add bytes in file:
	// bytes := []byte("HELLO_GOLANG")
	// newFile.Write(bytes)

	/*
		------------------------------------
		STREAMING
		Copy content of exmaple.txt and paste to example3.txt
		------------------------------------
	*/

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example3.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	// create reader & writer using bufio
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	// infinite loop:
	for {
		b, err := reader.ReadByte()
		if err != nil {
			//handle when end of file:
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	//flush:
	writer.Flush()

	fmt.Println("WRITTEN TO NEW FILE SUCCESSFULLY")

}
