package yahooweather

import "testing"

func TestMakeQuery(t *testing.T) {
	//Test Url
	url := "https://query.yahooapis.com/v1/public/yql?q=select%20item.condition%20from%20weather.forecast%20where%20woeid%20%3D%202487889&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"
	if w := MakeQuery(url) ; w == nil {
		t.Log("Query Failed")
		t.Fail()
	}
}