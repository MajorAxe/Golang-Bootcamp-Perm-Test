package main

import (
	"flag"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func main() {
    var value int
    var convertFrom, convertTo string

    flag.StringVar(&convertFrom, "from", "RUB", "currency FROM which we will convert")
    flag.StringVar(&convertTo, "to", "RUB", "currency TO which we will convert")
    flag.IntVar(&value, "value", 0, "amount of money")

    flag.Parse()

	resp, err := http.Get("http://api.fixer.io/latest")
	if err != nil {
		fmt.Println("Error getting exchange rates from fixer.io")
		return
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading exchange rates json")
			return
		}
		var bodyParsed map[string]interface{}
		if err := json.Unmarshal(body, &bodyParsed); err != nil {
			fmt.Println("Error decoding exchange rates json")
			return
		}

		var rates = bodyParsed["rates"].(map[string]interface{})
		var inBase, inTo float64

		if bodyParsed["base"].(string) == convertFrom {
			inBase = float64(value)
		} else {
			rateFrom := rates[convertFrom]
			if rateFrom == nil {
				fmt.Println("No data for FROM currency rate found")
				return
			}
			inBase = float64(value) / rateFrom.(float64)
		}
		if bodyParsed["base"].(string) == convertTo {
			inTo = inBase
		} else {
			rateTo := rates[convertTo]
			if rateTo == nil {
				fmt.Println("No data for TO currency rate found")
				return
			}
			inTo = inBase * rateTo.(float64)
		}
		fmt.Println("Converted", value, convertFrom, "to", inTo, convertTo)
	}
}
