package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// checkPrices periodically checks the price of BTC and sends a mail to the user that made an alert
func checkPrices(db *DB) {
	for {
		resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}
		var priceResp PriceResponse
		if err := json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
			time.Sleep(time.Minute)
			continue
		}

		currentPrice := priceResp.Bitcoin.USD
		fmt.Println("The Current Price of BTC is", currentPrice)
		alerts, err := db.GettingFromAlert()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, values := range alerts {
			if values.Price == int(currentPrice) && values.Status == "Created" {
				err = db.UpdatingFromAlert(values.AlertID, "Triggered")
				if err != nil {
					fmt.Println(err)
					continue
				}
				go sendMail(values.AlertID)
			}
		}
		time.Sleep(time.Minute)
	}
}
