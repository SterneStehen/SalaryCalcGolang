package conver

//import "fmt"
import "strconv"
import "strings"
import "errors"

func Conver(arrStr []string) ([]float64, error){
	readFloat := make([]float64, len(arrStr))
	//fmt.Println(arrStr)
	for i, val := range arrStr{
		valTrim := strings.TrimSpace(val)
		res, err := strconv.ParseFloat(valTrim, 64)
		if(err != nil){
			return nil, errors.New( "Parse float failed")
		}
		readFloat[i] = res
	}
	return readFloat, nil
}