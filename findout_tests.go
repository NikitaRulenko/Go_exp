package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	raw_data := getPairs()
	data := PyListToArray(raw_data)

	fmt.Println(data[0])

	for i := range data {

		if strings.HasPrefix(data[i][0], `"t`) {

			//fmt.Println("YES")



		} else if strings.HasPrefix(data[i][0], `"f`) {

			//fmt.Println("NO")

		} else {
			fmt.Println("UNKNOWN SYMBOL")
		}
	}
}

func getPairs() string {

	url := "https://api-pub.bitfinex.com/v2/tickers?symbols=ALL"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func PyListToArray(pylist string) [][]string{

	var result [][]string

	x := strings.TrimLeft(pylist, "[[")
	x = strings.TrimRight(x, "]]")
	d1 := strings.Split(x, "],[")

	for i := range d1{
		d2 := strings.Split(d1[i], ",")
			result = append(result, d2)
	}

	return result
}

//func httpPushToInflux(data [][]string) string {
//
//	for i := range data {
//
//	}
//
//}