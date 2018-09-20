package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	//"reflect"
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

//TODO: this function returns the header that is to be written
// to the ko test
func ImportString() string {
	return `import "testing"`
}

//TODO: this function returns a list of all the functions that are called
//while testing in the current directory
func ListFunctions() []string {

	files := []string{"Test", "Test1", "Test2"}
	return files

}

/*
this function reads the file .testignore and reads
the function which are to be ignored*/
func ExcludeList() (map[string]bool, error) {
	//read contents from file
	f, err := os.Open(".testignore")
	if err != nil {
		return nil, err
	}
	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	functionString := string(fileContent)
	exfunctions := strings.Split(functionString, "\n")
	var exmap = make(map[string]bool)
	for _, val := range exfunctions {
		exmap[val] = true
	}
	return exmap, nil
}

/* write content to temporary go file for ko testing*/
func WriteContent(cont []byte) error {
	// open file
	f, err := os.OpenFile("ko_test.go", os.O_APPEND|os.O_RDWR|os.O_CREATE, 655)
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

/*invoking go test*/
func CmdExecTest() error {
	cmd := exec.Command("go", "test", "ko_test.go")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

/*remove file*/

func RemoveFile() error {
	_, err := os.Stat("ko_test.go")
	if os.IsNotExist(err) {
		return nil
	}
	cmd := exec.Command("rm", "ko_test.go")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	functions := ListFunctions()
	exfunctions, err := ExcludeList()
	if err != nil {
		panic(err)
	}
	var preparefunctions = make([]string, 1)
	for _, val := range functions {
		if !exfunctions[val] {
			preparefunctions = append(preparefunctions, val)
		}
	}
	fileContent := []byte(ImportString() + GetFunctionString(preparefunctions))
	err = WriteContent(fileContent)
	if err != nil {
		panic(err)
	}

	defer RemoveFile() // remove the file after checking if it exists
	err = CmdExecTest()

	if err != nil {
		panic(err)
	}

	return
}
