package testflow

//go:generate go run main.go

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const templateFile = "./templates/_test.gotpl"

type methodProperty struct {
	Name      string
	Arguments []argument
	Responses []response
}

type argument struct {
	Name string
	Type string
}

type response struct {
	Name string
	Type string
}

func GenerateTestTemplate(targetPath string) {
	projectDir, _ := os.Getwd()
	fullTargetPath := fmt.Sprintf("%s/%s", projectDir, targetPath)
	packageName := targetPath[strings.LastIndex(targetPath, "/")+1:]
	serviceFiles, err := ioutil.ReadDir(fullTargetPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range serviceFiles {
		fileName := file.Name()
		if strings.Contains(fileName, "_test.go") || !strings.Contains(fileName, ".go") {
			// skip already generated test file or file except go
			continue
		}
		fileAlias := strings.Split(fileName, ".")
		fileNameNoExtension := fileAlias[0]
		outFilePath := fmt.Sprintf("%s/%s_test.go", fullTargetPath, fileNameNoExtension)
		methodProperties := parseFuncSignature(fmt.Sprintf("%s/%s", targetPath, fileName))
		if !existFile(outFilePath) {
			// not yet generate file
			generateTestFile(packageName, fileNameNoExtension, outFilePath, methodProperties)
			fmt.Println(fmt.Sprintf("Create Test File: %s", outFilePath))
		}
	}
}

// parseFuncSignature extracts struct name, argument property(param name, type), response property.
func parseFuncSignature(inFilePath string) []methodProperty {

	methodProperties := make([]methodProperty, 0)
	fp, err := os.Open(inFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	rep := regexp.MustCompile(`func \((.+)\) (.+)\((.*)\) (\((.*)\)|(.*)) {$`)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "func") {
			result := rep.FindStringSubmatch(line)
			if len(result) == 7 {
				// only struct method signature
				methodName := result[2]
				arguments := parseArgument(result[3])
				var dirtyResponses string
				if strings.Contains(result[4], "(") {
					// ex　func a() (a, p int), func a() (err error)
					dirtyResponses = result[5]
				} else {
					// ex) func a() error
					dirtyResponses = result[6]
				}
				responses := parseReturn(dirtyResponses)
				methodProperties = append(methodProperties, methodProperty{
					Name:      methodName,
					Arguments: arguments,
					Responses: responses,
				})
			}
		}
	}

	return methodProperties
}

// parseArgument returns arguments property(name, type) for struct method.
func parseArgument(dirtyArguments string) []argument {
	var arguments []argument

	if len(dirtyArguments) == 0 {
		return []argument{}
	}

	nameAndTypes := strings.Split(dirtyArguments, ", ")
	indexToTypeMap := make(map[int]string)
	shortTypeIndexes := make([]int, 0)

	//　func create(width int, height int) →　func create(width, height int)
	for i, nameAndType := range nameAndTypes {
		temp := strings.Split(nameAndType, " ")
		if len(temp) == 1 {
			// if omit type
			shortTypeIndexes = append(shortTypeIndexes, i)
		} else {
			// if exist type
			for _, index := range shortTypeIndexes {
				indexToTypeMap[index] = temp[1]
			}
			shortTypeIndexes = []int{}
		}
	}

	for i, nameAndType := range nameAndTypes {
		temp := strings.Split(nameAndType, " ")
		if len(temp) == 1 {
			arguments = append(arguments, argument{
				Name: temp[0],
				Type: indexToTypeMap[i],
			})
		} else {
			arguments = append(arguments, argument{
				Name: temp[0],
				Type: temp[1],
			})
		}
	}

	return arguments
}

// parseReturn returns ret values for strcut method
func parseReturn(dirtyResponses string) []response {
	var responses []response
	nameAndTypes := strings.Split(dirtyResponses, ", ")
	indexToTypeMap := make(map[int]string) // key: omit param index, value: type name
	shortTypeIndexes := make([]int, 0)
	returnOnly := true

	for _, nameAndType := range nameAndTypes {
		temp := strings.Split(nameAndType, " ")
		if len(temp) > 1 {
			returnOnly = false
		}
	}

	if !returnOnly {
		//　func a() (p int,q int)→　func a() (p, q int)
		for i, nameAndType := range nameAndTypes {
			temp := strings.Split(nameAndType, " ")
			if len(temp) == 1 {
				// if omit type
				shortTypeIndexes = append(shortTypeIndexes, i)
			} else {
				// if exist type
				for _, index := range shortTypeIndexes {
					// assign type to omit type param name
					indexToTypeMap[index] = temp[1]
				}
				shortTypeIndexes = []int{}
			}
		}
	}

	for i, nameAndType := range nameAndTypes {
		temp := strings.Split(nameAndType, " ")
		if returnOnly {
			responses = append(responses, response{
				Name: "",
				Type: temp[0],
			})
		} else if len(temp) == 1 {
			responses = append(responses, response{
				Name: temp[0],
				Type: indexToTypeMap[i],
			})
		} else {
			responses = append(responses, response{
				Name: temp[0],
				Type: temp[1],
			})
		}
	}

	return responses
}

// generateTestFile generates _test.go template file
func generateTestFile(packageName string, fileName string, filePath string, methodProperties []methodProperty) {
	w, err := os.Create(filePath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Print(err)
		return
	}

	err = t.Execute(w, struct {
		PackageName      string
		LowerCaseName    string
		UpperCaseName    string
		MethodProperties []methodProperty
	}{
		PackageName:      packageName,
		LowerCaseName:    ToLowerCamelCase(fileName),
		UpperCaseName:    ToUpperCamelCase(fileName),
		MethodProperties: methodProperties,
	})
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	w.Close()
}

func existFile(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
