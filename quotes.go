package tradierv2

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetQuotes(token string, symbols string) {

	apiUrl := "https://api.tradier.com/v1/markets/quotes?symbols=" + symbols + "&greeks=false"

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

	fmt.Println(resp.Status)
	fmt.Println(string(responseData))
}
