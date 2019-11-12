package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"runtime"
)
func main(){
	fmt.Println("GO")
	runtime.GOMAXPROCS(1)
	numComplete := 0
	stockSymbols := []string { "AMZN", "AAPL", "DIS", "NVDA", "TSLA"}
    start := time.Now()
	for _, symbol := range stockSymbols {
		go func(symbol string){
	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + "NVBWKBA9R7V7WQR9"
	fmt.Println(url)
	data := getDataFromUrl(url)
	fmt.Println(data)
	numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		time.Sleep(1 * time.Millisecond)
	}
	elasped := time.Since(start)
	fmt.Printf("execution time: %s", elasped)

}
func UnmarshalQuote(data []byte) (QuoteResponse, error) {
	var r QuoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *QuoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
type QuoteResponse struct {

	
		The01Symbol           string `json:"01. symbol"`            
		The02Open             string `json:"02. open"`              
		The03High             string `json:"03. high"`              
		The04Low              string `json:"04. low"`               
		The05Price            string `json:"05. price"`             
		The06Volume           string `json:"06. volume"`            
		The07LatestTradingDay string `json:"07. latest trading day"`
		The08PreviousClose    string `json:"08. previous close"`    
		The09Change           string `json:"09. change"`            
		The10ChangePercent    string `json:"10. change percent"`    
	
}

func getDataFromUrl(url string) string {

	var data string
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}

	return string(data)
}
