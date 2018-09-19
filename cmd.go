package main

import (
	"fmt"
	"os"
)

/*
this function creates a temporary test file
writes into it all the files the functions
that need to be included and closes the file*/
func GetFunctionString(functions []string) string {
	var functionString string
	for _, val := range functions {
		functionString = fmt.Sprintf("%s%s(t)\n", functionString, val)
	}
	return functionString
}
func ImportString() string {
	var ImportString string
	ImportString = fmt.Sprintf("%s\n", reflect.TypeOf(Empty{}).PkgPath())
	return `import "testing"` + ImportString
}

//TODO: this function returns a list of all the functions that are called
//while testing in the current directory
func ListFunctions() []strings {

}

// TODO: this is a list read from .testignore file which lists all the
// functions which are to be ignored in the testing
func ExcludeList() map[string]bool {
	//read contents from file
	// covnert to string
	// functionString=string(flestring)
	exfunctions := strings.Split(functionString, "\n")
	exmap = make(map[string]bool)
	for key, val := range exfunctions {
		exmap[val] = true
	}
	return exmap
}

// TODO: writes contents to a file
func WriteContent(cont []byte) error {
	// open file
	f := os.Open("ko_test.go", os.Create || os.RDWR || os.Write, 655)
	// defer file close
	defer f.Close()
	// write to file
	err := f.Write(cont)
	if err != nil {
		return err
	}
	// close file
	return nil
}

// TODO: execute the go test command on the test file generated
func CmdExecTest() error {

}

func RemoveFile() error {
	// remove file
}

func main() {
	functions := ListFunctions()
	exfunctions := ExcludedList()
	preparefunctions = make([]string)
	for key, val := range functions {
		if !exfunctions[val] {
			preparefunctions = append(preparefunctions, val)
		}
	}
	fileContent := []byte(ImportString() + GetFunctionString(preparefunctions))
	err = WriteContent(fileContent)
	if err != nil {
		panic(err)
	}

	defer RemoveFile() // remove the file after testing
	err = CmdExecTest()

	if err != nil {
		panic(err)
	}

	return
}
