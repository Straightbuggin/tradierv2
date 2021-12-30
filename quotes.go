package tradierv2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetQuotes(token string, symbols string, greeks string, target interface{}) error {

	apiUrl := "https://api.tradier.com/v1/markets/quotes?symbols=" + symbols + "&greeks=true" + greeks

	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlStr, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Accept", "application/json")

	resp, _ := client.Do(r)
	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return json.Unmarshal(responseData, &target)
}

type QuotesResp struct {
	Quotes struct {
		Quote []struct {
			Symbol           string  `json:"symbol"`
			Description      string  `json:"description"`
			Exch             string  `json:"exch"`
			Type             string  `json:"type"`
			Last             float64 `json:"last"`
			Change           float64 `json:"change"`
			Volume           int     `json:"volume"`
			Open             float64 `json:"open"`
			High             float64 `json:"high"`
			Low              float64 `json:"low"`
			Close            float64 `json:"close"`
			Bid              float64 `json:"bid"`
			Ask              float64 `json:"ask"`
			ChangePercentage float64 `json:"change_percentage"`
			AverageVolume    int     `json:"average_volume"`
			LastVolume       int     `json:"last_volume"`
			TradeDate        int64   `json:"trade_date"`
			Prevclose        float64 `json:"prevclose"`
			Week52High       float64 `json:"week_52_high"`
			Week52Low        float64 `json:"week_52_low"`
			Bidsize          int     `json:"bidsize"`
			Bidexch          string  `json:"bidexch"`
			BidDate          int64   `json:"bid_date"`
			Asksize          int     `json:"asksize"`
			Askexch          string  `json:"askexch"`
			AskDate          int64   `json:"ask_date"`
			RootSymbols      string  `json:"root_symbols,omitempty"`
			Underlying       string  `json:"underlying,omitempty"`
			Strike           float64 `json:"strike,omitempty"`
			OpenInterest     int     `json:"open_interest,omitempty"`
			ContractSize     int     `json:"contract_size,omitempty"`
			ExpirationDate   string  `json:"expiration_date,omitempty"`
			ExpirationType   string  `json:"expiration_type,omitempty"`
			OptionType       string  `json:"option_type,omitempty"`
			RootSymbol       string  `json:"root_symbol,omitempty"`
		} `json:"quote"`
	} `json:"quotes"`
}
