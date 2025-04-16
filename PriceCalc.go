package main

import "fmt"
import "SalaryCalcGolang/prices"
import "SalaryCalcGolang/manageFile"
//import "time"

func main(){
	//prices := []float64{ 10, 20, 32}
	taxs := []float64{ 0, 0.07, 0.1, 0.15}
	done := make([]chan bool, len(taxs))
	//mapTax := make(map[float64][]float64, len(taxs))
	
	for i, tax := range taxs{
		done[i] = make(chan bool)
		fm := manageFile.NewPath("price.txt", fmt.Sprintf("result_%.0f.json", tax*100) )
		priceJob := prices.NewTaxIncludedPriceJob(fm, tax)
		
		go priceJob.Process(done[i])
	}
	//time.Sleep(time.Second*3)

	for _, val := range done{
		<-val
	}
}