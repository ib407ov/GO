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

type CoinMarketCapResponseETH struct {
	Status struct {
		Timestamp    string `json:"timestamp"`
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Elapsed      int    `json:"elapsed"`
		CreditCount  int    `json:"credit_count"`
		Notice       string `json:"notice"`
	} `json:"status"`

	Data struct {
		ETH struct {
			ID                            int         `json:"id"`
			Name                          string      `json:"name"`
			Symbol                        string      `json:"symbol"`
			Slug                          string      `json:"slug"`
			NumMarketPairs                int         `json:"num_market_pairs"`
			DateAdded                     string      `json:"date_added"`
			Tags                          []string    `json:"tags"`
			MaxSupply                     interface{} `json:"max_supply"`
			CirculatingSupply             float64     `json:"circulating_supply"`
			TotalSupply                   float64     `json:"total_supply"`
			IsActive                      int         `json:"is_active"`
			InfiniteSupply                bool        `json:"infinite_supply"`
			Platform                      interface{} `json:"platform"`
			CMCRank                       int         `json:"cmc_rank"`
			IsFiat                        int         `json:"is_fiat"`
			SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
			SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
			TvlRatio                      interface{} `json:"tvl_ratio"`
			LastUpdated                   string      `json:"last_updated"`
			Quote                         struct {
				USD struct {
					Price                 float64     `json:"price"`
					Volume24h             float64     `json:"volume_24h"`
					VolumeChange24h       float64     `json:"volume_change_24h"`
					PercentChange1h       float64     `json:"percent_change_1h"`
					PercentChange24h      float64     `json:"percent_change_24h"`
					PercentChange7d       float64     `json:"percent_change_7d"`
					PercentChange30d      float64     `json:"percent_change_30d"`
					PercentChange60d      float64     `json:"percent_change_60d"`
					PercentChange90d      float64     `json:"percent_change_90d"`
					MarketCap             float64     `json:"market_cap"`
					MarketCapDominance    float64     `json:"market_cap_dominance"`
					FullyDilutedMarketCap float64     `json:"fully_diluted_market_cap"`
					Tvl                   interface{} `json:"tvl"`
					LastUpdated           string      `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"ETH"`
	} `json:"data"`
}

func main() {
	fmt.Println("Для припинення роботи натисніть CTRl+C")
	for {
		getPriceCoinBTC("1d4c8f1c-3c85-407e-83f8-2b22df3230cc")
		getPriceCoinETH("1d4c8f1c-3c85-407e-83f8-2b22df3230cc")
		currentTime := time.Now()
		fmt.Println(currentTime.Format("2006-01-02 15:04:05") + "\n")

		time.Sleep(time.Second * 10)
	}
}

func getPriceCoinBTC(apiKey string) {
	urlBTC := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=BTC&CMC_PRO_API_KEY=%s", apiKey)

	response, err := http.Get(urlBTC)
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
}
func getPriceCoinETH(apiKey string) {
	urlETH := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=ETH&CMC_PRO_API_KEY=%s", apiKey)

	response, err := http.Get(urlETH)
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

	var cmcResponseETH CoinMarketCapResponseETH
	err = json.Unmarshal(body, &cmcResponseETH)
	if err != nil {
		fmt.Println("Помилка при розборі JSON:", err)
		return
	}

	priceETH := cmcResponseETH.Data.ETH.Quote.USD.Price
	fmt.Printf("Курс біткоїна (ETH) в USD: $%.2f\n", priceETH)
}
