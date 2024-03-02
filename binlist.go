package binlist

import (
	"fmt"
	"github.com/justinmoonca/everybot-core/pkg/helpers"
	"github.com/tidwall/gjson"
	"log"
)

type BinInfo struct {
	CardType   string `json:"card_type"`
	CardScheme string `json:"card_scheme"`
	CardLevel  string `json:"card_level"`
	Country    string `json:"country"`
	Bank       string `json:"bank"`
	DataSource string `json:"data_source"`
}

const (
	API_BINLIST_IO  = "https://binlist.io/lookup/"
	API_BINLIST_NET = "https://lookup.binlist.net/"
)

func GetBinInfo(cardNumber string, proxyUrl string) (binInfo BinInfo, err error) {
	cardNumber = cardNumber[:6]
	resBody, err := helpers.HttpGet(fmt.Sprintf("%s%s", API_BINLIST_NET, cardNumber), proxyUrl)
	if err != nil {
		log.Println("request error: ", err)
		return binInfo, err
	}

	cardType := gjson.GetBytes(resBody, "type").String()
	cardScheme := gjson.GetBytes(resBody, "scheme").String()
	cardLevel := gjson.GetBytes(resBody, "brand").String()
	country := gjson.GetBytes(resBody, "country.name").String()
	bank := gjson.GetBytes(resBody, "bank.name").String()
	datasource := "binlist.net"

	if cardType == "" {
		resBody, err = helpers.HttpGet(fmt.Sprintf("%s%s", API_BINLIST_IO, cardNumber), proxyUrl)
		if err != nil {
			log.Println("request error: ", err)
			return binInfo, err
		}

		log.Println("resBody: ", string(resBody))
		cardType = gjson.GetBytes(resBody, "type").String()
		cardScheme = gjson.GetBytes(resBody, "scheme").String()
		cardLevel = gjson.GetBytes(resBody, "brand").String()
		country = gjson.GetBytes(resBody, "country.name").String()
		bank = gjson.GetBytes(resBody, "bank.name").String()
		datasource = "binlist.io"
	}

	return BinInfo{
		CardType:   cardType,
		CardScheme: cardScheme,
		CardLevel:  cardLevel,
		Country:    country,
		Bank:       bank,
		DataSource: datasource,
	}, err
}
