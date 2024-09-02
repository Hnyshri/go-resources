package main

import (
	"bufio"
	"fmt"
	"os"
)

func filesFunction() {
	file, err := os.Open("files/readFile.txt")

	fmt.Println(err)

	if err != nil {
		panic(err)
	}

	defer file.Close() // Close file after open file

	fmt.Println(file.Name())
	fmt.Println(file.Stat())

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name", fileInfo.Name())
	fmt.Println("File Size", fileInfo.Size())
	fmt.Println("File or floder", fileInfo.IsDir())
	fmt.Println("File Last Modify time", fileInfo.ModTime())
	fmt.Println("File Permission", fileInfo.Mode())

	buffer := make([]byte, fileInfo.Size())
	readFile, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}

	println("File readFile", readFile, buffer, string(buffer))

	// ======================== Method 2 to Read File =======================

	// Not use for Large file or it is not good pratice
	fileData, err := os.ReadFile("files/readFile.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Method 2 to Read File ", string(fileData))

	// ======================== Read folder =======================

	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	dirInfo, err := dir.ReadDir(-1) // -1 means get all files // if 1 then get only 1 file and 2 then 2 files so on
	if err != nil {
		panic(err)
	}
	for _, value := range dirInfo {
		fmt.Println("Read folder ", value.Name(), value.IsDir())
	}

	// ======================== Write New File =======================

	writeFile, err := os.Create("files/writeFile.txt")
	if err != nil {
		panic(err)
	}
	defer writeFile.Close()

	// Write file simple way
	writeFile.WriteString("How are you") // writeString is use for append data
	// // Reset file pointer to the beginning of the file
	// writeFile.Seek(0, 0)
	writeFile.WriteString("Shrityash")

	// Write file bytes way
	bytes := []byte("Hello")
	fmt.Println("Write New File ", bytes)
	writeFile.Write(bytes)

	// ====================================Read and Write to another File using streaming fashion =======================

	sourceFile, err := os.Open("files/readFile.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create("files/streamingFashion.txt")
	if err != nil {
		panic(err)
	}
	defer destinationFile.Close()

	reader := bufio.NewReader(sourceFile)      // defaultBufSize = 4096
	writer := bufio.NewWriter(destinationFile) // defaultBufSize = 4096

	for { // infinite loop
		byteData, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" { // EOF => End of File
				panic(err)
			}
			break
		}

		err2 := writer.WriteByte(byteData)
		if err2 != nil {
			panic(err2)
		}
	}
	writer.Flush()
	fmt.Println("streamingFashion Created")

	// ====================================Delete the File =======================
	err3 := os.Remove("files/streamingFashion.txt")
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("File Deleted Done")
}
