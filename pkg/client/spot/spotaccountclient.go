package spot

import (
	"encoding/json"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/domain"
)

type SpotAccountService struct {
	Client domain.RestClient
}

func NewSpotAccountService(c domain.RestClient) *SpotAccountService {
	return &SpotAccountService{Client: c}
}

type GetAssetsLiteParam struct {
	coin *string
}

type GetAssetsLiteResponse struct {
	common.CommonResponse
	Data []GetAssetsLiteData `json:"data"`
}

type GetAssetsLiteData struct {
	CoinId    int    `json:"coinId"`
	CoinName  string `json:"coinName"`
	Available string `json:"available"`
	Frozen    string `json:"frozen"`
	Lock      string `json:"lock"`
	UTime     string `json:"uTime"`
}

func (s *SpotAccountService) GetAssetsLite(param GetAssetsLiteParam) (GetAssetsLiteResponse, error) {
	var res GetAssetsLiteResponse

	params := internal.NewParams()
	if param.coin != nil {
		params["coin"] = *param.coin
	}

	uri := constants.SpotAccount + "/assets-lite"
	resp, err := s.Client.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil

}
