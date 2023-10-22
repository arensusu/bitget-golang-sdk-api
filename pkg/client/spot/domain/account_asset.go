package domain

import "github.com/arensusu/bitget-golang-sdk-api/internal/common"

type SpotGetAccountAssetsService interface {
	Coin(string) *SpotGetAccountAssetsService
	Do() (SpotGetAccountAssetsResponse, error)
}

type SpotGetAccountAssetsResponse struct {
	common.CommonResponse
	Data []SpotGetAccountAssetsData `json:"data"`
}

type SpotGetAccountAssetsData struct {
	CoinId    int    `json:"coinId"`
	CoinName  string `json:"coinName"`
	Available string `json:"available"`
	Frozen    string `json:"frozen"`
	Lock      string `json:"lock"`
	UTime     string `json:"uTime"`
}
