package prices

import "fmt"
import "os"
//import "strconv"
import "strings"
import "SalaryCalcGolang/conver"

type TaxJob struct{
	TaxRate float64
	InputPrice []float64
	TaxIncludedPrices map[string]float64
}


func (job *TaxJob) LoadData() {
	read, err := os.ReadFile("prices/price.txt")
	//fmt.Println(read)
	if err != nil{
		fmt.Println("readerror")
		return 
	}
	readStr := string(read)
	arrStr := strings.Split(readStr, "\n")
	//fmt.Println(arrStr)
	
	job.InputPrice, err = conver.Conver(arrStr)
	if err != nil{
		fmt.Println(err)
		return
	}
	// TaxJob{
		// 	//TaxRate: taxeRate,
		// 	InputPrice: readFloat,
		// 	//InputPrice: []float64{10, 20, 30},
		// 	}
		
	}


	func (job *TaxJob) Process(){
		job.LoadData()
		result := make(map[string]string)
		for _, price := range job.InputPrice{
			resString := fmt.Sprintf("%.2f", price * (1 + job.TaxRate))
			result[fmt.Sprintf("%.2f", price)] = resString
		}
		fmt.Println(result)
	}
	
	func NewTaxIncludedPriceJob(taxeRate float64) *TaxJob {
		
		
		
		// readFloat, err := strconv.ParseFloat(readStr, 64)
		// if err != nil{
			// 	return &TaxJob{}
			// }
			
			return &TaxJob{
	TaxRate: taxeRate,
	//InputPrice: readFloat,
	//InputPrice: []float64{10, 20, 30},
	}
}