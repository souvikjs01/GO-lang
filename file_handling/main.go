package main

import (
	"fmt"
	"os"
)

func main() {
	/* // file creating:
	file, err := os.Create("golang.txt")
	if err != nil {
		fmt.Println("error while file creating", err)
		return
	}
	// file created successfully-
	// close the resources-
	defer file.Close()

	fmt.Println("file created successfully")

	// write into file:
	content := "Hey there! i am using golang"

	byte, err2 := io.WriteString(file, content+"\n")
	if err2 != nil {
		fmt.Println("error while writing file", err2)
		return
	}
	fmt.Println("byte written", byte)

	*/

	// file reading:
	/*file, err := os.Open("golang.txt")
	if err != nil {
		fmt.Println("error while opening the file", err)
		return
	}
	defer file.Close()

	// creating buffer to read the file

	buffer := make([]byte, 1024)

	for {
		n, err3 := file.Read(buffer)
		if err3 == io.EOF {
			break
		}

		if err3 != nil {
			fmt.Println("error while reading the file", err3)
			return
		}

		// process the file

		fmt.Println(string(buffer[:n]))
	}
	*/

	// read the file into byte slices
	// not recomended for large size file-
	content, err := os.ReadFile("golang.txt")
	if err != nil {
		fmt.Println("error while reading the file")
		return
	}
	fmt.Println(string(content))

}
