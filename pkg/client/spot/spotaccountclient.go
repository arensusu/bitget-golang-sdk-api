package spot

import (
	"encoding/json"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/model/spot/account"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init() *SpotAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

type SpotGetAccountAssetsService struct {
	Client *common.BitgetRestClient
	coin   *string
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

func (p *SpotGetAccountAssetsService) Coin(coin string) *SpotGetAccountAssetsService {
	p.coin = &coin
	return p
}

/**
 * Obtain account assets
 * @return ResponseResult
 */
func (p *SpotGetAccountAssetsService) Do() (SpotGetAccountAssetsResponse, error) {
	var res SpotGetAccountAssetsResponse

	params := internal.NewParams()
	if p.coin != nil {
		params["coin"] = *p.coin
	}

	uri := constants.SpotAccount + "/assets"
	resp, err := p.Client.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, err

}

/**
 * Obtain transfer records
 * @param coinId
 * @param fromType
 * @param limit
 * @param after
 * @param before
 * @return ResponseResult
 */
func (p *SpotAccountClient) TransferRecords(coinId string, fromType string, limit string, after string, before string) (string, error) {

	params := internal.NewParams()
	if len(coinId) > 0 {
		params["coinId"] = coinId
	}
	if len(fromType) > 0 {
		params["fromType"] = fromType
	}
	if len(limit) > 0 {
		params["limit"] = limit
	}
	if len(after) > 0 {
		params["after"] = after
	}
	if len(before) > 0 {
		params["before"] = before
	}

	uri := constants.SpotAccount + "/transferRecords"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/**
 * Get the bill flow
 * @param BillsReq
 * @return ResponseResult
 */
func (p *SpotAccountClient) Bills(params account.BillsReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.SpotAccount + "/bills"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}
