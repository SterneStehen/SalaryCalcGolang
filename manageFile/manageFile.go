package manageFile

import "os"
//import "fmt"
import "strings"
import "errors"
import "encoding/json"


type FileManagePath struct{
	InputPath string
	OutputPath string
}



func(fm FileManagePath) ReadMyFile() ([]string, error){
	read, err := os.ReadFile(fm.InputPath)
	//fmt.Println(read)
	if err != nil{
		return nil, errors.New("Read File failed")
	}
	readStr := string(read)
	arrStr := strings.Split(readStr, "\n")

	//fmt.Println(arrStr)
	return arrStr, nil
}

func (fm FileManagePath) WriteJson(data interface{}) error{
	file, err := os.Create(fm.OutputPath)
	if err != nil{
		return  errors.New("Create File failed")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil{
		file.Close()
		return  errors.New("Convert File failed")
	}
	file.Close()
	return nil
}

func NewPath(input string, output string) FileManagePath  {
	
	return FileManagePath{	
	InputPath: input,
	OutputPath: output,
	}
		

}