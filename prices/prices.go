package prices

import "fmt"
//import "os"
//import "strconv"
//mport "strings"
import "SalaryCalcGolang/conver"
import "SalaryCalcGolang/manageFile"

type TaxJob struct{
	TaxRate float64
	InputPrice []float64
	TaxIncludedPrices map[string]string
}


func (job *TaxJob) LoadData() {
	arrStr, err := manageFile.ReadMyFile("price.txt")
	resFloat, err := conver.Conver(arrStr)
	if err != nil{
		fmt.Println(err)
		return
	}
	job.InputPrice = resFloat
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
		job.TaxIncludedPrices = result
		manageFile.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job )
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