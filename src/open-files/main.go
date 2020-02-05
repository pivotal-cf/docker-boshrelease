package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const MAX int = 100000

func main() {
	var files = make([]*os.File, MAX)
	defer closeFiles(files)
	for i := 0; i < MAX; i++ {
		file, err := ioutil.TempFile("", "ulimit-test-")
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Failed to open file %d, see error above\n", i)
			return
		}
		files[i] = file
	}
}

func closeFiles(files []*os.File) {
	for i := 0; i < len(files); i++ {
		if files[i] != nil {
			err := files[i].Close()
			if err != nil {
				fmt.Println(err)
				fmt.Println("Failed to close file, see error above")
			}
			err = os.Remove(files[i].Name())
			if err != nil {
				fmt.Println(err)
				fmt.Println("Failed to remove file, see error above")
			}
		} else {
			return
		}
	}
}