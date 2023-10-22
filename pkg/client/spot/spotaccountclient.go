package spot

import (
	"encoding/json"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	common_domain "github.com/arensusu/bitget-golang-sdk-api/internal/common/domain"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/client/spot/domain"
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
	Client common_domain.RestClient
	coin   *string
}

func NewSpotGetAccountAssetsService(c common_domain.RestClient) *SpotGetAccountAssetsService {
	return &SpotGetAccountAssetsService{
		Client: c,
	}
}

func (p *SpotGetAccountAssetsService) Coin(coin string) *SpotGetAccountAssetsService {
	p.coin = &coin
	return p
}

/**
 * Obtain account assets
 * @return ResponseResult
 */
func (p *SpotGetAccountAssetsService) Do() (domain.SpotGetAccountAssetsResponse, error) {
	var res domain.SpotGetAccountAssetsResponse

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
