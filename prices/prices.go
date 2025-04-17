package prices

import "fmt"
//import "os"
//import "strconv"
//mport "strings"
import "SalaryCalcGolang/conver"
import "SalaryCalcGolang/manageFile"
//import "time"
import "errors"

type TaxJob struct{
	PathFile manageFile.FileManagePath	`json:"-"`
	TaxRate float64 					`json:"tax_rate"`
	InputPrice []float64 				`json:"input_price"`
	TaxIncludedPrices map[string]string `json:"tax_included_price"`
}


func (job *TaxJob) LoadData() (error){
	arrStr, err := job.PathFile.ReadMyFile()
	resFloat, err := conver.Conver(arrStr)
	if err != nil{
		fmt.Println(err)
		return errors.New("failed Convert")
	}
	job.InputPrice = resFloat
	// TaxJob{
		// 	//TaxRate: taxeRate,
		// 	InputPrice: readFloat,
		// 	//InputPrice: []float64{10, 20, 30},
		// 	}
	return nil
	}


	func (job *TaxJob) Process(done chan bool, errChan chan error) {
		err := job.LoadData()
		if err != nil{
			fmt.Println(err)
			errChan <- err 
			return
		}
		result := make(map[string]string)
		for _, price := range job.InputPrice{
			resString := fmt.Sprintf("%.2f", price * (1 + job.TaxRate))
			result[fmt.Sprintf("%.2f", price)] = resString
		}
		fmt.Println(result)
		job.TaxIncludedPrices = result
		//manageFile.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job )
		job.PathFile.WriteJson(job)
		//time.Sleep(time.Second *10)
		done <- true
	}
	
	func NewTaxIncludedPriceJob(fm manageFile.FileManagePath, taxeRate float64) *TaxJob {
		
		
		
		// readFloat, err := strconv.ParseFloat(readStr, 64)
		// if err != nil{
			// 	return &TaxJob{}
			// }
			
			return &TaxJob{
			PathFile: fm,
			TaxRate: taxeRate,
	//InputPrice: readFloat,
	//InputPrice: []float64{10, 20, 30},
	}
}