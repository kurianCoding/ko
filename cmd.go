package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

/*
this function creates a temporary test file
writes into it all the files the functions
that need to be included and closes the file*/

/* strings to echo into temp file*/
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

/*
this function reads the file .testignore and reads
the function which are to be ignored*/
func ExcludeList() map[string]bool {
	//read contents from file
	f := os.Open(".testignore", os.READ)
	fileContent, err := ioutil.ReadAll(f)
	functionString = string(fileContent)
	exfunctions := strings.Split(functionString, "\n")
	exmap = make(map[string]bool)
	for key, val := range exfunctions {
		exmap[val] = true
	}
	return exmap
}

func WriteContent(cont []byte) error {
	// open file
	f, err := os.Open("ko_test.go", os.Create || os.RDWR || os.Write, 655)
	if err != nil {
		return err
	}
	// defer file close
	defer f.Close()
	// write to file
	_, err = f.Write(cont)
	if err != nil {
		return err
	}
	// close file
	return nil
}

func CmdExecTest() error {
	cmd := exec.Command("go", "test", "ko_test.go")
	err := cmd.Run()
	if err != nil {
		return err
	}
}

func RemoveFile() error {
	// remove file
	cmd := exec.Command("rm", "ko_test.go")
	err := cmd.Run()
	if err != nil {
		return err
	}
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
