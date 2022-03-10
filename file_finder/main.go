package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Please give directory absolute path\n")
		return
	}
	filesInfo, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var names []byte
	for _, v := range filesInfo {
		if v.Size() == 0 {
			name := v.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}
	}
	errWriteFile := ioutil.WriteFile("emtpyfiles.txt", names, 0644)
	if errWriteFile != nil {
		fmt.Println(errWriteFile)
		return
	}

}
