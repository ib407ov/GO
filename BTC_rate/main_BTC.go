package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CoinMarketCapResponse struct {
	Data map[string]struct {
		Quote map[string]Quote `json:"quote"`
	} `json:"data"`
}

type Quote struct {
	Price float64 `json:"price"`
}

func main() {
	fmt.Println("Для припинення роботи натисніть CTRl+C")
	for {

		apiKey := "1d4c8f1c-3c85-407e-83f8-2b22df3230cc"

		url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=BTC&CMC_PRO_API_KEY=%s", apiKey)

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Помилка при виконанні HTTP-запиту:", err)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Помилка при читанні відповіді:", err)
			return
		}

		var cmcResponse CoinMarketCapResponse
		err = json.Unmarshal(body, &cmcResponse)
		if err != nil {
			fmt.Println("Помилка при розборі JSON:", err)
			return
		}

		// Отримання тільки курсу біткоїна в USD та виведення його
		priceUSD := cmcResponse.Data["BTC"].Quote["USD"].Price
		fmt.Printf("Курс біткоїна (BTC) в USD: $%.2f\n", priceUSD)
		time.Sleep(time.Second * 5)
	}
}
