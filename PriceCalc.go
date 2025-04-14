package main

//import "fmt"
import "SalaryCalcGolang/prices"

func main(){
	//prices := []float64{ 10, 20, 32}
	taxs := []float64{ 0, 0.07, 0.1, 0.15}
	//mapTax := make(map[float64][]float64, len(taxs))

	for _, tax := range taxs{
		priceJob := prices.NewTaxIncludedPriceJob(tax)
		
		priceJob.Process()
	}
	
	//fmt.Println(mapTax)

}