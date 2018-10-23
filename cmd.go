package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	//"reflect"
	"regexp"
	"strings"
)

/*
this function creates a temporary test file
writes into it all the files the functions
that need to be included and closes the file*/

/* strings to echo into temp file*/
func GetFunctionString(functions []string) string {
	var functionString string
	for _, val := range functions {
		functionString = fmt.Sprintf("%s%s\n", functionString, val)
	}
	return functionString
}

//TODO: this function returns the header that is to be written
// to the ko test
func ImportString(testFile []byte) string {
	/*
		get the package name
	*/
	packageName := strings.Split(string(testFile), "\n")[0]
	/*
	   get the import string
	*/

	/*
		TODO
			rgx := regexp.MustCompile(`import\(\n[A-z,0-9]*\)`)
			imports := rgx.FindAllString(string(testfile), -1)
	*/
	importString := `import "testing"`
	return fmt.Sprintf("%s\n%s", packageName, importString)
}

//his function returns a list of all the functions that are called
//while testing in the current directory
func ListFunctions(fileContents string) map[string]string {
	//regex for functions
	//func\sTest[A-z]*\(.*?\)\s\{[^*]+\}
	//func\sTest[A-z]+\s\([A-z]\s\*testing.[A-z]\)\{\n[^\{]*\}

	//func\s[A-z]* //split the file using this and write the values as key value pair
	rgx := regexp.MustCompile(`func\s[A-z,0-9]*`)
	array := rgx.Split(fileContents, -1)
	functionNames := rgx.FindAllString(fileContents, -1)
	functionMaps := make(map[string]string, 1)
	for i := 1; i < len(functionNames); i++ {
		functionMaps[functionNames[i]] = array[i+1]
	}
	return functionMaps

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
		exmap["func "+val] = true
	}
	return exmap, nil
}

/* write content to temporary go file for ko testing*/
func WriteContent(cont []byte) error {
	// open file
	f, err := os.OpenFile("ko_test.go", os.O_APPEND|os.O_RDWR|os.O_CREATE, 777)
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
func readTestFile(filePath string) ([]byte, error) {
	//TODO: generic file name so that it takes all
	// ko files in the directory
	f, err := os.Open(filePath + "test.ko")
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func main() {
	test, err := readTestFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	funcMap := ListFunctions(string(test))
	exfunctions, err := ExcludeList()
	if err != nil {
		panic(err)
	}
	var preparefunctions = make([]string, 1)
	for key, val := range funcMap {
		if exfunctions[key] == true {
			continue
		} else {
			preparefunctions = append(preparefunctions, key+val)
		}
	}

	fileContent := []byte(ImportString(test) + GetFunctionString(preparefunctions))
	err = WriteContent(fileContent)
	if err != nil {
		panic(err)
	}

	//defer RemoveFile() // remove the file after checking if it exists
	err = CmdExecTest()

	if err != nil {
		panic(err)
	}

	return
}
