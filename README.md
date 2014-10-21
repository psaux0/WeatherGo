YahooWeather
--------------
YahooWeather is a **Golang** package for command line weather tool for users to get weather information in the U.S. Built with **Go** programming language, YahooWeather is able to query Yahoo Weather API with high speed. 

####Install
```
cd $GOPATH/src
go build yahooweather
go install yahooweather
```

####Sample Usage
```
package main
import (
	"yahooweather"
	"os"
	"fmt"
)

func main() {
	city, state := os.Args[1], os.Args[2]
	loc := yahooweather.BuildLocation(city,state)
	queryUrl := yahooweather.BuildUrl(loc)
	w := yahooweather.MakeQuery(queryUrl)
	if w == nil {
		fmt.Printf("Program Error")
	} else {
		fmt.Printf("Temperature: %s %s, %s, Humidity: %s",w.Temp,w.Tp,w.Weth,w.Humidity)
	}
}
```
####Sample Command and output
```
./yahooweather boston ma
Temperature: 60 F, Cloudy, Humidity: 83

#You may also get more information through edit weather.go
```
