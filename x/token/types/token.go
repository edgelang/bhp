package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Token struct {
	Description    string `json:"description" v2:"description"`         // e.g. "OK Group Global Utility Token"
	Symbol         string `json:"symbol" v2:"symbol"`                   // e.g. "okt"
	OriginalSymbol string `json:"original_symbol" v2:"original_symbol"` // e.g. "OKT"
	WholeName      string `json:"whole_name" v2:"whole_name"`           // e.g. "OKT"
	//bhp
	OriginalTotalSupply sdk.Int `json:"original_total_supply" v2:"original_total_supply"` // e.g. 1000000000.00000000
	//bhp
	TotalSupply sdk.Int        `json:"total_supply" v2:"total_supply"` // e.g. 1000000000.00000000
	Owner       sdk.AccAddress `json:"owner" v2:"owner"`               // e.g. okchain1upyg3vl6vqaxqvzts69zpus2c027p7paw63s99
	Mintable    bool           `json:"mintable" v2:"mintable"`         // e.g. false
}

func (token Token) String() string {
	b, err := json.Marshal(token)
	if err != nil {
		return "{}"
	}
	return string(b)
}

type Tokens []Token

func (tokens Tokens) String() string {
	b, err := json.Marshal(tokens)
	if err != nil {
		return "[{}]"
	}
	return string(b)
}

type Currency struct {
	Description string  `json:"description"`
	Symbol      string  `json:"symbol"`
	TotalSupply sdk.Int `json:"total_supply"` //bhp3
}

func (currency Currency) String() string {
	b, err := json.Marshal(currency)
	if err != nil {
		return "[{}]"
	}
	return string(b)
}

//type ByDenom sdk.DecCoins
//
//func (d ByDenom) Len() int           { return len(d) }
//func (d ByDenom) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
//func (d ByDenom) Less(i, j int) bool { return d[i].Denom < d[j].Denom }

type Transfer struct {
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type TransferUnit struct {
	To    sdk.AccAddress `json:"to"`
	Coins sdk.Coins      `json:"coins"` //Bhp4
}

type CoinsInfo []CoinInfo

func (d CoinsInfo) Len() int           { return len(d) }
func (d CoinsInfo) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d CoinsInfo) Less(i, j int) bool { return d[i].Symbol < d[j].Symbol }

type AccountResponse struct {
	Address    string    `json:"address"`
	Currencies CoinsInfo `json:"currencies"`
}

func NewAccountResponse(addr string) AccountResponse {
	var accountResponse AccountResponse
	accountResponse.Address = addr
	accountResponse.Currencies = []CoinInfo{}
	return accountResponse
}

type CoinInfo struct {
	Symbol    string `json:"symbol" v2:"currency"`
	Available string `json:"available" v2:"available"`
	Locked    string `json:"locked" v2:"locked"`
}

func NewCoinInfo(symbol, available, locked string) *CoinInfo {
	return &CoinInfo{
		Symbol:    symbol,
		Available: available,
		Locked:    locked,
	}
}

//type QueryPage struct {
//	Page    int `json:"page"`
//	PerPage int `json:"per_page"`
//}

type AccountParam struct {
	Symbol string `json:"symbol"`
	Show   string `json:"show"`
	//QueryPage
}

type AccountParamV2 struct {
	Currency string `json:"currency"`
	HideZero string `json:"hide_zero"`
}

type AccCoins struct {
	Acc   sdk.AccAddress `json:"address"`
	Coins sdk.Coins      `json:"coins"` //Bhp5
}
