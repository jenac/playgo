package main

import (
	"fmt"
	"os"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	fmt.Print(os.Getenv("GOPATH"))
	spy, _ := quote.NewQuoteFromYahoo("spy", "2017-09-01", "2018-02-27", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 14)
	fmt.Println(rsi2)
}
