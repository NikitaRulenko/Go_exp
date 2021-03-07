package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	//inflx "github.com/influxdata/influxdb-client-go/v2"
)

func main() {

	//for {

		raw_data := getPairs()
		data := PyListToArray(raw_data)
		httpPushToInflux(data)
		time.Sleep(time.Second * 10)
		fmt.Println("POSTED--->" + time.Now().String())
	//}
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
		d2[0] = strings.TrimLeft(d2[0], `"`)
		d2[0] = strings.TrimRight(d2[0], `"`)
		result = append(result, d2)
	}

	return result
}

func httpPushToInflux(data [][]string) string {

	var (
		t_count = 0
		//f_count int = 0
	)

	for i := range data {

		if strings.HasPrefix(data[i][0], `t`) {

			body := "bitfinex,SYMBOL="+data[i][0]+" BID="+data[i][1]+",BID_SIZE="+data[i][2]+",ASK="+data[i][3]+",ASK_SIZE="+data[i][4]+",DAILY_CHANGE="+data[i][5]+",DAILY_CHANGE_RELATIVE="+data[i][6]+",LAST_PRICE="+data[i][7]+",VOLUME="+data[i][8]+",HIGH="+data[i][9]+",LOW="+data[i][10]+""
			fmt.Println(body)

			resp, err := http.Post("http://localhost:8086/write?db=bitfinex_db", "application/json", bytes.NewBufferString(body))

			if err != nil {
				log.Fatalln(err)
			} else {
				//fmt.Println(resp)
				t_count++
				defer resp.Body.Close()
			}

		} else if strings.HasPrefix(data[i][0], `f`) {

			//fmt.Println("NO")


		} else {
			fmt.Println("UNKNOWN SYMBOL")
		}
	}

	fmt.Println("Number of t pairs posted = " + strconv.Itoa(t_count))

	return ""
}