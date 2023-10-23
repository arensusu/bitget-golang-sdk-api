package mix

import (
	"encoding/json"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/domain"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/model/mix/account"
)

type MixAccountGetAccountListService struct {
	client domain.RestClient
	params map[string]string
}

func NewMixAccountGetAccountListService(c domain.RestClient) *MixAccountGetAccountListService {
	return &MixAccountGetAccountListService{
		client: c,
		params: internal.NewParams(),
	}
}

func (s *MixAccountGetAccountListService) ProductType(productType string) *MixAccountGetAccountListService {
	s.params["productType"] = productType
	return s
}

type MixAccountGetAccountListReponse struct {
	common.CommonResponse
	Data []MixAccountGetAccountListData
}

type MixAccountGetAccountListData struct {
	MarginCoin        string `json:"marginCoin"`
	Locked            string `json:"locked"`
	Available         string `json:"available"`
	CrossMaxAvailable string `json:"crossMaxAvailable"`
	FixedMaxAvailable string `json:"fixedMaxAvailable"`
	MaxTransferOut    string `json:"maxTransferOut"`
	Equity            string `json:"equity"`
	UsdtEquity        string `json:"usdtEquity"`
	BtcEquity         string `json:"btcEquity"`
	CrossRiskRate     string `json:"crossRiskRate"`
	UnrealizedPL      string `json:"unrealizedPL"`
	Bonus             string `json:"bonus"`
}

func (s *MixAccountGetAccountListService) Do() (MixAccountGetAccountListReponse, error) {
	var res MixAccountGetAccountListReponse

	uri := constants.MixAccount + "/accounts"

	resp, err := s.client.DoGet(uri, s.params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil
}

type MixAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MixAccountClient) Init() *MixAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

/**
 * Get account  information
 * @param symbol
 * @param marginCoin
 * @return ResponseResult
 */
func (p *MixAccountClient) Account(symbol string, marginCoin string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol
	params["marginCoin"] = marginCoin

	uri := constants.MixAccount + "/account"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/**
 * set lever
 * @param SetLeveragerReq
 * @return ResponseResult
 */
func (p *MixAccountClient) SetLeverage(params account.SetLeveragerReq) (string, error) {
	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/setLeverage"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

/**
 * Adjustment margin
 * @param SetMarginReq
 * @return ResponseResult
 */
func (p *MixAccountClient) SetMargin(params account.SetMarginReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/setMargin"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

/**
 * Adjust margin mode
 * @param SetMarginModeReq
 * @return ResponseResult
 */
func (p *MixAccountClient) SetMarginMode(params account.SetMarginModeReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/setMarginMode"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

/**
 * Get the openable quantity
 * @param OpenCountReq
 * @return ResponseResult
 */
func (p *MixAccountClient) OpenCount(params account.OpenCountReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/open-count"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

func (p *MixAccountClient) SetPositionMode(params account.SetPositionModeReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/setPositionMode"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}
