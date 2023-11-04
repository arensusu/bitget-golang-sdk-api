package mix

import (
	"encoding/json"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/domain"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/model/mix/account"
)

type MixAccountService struct {
	BitgetRestClient domain.RestClient
}

func NewMixAccountService(c domain.RestClient) *MixAccountService {
	return &MixAccountService{BitgetRestClient: c}
}

type GetAccountsReponse struct {
	common.CommonResponse
	Data []GetAccountsData
}

type GetAccountsData struct {
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

func (s *MixAccountService) GetAccounts(productType string) (GetAccountsReponse, error) {
	var res GetAccountsReponse

	uri := constants.MixAccount + "/accounts"
	params := internal.NewParams()
	params["productType"] = productType

	resp, err := s.BitgetRestClient.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil
}

/**
 * Get account  information
 * @param symbol
 * @param marginCoin
 * @return ResponseResult
 */
func (p *MixAccountService) Account(symbol string, marginCoin string) (string, error) {

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
func (p *MixAccountService) SetLeverage(params account.SetLeveragerReq) (string, error) {
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
func (p *MixAccountService) SetMargin(params account.SetMarginReq) (string, error) {

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
func (p *MixAccountService) SetMarginMode(params account.SetMarginModeReq) (string, error) {

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
func (p *MixAccountService) OpenCount(params account.OpenCountReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/open-count"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}

func (p *MixAccountService) SetPositionMode(params account.SetPositionModeReq) (string, error) {

	postBody, jsonErr := internal.ToJson(params)

	if jsonErr != nil {
		return "", jsonErr
	}

	uri := constants.MixAccount + "/setPositionMode"

	resp, err := p.BitgetRestClient.DoPost(uri, postBody)

	return resp, err
}
