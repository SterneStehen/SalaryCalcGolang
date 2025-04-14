package manageFile

import "os"
//import "fmt"
import "strings"
import "errors"
import "encoding/json"


func ReadMyFile(path string) ([]string, error){
	read, err := os.ReadFile(path)
	//fmt.Println(read)
	if err != nil{
		return nil, errors.New("Read File failed")
	}
	readStr := string(read)
	arrStr := strings.Split(readStr, "\n")

	//fmt.Println(arrStr)
	return arrStr, nil
}

func WriteJson(path string, data interface{}) error{
	file, err := os.Create(path)
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